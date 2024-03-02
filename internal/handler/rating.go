package handler

import (
	"GroupUp/internal/service"
)

type RatingHandler struct {
	*Handler
	ratingService service.RatingService
}

func NewRatingHandler(handler *Handler, ratingService service.RatingService) *RatingHandler {
	return &RatingHandler{
		Handler:       handler,
		ratingService: ratingService,
	}
}
