package controllers

import (
	"context"
	"github.com/anton-uvarenko/cinema/authorization-service/internal/pkg"
	"github.com/anton-uvarenko/cinema/authorization-service/protobufs/auth"
	"github.com/anton-uvarenko/cinema/authorization-service/protobufs/general"
	"github.com/sirupsen/logrus"
)

type VerificationController struct {
	auth.UnimplementedVerificationServer
	verificationService iVerificationService
}

type iVerificationService interface {
	SendCode(id int) error
	VerifyCode(code int, id int) error
}

func NewVerificationController(service iVerificationService) *VerificationController {
	return &VerificationController{
		verificationService: service,
	}
}

func (c *VerificationController) SendCode(ctx context.Context, id *auth.IdPayload) (*general.Empty, error) {
	logrus.Info("id is ", id.Id)
	err := c.verificationService.SendCode(int(id.Id))
	if err != nil {
		logrus.Error(err)
		fail := err.(pkg.Error)
		return nil, pkg.NewRpcError(fail.Error(), fail.Code())
	}

	return &general.Empty{}, nil
}

func (c *VerificationController) VerifyCode(ctx context.Context, payload *auth.VerificationPayload) (*general.Empty, error) {
	err := c.verificationService.VerifyCode(int(payload.Code), int(payload.Id))
	if err != nil {
		logrus.Error(err)
		fail := err.(pkg.Error)
		return nil, pkg.NewRpcError(fail.Error(), fail.Code())
	}

	return &general.Empty{}, nil
}
