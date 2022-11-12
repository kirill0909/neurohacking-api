package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/kirill0909/neurohacking-api/models"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
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

func (h *Handler) getAllCategories(c *gin.Context) {
	userId, err := GetUserId(c)
	if err != nil {
		logrus.Println(err)
		return
	}

	categories, err := h.services.Category.GetAll(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"categories": categories,
	})

}

func (h *Handler) getCategoryById(c *gin.Context) {
	userId, err := GetUserId(c)
	if err != nil {
		logrus.Println(err)
		return
	}

	categoryId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id parameter")
		return
	}

	category, err := h.services.GetById(userId, categoryId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"category": category,
	})
}

func (h *Handler) updateCategory(c *gin.Context) {
	userId, err := GetUserId(c)
	if err != nil {
		logrus.Println(err)
		return
	}

	categoryId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id parameter")
		return
	}

	var input models.CategoryUpdateInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	if ok := checkEmptyValueCategoryUpdateInput(input); !ok {
		newErrorResponse(c, http.StatusBadRequest, "invalid input value")
		return
	}

	updatedCategory, err := h.services.Category.Update(input, userId, categoryId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"category": updatedCategory,
	})
}

func (h *Handler) deleteCategory(c *gin.Context) {}
