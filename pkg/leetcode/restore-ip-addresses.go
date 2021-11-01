package leetcode

import (
	"strconv"
	"strings"
)

/*
	93. Restore IP Addresses
*/

func restoreIPAddresses(s string) []string {
	rs := []string{}
	ip := []string{}
	if len(s) < 4 {
		return rs
	}

	for i := 0; i < len(s) && i < 3; i++ {
		if checkIPSep(s[0 : i+1]) {
			ip = append(ip, s[0:i+1])
			rs = dfsIPAddr(s[i+1:], rs, ip)
			ip = ip[:len(ip)-1]
		}
	}
	return rs
}

func checkIPSep(s string) bool {
	if s == "" || len(s) != 1 && s[0] == '0' {
		return false
	}
	n, err := strconv.Atoi(s)
	if err != nil || n < 0 || n > 255 {
		return false
	}
	return true
}

func dfsIPAddr(s string, rs []string, ip []string) []string {
	if len(ip) == 3 {
		if checkIPSep(s) {
			ip = append(ip, s)
			ipStr := strings.Join(ip, ".")
			rs = append(rs, ipStr)
			return rs
		}
		return rs
	}

	for i := 0; i < len(s) && i < 3; i++ {
		if checkIPSep(s[0 : i+1]) {
			ip = append(ip, s[0:i+1])
			rs = dfsIPAddr(s[i+1:], rs, ip)
			ip = ip[:len(ip)-1]
		}
	}
	return rs
}
