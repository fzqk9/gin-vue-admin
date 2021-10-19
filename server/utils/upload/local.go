package upload

import (
	"errors"
	"io"
	"mime/multipart"
	"os"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"go.uber.org/zap"
)

type Local struct{}

//@author: [piexlmax](https://github.com/piexlmax)
//@author: [ccfish86](https://github.com/ccfish86)
//@author: [SliverHorn](https://github.com/SliverHorn)
//@object: *Local
//@function: UploadFile
//@description: 上传文件
//@param: file *multipart.FileHeader
//@return: string, string, error

func (*Local) UploadFile(file *multipart.FileHeader, module int, userType int) (string, string, error) {
	// 读取文件后缀
	ext := path.Ext(file.Filename)
	// 读取文件名并加密
	//name := strings.TrimSuffix(file.Filename, ext)
	//name = utils.MD5V([]byte(name))

	name := utils.GUID()
	// 拼接新文件名
	// filename := name + "_" + time.Now().Format("20060102150405") + ext

	//filename := name + "_" + time.Now().Format("20060102150405") + ext
	var path string
	if userType == 1 { //管理用户
		path = global.GVA_CONFIG.Local.Path + "/" + strconv.Itoa(module) +
			"/" + time.Now().Format("20060102")
	} else if userType == 2 { //普通用户
		path = global.GVA_CONFIG.Local.PathUser + "/" + strconv.Itoa(module) +
			"/" + time.Now().Format("20060102")
	}

	// 尝试创建此路径
	//mkdirErr := os.MkdirAll(global.GVA_CONFIG.Local.Path, os.ModePerm)
	mkdirErr := os.MkdirAll(path, os.ModePerm)
	if mkdirErr != nil {
		global.GVA_LOG.Error("function os.MkdirAll() Filed", zap.Any("err", mkdirErr.Error()))
		return "", "", errors.New("function os.MkdirAll() Filed, err:" + mkdirErr.Error())
	}
	// 拼接路径和文件名
	//path := global.GVA_CONFIG.Local.Path + "/" + filename
	path = path + "/" + name + ext

	f, openError := file.Open() // 读取文件
	if openError != nil {
		global.GVA_LOG.Error("function file.Open() Filed", zap.Any("err", openError.Error()))
		return "", "", errors.New("function file.Open() Filed, err:" + openError.Error())
	}
	defer f.Close() // 创建文件 defer 关闭

	out, createErr := os.Create(path)
	if createErr != nil {
		global.GVA_LOG.Error("function os.Create() Filed", zap.Any("err", createErr.Error()))
		return "", "", errors.New("function os.Create() Filed, err:" + createErr.Error())
	}
	defer out.Close() // 创建文件 defer 关闭

	_, copyErr := io.Copy(out, f) // 传输（拷贝）文件
	if copyErr != nil {
		global.GVA_LOG.Error("function io.Copy() Filed", zap.Any("err", copyErr.Error()))
		return "", "", errors.New("function io.Copy() Filed, err:" + copyErr.Error())
	}
	return path, name, nil
}

//@author: [piexlmax](https://github.com/piexlmax)
//@author: [ccfish86](https://github.com/ccfish86)
//@author: [SliverHorn](https://github.com/SliverHorn)
//@object: *Local
//@function: DeleteFile
//@description: 删除文件
//@param: key string
//@return: error

func (*Local) DeleteFile(key string) error {
	p := global.GVA_CONFIG.Local.Path + "/" + key
	if strings.Contains(p, global.GVA_CONFIG.Local.Path) {
		if err := os.Remove(p); err != nil {
			return errors.New("本地文件删除失败, err:" + err.Error())
		}
	}
	return nil
}
