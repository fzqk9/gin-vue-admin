package common

import (
	"fmt"
	"strconv"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CommonFileApi struct {
}

var commonFileService = service.ServiceGroupApp.CommonServiceGroup.CommonFileService

// @Tags UploadFile
// @Summary 上传文件
// @Security ApiKeyAuth
// @accept multipart/form-data
// @Produce  application/json
// @Param file formData file true "上传文件 "
// @Success 200 {string} string "{"success":true,"data":{},"msg":"上传成功"}"
// @Router /fileUploadAndDownload/upload [post]
func (commonFileApi *CommonFileApi) UploadFile(c *gin.Context) {
	module := c.DefaultPostForm("module", "0")
	media_type := c.DefaultPostForm("media_type", "0") // 此方法可以设置默认值
	//noSave := c.DefaultQuery("noSave", "0")
	//media_type := c.DefaultQuery("media_type", "0")
	//module := c.DefaultQuery("module", "0")
	fmt.Println("media_type=%s", media_type)
	fmt.Println("module=%s", module)
	i, err := strconv.Atoi(media_type)
	j, err := strconv.Atoi(module)

	var file autocode.BasicFile
	var pMediaType *int
	var pModule *int
	pMediaType = &i
	pModule = &j
	module = c.DefaultQuery("module", "0")
	_, header, err := c.Request.FormFile("file")
	if err != nil {
		global.GVA_LOG.Error("接收文件失败!", zap.Any("err", err))
		response.FailWithMessage("接收文件失败", c)
		return
	}
	//var basicFile autocode.BasicFile
	//basicFileService.CreateBasicFile(basicFile)
	//err, file = commFileService.UploadFile(header, noSave)
	err, file = commonFileService.UploadFile(header, noSave, pMediaType, pModule) // 文件上传后拿到文件路径
	//err, file = basicFileService.CreateBasicFile(header, noSave) // 文件上传后拿到文件路径
	if err != nil {
		global.GVA_LOG.Error("修改数据库链接失败!", zap.Any("err", err))
		response.FailWithMessage("修改数据库链接失败", c)
		return
	}
	response.OkWithDetailed(file, "上传成功", c)
	//response.OkWithDetailed({File: file}, "上传成功", c)

}
