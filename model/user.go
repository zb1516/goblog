package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
	Email string `json:"email"`
	LoginCount int  `json:"login_count" sql:"-"`
	LastTime time.Time `json:"last_time" sql:"-"`
	LastIp string `json:"last_ip" sql:"-"`
	State int8 `json:"state" sql:"-"`
}

//获取完整表名
func (m *User)TableName() string {
	return TableName("user")
}

/**
查询列表
*/
func (u *User)GetUserlist() (list []*User, err error) {
	db := Db
	var user []*User
	errs :=db.Table(u.TableName()).Find(&user).Error
	return user , errs
}

//添加用户
func (u *User)CreateUser(user User) (User,error) {
	err:=Db.Table(u.TableName()).Create(&user).Error
	return user,err
}

//获取用户
func (u *User)GetUserByName(username string) ([]*User,error) {
	var (
		user []*User
	)
	err:=Db.Table(u.TableName()).Where("username = ?" ,username).First(&user).Error
	return user , err
}

//获取用户
func (u *User)GetUser(username string , password string) ([]*User,error) {
	var (
		user []*User
	)
	err:=Db.Table(u.TableName()).Where("username = ?" ,username).Where("password = ?" ,password).First(&user).Error
	return user , err
}

//更新用户
func (u *User)UpdateUser(username string,update interface{})(User , error){
	var (
		user User
	)
	err:=Db.Table(u.TableName()).Where("username = ?",username).Update(update).First(&user).Error
	return user,err
}

//获取用户
func (u *User)GetUserById(userid int) (User,error) {
	var (
		user User
	)
	err:=Db.Table(u.TableName()).Where("id = ?" ,userid).First(&user).Error
	return user , err
}

//删除
func (u *User)DelUser(user User) error {
	err:=Db.Table(u.TableName()).Where("id =?",user.ID).Delete(&user).Error
	return err
}

//更新用户
func (u *User)UpdateUserById(user User)  error{
	err:=Db.Table(u.TableName()).Where("id = ?",user.ID).Update(&user).Error
	return err
}