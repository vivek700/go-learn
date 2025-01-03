package main

import (
	"fmt"
	"strings"
)

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

type IPAddr [4]byte

func (ip IPAddr) String() string {
	parts := make([]string, len(ip))
	for i, part := range ip {
		parts[i] = fmt.Sprintf("%d", part)
	}
	return strings.Join(parts, ".")

}

func main() {
	a := Person{"vivek", 23}
	b := Person{"Zaphod", 9001}

	fmt.Println(a, b)

	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}

}
