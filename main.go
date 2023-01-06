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

import (
	"github.com/easbarba/apito_backend/controllers"
	"github.com/easbarba/apito_backend/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.EnvironmentVariables()
	initializers.Database()
}

func Index(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Apito final!"})
}

func main() {
	router := gin.Default()
	v1 := router.Group("api/v1")
	{
		// index
		v1.GET("/", Index)

		// referees
		v1.GET("/referees", controllers.ListReferees)
		v1.POST("/referees", controllers.NewReferee)
		v1.GET("/referees/:id", controllers.GetReferee)
		v1.PUT("/referees/:id", controllers.UpdateReferee)
		v1.PATCH("/referees/:id", controllers.PartialUpdateReferee)
		v1.DELETE("/referees/:id", controllers.DeleteReferee)
		v1.GET("/referees/search", controllers.FindReferee)
	}
	router.Run(":8080")
}
