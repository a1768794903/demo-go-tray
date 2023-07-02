package config

import (
	"demo-go-tray/utils"
	"fmt"
	"github.com/spf13/viper"
	"log"
	"testing"
)

func TestName(t *testing.T) {
	// 创建一个新的 Viper 实例
	v := viper.New()

	// 设置要读取的配置文件名和路径
	v.SetConfigFile("./config.toml")

	// 读取配置文件并处理错误
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("无法加载配置文件：%s", err))
	}
	// 加载配置文件内容到结构体对象
	if err := v.Unmarshal(&Cfg); err != nil {
		log.Panic("配置文件内容加载失败: ", err)
	}
	width, heigth := utils.GetScreenSize()
	url := fmt.Sprintf(Cfg.Url.Picsum, width, heigth)
	fmt.Println(url)
}
