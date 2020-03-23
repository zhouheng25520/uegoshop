package model

type AdvertItem struct {
	ID        int    `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	AdID      int    `json:"ad_id" gorm:"default:0;column:ad_id"`
	Name      string `json:"name" gorm:"default:''"`
	Link      string `json:"link" gorm:"default:''"`
	Sort      int    `json:"sort" gorm:"default:1"`
	IsEnabled int    `json:"is_enabled" gorm:"default:1"`
	Describe  string `json:"describe"`
}

func (ai *AdvertItem) TableName() string {
	return "cc_advertisement_items"
}
