package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"strings"
	"sync"
	"time"
	"unicode"
)

var allclientName []string = make([]string, 0)
var allclientNameMutex sync.Mutex

func handleClient(conn net.Conn) {

	if len(clients) < 10 {
		defer conn.Close()
		logoBytes, err := ioutil.ReadFile("welcome.txt")
		if err != nil {
			fmt.Println("Error reading ASCII logo file:", err)
			return
		}
		coloredLogo := "\x1b[34m" + string(logoBytes) + "\x1b[0m" // Blue color
		conn.Write([]byte(coloredLogo + "\n"))
		conn.Write([]byte("[ENTER YOUR NAME]: "))
		reader := bufio.NewReader(conn)
		clientName, err := reader.ReadString('\n')
		if err != nil {
			conn.Close()
			return
		}
		clientName = espace(clientName)
		for clientName == "" || len(clientName) == 0 || !isPrintable(clientName) && checkName(clientName, allclientName) {
			conn.Write([]byte("[ENTER YOUR NAME]:"))
			reader := bufio.NewReader(conn)
			clientName, err = reader.ReadString('\n')
			clientName = espace(clientName)
		}

		for clientName == "" || len(clientName) == 0 || !isPrintable(clientName) || checkName(clientName, allclientName) {
			conn.Write([]byte("[ENTER ANOTHER YOUR NAME THE NAME YOU ENTER IS ALREADY USED BY OTHER USER OR IS INVALID] \n"))
			conn.Write([]byte("[ENTER YOUR NAME]: "))
			reader = bufio.NewReader(conn)
			clientName, err = reader.ReadString('\n')
			clientName = espace(clientName)

		}
		if clientName != "" && isPrintable(clientName) {
			fileContent, err := ioutil.ReadFile("filename.txt")
			if err != nil {
				fmt.Println("Error reading ASCII logo file:", err)

				conn.Close()
				return
			}
			fileContents := "\x1b[31m" + string(fileContent) + "\x1b[0m" // Red color
			conn.Write([]byte(fileContents + "\n"))
		}
		if err != nil {
			fmt.Println("Error reading client name:", err)
			return
		}
		client := &Client{
			conn:        conn,
			name:        clientName,
			messages:    make(chan string),
			connectedAt: time.Now(),
			index:       len(clients),
		}
		var er error

		clients = append(clients, client)
		allclientNameMutex.Lock()
		allclientName = append(allclientName, clientName)
		allclientNameMutex.Unlock()

		logText := fmt.Sprintf("%s has joined our chat... at %s\n", clientName, formatDate(time.Now()))
		clientSendMessage(logText, conn)

		writeTolog(logText)
		go func(client22 *Client) {
			for message := range client.messages {
				_, err := client.conn.Write([]byte(message))
				if err != nil {
					break
				}

			}
			allclientName1 := allclientName
			supprimerName(allclientName1, client.name)
			allclientName = allclientName1

		}(client)
		for {
			logText11 := fmt.Sprintf("[%s][%s]:", formatDate(time.Now()), client.name)
			conn.Write([]byte(logText11))
			message, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Error reading message:", err)
				break
			}
			for err != nil {
				message, err = reader.ReadString('\n')
				er = err
			}
			message = strings.TrimSpace(message)
			if message == "exit" {
				logText = fmt.Sprintf("Client %s exited at: %s", client.name, formatDate(time.Now())) + "\n"
				writeTolog(logText)
				break
			}
			// conn.Write([]byte("Enter a message or press '--change' to change your name: \n"))

			if message == "--change" {
				olderName := client.name
				conn.Write([]byte("[ENTER YOUR NEW NAME]: "))
				msg, err := reader.ReadString('\n')
				for msg == "" || len(msg) == 0 || !isPrintable(msg) || checkName(msg, allclientName) {
					conn.Write([]byte("[ENTER ANOTHER YOUR NAME THE NAME YOU ENTER IS ALREADY USED BY OTHER USER OR IS INVALID]: "))
					reader = bufio.NewReader(conn)
					msg, err = reader.ReadString('\n')
					msg = espace(msg)

					//clientName = clientName[:len(clientName)-1]
				}
				if err != nil {
					fmt.Println("Error reading the name:", err)
					break
				}
				msg = espace(msg)
				client.name = msg
				clientName = msg
				logText = fmt.Sprintf("[%s] change his name to: [%s] at : [%s]", olderName, client.name, formatDate(time.Now()))
				clientSendMessage(logText, conn)
				writeTolog(logText)
			} else if message != "" && isPrintable(message) {
				// Handle the received message (you can send it to other clients, etc.)
				logText = fmt.Sprintf("[%s][%s]:%s", formatDate(time.Now()), client.name, message+"\n")
				if er != nil {
					clientSendMessage("", conn)
					break
				} else {
					clientSendMessage(logText, conn)
					writeTolog(logText)
				}
			}
		}
		logText22 := fmt.Sprintf(clientName + " has left our chat... at " + formatDate(time.Now()) + "\n")
		clientSendMessage(logText22, conn)
		fmt.Println(clientName, "has left our chat...")
		clients1 := supprimer(clients, client.index)
		clients = clients1
		allclientNameMutex.Lock()
		allclientName = supprimerName(allclientName, client.name)
		allclientNameMutex.Unlock()

		close(client.messages)
		conn.Close()

	} else {
		conn.Write([]byte("The limit of the chat is reached: "))
		conn.Close()
		return
	}
}
func writeTolog(s string) {
	file, err := os.OpenFile("filename.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
	}
	_, err = file.WriteString(s)
	if err != nil {
		fmt.Println("Error writing to file:", err)
	}
}
func formatDate(currentTime time.Time) string {
	layout := "2006-01-02 15:04:05"
	formattedDate := currentTime.Format(layout)
	return formattedDate
}

func clientSendMessage(message string, conn net.Conn) {
	for _, client := range clients {
		if conn != client.conn {
			msg := fmt.Sprintf("\n%s[%s][%s]:", message, formatDate(time.Now()), client.name)
			client.messages <- msg
		}
	}
}
func checkName(sname string, tab []string) bool {
	allclientNameMutex.Lock()
	//exists := checkName(clientName, allclientName)
	allclientNameMutex.Unlock()
	for i := 0; i < len(tab); i++ {
		if tab[i] == sname {
			return true
		}
	}
	return false
}
func supprimer(clients []*Client, i int) []*Client {
	clients[i] = clients[len(clients)-1]
	clients[i].index = i
	return clients[:len(clients)-1]
}

func supprimerName(slice []string, str string) []string {
	index := -1
	for i, v := range slice {
		if v == str {
			index = i
			break
		}
	}
	if index != -1 {
		return append(slice[:index], slice[index+1:]...)
	}
	return slice
}

func isPrintable(str string) bool {

	for _, char := range str {
		if !unicode.IsPrint(char) {
			return false
		}
	}
	return true
}

func espace(str string) string {
	str = strings.ReplaceAll(str, " ", "")
	if len(str) > 0 {
		str = strings.TrimSpace(str[:len(str)-1])
	}
	return str
}
