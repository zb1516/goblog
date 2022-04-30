package config

import "C"
import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type MysqlConfig struct {
	Host     string `yaml:"host"`
	Name     string `yaml:"name"`
	Port     string `yaml:"port"`
	Password string `yaml:"password"`
	User     string `yaml:"user"`
	Chartset string `yaml:"chartset"`
	Prefix 	 string `yaml:"prefix"`
}

type CacheConfig struct {
	Host     string `yaml:"host"`
	Port     string    `yaml:"port"`
	Password string `yaml:"password"`
	Db       string    `yaml:"db"`
}

type Config struct {
	Mysql MysqlConfig `yaml:"mysql"`
	Cache CacheConfig `yaml:"cache"`
}

//定义解析后的配置文件路径
var activePath string

//初始化配置文件
func (c *Config)_Init() *Config{
	InitActivePath("./","env")
	config, err := ioutil.ReadFile(activePath)
	if err != nil {
		fmt.Print(err)
	}
	//result := make(map[string]interface{})
	errs:=yaml.Unmarshal(config, &c)
	if errs != nil {
		fmt.Println(errs.Error())
	}
	return c
}

//加载配置文件
func LoadConfig() *Config {
	//定义从配置文件转换成的对象
	var c Config
	return c._Init()
}

//加载配置项路径
func InitActivePath(relativePath, activeEnv string) {
	if activeEnv == "env" {
		activeEnv = "config.env"
	}else {
		activePath = "config.prod"
	}
	activePath = relativePath + activeEnv + ".yaml"
}


//定义连接数据库的函数
func DatabaseDialString() string {
	conf:=LoadConfig()
	return fmt.Sprintf("%s:%s@%s(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		conf.Mysql.User,
		conf.Mysql.Password,
		"tcp",
		conf.Mysql.Host,
		conf.Mysql.Port,
		conf.Mysql.Name,
		"utf8",
	)
}
