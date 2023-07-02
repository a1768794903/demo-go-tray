package config

var Cfg Config

// https://picsum.photos/id/123/200/300
// https://source.unsplash.com/collection/c1jfZtifQgU/%dx%d
// https://unsplash.it/2560/1440?random

// 配置文件的结构体
type Config struct {
	Url Url
}
type Url struct {
	Unsplash string
	Picsum   string
}
