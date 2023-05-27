package http

import (
	"github.com/anton-uvarenko/cinema/broker-service/internal/transport/grpc"
	"github.com/anton-uvarenko/cinema/broker-service/internal/transport/http/controllers"
	"net/http"
)

type Controllers struct {
	AuthController         *controllers.AuthController
	VerificationController *controllers.VerificationController
	PassRecoveryController *controllers.PassRecoveryController
	FilmsController        *controllers.FilmController
	SocialController       *controllers.SocialAuthController
	CommentsController     *controllers.CommentController
	UserDataController     *controllers.UserDataController
}

func NewControllers(clients *grpc.AllClients, httpClient *http.Client) *Controllers {
	return &Controllers{
		AuthController:         controllers.NewAuthController(clients.AuthClients.AuthClient),
		VerificationController: controllers.NewVerificationController(clients.AuthClients.VerificationClient),
		PassRecoveryController: controllers.NewPassRecoveryController(clients.AuthClients.PassRecoveryClient),
		FilmsController:        controllers.NewFilmController(httpClient),
		SocialController:       controllers.NewSocialAuthController(clients.AuthClients.SocialClient),
		CommentsController:     controllers.NewCommentController(clients.UserClients.CommentsClient),
		UserDataController:     controllers.NewUserDataController(clients.UserClients.UserDataClients),
	}
}
