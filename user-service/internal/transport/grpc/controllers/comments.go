package controllers

import (
	"context"
	"github.com/anton-uvarenko/cinema/user-service/internal/core/repo/entities"
	"github.com/anton-uvarenko/cinema/user-service/internal/pkg"
	"github.com/anton-uvarenko/cinema/user-service/protobufs/users"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type CommentController struct {
	users.UnimplementedCommentsServer
	service iCommentService
}

func NewCommentController(service iCommentService) *CommentController {
	return &CommentController{
		service: service,
	}
}

type iCommentService interface {
	AddComment(comment *entities.Comment) (*entities.Comment, error)
}

func (c *CommentController) AddComment(ctx context.Context, paylaod *users.CommentPayload) (*users.Comment, error) {
	replyTo := int(paylaod.ReplyTo)
	cmType := entities.CommentType(paylaod.CommentType)

	if cmType != entities.PrivateComment && cmType != entities.PublicComment {
		return nil, pkg.NewRpcError("wrong comment type", 400)
	}

	comment := &entities.Comment{
		UserId:      int(paylaod.UserId),
		FilmId:      int(paylaod.FilmId),
		Text:        string(paylaod.Text),
		ReplyTo:     &replyTo,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		WasEdited:   false,
		Rating:      0,
		CommentType: cmType,
	}

	cm, err := c.service.AddComment(comment)
	if err != nil {
		fail := err.(pkg.Error)
		return nil, pkg.NewRpcError(fail.Error(), fail.Code())
	}

	response := &users.Comment{
		Id:          int32(cm.Id),
		UserId:      int32(cm.UserId),
		FilmId:      int32(cm.FilmId),
		Text:        cm.Text,
		ReplyTo:     int32(*cm.ReplyTo),
		CommentType: string(cm.CommentType),
		WasEdited:   cm.WasEdited,
		CreatedAt:   timestamppb.New(cm.CreatedAt),
		UpdatedAt:   timestamppb.New(cm.UpdatedAt),
	}

	return response, nil
}
