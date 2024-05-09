package handler

import (
	"net/http"
	"strconv"

	"gin-app/internal/repository"
	"github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	repo repository.UserRepository
	l    *logrus.Logger
}

func NewUserHandler(repo repository.UserRepository, l *logrus.Logger) *UserHandler {
	return &UserHandler{repo: repo, l: l}
}

func (h *UserHandler) GetUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		h.l.Errorf("Invalid id: %v", id)
		c.JSON(http.StatusBadRequest, gin.H{"error.tsx": "invalid ID"})
		return
	}

	user, err := h.repo.FindByID(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error.tsx": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}
