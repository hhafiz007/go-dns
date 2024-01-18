package main

import (
	// "bytes"
	// "encoding/binary"
	"fmt"

	// Uncomment this block to pass the first stage
	"net"
	"os"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	argumentOne := os.Args[0]

	if string(argumentOne) == "resolver" {
		fmt.Println("hi")

		os.Exit(2)
	}

	// Uncomment this block to pass the first stage
	//
	udpAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:2053")
	if err != nil {
		fmt.Println("Failed to resolve UDP address:", err)
		return
	}

	udpConn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		fmt.Println("Failed to bind to address:", err)
		return
	}
	defer udpConn.Close()

	buf := make([]byte, 512)

	for {
		size, source, err := udpConn.ReadFromUDP(buf)
		if err != nil {
			fmt.Println("Error receiving data:", err)
			break
		}

		receivedData := string(buf[:size])
		fmt.Printf("Received %d bytes from %s: %s\n", size, source, receivedData)

		// Create an empty response

		// header := NewDNSHeader()
		// header.Flags |= (1 << 15)
		// question := NewDNSQuestion()

		// qBytes := question.createQuestion()

		// answer := NewDNSAnswer()
		// aBytes := answer.createAnswer()

		// fmt.Println(qBytes)

		// response := header.createHeader()

		DNSMessage := NewDNSMessage()

		reply := DNSMessage.createMessage(buf)

		// fmt.Println(reply)

		// var reply []byte

		// reply = append(reply, response...)
		// reply = append(reply, qBytes...)
		// reply = append(reply, aBytes...)

		_, err = udpConn.WriteToUDP(reply, source)
		if err != nil {
			fmt.Println("Failed to send response:", err)
		}
	}
}
