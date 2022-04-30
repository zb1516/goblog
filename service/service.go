package service

import (
	"goblog/model"
	"net/http"
)

type Service struct {}

//获取右侧公共
func (s *Service)GetRight(w http.ResponseWriter) map[string]interface{} {
	var (
		post model.Post
		tag model.Tag
		category model.Category
		config model.Config
	)
	//查询最新的五条数据
	newRes , newErr:=post.GetPostListNew()
	if newErr != nil {
		w.Write([]byte(newErr.Error()))
	}
	//查询出所有标签
	tagRes,tagErr:=tag.GetTagList()
	if tagErr != nil {
		w.Write([]byte(tagErr.Error()))
	}
	cateRes,cateErr:=category.GetCategoryList()
	if cateErr != nil {
		w.Write([]byte(cateErr.Error()))
	}
	//查询配置信息
	configRes , err :=config.GetConfig()
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	confData:=make(map[string]interface{})
	if len(configRes) > 0 {
		for _,val := range configRes {
			confData[val.Name]=val.Value
		}
	}
	data:=make(map[string]interface{})
	data["tag"]=tagRes
	data["newList"]=newRes
	data["category"] = cateRes
	data["config"] = confData
	return data
}

//查询出最新的一条分类
func (s *Service)GetTagFirst() (uint , error){
	var (
		tag model.Tag
	)
	tagId,err:=tag.GetTagId()
	return tagId , err
}


//查询出最新的一条分类
func (s *Service)GetCategoryFirst() (uint,error) {
	var (
		category model.Category
	)
	cateId,err:=category.GetCategoryId()
	return cateId , err
}