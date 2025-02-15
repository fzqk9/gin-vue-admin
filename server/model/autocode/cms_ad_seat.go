// 自动生成模板CmsAdSeat
package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// CmsAdSeat 结构体
// 如果含有time.Time 请自行import time包
type CmsAdSeat struct {
     CmsAdSeatMini
      Desc  string `json:"desc" form:"desc" gorm:"column:desc;comment:广告描述;type:varchar(500);"`
}

// CmsAdSeatMini 结构体
type CmsAdSeatMini struct {
      global.GVA_MODEL
      Name  string `json:"name" form:"name" gorm:"column:name;comment:位置名称;type:varchar(60);"`
      Width  *int `json:"width" form:"width" gorm:"column:width;comment:广告位宽度;type:smallint"`
      Height  *int `json:"height" form:"height" gorm:"column:height;comment:广告位高度;type:smallint"`
      Style  string `json:"style" form:"style" gorm:"column:style;comment:模板;type:text;"`
      Status  *int `json:"status" form:"status" gorm:"column:status;comment:状态:0关闭1开启;type:smallint"`

}


// TableName CmsAdSeat 表名
func (CmsAdSeat) TableName() string {
  return "cms_ad_seat"
}

