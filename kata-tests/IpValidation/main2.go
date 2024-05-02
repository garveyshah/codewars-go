package main

import "strconv"

func split(value string, separator string) []string {

	var parts []string
	last := 0

	for i, v := range value {
		if string(v) == separator {
			parts = append(parts, value[last:i])
			last = i + 1
		}
	}

	parts = append(parts, value[last:])

	return parts
}

func Is_valid_ip(ip string) bool {

	parts := split(ip, ".")

	if len(parts) != 4 {
		return false
	}

	for _, value := range parts {
		if value[0] == '0' && len(value) > 1 {
			return false
		}

		num, err := strconv.Atoi(value)

		if err != nil {
			return false
		}

		if num > 255 || num < 0 {
			return false
		}

	}

	return true
}
