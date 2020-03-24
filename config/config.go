package config

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

// 验证数据包
var (
	OSMap             map[string]string // RDP OS列表
	Timeouted         time.Duration     // 扫描超时时间
	Ports             []int
	ThreadCount       int
	ResultFileHandler *os.File
)

func Init() {
	OSMap = map[string]string{
		"030000130ed000001234000200080002000000": "2008",
		"030000130ed000001234000209080002000000": "win7OR2008R2",
		"030000130ed000001234000201080002000000": "2008R2DC",
		"030000130ed00000123400020f080002000000": "2012R2OR8",
		"030000130ed00000123400021f080002000000": "2016",
	}

	envTimeouted := os.Getenv("timeout")
	if envTimeouted != "" {
		number, _ := strconv.Atoi(envTimeouted)

		Timeouted = time.Duration(number)
	} else {
		Timeouted = 2
	}


	Ports = []int{22, 3389}

	envThreadCount := os.Getenv("goroutine_count")
	if envThreadCount != "" {
		number, _ := strconv.Atoi(envThreadCount)

		ThreadCount = number
	} else {
		ThreadCount = 5000
	}

	file, err := os.OpenFile("success.txt", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}

	ResultFileHandler = file

	fmt.Printf(`
		当前设置: 
			超时时间: %d秒;
			最大协程数: %d;
				
				将在3秒后开始...
	`, Timeouted, ThreadCount)
	fmt.Println()
	time.Sleep(3 * time.Second)
}
