package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/kirill0909/neurohacking-api/models"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

func (h *Handler) createWord(c *gin.Context) {
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

	var input models.Word
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input boyd")
		return
	}

	if ok := checkEmptyValueWord(input); !ok {
		newErrorResponse(c, http.StatusBadRequest, "invalid input value")
		return
	}

	word, err := h.services.Word.Create(input, userId, categoryId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"word": word,
	})
}

func (h *Handler) getAllWords(c *gin.Context) {
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

	words, err := h.services.Word.GetAll(userId, categoryId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"words": words,
	})

}

func (h *Handler) getWordById(c *gin.Context) {}

func (h *Handler) updateWord(c *gin.Context) {}

func (h *Handler) deleteWord(c *gin.Context) {}
