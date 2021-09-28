package common

import (
	"mime/multipart"

	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	autocodeSev "github.com/flipped-aurora/gin-vue-admin/server/service/autocode"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/upload"
)

type CommonFileService struct {
}

//@author: [ljd]
//@function: UploadFile 上传文件
//@description: 根据配置文件判断是文件上传到本地或者七牛云
//@param: header *multipart.FileHeader, noSave string
//@return: err error, file model.ExaFileUploadAndDownload

func (e *CommonFileService) UploadFile(header *multipart.FileHeader, basicFile autocode.BasicFile) (err error, file autocode.BasicFile) {
	oss := upload.NewOss()

	//var dictSer = new(systemService.DictionaryDetailService)
	//moduleName, _ = dictSer.GetNameByValue(13, module)
	var module int
	module = *basicFile.Module
	filePath, guid, uploadErr := oss.UploadFile(header, module)
	if uploadErr != nil {
		panic(err)
	}
	//是否保存到数据库表 basic_file
	//s := strings.Split(header.Filename, ".")
	basicFile.Guid = guid
	basicFile.Path = filePath
	basicFile.Name = header.Filename

	basicFS := new(autocodeSev.BasicFileService)
	return basicFS.CreateBasicFile(basicFile), basicFile

}
