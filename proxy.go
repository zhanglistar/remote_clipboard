package main

import (
	"fmt"
	"net"
	"os"
	"reflect"
	"strconv"
	"strings"
)

var ips []net.UDPAddr

func heartBeat() {

}

func handleSub(conn *net.UDPConn) {
	data := make([]byte, 1024)
	for {
		n, _, err := conn.ReadFromUDP(data)
		if err == nil {
			//fmt.Println(peerAddr)
			find := false
			s := strings.Split(string(data[:n]), ":")
			if len(s) != 2 {
				fmt.Println("format error")
				continue
			}
			port, err := strconv.Atoi(s[1])
			if err != nil {
				fmt.Println("port format error", err)
				continue
			}
			v := net.UDPAddr{IP: net.ParseIP(s[0]), Port: port}
			for _, item := range ips {
				if reflect.DeepEqual(item, v) {
					find = true
					break
				}
			}
			if find == false {
				ips = append(ips, v)
				fmt.Println("add ", v)
			}
		}
	}
}

func handleData(data []byte, addr *net.UDPAddr) {
	//fmt.Println(string(data), len(data))
	for _, ip := range ips {
		//fmt.Println(ip)
		conn, err := net.DialUDP("udp", nil, &ip)
		if err != nil {
			continue
		}
		//fmt.Println("write to aaa")
		_, err = conn.Write(data)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func main() {
	ln, err := net.ListenUDP("udp", &net.UDPAddr{IP: net.ParseIP("0.0.0.0"), Port: 9110})
	if err != nil {
		fmt.Println("Can't listen udp: ", err)
		os.Exit(1)
	}
	defer ln.Close()

	subConn, err1 := net.ListenUDP("udp", &net.UDPAddr{IP: net.ParseIP("0.0.0.0"), Port: 9111})
	if err1 != nil {
		fmt.Println("Can't listen udp: ", err1)
		os.Exit(1)
	}
	defer subConn.Close()

	go heartBeat()
	go handleSub(subConn)

	data := make([]byte, 1024)
	for {
		n, remoteAddr, err := ln.ReadFromUDP(data)
		if err != nil {
			fmt.Println("err")
			continue
		} else {
			//fmt.Println(remoteAddr)
			go handleData(data[:n], remoteAddr)
		}
	}
	os.Exit(0)
}
