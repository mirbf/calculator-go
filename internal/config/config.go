// 配置管理模块
package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

// Config 应用配置结构
type Config struct {
	WindowWidth  int      `json:"window_width"`
	WindowHeight int      `json:"window_height"`
	Theme        string   `json:"theme"`
	History      []string `json:"history"`
}

// DefaultConfig 返回默认配置
func DefaultConfig() *Config {
	return &Config{
		WindowWidth:  400,
		WindowHeight: 600,
		Theme:        "light",
		History:      make([]string, 0),
	}
}

// Save 保存配置到文件
func (c *Config) Save() error {
	configDir, err := getConfigDir()
	if err != nil {
		return fmt.Errorf("获取配置目录失败: %v", err)
	}

	// 确保配置目录存在
	if err := os.MkdirAll(configDir, 0755); err != nil {
		return fmt.Errorf("创建配置目录失败: %v", err)
	}

	configFile := filepath.Join(configDir, "config.json")
	data, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return fmt.Errorf("序列化配置失败: %v", err)
	}

	if err := os.WriteFile(configFile, data, 0644); err != nil {
		return fmt.Errorf("写入配置文件失败: %v", err)
	}

	return nil
}

// Load 从文件加载配置
func Load() (*Config, error) {
	configDir, err := getConfigDir()
	if err != nil {
		return DefaultConfig(), nil // 如果无法获取配置目录，返回默认配置
	}

	configFile := filepath.Join(configDir, "config.json")
	data, err := os.ReadFile(configFile)
	if err != nil {
		if os.IsNotExist(err) {
			return DefaultConfig(), nil // 配置文件不存在，返回默认配置
		}
		return nil, fmt.Errorf("读取配置文件失败: %v", err)
	}

	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("解析配置文件失败: %v", err)
	}

	return &config, nil
}

// getConfigDir 获取配置目录路径
func getConfigDir() (string, error) {
	var configDir string

	switch runtime.GOOS {
	case "windows":
		appData := os.Getenv("APPDATA")
		if appData == "" {
			return "", fmt.Errorf("无法获取 APPDATA 环境变量")
		}
		configDir = filepath.Join(appData, "calculator-go")
	case "darwin":
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return "", fmt.Errorf("无法获取用户主目录: %v", err)
		}
		configDir = filepath.Join(homeDir, "Library", "Application Support", "calculator-go")
	default: // Linux 和其他 Unix 系统
		configHome := os.Getenv("XDG_CONFIG_HOME")
		if configHome == "" {
			homeDir, err := os.UserHomeDir()
			if err != nil {
				return "", fmt.Errorf("无法获取用户主目录: %v", err)
			}
			configHome = filepath.Join(homeDir, ".config")
		}
		configDir = filepath.Join(configHome, "calculator-go")
	}

	return configDir, nil
}

// GetDataDir 获取数据目录路径
func GetDataDir() (string, error) {
	// 对于这个应用，数据目录和配置目录相同
	return getConfigDir()
}

// GetHistoryFilePath 获取历史文件路径
func GetHistoryFilePath() (string, error) {
	dataDir, err := GetDataDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(dataDir, "history.json"), nil
}