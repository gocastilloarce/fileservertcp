package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
)

func main() {
	if len(os.Args) > 0 {
		switch os.Args[1] {
		case "startserver":
			{
				startServer()
				break
			}
		}
	}
}

func startServer() {
	l, err := net.Listen("tcp", "localhost:9090")
	if err != nil {
		return
	}
	defer l.Close()
	var connMap = &sync.Map{}
	var chanMap = &sync.Map{}
	chanMap.Store(1, &sync.Map{})
	chanMap.Store(2, &sync.Map{})
	for {
		conn, err := l.Accept()
		if err != nil {
			println("error accepting connection", err)
			return
		}
		id := uuid.New().String()
		connMap.Store(id, conn)
		go handleUserConnection(id, conn, connMap, chanMap)
	}
}

func handleUserConnection(id string, c net.Conn, connMap *sync.Map, chanMap *sync.Map) {
	defer func() {
		c.Close()
		connMap.Delete(id)
		chanMap.Range(func(key, value interface{}) bool {
			clients_map, ok := value.(*sync.Map)
			if ok {
				clients_map.Delete(id)
			}
			return true
		})
	}()
	action := make([]byte, 200)
	l, err := bufio.NewReader(c).Read(action)
	if err != nil {
		fmt.Println("error reading mode", err)
		return
	}
	action_readed := string(action[:l])
	splits := strings.Split(action_readed, " ")
	if splits[0] == "subscribe" {
		channel, err := strconv.Atoi(splits[1])
		if err != nil {
			fmt.Println(err)
		}
		subscribe(id, channel, chanMap, c)
		print("subscribe " + id + " to " + splits[1])
	} else if splits[0] == "send" {
		channel_number, err := strconv.Atoi(splits[1])
		if err != nil {
			fmt.Println(err)
		}
		channel_Load, _ := chanMap.Load(channel_number)
		channel := channel_Load.(*sync.Map)
		size, err := strconv.Atoi(splits[2])
		if err != nil {
			fmt.Println(err)
		}
		filename := splits[3]
		fmt.Println("Sending " + filename + " to channel " + strconv.Itoa(channel_number))
		sendFile(id, channel, filename, size, c)
		fmt.Println("File " + filename + " sent to " + strconv.Itoa(channel_number))
	}
	time.Sleep(10 * time.Minute)
}

func subscribe(id string, channel int, chanMap *sync.Map, conn net.Conn) {
	chane, _ := chanMap.Load(channel)
	clients, ok := chane.(*sync.Map)
	if ok {
		clients.Store(id, conn)
	}

}

func sendFile(id string, channel *sync.Map, filename string, size int, c net.Conn) {
	writeString(filename, channel)
	writeString(strconv.Itoa(size), channel)
	buff := make([]byte, 100000)
	for {
		l, err := c.Read(buff)
		if err != nil {
			println(err.Error())
		}
		writeBytes(buff[:l], channel)
		if l < 1 {
			break
		}
	}
}

func writeString(message string, channel *sync.Map) {
	channel.Range(func(key, value interface{}) bool {
		conn, ok := value.(net.Conn)
		if ok {
			_, err := conn.Write([]byte(message))
			if err != nil {
				println("error on writing to connection", err.Error())
			}
		}
		return true
	})
}

func writeBytes(message []byte, channel *sync.Map) {
	channel.Range(func(key, value interface{}) bool {
		conn, ok := value.(net.Conn)
		if ok {
			_, err := conn.Write(message)
			if err != nil {
				println("error on writing to connection", err.Error())
			}
		}
		return true
	})
}
