/*
@author:Deng.l.w
@version:1.20
@date:2023-02-23 18:12
@file:settings.go
*/

package settings

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// Conf 全局变量，保存储层序所有配置信息
var Conf = new(MultiConfig)

type MultiConfig struct {
	*AppConfig   `mapstructure:"app"`
	*LogConfig   `mapstructure:"log"`
	*MySOLConfig `mapstructure:"mysql"`
	*RedisConfig `mapstructure:"redis"`
}

type AppConfig struct {
	Name      string `mapstructure:"name"`
	Mode      string `mapstructure:"mode"`
	Version   string `mapstructure:"version"`
	Port      int    `mapstructure:"port"`
	StartTime string `mapstructure:"start_time"`
	MachineID int64  `mapstructure:"machine_id"`
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}

type MySOLConfig struct {
	Host        string `mapstructure:"host"`
	User        string `mapstructure:"user"`
	Password    string `mapstructure:"password"`
	DbName      string `mapstructure:"dbname"`
	Port        int    `mapstructure:"port"`
	MaxOpenConn int    `mapstructure:"max_open_conn"`
	MaxIdleConn int    `mapstructure:"max_idle_conn"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Password string `mapstructure:"password"`
	Port     int    `mapstructure:"port"`
	DB       int    `mapstructure:"db"`
	PoolSize int    `mapstructure:"pool_size"`
}

func Init(filePath string) (err error) {
	//viper.SetConfigFile("F:/go_project/web_app/config.yaml")
	//viper.SetConfigName("./config")             //配置文件名称（无扩展名）
	//viper.AddConfigPath(".")                    // 配置搜索路径（相对路径），可配置多个
	viper.SetConfigFile(filePath)
	if err = viper.ReadInConfig(); err != nil { // 读取配置文件
		fmt.Printf("failed,%v\n", err)
		return
	}

	// 把读取到的配置信息反序列化到conf 变量中
	if err := viper.Unmarshal(Conf); err != nil {
		fmt.Printf("unmarshal failed,err:%v\n", err)
	}

	viper.WatchConfig()                            // 监控
	viper.OnConfigChange(func(in fsnotify.Event) { //回调函数
		fmt.Println("配置文件已修改")
		if err := viper.Unmarshal(Conf); err != nil {
			fmt.Printf("unmarshal failed,err:%v\n", err)
		}
	})
	return
}
