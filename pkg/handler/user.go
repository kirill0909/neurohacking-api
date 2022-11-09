package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/kirill0909/neurohacking-api/models"
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

func (h *Handler) signIn(c *gin.Context) {}

func (h *Handler) userUpdate(c *gin.Context) {}

func (h *Handler) userDelete(c *gin.Context) {}
