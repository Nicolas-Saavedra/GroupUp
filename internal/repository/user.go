package repository

import (
	"context"

	"gorm.io/gorm/clause"

	"GroupUp/internal/model"
)

type UserRepository interface {
	FindById(id string, ctx context.Context) (*model.User, error)
	// Fetches the user information. Will fill associations, so must
	// be called by the same user, hence, logged in
	PrivilegedFindById(id string, ctx context.Context) (*model.User, error)
	Create(user *model.User, ctx context.Context) (*model.User, error)
	Update(user *model.User, ctx context.Context) (*model.User, error)
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

func (r *userRepository) FindById(id string, ctx context.Context) (*model.User, error) {
	var user model.User
	err := r.DB(ctx).First(&user, "id = ?", id).Error
	return &user, err
}

func (r *userRepository) PrivilegedFindById(id string, ctx context.Context) (*model.User, error) {
	var user model.User
	err := r.DB(ctx).
		Preload(clause.Associations).
		First(&user, "id = ?", id).
		Error
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
