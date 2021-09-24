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

type BasicFileApi struct {
}

var basicFileService = service.ServiceGroupApp.AutoCodeServiceGroup.BasicFileService
 

// CreateBasicFile 创建BasicFile
// @Tags BasicFile
// @Summary 创建BasicFile
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.BasicFile true "创建BasicFile"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /basicFile/createBasicFile [post]
func (basicFileApi *BasicFileApi) CreateBasicFile(c *gin.Context) {
	var basicFile autocode.BasicFile
	_ = c.ShouldBindJSON(&basicFile)
	if err := basicFileService.CreateBasicFile(basicFile); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteBasicFile 删除BasicFile
// @Tags BasicFile
// @Summary 删除BasicFile
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.BasicFile true "删除BasicFile"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /basicFile/deleteBasicFile [delete]
func (basicFileApi *BasicFileApi) DeleteBasicFile(c *gin.Context) {
	var basicFile autocode.BasicFile
	_ = c.ShouldBindJSON(&basicFile)
	if err := basicFileService.DeleteBasicFile(basicFile); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteBasicFileByIds 批量删除BasicFile
// @Tags BasicFile
// @Summary 批量删除BasicFile
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除BasicFile"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /basicFile/deleteBasicFileByIds [delete]
func (basicFileApi *BasicFileApi) DeleteBasicFileByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := basicFileService.DeleteBasicFileByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateBasicFile 更新BasicFile
// @Tags BasicFile
// @Summary 更新BasicFile
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.BasicFile true "更新BasicFile"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /basicFile/updateBasicFile [put]
func (basicFileApi *BasicFileApi) UpdateBasicFile(c *gin.Context) {
	var basicFile autocode.BasicFile
	_ = c.ShouldBindJSON(&basicFile)
	if err := basicFileService.UpdateBasicFile(basicFile); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindBasicFile 用id查询BasicFile
// @Tags BasicFile
// @Summary 用id查询BasicFile
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.BasicFile true "用id查询BasicFile"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /basicFile/findBasicFile [get]
func (basicFileApi *BasicFileApi) FindBasicFile(c *gin.Context) {
	var basicFile autocode.BasicFile
	_ = c.ShouldBindQuery(&basicFile)
	if err, rebasicFile := basicFileService.GetBasicFile(basicFile.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rebasicFile": rebasicFile}, c)
	}
}

// GetBasicFileList 分页获取BasicFile列表
// @Tags BasicFile
// @Summary 分页获取BasicFile列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.BasicFileSearch true "分页获取BasicFile列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /basicFile/getBasicFileList [get]
func (basicFileApi *BasicFileApi) GetBasicFileList(c *gin.Context) {
	createdAtBetween, _ := c.GetQueryArray("createdAtBetween[]")

	var pageInfo autocodeReq.BasicFileSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := basicFileService.GetBasicFileInfoList(pageInfo,createdAtBetween); err != nil {
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
// @Param data body autocode.BasicFile true "快速更新BasicFile" 
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router  /basicFile/quickEdit [post] 
func (basicFileApi *BasicFileApi) QuickEdit(c *gin.Context) {
	var quickEdit request.QuickEdit
	_ = c.ShouldBindJSON(&quickEdit)
	quickEdit.Table = "basic_file"
	//var_dump.Dump(quickEdit)
	if err := commDbService.QuickEdit(quickEdit); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}