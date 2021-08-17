package httputils

import (
	"fmt"
	"net"
	"regexp"
	"strconv"
)

// GetLocalIP get local network ip
func GetLocalIP() ([]string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return nil, err
	}

	ips := make([]string, 0, len(addrs))
	for _, address := range addrs {
		if ipNet, ok := address.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				ips = append(ips, ipNet.IP.String())
			}
		}
	}

	return ips, nil
}

// IP2Long IP convert to long int
func IP2Long(ipStr string) (uint64, error) {
	var ip uint64
	r := `^(\d{1,3})\.(\d{1,3})\.(\d{1,3})\.(\d{1,3})`
	reg, err := regexp.Compile(r)
	if err != nil {
		return 0, err
	}
	ips := reg.FindStringSubmatch(ipStr)
	if ips == nil {
		return 0, fmt.Errorf("Error ip addr:" + ipStr)
	}

	ipInt := make([]int, 0, 4)
	for index, i := range ips {
		d, err := strconv.Atoi(i)
		if err != nil {
			return 0, nil
		}
		if d < 0 || d > 255 {
			return 0, fmt.Errorf("Error ip addr:%s in segment[%d]", ipStr, index)
		}
		ipInt = append(ipInt, d)
	}

	ip += uint64(ipInt[0] * 0x1000000)
	ip += uint64(ipInt[1] * 0x10000)
	ip += uint64(ipInt[2] * 0x100)
	ip += uint64(ipInt[3])

	return ip, nil
}

// Long2IP longint convert to IP
func Long2IP(ip uint32) string {
	return fmt.Sprintf("%d.%d.%d.%d", ip>>24, ip<<8>>24, ip<<16>>24, ip<<24>>24)
}
