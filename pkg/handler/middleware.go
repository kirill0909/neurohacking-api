package handler

import (
	"github.com/kirill0909/neurohacking-api/models"
	"strings"
)

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
