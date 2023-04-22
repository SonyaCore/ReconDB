package check

import (
	"fmt"
	"net"
	"strconv"
	"strings"
)

func CheckIPAddress(ip string) bool {
	if strings.Contains(ip, ":") {
		parts := strings.Split(ip, ":")
		ip = parts[0]
		port := parts[1]

		err := CheckPort(port)
		if err != nil {
			return false
		}
	}

	if net.ParseIP(ip) == nil {
		return false
	}
	return true
}

func CheckPort(port string) error {
	portNum, err := strconv.Atoi(port)
	if err != nil {
		return fmt.Errorf("invalid port number: %s", port)
	}
	if portNum < 1 || portNum > 65535 {
		return fmt.Errorf("port number out of range: %d", portNum)
	}
	return nil
}
