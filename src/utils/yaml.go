package utils

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type YmlConfig struct {
	Ip       string `yaml:"ip"`        //服务器ip
	User     string `yaml:"user"`      //用户
	Pwd      string `yaml:"pwd"`       //用户密码
	DbName   string `yaml:"db_name"`   //数据库名字
	DbPort   string `yaml:"db_port"`   //数据库端口
	HttpPort string `yaml:"http_port"` //HttpServer port
}

func (c *YmlConfig) GetYml() *YmlConfig {
	yamlFile, err := ioutil.ReadFile("src/config/config.yml")
	if err != nil {
		fmt.Println(err.Error())
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		fmt.Println(err.Error())
	}
	return c
}
