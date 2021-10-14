package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
     "github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type CmsAdSeatService struct {
}

// CreateCmsAdSeat 创建CmsAdSeat记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmsAdSeatService *CmsAdSeatService) CreateCmsAdSeat(cmsAdSeat autocode.CmsAdSeat) (err error) {
	err = global.GVA_DB.Create(&cmsAdSeat).Error
	return err
}

// DeleteCmsAdSeat 删除CmsAdSeat记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmsAdSeatService *CmsAdSeatService)DeleteCmsAdSeat(cmsAdSeat autocode.CmsAdSeat) (err error) {
	err = global.GVA_DB.Delete(&cmsAdSeat).Error
	return err
}

// DeleteCmsAdSeatByIds 批量删除CmsAdSeat记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmsAdSeatService *CmsAdSeatService)DeleteCmsAdSeatByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.CmsAdSeat{},"id in ?",ids.Ids).Error
	return err
}

// UpdateCmsAdSeat 更新CmsAdSeat记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmsAdSeatService *CmsAdSeatService)UpdateCmsAdSeat(cmsAdSeat autocode.CmsAdSeat) (err error) {
	err = global.GVA_DB.Save(&cmsAdSeat).Error
	return err
}

// GetCmsAdSeat 根据id获取CmsAdSeat记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmsAdSeatService *CmsAdSeatService)GetCmsAdSeat(id uint) (err error, obj autocode.CmsAdSeat) {
	err = global.GVA_DB.Where("id = ?", id).First(&obj).Error 
    obj.MapData = make(map[string]string)  
    return err, obj
}

// GetCmsAdSeatInfoList 分页获取CmsAdSeat记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmsAdSeatService *CmsAdSeatService)GetCmsAdSeatInfoList(info autoCodeReq.CmsAdSeatSearch, createdAtBetween []string) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    //修改 by ljd  增加查询排序 
    order := info.OrderKey
	desc := info.OrderDesc
    // 创建db
	db := global.GVA_DB.Model(&autocode.CmsAdSeat{})
    var cmsAdSeats []autocode.CmsAdSeat

    //修改 by ljd  
    if info.ID > 0 {
		db = db.Where("`id` = ?", info.ID)
	}
	if createdAtBetween != nil && len(createdAtBetween) > 0 {
		db = db.Where("`created_at` BETWEEN ? AND ?", createdAtBetween[0], createdAtBetween[1])
	}

    // 如果有条件搜索 下方会自动创建搜索语句
    if info.Status != nil {
        db = db.Where("`status` = ?",info.Status)
    }
	err = db.Count(&total).Error
	//err = db.Limit(limit).Offset(offset).Find(&cmsAdSeats).Error
    //修改 by ljd  增加查询排序 
     OrderStr := "id desc"
     if !utils.IsEmpty(order) { 
		if desc {
			OrderStr = order + " desc"
		} else {
			OrderStr = order
		} 
	}  
     err = db.Order(OrderStr).Limit(limit).Offset(offset).Find(&cmsAdSeats).Error
     //更新图片path
	for i, v := range cmsAdSeats {
	 v.MapData = make(map[string]string)
	  cmsAdSeats[i] = v
	}
	return err, cmsAdSeats, total
}
