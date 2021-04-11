package v1

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"githug.com/bob118/fm/config/fmconfig"
	"githug.com/bob118/fm/ec"
	"githug.com/bob118/fm/models"
	"githug.com/bob118/fm/utils"
)

//GetAccounts function.
//request:	GET /api/v1/accounts?uuid=xxx&id=xxx&domain=xxx&page=xxx
//response:	json
func GetAccounts(c *gin.Context) {

	code := ec.SUCCESS
	condition := "true"
	data := make(map[string]interface{})

	if uuid := c.Query("uuid"); len(uuid) > 0 {
		condition += fmt.Sprintf(" and account_uuid ='%s'", uuid)
	}
	if id := c.Query("id"); len(id) > 0 {
		condition += fmt.Sprintf(" and account_id ='%s'", id)
	}
	if domain := c.Query("domain"); len(domain) > 0 {
		condition += fmt.Sprintf(" and account_domain='%s'", domain)
	}
	if page := c.Query("page"); len(page) > 0 {
		if p, err := strconv.Atoi(page); err == nil {
			offset := (uint)(p-1) * fmconfig.CFG.Runtime.Pagesize
			condition += fmt.Sprintf(" offset %d limit %d", offset, fmconfig.CFG.Runtime.Pagesize)
		}
	}

	data["count"] = models.GetAccountsCount(condition)
	data["lists"] = models.GetAccounts(condition)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"data": data,
	})
}

//PostAccount function.
//request:	POST /api/v1/accounts
//request param:	MEMEJSON;MIMEPOSTFORM
//response: json/form.
func PostAccount(c *gin.Context) {
	switch c.ContentType() {
	case gin.MIMEJSON:
		PostAccountJSON(c)
	case gin.MIMEPOSTForm:
	}
}

//PostAccountJSON
func PostAccountJSON(c *gin.Context) {
	code := ec.SUCCESS
	ua := models.Account{}

	//A verify
	if err := c.BindJSON(&ua); err != nil {
		code = ec.ERROR_BIND_JSON
	} else {
		if ua.Aid == "" { //not null
			code = ec.ERROR_PARAM_NULL
		} else {
			if ua.Aname == "" {
				ua.Aname = ua.Aid
			}
			if ua.Aauth == "" {
				ua.Aauth = ua.Aid
			}
			if ua.Apassword == "" {
				ua.Apassword = ua.Aid
			}
			if ua.Aa1hash == "" { // md5(user:domain:password)
				s := fmt.Sprintf("%s:%s:%s", ua.Aid, ua.Adomain, ua.Apassword)
				ua.Aa1hash = utils.MakeA1Hash(s)
			}
			if ua.Agroup == "" {
				ua.Agroup = "default"
			}
			if ua.Adomain == "" {
				code = ec.ERROR_PARAM_NULL
			} else {
				if ua.Aproxy == "" {
					ua.Aproxy = ua.Adomain
				}
			}
			//cacheable notice, default ""
			//1,cacheable = "true"; mod_xml_curl requst reduce.
			//2,cacheable = "60000" mod_xml_curl request cache timer 60s.
			//3,xml_flush_cache;xml_flush_cache id 1002 10.10.10.250
			if ua.Acacheable == "" {
				ua.Acacheable = ""
			}
		}
	}

	if code == ec.SUCCESS {
		isExist, old := models.IsExistByiddomain(ua)
		if isExist { //id@domain exist already.
			ua = old
			code = ec.ERROR_ITEM_EXIST
		} else { //insert ua into account.
			models.CreateAccount(&ua)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"data": ua,
	})
}

//PutAccount function.
//request:	PUT /api/v1/accounts/:uuid
//param:	json
func PutAccount(c *gin.Context) {
	switch c.ContentType() {
	case gin.MIMEJSON:
		PutAccountJSON(c)
	case gin.MIMEPOSTForm:
	}
}

//PutAccountJSON
func PutAccountJSON(c *gin.Context) {
	code := ec.SUCCESS
	ua := models.Account{}

	if uuid := c.Param("uuid"); uuid == "" {
		code = ec.ERROR_PARAM_NULL
	} else {
		isExist, old := models.IsExistByuuid(uuid)
		if !isExist {
			code = ec.ERROR_ITEM_NOTEXIST
		} else {
			if err := c.BindJSON(&ua); err != nil {
				code = ec.ERROR_BIND_JSON
			} else {
				if ua.Auuid != uuid { //param uuid !=ua.uuid
					ua.Auuid = uuid
				}
				if ua.Aid == "" &&
					ua.Aname == "" &&
					ua.Aauth == "" &&
					ua.Apassword == "" &&
					ua.Aa1hash == "" &&
					ua.Agroup == "" &&
					ua.Adomain == "" &&
					ua.Aproxy == "" &&
					ua.Acacheable == "" &&
					true { //all param null
					code = ec.ERROR_PARAM_NULL
				} else {
					models.ModifyAccount(old, &ua)
				}
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"data": ua,
	})
}

//DeleteAccount function.
//request:			DELETE /api/v1/accounts/:uuid
//request param:	uuid
//response:
func DeleteAccount(c *gin.Context) {
	code := ec.SUCCESS
	if uuid := c.Param("uuid"); uuid == "" {
		code = ec.ERROR_PARAM_NULL
	} else {
		isExist, _ := models.IsExistByuuid(uuid)
		if !isExist {
			code = ec.ERROR_ITEM_NOTEXIST
		} else {
			models.DeleteAccount(uuid)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"data": nil,
	})
}
