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
	        $1,
			$2,
			$3,
			$4,
			$5,
			$6,
			$7,
			$8,
			$9
	)
	RETURNING id
`
	var id int
	err := r.db.QueryRowContext(
		ctx,
		query,
		c.Text,
		c.UserId,
		c.ReplyTo,
		c.Rating,
		c.CommentType,
		c.CreatedAt,
		c.UpdatedAt,
		c.WasEdited,
		c.FilmId,
	).Scan(&id)
	if err != nil {
		logrus.Error("error adding comment: ", err)
		return nil, err
	}

	ctx, cancel = context.WithTimeout(context.Background(), DefaultExecTimeout)
	defer cancel()
	rowComment := &Comment{}

	getQuery := `
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
	WHERE id=$1
`

	err = r.db.GetContext(ctx, rowComment, getQuery, id)
	if err != nil {
		logrus.Error(err)
		return nil, err
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

func (r *CommentRepo) GetPublicCommentsByFilmId(filmId int, limit int, offset int, repliesAmmount int) ([]*Comment, error) {
	ctx, cancel := context.WithTimeout(context.Background(), DefaultExecTimeout)
	defer cancel()

	// $1 is film id
	// $2 is comment type
	// $3 is amount of replies
	// $4 is amount of root comments
	// $5 us offset - page
	query := `
	SELECT 
	    c1.id,
	    c1.text,
	    c1.user_id,
	    c1.created_at,
	    c1.updated_at,
	    c1.reply_to,
	    c1.rating,
	    c1.was_edited,
	    c1.film_id
	FROM comments c1
	WHERE c1.film_id = $1
	AND c1.comment_type = $2
	AND c1.reply_to IS NULL
	ORDER BY created_at
	LIKE $4
	OFFSET $5
	
	UNION ALL 
	    
	SELECT
	    c1.id,
	    c1.text,
	    c1.user_id,
	    c1.created_at,
	    c1.updated_at,
	    c1.reply_to,
	    c1.rating,
	    c1.was_edited,
	    c1.film_id
	FROM comments c1
	JOIN (
	    SELECT c.id, c.reply_to, row_number() over (partition by c.reply_to order by created_at) as row_num
	    FROM comments c
	    WHERE c.reply_to IS NOT NULL
		AND c.film_id = $1
	    AND c.comment_type = $2
	) c2
	ON c1.id = c2.reply_to
	WHERE c2.row_num <= $3
	AND c1.film_id = $1
	AND c1.comment_type = $2
	ORDER BY created_at
	OFFSET $5
`

	cms := []*Comment{}
	err := r.db.GetContext(ctx, cms, query, filmId, PublicComment, repliesAmmount, limit, offset)
	if err != nil {
		return nil, err
	}

	return cms, nil
}

func (r *CommentRepo) GetPrivateCommentsByFilmId(filmId int, userId int, offset int, limit int) ([]*Comment, error) {
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
	WHERE film_id = $1 and comment_type = $2 and user_id = $3
	LIMIT $4
	OFFSET $5
`

	cms := []*Comment{}
	err := r.db.GetContext(ctx, cms, query, filmId, PrivateComment, userId, limit, offset)
	if err != nil {
		return nil, err
	}

	return cms, nil
}
