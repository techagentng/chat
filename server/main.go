package main

import (
	"flag"
	"fmt"
	"log"
	"net"
)

var (
	port *int
)

// Basic flag declarations are available for string, integer, and boolean options.
func init() {
	port = flag.Int("port", 3000, "an int")
}

func main() {

	// Once all flags are declared, call flag.Parse() to execute the command-line parsing.
	flag.Parse()

	fmt.Println("port:", *port)
	listener, err := net.ListenTCP("tcp", &net.TCPAddr{Port: *port}) //Listen for conn and check error
	if err != nil{
		log.Fatalf("error creating server %v\n", err)
	}
	defer listener.Close()  //A var with tcp and addr params
	_, err = listener.Accept()
	if err != nil{
		log.Printf("error accepting connection %v", err)
	}
}