package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
    "github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type CmsCatService struct {
}

// CreateCmsCat 创建CmsCat记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmsCatService *CmsCatService) CreateCmsCat(cmsCat autocode.CmsCat) (err error) {
	err = global.GVA_DB.Create(&cmsCat).Error
	return err
}

// DeleteCmsCat 删除CmsCat记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmsCatService *CmsCatService)DeleteCmsCat(cmsCat autocode.CmsCat) (err error) {
	err = global.GVA_DB.Delete(&cmsCat).Error
	return err
}

// DeleteCmsCatByIds 批量删除CmsCat记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmsCatService *CmsCatService)DeleteCmsCatByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.CmsCat{},"id in ?",ids.Ids).Error
	return err
}

// UpdateCmsCat 更新CmsCat记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmsCatService *CmsCatService)UpdateCmsCat(cmsCat autocode.CmsCat) (err error) {
	err = global.GVA_DB.Save(&cmsCat).Error
	return err
}

// GetCmsCat 根据id获取CmsCat记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmsCatService *CmsCatService)GetCmsCat(id uint,fields string ) (err error, obj autocode.CmsCat) {
	err = global.GVA_DB.Where("id = ?", id).First(&obj).Error 
    if utils.IsEmpty(fields) {
        err = global.GVA_DB.Where("id = ?", id).First(&obj).Error 
        	} else {
        err = global.GVA_DB.Select(fields).Where("id = ?", id).First(&obj).Error  
	}

    obj.MapData = make(map[string]string) 
    if !utils.IsEmpty(obj.Thumb) {
        _,obj.MapData[obj.Thumb] = commFileService.GetPathByGuid(obj.Thumb)
    }  
    return err, obj
}

// GetCmsCatInfoList 分页获取CmsCat记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmsCatService *CmsCatService)GetCmsCatInfoList(info autoCodeReq.CmsCatSearch, createdAtBetween []string,fields string) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    //修改 by ljd  增加查询排序 
    order := info.OrderKey
	desc := info.OrderDesc
    // 创建db
	db := global.GVA_DB.Model(&autocode.CmsCat{})
    //var cmsCats []autocode.CmsCat
    var cmsCats []autocode.CmsCatMini

    //修改 by ljd  
    if info.ID > 0 {
		db = db.Where("`id` = ?", info.ID)
	}
	if createdAtBetween != nil && len(createdAtBetween) > 0 {
		db = db.Where("`created_at` BETWEEN ? AND ?", createdAtBetween[0], createdAtBetween[1])
	}

    // 如果有条件搜索 下方会自动创建搜索语句
    if info.BeSys != nil {
        db = db.Where("`be_sys` = ?",info.BeSys)
    }
    if info.GroupId != nil {
        db = db.Where("`group_id` = ?",info.GroupId)
    }
    if info.MediaType != nil {
        db = db.Where("`media_type` = ?",info.MediaType)
    }
    if info.Name != "" {
        db = db.Where("`name` LIKE ?","%"+ info.Name+"%")
    }
    if info.Desc != "" {
        db = db.Where("`desc` LIKE ?","%"+ info.Desc+"%")
    }
    if info.Keywords != "" {
        db = db.Where("`keywords` LIKE ?","%"+ info.Keywords+"%")
    }
    if info.Status != nil {
        db = db.Where("`status` = ?",info.Status)
    }
	err = db.Count(&total).Error
	//err = db.Limit(limit).Offset(offset).Find(&cmsCats).Error
    //修改 by ljd  增加查询排序 
     OrderStr := "id desc"
     if !utils.IsEmpty(order) { 
		if desc {
			OrderStr = order + " desc"
		} else {
			OrderStr = order
		} 
	}  
    if utils.IsEmpty(fields) {
      err = db.Order(OrderStr).Limit(limit).Offset(offset).Find(&cmsCats).Error
    } else {
      err = db.Select(fields).Order(OrderStr).Limit(limit).Offset(offset).Find(&cmsCats).Error
    }         
     //更新图片path
	for i, v := range cmsCats {
	 v.MapData = make(map[string]string) 
        if !utils.IsEmpty(v.Thumb) {
            _, v.MapData[v.Thumb] = commFileService.GetPathByGuid(v.Thumb)
        }
	  cmsCats[i] = v
	}
	return err, cmsCats, total
}
