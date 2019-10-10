package cmd

import "github.com/gin-gonic/gin"

func App() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())

	//Dependency Injection & Route Register
	routes.Configure(r)

	return r
}
