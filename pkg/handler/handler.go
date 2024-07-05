package handler

import (
	"time-tracker/pkg/service"

	_ "time-tracker/docs"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	logrus.Info("Initializing routes")

	user := router.Group("/api/user")
	{
		user.POST("/", h.createUser)
		user.GET("/", h.getAllUsers)
		user.DELETE("/:id", h.deleteUser)
		user.PUT("/:id", h.updateUser)
		user.GET("/search", h.searchUsers)
	}

	item := router.Group("/api/item")
	{
		item.POST("/", h.createItem)
		item.GET("/:id", h.getItemsById)
		item.DELETE("/:id", h.deleteItem)
		item.PUT("/:id", h.updateItem)
		item.PUT("/:id/time/:flag", h.updateItemTime)
		item.GET("/:id/time", h.getItemsByDate)
	}

	logrus.Info("Routes initialized successfully")
	return router
}

type errorResponse struct {
	Message string `json:"message"`
}
