package main

import (
	"demo-go-tray/signal"
	"demo-go-tray/timer"
	"demo-go-tray/utils"
	"fmt"
	"github.com/riftbit/go-systray"
	"io/ioutil"
	"time"
)

// go build -ldflags="-H windowsgui"
//var (
//	//url string = "https://source.unsplash.com/random/1980x1080"
//	//url string = "https://picsum.photos/%s/%s"
//)

func main() {
	//systray.SetCustomLeftClickAction()
	//systray.SetCustomRightClickAction()
	utils.InitConfig()
	sig := signal.GetSignal()
	sig.ListenSignal()
	err := systray.Run(onReady, onExit)
	if err != nil {
		println(err)
	}
}

func onReady() {
	// 设置图标
	err := systray.SetIcon(getIcon("./desktop.ico"))
	if err != nil {
		println(err)
	}
	reFlash := systray.AddMenuItem("立即刷新", "立即刷新", 0)

	flashTime := systray.AddSubMenu("刷新间隔")
	timeHalfHour := flashTime.AddSubMenuItem("30分钟", "", 0)
	timeOneHour := flashTime.AddSubMenuItem("1个小时", "", 0)
	timeTwoHour := flashTime.AddSubMenuItem("2个小时", "", 0)
	timeFourHour := flashTime.AddSubMenuItem("4个小时", "", 0)

	pictureSrc := systray.AddSubMenu("壁纸源")
	Unsplash := pictureSrc.AddSubMenuItem("Unsplash", "", 0)
	Picsum := pictureSrc.AddSubMenuItem("Picsum", "", 0)

	sort := systray.AddSubMenu("分类")
	all := sort.AddSubMenuItem("所有", "", 0)
	wallpaper := sort.AddSubMenuItem("壁纸", "", 0)
	person := sort.AddSubMenuItem("人物", "", 0)
	textures := sort.AddSubMenuItem("纹理", "", 0)
	nature := sort.AddSubMenuItem("自然", "", 0)
	architecture := sort.AddSubMenuItem("建筑", "", 0)
	film := sort.AddSubMenuItem("电影", "", 0)
	animals := sort.AddSubMenuItem("动物", "", 0)
	travel := sort.AddSubMenuItem("旅行", "", 0)
	food := sort.AddSubMenuItem("食物", "", 0)

	save := systray.AddMenuItem("保存", "保存", 0)
	about := systray.AddMenuItem("关于", "关于", 0)

	systray.AddSeparator()

	mQuit := systray.AddMenuItem("退出", "退出", 0)

	//go func() {
	//	for {
	//		systray.SetTitle(timezone)
	//		systray.SetTooltip(timezone)
	//		time.Sleep(1 * time.Second)
	//	}
	//}()

	go func() {
		for {
			select {
			case <-reFlash.OnClickCh():
				sig := signal.GetSignal()
				sig.SendSignal()
			case <-timeHalfHour.OnClickCh():
				timer.SetTimer(30 * time.Minute)
			case <-timeOneHour.OnClickCh():
				timer.SetTimer(1 * time.Hour)
			case <-timeTwoHour.OnClickCh():
				timer.SetTimer(2 * time.Hour)
			case <-timeFourHour.OnClickCh():
				timer.SetTimer(4 * time.Hour)

			case <-Unsplash.OnClickCh():
				utils.SetUrl("Unsplash")
			case <-Picsum.OnClickCh():
				utils.SetUrl("Picsum")

			case <-all.OnClickCh():
				utils.SetKeyWord("all")
			case <-wallpaper.OnClickCh():
				utils.SetKeyWord("wallpaper")
			case <-person.OnClickCh():
				utils.SetKeyWord("person")
			case <-textures.OnClickCh():
				utils.SetKeyWord("textures")
			case <-nature.OnClickCh():
				utils.SetKeyWord("nature")
			case <-architecture.OnClickCh():
				utils.SetKeyWord("architecture")
			case <-film.OnClickCh():
				utils.SetKeyWord("film")
			case <-animals.OnClickCh():
				utils.SetKeyWord("animals")
			case <-travel.OnClickCh():
				utils.SetKeyWord("travel")
			case <-food.OnClickCh():
				utils.SetKeyWord("food")

			case <-save.OnClickCh():
			case <-about.OnClickCh():

			case <-mQuit.OnClickCh():
				signal.Close()
				systray.Quit()
				return
			}
		}
	}()
}

func onExit() {
	// 清除销毁
}

func getIcon(s string) []byte {
	b, err := ioutil.ReadFile(s)
	if err != nil {
		fmt.Print(err)
	}
	return b
}
