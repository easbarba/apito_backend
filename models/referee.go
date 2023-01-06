package models

import "gorm.io/gorm"

// swagger:parameters referees newRecipe
type Referee struct {
	gorm.Model
	Name  string `json:"name"`
	State string `json:"state"`
}
