package main

import (
	"MP2/structs"
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

var conns = make(chan net.Conn)
var messages = make(chan structs.Message)

func connectToServer(client structs.Client) net.Conn {
	address := client.IP + ":" + client.PORT
	conn, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Println(err)

	}
	return conn
}

func readMessageFromClient(user string) structs.Message {
	reader, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	words := strings.Split(reader, " ")
	message := structs.Message{
		To:      words[1],
		From:    user,
		Content: strings.Join(words[2:], " "),
	}

	return message

}

func sendMessageToServer(conn net.Conn, message structs.Message) {
	conns <- conn
	messages <- message
}
