// Package topic 模型
package topic

import (
	"gohub-api/app/models"
	"gohub-api/app/models/category"
	"gohub-api/app/models/user"
	"gohub-api/pkg/database"
)

type Topic struct {
	models.BaseModel
	Title      string `json:"title,omitempty"`
	Body       string `json:"body,omitempty"`
	UserID     string `json:"userID,omitempty"`
	CategoryID string `json:"categoryID,omitempty"`
	Status     int    `json:"status,omitempty"` //状态

	//通过user_id 关联用户
	User user.User `json:"user"`
	//通过category_id 关联分类
	Category category.Category `json:"category"`

	models.CommonTimestampsField
}

//Create
func (topic *Topic) Create() {
	database.DB.Create(&topic)
}

//Save
func (topic *Topic) Save() (RowsAffected int64) {
	result := database.DB.Save(&topic)
	return result.RowsAffected
}

//Delete
func (topic *Topic) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&topic)
	return result.RowsAffected
}
