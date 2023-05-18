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
}

func NewControllers(clients grpc.AuthClients, httpClient *http.Client) *Controllers {
	return &Controllers{
		AuthController:         controllers.NewAuthController(clients.AuthClient),
		VerificationController: controllers.NewVerificationController(clients.VerificationClient),
		PassRecoveryController: controllers.NewPassRecoveryController(clients.PassRecoveryClient),
		FilmsController:        controllers.NewFilmController(httpClient),
		SocialController:       controllers.NewSocialAuthController(clients.SocialClient),
		CommentsController:     controllers.NewCommentController(clients.CommentsClient),
	}
}
