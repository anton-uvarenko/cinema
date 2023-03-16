package controllers

import (
	"github.com/anton-uvarenko/cinema/authorization-service/internal/core/repo/entities"
	"github.com/anton-uvarenko/cinema/authorization-service/internal/pkg"
	"github.com/sirupsen/logrus"
	"net/http"
)

type PassRecoveryController struct {
	service iPassRecoveryService
}

type iPassRecoveryService interface {
	SendRecoveryCode(email string) error
	Verify(email string, code int) (*entities.User, error)
	UpdatePassword(id int, password string) error
}

func NewPassRecoveryController(service iPassRecoveryService) *PassRecoveryController {
	return &PassRecoveryController{
		service: service,
	}
}

type CodePayload struct {
	Code  int    `json:"code"`
	Email string `json:"email"`
}

type PasswordPayload struct {
	Id       int
	Password string `json:"password"`
}

type PassRecoveryResponse struct {
	Jwt string
}

func (c *PassRecoveryController) SendRecoveryCode(email string, resp *int) error {
	err := c.service.SendRecoveryCode(email)
	if err != nil {
		logrus.Error(err)
		fail := err.(pkg.Error)
		return pkg.NewRpcError(fail.Error(), fail.Code())
	}
	return nil
}

func (c *PassRecoveryController) VerifyRecoveryCode(code CodePayload, resp *PassRecoveryResponse) error {
	user, err := c.service.Verify(code.Email, code.Code)
	if err != nil {
		logrus.Error(err)
		fail := err.(pkg.Error)
		return pkg.NewRpcError(fail.Error(), fail.Code())
	}

	token, err := pkg.NewJwt(user.Id, user.UserType, true)
	if err != nil {
		logrus.Error(err)
		return pkg.NewRpcError("error creating jwt", http.StatusInternalServerError)
	}

	resp.Jwt = token
	return nil
}

func (c *PassRecoveryController) UpdatePassword(pass PasswordPayload, resp *int) error {
	err := c.service.UpdatePassword(pass.Id, pass.Password)
	if err != nil {
		logrus.Error(err)
		fail := err.(pkg.Error)
		return pkg.NewRpcError(fail.Error(), fail.Code())
	}

	return nil
}
