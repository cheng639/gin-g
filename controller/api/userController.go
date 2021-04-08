package api

import (
	"errors"
	"gin-g/controller"
	"gin-g/global"
	"gin-g/model"
	"gin-g/model/request"
	"gin-g/model/response"
	"github.com/gin-gonic/gin"
	"github.com/sanxia/glib"
	"gorm.io/gorm"
	"time"
)

type UserController struct {
	ApiController
}

func NewUserController() *UserController{
	return &UserController{
		ApiController{
			controller.Controller{
				Binding:   &request.Login{},
				Model:     &model.User{},
				Relations: nil,
			},
		},
	}
}

/**
 *用户登录
 */
func (ctr *UserController) Login(c *gin.Context){
	var r = ctr.Binding.(*request.Login)
	if err := c.ShouldBindJSON(r); err != nil{
		response.Message(err.Error(), c)
		return
	}

	_code, err := global.REDIS.Get(r.Phone).Result(); if err != nil {
		response.Error("手机号或验证码错误", c)
		return
	}

	if r.Code != _code {
		response.Error("验证码错误", c)
		return
	}

	result := global.DB.Where("phone=?", r.Phone).Take(ctr.Model)
	// 检查 ErrRecordNotFound 错误
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		response.Error("用户不存在", c)
		return
	}

	//缓存用户token
	uid := ctr.Model.(*model.User).UUID.String()
	token := glib.Md5(uid)
	global.REDIS.Set(token, uid, 4*60*60*time.Second)

	c.Header("User-Token", token)
	response.Result(response.SUCCESS, "登录成功！", ctr.Model, c)
}

