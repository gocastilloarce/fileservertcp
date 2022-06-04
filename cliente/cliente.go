package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"path"
	"strconv"
	"time"
)

func cliente() {
	fmt.Println(os.Args)
	switch os.Args[1] {
	case "receive":
		{
			channel, err := strconv.Atoi(os.Args[2])
			if err != nil {
				fmt.Println(err)
				return
			}
			receive(channel)
			break
		}
	case "send":
		{
			filename := os.Args[2]
			channel, err := strconv.Atoi(os.Args[3])
			if err != nil {
				fmt.Println(err)
				return
			}
			sendFile(filename, channel)
			break
		}
	}
	/* nombre := flag.String("f", "", "file name")
	flag.Parse()

	readFile, err := os.ReadFile(*nombre)
	if err != nil {
		log.Fatal(err)
	}

	c, err := net.Dial("tcp", ":9090")
	if err != nil {
		fmt.Print("error: ")
		fmt.Println(err)
		return
	}
	l, err := c.Write(readFile)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Print("escrito: ", l)
	if err != nil {
		fmt.Println("error:", err)
	}
	p := make([]byte, 2000000)
	longi, err := c.Read(p)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("longi", longi) */
	/* longi, err = c.Read(p)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("longi", longi) */
	/* c.Close()
	defer c.Close() */
	/* for {
		fmt.Print("Escriba algo: ")
		reader := bufio.NewReader(os.Stdin)
		msg, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("error:", err)
		}
		c.Write([]byte(fmt.Sprint(msg)))
	} */

}

func receive(channel int) {
	c, err := net.Dial("tcp", ":9090")
	defer c.Close()
	if err != nil {
		fmt.Println("error: ", err)
		return
	}
	_, err = c.Write([]byte("subscribe " + strconv.Itoa(channel)))
	if err != nil {
		fmt.Println("error:", err)
	}
	name := make([]byte, 200)
	size := make([]byte, 200)
	c.SetReadDeadline(time.Now().Add(10 * time.Minute))
	for {
		l, err := c.Read(name)
		if err != nil {
			fmt.Println("error reading name", err)
			return
		}
		name_readed := string(name[:l])
		l, err = c.Read(size)
		if err != nil {
			fmt.Println("error reading name", err)
			return
		}
		size_readed, _ := strconv.Atoi(string(size[:l]))
		fmt.Println("Por recibir: ", name_readed, " de "+strconv.Itoa(size_readed)+"Bytes")

		file_mem := make([]byte, 0, size_readed)
		buff := make([]byte, 100000)
		for {
			l, err = c.Read(buff)
			if err != nil {
				print(err.Error())
			}
			println(len(file_mem))
			file_mem = append(file_mem, buff[:l]...)
			println(len(file_mem), cap(file_mem))
			if len(file_mem) == size_readed {
				break
			}
		}
		err = os.WriteFile(path.Base(name_readed)[:len(path.Base(name_readed))-len(path.Ext(name_readed))]+"_"+time.Now().Format("2006_01_02_15_04_05_0700")+path.Ext(name_readed), file_mem, 0644)
		if err != nil {
			print(err.Error())
		}
	}
}

func sendFile(filename string, channel int) {
	c, err := net.Dial("tcp", ":9090")
	c.SetWriteDeadline(time.Now().Add(2 * time.Second))
	if err != nil {
		fmt.Print("error: ")
		fmt.Println(err)
		return
	}

	fi, err := os.Stat(filename)
	if err != nil {
		fmt.Print("error: ")
		fmt.Println(err)
		return
	}
	size := fi.Size()
	_, err = c.Write([]byte("send " + strconv.Itoa(channel) + " " + strconv.Itoa(int(size)) + " " + filename))
	if err != nil {
		fmt.Println("error:", err)
	}

	readFile, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	l, err := c.Write(readFile)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Print("escrito: ", l)
	reader := bufio.NewReader(os.Stdin)
	_, _ = reader.ReadString('\n')

}

func main() {
	cliente()
}
