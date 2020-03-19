package v1

import (
	"ccshop/common"
	"ccshop/database"
	"ccshop/errorcode"
	"ccshop/response"
	"ccshop/service"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Advert struct {
	Db *gorm.DB
}

func NewAdvert(db *gorm.DB) *Advert {
	return &Advert{Db: database.Db()}
}

// register advert struct
func AdvertRegister(router *gin.RouterGroup) {
	advert := Advert{}
	router.GET("/info", advert.GetInfo)
	router.GET("/info-orm",advert.GetInfoFromOrm)
}

func (advert *Advert) GetInfoFromOrm(c *gin.Context)  {
	code := c.DefaultQuery("code", "")
	resp := response.NewResponse(c)
	// todo validator
	if common.IsEmptyString(code) {
		resp.Body.Code = errorcode.ParamsInvalid.Code()
		resp.ReturnJsonError()
		return
	}

	advertService := service.NewAdvertService(advert.Db)
	adverts, ecode := advertService.FetchAdvertingFromOrm(code)
	if ecode != errorcode.OK {
		resp.Body.Code = ecode.Code()
		resp.ReturnJsonError()
		return
	}
	resp.ReturnJsonSuccess(adverts)
}

// return advert information, if has error well be return error message
func (advert *Advert) GetInfo(c *gin.Context) {
	code := c.DefaultQuery("code", "")
	resp := response.NewResponse(c)
	if common.IsEmptyString(code) {
		resp.Body.Code = errorcode.ParamsInvalid.Code()
		resp.ReturnJsonError()
		return
	}

	advertService := service.NewAdvertService(advert.Db)
	adverts, ecode := advertService.FetchAdvertingFromModel(code)
	if ecode != errorcode.OK {
		resp.Body.Code = ecode.Code()
		resp.ReturnJsonError()
		return
	}
	resp.ReturnJsonSuccess(adverts)
}
