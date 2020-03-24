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
	config.ResultFileHandler.Write([]byte(fmt.Sprintf("%s:%d:%s \n", ip, port, service)))

	Lock.Unlock()
}
