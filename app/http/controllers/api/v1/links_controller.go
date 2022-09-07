package v1

import (
	"gohub-api/app/models/link"
	"gohub-api/pkg/response"

	"github.com/gin-gonic/gin"
)

type LinksController struct {
	BaseAPIController
}

//Index Links 列表
func (ctrl *LinksController) Index(c *gin.Context) {
	links := link.GetByLimit(5)
	response.Data(c, links)
}