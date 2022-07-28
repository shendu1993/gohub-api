package requests

import (
	"gohub-api/pkg/locale"

	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type CategoryRequest struct {
	Name        string `valid:"name" json:"name"`
	Description string `valid:"description" json:"description,omitempty"`
}

func CategorySave(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"name":        []string{"required", "min_cn:2", "max_cn:8", "not_exists:categories,name"},
		"description": []string{"min_cn:3", "max_cn:255"},
	}
	messages := govalidator.MapData{
		"name": []string{
			"required:" + locale.Translate(c, "category_validate_name_required"),
			"min_cn:" + locale.Translate(c, "category_validate_name_min_cn"),
			"max_cn:" + locale.Translate(c, "category_validate_name_max_cn"),
			"not_exists:名称已存在" + locale.Translate(c, "category_validate_name_not_exists"),
		},
		"description": []string{
			"min_cn:描述长度需至少 3 个字" + locale.Translate(c, "category_validate_description_min_cn"),
			"max_cn:描述长度不能超过 255 个字" + locale.Translate(c, "category_validate_name_max_cn"),
		},
	}
	return validate(data, rules, messages)
}
