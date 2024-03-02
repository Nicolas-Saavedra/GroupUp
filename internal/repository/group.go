package repository

import (
	"context"

	"GroupUp/internal/model"
)

type GroupRepository interface {
	FirstById(id int64, ctx context.Context) (*model.Group, error)
	Create(course *model.Group, ctx context.Context) (*model.Group, error)
	Update(course *model.Group, ctx context.Context) (*model.Group, error)
	Delete(group *model.Group, ctx context.Context) error
}

func NewGroupRepository(repository *Repository) GroupRepository {
	return &groupRepository{
		Repository: repository,
	}
}

type groupRepository struct {
	*Repository
}

func (r *groupRepository) FirstById(id int64, ctx context.Context) (*model.Group, error) {
	var group model.Group
	err := r.DB(ctx).First(&group, id).Error
	return &group, err
}

func (r *groupRepository) Create(group *model.Group, ctx context.Context) (*model.Group, error) {
	err := r.DB(ctx).Create(group).Error
	return group, err
}

func (r *groupRepository) Update(group *model.Group, ctx context.Context) (*model.Group, error) {
	err := r.DB(ctx).Save(group).Error
	return group, err
}

func (r *groupRepository) Delete(group *model.Group, ctx context.Context) error {
	err := r.DB(ctx).Delete(group).Error
	return err
}
