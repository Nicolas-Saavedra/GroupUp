package service

import (
	"context"
	"errors"
	"fmt"

	"firebase.google.com/go/v4/auth"
	"gorm.io/gorm"

	"GroupUp/internal/model"
	"GroupUp/internal/repository"
)

type UserService interface {
	FetchCurrentUserOrCreate(id string, authHeader string, ctx context.Context) (*model.User, error)
}

type userService struct {
	*Service
	userRepository repository.UserRepository
	authRepository repository.AuthRepository
}

func NewUserService(
	service *Service,
	userRepository repository.UserRepository,
	authRepository repository.AuthRepository,
) UserService {
	return &userService{
		Service:        service,
		userRepository: userRepository,
		authRepository: authRepository,
	}
}

func (s *userService) FetchCurrentUserOrCreate(
	id string,
	authHeader string,
	ctx context.Context,
) (*model.User, error) {
	token, err := s.authRepository.VerifyHeader(authHeader, ctx)
	if err != nil {
		return nil, fmt.Errorf("requested auth header is invalid")
	}
	user, err := s.userRepository.PrivilegedFindById(token.UID, ctx)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return s.createUserFromToken(token, ctx)
		}
		return nil, fmt.Errorf("could not find user with said id")
	}
	return user, nil
}

func (s *userService) createUserFromToken(
	token *auth.Token,
	ctx context.Context,
) (*model.User, error) {
	userLoginInfo, err := s.authRepository.GetLoginInformation(token.UID, ctx)
	if err != nil {
		return nil, fmt.Errorf("could not grab login information from user %s", token.UID)
	}
	user := &model.User{}
	user.ID = token.UID
	user.Name = userLoginInfo.DisplayName
	user.Email = userLoginInfo.Email
	_, err = s.userRepository.Create(user, ctx)
	if err != nil {
		return nil, fmt.Errorf("could not create new user %s on the database", token.UID)
	}
	return user, nil
}
