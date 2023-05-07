package controllers

import (
	"context"
	"github.com/anton-uvarenko/cinema/authorization-service/internal/pkg"
	"github.com/anton-uvarenko/cinema/authorization-service/protobufs/auth"
)

type SocialAuthController struct {
	auth.UnimplementedSocialAuthServer
	service iSocialAuthService
}

type iSocialAuthService interface {
	GoogleAuth(code string) (string, error)
}

func NewSocialAuthController(service iSocialAuthService) *SocialAuthController {
	return &SocialAuthController{
		service: service,
	}
}

func (c *SocialAuthController) GoogleAuth(ctx context.Context, payload *auth.SocialAuthPayload) (*auth.JwtResponse, error) {
	jwt, err := c.service.GoogleAuth(payload.Code)
	if err != nil {
		fail := err.(pkg.Error)
		return nil, pkg.NewRpcError(fail.Error(), fail.Code())
	}

	return &auth.JwtResponse{
		Jwt: jwt,
	}, nil
}
