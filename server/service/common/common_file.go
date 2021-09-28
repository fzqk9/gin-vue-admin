package common

import (
	"mime/multipart"
	"strings"

	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	autocodeSev "github.com/flipped-aurora/gin-vue-admin/server/service/autocode"
	systemService "github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/upload"
)

type CommonFileService struct {
}

//@author: [ljd]
//@function: UploadFile 上传文件
//@description: 根据配置文件判断是文件上传到本地或者七牛云
//@param: header *multipart.FileHeader, noSave string
//@return: err error, file model.ExaFileUploadAndDownload

func (e *CommonFileService) UploadFile(header *multipart.FileHeader, noSave string, media_type *int, module *int) (err error, file autocode.BasicFile) {
	oss := upload.NewOss()

	var moduleName string
	var dictSer = new(systemService.DictionaryDetailService)
	var moduleVal int
	moduleVal = *module
	moduleName, _ = dictSer.GetNameByValue(13, moduleVal)

	filePath, guid, uploadErr := oss.UploadFile(header, moduleName)
	if uploadErr != nil {
		panic(err)
	}
	//是否保存到数据库表 basic_file
	if noSave == "0" {
		s := strings.Split(header.Filename, ".")
		// f := request.FileUpload{
		// 	Url:  filePath,
		// 	Name: header.Filename,
		// 	Tag:  s[len(s)-1],
		// 	Key:  key,
		// }
		// 保存到数据库表 basic_file
		basicFile := autocode.BasicFile{
			Path:      filePath,
			Guid:      guid,
			Name:      header.Filename,
			Ext:       s[len(s)-1],
			MediaType: media_type,
			Module:    module,
		}
		basicFS := new(autocodeSev.BasicFileService)
		return basicFS.CreateBasicFile(basicFile), basicFile
	}
	return
}
