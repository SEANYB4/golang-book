package main

import (
	"encoding/gob"
	"fmt"
	"net"
)

func server() {

	// listen on a port
	ln, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Server listening...")

	for {

		// accept a connection
		c, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		// handle the connection
		go handleServerConnection(c)
	}

}


func handleServerConnection(c net.Conn) {

	// receive the message 
	var msg string 
	err := gob.NewDecoder(c).Decode(&msg)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Received: ", msg)
	}
	c.Close()
}


func client() {
	// connect to the server
	c, err := net.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println(err)
		return
	}


	// send the message

	msg := "Hello world"
	fmt.Println("Sending", msg)
	var i int = 1
	for i < 5 {
		err = gob.NewEncoder(c).Encode(msg)
		if err != nil {
			fmt.Println(err)
		}
	}
	
	c.Close()
}

func main() {
	go server()
	go client()

	var input string
	fmt.Scanln(&input)
}





