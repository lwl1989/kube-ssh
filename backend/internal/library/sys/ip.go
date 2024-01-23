package sys

import (
	"errors"
	"github.com/lwl1989/kube-ssh/backend/internal/config"
	"net"
	"os"
	"strconv"
)

func GetLocalHostPort() (string, int) {
	ip, err := externalIP()
	if err != nil {
		return config.GlobalConfig.Host, 443
	}

	return ip.String(), config.GlobalConfig.Port
}

func externalIP() (net.IP, error) {
	iFaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	for _, iFace := range iFaces {
		if iFace.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iFace.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}
		address, err := iFace.Addrs()
		if err != nil {
			return nil, err
		}
		for _, addr := range address {
			ip := getIpFromAddr(addr)
			if ip == nil {
				continue
			}
			return ip, nil
		}
	}
	return nil, errors.New("connected to the network")
}

func getIpFromAddr(addr net.Addr) net.IP {
	var ip net.IP
	switch v := addr.(type) {
	case *net.IPNet:
		ip = v.IP
	case *net.IPAddr:
		ip = v.IP
	}
	if ip == nil || ip.IsLoopback() {
		return nil
	}
	ip = ip.To4()
	if ip == nil {
		return nil // not an ipv4 address
	}

	return ip
}

type Process struct {
	Name  string
	Pid   string
	State string
	Port  int64
}

func (p *Process) Kill() error {
	pid, _ := strconv.Atoi(p.Pid)
	proc, _ := os.FindProcess(pid)
	return proc.Kill()
}
