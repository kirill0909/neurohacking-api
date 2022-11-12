package handler

import (
	"errors"
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

func GetUserId(c *gin.Context) (int, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		newErrorResponse(c, http.StatusUnauthorized, "user id not found")
		return 0, errors.New("user id not found")
	}

	idInt, ok := id.(int)
	if !ok {
		newErrorResponse(c, http.StatusUnauthorized, "user id is of invalid type")
		return 0, errors.New("user id is of invalid type")
	}

	return idInt, nil
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

func checkEmptyValuesUserSignInInput(input models.UserSignInInput) bool {
	if len(strings.TrimSpace(input.Email)) == 0 ||
		len(strings.TrimSpace(input.Password)) == 0 {
		return false
	}
	return true
}

func checkEmptyValueUserUpdateInput(input models.UserUpdateInput) bool {
	if (input.Name != nil && len(strings.TrimSpace(*input.Name)) == 0) ||
		(input.Email != nil && len(strings.TrimSpace(*input.Email)) == 0) ||
		(input.Password != nil && len(strings.TrimSpace(*input.Password)) == 0) {
		return false
	}
	return true
}

func checkEmptyValueCategoryInput(input models.Category) bool {
	if len(strings.TrimSpace(input.Name)) == 0 {
		return false
	}
	return true
}

func checkEmptyValueCategoryUpdateInput(input models.CategoryUpdateInput) bool {
	if input.Name != nil && len(strings.TrimSpace(*input.Name)) == 0 {
		return false
	}
	return true
}

func checkEmptyValueWord(input models.Word) bool {
	if len(strings.TrimSpace(input.Name)) == 0 {
		return false
	}
	return true
}

func stringToPointer(str string) *string {
	return &str
}
