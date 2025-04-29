package handler

import (
	"hazar_tracking/internal/model"
	"hazar_tracking/utilits"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) updateProfile(c *gin.Context) {
	_, ok := c.Get(UserCtx)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not found"})
		return
	}
	IntUserId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id params"})
		return
	}

	var dataInput model.UpdateProfileInput
	if err := c.ShouldBind(&dataInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Faýllary almak şowsuz boldy"})
		return
	}
	data, err := h.service.Profile.GetById(IntUserId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if data.Image_Url != nil {
		err := utilits.RemoveFile(*data.Image_Url)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "profildaki suraty pozup bolmady"})
			return
		}
	}

	files := form.File["Images"]
	if len(files) > 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "should must be one photo"})
		return
	}

	imageUrl, err := utilits.CreateFile(files)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "profildaki suraty tazelap bolmady"})
		return
	}
	dataInput.Image_Url = &imageUrl[0]
	err = h.service.Profile.Update(IntUserId, dataInput)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func (h *Handler) getAllUser(c *gin.Context) {
	_, ok := c.Get(UserCtx)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not found"})
		return
	}
	data, err := h.service.Profile.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": data})
}

func (h *Handler) getByIdUser(c *gin.Context) {
	_, ok := c.Get(UserCtx)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not found"})
		return
	}

	IntUserId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id params"})
		return
	}

	data, err := h.service.Profile.GetById(IntUserId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": data})
}

func (h *Handler) deleteUser(c *gin.Context) {
	_, ok := c.Get(UserCtx)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not found"})
		return
	}

	IntUserId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id params"})
		return
	}

	err = h.service.Profile.Delete(IntUserId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
