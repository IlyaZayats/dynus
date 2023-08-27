package main

import (
	v1 "github.com/IlyaZayats/dynus/internal/api/v1"
)

func main() {
	r := v1.NewHandlers()
	r.Init()
	defer r.CloseConnection()

	//r := gin.New()

	//r.GET("/ping", func(c *gin.Context) {
	//	c.JSON(http.StatusOK, gin.H{
	//		"message": "pong",
	//	})
	//})

}
