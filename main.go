package main

import (
	"fmt"
	"log"
	"net"
	"os"
	//	"strconv"
	//	"text/tabwriter"
)

type socket struct {
	iface string
	port  string
}

type InterfaceInfo struct {
	Name string
	IPv4 []string
	Ipv6 []string
}

func getInterfacesInfo() {
	interfaces, err := net.Interfaces()
	if err != nil {
		log.Fatal(err)
	}
	//	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	// fmt.Fprintf(w, "Interface: %-10s\t \tMAC Address: %s\tIP address: \n", iface.Name, iface.HardwareAdd r.String())
	//w.Flush()

	interfaceInfoMap := []map[string]string{}

	for _, iface := range interfaces[1:] {
		m := make(map[string]string)
		m["Name"] = iface.Name
		m["MAC"] = iface.HardwareAddr.String()
		interfaceInfoMap = append(interfaceInfoMap, m)
		addrs, err := iface.Addrs()
		if err != nil {
			log.Fatal(err)
		}
		for _, addr := range addrs {
			ipNet, ok := addr.(*net.IPNet)
			if !ok {
				continue
			}

			ip := ipNet.IP
			if ip.To4() != nil {
				m["IPv4"] = ip.String()

			} else {
				m["IPv6"] = ip.String()
			}

		}

	}
	for _, iface := range interfaceInfoMap {
		fmt.Printf(
			"Name: %10s MAC: %10s IPv4: %s IPv6: %s\n",
			iface["Name"],
			iface["MAC"],
			iface["IPv4"],
			iface["IPv6"],
		)
	}
}

func checkHost() {
	//	defer c.Close()
	getInterfacesInfo()
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

func main() {
	checkHost()
}
