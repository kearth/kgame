package utils

import (
	"encoding/json"
	"os"
)

// Config 配置结构体
type Config struct {
	GameName     string `json:"game_name"`
	MaxPlayers   int    `json:"max_players"`
	MinPlayers   int    `json:"min_players"`
	LogLevel     int    `json:"log_level"`
	MaxLevel     int    `json:"max_level"`
	DefaultRules bool   `json:"default_rules"`
}

// DefaultConfig 默认配置
var DefaultConfig = Config{
	GameName:     "升官图",
	MaxPlayers:   4,
	MinPlayers:   2,
	LogLevel:     LogLevelInfo,
	MaxLevel:     10,
	DefaultRules: true,
}

// LoadConfig 从文件加载配置
func LoadConfig(filePath string) (*Config, error) {
	// 尝试读取配置文件
	file, err := os.Open(filePath)
	if err != nil {
		// 如果文件不存在，返回默认配置
		if os.IsNotExist(err) {
			return &DefaultConfig, nil
		}
		return nil, err
	}
	defer file.Close()

	// 解析配置文件
	var config Config
	if err := json.NewDecoder(file).Decode(&config); err != nil {
		return nil, err
	}

	return &config, nil
}

// SaveConfig 保存配置到文件
func SaveConfig(config *Config, filePath string) error {
	// 创建或打开配置文件
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// 编码配置文件
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(config)
}

// GetConfig 获取配置
func GetConfig() *Config {
	// 尝试从默认路径加载配置
	config, err := LoadConfig("./config.json")
	if err != nil {
		// 如果加载失败，返回默认配置
		return &DefaultConfig
	}

	return config
}
