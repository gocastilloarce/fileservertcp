package main

import (
	"log"
	//"flag"
	"fmt"
	"net"
	"os"
)

func cliente() {
	//nombre := flag.Int("nombre", 0, "El nombre de la persona")
	//flag.Parse()
	//fmt.Println("Nombre: ", *nombre)
	readFile, err := os.ReadFile("a.png")
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
	archivo, err := os.Open("a.png")
	fi, err := archivo.Stat()
	fmt.Println(fi.Size())
	if err != nil {
		fmt.Println("error:", err)
	}
	p := make([]byte, 2000000)
	longi, err := c.Read(p)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("longi", longi)
	p = p[0:longi]
	defer c.Close()
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

func main() {
	cliente()
}
