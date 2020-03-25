package v1

import (
	"ccshop/common"
	"ccshop/database"
	"ccshop/errorcode"
	"ccshop/response"
	"ccshop/service"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"time"
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
}
// return advert information, if has error well be return error message
func (advert *Advert) GetInfo(c *gin.Context)  {
	time.Sleep(time.Second*6)
	code := c.DefaultQuery("code", "")
	resp := response.NewResponse(c)
	// todo validator
	if common.IsEmptyString(code) {
		resp.Body.Code = errorcode.ParamsInvalid.Code()
		resp.ReturnJsonError()
		return
	}

	advertService := service.NewAdvertService()
	adverts, ecode := advertService.FetchAdverting(code)
	if ecode != errorcode.OK {
		resp.Body.Code = ecode.Code()
		resp.ReturnJsonError()
		return
	}
	resp.ReturnJsonSuccess(adverts)
}
