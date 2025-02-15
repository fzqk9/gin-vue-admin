package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CmsArticleRouter struct {
}

// InitCmsArticleRouter 初始化 CmsArticle 路由信息
func (s *CmsArticleRouter) InitCmsArticleRouter(Router *gin.RouterGroup) {
	cmsArticleRouter := Router.Group("cmsArticle").Use(middleware.OperationRecord())
	cmsArticleRouterWithoutRecord := Router.Group("cmsArticle")
	var cmsArticleApi = v1.ApiGroupApp.AutoCodeApiGroup.CmsArticleApi
	{
		cmsArticleRouter.POST("createCmsArticle", cmsArticleApi.CreateCmsArticle)   // 新建CmsArticle
		cmsArticleRouter.DELETE("deleteCmsArticle", cmsArticleApi.DeleteCmsArticle) // 删除CmsArticle
		cmsArticleRouter.DELETE("deleteCmsArticleByIds", cmsArticleApi.DeleteCmsArticleByIds) // 批量删除CmsArticle
		cmsArticleRouter.PUT("updateCmsArticle", cmsArticleApi.UpdateCmsArticle)    // 更新CmsArticle
	    cmsArticleRouter.POST("quickEdit", cmsArticleApi.QuickEdit)  // 快速编辑
	}
	{
		cmsArticleRouterWithoutRecord.GET("findCmsArticle", cmsArticleApi.FindCmsArticle)        // 根据ID获取CmsArticle
		cmsArticleRouterWithoutRecord.GET("getCmsArticleList", cmsArticleApi.GetCmsArticleList)  // 获取CmsArticle列表
	}
}
