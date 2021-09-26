package commom

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type CommFileRouter struct {
}

// InitCommFileRouter 初始化 路由信息
func (s *CommFileRouter) InitCommFileRouter(Router *gin.RouterGroup) {
	//commFileRouter := Router.Group("commFile").Use(middleware.OperationRecord())
	commFileRouterWithoutRecord := Router.Group("commFile")
	var basicFileApi = v1.ApiGroupApp.AutoCodeApiGroup.BasicFileApi
	{
		//commFileRouter
	}
	{
		commFileRouterWithoutRecord.GET("get", basicFileApi.FindBasicFile)        // 根据ID获取BasicFile
		commFileRouterWithoutRecord.POST("upload", basicFileApi.GetBasicFileList) // 获取BasicFile列表
	}
}
