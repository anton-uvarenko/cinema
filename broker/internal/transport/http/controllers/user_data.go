package controllers

import (
	"context"
	"encoding/json"
	"github.com/anton-uvarenko/cinema/broker-service/internal/pkg"
	"github.com/anton-uvarenko/cinema/broker-service/protobufs/users"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"strconv"
	"strings"
)

type UserDataController struct {
	client users.UserDataUploaderClient
}

func NewUserDataController(client users.UserDataUploaderClient) *UserDataController {
	return &UserDataController{
		client: client,
	}
}

func (c *UserDataController) AddData(w http.ResponseWriter, r *http.Request) {
	id, _ := pkg.ParseWithId(strings.Split(r.Header.Get("Authorization"), " ")[1])
	userName, _ := pkg.ParseWithUsername(strings.Split(r.Header.Get("Authorization"), " ")[1])

	favActor, _ := strconv.Atoi(r.FormValue("fav-actor"))
	favGenre, _ := strconv.Atoi(r.FormValue("fav-genre"))
	favFilm, _ := strconv.Atoi(r.FormValue("fav-film"))

	img, head, err := r.FormFile("img")
	if err != nil && err != http.ErrMissingFile {
		logrus.Error(err)
		http.Error(w, "couldn't retrieve image", http.StatusInternalServerError)
		return
	}

	if err == http.ErrMissingFile {
		stream, err := c.client.UploadData(context.Background())
		if err != nil {
			logrus.Error(err)
			http.Error(w, "couldn't establish connection with userdata service", http.StatusServiceUnavailable)
			return
		}

		err = stream.Send(&users.UploadDataPayload{
			Image:    nil,
			FavActor: int32(favActor),
			FavGenre: int32(favGenre),
			FavFilm:  int32(favFilm),
			UserId:   int32(id),
			//Username:
		})
		if err != nil {
			logrus.Error(err)
			http.Error(w, "error sending image", http.StatusInternalServerError)
			return
		}

		_, err = stream.CloseAndRecv()
		if err != nil {
			fail := pkg.CustToPkgError(err.Error())
			http.Error(w, fail.Error(), fail.Code())
			return
		}

		return
	}

	// image has to be less than 500 kB
	if head.Size > (1 << 22) {
		http.Error(w, "too large image", http.StatusRequestEntityTooLarge)
		return
	}

	stream, err := c.client.UploadData(context.Background())
	if err != nil {
		logrus.Error(err)
		http.Error(w, "couldn't establish connection with userdata service", http.StatusServiceUnavailable)
		return
	}

	buffer := make([]byte, 10240)
	for {
		n, err := img.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			logrus.Error(err.Error())
			http.Error(w, "error sending image", http.StatusInternalServerError)
			return
		}

		err = stream.Send(&users.UploadDataPayload{
			Image:    buffer[:n],
			FavActor: int32(favActor),
			FavGenre: int32(favGenre),
			FavFilm:  int32(favFilm),
			UserId:   int32(id),
			Username: userName,
		})
		if err != nil {
			logrus.Error(err)
			http.Error(w, "error sending image", http.StatusInternalServerError)
			return
		}
	}

	logrus.Info("exited from infinite loop")
	_, err = stream.CloseAndRecv()
	if err != nil {
		fail := pkg.CustToPkgError(err.Error())
		http.Error(w, fail.Error(), fail.Code())
		return
	}
}

func (c *UserDataController) GetData(w http.ResponseWriter, r *http.Request) {
	payload := users.GetDataPayload{}
	userId, err := strconv.Atoi(r.URL.Query().Get("user_id"))
	if err != nil {
		http.Error(w, "wrong format of user_id specified", http.StatusBadRequest)
		return
	}

	payload.UserId = int32(userId)

	resp, err := c.client.GetData(context.Background(), &payload)
	if err != nil {
		fail := pkg.CustToPkgError(err.Error())
		http.Error(w, fail.Error(), fail.Code())
		return
	}

	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, "error encoding data", http.StatusInternalServerError)
		return
	}
}

func (c *UserDataController) DeleteImage(w http.ResponseWriter, r *http.Request) {
	payload := users.DeleteImagePayload{}
	id, _ := pkg.ParseWithId(strings.Split(r.Header.Get("Authorization"), " ")[1])
	payload.UserId = int32(id)

	_, err := c.client.DeleteImage(context.Background(), &payload)
	if err != nil {
		fail := pkg.CustToPkgError(err.Error())
		http.Error(w, fail.Error(), fail.Code())
		return
	}
}
