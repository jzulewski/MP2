package main

import (
	"MP2/structs"
	"fmt"
	"net"
)

func startServer(address string) {
	listener, err := net.Listen("tcp", ":"+address)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}

		go handleConnection(conn)

	}

}

func handleConnection(conn net.Conn) {

	connection := <-conns
	message := <-messages

	to := message.To

	if to == "" {
		print("Error: To field was not provided")
	}

	if message.Content == "EXIT" {
		connection.Close()
	} else {
		sendMessageToClient(to, message)
	}

	conn.Close()

}

func sendMessageToClient(to string, message structs.Message) {
	fmt.Printf("To: %s, From: %s, \n %s", to, message.From, message.Content)
}
