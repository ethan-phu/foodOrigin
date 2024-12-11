package config

import (
	"log"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var GlobalConfig *Config

// Config is application global config
type Config struct {
	Mode         string      `mapstructure:"mode"`           // gin启动模式
	Port         string      `mapstructure:"port"`           // 启动端口
	AppName      string      `mapstructure:"app-name"`       //应用名称
	Url          string      `mapstructure:"url"`            // 应用地址,用于自检 eg. http://127.0.0.1
	MaxPingCount int         `mapstructure:"max-ping-count"` // 最大自检次数，用户健康检查
	JwtSecret    string      `mapstructure:"jwt-secret"`
	Language     string      `mapstructure:"language"` // 项目语言
	DBConfig     DBConfig    `mapstructure:"database"` // 数据库信息
	RedisConfig  RedisConfig `mapstructure:"redis"`    // redis
	LogConfig    LogConfig   `mapstructure:"log"`      // uber zap
	Wechat       WechatConfig `mapstructure:"wechat"`  // 微信小程序配置
}

// DBConfig is used to configure mysql database
type DBConfig struct {
	Dbname          string `mapstructure:"dbname"`
	Host            string `mapstructure:"host"`
	Port            string `mapstructure:"port"`
	Username        string `mapstructure:"username"`
	Password        string `mapstructure:"password"`
	MaximumPoolSize int    `mapstructure:"maximum-pool-size"`
	MaximumIdleSize int    `mapstructure:"maximum-idle-size"`
	LogMode         bool   `mapstructure:"log-mode"`
	Sqlite          bool   `mapstructure:"sqlite"`
}

// RedisConfig is used to configure redis
type RedisConfig struct {
	Addr         string `mapstructure:"address"`
	Password     string `mapstructure:"password"`
	Db           int    `mapstructure:"db"`
	PoolSize     int    `mapstructure:"pool-size"`
	MinIdleConns int    `mapstructure:"min-idle-conns"`
	IdleTimeout  int    `mapstructure:"idle-timeout"`
}

// LogConfig is used to configure uber zap
type LogConfig struct {
	Level      string `mapstructure:"level"`
	FileName   string `mapstructure:"file-name"`
	TimeFormat string `mapstructure:"xtime-format"`
	MaxSize    int    `mapstructure:"max-size"`
	MaxBackups int    `mapstructure:"max-backups"`
	MaxAge     int    `mapstructure:"max-age"`
	Compress   bool   `mapstructure:"compress"`
	LocalTime  bool   `mapstructure:"local-xtime"`
	Console    bool   `mapstructure:"console"`
}

// WechatConfig 微信小程序配置
type WechatConfig struct {
	AppID     string `mapstructure:"app-id"`     // 小程序AppID
	AppSecret string `mapstructure:"app-secret"` // 小程序AppSecret
}

// Load is a loader to load config file.
func Load(configFilePath string) *Config {
	resolveRealPath(configFilePath)
	// 初始化配置文件
	if err := initConfig(); err != nil {
		panic(err)
	}
	// 监控配置文件，并热加载
	watchConfig()

	return GlobalConfig
}

func initConfig() error {
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	viper.SetEnvPrefix("APPLICATION")
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	// 解析配置文件
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	// 解析到struct
	GlobalConfig = &Config{}
	if err := viper.Unmarshal(GlobalConfig); err != nil {
		panic(err)
	}
	log.Println("The application configuration file is loaded successfully!")
	return nil
}

// 监控配置文件变动
// 注意：有些配置修改后，及时重新加载也要重新启动应用才行，比如端口
func watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		log.Printf("Configuration file changed: %s, reload it", in.Name)
		// 忽略错误
		// 配置文件的路径
		Load(in.Name)
	})
}

// 如果未传递配置文件路径将使用约定的环境配置文件
func resolveRealPath(path string) {
	if path != "" {
		viper.SetConfigFile(path)
	} else {
		// 设置默认的config
		// 路径
		viper.AddConfigPath("conf")
		// 配置文件的名字
		viper.SetConfigName("config")
	}
}
