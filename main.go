package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
)

type Client struct {
	conn        net.Conn
	name        string
	messages    chan string
	connectedAt time.Time
	index       int
}

var (
	clients []*Client
)

func main() {
	var port string
	l := len(os.Args)
	if l == 1 || l == 2 {
		if l == 1 {
			port = "8989"
		} else {
			port = os.Args[1]
		}
		_, err := strconv.Atoi(port)
		if err != nil {
			fmt.Println("[USAGE]: ./TCPChat $port")
			os.Exit(0)
		}
		address := ":" + port
		listener, err := net.Listen("tcp", address)
		if err != nil {
			fmt.Println("Error listening:", err)
			return
		}
		videFichier()
		defer listener.Close()
		fmt.Printf("Server listening on localhost:%s\n", port)
		for {
			conn, err := listener.Accept()
			if err != nil {
				fmt.Println("Error accepting:", err)
				continue
			}

			go handleClient(conn)
		}
	} else {
		fmt.Println("[USAGE]: ./TCPChat $port")
	}
}
func videFichier() {
	// Write the string to a file
	filepath := "filename.txt"
	file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
	defer file.Close()
}
