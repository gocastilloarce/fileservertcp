package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"

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
	var chanMap = make([]string, 200)
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
	}()
	action := make([]byte, 20)
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
		subscribe(id, channel, chanMap)
		print("subscribe to " + splits[1])
	}
	/* for {
	p := make([]byte, 2000000)
	len, err := bufio.NewReader(c).Read(p)
	fmt.Println("Longitud archivo:", len)
	if err != nil {
		println("error reading from client")
		fmt.Println(err)
		return
	}
	os.WriteFile("recibido.txt", p[0:len], 0644)
	/* _, err = c.Write(p[0:len])
	if err != nil {
		println("error reading from client")
		fmt.Println(err)
		return
	} */
	//println("mensaje:", string(p), id)
	/* connMap.Range(func(key, value interface{}) bool {
			conn, ok := value.(net.Conn)
			if ok {
				_, err := conn.Write(p[0:len])
				if err != nil {
					println("error on writing to connection", err)
				}
			}
			return true
		})
	} */
}

func subscribe(id string, channel int, chanMap *sync.Map) {
	chanMap.Store(id)
}
