package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/kirill0909/neurohacking-api/models"
	"github.com/sirupsen/logrus"
	"net/http"
)

func (h *Handler) signUp(c *gin.Context) {
	var input models.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	if ok := checkEmptyValuesUser(input); !ok {
		newErrorResponse(c, http.StatusBadRequest, "invalid input value")
		return
	}

	id, err := h.services.User.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"id": id,
	})

}

type signInInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) signIn(c *gin.Context) {
	var input signInInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	if ok := checkEmptyValuesSignInInput(input); !ok {
		newErrorResponse(c, http.StatusBadRequest, "invalid input value")
		return
	}

	token, err := h.services.User.GenerateToken(input.Email, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func (h *Handler) userUpdate(c *gin.Context) {
	id, err := GetUserId(c)
	if err != nil {
		logrus.Println(err)
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"id": id,
	})
}

func (h *Handler) userDelete(c *gin.Context) {}
