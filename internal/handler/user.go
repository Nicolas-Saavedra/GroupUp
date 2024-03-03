package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"

	"GroupUp/internal/service"
	"GroupUp/pkg/context"
	"GroupUp/pkg/helper/resp"
)

type UserHandler interface {
	GetUserWithID(ctx *gin.Context)
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

func (h *userHandler) GetUserWithID(ctx *gin.Context) {
	id := ctx.Param("id")
	token := context.GrabToken(ctx)
	if token != nil && id == token.UID {
		h.getCurrentUser(id, ctx)
	} else {
		h.getForeignUserWithID(id, ctx)
	}
}

func (h *userHandler) getForeignUserWithID(id string, ctx *gin.Context) {
	user, err := h.userService.FetchUserById(id, ctx)
	h.logger.Info("GetCurrentUser", zap.Any("user", user))
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

func (h *userHandler) getCurrentUser(id string, ctx *gin.Context) {
	user, err := h.userService.FetchCurrentUserOrCreate(id, ctx)
	h.logger.Info("GetCurrentUser", zap.Any("user", user))
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
