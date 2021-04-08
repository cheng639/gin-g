package controller

import (
	"errors"
	"gin-g/global"
	"gin-g/model/response"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Controller struct {
	Binding interface{}     //绑定的验证结构体
	Model interface{}       //关联的ORM结构体
	Models interface{}       //关联的ORM结构体切片
	Relations []string      //需要预加载的关联模型
	Search func(c *gin.Context, tx *gorm.DB) (*gorm.DB)  //列表筛选函数
}

/**
*返回模型数据列表
 */
func (ctr *Controller) Index(c *gin.Context){
	tx := global.DB.Model(ctr.Model)
	tx = ctr.Search(c, tx)
	tx = ctr.Preload(tx)
	tx.Find(ctr.Models)

	response.Data(ctr.Models, c)
}

/**
*列表查询条件
*/
func Search(c *gin.Context, tx *gorm.DB) (*gorm.DB){

	return tx
}

/**
*
 */
func (ctr *Controller) Preload(tx *gorm.DB) (*gorm.DB){
	if len(ctr.Relations) > 0 {
		for _, v := range ctr.Relations{
			tx.Preload(v)
		}
	}

	return tx
}

/**
*添加数据
*/
func (ctr *Controller) Store(c *gin.Context){
	item := ctr.Binding
	if err := c.ShouldBindJSON(item); err != nil{
		response.Message(err.Error(), c)
		return
	}
	result := global.DB.Create(item); if(result.Error != nil){
		response.Error("添加失败!", c)
		return
	}
	response.Message("添加成功!", c)
}

/**
*修改数据
 */
func (ctr *Controller) Update(c *gin.Context){
	item := ctr.Binding
	if err := c.ShouldBindJSON(item); err != nil{
		response.Message(err.Error(), c)
		return
	}

	id := c.Param("id")
	result :=global.DB.Model(ctr.Model).Where("id=?", id).First(ctr.Model)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		response.Error("数据不存在", c)
		return
	}

	result = global.DB.Model(ctr.Model).Where("id=?", id).Updates(item)
	if(result.Error != nil){
		response.Error("修改失败!", c)
		return
	}

	response.Message("修改成功！", c)
}

/**
*返回数据详情
 */
func (ctr *Controller) Show(c *gin.Context){

	id := c.Param("id")
	result :=global.DB.Model(ctr.Model).First(ctr.Model, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		response.Data(nil, c)
		return
	}

	response.Data(ctr.Model, c)
}

/**
*删除数据
 */
func (ctr *Controller) Destroy(c *gin.Context){
	id := c.Param("id")
	result :=global.DB.Delete(ctr.Model, id)
	if(result.RowsAffected > 0){
		response.Message("删除成功！", c)
		return
	}

	response.Message("删除失败！", c)
}