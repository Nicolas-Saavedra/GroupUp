package server

import (
	"github.com/gin-gonic/gin"

	"GroupUp/internal/handler"
	"GroupUp/internal/middleware"
	"GroupUp/pkg/helper/resp"
	"GroupUp/pkg/log"
)

func NewServerHTTP(
	logger *log.Logger,
	userHandler handler.UserHandler,
	groupHandler handler.GroupHandler,
	courseHandler handler.CourseHandler,
) *gin.Engine {
	r := gin.Default()
	r.Use(middleware.CORSMiddleware())
	v1 := r.Group("/api/v1")
	{
		v1.GET("/", func(ctx *gin.Context) {
			resp.HandleSuccess(ctx, map[string]interface{}{
				"say": "backend running successfully",
			})
		})
		v1.GET("/users/:id", userHandler.GetUserById)
	}
	return r
}
