package model

import (
	"uegoshop/common"
	"uegoshop/errorcode"
	"github.com/jinzhu/gorm"
	"time"
)

type Advert struct {
	ID           int           `json:"id" gorm:"AUTO_INCREMENT;primary_key;"`
	Name         string        `json:"name"`
	Code         string        `json:"code" gorm:"unique"`
	EnabledCycle int           `json:"enabled_cycle"`
	IsEnabled    int           `json:"is_enabled"`
	FromDate     time.Time     `json:"-"`
	FromDateTime int64         `json:"from_date_time" gorm:"-"`
	ExpireTime   int64         `json:"expired" gorm:"-"`
	Items        []*AdvertItem `json:"items" gorm:"FOREIGNKEY:AdID;ASSOCIATION_FOREIGNKEY:ID"`
	ToDate       time.Time     `json:"-"`
	UpdatedAt    time.Time     `json:"-"`
	CreatedAt    time.Time     `json:"-"`
}

func (a *Advert) TableName() string {
	return "advertisements"
}

func (a *Advert) FetchAdvertFromOrm(db *gorm.DB, condition interface{}) (*Advert, errorcode.Code) {
	advert := &Advert{}

	/*
	 |-------------------------------------------------------------------------
	 | 关联模型 -- Related方式实现, 并按字段排序
	 |-------------------------------------------------------------------------
	*/
	recordDB := db.Where(condition).First(&advert)
	if recordDB.RecordNotFound() == true {
		return nil, errorcode.OK
	}
	result := db.Model(&advert).Related(&advert.Items, "Items").
		Where(AdvertItem{AdID: advert.ID, IsEnabled: 1}).
		Order("sort asc").Find(&advert.Items)
	/*
	 |--------------------------------------------------------------------------
	*/

	/*
	 |-------------------------------------------------------------------------
	 | 关联模型 -- association实现方式 并按字段排序
	 |-------------------------------------------------------------------------
	 | recordDB := db.Where(condition).First(&advert)
	 | if recordDB.RecordNotFound() == true {
	 |    return nil, errorcode.OK
	 | }
	 | db.Model(&advert).Association("Items")
	 | result := db.Where(AdvertItem{AdID:advert.ID,IsEnabled:1}).Order("sort asc").Find(&advert.Items)
	 |--------------------------------------------------------------------------
	*/

	/*
	 |-------------------------------------------------------------------------
	 | 预加载数据查询 -- Preload实现方式, 并按字段排序
	 |-------------------------------------------------------------------------
	 | recordDB := db.Where(condition).First(&advert)
	 | if recordDB.RecordNotFound() == true {
	 |	  return nil, errorcode.OK
	 | }
	 | result = recordDB.Preload("Items", func(db *gorm.DB) *gorm.DB {
	 |		  item := &AdvertItem{}
	 |		  return db.Where("is_enabled = ? ", "1").
	 |				Order(fmt.Sprintf("%s.sort asc", item.TableName()))
	 |	  }).Find(&advert)
	 |--------------------------------------------------------------------------
	*/

	if err := result.Error; err != nil {
		return nil, errorcode.DataFailed
	}

	fromDateTime, err := common.FormatTime(common.CENTRAL_STANDARD_TIME_LAYOUT,
		advert.FromDate.String())
	if err != nil {
		return advert, errorcode.TimeStampFormat
	}
	advert.FromDateTime = fromDateTime.Unix()

	endTime, err := common.FormatTime(common.CENTRAL_STANDARD_TIME_LAYOUT,
		advert.ToDate.String())
	if err != nil {
		return advert, errorcode.TimeStampFormat
	}

	if time.Now().Unix()-endTime.Unix() < 0 {
		advert.ExpireTime = time.Now().Unix() - endTime.Unix()
	}

	return advert, errorcode.OK
}
