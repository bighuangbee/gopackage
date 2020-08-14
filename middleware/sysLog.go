package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"strings"
)

/**
* @Author: bigHuangBee
* @Date: 2020/3/24 19:04
 */

//1=新增 2=查看 3=编辑 4=删除 5=登录 6=注销 7=导出
const (
	LOG_TYPE_ADD = 1
	LOG_TYPE_READ = 2
	LOG_TYPE_EDIT = 3
	LOG_TYPE_DEL = 4
	LOG_TYPE_LOGIN = 5
	LOG_TYPE_LOGOUT = 6
	LOG_TYPE_EXPORT = 7
)

func SysLog() gin.HandlerFunc {

	return func(c *gin.Context) {

		//writer := responeWriter{
		//	c.Writer,
		//	bytes.NewBuffer([]byte("")),
		//}
		//c.Writer = &writer
		//
		//var requestParams string;
		//var username string
		//if c.Request.Method == "GET"{
		//	if len(c.Request.URL.Query()) > 0{
		//		params, _ := json.Marshal(c.Request.URL.Query())
		//		requestParams = string(params)
		//	}
		//}else{
		//	if strings.Contains(c.Request.Header.Get("Content-Type"), "application/json"){
		//		body,_ := ioutil.ReadAll(c.Request.Body)
		//		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		//		requestParams = string(body)
		//
		//		if c.Request.RequestURI == "/user/login"{
		//			type User struct {
		//				Username	string `form:"username"`
		//			}
		//			var user User
		//			json.Unmarshal(body, &user)
		//			username = user.Username
		//		}
		//	} else{
		//		c.DefaultPostForm("default", "")
		//		c.Request.ParseForm()
		//
		//		if len(c.Request.PostForm) > 0{
		//			result,_ := json.Marshal(c.Request.PostForm)
		//			requestParams = string(result)
		//		}
		//	}
		//}
		//
		//c.Next()
		//
		//type Res struct {
		//	Code int `json:"code"`
		//	Msg string `json:"msg"`
		//
		//}
		//res := Res{}
		//err := json.Unmarshal(writer.WriterBuff.Bytes(), &res)
		//if err != nil {
		//	loger.Error(err)
		//	return
		//}
		//
		//if res.Code != http.StatusOK{
		//	return
		//}
		//
		//
		//sysLog := sys.SysLog{
		//	Type:       getLogType(c),
		//	CreateTime: storage.Time(time.Now()),
		//	IP:         c.ClientIP(),
		//	Url:        c.Request.URL.Path,
		//	UrlTitle:   sys.GetSysResource(c.Request.URL.Path).Title + " " + c.Request.URL.Path,
		//}
		//
		//claims, ok := c.Get("claims")
		//if ok {
		//	userclaims := claims.(*jwtService.UserClaims)
		//
		//	sysLog.UserId = userclaims.UserId
		//	sysLog.Username = userclaims.UserName
		//	sysLog.Nickname = userclaims.NickName
		//	sysLog.Params = requestParams
		//}else{
		//	sysLog.Params = ""
		//	sysLog.Username = username
		//}
		//
		//sysLog.Add()
	}
}

func getLogType(c *gin.Context) int{
	logType := 0
	if c.Request.Method == "GET"{
		logType = LOG_TYPE_READ
	}else if c.Request.Method == "POST"{
		if strings.LastIndex(c.Request.RequestURI, "add") > 0{
			logType = LOG_TYPE_ADD
		}else if strings.LastIndex(c.Request.RequestURI, "del") > 0{
			logType = LOG_TYPE_DEL
		}else if strings.LastIndex(c.Request.RequestURI, "edit") > 0 || strings.LastIndex(c.Request.RequestURI, "update") > 0{
			logType = LOG_TYPE_EDIT
		}
	}
	if c.Request.RequestURI == "/user/login"{
		logType = LOG_TYPE_LOGIN
	}
	if c.Request.RequestURI == "/user/logout"{
		logType = LOG_TYPE_LOGOUT
	}
	return logType
}

/**
重新实现ResponseWriter接口的Write方法
保存请求回复的数据副本
*/
type responeWriter struct {
	gin.ResponseWriter
	WriterBuff *bytes.Buffer
}

func (r *responeWriter) Write(body []byte) (size int, err error){
	r.WriterBuff.Write(body)
	size, err = r.ResponseWriter.Write(body)	//调用ResponseWriter接口的原write
	return
}
