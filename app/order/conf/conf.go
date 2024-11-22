package conf

import (
	"github.com/bytedance/go-tagexpr/v2/validator"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"net"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/kr/pretty"
)

var (
	conf *Config
	once sync.Once
)

type Config struct {
	Env      string
	Kitex    Kitex    `yaml:"kitex"`
	MySQL    MySQL    `yaml:"mysql"`
	Redis    Redis    `yaml:"redis"`
	Registry Registry `yaml:"registry"`
}

type MySQL struct {
	DSN string `yaml:"dsn"`
}

type Redis struct {
	Address  string `yaml:"address"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

type Kitex struct {
	Service       string `yaml:"service"`
	Address       string `yaml:"address"`
	LogLevel      string `yaml:"log_level"`
	LogFileName   string `yaml:"log_file_name"`
	LogMaxSize    int    `yaml:"log_max_size"`
	LogMaxBackups int    `yaml:"log_max_backups"`
	LogMaxAge     int    `yaml:"log_max_age"`
}

type Registry struct {
	RegistryAddress []string `yaml:"registry_address"`
	Username        string   `yaml:"username"`
	Password        string   `yaml:"password"`
}

// GetConf gets configuration instance
func GetConf() *Config {
	once.Do(initConf)
	return conf
}

// 新增：获取本机的有效 IP 地址
func GetLocalIP() string { // 新增
	addrs, err := net.InterfaceAddrs() // 新增：获取所有网络接口地址
	if err != nil {                    // 新增：错误处理
		klog.Error("Error getting local IP: %v", err) // 新增
		return ""                                     // 新增
	}

	// 遍历所有接口地址
	for _, addr := range addrs { // 新增
		// 判断是否是有效的 IPv4 地址且非回环地址
		if ip, ok := addr.(*net.IPNet); ok && !ip.IP.IsLoopback() && ip.IP.To4() != nil { // 新增
			// 过滤掉自动私有 IP 地址（169.254.x.x）
			if ip.IP.IsPrivate() { // 新增
				return ip.IP.String() // 新增
			}
		}
	}
	return "" // 新增
}

func initConf() {
	prefix := "conf"
	confFileRelPath := filepath.Join(prefix, filepath.Join(GetEnv(), "conf.yaml"))
	content, err := ioutil.ReadFile(confFileRelPath)
	if err != nil {
		panic(err)
	}
	conf = new(Config)
	err = yaml.Unmarshal(content, conf)
	if err != nil {
		klog.Error("parse yaml error - %v", err)
		panic(err)
	}
	if err := validator.Validate(conf); err != nil {
		klog.Error("validate config error - %v", err)
		panic(err)
	}

	// 新增：获取本机 IP 并更新到配置中（仅设置 IP，保留原有端口）
	localIP := GetLocalIP() // 调用获取本机 IP 的方法
	if localIP != "" {      // 判断 IP 是否有效
		conf.Kitex.Address = localIP + ":" + conf.Kitex.Address[strings.LastIndex(conf.Kitex.Address, ":")+1:] // 替换地址中的 IP 部分
	} else { // 处理无法获取 IP 的情况
		klog.Warn("Failed to get local IP, using default config")
	}

	conf.Env = GetEnv()
	pretty.Printf("%+v\n", conf)
}

// GetEnv 获取环境变量（默认为 test）
func GetEnv() string {
	e := os.Getenv("GO_ENV")
	if len(e) == 0 {
		return "test"
	}
	return e
}

// LogLevel 获取日志级别
func LogLevel() klog.Level {
	level := GetConf().Kitex.LogLevel
	switch level {
	case "trace":
		return klog.LevelTrace
	case "debug":
		return klog.LevelDebug
	case "info":
		return klog.LevelInfo
	case "notice":
		return klog.LevelNotice
	case "warn":
		return klog.LevelWarn
	case "error":
		return klog.LevelError
	case "fatal":
		return klog.LevelFatal
	default:
		return klog.LevelInfo
	}
}
