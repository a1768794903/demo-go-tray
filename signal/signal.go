package signal

import (
	"demo-go-tray/download"
	"demo-go-tray/utils"
	"fmt"
	"time"
)

var downloadSignal *Signal

// MySignal 自定义结构体类型作为信号
type Signal struct {
	SignalChan chan struct{}
}

func newSignal() *Signal {
	return &Signal{
		SignalChan: make(chan struct{}, 0),
	}
}

func GetSignal() *Signal {
	if downloadSignal == nil {
		downloadSignal = newSignal()
	}
	return downloadSignal
}

// ListenSignal 监听信号的函数
func (sig *Signal) ListenSignal() {
	go func() {
		defer close(sig.SignalChan)
		for {
			select {
			case <-sig.SignalChan:
				fmt.Println("收到信号")
				url := utils.GetUrl()
				err := download.LoadImage(url)
				if err != nil {
					fmt.Printf("下载图片失败，err:%+v", err)
				}

			default:
				// 在没有信号时执行的操作
				fmt.Println("等待信号...")
				time.Sleep(time.Second)
			}
		}
	}()
}

// SendSignal 发送信号的函数
func (sig *Signal) SendSignal() {
	sig.SignalChan <- struct{}{}
}

func (sig *Signal) close() {
	close(sig.SignalChan)
}
func Close() {
	downloadSignal.close()
}
