package config

import (
	"github.com/spf13/viper"
	"sync"
)

var (
	confOnce sync.Once
	config   *Config
)

type Config struct {
	Host     string         `json:"host"`
	Port     string         `json:"port"`
	RPCPort  string         `json:"rpcport"`
	Develop  bool           `json:"develop"`
	Log      LoggerConfig   `json:"log"`
	Database DatabaseConfig `json:"database"`
}

// database config
type DatabaseConfig struct {
	Source  string `json:"source"`
	Driver  string `json:"driver"`
	MaxOpen int    `json:"maxOpen"`
	MaxIdle int    `json:"maxIdle"`
}

// logger config
type LoggerConfig struct {
	Type  string `json:"type"`  // options: file, stdout
	Level string `json:"level"` // debug, info, error...
}

type KafkaWriteConfig struct {
	Brokers []string `json:"brokers"`
	Topic   string   `json:"topic"`
}

type KafkaReadConfig struct {
	Brokers []string `json:"brokers"`
	Topic   string   `json:"topic"`
	GroupID string   `json:"groupId"`
	Timeout int64    `json:"timeout"`
}

// config 파일(yaml)을 읽고 global struct 에 저장한다.
func LoadConfigFile() *Config {
	filename := "env.yaml"
	confOnce.Do(func() {
		viper.SetConfigType("yaml")
		viper.SetConfigFile(filename)
		err := viper.ReadInConfig() // Find and read the config file
		if err != nil {             // Handle errors reading the config file
			panic(err)
		}
		err = viper.Unmarshal(&config)

		if err != nil {
			panic(err)
		}
	})
	return config
}
