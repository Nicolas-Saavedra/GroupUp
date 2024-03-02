//go:build wireinject
// +build wireinject

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/spf13/viper"

	"GroupUp/internal/handler"
	"GroupUp/internal/repository"
	"GroupUp/internal/server"
	"GroupUp/internal/service"
	"GroupUp/pkg/log"
)

var ServerSet = wire.NewSet(server.NewServerHTTP)

var RepositorySet = wire.NewSet(
	repository.NewDb,
	repository.NewRepository,
	repository.NewUserRepository,
	repository.NewGroupRepository,
	repository.NewCourseRepository,
	repository.NewRatingRepository,
)

var ServiceSet = wire.NewSet(
	service.NewService,
	service.NewUserService,
	service.NewGroupService,
	service.NewCourseService,
	service.NewRatingService,
)

var HandlerSet = wire.NewSet(
	handler.NewHandler,
	handler.NewUserHandler,
	handler.NewGroupHandler,
	handler.NewCourseHandler,
	handler.NewRatingHandler,
)

func newApp(*viper.Viper, *log.Logger) (*gin.Engine, func(), error) {
	panic(wire.Build(
		ServerSet,
		RepositorySet,
		ServiceSet,
		HandlerSet,
	))
}
