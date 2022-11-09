package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/kirill0909/neurohacking-api/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{services: service}
}

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

		data := user.Group("/data", h.userIdentity)
		{
			data.PUT("/update", h.userUpdate)
			data.DELETE("/delete", h.userDelete)
		}

	}

	category := router.Group("/category", h.userIdentity)
	{
		category.POST("/", h.createCategory)
		category.GET("/", h.getAllCategories)
		category.GET("/:id", h.getCategoryById)
		category.PUT("/:id", h.updateCategory)
		category.DELETE("/:id", h.deleteCategory)

		word := category.Group("/:id/word", h.userIdentity)
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
