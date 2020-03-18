package v1

import (
	"ccshop/common"
	"ccshop/database"
	"ccshop/errorcode"
	"ccshop/response"
	"ccshop/service"
	"fmt"
	"github.com/gin-gonic/gin"
)

type Advert struct {
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
	if common.IsEmptyString(code) {
		resp.Body.Code = errorcode.ParamsInvalid.Code()
		resp.ReturnJsonError()
		return
	}
	config := &database.Config{
		User:        "root",
		Password:    "123456",
		DBName:      "ccshop",
		Charset:     "utf8",
		//MaxOpenConn: 50,
		//MaxLifetime: int(time.Second * 3),
	}
	db := database.NewMysqlDB(config)
	if err := db.Create(); err != nil {
		fmt.Println("err is :", err)
		return
	}
	defer db.DB.Close()

	advertService := service.NewAdvertService(db.DB)
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
	config := &database.Config{
		User:        "root",
		Password:    "123456",
		DBName:      "ccshop",
		Charset:     "utf8",
		//MaxOpenConn: 50,
		//MaxLifetime: int(time.Second * 3),
	}
	db := database.NewMysqlDB(config)
	if err := db.Create(); err != nil {
		fmt.Println("err is :", err)
		return
	}
	defer db.DB.Close()

	advertService := service.NewAdvertService(db.DB)
	adverts, ecode := advertService.FetchAdvertingFromModel(code)
	if ecode != errorcode.OK {
		resp.Body.Code = ecode.Code()
		resp.ReturnJsonError()
		return
	}
	resp.ReturnJsonSuccess(adverts)
}
