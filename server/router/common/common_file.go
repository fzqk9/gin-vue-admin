package commom

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type CommonFileRouter struct {
}

// InitCommonFileRouter 初始化 路由信息
func (s *CommonFileRouter) InitCommonFileRouter(Router *gin.RouterGroup) {
	//commFileRouter := Router.Group("commFile").Use(middleware.OperationRecord())
	commonFileRouterWithoutRecord := Router.Group("commFile")
	var basicFileApi = v1.ApiGroupApp.AutoCodeApiGroup.BasicFileApi
	var fileApi = v1.ApiGroupApp.CommonApiGroup.CommonFileApi
	{
		//commFileRouter
	}
	{
		commonFileRouterWithoutRecord.GET("get", basicFileApi.FindBasicFile) // 根据ID获取BasicFile
		commonFileRouterWithoutRecord.POST("upload", fileApi.UploadFile)     // 获取BasicFile列表
	}
}
