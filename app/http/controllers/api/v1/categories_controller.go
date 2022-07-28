package v1

import (
	"gohub-api/app/models/category"
	"gohub-api/app/requests"
	"gohub-api/pkg/locale"
	"gohub-api/pkg/response"

	"github.com/gin-gonic/gin"
)

type CategoriesController struct {
	BaseAPIController
}

func (ctrl *CategoriesController) Index(c *gin.Context) {
	categories := category.All()
	response.Data(c, categories)
}

func (ctrl *CategoriesController) Show(c *gin.Context) {
	categoryModel := category.Get(c.Param("id"))
	if categoryModel.ID == 0 {
		response.Abort404(c)
		return
	}
	response.Data(c, categoryModel)
}

func (ctrl *CategoriesController) Update(c *gin.Context) {
	//查询是否存在该分类
	categoryModel := category.Get(c.Param("id"))
	if categoryModel.ID < 1 {
		response.Abort404(c)
		return
	}
	//参数校验
	request := requests.CategoryRequest{}
	if ok := requests.Validate(c, &request, requests.CategorySave); !ok {
		return
	}
	//更改保存数据
	categoryModel.Name = request.Name
	categoryModel.Description = request.Description
	rowsAffected := categoryModel.Save()
	if rowsAffected > 0 {
		response.Data(c, categoryModel)
	} else {
		response.Abort500(c, locale.Translate(c, "category_validate_update"))
	}

}

//Store 新增分类
func (ctrl *CategoriesController) Store(c *gin.Context) {

	request := requests.CategoryRequest{}
	if ok := requests.Validate(c, &request, requests.CategorySave); !ok {
		return
	}

	categoryModel := category.Category{
		Name:        request.Name,
		Description: request.Description,
	}
	categoryModel.Create()
	if categoryModel.ID > 0 {
		response.Created(c, categoryModel)
	} else {
		response.Abort500(c, "创建失败，请稍后尝试~")
	}
}
