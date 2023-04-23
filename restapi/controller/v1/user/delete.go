package user

import (
	"Go-IMS/param"
	"Go-IMS/service/user"
	"Go-IMS/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// ConDeleteUser 控制层：删除用户（管理员权限）
// 参数：
//		c *gin.Context：gin.Context的指针
// 返回值：
//		无
func ConDeleteUser(c *gin.Context) {
	userId := param.ReqId{}
	if err := c.ShouldBindUri(&userId); err != nil {
		parError := utils.HandleValidatorError(err)
		c.JSON(http.StatusOK, parError)
		return
	}
	resStruct := user.SerDeleteUser(userId)
	c.JSON(http.StatusOK, resStruct)
}
