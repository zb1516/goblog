package model

type Config struct {
	ID int
	Name string `json:"name"`
	Value string `json:"value"`
}

type CofnigParam struct {
	Title string `json:"title"`
	Url string `json:"url"`
	Keywords string `json:"keywords"`
	Description string `json:"description"`
	Email string `json:"email"`
	Timezone string `json:"timezone"`
	Start string `json:"start"`
	QQ string `json:"qq"`
	Github string `json:"github"`
	Csdn string `json:"csdn"`
}

func (c *Config)TableName() string {
	return TableName("config")
}

//获取配置
func (c *Config)GetConfig() ([]*Config , error) {
	var config []*Config
	err:=Db.Table(c.TableName()).Find(&config).Error
	return config , err
}

//更新配置信息
func (c *Config)UpdateConfig(data map[string]string) bool {
	var config Config
	tx :=Db.Begin()
	for k,v:=range data{
		config.Name= k
		config.Value= v
		err:=Db.Table(c.TableName()).Where("name = ?",k).Update(&config).Error
		if err != nil {
			tx.Rollback()
			return false
		}
	}
	tx.Commit()
	return true
}