package httputil

import (
	"net"
	"strconv"
)

func GetPort(addr net.Addr) (int, error) {
	_, port, err := net.SplitHostPort(addr.String())
	if err != nil {
		return -1, err
	}
	parsedPort, err := strconv.Atoi(port)
	if err != nil {
		return -1, err
	}
	return parsedPort, nil
}
