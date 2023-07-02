package timer

import (
	"testing"
	"time"
)

func TestSetTimer(t *testing.T) {
	SetTimer(10)

	SetTimer(2)
	time.Sleep(10 * time.Second)
}
