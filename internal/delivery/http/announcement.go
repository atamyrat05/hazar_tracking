package handler

import (
	"hazar_tracking/internal/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createAnnouncement(c *gin.Context) {
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

	var input model.AnnouncementInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body params"})
		return
	}
	id, err := h.service.Announcement.Create(input, IntUserId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Id": id})

}

func (h *Handler) getAllAnnouncement(c *gin.Context) {
	_, ok := c.Get(UserCtx)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not found"})
		return
	}

	data, err := h.service.Announcement.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": data})
}

func (h *Handler) getByIdAnnouncement(c *gin.Context) {
	_, ok := c.Get(UserCtx)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not found"})
		return
	}
	announcementId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id params"})
		return
	}
	data, err := h.service.Announcement.GetById(announcementId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": data})
}
