package config

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

// Config 结构体表示配置文件的结构
type Config struct {
	AppID     string `yaml:"appid"`
	APISecret string `yaml:"apisecret"`
	APIKey    string `yaml:"apiKey"`
	HostURL   string `yaml:"hosturl"`
	DoMain    string `yaml:"domain"`
}

// LoadConfig 从 YAML 文件加载配置
func LoadConfig(filePath string) (*Config, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("无法读取配置文件: %v", err)
	}

	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("无法解析配置文件: %v", err)
	}

	return &config, nil
}

// InitConfig 初始化配置
func InitConfig(configPath string) *Config {
	config, err := LoadConfig(configPath)
	if err != nil {
		log.Fatalf("配置加载失败: %v", err)
	}
	return config
}
