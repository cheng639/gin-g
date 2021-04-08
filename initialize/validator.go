package initialize

import (
	"gin-g/global"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

func InitValidator(){
	RegisterDefaultTranslations()
}

//设置验证器本地语言
func RegisterDefaultTranslations(){
	unt := ut.New(zh.New())
	trans, _ := unt.GetTranslator("zh")
	//validate := validator.New()
	validate, _ := binding.Validator.Engine().(*validator.Validate);
	//注册一个函数，获取struct tag里自定义的label作为字段名
	//validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
	//	name:=fld.Tag.Get("label")
	//	return name
	//})
	//验证器注册翻译器
	err := zh_translations.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		global.LOG.Info(err.Error())
	}

}

//func init() {
//	_ = utils.RegisterRule("PageVerify",
//		utils.Rules{
//			"Page":     {utils.NotEmpty()},
//			"PageSize": {utils.NotEmpty()},
//		},
//	)
//	_ = utils.RegisterRule("IdVerify",
//		utils.Rules{
//			"Id": {utils.NotEmpty()},
//		},
//	)
//	_ = utils.RegisterRule("AuthorityIdVerify",
//		utils.Rules{
//			"AuthorityId": {utils.NotEmpty()},
//		},
//	)
//}
