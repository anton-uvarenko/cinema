package services

import (
	"github.com/anton-uvarenko/cinema/user-service/internal/core/repo/entities"
	"github.com/anton-uvarenko/cinema/user-service/internal/pkg"
	"github.com/sirupsen/logrus"
)

type CommentsService struct {
	commentRepo iCommentRepo
}

type iCommentRepo interface {
	AddComment(c entities.Comment) (*entities.Comment, error)
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
