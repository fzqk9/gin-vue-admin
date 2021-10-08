package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
)

type CmsCatArtService struct {
}

// CreateCmsCatArt 创建CmsCatArt记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmsCatArtService *CmsCatArtService) CreateCmsCatArt(cmsCatArt autocode.CmsCatArt) (err error) {
	err = global.GVA_DB.Create(&cmsCatArt).Error
	return err
}

// DeleteCmsCatArt 删除CmsCatArt记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmsCatArtService *CmsCatArtService)DeleteCmsCatArt(cmsCatArt autocode.CmsCatArt) (err error) {
	err = global.GVA_DB.Delete(&cmsCatArt).Error
	return err
}

// DeleteCmsCatArtByIds 批量删除CmsCatArt记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmsCatArtService *CmsCatArtService)DeleteCmsCatArtByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.CmsCatArt{},"id in ?",ids.Ids).Error
	return err
}

// UpdateCmsCatArt 更新CmsCatArt记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmsCatArtService *CmsCatArtService)UpdateCmsCatArt(cmsCatArt autocode.CmsCatArt) (err error) {
	err = global.GVA_DB.Save(&cmsCatArt).Error
	return err
}

// GetCmsCatArt 根据id获取CmsCatArt记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmsCatArtService *CmsCatArtService)GetCmsCatArt(id uint) (err error, cmsCatArt autocode.CmsCatArt) {
	err = global.GVA_DB.Where("id = ?", id).First(&cmsCatArt).Error
	return
}

// GetCmsCatArtInfoList 分页获取CmsCatArt记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmsCatArtService *CmsCatArtService)GetCmsCatArtInfoList(info autoCodeReq.CmsCatArtSearch, createdAtBetween []string) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    //修改 by ljd  增加查询排序 
    order := info.OrderKey
	desc := info.OrderDesc
    // 创建db
	db := global.GVA_DB.Model(&autocode.CmsCatArt{})
    var cmsCatArts []autocode.CmsCatArt

    //修改 by ljd  
    if info.ID > 0 {
		db = db.Where("`id` = ?", info.ID)
	}
	if createdAtBetween != nil && len(createdAtBetween) > 0 {
		db = db.Where("`created_at` BETWEEN ? AND ?", createdAtBetween[0], createdAtBetween[1])
	}

    // 如果有条件搜索 下方会自动创建搜索语句
    if info.ArticleId != nil {
        db = db.Where("`article_id` = ?",info.ArticleId)
    }
    if info.CatId != nil {
        db = db.Where("`cat_id` = ?",info.CatId)
    }
    if info.Status != nil {
        db = db.Where("`status` = ?",info.Status)
    }
	err = db.Count(&total).Error
	//err = db.Limit(limit).Offset(offset).Find(&cmsCatArts).Error
    //修改 by ljd  增加查询排序 
     if order != "" {
		var OrderStr string
		if desc {
			OrderStr = order + " desc"
		} else {
			OrderStr = order
		}
        err = db.Order(OrderStr).Limit(limit).Offset(offset).Find(&cmsCatArts).Error 
	} else {
		err = db.Limit(limit).Offset(offset).Find(&cmsCatArts).Error
	}
	return err, cmsCatArts, total
}
