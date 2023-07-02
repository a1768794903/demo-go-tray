package download

import (
	"demo-go-tray/utils"
	"fmt"
	"testing"
)

func TestDownloadImage(t *testing.T) {
	//width := utils.GetSystemMetrics(0)
	//height := utils.GetSystemMetrics(1)
	width, height := utils.GetScreenSize()
	url := fmt.Sprintf("https://picsum.photos/%d/%d", width, height)
	err := LoadImage(url)
	if err != nil {
		println(err)
	}
}
