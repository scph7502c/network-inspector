package main

import (
	"fmt"
	"log"
	"net"

	"github.com/vishvananda/netlink"
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

func fetchRoutes() ([]netlink.Route, error) {
	var allRoutes []netlink.Route

	routesV4, errV4 := netlink.RouteList(nil, netlink.FAMILY_V4)
	if errV4 == nil {
		allRoutes = append(allRoutes, routesV4...)
	}
	routesV6, errV6 := netlink.RouteList(nil, netlink.FAMILY_V6)
	if errV6 == nil {
		allRoutes = append(allRoutes, routesV6...)
	}

	if errV4 != nil && errV6 != nil {
		return nil, fmt.Errorf("cannot fetch IPv4 nor IPv6 routes")
	}
	return allRoutes, nil

}

func groupRoutesByLink(routes []netlink.Route) (map[int][]netlink.Route, error) {

	routesByLink := make(map[int][]netlink.Route)
	for _, route := range routes {
		routesByLink[route.LinkIndex] = append(routesByLink[route.LinkIndex], route)
	}

	return routesByLink, nil

}

func listInterfacesByNetlink() {
	links, err := netlink.LinkList()
	if err != nil {
		fmt.Printf("Can't fetch interface list: %v", err)
	}
	for _, link := range links {
		if err != nil {
			fmt.Printf("Error reading route table: %v", err)
		}
		attrs := link.Attrs()
		state := "DOWN"
		if attrs.Flags&net.FlagUp != 0 {
			state = "UP"
		}

		fmt.Printf("\n- Name: %-10s | Index: %-3d | State: %-5s | MTU: %d | MAC: %s\n",
			attrs.Name,
			attrs.Index,
			state,
			attrs.MTU,
			attrs.HardwareAddr.String(),
		)
		if routes, ok := routesByLink[attrs.Index]; ok {
			for _, r := range routes {
				fmt.Printf("    Route: dst=%v gw=%v src=%v\n", r.Dst, r.Gw, r.Src)
			}
		} else {
			fmt.Println("    (no routes)")
		}
	}
}

func main() {
	listInterfacesByNetlink()
}
