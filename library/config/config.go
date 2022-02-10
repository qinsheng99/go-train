package config

import (
	"context"
	"encoding/json"

	"github.com/qinsheng99/goWeb/library/etcd"
	"go.uber.org/zap"
)

// Config 整个项目的配置
type Config struct {
	Mode       string `json:"mode"`
	Port       int    `json:"port"`
	*LogConfig `json:"log"`
}

// LogConfig 日志配置
type LogConfig struct {
	Level      string `json:"level"`
	Filename   string `json:"filename"`
	MaxSize    int    `json:"maxsize"`
	MaxAge     int    `json:"max_age"`
	MaxBackups int    `json:"max_backups"`
}

// Conf 全局配置变量
var Conf = new(Config)

// Init 初始化配置；从指定文件加载配置文件
func Init() error {
	err := json.Unmarshal(getConfig(), Conf)
	if err != nil {
		zap.S().Errorf("get conf for etcd failed, err : %v", err)
		panic(err)
	}
	return nil
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
