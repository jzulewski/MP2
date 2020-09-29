package main

import (
	"MP2/structs"
	"os"
)

func main() {

	arguments := os.Args

	switch arguments[1] {
	case "server":
		startServer(arguments[2])

	case "client":
		client := structs.Client{IP: arguments[3], PORT: arguments[4], Username: arguments[5]}
		conn := connectToServer(client)
		message := readMessageFromClient(client.Username)
		go sendMessageToServer(conn, message)
	default:
		print("Please specify if you are using a server or client")
	}

}
