package services

import (
	"database/sql"
	"encoding/json"
	"github.com/anton-uvarenko/cinema/authorization-service/internal/core/config"
	"github.com/anton-uvarenko/cinema/authorization-service/internal/core/repo/entities"
	"github.com/anton-uvarenko/cinema/authorization-service/internal/pkg"
	"github.com/sirupsen/logrus"
	"net/http"
)

type SocialAuthService struct {
	userRepo iSocialUserRepo
}

func NewSocialAuthService(repo iUserRepo) *SocialAuthService {
	return &SocialAuthService{
		userRepo: repo,
	}
}

type iSocialUserRepo interface {
	GetUserByEmail(email string) (*entities.User, error)
}

type socialResponse struct {
	Email string `json:"email"`
}

func (s *SocialAuthService) GoogleAuth(code string) (string, error) {
	response, err := http.Get(config.GoogleUrl + code)
	if err != nil {
		return "", pkg.NewError("couldn't retrieve data from google api", http.StatusServiceUnavailable)
	}

	email := &socialResponse{}
	err = json.NewDecoder(response.Body).Decode(email)
	if err != nil {
		return "", pkg.NewError("couldn't retrieve data from google api", http.StatusServiceUnavailable)
	}

	user, err := s.userRepo.GetUserByEmail(email.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", pkg.NewError("user is not registered", http.StatusNotFound)
		}
		logrus.Error(err.Error())
		return "", pkg.NewError("db error", http.StatusInternalServerError)
	}

	jwt, err := pkg.NewJwt(user.Id, user.Email, user.Username, user.UserType, false, user.IsVerified)
	if err != nil {
		return "", pkg.NewError("couldn't create jwt", http.StatusInternalServerError)
	}

	return jwt, nil
}

func (s *SocialAuthService) FacebookAuth(code string) (string, error) {
	response, err := http.Get(config.FacebookUrl + code)
	if err != nil {
		return "", pkg.NewError("couldn't retrieve data from google api", http.StatusServiceUnavailable)
	}

	email := &socialResponse{}
	err = json.NewDecoder(response.Body).Decode(email)
	if err != nil {
		return "", pkg.NewError("couldn't retrieve data from google api", http.StatusServiceUnavailable)
	}

	user, err := s.userRepo.GetUserByEmail(email.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", pkg.NewError("user is not registered", http.StatusNotFound)
		}
		logrus.Error(err.Error())
		return "", pkg.NewError("db error", http.StatusInternalServerError)
	}

	jwt, err := pkg.NewJwt(user.Id, user.Email, user.Username, user.UserType, false, user.IsVerified)
	if err != nil {
		return "", pkg.NewError("couldn't create jwt", http.StatusInternalServerError)
	}

	return jwt, nil
}
