package pkg

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
	"net/url"
	"os"
)

func SendBasicPlaylists(id int) error {
	rawUrl := "http://"
	rawUrl += os.Getenv("DNS_FILMS") + ":8000"
	rawUrl += "/playlists"

	parsedUrl, err := url.Parse(rawUrl)
	if err != nil {
		logrus.Error(err)
		return NewError("can't parse playlist service url", http.StatusInternalServerError)
	}

	req, err := http.NewRequest(http.MethodPost, parsedUrl.String(), nil)
	if err != nil {
		logrus.Error(err)
		return NewError("can't create request", http.StatusInternalServerError)
	}

	req.Form.Add("title", "Watch later")
	req.Form.Add("user_id", fmt.Sprintf("%d", id))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return NewError("error calling service", http.StatusInternalServerError)
	}

	if resp.StatusCode >= 400 {
		return NewError("error calling playlist service", http.StatusInternalServerError)
	}

	req, err = http.NewRequest(http.MethodPost, parsedUrl.String(), nil)
	if err != nil {
		logrus.Error(err)
		return NewError("can't create request", http.StatusInternalServerError)
	}

	req.Form.Add("title", "Watched")
	req.Form.Add("user_id", fmt.Sprintf("%d", id))

	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		return NewError("error calling service", http.StatusInternalServerError)
	}

	if resp.StatusCode >= 400 {
		return NewError("error calling playlist service", http.StatusInternalServerError)
	}

	return nil
}
