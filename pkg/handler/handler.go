package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/vpnvsk/amunet_auth_microservices/pkg/service"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	api := router.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/sign-up", h.T)
		}
	}
	return router
}

func (h *Handler) T(c *gin.Context) {}
