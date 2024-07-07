package main

import (
	"gosvr/db"

	"github.com/gin-gonic/gin"
)

func pingHandler(c *gin.Context) {
	c.JSON(200, gin.H{"message": "pong"})
}
func main() {
	db.MysqlInit()
	db.RedisInit()

	r := gin.Default()
	r.GET("/ping", pingHandler)
	r.Run()
}
