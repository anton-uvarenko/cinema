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
	commentRepo  iCommentRepo
	userDataRepo iCMUserDataRepo
}

const s3DefaultLink = "https://cinema-avatar-photos.s3.eu-central-1.amazonaws.com/"

type iCommentRepo interface {
	AddComment(c entities.Comment) (*entities.Comment, error)
	LikeComment(userId int, commentId int) (*entities.Comment, error)
	GetLikedCommentsIds(userId int, filmId int) ([]int, error)
	GetPrivateCommentsByFilmId(filmId int, userId int, offset int, limit int) ([]*entities.Comment, error)
	GetPublicCommentsByFilmId(filmId int, limit int, offset int, repliesAmmount int) ([]*entities.Comment, error)
	DeleteAllUserComments(userId int) error
	DeleteSingleComment(commentId int) error
}

type iCMUserDataRepo interface {
	GetUserDataByIds(ids []int) ([]entities.UserData, error)
}

func NewCommentService(repo iCommentRepo, userDataRepo iCMUserDataRepo) *CommentsService {
	return &CommentsService{
		commentRepo:  repo,
		userDataRepo: userDataRepo,
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
		logrus.Error(err)
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
		logrus.Error(err)
		return nil, pkg.NewError("error retrieving data from database", http.StatusInternalServerError)
	}

	ids := []int{}
	for _, v := range comments {
		ids = append(ids, v.UserId)
	}

	usersData, err := s.userDataRepo.GetUserDataByIds(ids)
	if err != nil {
		logrus.Error(err)
		return nil, pkg.NewError("couldn't get users data", http.StatusInternalServerError)
	}

	// assign username and avatarlink for each comment
	for i := range comments {
		for _, v := range usersData {
			if v.Userid == comments[i].UserId {
				if v.AvatarName != "" {
					comments[i].AvatarLink = s3DefaultLink + v.AvatarName
				}
				comments[i].Username = v.UserName
			}
		}
	}

	liked, err := s.commentRepo.GetLikedCommentsIds(int(payload.UserId), int(payload.FilmId))
	if err != nil {
		if err != sql.ErrNoRows {
			return nil, pkg.NewError("error requesting database data", http.StatusInternalServerError)
		}
	}

	for i := range comments {
		for _, v := range liked {
			logrus.Info(comments[i].Id, v)
			if v == comments[i].Id {
				comments[i].IsLiked = true
			}
		}
	}

	return comments, nil
}

func (s *CommentsService) LikeComment(payload *users.LikeCommentPayload) (*entities.Comment, error) {
	comment, err := s.commentRepo.LikeComment(int(payload.UserId), int(payload.CommentId))
	if err != nil {
		logrus.Error(err)
		return nil, pkg.NewError("error sending data to db", http.StatusInternalServerError)
	}

	userData, err := s.userDataRepo.GetUserDataByIds([]int{int(payload.UserId)})
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, pkg.NewError("no user data found", http.StatusNotFound)
		}
		logrus.Error(err)
		return nil, pkg.NewError("error getting user data", http.StatusInternalServerError)
	}

	if userData[0].AvatarName != "" {
		comment.AvatarLink = s3DefaultLink + userData[0].AvatarName
	}

	comment.Username = userData[0].UserName

	return comment, nil
}

func (s *CommentsService) DeleteAllUserComments(userId int) error {
	err := s.commentRepo.DeleteAllUserComments(userId)
	if err != nil {
		if err == sql.ErrNoRows {
			return pkg.NewError("no such user", http.StatusNotFound)
		}

		logrus.Error(err)
		return pkg.NewError("error deleteing comments", http.StatusInternalServerError)
	}
	return nil
}

func (s *CommentsService) DeleteSingleComment(commentId int) error {
	err := s.DeleteSingleComment(commentId)
	if err != nil {
		if err == sql.ErrNoRows {
			return pkg.NewError("no such user", http.StatusNotFound)
		}

		logrus.Error(err)
		return pkg.NewError("error deleteing comments", http.StatusInternalServerError)
	}
	return nil
}
