package v1

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
)

func (h *Handlers) InitSlug(c *gin.Context) {
	var data NewSlugRequest
	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if matched, err := regexp.MatchString("^[a-zA-Z0-9_]+$", data.Name); !matched || err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "validation error"})
		return
	}
	//db := postgres.NewPostgresConnection(h.url)
	//defer db.Close(context.Background())
	if _, err := h.db.Query(context.Background(), "INSERT INTO Slugs (name) VALUES ($1)", data.Name); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "insert error", "text": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
