package entities

import (
	"context"
	"github.com/jmoiron/sqlx"
)

type FilmRateRepo struct {
	db *sqlx.DB
}

func NewFilmRateRepo(db *sqlx.DB) *FilmRateRepo {
	return &FilmRateRepo{
		db: db,
	}
}

func (r *FilmRateRepo) RateFilm(userId int, filmId, rating int) error {
	ctx, cancel := context.WithTimeout(context.Background(), DefaultExecTimeout)
	defer cancel()

	query := `
	INSERT INTO film_rate (
	                       users_id, 
	                       film_id, 
	                       rating
	                       )
	VALUES (
	        $1,
	        $2,
	        $3
	) ON CONFLICT DO 
	    UPDATE SET
	               rating = $3
	WHERE users_id = $1 AND film_id = $2
`

	_, err := r.db.ExecContext(ctx, query, userId, filmId, rating)
	if err != nil {
		return err
	}
	return nil
}
