package controllers

import (
	"bytes"
	"errors"
	"fmt"
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
	req.Form.Set("user_id", id)

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
