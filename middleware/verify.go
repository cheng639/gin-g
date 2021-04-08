package middleware

import (
	"gin-g/global"
	"gin-g/model/response"
	"github.com/gin-gonic/gin"
	"sort"
)

// 检查api请求的用户token,
func Verify() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		fullPath := c.FullPath()
		result := exempt(method,fullPath);if result == true{
			// 处理请求
			c.Next()
			return
		}

		token := c.GetHeader("User-Token")
		uid, err := global.REDIS.Get(token).Result(); if err != nil {
			response.Error("用户未登录!", c)
			c.Abort()
			return
		}
		c.Set("uid", uid)

	}
}

//无需登录的接口
func getExemptions()  map[string][]string  {
	var exemptions map[string][]string
	exemptions = make(map[string][]string)
	exemptions["GET"] = []string{}
	exemptions["POST"] = []string{
		"/api/users/login",
		"/api/smsCode",
	}
	exemptions["PUT"] = []string{}
	exemptions["DELETE"] = []string{}

	return exemptions
}

//检查接口是否需要token
func exempt(method, fullPath string) bool{
	exemptions := getExemptions()
	paths := exemptions[method]
	sort.Strings(paths)
	exist := sort.SearchStrings(paths, fullPath); if exist == len(exemptions[method]){
		return false
	}

	return true
}



