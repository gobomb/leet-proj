package net

import (
	"log"
	"net"
)

func ParceCidr() {
	_, c1, err := net.ParseCIDR("10.20.128.0")
	if err != nil {
		log.Fatalf("can't parse pod cidr:%s", err)
	}

	_, c2, err := net.ParseCIDR("10.46.0.0/12")
	if err != nil {
		log.Fatalf("can't parse pod cidr:%s", err)
	}

	log.Println(intersect(c1, c2))
}

func intersect(n1, n2 *net.IPNet) bool {
	return n2.Contains(n1.IP) || n1.Contains(n2.IP)
}
