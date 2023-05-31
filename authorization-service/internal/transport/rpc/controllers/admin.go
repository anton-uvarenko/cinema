package controllers

import (
	"context"
	"github.com/anton-uvarenko/cinema/authorization-service/internal/core/repo/entities"
	"github.com/anton-uvarenko/cinema/authorization-service/internal/pkg"
	"github.com/anton-uvarenko/cinema/authorization-service/protobufs/auth"
	"github.com/anton-uvarenko/cinema/authorization-service/protobufs/general"
)

type AdminController struct {
	auth.UnimplementedAdminHandlerServer
	service iAdminService
}

func NewAdminController(service iAdminService) *AdminController {
	return &AdminController{
		service: service,
	}
}

type iAdminService interface {
	ChangeUserType(payload *auth.UpdateUserTypePayload) (string, error)
	DeleteUser(payload *auth.DeleteUserPayload) error
	GetAllUsers() ([]entities.User, error)
}

func (c *AdminController) UpdateUserType(ctx context.Context, payload *auth.UpdateUserTypePayload) (*general.JwtResponse, error) {
	jwt, err := c.service.ChangeUserType(payload)
	if err != nil {
		fail := err.(pkg.Error)
		return nil, pkg.NewRpcError(fail.Error(), fail.Code())
	}
	return &general.JwtResponse{
		Jwt: jwt,
	}, nil
}

func (c *AdminController) DelteUser(ctx context.Context, payload *auth.DeleteUserPayload) (*general.Empty, error) {
	err := c.service.DeleteUser(payload)
	if err != nil {
		fail := err.(pkg.Error)
		return nil, pkg.NewRpcError(fail.Error(), fail.Code())
	}

	err = pkg.SendDeleteUserDataRequst(int(payload.UserId))
	if err != nil {
		fail := err.(pkg.Error)
		return nil, pkg.NewRpcError(fail.Error(), fail.Code())
	}

	return &general.Empty{}, nil
}

func (c *AdminController) GetAllUsers(ctx context.Context, payload *general.Empty) (*auth.GetAllUsersResponse, error) {
	response := &auth.GetAllUsersResponse{
		Users: []*auth.UserResponse{},
	}
	users, err := c.service.GetAllUsers()
	if err != nil {
		fail := err.(pkg.Error)
		return nil, pkg.NewRpcError(fail.Error(), fail.Code())
	}

	for _, v := range users {
		u := &auth.UserResponse{
			UserId:   int32(v.Id),
			Email:    v.Email,
			UserType: string(v.UserType),
			Username: v.Username,
		}
		response.Users = append(response.Users, u)
	}

	return response, nil
}
