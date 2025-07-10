package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDefaultConfig(t *testing.T) {
	config := DefaultConfig()
	if config.WindowWidth != 400 {
		t.Errorf("Expected WindowWidth to be 400, got %d", config.WindowWidth)
	}
	if config.WindowHeight != 600 {
		t.Errorf("Expected WindowHeight to be 600, got %d", config.WindowHeight)
	}
	if config.Theme != "light" {
		t.Errorf("Expected Theme to be 'light', got '%s'", config.Theme)
	}
	if len(config.History) != 0 {
		t.Errorf("Expected History to be empty, got length %d", len(config.History))
	}
}

func TestConfigSaveAndLoad(t *testing.T) {
	// 创建临时目录用于测试
	tempDir := t.TempDir()
	originalGetConfigDir := getConfigDir
	
	// 模拟 getConfigDir 函数返回临时目录
	getConfigDir = func() (string, error) {
		return tempDir, nil
	}
	defer func() {
		getConfigDir = originalGetConfigDir
	}()

	// 创建测试配置
	config := &Config{
		WindowWidth:  800,
		WindowHeight: 1000,
		Theme:        "dark",
		History:      []string{"5 + 3 = 8", "10 - 2 = 8"},
	}

	// 保存配置
	err := config.Save()
	if err != nil {
		t.Fatalf("Failed to save config: %v", err)
	}

	// 验证配置文件是否存在
	configFile := filepath.Join(tempDir, "config.json")
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		t.Fatalf("Config file was not created")
	}

	// 加载配置
	loadedConfig, err := Load()
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}

	// 验证加载的配置
	if loadedConfig.WindowWidth != 800 {
		t.Errorf("Expected WindowWidth to be 800, got %d", loadedConfig.WindowWidth)
	}
	if loadedConfig.WindowHeight != 1000 {
		t.Errorf("Expected WindowHeight to be 1000, got %d", loadedConfig.WindowHeight)
	}
	if loadedConfig.Theme != "dark" {
		t.Errorf("Expected Theme to be 'dark', got '%s'", loadedConfig.Theme)
	}
	if len(loadedConfig.History) != 2 {
		t.Errorf("Expected History length to be 2, got %d", len(loadedConfig.History))
	}
}

func TestLoadNonExistentConfig(t *testing.T) {
	// 创建临时目录用于测试
	tempDir := t.TempDir()
	originalGetConfigDir := getConfigDir
	
	// 模拟 getConfigDir 函数返回临时目录
	getConfigDir = func() (string, error) {
		return tempDir, nil
	}
	defer func() {
		getConfigDir = originalGetConfigDir
	}()

	// 尝试加载不存在的配置文件
	config, err := Load()
	if err != nil {
		t.Fatalf("Expected no error when loading non-existent config, got: %v", err)
	}

	// 应该返回默认配置
	defaultConfig := DefaultConfig()
	if config.WindowWidth != defaultConfig.WindowWidth {
		t.Errorf("Expected default WindowWidth %d, got %d", defaultConfig.WindowWidth, config.WindowWidth)
	}
}

func TestGetDataDir(t *testing.T) {
	dataDir, err := GetDataDir()
	if err != nil {
		t.Fatalf("Failed to get data directory: %v", err)
	}
	if dataDir == "" {
		t.Error("Data directory should not be empty")
	}
}

func TestGetHistoryFilePath(t *testing.T) {
	historyPath, err := GetHistoryFilePath()
	if err != nil {
		t.Fatalf("Failed to get history file path: %v", err)
	}
	if historyPath == "" {
		t.Error("History file path should not be empty")
	}
	if filepath.Base(historyPath) != "history.json" {
		t.Errorf("Expected history file name to be 'history.json', got '%s'", filepath.Base(historyPath))
	}
}