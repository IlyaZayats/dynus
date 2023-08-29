package handlers

import (
	"github.com/IlyaZayats/dynus/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/golang-module/carbon/v2"
	"net/http"
	"os"
	"strings"
)

type SlugHandlers struct {
	svc       *services.SlugService
	engine    *gin.Engine
	validator *Validate
}

func NewSlugHandlers(engine *gin.Engine, svc *services.SlugService, v *Validate) (*SlugHandlers, error) {
	h := &SlugHandlers{
		svc:       svc,
		engine:    engine,
		validator: v,
	}
	err := h.validator.InitValidator()
	h.initRoute()
	return h, err
}

func (h *SlugHandlers) initRoute() {
	h.engine.PUT("/slugs", h.InsertSlug)
	h.engine.DELETE("/slugs", h.DeleteSlug)
	h.engine.GET("/slugs/:user_id", h.GetActiveSlugs)
	h.engine.POST("/slugs/:user_id", h.UpdateUserSlugs)
	h.engine.GET("/slugs/history/:date", h.GetHistory)
}

func (h *SlugHandlers) InsertSlug(c *gin.Context) {
	req, ok := GetRequest[InsertSlugRequest](c, BindParams{WithUri: false, WithJson: true}, h.validator)
	if !ok {
		return
	}

	if err := h.svc.InsertSlug(req.Name, req.Chance); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "insert slug error", "text": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func (h *SlugHandlers) DeleteSlug(c *gin.Context) {

	req, ok := GetRequest[DeleteSlugRequest](c, BindParams{WithUri: false, WithJson: true}, h.validator)
	if !ok {
		return
	}

	if err := h.svc.DeleteSlug(req.Name); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "delete slug error", "text": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func (h *SlugHandlers) GetActiveSlugs(c *gin.Context) {

	req, ok := GetRequest[GetActiveSlugsRequest](c, BindParams{WithUri: true, WithJson: false}, h.validator)
	if !ok {
		return
	}

	activeSlugs, err := h.svc.GetActiveSlugs(req.UserId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "get active error", "text": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok", "user_id": req.UserId, "slugs": activeSlugs})
}

func (h *SlugHandlers) UpdateUserSlugs(c *gin.Context) {

	req, ok := GetRequest[UpdateUserSlugsRequest](c, BindParams{WithUri: true, WithJson: true}, h.validator)
	if !ok {
		return
	}

	//if !validateSlice(req.InsertSlugs) || !validateSlice(req.DeleteSlugs) || (len(req.InsertSlugs)+len(req.DeleteSlugs) == 0) {
	//	c.JSON(http.StatusBadRequest, gin.H{"error": "slugs validation error"})
	//	return
	//}
	if err := h.svc.UpdateUserSlugs(req.UserId, req.InsertSlugs, req.DeleteSlugs, req.Ttl); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "update slugs error", "text": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func (h *SlugHandlers) GetHistory(c *gin.Context) {
	req, ok := GetRequest[GetHistoryRequest](c, BindParams{WithUri: true, WithJson: false}, h.validator)
	if !ok {
		return
	}
	data, err := h.svc.GetHistory(req.Date)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "get history", "text": err.Error()})
		return
	}
	filename := "history" + strings.Replace(strings.Join(strings.Split(carbon.Now().String(), " "), "_"), ":", "-", -1) + ".csv"
	file, err := os.Create(filename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "open csv", "text": err.Error()})
		return
	}
	if _, err := file.WriteString(strings.Join(data, "")); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "write csv", "text": err.Error()})
		return
	}
	if err = file.Close(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "close csv", "text": err.Error()})
		return
	}
	c.FileAttachment(filename, filename)
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
