package controllers

import (
	"bytes"
	"errors"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
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
	data, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	req, err := http.NewRequest(r.Method, "http://"+os.Getenv("DNS_FILMS")+":8000/"+r.URL.Path, bytes.NewBuffer(data))
	if err != nil {
		logrus.Errorf("request create: %v", err)
		http.Error(w, "couldn't create request", http.StatusInternalServerError)
		return
	}
	req.Header = r.Header.Clone()

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
