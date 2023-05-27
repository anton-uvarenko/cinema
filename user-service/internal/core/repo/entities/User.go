package entities

import (
	"context"
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type UserData struct {
	Userid         int    `db:"user_id"`
	FavouriteGenre int    `db:"favourite_genre"`
	FavouriteActor int    `db:"favourite_actor"`
	FavouriteFilm  int    `db:"favourite_film"`
	AvatarName     string `db:"avatar_name"`
}

type UserRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (r *UserRepo) AddData(u *UserData) (*UserData, error) {
	ctx, cancel := context.WithTimeout(context.Background(), DefaultExecTimeout)
	defer cancel()

	query := `
	INSERT INTO user_data (
	                       user_id,
	                       favourite_genre,
	                       favourite_actor,
	                       favourite_film,
	                       avatar_name
	)
	VALUES (
	        $1,
			$2,
			$3,
			$4,
			$5
	)
	ON CONFLICT (user_id) DO UPDATE 
	SET 
		favourite_genre = $2,
		favourite_actor = $3,
		favourite_film = $4,
		avatar_name = $5
	RETURNING user_id
`

	var user_id int
	err := r.db.QueryRowContext(
		ctx,
		query,
		u.Userid,
		u.FavouriteGenre,
		u.FavouriteActor,
		u.FavouriteFilm,
		u.AvatarName,
	).Scan(&user_id)
	if err != nil {
		logrus.Error(err)
		return nil, errors.New("error inserting into db: " + err.Error())
	}

	ctx, cancel = context.WithTimeout(context.Background(), DefaultExecTimeout)
	defer cancel()
	resp := &UserData{}

	query = `
	SELECT 
	    user_id,
		favourite_genre,
		favourite_actor,
		favourite_film,
		avatar_name
	FROM user_data
	WHERE user_id = $1
`
	err = r.db.GetContext(ctx, resp, query, user_id)
	if err != nil {
		logrus.Error(err)
		return nil, errors.New("error retrieving data from db: " + err.Error())
	}

	return resp, nil
}

func (r *UserRepo) GetUserDataById(id int) (*UserData, error) {
	ctx, cancel := context.WithTimeout(context.Background(), DefaultExecTimeout)
	defer cancel()

	query := `
	SELECT 
	    user_id,
		favourite_genre,
		favourite_actor,
		favourite_film,
		avatar_name
	FROM user_data
	WHERE user_id = $1
`
	resp := &UserData{}
	err := r.db.GetContext(ctx, resp, query, id)
	if err != nil {
		logrus.Error(err)
		return nil, errors.New("error retrieving data from db: " + err.Error())
	}

	return resp, nil
}
