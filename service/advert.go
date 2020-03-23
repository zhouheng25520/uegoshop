package service

import (
	"ccshop/database"
	"ccshop/errorcode"
	"ccshop/model"
	"github.com/jinzhu/gorm"
)

type AdvertService struct {
	Db     *gorm.DB
	Advert *model.Advert
}

func (as *AdvertService) FetchAdverting(code string) (*model.Advert, errorcode.Code) {
	condition := model.Advert{
		Code: code,
		IsEnabled: 1,
	}
	return as.Advert.FetchAdvertFromOrm(as.Db, condition)
}

func NewAdvertService() *AdvertService {
	return &AdvertService{Db: database.Db()}
}
