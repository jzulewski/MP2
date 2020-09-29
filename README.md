**MP2**

**To Run**

To start the server run:

`go run ./ server 1234`

To start a client run in a separate terminal:

`go run ./ client connect 127.0.0.1 1234 John`

To start other clients:

`go run ./ client connect IP 1234 Username`
**Currently the Client is in an infinite loop**

To send a message:

`go run ./ client send Username Message`

**Structure and Design**

To create the chat room, this design uses one main server that multiple clients can connect to. Having this central server makes the travel of messages simple to understand (since it is from one client to a common server to the destination client)
and makes it possible for one server to send an exit signal to all clients gracefully

main.go uses a switch case to separate commands in a straightforward way. To create or communicate with the server the first word in the command is server, and it is the same for the client

After creating the server and connecting the client, sending a message begins with reading the user input. From there the input is parsed into To, From, and Content fields and placed into a struct that can be sent to the server.

Sending a message to the server uses a channel from the client to the server. Once recieved from the channel, the message is displayed by the intended client

If the message is EXIT then the connection will be closed

**Improvements**
Fix Client Infinite Loop
Work on Exiting gracefully

**Citations**
I collaborated with Colin earlier in the week on some basic-level structure for the program flow
