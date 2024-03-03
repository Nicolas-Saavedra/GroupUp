package repository

import (
	"context"

	"GroupUp/internal/model"
)

type RatingRepository interface {
	FirstById(id string, ctx context.Context) (*model.Rating, error)
	Create(course *model.Rating, ctx context.Context) (*model.Rating, error)
	Update(course *model.Rating, ctx context.Context) (*model.Rating, error)
	Delete(rating *model.Rating, ctx context.Context) error
}

func NewRatingRepository(repository *Repository) RatingRepository {
	return &ratingRepository{
		Repository: repository,
	}
}

type ratingRepository struct {
	*Repository
}

func (r *ratingRepository) FirstById(id string, ctx context.Context) (*model.Rating, error) {
	var rating model.Rating
	err := r.DB(ctx).First(&rating, id).Error
	return &rating, err
}

func (r *ratingRepository) Create(
	rating *model.Rating,
	ctx context.Context,
) (*model.Rating, error) {
	err := r.DB(ctx).Create(rating).Error
	return rating, err
}

func (r *ratingRepository) Update(
	rating *model.Rating,
	ctx context.Context,
) (*model.Rating, error) {
	err := r.DB(ctx).Save(rating).Error
	return rating, err
}

func (r *ratingRepository) Delete(rating *model.Rating, ctx context.Context) error {
	err := r.DB(ctx).Delete(rating).Error
	return err
}
