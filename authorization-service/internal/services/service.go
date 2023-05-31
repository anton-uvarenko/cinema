package services

import "github.com/anton-uvarenko/cinema/authorization-service/internal/core/repo"

type Service struct {
	AuthService             *AuthService
	VerificationService     *VerificationService
	PasswordRecoveryService *PassRecoverService
	SocialAuthService       *SocialAuthService
	AdminService            *AdminService
}

func NewService(repo *repo.Repo) *Service {
	return &Service{
		AuthService:             NewAuthService(repo.UserRepo),
		VerificationService:     NewVerificationService(repo.UserRepo),
		PasswordRecoveryService: NewPasswordRecoveryService(repo.UserRepo),
		SocialAuthService:       NewSocialAuthService(repo.UserRepo),
		AdminService:            NewAdminService(repo.UserRepo),
	}
}
