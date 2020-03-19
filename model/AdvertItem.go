package model

import (
	"ccshop/errorcode"
	"fmt"
	"github.com/jinzhu/gorm"
)

type AdvertItem struct {
	ID        int    `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	AdID      int    `json:"ad_id" gorm:"default:0"`
	Name      string `json:"name" gorm:"default:''"`
	Link      string `json:"link" gorm:"default:''"`
	Sort      int    `json:"sort" gorm:"default:1"`
	IsEnabled int    `json:"is_enabled" gorm:"default:1"`
	Describe  string `json:"describe"`
}

func (ai *AdvertItem) TableName() string {
	return "cc_advertising_items"
}


func (ai *AdvertItem) FetchList(db *gorm.DB, condition interface{}) ([]*AdvertItem, errorcode.Code) {
	var AdvertItems []*AdvertItem
	ecode := errorcode.OK
	err := db.Model(&AdvertItem{}).Where(condition).Order("sort asc").Find(&AdvertItems).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return AdvertItems, ecode
		}
		fmt.Println("got advert items has err :", err)
		return AdvertItems, errorcode.DataFailed
	}
	return AdvertItems, ecode
}
