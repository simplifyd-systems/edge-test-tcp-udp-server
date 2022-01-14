package main

import (
	"bufio"
	"edge-test-tcp-udp-server/env"
	"fmt"
	"log"
	"math/rand"
	"net"
	"strconv"
	"strings"
	"time"
)

func main() {
	ch := make(chan struct{})
	env.Setup()

	go tcpServer()
	go udpServer()

	<-ch
}

func tcpServer() {
	/*
		arguments := os.Args
		if len(arguments) == 1 {
			fmt.Println("Please provide port number")
			return
		}
		PORT := ":" + arguments[1]
	*/

	PORT := ":" + env.Var.Port
	l, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Fatalln(err)
		return
	}
	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		log.Println("New connection", c.RemoteAddr())
		if err != nil {
			log.Println("error accepting connection", err)
			continue
		}

		go func(conn net.Conn) {
			for {
				netData, err := bufio.NewReader(conn).ReadString('\n')
				if err != nil {
					fmt.Println(err)
					return
				}
				if strings.TrimSpace(string(netData)) == "STOP" {
					fmt.Println("Exiting TCP server!")
					return
				}

				fmt.Print("-> ", string(netData))
				t := time.Now()
				myTime := t.Format(time.RFC3339)

				netData = strings.TrimSuffix(netData, "\n")
				// c.Write([]byte(myTime))
				c.Write([]byte(fmt.Sprintf("--> '%s' - received at time: %s - from IP address %s\n\n", netData, myTime, c.RemoteAddr())))
			}
		}(c)
	}
}

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func udpServer() {
	/*
		arguments := os.Args
		if len(arguments) == 1 {
			fmt.Println("Please provide a port number!")
			return
		}
		PORT := ":" + arguments[1]
	*/
	PORT := ":" + env.Var.Port

	s, err := net.ResolveUDPAddr("udp4", PORT)
	if err != nil {
		log.Fatalln(err)
		return
	}

	connection, err := net.ListenUDP("udp4", s)
	if err != nil {
		log.Fatalln(err)
		return
	}

	defer connection.Close()
	buffer := make([]byte, 1024)
	rand.Seed(time.Now().Unix())

	for {
		n, addr, err := connection.ReadFromUDP(buffer)
		fmt.Print("-> ", string(buffer[0:n-1]))

		if strings.TrimSpace(string(buffer[0:n])) == "STOP" {
			fmt.Println("Exiting UDP server!")
			return
		}

		data := []byte(strconv.Itoa(random(1, 1001)))
		fmt.Printf("data: %s\n", string(data))
		_, err = connection.WriteToUDP(data, addr)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
