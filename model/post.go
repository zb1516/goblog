package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"strconv"
)

type Post struct {
	gorm.Model
	UserId int `json:"user_id"`
	Title string `json:"title"`
	Url string `json:"url"`
	Content string `json:"content"`
	Tags string	`json:"tags"`
	Views int `json:"views"`
	IsTop int `json:"is_top"`
	CategoryId int `json:"category_id"`
	Status int8 `json:"status"`
	Types int8 `json:"types"`
	Info string `json:"info"`
	Image string `json:"image"`
	Love int `json:"love"`
	Username string `json:"username"`
}

func (m *Post)TableName() string {
	return TableName("post")
}

//获取文章列表
func (p *Post)GetPostList(page int,limit int,is_top int) ( []*Post , int , error){
	offset := (page - 1) * limit
	query:=Db.Table(p.TableName())
	var (
		post []*Post
		total int
	)
	//查询是否置顶,置顶会显示在首页
	if is_top > 0 {
		query=query.Where("is_top = ?",is_top)
	}
	if page > 0 {
		query = query.Offset(offset)
	}
	if limit > 0 {
		query = query.Limit(limit)
	}
	res := query.Order("created_at desc").Count(&total).Find(&post)
	return post , total , res.Error
}

//通过分类查询文章
func (p *Post)GetPostByCategory(page int , limit int , cateGoryId int) (posts []*Post , count int,  err error)   {
	offset:=(page-1)*10
	var (
		post []*Post
		total int
	)
	query:=Db.Table(p.TableName())
	//通过分类id查询
	if cateGoryId > 0 {
		query=query.Where("category_id = ?",cateGoryId)
	}
	errors := query.Offset(offset).Limit(limit).Find(&post).Count(&total).Error
	return post,total,errors
}

//最新的五条数据
func (p *Post)GetPostListNew() (posts []*Post ,  err error)   {
	var post []*Post
	errors := Db.Table(p.TableName()).Limit(5).Find(&post).Error
	return post,errors
}

//通过分类查询文章
func (p *Post)GetPostByTagPostId(page int , limit int , postId []int) (posts []*Post , count int ,  err error)   {
	offset:=(page-1)*10
	var (
		post []*Post
		total int
	)
	query:=Db.Table(p.TableName())
	//通过分类id查询
	if len(postId) > 0 {
		query=query.Where("id in (?)",postId)
	}
	errors := query.Order("created_at desc").Offset(offset).Limit(limit).Find(&post).Count(&total).Error
	return post,total,errors
}

//查询文章信息
func (p *Post)GetPostInfo(id int) ([]*Post , error) {
	var (
		post []*Post
		view int
	)
	res:=Db.Table(p.TableName()).Where("id = ?",id).First(&post)
	for _,v := range post {
		view = v.Views
	}
	//更新点击量
	update:=make(map[string]interface{})
	update["views"]=view+1
	Db.Table(p.TableName()).Where("id = ?",id).Update(update)
	//把点击量重新赋值给结构
	for _,v := range post {
		v.Views = view+1
	}
	return post , res.Error
}

//更新文章喜欢数
func (p *Post)PostLoveUpate(id int) ([]*Post,error)  {
	var (
		post []*Post
		love int
	)
	err :=Db.Table(p.TableName()).Where("id = ?",id).First(&post).Error
	if err != nil {
		return nil,err
	}
	for _ , v:= range post {
		love = v.Love
	}
	data := make(map[string]int)
	data["Love"]= love + 1
	res:=Db.Table(p.TableName()).Where("id = ?",id).Update(data).Error
	for _ , v:= range post {
		v.Love = love + 1
	}
	return post,res
}

//添加文章
func (p *Post)CreatePost(post Post) (interface{},error) {
	//查询标签
	var (
		tag Tag
		tagPost TagPost
	)
	tx :=Db.Begin()
	tagId,_:=strconv.Atoi(post.Tags)
	post.Tags,_=tag.GetTagName(tagId)
	err:=Db.Table(p.TableName()).Create(&post).Error
	if err != nil {
		tx.Rollback()
		return nil , err
	}
	//添加绑定关系
	tagPost.PostId = int(post.ID)
	tagPost.TagId = tagId
	_ , err =tagPost.CreateTagPost(tagPost)
	if err != nil {
		tx.Rollback()
		return nil ,err
	}
	//查询标签结构体
	tagInfo,_:=tag.GetTagById(tagId)
	fmt.Println(tagInfo)
	tag.ID = uint(tagId)
	tag.Count=tagInfo.Count+1
	fmt.Println(tag)
	//增加标签使用次数
	tag.UpdateTag(tag)
	tx.Commit()
	return post , nil
}


//通过id查询一条数据
func (p *Post)GetPostById(id int)(Post, error)  {
	var (
		post Post
	)
	res:=Db.Table(p.TableName()).Where("id = ?",id).First(&post)
	return post,res.Error
}


//删除
func (p *Post)DelPost(post Post) error {
	err:=Db.Table(p.TableName()).Where("id =?",post.ID).Delete(&post).Error
	return err
}

//通过名称查询
func (p *Post)GetPostByName(title string)([]*Post, error)  {
	var (
		post []*Post
	)
	res:=Db.Table(p.TableName()).Where("title = ?",title).First(&post)
	return post,res.Error
}

//更新文章
func (p *Post)UpdatePost(post Post) (Post,error) {
	err:=Db.Table(p.TableName()).Where("id =?",post.ID).Update(&post).Error
	return post,err
}
