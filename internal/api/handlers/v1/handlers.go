package v1

import (
	"fmt"
	"github.com/IlyaZayats/dynus/internal/db"
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
)

type Handlers struct {
	g  *gin.Engine
	db db.DataBaseInterface
}

//"postgres://dynus:dynus@postgres:5432/dynus"

func NewHandlers(db db.DataBaseInterface) *Handlers {
	return &Handlers{g: gin.New(), db: db}
}

func (h *Handlers) Init() {
	h.g.PUT("/slugs", h.InsertSlug)
	h.g.DELETE("/slugs", h.DeleteSlug)
	h.g.GET("/slugs/:user_id", h.GetActiveSlugs)
	h.g.POST("/slugs/:user_id", h.UpdateUserSlugs)
}

func (h *Handlers) Run() {
	if err := h.g.Run(":8080"); err != nil {
		fmt.Println(err.Error())
	}
}

func (h *Handlers) CloseConnection() {
	h.db.CloseConnection()
}

func (h *Handlers) InsertSlug(c *gin.Context) {
	var req NewSlugRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if matched, err := regexp.MatchString("^[a-zA-Z0-9_]+$", req.Name); !matched || err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "validation error"})
		return
	}
	if err := h.db.InsertSlug(req.Name); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "insert slug error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func (h *Handlers) DeleteSlug(c *gin.Context) {
	var req RemoveSlugRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if matched, err := regexp.MatchString("^[a-zA-Z0-9_]+$", req.Name); !matched || err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "validation error"})
		return
	}
	if err := h.db.DeleteSlug(req.Name); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "delete slug error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func (h *Handlers) GetActiveSlugs(c *gin.Context) {
	var data ActiveUserSlugsRequest
	if err := c.BindUri(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if matched, err := regexp.MatchString("^[\\d]+$", data.UserId); !matched || err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "validation error"})
		return
	}
	activeSlugs, err := h.db.GetActiveSlugs(data.UserId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok", "user_id": data.UserId, "slugs": activeSlugs})
}

func validateSlice(slice []string) bool {
	if len(slice) != 0 {
		for _, v := range slice {
			if matched, err := regexp.MatchString("^[a-zA-Z0-9_]+$", v); !matched || err != nil {
				return false
			}
		}
	}
	return true
}

func (h *Handlers) UpdateUserSlugs(c *gin.Context) {
	var req UpdateUserSlugsRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if matched, err := regexp.MatchString("^[\\d]+$", req.UserId); !matched || err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id validation error"})
		return
	}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if !validateSlice(req.InsertSlugs) || !validateSlice(req.DeleteSlugs) || (len(req.InsertSlugs)+len(req.DeleteSlugs) == 0) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "slugs validation error"})
		return
	}
	if err := h.db.UpdateUserSlugs(req.UserId, req.InsertSlugs, req.DeleteSlugs); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "update slugs error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
