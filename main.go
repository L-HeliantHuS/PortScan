package main

import (
	"PortScan/config"
	"PortScan/utils"
	"bufio"
	"io"
	"os"
	"strings"
	"sync"
)

var mainChannel = make(chan string, config.ThreadCount)
var temp string
var wg sync.WaitGroup

func NewScan(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic("读取文件失败, 请检查路径是否正确")
	}

	reader := bufio.NewReader(file)
	go GOScan(mainChannel)
	for {
		ip, err := reader.ReadString('\n')
		if err != nil || err == io.EOF {
			break
		}
		ip = strings.TrimSpace(ip)
		mainChannel <- ip
	}
}

func GOScan(channel chan string)  {
	for {
		ip := <- channel
		wg.Add(1)
		go utils.Scan(ip, &wg)
	}
}

func main() {
	// 初始化配置
	config.Init()


	NewScan("ip.txt")


	wg.Wait()
}
