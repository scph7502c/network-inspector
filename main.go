package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

type socket struct {
	iface string
	port  string
}

type InterfaceInfo struct {
	Name string
	MAC  string
	IPv4 []string
	IPv6 []string
}

func checkHost() {
	arguments := os.Args
	if len(arguments) != 3 {
		log.Fatalf("Usage: <ip> <port>", arguments[0])
	}
	s := socket{}
	s.iface = arguments[1]
	s.port = arguments[2]
	address := s.iface + ":" + s.port
	fmt.Println("\nHost is listening on interface:port --> ", address)
}

func collectInterfacesInfo() ([]InterfaceInfo, error) {
	var result []InterfaceInfo

	interfaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	for _, iface := range interfaces[1:] {
		info := InterfaceInfo{
			Name: iface.Name,
			MAC:  iface.HardwareAddr.String(),
		}

		addrs, err := iface.Addrs()
		if err != nil {
			return nil, err
		}

		info.IPv4, info.IPv6 = parseIPAddresses(addrs)
		result = append(result, info)
	}

	return result, nil
}

func parseIPAddresses(addrs []net.Addr) (ipv4 []string, ipv6 []string) {
	for _, addr := range addrs {
		ipNet, ok := addr.(*net.IPNet)
		if !ok {
			continue
		}

		ip := ipNet.IP
		if ip4 := ip.To4(); ip4 != nil {
			ipv4 = append(ipv4, ip4.String())
		} else {
			ipv6 = append(ipv6, ip.String())
		}
	}
	return
}

func printInterfaces(ifaces []InterfaceInfo) {
	for _, iface := range ifaces {
		ipv4 := "-"
		if len(iface.IPv4) > 0 {
			ipv4 = strings.Join(iface.IPv4, ", ")
		}

		ipv6 := "-"
		if len(iface.IPv6) > 0 {
			ipv6 = strings.Join(iface.IPv6, ", ")
		}

		fmt.Printf(
			"Name: %20s || MAC: %20s || IPv4: %20s || IPv6: %20s\n",
			iface.Name,
			iface.MAC,
			ipv4,
			ipv6,
		)
	}
}

func getInterfacesInfo() {
	ifaces, err := collectInterfacesInfo()
	if err != nil {
		log.Fatal(err)
	}

	printInterfaces(ifaces)
}

func main() {
	getInterfacesInfo()
}
