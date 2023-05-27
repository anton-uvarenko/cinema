package grpc

import (
	"github.com/anton-uvarenko/cinema/user-service/internal/services"
	"github.com/anton-uvarenko/cinema/user-service/internal/transport/grpc/controllers"
	"github.com/anton-uvarenko/cinema/user-service/protobufs/users"
	"google.golang.org/grpc"
)

func SetUpServerControllers(server *grpc.Server, services *services.Service) {
	users.RegisterCommentsServer(server, controllers.NewCommentController(services.CommentsService))
	users.RegisterUserDataUploaderServer(server, controllers.NewUserDataController(services.UserDataService))
}
