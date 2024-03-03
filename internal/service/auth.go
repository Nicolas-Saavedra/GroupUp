package service

import (
	"context"

	"firebase.google.com/go/v4/auth"

	"GroupUp/internal/repository"
)

type AuthService interface {
	VerifyToken(authHeader string, ctx context.Context) (*auth.Token, error)
}

func NewAuthService(service *Service, authRepository repository.AuthRepository) AuthService {
	return &authService{
		Service:        service,
		authRepository: authRepository,
	}
}

type authService struct {
	*Service
	authRepository repository.AuthRepository
}

func (s *authService) VerifyToken(authHeader string, ctx context.Context) (*auth.Token, error) {
	return s.authRepository.VerifyHeader(authHeader, ctx)
}
