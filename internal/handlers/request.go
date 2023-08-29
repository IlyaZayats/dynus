package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type InsertSlugRequest struct {
	Name   string `json:"name" binding:"required" validate:"required,slug"`
	Chance string `json:"chance" binding:"required" validate:"required,chance"`
}

type DeleteSlugRequest struct {
	Name string `json:"name" binding:"required" validate:"required,slug"`
}

type GetActiveSlugsRequest struct {
	UserId string `uri:"user_id" binding:"required" validate:"required,number"`
}

type UpdateUserSlugsRequest struct {
	UserId      string            `uri:"user_id" validate:"required,number"`
	InsertSlugs []string          `json:"insert_slugs" validate:"slugslice"`
	DeleteSlugs []string          `json:"delete_slugs" validate:"slugslice"`
	Ttl         map[string]string `json:"ttl" validate:"ttl"`
}

type GetHistoryRequest struct {
	Date string `uri:"date" validate:"required,datem"`
}

type BindParams struct {
	WithUri  bool
	WithJson bool
}

func GetRequest[T any](c *gin.Context, bp BindParams, v *Validate) (T, bool) {
	var request T

	if bp.WithUri {
		if err := c.BindUri(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return request, false
		}
	}

	if bp.WithJson {
		if err := c.BindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return request, false
		}
	}

	if err := v.validate.Struct(request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "validation error", "text": err.Error()})
		return request, false
	}

	return request, true
}
