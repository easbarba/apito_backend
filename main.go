// Apito backend API
//
// Evaluate soccer referees' performance.
//
//	Schemes: http
//	Host: localhost:8080
//	BasePath: /
//	Version: 0.1.0
//	Contact: EAS Barba <easbarba@outlook.com> https://easbarba.github.io
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
// swagger:meta
package main

import "github.com/gin-gonic/gin"

func Index(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Apito final!"})
}

func main() {
	router := gin.Default()
	v1 := router.Group("api/v1")
	{
		// index
		v1.GET("/", Index)
	}

	router.Run()
}
