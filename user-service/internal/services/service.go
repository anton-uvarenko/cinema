package services

import "github.com/anton-uvarenko/cinema/user-service/internal/core/repo"

type Service struct {
	CommentsService *CommentsService
}

func NewServices(repo *repo.Repo) *Service {
	return &Service{
		CommentsService: NewCommentService(repo.CommentsRepo),
	}
}
