package repo

import (
	"github.com/anton-uvarenko/cinema/user-service/internal/core/repo/entities"
	"github.com/jmoiron/sqlx"
)

type Repo struct {
	UserRepo     *entities.UserRepo
	CommentsRepo *entities.CommentRepo
}

func NewRepo(db *sqlx.DB) *Repo {
	return &Repo{
		UserRepo:     entities.NewUserRepo(db),
		CommentsRepo: entities.NewCommentRepo(db),
	}
}
