package model

import "github.com/jinzhu/gorm"

type Tag struct {
	gorm.Model
	Name string `json:"name"`
	Count int `json:"count"`
}

func (m *Tag)TableName() string {
	return TableName("tag")
}

//获取标签列表
func (t *Tag)GetTagList()(list []*Tag , err error)  {
	var tag []*Tag
	res:=Db.Table(t.TableName()).Order("created_at desc").Find(&tag)
	return tag,res.Error
}

//返回标签名
func (t *Tag)GetTagName(tagId int) (string , error) {
	var tag []*Tag
	res:=Db.Table(t.TableName()).Where("id = ?",tagId).Order("created_at desc").First(&tag)
	var tagName string
	for _,v:= range tag {
		tagName = v.Name
	}
	return tagName,res.Error
}

//返回标签id
func (t *Tag)GetTagId() (uint ,error) {
	var (
		tag []*Tag
		tagId uint
	)
	res:=Db.Table(t.TableName()).First(&tag)
	for _,v := range tag {
		tagId = v.ID
	}
	return tagId , res.Error
}

//创建分类
func (t *Tag)CreateTag(tag Tag) (Tag,error) {
	err:=Db.Table(t.TableName()).Create(&tag).Error
	return tag,err
}

//通过名称查询分类
func (t *Tag)GetTagByName(name string)([]*Tag, error)  {
	var (
		tag []*Tag
	)
	res:=Db.Table(t.TableName()).Where("name = ?",name).Order("created_at desc").First(&tag)
	return tag,res.Error
}


//更新标签
func (t *Tag)UpdateTag(tag Tag) (Tag,error) {
	err:=Db.Table(t.TableName()).Where("id =?",tag.ID).Update(&tag).Error
	return tag,err
}

//删除
func (t *Tag)DelTag(tag Tag) error {
	err:=Db.Table(t.TableName()).Where("id =?",tag.ID).Delete(&tag).Error
	return err
}

//根据id查询一条数据
func (t *Tag)GetTagById(id int)(Tag, error)  {
	var (
		tag Tag
	)
	res:=Db.Table(t.TableName()).Where("id = ?",id).First(&tag)
	return tag,res.Error
}