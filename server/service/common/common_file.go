package common

import (
	"mime/multipart"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/upload"
)

type CommonFileService struct {
}

//@author: [ljd]
//@function: UploadFile 上传文件
//@description: 根据配置文件判断是文件上传到本地或者七牛云
//@param: header *multipart.FileHeader,basicFile autocode.BasicFile
//@return: err error, file model.ExaFileUploadAndDownload

func (e *CommonFileService) UploadFile(header *multipart.FileHeader, basicFile autocode.BasicFile) (err error, file autocode.BasicFile) {
	oss := upload.NewOss()

	//var dictSer = new(systemService.DictionaryDetailService)
	//moduleName, _ = dictSer.GetNameByValue(13, module)
	var module int
	module = *basicFile.Module
	filePath, guid, uploadErr := oss.UploadFile(header, module, 2)
	if uploadErr != nil {
		panic(err)
	}
	//是否保存到数据库表 basic_file
	//s := strings.Split(header.Filename, ".")
	basicFile.Guid = guid
	basicFile.Path = filePath
	basicFile.Name = header.Filename

	driver := 0
	switch global.GVA_CONFIG.System.OssType {
	case "local":
		driver = 0
	case "qiniu":
		driver = 1
	case "tencent-cos":
		driver = 2
	case "aliyun-oss":
		driver = 3
	default:
		driver = 4
	}
	basicFile.Driver = &driver
	err = global.GVA_DB.Create(&basicFile).Error
	return err, basicFile
}

func (e *CommonFileService) GetPathByGuid(guid string) (err error, path string) {
	//basicFS := new(autocodeSev.BasicFileService)
	// 先从 redis 获取 ，没有则读取数据库 ，缓存2小时
	if utils.IsEmpty(guid) {
		return nil, ""
	}

	basicFile := autocode.BasicFile{}
	err = global.GVA_DB.Where("guid = ?", guid).First(&basicFile).Error
	if err != nil {
		return err, ""
	}
	path = basicFile.Path
	if !utils.IsEmpty(path) {
		path = global.GVA_CONFIG.Local.BaseUrl + path
	}
	return err, path
}
