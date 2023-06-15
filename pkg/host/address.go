package host

import (
	"fmt"
	"net"
	"strconv"
	"strings"
)

// IpAddress checks IpAddress and host the port if input contain ":" prefix.
func IpAddress(ip string) error {
	if strings.Contains(ip, ":") {
		parts := strings.Split(ip, ":")
		ip = parts[0]
		port := parts[1]

		err := Port(port)
		if err != nil {
			return fmt.Errorf("%s", err.Error())
		}
	}
	if net.ParseIP(ip) == nil {
		return fmt.Errorf("invalid IP Address: %s", ip)
	}
	return nil
}

// Port checks the valid port
func Port(port string) error {
	portNum, err := strconv.Atoi(port)
	if err != nil {
		return fmt.Errorf("invalid port number: %s", port)
	}
	if portNum < 1 || portNum > 65535 {
		return fmt.Errorf("port number out of range: %d", portNum)
	}
	return nil
}

// ParseCidr parses valid cidr
func ParseCidr(cidr string) (net.IP, *net.IPNet, error) {
	ip, n, err := net.ParseCIDR(cidr)
	return ip, n, err
}
