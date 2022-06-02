package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
)

func cliente() {
	fmt.Println(os.Args)
	if len(os.Args) == 3 {
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
	if err != nil {
		fmt.Print("error: ")
		fmt.Println(err)
		return
	}
	_, err = c.Write([]byte("subscribe " + strconv.Itoa(channel)))
	if err != nil {
		fmt.Println("error:", err)
	}
}

func main() {
	cliente()
}
