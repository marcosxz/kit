package kit

import (
	"errors"
	"net"
	"strings"
	"time"
)

var LocalIpEmptyErr = errors.New("local ip empty")

func LocalIP() (string, error) {

	ads, err := net.InterfaceAddrs()

	if err != nil {
		return "", err
	}

	for _, addr := range ads {
		if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() && ipNet.IP.IsGlobalUnicast() {
			if ipNet.IP.To4() != nil {
				return ipNet.IP.String(), nil
			}
		}
	}

	return "", LocalIpEmptyErr
}

func LocalIPByTcpDial(address string, timeout time.Duration) (string, error) {
	if conn, err := net.DialTimeout("tcp", address, timeout); err != nil {
		return "", err
	} else {
		ip := strings.Split(conn.LocalAddr().String(), ":")[0]
		return ip, conn.Close()
	}
}
