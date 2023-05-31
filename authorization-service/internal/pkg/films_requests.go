package pkg

import (
	"bytes"
	"fmt"
	"github.com/anton-uvarenko/cinema/authorization-service/internal/core/repo/entities"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"net/url"
	"os"
)

func SendBasicPlaylists(id int, userType entities.UserType) error {
	rawUrl := "http://"
	rawUrl += os.Getenv("DNS_FILMS") + ":8000"
	rawUrl += "/playlists/"

	parsedUrl, err := url.Parse(rawUrl)
	if err != nil {
		logrus.Error(err)
		return NewError("can't parse playlist service url", http.StatusInternalServerError)
	}

	formVals := url.Values{}
	formVals.Set("title", "Watch later")
	formVals.Set("user_id", fmt.Sprintf("%d", id))
	formVals.Set("user_type", string(userType))

	req, err := http.NewRequest(http.MethodPost, parsedUrl.String(), bytes.NewBuffer([]byte(formVals.Encode())))
	if err != nil {
		logrus.Error(err)
		return NewError("can't create request", http.StatusInternalServerError)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	clientHttp := http.Client{}
	resp, err := clientHttp.Do(req)
	if err != nil {
		logrus.Error(err)
		return NewError("error calling service", http.StatusInternalServerError)
	}

	body, _ := io.ReadAll(resp.Body)

	if resp.StatusCode >= 400 {
		logrus.Info(string(body))
		return NewError("error calling playlist service", http.StatusInternalServerError)
	}

	formVals.Set("title", "Watched")

	req, err = http.NewRequest(http.MethodPost, parsedUrl.String(), bytes.NewBuffer([]byte(formVals.Encode())))
	if err != nil {
		logrus.Error(err)
		return NewError("can't create request", http.StatusInternalServerError)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		return NewError("error calling service", http.StatusInternalServerError)
	}

	if resp.StatusCode >= 400 {
		return NewError("error calling playlist service", http.StatusInternalServerError)
	}

	return nil
}
