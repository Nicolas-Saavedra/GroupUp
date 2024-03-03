package repository

import (
	"context"

	"GroupUp/internal/model"
)

type UserRepository interface {
	FirstById(id string, ctx context.Context) (*model.User, error)
	Create(course *model.User, ctx context.Context) (*model.User, error)
	Update(course *model.User, ctx context.Context) (*model.User, error)
	Delete(user *model.User, ctx context.Context) error
}

type userRepository struct {
	*Repository
}

func NewUserRepository(repository *Repository) UserRepository {
	return &userRepository{
		Repository: repository,
	}
}

func (r *userRepository) FirstById(id string, ctx context.Context) (*model.User, error) {
	var user model.User
	err := r.DB(ctx).First(&user, "id = ?", id).Error
	return &user, err
}

func (r *userRepository) Create(user *model.User, ctx context.Context) (*model.User, error) {
	err := r.DB(ctx).Create(user).Error
	return user, err
}

func (r *userRepository) Update(user *model.User, ctx context.Context) (*model.User, error) {
	err := r.DB(ctx).Save(user).Error
	return user, err
}

func (r *userRepository) Delete(user *model.User, ctx context.Context) error {
	err := r.DB(ctx).Delete(user).Error
	return err
}
