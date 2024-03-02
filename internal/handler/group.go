package handler

import (
	"GroupUp/internal/service"
)

type GroupHandler interface{}

type groupHandler struct {
	*Handler
	groupService service.GroupService
}

func NewGroupHandler(handler *Handler, groupService service.GroupService) GroupHandler {
	return &groupHandler{
		Handler:      handler,
		groupService: groupService,
	}
}
