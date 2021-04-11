package v1

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"githug.com/bob118/fm/config/fmconfig"
	"githug.com/bob118/fm/ec"
	"githug.com/bob118/fm/models"
)

//GetE164s function.
//request:	GET /api/v1/e164s?uuid=xxx&number=xxx&guuid=xxx&page=xxx
//response:	json
func GetE164s(c *gin.Context) {

	code := ec.SUCCESS
	condition := "true"
	data := make(map[string]interface{})

	if uuid := c.Query("uuid"); len(uuid) > 0 {
		condition += fmt.Sprintf(" and e164_uuid='%s'", uuid)
	}
	if number := c.Query("number"); len(number) > 0 {
		condition += fmt.Sprintf(" and e164_number='%s'", number)
	}
	if guuid := c.Query("guuid"); len(guuid) > 0 {
		condition += fmt.Sprintf(" and gateway_uuid='%s'", guuid)
	}
	if page := c.Query("page"); len(page) > 0 {
		if p, err := strconv.Atoi(page); err == nil {
			offset := (uint)(p-1) * fmconfig.CFG.Runtime.Pagesize
			condition += fmt.Sprintf(" offset %d limit %d", offset, fmconfig.CFG.Runtime.Pagesize)
		}
	}

	data["count"] = models.GetE164sCount(condition)
	data["lists"] = models.GetE164s(condition)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"data": data,
	})
}

//PostE164 function.
//request:	POST /api/v1/e164s
//response:	json
func PostE164(c *gin.Context) {
	switch c.ContentType() {
	case gin.MIMEJSON:
		PostE164JSON(c)
	case gin.MIMEPOSTForm:
	}
}

//PostE164JSON function.
func PostE164JSON(c *gin.Context) {
	code := ec.SUCCESS
	e164 := models.E164{}

	//input verify
	if err := c.BindJSON(&e164); err != nil {
		code = ec.ERROR_BIND_JSON
	} else {
		if e164.Enumber == "" {
			code = ec.ERROR_PARAM_NULL
		} else {
			exist, _ := models.IsExistE164Bynumber(e164)
			if exist {
				code = ec.ERROR_ITEM_EXIST
			} else {
				models.CreateE164(&e164)
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"data": e164,
	})
}

//PutE164 function.
//request:	PUT /api/v1/e164s/:uuid
//response:	json
func PutE164(c *gin.Context) {
	switch c.ContentType() {
	case gin.MIMEJSON:
		PutE164JSON(c)
	case gin.MIMEPOSTForm:
	}
}

//PutE164JSON fuction.
func PutE164JSON(c *gin.Context) {
	code := ec.SUCCESS
	e164 := models.E164{}

	if uuid := c.Param("uuid"); uuid == "" {
		code = ec.ERROR_PARAM_NULL
	} else {
		isExist, old := models.IsExistE164Byuuid(uuid)
		if !isExist {
			code = ec.ERROR_ITEM_NOTEXIST
		} else {
			if err := c.BindJSON(&e164); err != nil {
				code = ec.ERROR_BIND_JSON
			} else {
				if e164.Euuid != uuid {
					e164.Euuid = uuid
				}
				if true &&
					e164.Enumber == "" {
					code = ec.ERROR_PARAM_NULL
				} else {
					models.ModifyE164(old, &e164)
				}
			}
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"data": e164,
	})
}

//DeleteE164 function.
//request:			DELETE /api/v1/e164s/:uuid
//request param:	uuid
//response:			nil
func DeleteE164(c *gin.Context) {
	code := ec.SUCCESS
	if uuid := c.Param("uuid"); uuid == "" {
		code = ec.ERROR_PARAM_NULL
	} else {
		exist, _ := models.IsExistE164Byuuid(uuid)
		if !exist {
			code = ec.ERROR_ITEM_NOTEXIST
		} else {
			models.DeleteE164(uuid)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"data": nil,
	})
}
