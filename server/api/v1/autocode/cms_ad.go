package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    autocodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type CmsAdApi struct {
}

var cmsAdService = service.ServiceGroupApp.AutoCodeServiceGroup.CmsAdService
 

// CreateCmsAd 创建CmsAd
// @Tags CmsAd
// @Summary 创建CmsAd
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.CmsAd true "创建CmsAd"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cmsAd/createCmsAd [post]
func (cmsAdApi *CmsAdApi) CreateCmsAd(c *gin.Context) {
	var cmsAd autocode.CmsAd
	_ = c.ShouldBindJSON(&cmsAd)
	if err := cmsAdService.CreateCmsAd(cmsAd); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCmsAd 删除CmsAd
// @Tags CmsAd
// @Summary 删除CmsAd
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.CmsAd true "删除CmsAd"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cmsAd/deleteCmsAd [delete]
func (cmsAdApi *CmsAdApi) DeleteCmsAd(c *gin.Context) {
	var cmsAd autocode.CmsAd
	_ = c.ShouldBindJSON(&cmsAd)
	if err := cmsAdService.DeleteCmsAd(cmsAd); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCmsAdByIds 批量删除CmsAd
// @Tags CmsAd
// @Summary 批量删除CmsAd
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除CmsAd"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /cmsAd/deleteCmsAdByIds [delete]
func (cmsAdApi *CmsAdApi) DeleteCmsAdByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := cmsAdService.DeleteCmsAdByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCmsAd 更新CmsAd
// @Tags CmsAd
// @Summary 更新CmsAd
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.CmsAd true "更新CmsAd"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /cmsAd/updateCmsAd [put]
func (cmsAdApi *CmsAdApi) UpdateCmsAd(c *gin.Context) {
	var cmsAd autocode.CmsAd
	_ = c.ShouldBindJSON(&cmsAd)
	if err := cmsAdService.UpdateCmsAd(cmsAd); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCmsAd 用id查询CmsAd
// @Tags CmsAd
// @Summary 用id查询CmsAd
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.CmsAd true "用id查询CmsAd"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /cmsAd/findCmsAd [get]
func (cmsAdApi *CmsAdApi) FindCmsAd(c *gin.Context) {
	var cmsAd autocode.CmsAd
	_ = c.ShouldBindQuery(&cmsAd)
	if err, recmsAd := cmsAdService.GetCmsAd(cmsAd.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		//response.OkWithData(gin.H{"recmsAd": recmsAd}, c)
		response.OkWithData(gin.H{"cmsAd": recmsAd}, c)
	}
}

// GetCmsAdList 分页获取CmsAd列表
// @Tags CmsAd
// @Summary 分页获取CmsAd列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.CmsAdSearch true "分页获取CmsAd列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cmsAd/getCmsAdList [get]
func (cmsAdApi *CmsAdApi) GetCmsAdList(c *gin.Context) {
	createdAtBetween, _ := c.GetQueryArray("createdAtBetween[]")

	var pageInfo autocodeReq.CmsAdSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := cmsAdService.GetCmsAdInfoList(pageInfo,createdAtBetween); err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
        response.FailWithMessage("获取失败", c)
    } else {
        response.OkWithDetailed(response.PageResult{
            List:     list,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, "获取成功", c)
    }
}



// QuickEdit 快速更新
// @Tags QuickEdit
// @Summary 快速更新
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.CmsAd true "快速更新CmsAd" 
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router  /cmsAd/quickEdit [post] 
func (cmsAdApi *CmsAdApi) QuickEdit(c *gin.Context) {
	var quickEdit request.QuickEdit
	_ = c.ShouldBindJSON(&quickEdit)
	quickEdit.Table = "cms_ad"
	//var_dump.Dump(quickEdit)
	if err := commDbService.QuickEdit(quickEdit); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}