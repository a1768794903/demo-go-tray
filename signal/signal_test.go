package signal

import (
	"testing"
	"time"
)

func TestListenSignal(t *testing.T) {
	sig := GetSignal()
	sig.ListenSignal()
	time.Sleep(2 * time.Second)
	sig.SendSignal()
	time.Sleep(5 * time.Second)
}
