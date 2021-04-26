package routers

import (
	"github.com/bob1118/fm/config/fmconfig"
	v1 "github.com/bob1118/fm/routers/api/v1"
	"github.com/bob1118/fm/routers/fsapi"
	"github.com/gin-gonic/gin"
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

		apiv1.GET("/gateways", v1.GetGateways)
		apiv1.POST("/gateways", v1.PostGateway)
		apiv1.PUT("/gateways/:uuid", v1.PutGateway)
		apiv1.DELETE("/gateways/:uuid", v1.DeleteGateway)

		apiv1.GET("/e164s", v1.GetE164s)
		apiv1.POST("/e164s", v1.PostE164)
		apiv1.PUT("/e164s/:uuid", v1.PutE164)
		apiv1.DELETE("/e164s/:uuid", v1.DeleteE164)

		apiv1.POST("/call/dial")
		apiv1.POST("/call/dialout")
	}
	return r
}
