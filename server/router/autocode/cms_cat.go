package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CmsCatRouter struct {
}

// InitCmsCatRouter 初始化 CmsCat 路由信息
func (s *CmsCatRouter) InitCmsCatRouter(Router *gin.RouterGroup) {
	cmsCatRouter := Router.Group("cmsCat").Use(middleware.OperationRecord())
	cmsCatRouterWithoutRecord := Router.Group("cmsCat")
	var cmsCatApi = v1.ApiGroupApp.AutoCodeApiGroup.CmsCatApi
	{
		cmsCatRouter.POST("createCmsCat", cmsCatApi.CreateCmsCat)   // 新建CmsCat
		cmsCatRouter.DELETE("deleteCmsCat", cmsCatApi.DeleteCmsCat) // 删除CmsCat
		cmsCatRouter.DELETE("deleteCmsCatByIds", cmsCatApi.DeleteCmsCatByIds) // 批量删除CmsCat
		cmsCatRouter.PUT("updateCmsCat", cmsCatApi.UpdateCmsCat)    // 更新CmsCat
	    cmsCatRouter.POST("quickEdit", cmsCatApi.QuickEdit)  // 快速编辑
	}
	{
		cmsCatRouterWithoutRecord.GET("findCmsCat", cmsCatApi.FindCmsCat)        // 根据ID获取CmsCat
		cmsCatRouterWithoutRecord.GET("getCmsCatList", cmsCatApi.GetCmsCatList)  // 获取CmsCat列表
	}
}
