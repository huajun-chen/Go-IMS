package response

import (
	"Go-IMS/param"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Response 统一封装响应函数
// 参数：
//		c：gin.Context的指针
//		response：相应的结构体
// 返回值：
//		无
func Response(c *gin.Context, response param.Resp) {
	// 所有请求的响应系统状态码都返回200，在code字段自定义状态码
	c.JSON(http.StatusOK, response)
	return
}
