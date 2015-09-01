package main

import (
	"flag"
	"log"
	"net"

	"github.com/metaputer/m16arch"
)

var memplug = flag.String("memplug", ":23550", "memory plug addr")
var cpuplug = flag.String("cpuplug", ":23551", "cpu plug addr")
var bios = []m16arch.Byte{
	0x0900,
	0x0004,
	0x0400,
	0x0a00,
	0x0100,
	0x0600,
	0x0a00,
	0x0800,
}

var (
	memcon net.Conn
	cpucon net.Conn
)

func main() {
	memcon = connect(*memplug)
	cpucon = connect(*cpuplug)
	for a, v := range bios {
		date := []m16arch.Byte{0x2, m16arch.Byte(a), m16arch.Byte(v), 0}
		d := m16arch.To8Byte(date)
		memcon.Write(d)
		memcon.Read(d)
	}
	for {
		in := make([]byte, 8)
		cpucon.Read(in)
		input := m16arch.From8Byte(in)
		out := make([]m16arch.Byte, 4)
		switch input[0] {
		case 1:
			out[0] = readMem(input[1])
		}
		cpucon.Write(m16arch.To8Byte(out))
	}
}

func readMem(addr m16arch.Byte) m16arch.Byte {
	date := []m16arch.Byte{0x1, m16arch.Byte(addr), 0, 0}
	d := m16arch.To8Byte(date)
	memcon.Write(d)
	memcon.Read(d)
	out := m16arch.From8Byte(d)
	return out[0]
}

func writeMem(addr, value m16arch.Byte) {
	date := []m16arch.Byte{0x2, m16arch.Byte(addr), m16arch.Byte(value), 0}
	d := m16arch.To8Byte(date)
	memcon.Write(d)
	memcon.Read(d)
}

func connect(addr string) net.Conn {
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, _ := ln.Accept()
		if conn != nil {
			return conn
		}
	}
}
