package net


import (
	"errors"
	"net"
)

// 获取内网地址
func GetIntranetIp() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ip4 := ipnet.IP.To4(); ip4 != nil {
				if ip4[0] == 10 ||
					(ip4[0] == 172 && ip4[1] >= 16 && ip4[1] <= 31) ||
					(ip4[0] == 192 && ip4[1] == 168) {
					return ipnet.IP.String(), nil
				}
			}
		}
	}
	return "", errors.New("找不到服务器内网地址")
}
