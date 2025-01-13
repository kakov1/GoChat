package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"sync"
)

var (
	PORT = "8000"
)

var (
	clients     = make(map[net.Conn]string)
	clientsLock sync.Mutex             
	messages    = make(chan string)        
)

func main() {
	listener, err := net.Listen("tcp", ":" + PORT)

	if err != nil {
		fmt.Println("Ошибка запуска сервера:", err)
		os.Exit(1)
	}

	defer listener.Close()
	fmt.Println("Сервер запущен на порту " + PORT + "...")

	go handleMessages()

	for {
		conn, err := listener.Accept()

		if err != nil {
			fmt.Println("Ошибка подключения клиента:", err)
			continue
		}

		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	conn.Write([]byte("Введите своё имя: "))
	name, _ := bufio.NewReader(conn).ReadString('\n')
	name = strings.TrimSpace(name)

	clientsLock.Lock()
	clients[conn] = name
	clientsLock.Unlock()

	messages <- fmt.Sprintf("%s присоединился к чату!", name)

	for {
		message, err := bufio.NewReader(conn).ReadString('\n')

		if err != nil || strings.TrimSpace(message) == "/exit" {
			clientsLock.Lock()
			delete(clients, conn)
			clientsLock.Unlock()
			messages <- fmt.Sprintf("%s покинул чат.", name)

			return
		}

		messages <- fmt.Sprintf("%s: %s", name, strings.TrimSpace(message))
	}
}

func handleMessages() {
	for message := range messages {
		clientsLock.Lock()

		for conn := range clients {
			conn.Write([]byte(message + "\n"))
		}

		clientsLock.Unlock()
	}
}
