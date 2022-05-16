package main

import (
	"flag"
	"fmt"

	"github.com/cmstack/go-portqry/src/port"
)

func main() {
	var protocol string
	var host string
	var portnumber int
	var timeout int

	flag.StringVar(&protocol, "protocol", "tcp", "Use the specified protocol, (possible values are tcp, udp, or both). The default value is tcp")
	flag.StringVar(&host, "name", " ", "The <name> value represents the name or IP address of the computer to query. This value cannot include spaces")
	flag.IntVar(&portnumber, "ports", 0, "The <port> value represents the port to query on the destination computer.")
	flag.IntVar(&timeout, "timeout", 5, "The <port> value represents the timeout duration")
	flag.Parse()

	if host == " " {
		fmt.Println("Please specify a hostname or IP address")
		return
	}

	if portnumber == 0 {
		fmt.Println("Please specify a port number")
		return
	}

	if protocol != "tcp" {
		fmt.Println("Only tcp protocol is supported for now, please specify tcp")
	} else {
		fmt.Println("Scanning TCP port", portnumber, "on", host)
	}

	fmt.Printf("The TCP PORT \"%v\" is: %v\n", portnumber, port.PortScannerTCP(host, portnumber, timeout))
}
