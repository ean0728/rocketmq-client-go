package utils

import "strings"

// 去除可能存在的http:// 或者https://
func MockAddrValid(addr string) string {
	if strings.HasPrefix(addr, "http://") {
		return strings.TrimPrefix(addr, "http://")
	}
	if strings.HasPrefix(addr, "https://") {
		return strings.TrimPrefix(addr, "https://")
	}

	return addr
}
