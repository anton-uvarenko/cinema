package grpc

import (
	"github.com/anton-uvarenko/cinema/broker-service/protobufs/auth"
	"github.com/anton-uvarenko/cinema/broker-service/protobufs/users"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"os"
	"time"
)

const DefaultRetriesAmount = 5
const DefaultTryTimeout = 10 * time.Second

type AuthClients struct {
	AuthClient         auth.AuthClient
	VerificationClient auth.VerificationClient
	PassRecoveryClient auth.PassVerifyClient
	SocialClient       auth.SocialAuthClient
}

func connectAuthServer() AuthClients {
	logrus.Info("auth dns name is: ", os.Getenv("DNS_AUTH"))
	conn, err := grpc.Dial(
		os.Getenv("DNS_AUTH")+":5000",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		logrus.Error(err)
		panic(err)
	}

	logrus.Info(conn.GetState())

	logrus.Info("connected to auth")

	clients := AuthClients{
		AuthClient:         auth.NewAuthClient(conn),
		VerificationClient: auth.NewVerificationClient(conn),
		PassRecoveryClient: auth.NewPassVerifyClient(conn),
		SocialClient:       auth.NewSocialAuthClient(conn),
	}

	return clients
}

type UserCilents struct {
	CommentsClient users.CommentsClient
}

func connectUsersServer() UserCilents {
	logrus.Info("auth dns name is: ", os.Getenv("DNS_COMMENTS"))
	conn, err := grpc.Dial(
		os.Getenv("DNS_COMMENTS")+":5000",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		logrus.Error(err)
		panic(err)
	}

	logrus.Info(conn.GetState())

	logrus.Info("connected to auth")

	clients := UserCilents{
		CommentsClient: users.NewCommentsClient(conn),
	}

	return clients
}

type AllClients struct {
	AuthClients AuthClients
	UserClients UserCilents
}

func ConnectAllClients() *AllClients {
	return &AllClients{
		AuthClients: connectAuthServer(),
		UserClients: connectUsersServer(),
	}
}
