package service

import (
	"GroupUp/internal/repository"
)

type CourseService interface{}

func NewCourseService(
	service *Service,
	courseRepository repository.CourseRepository,
) CourseService {
	return &courseService{
		Service:          service,
		courseRepository: courseRepository,
	}
}

type courseService struct {
	*Service
	courseRepository repository.CourseRepository
}
