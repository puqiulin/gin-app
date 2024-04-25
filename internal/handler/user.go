package handler

import (
	"net/http"
	"strconv"

	"gin-app/internal/repository"
	"gin-app/pkg/logger"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	repo repository.UserRepository
}

func NewUserHandler(repo repository.UserRepository) *UserHandler {
	return &UserHandler{repo: repo}
}

func (h *UserHandler) GetUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		logger.Log.Errorf("Invalid id: %v", id)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}

	user, err := h.repo.FindByID(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}
