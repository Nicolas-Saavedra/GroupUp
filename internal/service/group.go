package service

import (
	"GroupUp/internal/repository"
)

type GroupService interface{}

func NewGroupService(service *Service, groupRepository repository.GroupRepository) GroupService {
	return &groupService{
		Service:         service,
		groupRepository: groupRepository,
	}
}

type groupService struct {
	*Service
	groupRepository repository.GroupRepository
}
