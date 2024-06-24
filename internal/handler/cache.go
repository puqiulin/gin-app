package handler

import (
	"net/http"

	"gin-app/internal/repository"
	"gin-app/pkg/uptrace"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type CacheHandler struct {
	repo *repository.Repository
	l    *logrus.Logger
	u    *uptrace.Client
}

func NewCacheHandler(repo *repository.Repository, l *logrus.Logger) *CacheHandler {
	l.WithField("handler", "CacheHandler")
	return &CacheHandler{repo: repo, l: l}
}

func (h *CacheHandler) GetRank(c *gin.Context) {
	ranking, err := h.repo.GetRank(c, c.Query("key"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, ranking)
}
