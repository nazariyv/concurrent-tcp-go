# concurrent-tcp-go

Concurrent TCP server implemented in Go using this <a href="https://opensource.com/article/18/5/building-concurrent-tcp-server-go">resource</a>

This is a TCPv4/IP server that spawns goroutines for each of the clients. Send messages to the server by terminating the message with \n character.

To start the server run

`$go conTCP.go 8001`

You can use a GUI tool like `PacketSender` to create TCP clients.
Alternatively, you can create a TCP connection by running

`$nc localhost 8001`

To check the served TCP clients by this script run

`netstat -anp TCP | grep 8001`
