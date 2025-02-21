package utils

import (
	"net"
	"sort"

	"github.com/c-robinson/iplib"
)

func NextIP(defaultIP string, ips []string) string {
	if len(ips) == 0 {
		return defaultIP
	}

	r := []net.IP{}
	for _, i := range ips {
		ip := net.ParseIP(i)
		r = append(r, ip)
	}

	sort.Sort(iplib.ByIP(r))

	last := r[len(r)-1]

	return iplib.NextIP(last).String()
}
