package utils

import (
	"errors"
	"net"
)

func ExternalIP() ([]string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return []string{}, err
	}
	var ipAddresses = make([]string, 0, 64)
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		//		if iface.Flags&net.FlagLoopback != 0 {
		//			continue // loopback interface
		//		}
		addrs, err := iface.Addrs()
		if err != nil {
			return []string{}, err
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip == nil {
				continue
			}
			ip = ip.To4()
			if ip == nil {
				continue // not an ipv4 address
			}
			ipAddresses = append(ipAddresses, ip.String())
		}
	}
	if len(ipAddresses) != 0 {
		return ipAddresses, nil
	}
	return []string{}, errors.New("are you connected to the network?")
}
