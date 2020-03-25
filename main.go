package main

import (
	"PortScan/config"
	"PortScan/utils"
	"bufio"
	"flag"
	"io"
	"os"
	"runtime"
	"strings"
	"sync"
	"time"
)

var mainChannel = make(chan string, config.ThreadCount)
var wg sync.WaitGroup

func NewScan(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic("Read file error, please check filepath.")
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

	// 每两秒检测一次goroutine数量
	for {
		if runtime.NumGoroutine() < config.ThreadCount * 2 {
			ip := <- channel
			wg.Add(1)
			go utils.Scan(ip, &wg)
		} else {
			time.Sleep(2 * time.Second)
		}
	}
}

func main() {

	flag.IntVar(&config.ThreadCount, "thread", 30000, "thread")
	flag.IntVar(&config.Timeouted, "timeout", 1, "timeout")
	flag.Parse()

	// 初始化配置
	config.Init()

	NewScan("ip.txt")

	wg.Wait()
}
