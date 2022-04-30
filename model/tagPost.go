package model

type TagPost struct {
	Id int `json:"id"`
	TagId int `json:"tag_id"`
	PostId int `json:"post_id"`
}

func (m *TagPost)TableName() string {
	return TableName("tag_post")
}

//取出文章id数组
func (t *TagPost)GetPostId(tagId int) ([]int,error) {
	var tagPost []*TagPost
	query:=Db.Table(t.TableName())
	if tagId > 0 {
		query=query.Where("tag_id = ?",tagId)
	}
	res:=query.Find(&tagPost)
	//创建切片，存储所有文章id
	var postID []int
	for _,v := range tagPost{
		postID = append(postID,v.PostId)
	}
	return postID,res.Error
}

//通过文章id取出tagid
func (t *TagPost)GetTagId(postId int) (int,error) {
	var tagPost []*TagPost
	res:=Db.Table(t.TableName()).Where("post_id = ?",postId).Find(&tagPost)
	//创建切片，存储所有文章id
	var tagID int
	for _,v := range tagPost{
		tagID =  v.TagId
	}
	return tagID,res.Error
}

//创建文章和标签绑定关系
func (t *TagPost)CreateTagPost(tagPost TagPost) (int , error) {
	err:=Db.Table(t.TableName()).Create(&tagPost).Error
	return tagPost.Id,err
}