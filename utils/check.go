package utils

import (
	"PortScan/config"
	"encoding/hex"
	"errors"
)

func Check(recvData []byte) (string, error) {
	// 检查是否是SSH
	if string(recvData[:3]) == "SSH" {
		return "SSH", nil
	}

	stringData := hex.EncodeToString(recvData)
	if _, ok := config.OSMap[stringData]; ok {
		return "RDP", nil
	}

	return "", errors.New("没有匹配到")
}