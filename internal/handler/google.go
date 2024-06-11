package handler

import (
	"net/http"

	"gin-app/internal/repository"
	"github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

type GoogleHandler struct {
	repo *repository.Repository
	l    *logrus.Logger
}

func NewGoogleHandler(repo *repository.Repository, l *logrus.Logger) *GoogleHandler {
	l.WithField("handler", "GoogleHandler")
	return &GoogleHandler{repo: repo, l: l}
}

func (h *GoogleHandler) AuthorizedCallBack(c *gin.Context) {
	const secret = "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"
	token := c.Query("access_token")
	//userInfo, err := getUserInfo(token)
	//if err != nil {
	//	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	//	return
	//}
	c.JSON(http.StatusOK, gin.H{"token": token})
}
