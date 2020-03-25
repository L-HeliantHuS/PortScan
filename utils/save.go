package utils

import (
	"PortScan/config"
	"fmt"
	"sync"
)

var (
	Lock sync.Mutex
)

func SaveToFile(ip string, port int, service string) {

	Lock.Lock()
	if service == "RDP" {
		config.RDPFileHandler.Write([]byte(fmt.Sprintf("%s:%d \n", ip, port)))
	} else if service == "SSH" {
		config.SSHFileHandler.Write([]byte(fmt.Sprintf("%s:%d \n", ip, port)))
	}
	Lock.Unlock()

	fmt.Printf("[+] %s -- %s:%d \n", service, ip, port)
}
