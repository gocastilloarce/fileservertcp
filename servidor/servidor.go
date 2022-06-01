package main

import (
	"bufio"
	"fmt"
	"net"
	"sync"

	"github.com/google/uuid"
)

func main() {
	l, err := net.Listen("tcp", "localhost:9090")
	if err != nil {
		return
	}
	defer l.Close()

	var connMap = &sync.Map{}

	for {
		conn, err := l.Accept()
		if err != nil {
			println("error accepting connection", err)
			return
		}

		id := uuid.New().String()
		connMap.Store(id, conn)
		go handleUserConnection(id, conn, connMap)
	}

}

func handleUserConnection(id string, c net.Conn, connMap *sync.Map) {
	defer func() {
		c.Close()
		connMap.Delete(id)
	}()
	for {
		p := make([]byte, 2000000)
		len, err := bufio.NewReader(c).Read(p)
		fmt.Println("Longitud archivo:", len)
		if err != nil {
			println("error reading from client")
			fmt.Println(err)
			return
		}
		_, err = c.Write(p[0:len])
		if err != nil {
			println("error reading from client")
			fmt.Println(err)
			return
		}
		/* println("mensaje:", string(p), id)
		connMap.Range(func(key, value interface{}) bool {
			if conn, ok := value.(net.Conn); ok {
				_, err := conn.Write([]byte(id[0:10] + ": " + string(p)))
				if err != nil {
					println("error on writing to connection", err)
				}
			}
			return true
		}) */
	}
}
