package controllers

import (
	"context"
	"github.com/anton-uvarenko/cinema/user-service/internal/core/repo/entities"
	"github.com/anton-uvarenko/cinema/user-service/internal/pkg"
	"github.com/anton-uvarenko/cinema/user-service/protobufs/general"
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
	LikeComment(payload *users.LikeCommentPayload) (*entities.Comment, error)
	GetPrivateComments(payload *users.GetPrivateCommentsPayload) ([]*entities.Comment, error)
	GetPublicComments(payload *users.GetPublicCommentsPayload) ([]*entities.Comment, error)
	DeleteAllUserComments(userId int) error
	DeleteSingleComment(commentId int) error
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
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		WasEdited:   false,
		Rating:      0,
		CommentType: cmType,
	}

	if replyTo != 0 {
		comment.ReplyTo = &replyTo
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
		CommentType: string(cm.CommentType),
		WasEdited:   cm.WasEdited,
		CreatedAt:   timestamppb.New(cm.CreatedAt),
		UpdatedAt:   timestamppb.New(cm.UpdatedAt),
	}

	if cm.ReplyTo != nil {
		response.ReplyTo = int32(*cm.ReplyTo)
	}

	return response, nil
}

func (c *CommentController) GetPrivateComments(ctx context.Context, payload *users.GetPrivateCommentsPayload) (*users.CommentsResponse, error) {
	comments, err := c.service.GetPrivateComments(payload)
	if err != nil {
		fail := err.(pkg.Error)
		return nil, pkg.NewRpcError(fail.Error(), fail.Code())
	}

	respComs := make([]*users.Comment, len(comments))
	for i, v := range comments {
		respComs[i] = &users.Comment{
			Id:          int32(v.Id),
			FilmId:      int32(v.FilmId),
			Text:        v.Text,
			UserId:      int32(v.UserId),
			CommentType: string(v.CommentType),
			WasEdited:   v.WasEdited,
			CreatedAt:   timestamppb.New(v.CreatedAt),
			UpdatedAt:   timestamppb.New(v.UpdatedAt),
		}
	}

	response := &users.CommentsResponse{
		Comments: respComs,
	}

	return response, nil
}

func (c *CommentController) GetPublicComments(ctx context.Context, payload *users.GetPublicCommentsPayload) (*users.CommentsResponse, error) {
	comments, err := c.service.GetPublicComments(payload)
	if err != nil {
		fail := err.(pkg.Error)
		return nil, pkg.NewRpcError(fail.Error(), fail.Code())
	}

	resp := users.CommentsResponse{
		Comments: []*users.Comment{},
	}

	for _, v := range comments {
		reply := int32(0)
		if v.ReplyTo != nil {
			reply = int32(*v.ReplyTo)
		}
		com := users.Comment{
			Id:          int32(v.Id),
			FilmId:      int32(v.FilmId),
			Text:        v.Text,
			UserId:      int32(v.UserId),
			ReplyTo:     reply,
			Username:    v.Username,
			AvatarLink:  v.AvatarLink,
			CommentType: string(v.CommentType),
			WasEdited:   v.WasEdited,
			CreatedAt:   timestamppb.New(v.CreatedAt),
			UpdatedAt:   timestamppb.New(v.UpdatedAt),
			IsLiked:     v.IsLiked,
			LikesAmount: int32(v.LikesAmount),
		}
		resp.Comments = append(resp.Comments, &com)
	}

	return &resp, nil
}

func (c *CommentController) LikeComment(ctx context.Context, payload *users.LikeCommentPayload) (*users.Comment, error) {
	comment, err := c.service.LikeComment(payload)
	if err != nil {
		fail := err.(pkg.Error)
		return nil, pkg.NewRpcError(fail.Error(), fail.Code())
	}

	resp := &users.Comment{
		Id:          int32(comment.Id),
		FilmId:      int32(comment.FilmId),
		Text:        comment.Text,
		UserId:      int32(comment.UserId),
		ReplyTo:     0,
		CommentType: string(comment.CommentType),
		WasEdited:   comment.WasEdited,
		CreatedAt:   timestamppb.New(comment.CreatedAt),
		UpdatedAt:   timestamppb.New(comment.UpdatedAt),
		AvatarLink:  comment.AvatarLink,
		Username:    comment.Username,
		IsLiked:     comment.IsLiked,
	}

	return resp, nil
}

func (c *CommentController) DeleteAllUserComments(ctx context.Context, payload *users.DeleteAllUserCommentsPayload) (*general.Empty, error) {
	err := c.service.DeleteAllUserComments(int(payload.UserId))
	if err != nil {
		fail := err.(pkg.Error)
		return nil, pkg.NewRpcError(fail.Error(), fail.Code())
	}

	return &general.Empty{}, nil
}

func (c *CommentController) DeleteSingleComment(ctx context.Context, payload *users.DeleteSingleCommentPayload) (*general.Empty, error) {
	err := c.service.DeleteSingleComment(int(payload.CommentId))
	if err != nil {
		fail := err.(pkg.Error)
		return nil, pkg.NewRpcError(fail.Error(), fail.Code())
	}

	return &general.Empty{}, nil
}
