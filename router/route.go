package router

import (
	"uegoshop/controller/v1"
	"github.com/gin-gonic/gin"
)

func InitRouter(middleware ...gin.HandlerFunc) *gin.Engine {
	r := gin.Default()
	r.Use(middleware...)
	api := r.Group("/api")
	{
		v1Controller := api.Group("/v1")
		{
			v1.AdvertRegister(v1Controller.Group("/advert"))
		}
	}

	return r
}
