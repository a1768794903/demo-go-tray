package utils

import (
	"demo-go-tray/config"
	"fmt"
	"testing"
)

func TestGetSystemMetrics(t *testing.T) {
	width := GetSystemMetrics(0)
	height := GetSystemMetrics(1)
	width = int(float32(width) * 1.5)
	height = int(float32(height) * 1.5)
	fmt.Println("wys:", width)
	fmt.Println("wys:", height)
}

func TestGetScreenSize(t *testing.T) {
	width, height := GetScreenSize()

	fmt.Println("wys:", width)
	fmt.Println("wys:", height)
}

func TestInitVip(t *testing.T) {
	InitVip()
	fmt.Println(config.Cfg)
}
