// Package user 存放用户 Model 相关逻辑
package user

import (
	"gohub-api/app/models"
	"gohub-api/pkg/database"
)

type User struct {
	models.BaseModel
	Name     string `json:"name,omitempty"`
	Email    string `json:"-"`
	Phone    string `json:"-"`
	Password string `json:"-"`

	models.CommonTimestampsField
}

func (userModel *User) Create() {
	database.DB.Create(&userModel)
}
