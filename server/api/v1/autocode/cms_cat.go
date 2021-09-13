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

type CmsCatApi struct {
}

var cmsCatService = service.ServiceGroupApp.AutoCodeServiceGroup.CmsCatService
var commDbService = service.ServiceGroupApp.CommonServiceGroup.CommDbService

// CreateCmsCat 创建CmsCat
// @Tags CmsCat
// @Summary 创建CmsCat
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.CmsCat true "创建CmsCat"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cmsCat/createCmsCat [post]
func (cmsCatApi *CmsCatApi) CreateCmsCat(c *gin.Context) {
	var cmsCat autocode.CmsCat
	_ = c.ShouldBindJSON(&cmsCat)
	if err := cmsCatService.CreateCmsCat(cmsCat); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCmsCat 删除CmsCat
// @Tags CmsCat
// @Summary 删除CmsCat
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.CmsCat true "删除CmsCat"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cmsCat/deleteCmsCat [delete]
func (cmsCatApi *CmsCatApi) DeleteCmsCat(c *gin.Context) {
	var cmsCat autocode.CmsCat
	_ = c.ShouldBindJSON(&cmsCat)
	if err := cmsCatService.DeleteCmsCat(cmsCat); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCmsCatByIds 批量删除CmsCat
// @Tags CmsCat
// @Summary 批量删除CmsCat
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除CmsCat"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /cmsCat/deleteCmsCatByIds [delete]
func (cmsCatApi *CmsCatApi) DeleteCmsCatByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := cmsCatService.DeleteCmsCatByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCmsCat 更新CmsCat
// @Tags CmsCat
// @Summary 更新CmsCat
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.CmsCat true "更新CmsCat"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /cmsCat/updateCmsCat [put]
func (cmsCatApi *CmsCatApi) UpdateCmsCat(c *gin.Context) {
	var cmsCat autocode.CmsCat
	_ = c.ShouldBindJSON(&cmsCat)
	if err := cmsCatService.UpdateCmsCat(cmsCat); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCmsCat 用id查询CmsCat
// @Tags CmsCat
// @Summary 用id查询CmsCat
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.CmsCat true "用id查询CmsCat"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /cmsCat/findCmsCat [get]
func (cmsCatApi *CmsCatApi) FindCmsCat(c *gin.Context) {
	var cmsCat autocode.CmsCat
	_ = c.ShouldBindQuery(&cmsCat)
	if err, recmsCat := cmsCatService.GetCmsCat(cmsCat.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recmsCat": recmsCat}, c)
	}
}

// GetCmsCatList 分页获取CmsCat列表
// @Tags CmsCat
// @Summary 分页获取CmsCat列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.CmsCatSearch true "分页获取CmsCat列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cmsCat/getCmsCatList [get]
func (cmsCatApi *CmsCatApi) GetCmsCatList(c *gin.Context) {
	createdAtBetween, _ := c.GetQueryArray("createdAtBetween[]")

	var pageInfo autocodeReq.CmsCatSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := cmsCatService.GetCmsCatInfoList(pageInfo,createdAtBetween); err != nil {
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
// @Param data body autocode.CmsAd true "快速更新"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router  /cmsCat/quickEdit [post] 
func (cmsCatApi *CmsCatApi) QuickEdit(c *gin.Context) {
	var quickEdit request.QuickEdit
	_ = c.ShouldBindJSON(&quickEdit)
	quickEdit.Table = "cms_cat"
	//var_dump.Dump(quickEdit)
	if err := commDbService.QuickEdit(quickEdit); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}