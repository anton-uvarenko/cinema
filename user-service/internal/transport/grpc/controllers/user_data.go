package controllers

import (
	"bytes"
	"context"
	"github.com/anton-uvarenko/cinema/user-service/internal/core/repo/entities"
	"github.com/anton-uvarenko/cinema/user-service/internal/pkg"
	"github.com/anton-uvarenko/cinema/user-service/protobufs/general"
	"github.com/anton-uvarenko/cinema/user-service/protobufs/users"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"time"
)

type UserDataController struct {
	users.UnimplementedUserDataUploaderServer
	service iUserDataService
}

const s3DefaultLink = "https://cinema-avatar-photos.s3.eu-central-1.amazonaws.com/"

type iUserDataService interface {
	AddData(img io.Reader, userData *entities.UserData) (*entities.UserData, error)
	GetData(userId int) (*entities.UserData, error)
}

func NewUserDataController(service iUserDataService) *UserDataController {
	return &UserDataController{
		service: service,
	}
}

func (c *UserDataController) UploadData(stream users.UserDataUploader_UploadDataServer) error {
	payload := &users.UploadDataPayload{
		Image: make([]byte, 0),
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*25)

	errorQuit := make(chan error)

	// get data from broker
	go func(errQuit chan<- error) {
		for {
			chunk, err := stream.Recv()
			if err == io.EOF {
				cancel()
				return
			}
			if err != nil {
				logrus.Info("EOF")
				errorQuit <- pkg.NewError(err.Error(), http.StatusInternalServerError)
				return
			}

			payload.Image = append(payload.Image, chunk.Image...)
			if payload.FavActor == 0 {
				payload.UserId = chunk.UserId
				payload.FavFilm = chunk.FavFilm
				payload.FavActor = chunk.FavActor
				payload.FavGenre = chunk.FavGenre
				payload.Username = chunk.Username
			}
		}
	}(errorQuit)

	// waiting the end of data retrieving or error
	select {
	case <-ctx.Done():
		if len(payload.Image) == 0 {
			uploadData := &entities.UserData{
				Userid:         int(payload.UserId),
				UserName:       payload.Username,
				FavouriteGenre: int(payload.FavGenre),
				FavouriteActor: int(payload.FavActor),
				FavouriteFilm:  int(payload.FavFilm),
				AvatarName:     "",
			}

			_, err := c.service.AddData(nil, uploadData)
			if err != nil {
				return err
			}
			stream.SendAndClose(&general.Empty{})
			break
		}

		uploadData := &entities.UserData{
			Userid:         int(payload.UserId),
			FavouriteGenre: int(payload.FavGenre),
			FavouriteActor: int(payload.FavActor),
			FavouriteFilm:  int(payload.FavFilm),
			AvatarName:     uuid.New().String(),
		}

		_, err := c.service.AddData(bytes.NewBuffer(payload.Image), uploadData)
		if err != nil {
			return err
		}

		stream.SendAndClose(&general.Empty{})

	case err := <-errorQuit:
		fail := err.(pkg.Error)
		return pkg.NewRpcError(fail.Error(), fail.Code())
	}
	return nil
}

func (c *UserDataController) GetData(ctx context.Context, payload *users.GetDataPayload) (*users.GetDataResponse, error) {
	data, err := c.service.GetData(int(payload.UserId))
	if err != nil {
		fail := err.(pkg.Error)
		return nil, pkg.NewRpcError(fail.Error(), fail.Code())
	}

	if data == nil {
		return nil, pkg.NewRpcError("not found", http.StatusNotFound)
	}

	imageLink := s3DefaultLink + data.AvatarName
	if data.AvatarName == "" {
		imageLink = ""
	}

	resp := &users.GetDataResponse{
		ImageLink: imageLink,
		FavActor:  int32(data.FavouriteActor),
		FavGenre:  int32(data.FavouriteGenre),
		FavFilm:   int32(data.FavouriteFilm),
		UserId:    int32(data.Userid),
	}

	return resp, nil
}
