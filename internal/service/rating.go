package service

import (
	"GroupUp/internal/repository"
)

type RatingService interface{}

func NewRatingService(
	service *Service,
	ratingRepository repository.RatingRepository,
) RatingService {
	return &ratingService{
		Service:          service,
		ratingRepository: ratingRepository,
	}
}

type ratingService struct {
	*Service
	ratingRepository repository.RatingRepository
}
