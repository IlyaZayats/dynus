package routing

import "github.com/gin-gonic/gin"

type HandlerInterface interface {
	Init()
	Run()
	CloseConnection()
	DeleteSlug(c *gin.Context)
	InsertSlug(c *gin.Context)
	UpdateUserSlugs(c *gin.Context)
	GetActiveSlugs(c *gin.Context)
}
