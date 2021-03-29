package routers

import (
	"github.com/gin-gonic/gin"
	"githug.com/bob118/fm/config/fmconfig"
	v1 "githug.com/bob118/fm/routers/api/v1"
	"githug.com/bob118/fm/routers/fsapi"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(fmconfig.CFG.Runtime.Runmode)

	//receive mod_xml_curl request
	r.POST("/fsapi", fsapi.PostFromXmlCurl)

	//receive open api request
	apiv1 := r.Group("/api/v1")
	{
		apiv1.GET("/accounts", v1.GetAccounts)
		apiv1.POST("/accounts", v1.PostAccount)
		apiv1.PUT("/accounts/:uuid", v1.PutAccount)
		apiv1.DELETE("/accounts/:uuid", v1.DeleteAccount)

		apiv1.GET("/gateways")
		apiv1.POST("/gateways")
		apiv1.PUT("/gateways/:uuid")
		apiv1.DELETE("/gateways/:uuid")

		apiv1.GET("/e164s")
		apiv1.POST("/e164s")
		apiv1.PUT("/e164s/:uuid")
		apiv1.DELETE("/e164s/:uuid")

		apiv1.POST("/call/dial")
		apiv1.POST("/call/dialout")
	}

	return r
}
