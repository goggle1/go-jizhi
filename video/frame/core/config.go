package core

import (
	"encoding/json"

	log "github.com/micro/go-log"
	"github.com/toolkits/file"
)

type Database struct {
	Connect     string `json:"connect"`
	MaxIdle     int    `json:"maxIdle"`
	TablePrefix string `json:"table_prefix"` //表名前缀
}

type Log struct {
	FileName string `json:"filename"`
	Level    string `json:"level"`
	Format   string `json:"format"`
}

type BaseConfig struct {
	Debug        bool         `json:"debug"`
	Name         string       `json:"service_name"`    //服务名称
	Version      string       `json:"service_version"` //
	Etcd         []string     `json:"etcd"`
	Database     *Database    `json:"database"`
	Log          *Log         `json:"log"`
	Redis        *RedisConfig `json:"redis"`
	AdminAddress string       `json:"admin_address"`
}

type RedisConfig struct {
	Addr         string `json:"addr"`
	Password     string `json:"password"`
	Db           int    `json:"db"`
	MaxIdle      int    `json:"maxIdle"`
	ConnTimeout  int    `json:"connTimeout"`
	ReadTimeout  int    `json:"readTimeout"`
	WriteTimeout int    `json:"writeTimeout"`
}

func NewConfig(menu string) (*BaseConfig, error) {
	path := menu + "res/base.json"
	if !file.IsExist(path) {
		log.Fatal("config file:", path, "is not existent")
	}
	configContent, err := file.ToTrimString(path)
	if err != nil {
		log.Fatal("read config file:", path, "fail:", err)
		return nil, err
	}
	var c BaseConfig
	err = json.Unmarshal([]byte(configContent), &c)
	if err != nil {
		log.Fatal("parse config file:", path, "fail:", err)
		return nil, err
	}
	return &c, nil
}
