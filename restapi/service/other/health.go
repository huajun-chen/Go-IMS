package other

import (
	"Go-WMS/global"
	"Go-WMS/param"
	"Go-WMS/param/resp"
	"Go-WMS/utils"
	"go.uber.org/zap"
	"net/http"
)

// SerGetSystemInfo 业务层：获取系统信息
// 参数：
//		无
// 返回值：
//		param.Resp：响应的结构体
func SerGetSystemInfo() param.Resp {
	// CPU
	cpuStruct, err := utils.CPUInfo()
	if err != nil {
		zap.S().Errorf("%s：%s", global.I18nMap["10023"], err)
		failStruct := param.Resp{
			Code: 10023,
			Msg:  global.I18nMap["10023"],
		}
		return failStruct
	}
	// 内存
	memStruct, err := utils.MemInfo()
	if err != nil {
		zap.S().Errorf("%s：%s", global.I18nMap["10024"], err)
		failStruct := param.Resp{
			Code: 10024,
			Msg:  global.I18nMap["10024"],
		}
		return failStruct
	}
	// 硬盘
	diskStruct, err := utils.DiskInfo()
	if err != nil {
		zap.S().Errorf("%s：%s", global.I18nMap["10025"], err)
		failStruct := param.Resp{
			Code: 10025,
			Msg:  global.I18nMap["10025"],
		}
		return failStruct
	}

	data := resp.RespSystem{
		CPU:    cpuStruct,
		Memory: memStruct,
		Disk:   diskStruct,
	}

	succStruct := param.Resp{
		Code: http.StatusOK,
		Data: data,
	}
	return succStruct
}
