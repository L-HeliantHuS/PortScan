package config

import (
	"fmt"
	"os"
	"time"
)

// 验证数据包
var (
	OSMap             map[string]string // RDP OS列表
	Timeouted         int     // 扫描超时时间
	Ports             []int
	ThreadCount       int
	SSHFileHandler *os.File
	RDPFileHandler *os.File
)

func Init() {
	OSMap = map[string]string{
		"030000130ed000001234000200080002000000": "2008",
		"030000130ed000001234000209080002000000": "win7OR2008R2",
		"030000130ed000001234000201080002000000": "2008R2DC",
		"030000130ed00000123400020f080002000000": "2012R2OR8",
		"030000130ed00000123400021f080002000000": "2016",
	}


	Ports = []int{22, 3389}

	ssh, err := os.OpenFile("SSH.txt", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}

	rdp, err := os.OpenFile("RDP.txt", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}

	SSHFileHandler = ssh
	RDPFileHandler = rdp

	fmt.Printf(`
		Config: 
			Timeout: %d s;
			Max Goroutine Count: %d;
				
	`, Timeouted, ThreadCount)
	fmt.Println()
	time.Sleep(3 * time.Second)
}
