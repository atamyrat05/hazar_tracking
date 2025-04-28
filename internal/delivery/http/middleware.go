package handler

import (
	"hazar_tracking/utilits"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const UserCtx = "userId"

func (h *Handler) UserIdentity(c *gin.Context) {
	header := c.GetHeader("Authorization")
	if header == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "auth header is not empty"})
		return
	}

	headerPart := strings.Split(header, " ")
	if len(headerPart) != 2 || headerPart[0] != "Bearer" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid type of auth"})
		return
	}

	userId, err := h.service.Authorization.ParseToken(headerPart[1])
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	c.Set(UserCtx, userId)
	c.Next()
}

func (h *Handler) UserIdentityWithEmail(c *gin.Context) {
	header := c.GetHeader("Authorization")
	if header == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "auth header is not empty"})
		return
	}

	headerPart := strings.Split(header, " ")
	if len(headerPart) != 2 || headerPart[0] != "Bearer" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid type of auth"})
		return
	}

	userId, err := utilits.ParseToken(headerPart[1])
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	c.Set(UserCtx, userId)
	c.Next()
}
