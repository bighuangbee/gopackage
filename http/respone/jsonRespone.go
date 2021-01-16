package respone

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code" 	: http.StatusOK,
		"msg" : StatusText[http.StatusOK],
		"data"	: data,
	})
}

func SuccessH(c *gin.Context, data gin.H){
	c.JSON(http.StatusOK, gin.H{
		"code" 	: http.StatusOK,
		"msg" : StatusText[http.StatusOK],
		"data"	: data,
	})
}

func Error(c *gin.Context, code int, attachMsg string) {
	if(attachMsg == ""){
		attachMsg = StatusText[code]
	}
	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"code" 	: code,
		"msg" : attachMsg,
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
		"code" 	: http.StatusMethodNotAllowed,
		"msg" :     StatusText[http.StatusMethodNotAllowed] + ", " + msg,
		"data"	: nil,
	})
}