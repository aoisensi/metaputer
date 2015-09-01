package main

import (
	"log"
	"math/rand"
	"net"

	"github.com/metaputer/m16arch"
)

var tests = make([]m16arch.Byte, 256)

func main() {
	for i := range tests {
		tests[i] = m16arch.Byte(rand.Intn(0xffff))
	}
	ln, err := net.Listen("tcp", ":23550")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			continue
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()
	mouth := make([]m16arch.Byte, 4)
	hip := make([]byte, 8)
	mouth[0] = 2
	for a, d := range tests {
		mouth[1] = m16arch.Byte(a)
		mouth[2] = m16arch.Byte(d)
		conn.Write(m16arch.To8Byte(mouth))
		conn.Read(hip)
	}
	s := 256
	mouth[0] = 1
	mouth[2] = 0
	for a, d := range tests {
		mouth[1] = m16arch.Byte(a)
		conn.Write(m16arch.To8Byte(mouth))
		conn.Read(hip)
		out := m16arch.From8Byte(hip)
		if out[0] == d {
			s--
		}
	}
	log.Printf("Failed test is %v times.\n", s)
}
