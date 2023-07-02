package download

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

var (
	filepath = "C:/Config/test.jpg"
)

func LoadImage(url string) error {
	// 发起 HTTP GET 请求
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	// 检查 HTTP 响应状态码
	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("下载失败，HTTP 状态码：%d", response.StatusCode)
	}

	// 创建文件
	file, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	// 将 HTTP 响应的内容写入文件
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	fmt.Println("图片下载完成:", filepath)
	return nil
}
