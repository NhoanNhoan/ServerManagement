package entity

import (
	"math"
	"net"
	"strconv"
)

type NetworkPortion struct {
	Id, Value string
	Netmask int
}

func (portion NetworkPortion) GenerateHosts()([]IpAddress, error) {
	cidr := portion.Value + "/" + strconv.Itoa(portion.Netmask)
	ip, ipnet, err := net.ParseCIDR(cidr)
	if err != nil {
		return nil, err
	}

	var ips []string
	for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); increase(ip) {
		ips = append(ips, ip.String())
	}

	// remove network address and broadcast address
	lenIPs := len(ips)

	listIp := make([]IpAddress, lenIPs)
	for i := range listIp {
		var ip IpAddress
		ip.State = "available"
		ip.NetworkPortion = portion
		if ip.Parse(ips[i] + "/" +strconv.Itoa(portion.Netmask)) {
			listIp[i] = ip
		}
	}

	return listIp, nil
}

func increase(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}

func (portion NetworkPortion) CalculateNumHosts() int {
	return int(math.Pow(2, float64(32 - portion.Netmask)))
}

