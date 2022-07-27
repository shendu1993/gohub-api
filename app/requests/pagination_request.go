package requests

import (
	"gohub-api/pkg/locale"

	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type PaginationRequest struct {
	Page    string ` valid:"page" form:"page"`
	Sort    string ` valid:"sort" form:"sort"`
	Order   string ` valid:"order" form:"order"`
	PerPage string `valid:"per_page" form:"per_page"`
}

func Pagination(data interface{}, c *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		"page":     []string{"numeric_between:1,1000"},
		"sort":     []string{"in:id,created_at,updated_at"},
		"order":    []string{"in:asc,desc"},
		"per_page": []string{"numeric_between:2,100"},
	}
	messages := govalidator.MapData{
		"page": []string{
			"numeric_between:" + locale.Translate(c, "users_validate_page"),
		},
		"sort": []string{
			"in:" + locale.Translate(c, "users_validate_sort"),
		},
		"order": []string{
			"in:" + locale.Translate(c, "users_validate_order"),
		},
		"per_page": []string{
			"numeric_between:" + locale.Translate(c, "users_validate_per_page"),
		},
	}
	return validate(data, rules, messages)
}
