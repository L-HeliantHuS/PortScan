package utils

import (
	"PortScan/config"
	"errors"
	"fmt"
	"net"
	"sync"
	"time"
)

func ScanProcess(ip string, port int) ([]byte, error) {
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", ip, port), time.Second * config.Timeouted)
	if err != nil {
		return nil, errors.New("Connect failed.")
	}

	defer conn.Close()

	_, err = conn.Write([]byte("\x03\x00\x00\x13\x0e\xe0\x00\x00\x00\x00\x00\x01\x00\x08\x00\x03\x00\x00\x00"))
	if err != nil {
		return nil, errors.New("Send data failed.")
	}


	buf := make([]byte, 100)
	n, err := conn.Read(buf)
	if err != nil {
		return nil, errors.New("Read data failed.")
	}

	return buf[:n], nil
}

func Scan(ip string, group *sync.WaitGroup) {
	flag := false
	// 先进行常用端口探测
	for _, port := range config.Ports {
		data, err := ScanProcess(ip, port)
		if err != nil {
			continue
		}
		service, err := Check(data)
		if err == nil {
			fmt.Printf("[+] %s, Result: %s \n", ip, service)
			SaveToFile(ip, port, service)
			flag = true
			break
		}
	}

	if flag == false {
		for port := 1; port < 65535; port++ {
			// 跳过已经扫描的端口
			if port == 22 || port == 3389 {
				continue
			}

			data, err := ScanProcess(ip, port)
			if err != nil {
				continue
			}
			service, err := Check(data)
			if err == nil {
				fmt.Printf("[+] %s, Result: %s \n", ip, service)

				SaveToFile(ip, port, service)

				break
			}
		}
	}
	group.Done()

}