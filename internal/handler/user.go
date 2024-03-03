package handler

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"

	"GroupUp/internal/service"
	"GroupUp/pkg/helper/resp"
)

type UserHandler interface {
	GetUserById(ctx *gin.Context)
	UpdateUser(ctx *gin.Context)
}

type userHandler struct {
	*Handler
	userService service.UserService
}

func NewUserHandler(handler *Handler, userService service.UserService) UserHandler {
	return &userHandler{
		Handler:     handler,
		userService: userService,
	}
}

func (h *userHandler) GetUserById(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		resp.HandleError(ctx, http.StatusBadRequest, 1, "no supplied id", nil)
		return
	}

	user, err := h.userService.GetUserById(id, ctx)
	h.logger.Info("GetUserByID", zap.Any("user", user))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			resp.HandleError(ctx, http.StatusNotFound, 1, "the requested user was not found", nil)
			return
		}
		resp.HandleError(ctx, http.StatusInternalServerError, 1, err.Error(), nil)
		return
	}
	resp.HandleSuccess(ctx, user)
}

func (h *userHandler) UpdateUser(ctx *gin.Context) {
	resp.HandleSuccess(ctx, nil)
}
