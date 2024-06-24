package handler

import (
	"net/http"
	"strconv"

	"gin-app/internal/repository"
	"gin-app/pkg/uptrace"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

type UserHandler struct {
	repo *repository.Repository
	l    *logrus.Logger
	u    *uptrace.Client
}

func NewUserHandler(repo *repository.Repository, l *logrus.Logger) *UserHandler {
	l.WithField("handler", "UserHandler")
	return &UserHandler{repo: repo, l: l}
}

func (h *UserHandler) GetUser(c *gin.Context) {
	ctx := c.Request.Context()
	span := trace.SpanFromContext(ctx)

	if span.IsRecording() {
		span.SetAttributes(
			attribute.String("test key", "test value"),
			attribute.String("service_name", "user"),
		)
	}

	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		h.l.Errorf("Invalid id: %v", id)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}

	user, err := h.repo.GetUserByID(c, int64(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
	//otelgin.HTML(c, http.StatusOK, "index", gin.H{
	//	"user": user,
	//})
}
