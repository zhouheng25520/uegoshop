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
		"code": code,
	}
	return as.Advert.FetchAdvert(as.Db, condition)
}

func (as *AdvertService) FetchAdvertingFromOrm(code string) (*model.Advert, errorcode.Code) {
	//condition := map[string]interface{}{
	//	"code":       code,
	//	"is_enabled": 1,
	//}
	return as.Advert.FetchAdvertFromOrm(as.Db, model.Advert{Code: code, IsEnabled: 1})
}

func NewAdvertService() *AdvertService {
	return &AdvertService{Db: database.Db()}
}
