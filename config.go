package chat

import "github.com/BurntSushi/toml"

// Config ... 設定全体
type Config struct {
	appName       string `toml:"app_name"`
	*ServerConfig `toml:"server"`
	*LogConfig    `toml:"logger"`
}

// ServerConfig ... サーバに関する設定
type ServerConfig struct {
	Host string
}

// LogConfig ... ログに関する設定
type LogConfig struct {
	Filepath string
	LogLevel string `toml:"log_level"`
}

// NewConfig ... 設定の取込
func NewConfig(configpath string) (*Config, error) {
	var config Config
	_, err := toml.DecodeFile(configpath, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
