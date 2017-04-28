package main

import (
	"bufio"
	//"encoding/binary"
	"flag"
	"fmt"
	"net"
	"os"
	//	"time"
)

var host = flag.String("host", "0.0.0.0", "host")
var port = flag.String("port", "9110", "port")

func main() {
	reader := bufio.NewReader(os.Stdin)
	input := ""
	for {
		if v, err := reader.ReadString('\n'); err == nil {
			input = input + v
		} else {
			break
		}
	}
	//input = input[:len(input)-1]
	//fmt.Printf("input is:%s", input)

	flag.Parse()

	conn, err := net.Dial("udp", *host+":"+*port)
	if err != nil {
		fmt.Println("Can't dial: ", err)
		os.Exit(1)
	}
	//fmt.Println("connected")
	defer conn.Close()
	_, err = conn.Write([]byte(input))
	if err != nil {
		fmt.Println("failed:", err)
		os.Exit(1)
	}
	os.Exit(0)
}
