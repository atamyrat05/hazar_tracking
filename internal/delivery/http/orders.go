package handler

import (
	"hazar_tracking/internal/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createOrders(c *gin.Context) {
	userId, ok := c.Get(UserCtx)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not found"})
		return
	}
	IntUserId, ok := userId.(int)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "userId can not converted int"})
		return
	}

	var input model.OrdersInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body params"})
		return
	}
	id, err := h.service.Orders.Create(input, IntUserId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Id": id})

}

func (h *Handler) getAllOrders(c *gin.Context) {
	userId, ok := c.Get(UserCtx)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not found"})
		return
	}
	IntUserId, ok := userId.(int)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "userId can not converted int"})
		return
	}

	data, err := h.service.Orders.GetAll(IntUserId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": data})
}

type outputData struct {
	Data      model.OrdersGet            `json:"data"`
	Locations []model.OrderTrackingSteps `json:"locations"`
}

func (h *Handler) getByIdOrders(c *gin.Context) {
	userId, ok := c.Get(UserCtx)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not found"})
		return
	}
	IntUserId, ok := userId.(int)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "userId can not converted int"})
		return
	}
	orderId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id params"})
		return
	}
	data, data2, err := h.service.Orders.GetById(IntUserId, orderId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, outputData{
		Data:      data,
		Locations: data2,
	})
}

type searchInput struct {
	Data string `json:"data" binding:"required"`
}

func (h *Handler) searching(c *gin.Context) {
	userId, ok := c.Get(UserCtx)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not found"})
		return
	}
	IntUserId, ok := userId.(int)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "userId can not converted int"})
		return
	}
	var input searchInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body params"})
		return
	}

	data, err := h.service.Orders.Search(IntUserId, input.Data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	order, location, err := h.service.Orders.GetById(IntUserId, data.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, outputData{
		Data:      order,
		Locations: location,
	})
}

func (h *Handler) getAllPoints(c *gin.Context) {
	_, ok := c.Get(UserCtx)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not found"})
		return
	}
	data, err := h.service.Orders.GetAllPoints()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"points": data})
}

func (h *Handler) updateOrdes(c *gin.Context) {
	_, ok := c.Get(UserCtx)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not found"})
		return
	}
	orderId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id params"})
		return
	}
	var input model.UpdateOrderInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body params"})
		return
	}
	err = h.service.Orders.Update(orderId, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func (h *Handler) createOrderPoints(c *gin.Context) {
	_, ok := c.Get(UserCtx)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not found"})
		return
	}
	var input model.OrderTrackingStepsInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body params"})
		return
	}
	id, err := h.service.Orders.CreatePoints(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Id": id})
}

func (h *Handler) UpdateOrderPoints(c *gin.Context) {
	_, ok := c.Get(UserCtx)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not found"})
		return
	}

	orderId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id params"})
		return
	}

	var input model.UpdateOrderTrackingStepsInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body params"})
		return
	}
	input.OrderId = orderId
	err = h.service.Orders.UpdatePoints(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
