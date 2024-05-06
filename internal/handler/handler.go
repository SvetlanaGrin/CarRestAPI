package handler

import (
	"carRestAPI/internal/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	//gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	api := router.Group("/api")
	{
		api.POST("/", h.Create)
		api.GET("/", h.GetAll)
		api.PUT("/:regNum", h.Update)
		api.DELETE("/:regNum", h.Delete)
	}

	return router
}
