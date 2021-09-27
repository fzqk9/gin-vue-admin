package common

import (
	"mime/multipart"
	"strings"

	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	autocodeSev "github.com/flipped-aurora/gin-vue-admin/server/service/autocode"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/upload"
)

type CommFileService struct {
}

//@author: [ljd]
//@function: UploadFile 上传文件
//@description: 根据配置文件判断是文件上传到本地或者七牛云
//@param: header *multipart.FileHeader, noSave string
//@return: err error, file model.ExaFileUploadAndDownload

func (e *CommFileService) UploadFile(header *multipart.FileHeader, noSave string) (err error, file request.FileUpload) {
	oss := upload.NewOss()
	filePath, key, uploadErr := oss.UploadFile(header)
	if uploadErr != nil {
		panic(err)
	}
	//是否保存到数据库表 basic_file
	if noSave == "0" {
		s := strings.Split(header.Filename, ".")
		f := request.FileUpload{
			Url:  filePath,
			Name: header.Filename,
			Tag:  s[len(s)-1],
			Key:  key,
		}
		// 保存到数据库表 basic_file
		basicFile := autocode.BasicFile{
			Path: filePath,
			Guid: utils.GUID(),
			Name: header.Filename,
			Ext:  s[len(s)-1],
		}
		basicFS := new(autocodeSev.BasicFileService)
		return basicFS.CreateBasicFile(basicFile), f
	}
	return
}
