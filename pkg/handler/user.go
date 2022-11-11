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

	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})

}

func (h *Handler) signIn(c *gin.Context) {
	var input models.UserSignInInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	if ok := checkEmptyValuesUserSignInInput(input); !ok {
		newErrorResponse(c, http.StatusBadRequest, "invalid input value")
		return
	}

	token, err := h.services.User.GenerateToken(input.Email, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func (h *Handler) userUpdate(c *gin.Context) {
	id, err := GetUserId(c)
	if err != nil {
		logrus.Println(err)
		return
	}

	var input models.UserUpdateInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if ok := checkEmptyValueUserUpdateInput(input); !ok {
		newErrorResponse(c, http.StatusBadRequest, "The value should not be empty")
		return
	}

	if err := h.services.User.Update(input, id); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"staus": "ok",
	})
}

func (h *Handler) userDelete(c *gin.Context) {
	userId, err := GetUserId(c)
	if err != nil {
		logrus.Println(err)
		return
	}

	if err := h.services.User.Delete(userId); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
