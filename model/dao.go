package model

import "C"
import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"goblog/config"
)

var Db *gorm.DB

func init() {
	//设置表名的前缀(只会修改默认的表名规则)
	gorm.DefaultTableNameHandler = func(db *gorm.DB,defaultTableName string)string {
		return "tb_"+defaultTableName
	}
	//读取数据库配置
	dns:=config.DatabaseDialString()
	//链接数据库
	conn,err:=gorm.Open("mysql",dns)
	if err != nil  {
		panic(err)
	}
	Db = conn
}

//返回带前缀的完整表名
func TableName(str string) string {
	conf:=config.LoadConfig()
	return conf.Mysql.Prefix + str
}