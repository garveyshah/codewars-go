package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	ips := []string{"0.0.0.0", "127.1.1.0", "192.168.0.1", "255.255.255.0", "10.10.10.10", "256.0.0.1", "192.168.1", "123.456.78.90", "123.045.067.089", "01.02.03.04", "1.2.3.4.5", "123.45.67.89"}

	// Test the validity of each IP address
	for _, ip := range ips {
		fmt.Println(ip, ":", Is_Valid_Ip(ip))
	}
}

func Is_Valid_Ip(ip string) bool {
	octets := strings.Split(ip, ".")

	if len(octets) != 4 {
		return false
	}

	for _, octetStr := range octets {
		if len(octetStr) > 1 && octetStr[0] == '0' {
			return false
		}
		octet, err := strconv.Atoi(octetStr)
		if err != nil || octet < 0 || octet > 255 {
			return false
		}

	}

	return true
}
