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

func (as *AdvertService) FetchAdvertingFromModel(code string) (*model.Advert, errorcode.Code) {
	condition := map[string]interface{}{
		"code":code,
	}
	return as.Advert.FetchAdvert(as.Db, condition)
}

func (as *AdvertService) FetchAdvertingFromOrm(code string) (*model.Advert, errorcode.Code) {
	condition := map[string]interface{}{
		"code":code,
	}
	return as.Advert.FetchAdvertFromOrm(as.Db, condition)
}



func NewAdvertService() *AdvertService {
	return &AdvertService{Db: database.Db()}
}


