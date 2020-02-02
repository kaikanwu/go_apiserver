package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
	"strings"
)
// 结构体
type Config struct {
	Name string
}

func Init(cfg string) error {

	c := Config{
		Name: cfg,
	}

	// 初始化配置文件
	if err := c.initConfig();err != nil {
		return err
	}
	return nil

}

// 参数前置
// 初始化配置文件
func (c *Config) initConfig() error {
	if c.Name != "" {
		viper.SetConfigFile(c.Name)
	} else{
		viper.AddConfigPath("conf")
		viper.SetConfigName("config")
	}
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	viper.SetEnvPrefix("APISERVER")
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	if err := viper.ReadInConfig(); err!= nil {
		return err
	}
	return nil
}

// 监控配置文件，实现热加载
func (c *Config)watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Printf("Config file changed: #{e.Name}")
	})

}