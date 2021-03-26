package v1

import "github.com/gin-gonic/gin"

//GetAccounts function.
//request:	GET /api/v1/accounts?uuid=xxx&id=xxx&domain=xxx&page=xxx
//response:	json
func GetAccounts(c *gin.Context) {}

//PostAccount function.
//request:	POST /api/v1/accounts
//request param:	MEMEJSON;MIMEPOSTFORM
//response: json/form.
func PostAccount(c *gin.Context) {
	switch c.ContentType() {
	case gin.MIMEJSON:
		PostAccountJSON(c)
	case gin.MIMEPOSTForm:
		PostAccountFORM(c)
	default:
	}
}
func PostAccountJSON(c *gin.Context) {}
func PostAccountFORM(c *gin.Context) {}

//PutAccount function.
//request:	PUT /api/v1/accounts/:uuid
//param:	json
func PutAccount(c *gin.Context) {
	switch c.ContentType() {
	case gin.MIMEJSON:
		PutAccountJSON(c)
	case gin.MIMEPOSTForm:
		PutAccountFORM(c)
	default:
	}
}
func PutAccountJSON(c *gin.Context) {}
func PutAccountFORM(c *gin.Context) {}

//DeleteAccount function.
//request:			DELETE /api/v1/accounts/:uuid
//request param:	uuid
//response:
func DeleteAccount(c *gin.Context) {}
