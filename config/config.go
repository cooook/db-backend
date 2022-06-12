package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Name string
}

// * 初始化配置文件
func Init(cfg string) error {
	c := Config{Name: cfg}
	if err := c.initConfig(); err != nil {
		return err
	}
	return nil
}

func (c *Config) initConfig() error {
	// * 指定配置文件
	if c.Name != "" {
		// * 设置配置文件名
		viper.SetConfigFile(c.Name)
	} else {
		// * 未指定配置文件，读取默认路径
		viper.AddConfigPath(".")
		viper.SetConfigName("config")
	}
	// * 设置config文件类型
	viper.SetConfigType("yaml")

	// * 读取配置文件内容，使用viper.get获取配置
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return nil
}
