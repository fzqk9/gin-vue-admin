package common

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type CommDbService struct {
}

// QuickEdit 更新QuickEdit记录
// Author [Linjd]
func (CommDbService *CommDbService) QuickEdit(quickEdit request.QuickEdit) (err error) {
	err = global.GVA_DB.Exec("UPDATE `"+quickEdit.Table+"` SET `"+quickEdit.Field+"`=? WHERE `id` = ?", quickEdit.Value, quickEdit.Id).Error
	//err = global.GVA_DB.(&quickEdit).Error
	return err
}
