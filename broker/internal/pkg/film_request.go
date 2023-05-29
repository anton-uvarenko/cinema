package pkg

import (
	"encoding/json"
	"github.com/anton-uvarenko/cinema/broker-service/internal/core"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
)

func DoFilmRequest(filmCh chan<- core.Films, reqStr string) {
	client := &http.Client{}
	req, _ := http.NewRequest(http.MethodGet, reqStr, nil)
	resp, _ := client.Do(req)

	r, _ := io.ReadAll(resp.Body)

	logrus.Info(string(r))
	films := core.Films{}
	_ = json.Unmarshal(r, &films)
	filmCh <- films
}
