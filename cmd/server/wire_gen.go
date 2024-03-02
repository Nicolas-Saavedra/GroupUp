// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"GroupUp/internal/handler"
	"GroupUp/internal/repository"
	"GroupUp/internal/server"
	"GroupUp/internal/service"
	"GroupUp/pkg/log"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

// Injectors from wire.go:

func newApp(viperViper *viper.Viper, logger *log.Logger) (*gin.Engine, func(), error) {
	handlerHandler := handler.NewHandler(logger)
	serviceService := service.NewService(logger)
	db := repository.NewDb(viperViper, logger)
	repositoryRepository := repository.NewRepository(logger, db)
	userRepository := repository.NewUserRepository(repositoryRepository)
	userService := service.NewUserService(serviceService, userRepository)
	userHandler := handler.NewUserHandler(handlerHandler, userService)
	groupRepository := repository.NewGroupRepository(repositoryRepository)
	groupService := service.NewGroupService(serviceService, groupRepository)
	groupHandler := handler.NewGroupHandler(handlerHandler, groupService)
	courseRepository := repository.NewCourseRepository(repositoryRepository)
	courseService := service.NewCourseService(serviceService, courseRepository)
	courseHandler := handler.NewCourseHandler(handlerHandler, courseService)
	engine := server.NewServerHTTP(logger, userHandler, groupHandler, courseHandler)
	return engine, func() {
	}, nil
}

// wire.go:

var ServerSet = wire.NewSet(server.NewServerHTTP)

var RepositorySet = wire.NewSet(repository.NewDb, repository.NewRepository, repository.NewUserRepository, repository.NewGroupRepository, repository.NewCourseRepository, repository.NewRatingRepository)

var ServiceSet = wire.NewSet(service.NewService, service.NewUserService, service.NewGroupService, service.NewCourseService, service.NewRatingService)

var HandlerSet = wire.NewSet(handler.NewHandler, handler.NewUserHandler, handler.NewGroupHandler, handler.NewCourseHandler, handler.NewRatingHandler)
