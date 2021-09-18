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
// STEP4 DEP ON (27)
var connections = []net.Conn{} //Connection variable as a map
func main() {
	// STEP1 DEP ON (15)
	// Once all flags are declared, call flag.Parse() to execute the command-line parsing.
	flag.Parse()

	fmt.Println("port:", *port)
	// STEP2 FOLLOWED BY (32)(38)
	listener, err := net.ListenTCP("tcp", &net.TCPAddr{Port: *port}) //Server Listener + error check
	if err != nil{
		log.Fatalf("error creating server %v\n", err)
	}
	// STEP3 DEP ON (27)
	conn, err := listener.Accept() //connection is waiting
	connections = append(connections, conn)
	// STEP5 DEP ON (27)
	log.Printf("Accepted connection from %v \n", conn.RemoteAddr()) //Log address of netCat localhost
	if err != nil{
		log.Printf("error accepting connection %v", err)
	}
	defer listener.Close()  // tcp with addr + net.TCPAddr{Port: *port}
}

func readFromConnection(conn net.Conn){
	message := make([]byte, 200)
	n, err := conn.Read(message)
	if err != nil {
		log.Println("Error reading from Connection", err)
		return
	}
	log.Printf("recieved %v bytes from connection", n)
}