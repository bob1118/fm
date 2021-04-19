package v1

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/bob1118/fm/config/fmconfig"
	"github.com/bob1118/fm/ec"
	"github.com/bob1118/fm/models"
	"github.com/gin-gonic/gin"
)

//GetGateways function.
//request:	GET /api/v1/gateways?uuid=xxx&name=xxx&realm=xxx&page=xxx
//response:	json
func GetGateways(c *gin.Context) {

	code := ec.SUCCESS
	condition := "true"
	data := make(map[string]interface{})

	if uuid := c.Query("uuid"); len(uuid) > 0 {
		condition += fmt.Sprintf(" and gateway_uuid='%s'", uuid)
	}
	if name := c.Query("name"); len(name) > 0 {
		condition += fmt.Sprintf(" and gateway_name='%s'", name)
	}
	if realm := c.Query("realm"); len(realm) > 0 {
		condition += fmt.Sprintf(" and gateway_realm='%s'", realm)
	}
	if page := c.Query("page"); len(page) > 0 {
		if p, err := strconv.Atoi(page); err == nil {
			offset := (uint)(p-1) * fmconfig.CFG.Runtime.Pagesize
			condition += fmt.Sprintf(" offset %d limit %d", offset, fmconfig.CFG.Runtime.Pagesize)
		}
	}

	data["count"] = models.GetGatewaysCount(condition)
	data["lists"] = models.GetGateways(condition)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"data": data,
	})
}

//PostGateway function.
//request:	POST /api/v1/gateways
//request param:	MEMEJSON;MIMEPOSTFORM
//response: json/form.
func PostGateway(c *gin.Context) {
	switch c.ContentType() {
	case gin.MIMEJSON:
		PostGatewayJSON(c)
	case gin.MIMEPOSTForm:
	}
}

//PostGatewayJSON function.
func PostGatewayJSON(c *gin.Context) {
	code := ec.SUCCESS
	gw := models.Gateway{}

	//input verify
	if err := c.BindJSON(&gw); err != nil {
		code = ec.ERROR_BIND_JSON
	} else {
		if gw.Gname == "" { //not null
			code = ec.ERROR_PARAM_NULL
		} else {
			//if gw.Gusername == "" { //nothing todo.
			//}
			if gw.Grealm == "" {
				gw.Grealm = gw.Gname
			}
			if gw.Gfromuser == "" {
				gw.Gfromuser = gw.Gusername
			}
			if gw.Gfromdomain == "" {
				gw.Gfromdomain = gw.Grealm
			}
			//if gw.Gpassword == "" { //nothing todo.
			//}
			if gw.Gextension == "" {
				gw.Gextension = gw.Gusername
			}
			if gw.Gproxy == "" {
				gw.Gproxy = gw.Grealm
			}
			if gw.Gregisterproxy == "" {
				gw.Gregisterproxy = gw.Gproxy
			}
			if gw.Gexpire == "" {
				gw.Gexpire = "3600"
			}
			if gw.Gregister == "" {
				gw.Gregister = "false"
			}
			if gw.Gcalleridinfrom == "" {
				gw.Gcalleridinfrom = "true"
			}
			if gw.Gextensionincontact == "" {
				gw.Gextensionincontact = "true"
			}
			//if gw.Goptionping == "" { //nothing todo.
			//}
		}
	}

	if code == ec.SUCCESS {
		isExist, _ := models.IsExistGatewayByname(gw)
		if isExist {
			code = ec.ERROR_ITEM_EXIST
		} else {
			models.CreateGateway(&gw)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"data": gw,
	})
}

//PutGateway function.
//request:	PUT /api/v1/gateways/:uuid
//param:	json
func PutGateway(c *gin.Context) {
	switch c.ContentType() {
	case gin.MIMEJSON:
		PutGatewayJSON(c)
	case gin.MIMEPOSTForm:
	}
}

//PutGatewayJSON function.
func PutGatewayJSON(c *gin.Context) {
	code := ec.SUCCESS
	gw := models.Gateway{}

	if uuid := c.Param("uuid"); uuid == "" {
		code = ec.ERROR_PARAM_NULL
	} else {
		isExist, old := models.IsExistGatewayByuuid(uuid)
		if !isExist {
			code = ec.ERROR_ITEM_NOTEXIST
		} else {
			if err := c.BindJSON(&gw); err != nil {
				code = ec.ERROR_BIND_JSON
			} else {
				if gw.Guuid == "" {
					gw.Guuid = old.Guuid
				}
				if gw.Gname == "" &&
					gw.Gusername == "" &&
					gw.Grealm == "" &&
					gw.Gfromuser == "" &&
					gw.Gfromdomain == "" &&
					gw.Gpassword == "" &&
					gw.Gextension == "" &&
					gw.Gproxy == "" &&
					gw.Gregisterproxy == "" &&
					gw.Gexpire == "" &&
					gw.Gregister == "" &&
					gw.Gcalleridinfrom == "" &&
					gw.Gextensionincontact == "" &&
					gw.Goptionping == "" &&
					true {
					code = ec.ERROR_PARAM_NULL
				} else {
					models.ModifyGateway(old, &gw)
				}
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"data": gw,
	})
}

//DeleteGateway function.
//request:			DELETE /api/v1/gateways/:uuid
//request param:	uuid
//response:			nil.
func DeleteGateway(c *gin.Context) {
	code := ec.SUCCESS
	if uuid := c.Param("uuid"); uuid == "" {
		code = ec.ERROR_PARAM_NULL
	} else {
		isExist, _ := models.IsExistGatewayByuuid(uuid)
		if !isExist {
			code = ec.ERROR_ITEM_NOTEXIST
		} else {
			models.DeleteGateway(uuid)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"data": nil,
	})
}
