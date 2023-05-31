package services

import (
	"database/sql"
	"github.com/anton-uvarenko/cinema/authorization-service/internal/core/repo/entities"
	"github.com/anton-uvarenko/cinema/authorization-service/internal/pkg"
	"github.com/anton-uvarenko/cinema/authorization-service/protobufs/auth"
	"github.com/sirupsen/logrus"
	"net/http"
)

type AdminService struct {
	repo iAdminUserRepo
}

type iAdminUserRepo interface {
	UpdateUserType(id int, userType entities.UserType) (*entities.User, error)
	DeleteUser(id int) error
	GetAllUsers() ([]entities.User, error)
}

func NewAdminService(repo iAdminUserRepo) *AdminService {
	return &AdminService{
		repo: repo,
	}
}

func (s *AdminService) ChangeUserType(payload *auth.UpdateUserTypePayload) (string, error) {
	tp := entities.UserType(payload.UserType)
	if tp == "" {
		return "", pkg.NewError("not existing user type", http.StatusBadRequest)
	}

	user, err := s.repo.UpdateUserType(int(payload.UserId), tp)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", pkg.NewError("no user in db", http.StatusNotFound)
		}

		logrus.Error(err)
		return "", pkg.NewError("error sending data into db", http.StatusInternalServerError)
	}

	jwt, err := pkg.NewJwt(user.Id, user.Email, user.Username, user.UserType, false, user.IsVerified)
	if err != nil {
		return "", pkg.NewError("error creating jwt", http.StatusInternalServerError)
	}

	return jwt, nil
}

func (s *AdminService) DeleteUser(payload *auth.DeleteUserPayload) error {
	err := s.repo.DeleteUser(int(payload.UserId))
	if err != nil {
		if err == sql.ErrNoRows {
			return pkg.NewError("no user with such id", http.StatusNotFound)
		}

		logrus.Error(err)
		return pkg.NewError("error deleting user", http.StatusInternalServerError)
	}

	return nil
}

func (s *AdminService) GetAllUsers() ([]entities.User, error) {
	users, err := s.repo.GetAllUsers()
	if err != nil {
		logrus.Error(err)
		return nil, pkg.NewError("error retrieving users from db", http.StatusInternalServerError)
	}

	return users, nil
}
