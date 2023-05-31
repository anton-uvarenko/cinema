package controllers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/anton-uvarenko/cinema/broker-service/internal/core"
	"github.com/anton-uvarenko/cinema/broker-service/internal/pkg"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type FilmController struct {
	client *http.Client
}

func NewFilmController(client *http.Client) *FilmController {
	return &FilmController{
		client: client,
	}
}

var ErrorCallFilms = errors.New("error calling films service")
var ErrReadingBody = errors.New("error reading response")

func (c *FilmController) RedirectRequest(w http.ResponseWriter, r *http.Request) {
	userId, _ := pkg.ParseWithId(strings.Split(r.Header.Get("Authorization"), " ")[1])

	data, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	rawUrl := "http://"
	rawUrl += os.Getenv("DNS_FILMS") + ":8000"
	rawUrl += r.URL.Path
	parsedUrl, err := url.Parse(rawUrl)
	if err != nil {
		logrus.Error(err)
	}

	logrus.Info(r.URL.Query().Get("user_id"))

	queryParams := url.Values{}

	for k, v := range r.URL.Query() {
		for _, val := range v {
			queryParams.Add(k, val)
		}
	}
	id := fmt.Sprintf("%d", userId)
	queryParams.Set("user_id", id)
	parsedUrl.RawQuery = queryParams.Encode()

	req, err := http.NewRequest(r.Method, parsedUrl.String(), bytes.NewBuffer(data))
	if err != nil {
		logrus.Errorf("request create: %v", err)
		http.Error(w, "couldn't create request", http.StatusInternalServerError)
		return
	}
	req.Header = r.Header.Clone()
	req.Form = r.Form
	if req.FormValue("user_id") != "" {
		req.Form.Set("user_id", id)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		logrus.Errorf("send requset to films: %v", err)
		http.Error(w, ErrorCallFilms.Error(), http.StatusServiceUnavailable)
		return
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, ErrReadingBody.Error(), http.StatusInternalServerError)
		return
	}

	if resp.StatusCode >= 400 {
		http.Error(w, string(respBody), resp.StatusCode)
		return
	}

	w.WriteHeader(resp.StatusCode)
	w.Write(respBody)
}

func (c *FilmController) ManePageRequest(w http.ResponseWriter, r *http.Request) {
	nFilms := make(chan core.Films)
	iFilms := make(chan core.Films)
	lFilms := make(chan core.Films)

	go pkg.DoFilmRequest(nFilms, "http://"+os.Getenv("DNS_FILMS")+":8000"+"/films/?release_date_after=2023&release&date_before=2023&order_by=imdb&page_size=3&page=1")
	go pkg.DoFilmRequest(iFilms, "http://"+os.Getenv("DNS_FILMS")+":8000"+"/films/?order_by=imdb_rating&page_size=3&page=1")
	go pkg.DoFilmRequest(lFilms, "http://"+os.Getenv("DNS_FILMS")+":8000"+"/films/?order_by=rating&page_size=3&page=1")

	resp := core.ManePageResponse{
		NewFilms:         <-nFilms,
		IBMRatingFilms:   <-iFilms,
		LocalRatingFilms: <-lFilms,
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(&resp)
	if err != nil {
		logrus.Error(err)
		http.Error(w, "error encoding response", http.StatusInternalServerError)
		return
	}
}
