package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`

}

const (
	SUCCESS = 0
	ERROR   = 101
)

func Result(code int, msg string, data interface{},  c *gin.Context) {
	// 开始时间
	c.JSON(http.StatusOK, Response{
		code,
		msg,
		data,
	})
}


func Success(c *gin.Context) {
	Result(SUCCESS, "操作成功", map[string]interface{}{}, c)
}

func Message(message string, c *gin.Context) {
	Result(SUCCESS, message, map[string]interface{}{}, c)
}

func Data(data interface{}, c *gin.Context) {
	Result(SUCCESS, "OK", data,  c)
}

func Fail(c *gin.Context) {
	Result(ERROR, "操作失败", map[string]interface{}{}, c)
}

func Error(message string, c *gin.Context) {
	Result(ERROR, message, map[string]interface{}{}, c)
}


