package service

import (
	"context"
	"errors"
	"fmt"

	"gorm.io/gorm"

	"GroupUp/internal/model"
	"GroupUp/internal/repository"
)

type UserService interface {
	FetchCurrentUserOrCreate(id string, ctx context.Context) (*model.User, error)
	FetchUserById(id string, ctx context.Context) (*model.User, error)
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
	ctx context.Context,
) (*model.User, error) {
	user, err := s.userRepository.PrivilegedFindById(id, ctx)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return s.createUserFromAuthService(id, ctx)
		}
		return nil, fmt.Errorf("could not find user with said id")
	}
	return user, nil
}

func (s *userService) FetchUserById(
	id string,
	ctx context.Context,
) (*model.User, error) {
	user, err := s.userRepository.FindById(id, ctx)
	if err != nil {
		return nil, fmt.Errorf("could not find user with said id")
	}
	return user, nil
}

func (s *userService) createUserFromAuthService(
	id string,
	ctx context.Context,
) (*model.User, error) {
	userLoginInfo, err := s.authRepository.GetLoginInformation(id, ctx)
	if err != nil {
		return nil, fmt.Errorf("could not grab login information from user %s", id)
	}
	user := &model.User{}
	user.ID = id
	user.Name = userLoginInfo.DisplayName
	user.Email = userLoginInfo.Email
	_, err = s.userRepository.Create(user, ctx)
	if err != nil {
		return nil, fmt.Errorf("could not create new user %s on the database", id)
	}
	return user, nil
}
