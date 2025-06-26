package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/vpnvsk/amunet_auth_microservices/internal/models"
	"net/http"
)

func (h *Handler) SignUp(c *gin.Context) {
	var input models.UserCreate
	authMethod := "email"
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error(), h.log)
		return
	}
	accessToken, refreshToken, err := h.service.SignUp(input.Email, input.Username, authMethod, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error(), h.log)
		return
	}

	c.SetCookie("amunet_refresh_token", refreshToken, int(h.settings.RefreshTTL), "/", "localhost",
		false, true)

	c.JSON(http.StatusOK, map[string]interface{}{
		"accessToken": accessToken,
	})
}

func (h *Handler) LogIn(c *gin.Context) {
	var input models.UserLogIn
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusUnprocessableEntity, err.Error(), h.log)
		return
	}
	accessToken, refreshToken, err := h.service.LogIn(input.Email, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error(), h.log)
		return
	}
	c.SetCookie("amunet_refresh_token", refreshToken, int(h.settings.RefreshTTL), "/", "localhost",
		false, true)

	c.JSON(http.StatusOK, map[string]interface{}{
		"accessToken": accessToken,
	})
}
