// Package link 模型
package link

import (
	"gohub-api/app/models"
	"gohub-api/pkg/database"
)

type Link struct {
	models.BaseModel
	Name   string `json:"name,omitempty"`
	URL    string `json:"URL,omitempty"`
	Status int    `json:"status,omitempty"`
	models.CommonTimestampsField
}

//Create
func (link *Link) Create() {
	database.DB.Create(&link)
}

//Save
func (link *Link) Save() (RowsAffected int64) {
	result := database.DB.Save(&link)
	return result.RowsAffected
}

//Delete
func (link *Link) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&link)
	return result.RowsAffected
}
