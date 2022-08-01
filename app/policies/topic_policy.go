//package policies 用户授权
package policies

import (
	"gohub-api/app/models/topic"
	"gohub-api/pkg/auth"

	"github.com/gin-gonic/gin"
)

//CanModifyTopic check can modify topic auth
func CanModifyTopic(c *gin.Context, _topic topic.Topic) bool {
	return auth.CurrentUID(c) == _topic.UserID
}
