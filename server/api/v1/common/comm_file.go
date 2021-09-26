package common

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/gin-gonic/gin"
)

type FileApi struct {
}

// @Tags UploadFile
// @Summary 上传文件
// @Security ApiKeyAuth
// @accept multipart/form-data
// @Produce  application/json
// @Param file formData file true "上传文件 "
// @Success 200 {string} string "{"success":true,"data":{},"msg":"上传成功"}"
// @Router /fileUploadAndDownload/upload [post]
func (u *FileApi) UploadFile(c *gin.Context) {
	//var file request.FileUpload
	var basicFile autocode.BasicFile
	// noSave := c.DefaultQuery("noSave", "0")
	// _, header, err := c.Request.FormFile("file")
	// if err != nil {
	// 	global.GVA_LOG.Error("接收文件失败!", zap.Any("err", err))
	// 	response.FailWithMessage("接收文件失败", c)
	// 	return
	// }
	basicFileService.CreateBasicFile(basicFile)
	//	err, file = fileUploadAndDownloadService.UploadFile(header, noSave) // 文件上传后拿到文件路径
	// err, file = basicFileService.CreateBasicFile(header, noSave) // 文件上传后拿到文件路径
	// if err != nil {
	// 	global.GVA_LOG.Error("修改数据库链接失败!", zap.Any("err", err))
	// 	response.FailWithMessage("修改数据库链接失败", c)
	// 	return
	// }
	//response.OkWithDetailed(exampleRes.ExaFileResponse{File: file}, "上传成功", c)
	response.OkWithDetailed(nil, "上传成功", c)
}
