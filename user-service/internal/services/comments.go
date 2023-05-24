package services

import (
	"database/sql"
	"github.com/anton-uvarenko/cinema/user-service/internal/core/repo/entities"
	"github.com/anton-uvarenko/cinema/user-service/internal/pkg"
	"github.com/anton-uvarenko/cinema/user-service/protobufs/users"
	"github.com/sirupsen/logrus"
	"net/http"
)

type CommentsService struct {
	commentRepo iCommentRepo
}

type iCommentRepo interface {
	AddComment(c entities.Comment) (*entities.Comment, error)
	GetPrivateCommentsByFilmId(filmId int, userId int, offset int, limit int) ([]*entities.Comment, error)
	GetPublicCommentsByFilmId(filmId int, limit int, offset int, repliesAmmount int) ([]*entities.Comment, error)
}

func NewCommentService(repo iCommentRepo) *CommentsService {
	return &CommentsService{
		commentRepo: repo,
	}
}

func (s *CommentsService) AddComment(comment *entities.Comment) (*entities.Comment, error) {
	cm, err := s.commentRepo.AddComment(*comment)
	if err != nil {
		logrus.Error(err)
		return nil, pkg.NewError("couldn't add comment", 500)
	}

	return cm, nil
}

func (s *CommentsService) GetPrivateComments(payload *users.GetPrivateCommentsPayload) ([]*entities.Comment, error) {
	limit := int(payload.Amount)
	comments, err := s.commentRepo.GetPrivateCommentsByFilmId(int(payload.FilmId), int(payload.UserId), int(payload.Page-1)*limit, limit)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, pkg.NewError("no comments found", http.StatusNotFound)
		}
		return nil, pkg.NewError("error calling database", http.StatusInternalServerError)
	}

	return comments, nil
}

func (s *CommentsService) GetPublicComments(payload *users.GetPublicCommentsPayload) ([]*entities.Comment, error) {
	comments, err := s.commentRepo.GetPublicCommentsByFilmId(
		int(payload.FilmId),
		int(payload.Amount),
		int(payload.Page-1)*int(payload.Amount),
		int(payload.ResponsesAmount),
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, pkg.NewError("no comments found", http.StatusNotFound)
		}
		return nil, pkg.NewError("error calling database", http.StatusInternalServerError)
	}

	return comments, nil
}
