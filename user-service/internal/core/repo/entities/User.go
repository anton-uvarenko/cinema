package entities

import "github.com/jmoiron/sqlx"

type User struct {
	FavouriteGenre string
	FavouriteActor int
	FavouriteFilm  int
	AvatarName     string
}

type UserRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}
