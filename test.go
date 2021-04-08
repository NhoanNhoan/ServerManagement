package main

import (
	"fmt"
	"net"
	"strings"
	"errors"
	"strconv"
)

func main() {
	content, err := StadardlizeNetworkPortion("192.169.2.")
	fmt.Println (content, err)
}

func GenerateHosts(value string)([]string, int, error) {
	ip, ipnet, err := net.ParseCIDR(value)
	if err != nil {
		return nil, 0, err
	}

	var ips []string
	for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); inc(ip) {
		ips = append(ips, ip.String())
	}

	// remove network address and broadcast address
	lenIPs := len(ips)
	switch {
	case lenIPs < 2:
		return ips, lenIPs, nil

	default:
		return ips[1 : len(ips)-1], lenIPs - 2, nil
	}
}

func inc(ip net.IP) {
    for j := len(ip) - 1; j >= 0; j-- {
        ip[j]++
     	if ip[j] > 0 {
            break
        }
    }
}

func StadardlizeNetworkPortion(raw string) (string, error) {
octets := strings.Split(raw, ".")
	if len(octets) > 4 {
		return strings.Join(octets[:5], "."), nil
	}

	// Check whether octets is number
	for i := range octets {
		if "" == octets[i] {
			octets[i] = "0"
		}
		
		if num, err := strconv.Atoi(octets[i]); nil != err || num < 0 || num > 255 {
			return "", errors.New(raw + " is not like format of ip address")
		}
	}

	// Insert 0 if octet is empty
	for len(octets) < 4 {
		octets = append(octets, "0")
	}

	return strings.Join(octets, "."), nil	
}
