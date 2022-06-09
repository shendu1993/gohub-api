// Package user 存放用户 Model 相关逻辑
package user

import (
	"gohub-api/app/models"
	"gohub-api/pkg/database"
	"gohub-api/pkg/hash"
)

type User struct {
	models.BaseModel
	Name     string `json:"name,omitempty"`
	Email    string `json:"-"`
	Phone    string `json:"-"`
	Password string `json:"-"`

	models.CommonTimestampsField
}

//添加用户
func (userModel *User) Create() {
	database.DB.Create(&userModel)
}

// ComparePassword 密码是否正确
func (userModel *User) ComparePassword(_password string) bool {
	return hash.BcryptCheck(_password, userModel.Password)
}

//更新用户信息
func (userModel *User) Save() (RowsAffected int64) {
	result := database.DB.Save(&userModel)
	return result.RowsAffected
}
