package handler

import (
	"hazar_tracking/internal/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.Static("/uploads", "./uploads")
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	profile := router.Group("/user", h.UserIdentity)
	{
		profile.GET("/", h.getAllUser)
		profile.GET("/:id", h.getByIdUser)
		profile.PUT("/:id", h.updateProfile)
		profile.DELETE("/:id", h.deleteUser)
	}

	orders := router.Group("/orders", h.UserIdentity)
	{
		orders.POST("/", h.createOrders)
		orders.GET("/", h.getAllOrders)
		orders.GET("/:id", h.getByIdOrders)
		orders.POST("/search", h.searching)
		orders.GET("/points", h.getAllPoints)
		orders.PUT("/:id", h.updateOrdes)
		orders.POST("/points", h.createOrderPoints)
		orders.PUT("/points/:id", h.UpdateOrderPoints)
	}
	email := router.Group("/email")
	{
		email.POST("/", h.validationEmail)
		email.POST("/test_code", h.UserIdentityWithEmail, h.testCode)
		email.PUT("/update_password", h.UserIdentityWithEmail, h.updatePassword)
	}
	announcement := router.Group("/announcement", h.UserIdentity)
	{
		announcement.POST("/", h.createAnnouncement)
		announcement.GET("/", h.getAllAnnouncement)
		announcement.GET("/:id", h.getByIdAnnouncement)
	}

	return router
}
