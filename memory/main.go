package main

import (
	"flag"
	"log"
	"net"

	"github.com/metaputer/m16arch"
)

var (
	debug    = flag.Bool("debug", false, "debug mode")
	capacity = flag.Int("cap", 1024, "memory size (default 1024)")
	plug     = flag.String("plug", "127.0.0.1:23550", "plug address")
)

var mem []m16arch.Byte

func main() {
	flag.Parse()
	mem = make([]m16arch.Byte, *capacity)
	var conn net.Conn
	for {
		var err error
		conn, err = net.Dial("tcp", *plug)
		if err == nil {
			break
		}
	}
	log.Println("Connected!!")
	defer conn.Close()
	for {
		mouth := make([]byte, 8)
		out := make([]m16arch.Byte, 4)

		_, err := conn.Read(mouth)
		if err != nil {
			break
		}
		input := m16arch.From8Byte(mouth)
		switch input[0] {
		case 1: //Read
			out[0] = mem[input[1]]
			if *debug {
				log.Printf("Readed value from %x is %x", input[1], mem[input[1]])
			}
		case 2: //Write
			mem[input[1]] = input[2]
			if *debug {
				log.Printf("Writed %x value to %x", input[2], input[1])
			}
		}
		_, err = conn.Write(m16arch.To8Byte(out))
		if err != nil {
			break
		}
	}
}
