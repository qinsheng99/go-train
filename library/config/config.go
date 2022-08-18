package config

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/qinsheng99/goWeb/library/etcd"
	"go.uber.org/zap"
)

// Config 整个项目的配置
type Config struct {
	Mode         string `json:"mode"`
	Port         int    `json:"port"`
	*LogConfig   `json:"log"`
	*MysqlConfig `json:"mysql"`
	*EsConfig    `json:"es"`
	*RedisConfig `json:"redis"`
}

// LogConfig 日志配置
type LogConfig struct {
	Level      string `json:"level"`
	Filename   string `json:"filename"`
	MaxSize    int    `json:"maxsize"`
	MaxAge     int    `json:"max_age"`
	MaxBackups int    `json:"max_backups"`
}

type MysqlConfig struct {
	DbHost    string `json:"db_host"`
	DbPort    int64  `json:"db_port"`
	DbUser    string `json:"db_user"`
	DbPwd     string `json:"db_pwd"`
	DbName    string `json:"db_name"`
	DbMaxConn int    `json:"db_max_conn"`
	DbMaxidle int    `json:"db_maxidle"`
}

type EsConfig struct {
	Host string `json:"host"`
	Port int64  `json:"port"`
}

type RedisConfig struct {
	Host string `json:"host"`
	Port int64  `json:"port"`
}

// Conf 全局配置变量
var Conf = new(Config)

// Init 初始化配置；从指定文件加载配置文件
func Init(flag bool) error {
	if flag {
		err := json.Unmarshal(getConfig(), Conf)
		if err != nil {
			zap.S().Errorf("get conf for etcd failed, err : %v", err)
			panic(err)
		}
		return nil
	} else {
		path, _ := os.Getwd()
		fmt.Println(path)
		bys, err := ioutil.ReadFile("../conf/conf.json")
		if err != nil {
			zap.S().Errorf("get conf for etcd failed, err : %v", err)
			panic(err)
		}
		err = json.Unmarshal(bys, Conf)
		if err != nil {
			zap.S().Errorf("get conf for etcd failed, err : %v", err)
			panic(err)
		}
		return nil
	}
}

func getConfig() (res []byte) {
	get, err := etcd.Get(context.Background(), "conf")
	if err != nil {
		panic(err)
	}
	for _, v := range get.Kvs {
		res = v.Value
	}
	return
}
