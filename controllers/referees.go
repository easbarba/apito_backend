package controllers

import (
	"net/http"

	"github.com/easbarba/apito_backend/initializers"
	"github.com/easbarba/apito_backend/models"
	"github.com/gin-gonic/gin"
)

// req data
type Input struct {
	Name  string `json:"name" binding:"required"`
	State string `json:"state"`
}

// swagger:operation GET /referees referees ListReferees
// Returns list of referees
// ---
// produces:
// - application/json
// responses:
//
//	'200':
//	    description: Successful operation
func ListReferees(c *gin.Context) {
	// set referees model
	var referees []models.Referee

	// query for all referees
	initializers.DB.Find(&referees)

	// return referees
	c.JSON(http.StatusOK, gin.H{"referees": referees})
}

// swagger:operation POST /referees referees NewReferee
// Create a new referee
// ---
// produces:
// - application/json
// responses:
//
//	'200':
//	    description: Successful operation
//	'400':
//	    description: Invalid input
func NewReferee(c *gin.Context) {
	// validate input
	var input Input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create Referee
	referee := models.Referee{Name: input.Name, State: input.State}
	result := initializers.DB.Create(&referee)

	if result == nil {
		c.Status(http.StatusBadRequest)
		return
	}

	// return Referee
	c.JSON(http.StatusOK, gin.H{"referee": referee})
}

// swagger:operation GET /referees referees GetReferee
// Returns list of referees
// ---
// produces:
// - application/json
// responses:
//
//	'200':
//	    description: Successful operation
func GetReferee(c *gin.Context) {
	// get referee
	id := c.Param("id")
	var referee models.Referee
	if err := initializers.DB.Where("id = ?", id).First(&referee).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Referee not found!"})
		return
	}

	// return referee
	c.JSON(http.StatusOK, gin.H{"referee": referee})
}

// swagger:operation PATCH /referees/{id} referees UpdateReferee
// Update an existing referee
// ---
// parameters:
//   - name: id
//     in: path
//     description: ID of the referee
//     required: true
//     type: string
//
// produces:
// - application/json
// responses:
//
//	'200':
//	    description: Successful operation
//	'400':
//	    description: Invalid input
//	'404':
//	    description: Invalid referee ID
func PartialUpdateReferee(c *gin.Context) {
	// Get referee if exist
	var referee models.Referee
	id := c.Param("id")
	if err := initializers.DB.Where("id = ?", id).First(&referee).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Referee not found!"})
		return
	}

	// Validate input
	var input Input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// update referee
	initializers.DB.Model(&referee).Updates(input)

	// return referee
	c.JSON(http.StatusOK, gin.H{"referee": referee})
}

// swagger:operation PUT /referees/{id} referees UpdateReferee
// Update an existing referee
// ---
// parameters:
//   - name: id
//     in: path
//     description: ID of the referee
//     required: true
//     type: string
//
// produces:
// - application/json
// responses:
//
//	'200':
//	    description: Successful operation
//	'400':
//	    description: Invalid input
//	'404':
//	    description: Invalid referee ID
func UpdateReferee(c *gin.Context) {
	// Get referee if exist
	var referee models.Referee

	id := c.Param("id")
	if err := initializers.DB.Where("id = ?", id).First(&referee).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Referee not found!"})
		return
	}

	// Validate input
	var input Input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// update referee
	initializers.DB.Model(&referee).Updates(input)

	// return referee
	c.JSON(http.StatusOK, gin.H{"referee": referee})
}

// swagger:operation DELETE /referees/{id} referees DeleteReferee
// Delete an existing referee
// ---
// produces:
// - application/json
// parameters:
//   - name: id
//     in: path
//     description: ID of the referee
//     required: true
//     type: string
//
// responses:
//
//	'200':
//	    description: Successful operation
//	'404':
//	    description: Invalid referee ID
func DeleteReferee(c *gin.Context) {
	// check if referee per id exist
	id := c.Param("id")
	var referee models.Referee
	if err := initializers.DB.Where("id = ?", id).First(&referee).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Referee not found!"})
		return
	}

	// delete referee
	initializers.DB.Delete(&models.Referee{}, id)

	// respond
	c.Status(200)
}

// swagger:operation GET /referees/search referees FindReferee
// Search referees based on tags
// ---
// produces:
// - application/json
// parameters:
//   - name: tag
//     in: query
//     description: referee tag
//     required: true
//     type: string
//
// responses:
//
//	'200':
//	    description: Successful operation
func FindReferee(c *gin.Context) {
	// tag := c.Query("tag")

	found := make([]models.Referee, 0)

	c.JSON(http.StatusOK, found)
}
