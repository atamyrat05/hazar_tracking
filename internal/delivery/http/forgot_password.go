package handler

import (
	"hazar_tracking/internal/model"
	"hazar_tracking/utilits"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) validationEmail(c *gin.Context) {
	var email model.Email
	if err := c.ShouldBindJSON(&email); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userId, err := h.service.Email.Validate(email.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := utilits.GenerateToken(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	code := utilits.GenerateRandomCode()
	err = h.service.Email.UpdateForgotCode(userId, code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = utilits.SendEmailSendGrid(code, email.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})

}

func (h *Handler) testCode(c *gin.Context) {
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

	var code model.Code
	if err := c.ShouldBindBodyWithJSON(&code); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body param"})
		return
	}
	err := h.service.Email.TestCode(code.Code, IntUserId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid code"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func (h *Handler) updatePassword(c *gin.Context) {
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

	var input model.UpdatePassword
	if err := c.ShouldBindBodyWithJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body param"})
		return
	}
	err := h.service.Email.UpdateUsersPassword(input, IntUserId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
