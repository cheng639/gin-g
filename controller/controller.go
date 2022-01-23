package controller

import (
	"errors"
	"gin-g/global"
	"gin-g/model/response"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"reflect"
	"strconv"
)

type Controller struct {
	Binding   interface{}                                //绑定的验证结构体
	Model     interface{}                                //关联的ORM结构体
	Models    interface{}                                //关联的ORM结构体切片
	Relations []string                                   //需要预加载的关联模型
	Search    func(c *gin.Context, tx *gorm.DB) *gorm.DB //列表筛选函数
}

/**
*返回模型数据列表
 */
func (ctr *Controller) Index(c *gin.Context) {
	ctr.Model = ctr.NewModel(ctr.Model)

	//获取分页参数
	p, err := ctr.GetPageParams(c)
	if err != nil {
		response.Error(err.Error(), c)
		return
	}

	//无分页参数不分页
	if p == nil {
		tx := global.DB.Model(ctr.Model)
		tx = ctr.Search(c, tx)
		tx = ctr.Preload(tx)
		tx.Order("id desc").Find(ctr.Models)

		response.Data(ctr.Models, c)
	} else {
		var total int64
		tx := global.DB.Model(ctr.Model)
		tx = ctr.Search(c, tx)
		tx = tx.Count(&total)
		p.Total = int(total)
		tx = ctr.Preload(tx)
		tx.Limit(p.Limit).Offset((p.Page - 1) * p.Limit).Order("id desc").Find(ctr.Models)

		response.Data(response.PagedData{List: ctr.Models, Pagination: p}, c)
	}

}

/**
*列表查询条件
 */
func Search(c *gin.Context, tx *gorm.DB) *gorm.DB {

	return tx
}

/**
*预加载
 */
func (ctr *Controller) Preload(tx *gorm.DB) *gorm.DB {
	if len(ctr.Relations) > 0 {
		for _, v := range ctr.Relations {
			tx.Preload(v)
		}
	}

	return tx
}

/**
*添加数据
 */
func (ctr *Controller) Store(c *gin.Context) {
	ctr.Binding = ctr.NewModel(ctr.Binding)

	if err := c.ShouldBindJSON(ctr.Binding); err != nil {
		response.Message(err.Error(), c)
		return
	}

	err := global.DB.Create(ctr.Binding).Error
	if err != nil {
		global.LOG.Error(err.Error())
		response.Error(err.Error(), c)
		return
	}
	response.Message("添加成功!", c)
}

/**
*修改数据
 */
func (ctr *Controller) Update(c *gin.Context) {
	ctr.Binding = ctr.NewModel(ctr.Binding)
	ctr.Model = ctr.NewModel(ctr.Model)
	if err := c.ShouldBindJSON(ctr.Binding); err != nil {
		response.Message(err.Error(), c)
		return
	}

	id := c.Param("id")
	err := global.DB.Model(ctr.Model).First(ctr.Model, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		response.Error("数据不存在", c)
		return
	}

	err = global.DB.Debug().Model(ctr.Model).Where("id=?", id).Updates(ctr.Binding).Error
	if err != nil {
		response.Error(err.Error(), c)
		return
	}

	response.Message("修改成功！", c)
}

/**
*返回数据详情
 */
func (ctr *Controller) Show(c *gin.Context) {
	ctr.Model = ctr.NewModel(ctr.Model)
	id := c.Param("id")

	tx := global.DB.Debug().Model(ctr.Model)
	tx = ctr.Preload(tx)

	err := tx.First(ctr.Model, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		response.Data(nil, c)
		return
	}

	response.Data(ctr.Model, c)
}

/**
 * 更新控制器所属的模型
 */
func (ctr *Controller) NewModel(i interface{}) interface{} {
	value := reflect.ValueOf(i)
	modelType := reflect.Indirect(value).Type()
	modelValue := reflect.New(modelType).Interface()

	return modelValue
}

/**
*删除数据
 */
func (ctr *Controller) Destroy(c *gin.Context) {
	id := c.Param("id")
	result := global.DB.Delete(ctr.Model, id)
	if result.RowsAffected > 0 {
		response.Message("删除成功！", c)
		return
	}

	response.Message("删除失败！", c)
}

func (ctr *Controller) GetPageParams(c *gin.Context) (*response.Pagination, error) {
	if c.Query("page") == "" {
		return nil, nil
	}

	p := &response.Pagination{}
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		return nil, err
	}
	p.Page = page

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "20"))
	if err != nil {
		return nil, err
	}
	p.Limit = limit

	return p, nil
}
