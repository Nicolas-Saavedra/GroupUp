package handler

import (
	"GroupUp/internal/service"
)

type CourseHandler interface {
}

type courseHandler struct {
	*Handler
	courseService service.CourseService
}

func NewCourseHandler(handler *Handler, courseService service.CourseService) CourseHandler {
	return &courseHandler{
		Handler:       handler,
		courseService: courseService,
	}
}
