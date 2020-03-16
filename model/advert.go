package model

import (
	"ccshop/errorcode"
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type Advert struct {
	ID           int       `json:"id" gorm:"AUTO_INCREMENT;primary_key;"`
	Name         string    `json:"name"`
	Code         string    `json:"code" gorm:"unique"`
	EnabledCycle int       `json:"enabled_cycle"`
	IsEnabled    int       `json:"is_enabled"`
	FromDate     time.Time `json:"from_date"`
	FromDateTime string    `json:"from_date_time" gorm:"-"`
	ExpireTime   int64     `json:"expire" gorm:"-"`
	ToDate       time.Time `json:"to_date"`
	UpdatedAt    time.Time `json:"-"`
	CreatedAt    time.Time `json:"-"`
}

func (a *Advert) TableName() string {
	return "cc_advertisings"
}

func (a *Advert) FetchAdvert(db *gorm.DB, condition map[string]interface{}) (*Advert, errorcode.Code) {
	advert := &Advert{}
	err := db.Table(a.TableName()).Where(condition).First(advert).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errorcode.OK
		}
		return nil, errorcode.DataFailed
	}
	//a.FromDateTime = a.FromDate.Unix() "2006-01-02T15:04:05+08:00"
	//parse, err := time.Parse(time.RFC3339, "2000-01-01T23:59:59+08:00")
	//fmt.Println(a.FromDate.Unix())
	stamp, err := time.Parse("2006-01-02T15:04:05+08:00", "2021-01-01T23:59:59+08:00")
	fmt.Println(stamp.Unix(), err)
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	timeNow := time.Now().Format("2006-01-02 15:04:05")
	a.FromDateTime = timeNow
	return advert, errorcode.OK
}
