package model

import (
	"github.com/jinzhu/gorm"
)

type Category struct {
	gorm.Model
	Name string `json:"name"`
}

//定义表名
func (m *Category)TableName() string {
	return TableName("category")
}

//返回所有分类
func (c *Category)GetCategoryList() (list []*Category,err error)  {
	var category []*Category
	res:=Db.Table(c.TableName()).Order("created_at desc").Find(&category)
	return category,res.Error
}

//查询一条数据
func (c *Category)GetCategoryName(id int)(string, error)  {
	var (
		category []*Category
		categoryName string
	)
	res:=Db.Table(c.TableName()).Where("id = ?",id).Order("created_at desc").First(&category)
	for _ , v:= range category {
		categoryName = v.Name
	}
	return categoryName,res.Error
}


//返回标签id
func (c *Category)GetCategoryId() (uint ,error) {
	var (
		category []*Category
		cateId uint
	)
	res:=Db.Table(c.TableName()).First(&category)
	for _,v := range category {
		cateId = v.ID
	}
	return cateId , res.Error
}

//创建分类
func (c *Category)CreateCategory(category Category) (Category,error) {
	err:=Db.Table(c.TableName()).Create(&category).Error
	return category,err
}

//通过名称查询分类
func (c *Category)GetCategoryByName(name string)([]*Category, error)  {
	var (
		category []*Category
	)
	res:=Db.Table(c.TableName()).Where("name = ?",name).Order("created_at desc").First(&category)
	return category,res.Error
}

//通过id查询一条数据
func (c *Category)GetCategoryById(id int)(Category, error)  {
	var (
		category Category
	)
	res:=Db.Table(c.TableName()).Where("id = ?",id).Order("created_at desc").First(&category)
	return category,res.Error
}

//更新分类
func (c *Category)UpdateCategory(category Category) (Category,error) {
	err:=Db.Table(c.TableName()).Where("id =?",category.ID).Update(&category).Error
	return category,err
}

//删除
func (c *Category)DelCategory(category Category) error {
	err:=Db.Table(c.TableName()).Where("id =?",category.ID).Delete(&category).Error
	return err
}
