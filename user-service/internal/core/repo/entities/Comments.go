package entities

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"time"
)

const DefaultExecTimeout = 5 * time.Second

type CommentType string

const (
	PrivateComment CommentType = "private"
	PublicComment  CommentType = "public"
)

type Comment struct {
	Id          int         `json:"id" db:"id"`
	Text        string      `json:"text" db:"text"`
	UserId      int         `json:"userId" db:"user_id"`
	CreatedAt   time.Time   `json:"createdAt" db:"created_at"`
	UpdatedAt   time.Time   `json:"updatedAt" db:"updated_at"`
	CommentType CommentType `json:"commentType" db:"comment_type"`
	ReplyTo     *int        `json:"replyTo" db:"reply_to"`
	Rating      int         `json:"rating" db:"rating"`
	WasEdited   bool        `json:"wasEdited" db:"was_edited"`
	FilmId      int         `json:"filmId" db:"film_id"`
}

type CommentRepo struct {
	db *sqlx.DB
}

func NewCommentRepo(db *sqlx.DB) *CommentRepo {
	return &CommentRepo{
		db: db,
	}
}

func (r *CommentRepo) AddComment(c Comment) (*Comment, error) {
	ctx, cancel := context.WithTimeout(context.Background(), DefaultExecTimeout)
	defer cancel()

	query := `
	INSERT INTO comments (
	                      text,
	                      user_id,
	                      reply_to,
	                      rating,
	                      comment_type,
	                      created_at,
	                      updated_at,
	                      was_edited,
	                      film_id
	                     )
	VALUES (
	        :text,
			:user_id,
			:reply_to,
			:rating,
			:comment_type,
			:created_at,
			:updated_at,
			:was_edited,
			:film_id
	)
`
	rows, err := r.db.NamedQueryContext(ctx, query, c)
	if err != nil {
		logrus.Error("error adding comment: ", err)
		return nil, err
	}

	rowComment := &Comment{}
	for rows.Next() {
		err = rows.StructScan(rowComment)
		if err != nil {
			logrus.Error("error scunning comment: ", err)
			return nil, err
		}
	}

	return rowComment, nil
}

func (r *CommentRepo) UpdateComment(c Comment) error {
	ctx, cancel := context.WithTimeout(context.Background(), DefaultExecTimeout)
	defer cancel()

	c.WasEdited = true
	c.UpdatedAt = time.Now()

	query := `
	UPDATE comments
	SET
	     text = :text,
	     user_id = :user_id,
	     reply_to = :reply_to,
	     rating = :rating,
	     updated_at = :updated_at,
	     was_edited = :was_edited
	WHERE id = :id
`

	_, err := r.db.NamedExecContext(ctx, query, c)
	if err != nil {
		logrus.Error("error updating comment: ", err)
		return err
	}

	return nil
}

func (r *CommentRepo) GetPublicCommentsByFilmId(filmId int) ([]*Comment, error) {
	ctx, cancel := context.WithTimeout(context.Background(), DefaultExecTimeout)
	defer cancel()

	query := `
	SELECT
	    id,
	    text,
	    user_id,
	    created_at,
	    updated_at,
	    reply_to,
	    rating,
	    was_edited,
	    film_id
	FROM comments
	WHERE film_id = $1 and commentType = $2
`

	cms := []*Comment{}
	err := r.db.GetContext(ctx, cms, query, filmId, PublicComment)
	if err != nil {
		return nil, err
	}

	return cms, nil
}

func (r *CommentRepo) GetPrivateCommentsByFilmId(filmId int) ([]*Comment, error) {
	ctx, cancel := context.WithTimeout(context.Background(), DefaultExecTimeout)
	defer cancel()

	query := `
	SELECT
	    id,
	    text,
	    user_id,
	    created_at,
	    updated_at,
	    film_id
	FROM comments
	WHERE film_id = $1 and commentType = $2
`

	cms := []*Comment{}
	err := r.db.GetContext(ctx, cms, query, filmId, PrivateComment)
	if err != nil {
		return nil, err
	}

	return cms, nil
}
