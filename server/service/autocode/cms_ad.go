package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
    "github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type CmsAdService struct {
}

// CreateCmsAd 创建CmsAd记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmsAdService *CmsAdService) CreateCmsAd(cmsAd autocode.CmsAd) (err error) {
	err = global.GVA_DB.Create(&cmsAd).Error
	return err
}

// DeleteCmsAd 删除CmsAd记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmsAdService *CmsAdService)DeleteCmsAd(cmsAd autocode.CmsAd) (err error) {
	err = global.GVA_DB.Delete(&cmsAd).Error
	return err
}

// DeleteCmsAdByIds 批量删除CmsAd记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmsAdService *CmsAdService)DeleteCmsAdByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.CmsAd{},"id in ?",ids.Ids).Error
	return err
}

// UpdateCmsAd 更新CmsAd记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmsAdService *CmsAdService)UpdateCmsAd(cmsAd autocode.CmsAd) (err error) {
	err = global.GVA_DB.Save(&cmsAd).Error
	return err
}

// GetCmsAd 根据id获取CmsAd记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmsAdService *CmsAdService)GetCmsAd(id uint,fields string ) (err error, obj autocode.CmsAd) {
	err = global.GVA_DB.Where("id = ?", id).First(&obj).Error 
    if utils.IsEmpty(fields) {
        err = global.GVA_DB.Where("id = ?", id).First(&obj).Error 
        	} else {
        err = global.GVA_DB.Select(fields).Where("id = ?", id).First(&obj).Error  
	}

    obj.MapData = make(map[string]string) 
    if !utils.IsEmpty(obj.Image) {
        _,obj.MapData[obj.Image] = commFileService.GetPathByGuid(obj.Image)
    }  
    return err, obj
}

// GetCmsAdInfoList 分页获取CmsAd记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmsAdService *CmsAdService)GetCmsAdInfoList(info autoCodeReq.CmsAdSearch, createdAtBetween []string,fields string) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    //修改 by ljd  增加查询排序 
    order := info.OrderKey
	desc := info.OrderDesc
    // 创建db
	db := global.GVA_DB.Model(&autocode.CmsAd{})
    //var cmsAds []autocode.CmsAd
    var cmsAds []autocode.CmsAdMini

    //修改 by ljd  
    if info.ID > 0 {
		db = db.Where("`id` = ?", info.ID)
	}
	if createdAtBetween != nil && len(createdAtBetween) > 0 {
		db = db.Where("`created_at` BETWEEN ? AND ?", createdAtBetween[0], createdAtBetween[1])
	}

    // 如果有条件搜索 下方会自动创建搜索语句
    if info.MediaType != nil {
        db = db.Where("`media_type` = ?",info.MediaType)
    }
    if info.Status != nil {
        db = db.Where("`status` = ?",info.Status)
    }
	err = db.Count(&total).Error
	//err = db.Limit(limit).Offset(offset).Find(&cmsAds).Error
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
      err = db.Order(OrderStr).Limit(limit).Offset(offset).Find(&cmsAds).Error
    } else {
      err = db.Select(fields).Order(OrderStr).Limit(limit).Offset(offset).Find(&cmsAds).Error
    }         
     //更新图片path
	for i, v := range cmsAds {
	 v.MapData = make(map[string]string) 
        if !utils.IsEmpty(v.Image) {
            _, v.MapData[v.Image] = commFileService.GetPathByGuid(v.Image)
        }
	  cmsAds[i] = v
	}
	return err, cmsAds, total
}
