package conf

import (
	"github.com/BurntSushi/toml"
)

var (
	global *Config
)

// Init 初始化配置
func Init(fpath string) {
	c, err := parse(fpath)
	if err != nil {
		panic("init config failed, " + err.Error())
	}
	global = c
}

// Global 获取全局配置
func Global() *Config {
	if global == nil {
		return &Config{}
	}
	return global
}

// Parse 解析配置文件
func parse(fpath string) (*Config, error) {
	var c Config
	_, err := toml.DecodeFile(fpath, &c)
	if err != nil {
		return nil, err
	}
	return &c, nil
}

// Config 配置参数
type Config struct {
	Version string  `toml:"version"`
	RunMode string  `toml:"run_mode"`
	WWW     string  `toml:"www"`
	Swagger Swagger `toml:"swagger"`
	HTTP    HTTP    `toml:"http"`
	Log     Log     `toml:"log"`
	CORS    CORS    `toml:"cors"`
}

// Swagger 接口文档配置
type Swagger struct {
	On          int      `toml:"on"`
	Title       string   `toml:"title"`
	Description string   `toml:"description"`
	Version     string   `toml:"version"`
	Host        string   `toml:"host"`
	BasePath    string   `toml:"base_path"`
	Schemes     []string `toml:"schemes"`
}

// Log 日志配置参数
type Log struct {
	AppNo                  int    `toml:"app_no"`
	AppName                string `toml:"app_name"`
	Level                  int    `toml:"log_level"`
	Format                 string `toml:"format"`
	Output                 string `toml:"output"`
	OutputFile             string `toml:"output_file"`
	DisableCustomTimestamp bool   `toml:"disable_custom_timestamp"`
	DisableLineHook        bool   `toml:"disable_line_hook"`
	MaxAge                 int    `toml:"log_file_max_age"`
	RotationTime           int    `toml:"log_file_rotation_time"`
}

// HTTP http配置参数
type HTTP struct {
	Host            string `toml:"host"`
	Port            int    `toml:"port"`
	CertFile        string `toml:"cert_file"`
	KeyFile         string `toml:"key_file"`
	ShutdownTimeout int    `toml:"shutdown_timeout"`
}

// CORS 跨域请求配置参数
type CORS struct {
	Enable           bool     `toml:"enable"`
	AllowOrigins     []string `toml:"allow_origins"`
	AllowMethods     []string `toml:"allow_methods"`
	AllowHeaders     []string `toml:"allow_headers"`
	AllowCredentials bool     `toml:"allow_credentials"`
	MaxAge           int      `toml:"max_age"`
}
