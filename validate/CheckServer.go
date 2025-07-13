package validate

import (
	"errors"
	"fmt"
	"net"
	"time"
)

// CheckServer validates IP/hostname and attempts to connect to ip:port
func CheckServer(ip, port string) error {
	if ip == "" || port == "" {
		return errors.New("IP or port cannot be empty")
	}

	address := net.JoinHostPort(ip, port)
	timeout := 3 * time.Second

	conn, err := net.DialTimeout("tcp", address, timeout)
	if err != nil {
		return fmt.Errorf("cannot connect to %s: %w", address, err)
	}
	defer conn.Close()

	return nil
}
