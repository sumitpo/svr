package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

var path2method map[string]string

// Function to show registered routes
func showRoutes(c *gin.Engine) map[string]string {
	routes := c.Routes()
	help := make(map[string]string)

	for _, route := range routes {
		help[route.Path] = route.Method
	}
	return help
}

func index(c *gin.Context) {
	ret := gin.H{"message": "hello world!"}
	c.JSON(http.StatusOK, ret)
}

func user(c *gin.Context) {
	name := c.Param("name")
	action := c.Param("action")
	action = strings.Trim(action, "/")
	ret := gin.H{"message": name + " is " + action}
	c.JSON(http.StatusOK, ret)
}

func routeHandler(c *gin.Context) {
	c.JSON(http.StatusOK, path2method)
}
func favicon(c *gin.Context) {
	c.File("./imgs/favicon.ico")
}

func toWelcom(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "/welcome")
}
func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.GET("/", toWelcom)
	r.GET("/welcome", index)
	r.GET("/user/:name/*action", user)
	r.GET("/favicon.ico", favicon)
	if !gin.IsDebugging() {
		path2method = showRoutes(r)
		r.GET("/help", routeHandler)
	}
	r.Run(":8080")
}
