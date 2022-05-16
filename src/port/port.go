package port

import (
	"net"
	"strconv"
	"time"
)

func PortScannerTCP(host string, port, duration int) string {
	address := net.JoinHostPort(host, strconv.Itoa(port))

	conn, err := net.DialTimeout("tcp", address, time.Second*time.Duration(duration))
	if err != nil {
		return "CLOSED"
	}
	defer conn.Close()

	return "OPEN"
}
