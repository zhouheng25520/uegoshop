package model

import (
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


func (ai *AdvertItem) FetchList(db *gorm.DB, condition map[string]interface{}) ([]*AdvertItem, error) {
	var itemList []*AdvertItem
	err := db.Model(ai).Where(condition).Order("sort asc").Find(itemList).Error
	if err != nil {
		fmt.Println("query advert items has failed , err:", err)
		return itemList, err
	}
	return itemList, nil
}
