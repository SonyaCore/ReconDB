package parser

import "net"

func ParseCidr(cidr string) (net.IP, *net.IPNet, error) {
	ip, n, err := net.ParseCIDR(cidr)
	return ip, n, err
}
