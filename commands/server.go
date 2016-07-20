package commands

import (
	"log"
	"net"
	"strconv"
	"strings"
)

func handleConn(c net.Conn) {
	buf := make([]byte, 1024)
	_, err := c.Read(buf)
	if err != nil {
		log.Println("Error reading:", err.Error())
	}
	split := strings.Split(string(buf), ",")
	interval, parseErr := strconv.Atoi(split[0])
	if parseErr != nil {
		log.Println("Error parsing interval:", parseErr.Error())
	}
	go WatchCmd(uint64(interval), split[1], split[2])
	c.Write([]byte("We're watching the game for you and we'll notify you when it's ready to watch!"))
}

// ServerCmd does things
func ServerCmd() {
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}
