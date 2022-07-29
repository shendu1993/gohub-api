// Package category 模型
package category

import (
	"gohub-api/app/models"
	"gohub-api/pkg/database"
)

type Category struct {
	models.BaseModel
	Name        string `json:"name,omitempty"`        //名称
	Description string `json:"description,omitempty"` //描述
	Status      int    `json:"status,omitempty"`      //描述
	models.CommonTimestampsField
}

//Create
func (category *Category) Create() {
	database.DB.Create(&category)
}

//Save
func (category *Category) Save() (RowsAffected int64) {
	result := database.DB.Save(&category)
	return result.RowsAffected
}

//Delete
func (category *Category) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&category)
	return result.RowsAffected
}
