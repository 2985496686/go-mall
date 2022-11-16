package conf

import (
	"gopkg.in/yaml.v2"
	"log"
	"mall/cache"
	"mall/dao"
	"os"
	"strings"
)

type Config struct {
	Mysql struct {
		DB         string `yaml:"Db"`
		DBHost     string `yaml:"DbHost"`
		DBPort     string `yaml:"DbPort"`
		DBUser     string `yaml:"DbUser"`
		DBPassWord string `yaml:"DbPassWord"`
		DBName     string `yaml:"DbName"`
	}

	Redis struct {
		Addr     string `yaml:"addr"`
		Password string `yaml:"password"`
		DB       int    `yaml:"db"`
	}
}

func Init() {
	config, err := os.ReadFile("./conf/config.yaml")
	if err != nil {
		log.Println("配置文件读取错误:", err)
		return
	}
	var c Config
	err = yaml.Unmarshal(config, &c)
	if err != nil {
		log.Println("反序列化配置文件失败:", err)
	}

	//配置mysql
	readPath := strings.Join([]string{c.Mysql.DBUser, ":", c.Mysql.DBPassWord, "@tcp(", c.Mysql.DBHost, ":", c.Mysql.DBPort, ")/", c.Mysql.DBName, "?charset=utf8&parseTime=true"}, "")
	writePath := strings.Join([]string{c.Mysql.DBUser, ":", c.Mysql.DBPassWord, "@tcp(", c.Mysql.DBHost, ":", c.Mysql.DBPort, ")/", c.Mysql.DBName, "?charset=utf8&parseTime=true"}, "")
	dao.DataBase(readPath, writePath)

	//配置redis
	cache.Cache(c.Redis.Addr, c.Redis.Password, c.Redis.DB)
}
