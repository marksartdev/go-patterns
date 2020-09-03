package main

import (
	"flag"
	"strings"
)

func main() {
	isServer := flag.Bool("s", false, "Server-mode")
	port := flag.String("p", "8000", "Port fot starting server (in server-mode) or ports which client has to listen")
	location := flag.String("l", "", "Location of GumballMachine")
	count := flag.Int("c", 0, "Count of gumballs")

	flag.Parse()

	if *isServer {
		startServer(*port, *location, *count)
	} else {
		ports := strings.Split(*port, ",")
		startClient(ports)
	}
}
