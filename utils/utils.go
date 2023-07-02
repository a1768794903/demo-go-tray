package utils

import (
	"demo-go-tray/config"
	"demo-go-tray/global"
	"fmt"
	"github.com/lxn/win"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"log"
	"syscall"
	"unsafe"
)

func GetSystemMetrics(nIndex int) int {
	ret, _, _ := syscall.NewLazyDLL(`User32.dll`).NewProc(`GetSystemMetrics`).Call(uintptr(nIndex))
	return int(ret)
}

func InitVip() {
	// 创建一个新的 Viper 实例
	v := viper.New()

	// 设置要读取的配置文件名和路径
	v.SetConfigFile("../config/config.toml")

	// 读取配置文件并处理错误
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("无法加载配置文件：%s", err))
	}
	// 加载配置文件内容到结构体对象
	if err := v.Unmarshal(&config.Cfg); err != nil {
		log.Panic("配置文件内容加载失败: ", err)
	}
}

func InitConfig() {
	config.Cfg.Url.Picsum = "https://picsum.photos/%d/%d"
	config.Cfg.Url.Unsplash = "https://unsplash.it/%d/%d?random"
}

func GetScreenSize() (int, int) {
	width := int(win.GetSystemMetrics(win.SM_CXSCREEN))
	height := int(win.GetSystemMetrics(win.SM_CYSCREEN))
	width = int(float32(width) * 1.5)
	height = int(float32(height) * 1.5)
	fmt.Printf("%dx%d\n", width, height)
	return width, height
}
func SetUrl(Source string) {
	width, height := GetScreenSize()
	if Source == "Unsplash" {
		global.URL = fmt.Sprintf(config.Cfg.Url.Unsplash, width, height)
		if GetKeyWord() != "all" {
			global.URL += fmt.Sprintf("/%s", GetKeyWord())
		}
	} else if Source == "Picsum" {
		global.URL = fmt.Sprintf(config.Cfg.Url.Picsum, width, height)
	}
}
func GetUrl() (url string) {
	width, height := GetScreenSize()
	url = fmt.Sprintf(config.Cfg.Url.Unsplash, width, height)
	if len(global.URL) != 0 {
		url = global.URL
	}
	return url
}

func SetKeyWord(keyword string) {
	global.KeyWord = keyword
}
func GetKeyWord() (keyword string) {
	keyword = "all"
	if len(global.KeyWord) == 0 {
		return
	}
	keyword = global.KeyWord
	return keyword
}

// SetWindowsWallpaper 设置windows壁纸
func SetWindowsWallpaper(imagePath string) error {
	dll := syscall.NewLazyDLL("user32.dll")
	proc := dll.NewProc("SystemParametersInfoW")
	_t, _ := syscall.UTF16PtrFromString(imagePath)
	ret, _, _ := proc.Call(20, 1, uintptr(unsafe.Pointer(_t)), 0x1|0x2)
	if ret != 1 {
		return errors.New("系统调用失败")
	}
	return nil
}
