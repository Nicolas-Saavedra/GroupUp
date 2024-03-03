package repository

import (
	"context"

	"GroupUp/internal/model"
)

type CourseRepository interface {
	FirstById(id string, ctx context.Context) (*model.Course, error)
	Create(course *model.Course, ctx context.Context) (*model.Course, error)
	Update(course *model.Course, ctx context.Context) (*model.Course, error)
	Delete(course *model.Group, ctx context.Context) error
}

func NewCourseRepository(repository *Repository) CourseRepository {
	return &courseRepository{
		Repository: repository,
	}
}

type courseRepository struct {
	*Repository
}

func (r *courseRepository) FirstById(id string, ctx context.Context) (*model.Course, error) {
	var course model.Course
	err := r.DB(ctx).First(&course, id).Error
	return &course, err
}

func (r *courseRepository) Create(
	course *model.Course,
	ctx context.Context,
) (*model.Course, error) {
	err := r.DB(ctx).Create(course).Error
	return course, err
}

func (r *courseRepository) Update(
	course *model.Course,
	ctx context.Context,
) (*model.Course, error) {
	err := r.DB(ctx).Save(course).Error
	return course, err
}

func (r *courseRepository) Delete(course *model.Group, ctx context.Context) error {
	err := r.DB(ctx).Delete(course).Error
	return err
}
