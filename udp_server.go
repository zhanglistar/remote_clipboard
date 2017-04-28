package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"os/exec"
)

var host = flag.String("host", "10.100.30.66", "host")
var port = flag.String("port", "9111", "port")
var local = ":9112"

func getLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

func main() {
	flag.Parse()
	ln, err := net.Dial("udp", "10.100.30.66:9111")
	if err != nil {
		fmt.Println("Can't resolve address: ", err)
		os.Exit(1)
	}
	defer ln.Close()
	n, err := ln.Write([]byte(getLocalIP() + local))
	if err != nil {
		fmt.Println(err, n)
		os.Exit(1)
	}
	dataln, err := net.ListenUDP("udp", &net.UDPAddr{IP: net.IPv4(0, 0, 0, 0), Port: 9112})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer dataln.Close()
	data := make([]byte, 1024)
	for {
		n, _, err := dataln.ReadFromUDP(data)
		//fmt.Println(remoteAddr)
		if err != nil {
			fmt.Println("err")
		} else {
			//fmt.Printf("remote , data %s", string(data))
			cmd := "echo -e \"" + string(data[:n]) + "\" |pbcopy"
			//fmt.Println(cmd)
			exec.Command("/bin/bash", "-c", cmd).Run()

		}
	}
}
