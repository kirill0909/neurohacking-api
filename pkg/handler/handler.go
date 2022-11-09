package handler

import (
	"github.com/gin-gonic/gin"
)

type Handler struct{}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	// This endpoint is only needed to check
	// the running server
	router.GET("/", h.ping)

	user := router.Group("/user")
	{
		auth := user.Group("/auth")
		{
			auth.POST("/sign-up", h.signUp)
			auth.POST("/sign-in", h.signIn)
		}

		user.PUT("/update", h.userUpdate)
		user.DELETE("/delete", h.userDelete)
	}

	category := router.Group("/category")
	{
		category.POST("/", h.createCategory)
		category.GET("/", h.getAllCategories)
		category.GET("/:id", h.getCategoryById)
		category.PUT("/:id", h.updateCategory)
		category.DELETE("/:id", h.deleteCategory)

		word := category.Group("/:id/word")
		{
			word.POST("/", h.createWord)
			word.GET("/", h.getAllWords)
			word.GET("/:word_id", h.getWordById)
			word.PUT("/:word_id", h.updateWord)
			word.DELETE("/:word_id", h.deleteWord)

		}
	}

	return router
}
