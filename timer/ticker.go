package timer

import (
	"demo-go-tray/signal"
	"fmt"
	"sync"
	"time"
)

var (
	timer     *time.Timer
	timerLock sync.Mutex
)

func SetTimer(duration time.Duration) {
	timerLock.Lock()
	defer timerLock.Unlock()

	if timer != nil {
		timer.Stop()
	}

	timer = time.AfterFunc(duration, func() {
		fmt.Println(duration, "秒后定时器触发")
		sig := signal.GetSignal()
		sig.SendSignal()
	})
}
