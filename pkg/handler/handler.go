package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/vpnvsk/amunet_auth_microservices"
	"github.com/vpnvsk/amunet_auth_microservices/pkg/service"
	"log/slog"
)

type Handler struct {
	log      *slog.Logger
	settings *amunet_auth_microservices.Config
	service  *service.Service
}

func NewHandler(log *slog.Logger, service *service.Service, setting *amunet_auth_microservices.Config) *Handler {
	return &Handler{
		log:      log,
		service:  service,
		settings: setting,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/signup", h.SignUp)
		auth.POST("/login", h.LogIn)
	}

	return router
}

func (h *Handler) T(c *gin.Context) {}
