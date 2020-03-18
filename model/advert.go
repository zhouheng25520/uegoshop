package model

import (
	"ccshop/common"
	"ccshop/errorcode"
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
	Items        []*AdvertItem `json:"items" gorm:"ForeignKey:AdID;ASSOCIATION_FOREIGNKEY:ID"`
	ToDate       time.Time     `json:"-"`
	UpdatedAt    time.Time     `json:"-"`
	CreatedAt    time.Time     `json:"-"`
}

func (a *Advert) TableName() string {
	return "cc_advertisings"
}

func (a *Advert) FetchAdvertFromOrm(db *gorm.DB, condition map[string]interface{}) (*Advert, errorcode.Code) {
	advert := &Advert{}
	err := db.Model(a).Where(condition).First(advert).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errorcode.OK
		}
		return nil, errorcode.DataFailed
	}
	fromDateTime, err := common.FormatTime(common.CENTRAL_STANDARD_TIME_LAYOUT, advert.FromDate.String())
	if err != nil {
		return advert, errorcode.TimeStampFormat
	}
	advert.FromDateTime = fromDateTime.Unix()

	endTime, err := common.FormatTime(common.CENTRAL_STANDARD_TIME_LAYOUT, advert.ToDate.String())
	if err != nil {
		return advert, errorcode.TimeStampFormat
	}

	if time.Now().Unix()-endTime.Unix() < 0 {
		advert.ExpireTime = time.Now().Unix() - endTime.Unix()
	}

	items := &AdvertItem{}
	advertItems, ecode := items.FetchList(db, map[string]interface{}{
		"is_enabled": 1,
	})
	if ecode.Code() != errorcode.OK.Code() {
		// 如果code非正常code, 将items置为空
		advertItems = []*AdvertItem{}
	}
	advert.Items = advertItems

	return advert, errorcode.OK
}



func (a *Advert) FetchAdvert(db *gorm.DB, condition map[string]interface{}) (*Advert, errorcode.Code) {
	advert := &Advert{}
	err := db.Model(a).Where(condition).First(advert).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errorcode.OK
		}
		return nil, errorcode.DataFailed
	}
	fromDateTime, err := common.FormatTime(common.CENTRAL_STANDARD_TIME_LAYOUT, advert.FromDate.String())
	if err != nil {
		return advert, errorcode.TimeStampFormat
	}
	advert.FromDateTime = fromDateTime.Unix()

	endTime, err := common.FormatTime(common.CENTRAL_STANDARD_TIME_LAYOUT, advert.ToDate.String())
	if err != nil {
		return advert, errorcode.TimeStampFormat
	}

	if time.Now().Unix()-endTime.Unix() < 0 {
		advert.ExpireTime = time.Now().Unix() - endTime.Unix()
	}

	items := &AdvertItem{}
	advertItems, ecode := items.FetchList(db, map[string]interface{}{
		"is_enabled": 1,
	})
	if ecode.Code() != errorcode.OK.Code() {
		// 如果code非正常code, 将items置为空
		advertItems = []*AdvertItem{}
	}
	advert.Items = advertItems

	return advert, errorcode.OK
}
