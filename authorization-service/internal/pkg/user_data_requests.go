package pkg

import (
	"context"
	"github.com/anton-uvarenko/cinema/authorization-service/protobufs/users"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net/http"
	"os"
)

type userClient struct {
	client         users.UserDataUploaderClient
	commentsClient users.CommentsClient
}

var client = &userClient{}

func InitGrpcConnection() error {
	logrus.Info("connecting to user_data service: ", os.Getenv("DNS_COMMENTS"))
	conn, err := grpc.Dial(
		os.Getenv("DNS_COMMENTS")+":5000",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)

	logrus.Info("connected, err: ", err)
	if err != nil {
		logrus.Error(err)
		panic(err)
		return err
	}

	logrus.Info("connecting users-data service")

	client.client = users.NewUserDataUploaderClient(conn)
	client.commentsClient = users.NewCommentsClient(conn)
	return nil
}

func SendBasicUserDataRequests(userId int, username string) error {
	if client.client == nil {
		return NewError("no connection is established", http.StatusInternalServerError)
	}

	stream, err := client.client.UploadData(context.Background())
	if err != nil {
		logrus.Error(err)
		return NewError("can't establish steaming connection", http.StatusServiceUnavailable)
	}
	payload := users.UploadDataPayload{
		Image:    nil,
		FavActor: 0,
		FavGenre: 0,
		FavFilm:  0,
		UserId:   int32(userId),
		Username: username,
	}
	err = stream.Send(&payload)
	if err != nil {
		logrus.Error(err)
		return NewError("error sending data", http.StatusInternalServerError)
	}

	_, err = stream.CloseAndRecv()
	if err != nil {
		logrus.Error(err.Error())
		return NewError("error closing streaming connection", http.StatusInternalServerError)
	}

	return nil
}

func SendDeleteUserDataRequst(userId int) error {
	dlPayload := users.DeleteDataPayload{
		UserId: int32(userId),
	}

	_, err := client.client.DeleteData(context.Background(), &dlPayload)
	if err != nil {
		return NewError("error deleting user data", http.StatusInternalServerError)
	}

	dp := users.DeleteAllUserCommentsPayload{UserId: int32(userId)}
	_, err = client.commentsClient.DeleteAllUserComments(context.Background(), &dp)
	if err != nil {
		return NewError("error deleteing comments", http.StatusInternalServerError)
	}

	return nil
}
