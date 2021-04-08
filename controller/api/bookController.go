package api

import (
	"gin-g/controller"
	"gin-g/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BookController struct {
	ApiController
}

func NewBookController() *BookController{
	return &BookController{
		ApiController{
			controller.Controller{
				Binding:   &model.Book{},
				Model:     &model.Book{},
				Models:    &[]model.Book{},
				Relations: []string{"Stage"},
				Search: controller.Search,    //选择或定义一个列表筛选函数
			},
		},
	}
}

/**
*列表查询条件
*/
func UserSearch(c *gin.Context, tx *gorm.DB) (*gorm.DB){

	tx.Where("status=?", 1)

	return tx
}





