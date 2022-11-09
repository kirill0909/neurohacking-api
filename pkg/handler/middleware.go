package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/kirill0909/neurohacking-api/models"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if len(strings.TrimSpace(header)) == 0 {
		newErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")
	if ok := validateAuthHeader(headerParts); !ok {
		newErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
		return
	}

	userId, err := h.services.ParseToken(headerParts[1])
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set(userCtx, userId)
}

func validateAuthHeader(headerParts []string) bool {
	if len(headerParts) != 2 ||
		headerParts[0] != "Bearer" ||
		len(strings.TrimSpace(headerParts[1])) == 0 {
		return false
	}
	return true
}

func checkEmptyValuesUser(user models.User) bool {
	if len(strings.TrimSpace(user.Name)) == 0 ||
		len(strings.TrimSpace(user.Email)) == 0 ||
		len(strings.TrimSpace(user.Password)) == 0 {
		return false
	}
	return true
}

func checkEmptyValuesSignInInput(input signInInput) bool {
	if len(strings.TrimSpace(input.Email)) == 0 ||
		len(strings.TrimSpace(input.Password)) == 0 {
		return false
	}
	return true
}
