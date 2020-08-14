package respone

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code" 	: http.StatusOK,
		"msg" : StatusText[http.StatusOK],
		"data"	: data,
	})
}

func Error(c *gin.Context, code int, attachMsg string) {
	msg := StatusText[code]
	if(attachMsg != ""){
		msg = msg + ", " + attachMsg
	}
	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"code" 	: code,
		"msg" : msg,
		"data"	: nil,
	})
}

func UnAuthorized(c *gin.Context, msg string){
	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"code" 	: http.StatusUnauthorized,
		"msg" :     StatusText[http.StatusUnauthorized] + ", " + msg,
		"data"	: nil,
	})
}

func UnPerminssion(c *gin.Context, msg string){
	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"code" 	: http.StatusUnauthorized,
		"msg" :     StatusText[http.StatusUnauthorized] + ", " + msg,
		"data"	: nil,
	})
}