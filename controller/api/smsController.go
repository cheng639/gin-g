package api

import (
	"fmt"
	"gin-g/global"
	"gin-g/model/request"
	"gin-g/model/response"
	"gin-g/utils/gsms"
	"github.com/gin-gonic/gin"
	"math/rand"
	"time"
)

func SendCode(c *gin.Context){
	var s request.SmsCode
	if err := c.ShouldBindJSON(&s); err != nil{
		response.Message(err.Error(), c)
		return
	}

	code := fmt.Sprintf("%d", int32(RangeRand(999999)))
	global.REDIS.Set(s.Phone, code, 60*60*time.Second)

	_, err := AliyunSmsSend(s.Phone, code)
	if err != nil{
		response.Message(err.Error(), c)
		return
	}

	response.Message("发送成功！", c)

}

// mobiles 接收短信的手机
// code 对应的模板替换内容
func AliyunSmsSend(mobiles string, code string) (*gsms.SmsResult, error) {
	SMS := global.CONFIG.SMS
	smsProvider := gsms.NewAliyunSms(SMS.AccessId, SMS.AccessKey, SMS.RegionId, SMS.SignName)
	smsProvider.SetTemplateCode(SMS.TemplateCode)
	smsProvider.SetTemplateParam(gsms.SmsTemplateParam{
		Code: code, //修改成你对应的模板码
	})
	return smsProvider.Send(mobiles)
}

//生成规定范围内的整数
//设置起始数字范围，0开始,n截止
func RangeRand(n int) int {

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(n)

}

