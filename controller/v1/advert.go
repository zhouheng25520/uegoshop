package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"uegoshop/common"
	"uegoshop/database"
	"uegoshop/errorcode"
	"uegoshop/response"
	"uegoshop/service"
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
