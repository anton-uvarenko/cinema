package controllers

import (
	"context"
	"github.com/anton-uvarenko/cinema/authorization-service/internal/core/repo/entities"
	"github.com/anton-uvarenko/cinema/authorization-service/internal/pkg"
	"github.com/anton-uvarenko/cinema/authorization-service/protobufs/auth"
	"github.com/anton-uvarenko/cinema/authorization-service/protobufs/general"
	"github.com/sirupsen/logrus"
	"net/http"
)

type PassRecoveryController struct {
	auth.UnimplementedPassVerifyServer
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

func (c *PassRecoveryController) SendRecoveryCode(ctx context.Context, email *auth.EmailPayload) (*general.Empty, error) {
	err := c.service.SendRecoveryCode(email.Email)
	if err != nil {
		logrus.Error(err)
		fail := err.(pkg.Error)
		return nil, pkg.NewRpcError(fail.Error(), fail.Code())
	}
	return &general.Empty{}, nil
}

func (c *PassRecoveryController) VerifyRecoveryCode(ctx context.Context, code *auth.CodePayload) (*general.JwtResponse, error) {
	user, err := c.service.Verify(code.Email, int(code.Code))
	if err != nil {
		logrus.Error(err)
		fail := err.(pkg.Error)
		return nil, pkg.NewRpcError(fail.Error(), fail.Code())
	}

	token, err := pkg.NewJwt(user.Id, user.UserType, true, user.IsVerified)
	if err != nil {
		logrus.Error(err)
		return nil, pkg.NewRpcError("error creating jwt", http.StatusInternalServerError)
	}

	return &general.JwtResponse{
		Jwt: token,
	}, nil
}

func (c *PassRecoveryController) UpdatePassword(ctx context.Context, pass *auth.PasswordPayload) (*general.Empty, error) {
	err := c.service.UpdatePassword(int(pass.Id), pass.Password)
	if err != nil {
		logrus.Error(err)
		fail := err.(pkg.Error)
		return nil, pkg.NewRpcError(fail.Error(), fail.Code())
	}

	return &general.Empty{}, nil
}
