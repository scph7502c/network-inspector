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

func getInterfacesInfo() {
	interfaces, err := net.Interfaces()
	//	addresses, err := net.InterfaceAddrs()
	if err != nil {
		log.Fatal(err)
	}
	//	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	var interfaceInfoMap = make(map[string]string)

	for _, iface := range interfaces[1:] {
		//		addr, err :=iface.Addrs()
		//		if err != nil {

		//		}
		interfaceInfoMap["Interface name"] = iface.Name
		interfaceInfoMap["Interface MAC address"] = iface.HardwareAddr.String()
		interfaceInfoMap["Interface IP address"] = ""

		//fmt.Fprintf(w, "Interface: %-10s\t \tMAC Address: %s\tIP address: \n", iface.Name, iface.HardwareAddr.String())
		//w.Flush()

	}
	fmt.Println(interfaceInfoMap)

}

func main() {
	checkHost()
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
