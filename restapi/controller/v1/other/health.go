package other

import (
	"Go-IMS/response"
	"Go-IMS/service/other"
	"github.com/gin-gonic/gin"
)

// ConGetSystemInfo 控制层：获取系统信息
// 参数：
//		c *gin.Context：gin.Context的指针
// 返回值：
//		无
func ConGetSystemInfo(c *gin.Context) {
	resStruct := other.SerGetSystemInfo()
	response.Response(c, resStruct)
}
