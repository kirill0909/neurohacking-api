package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/kirill0909/neurohacking-api/models"
	"github.com/sirupsen/logrus"
	"net/http"
)

func (h *Handler) createCategory(c *gin.Context) {
	userId, err := GetUserId(c)
	if err != nil {
		logrus.Println(err)
		return
	}

	var input models.Category
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	if ok := checkEmptyValueCategoryInput(input); !ok {
		newErrorResponse(c, http.StatusBadRequest, "invalid input value")
		return
	}

	category, err := h.services.Category.Create(input, userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"category": category,
	})
}

func (h *Handler) getAllCategories(c *gin.Context) {}

func (h *Handler) getCategoryById(c *gin.Context) {}

func (h *Handler) updateCategory(c *gin.Context) {}

func (h *Handler) deleteCategory(c *gin.Context) {}
