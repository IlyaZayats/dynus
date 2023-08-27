package v1

import (
	"fmt"
	postgres "github.com/IlyaZayats/dynus/internal/db"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Handlers struct {
	g  *gin.Engine
	db *pgxpool.Pool
}

//"postgres://dynus:dynus@postgres:5432/dynus"

func NewHandlers() *Handlers {
	return &Handlers{g: gin.New(), db: postgres.NewPostgresPool("postgres://dynus:dynus@postgres:5432/dynus")}
}

func (h *Handlers) Init() {
	h.InitRoute()
	h.Run()
}

func (h *Handlers) InitRoute() {
	h.g.PUT("/slugs", h.InitSlug)
	h.g.DELETE("/slugs", h.RemoveSlug)
	h.g.GET("/slugs/:user_id", h.ActiveUserSlugs)
	h.g.POST("/slugs/:user_id", h.UpdateUserSlugs)

}

func (h *Handlers) Run() {
	if err := h.g.Run(":8080"); err != nil {
		fmt.Println(err.Error())
	}
}

func (h *Handlers) CloseConnection() {
	h.db.Close()
}
