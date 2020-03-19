package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

func handleConnection(c net.Conn) {
	fmt.Printf("Serving %s\n", c.RemoteAddr())
	for {
		netData, err := bufio.NewReader(c).ReadString('\n')

		if err != nil {
			fmt.Printf("connection with %s terminated. reason %s\n", c.RemoteAddr(), err)
			// if I leave this in, then using `netstat -anp TCP | grep 8001`
			// in the terminal shows that the tcp client is in state `FIN_WAIT_2`
			// however when using break like below, it removes the tcp client from netstat
			return
		}

		temp := strings.TrimSpace(string(netData))

		if temp == "STOP" {
			break
		} else {
			fmt.Printf("%s sent %s\n", c.RemoteAddr(), temp)
		}

		result := strconv.Itoa(rand.Int()) + "\n"
		c.Write([]byte(string(result)))
	}
	c.Close()
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide a port number!")
		return
	}

	PORT := ":" + arguments[1]
	l, err := net.Listen("tcp4", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()
	rand.Seed(time.Now().Unix())

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go handleConnection(c)
	}
}
