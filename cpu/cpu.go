package main

import (
	"flag"
	"log"
	"net"

	"github.com/metaputer/m16arch"
)

var (
	plug = flag.String("plug", "127.0.0.1:23551", "plug address")
)

var (
	rgPtrN m16arch.Byte
	rgPtrW m16arch.Byte
	rgNumP m16arch.Byte
	rgNumN m16arch.Byte
	rgNumS [16]m16arch.Byte
)

var conn net.Conn

func main() {
	flag.Parse()
	for {
		var err error
		conn, err = net.Dial("tcp", *plug)
		if err == nil {
			break
		}
	}
	log.Println("Connected!!")
	defer conn.Close()
	for i := 0; i < 30; i++ {
		inst := readMem(rgPtrW)
		rgPtrW++
		if int(inst>>8) < len(ops) {
			ops[inst>>8](byte(inst & 0xff))
		}
	}
}

func readMem(addr m16arch.Byte) m16arch.Byte {
	conn.Write(m16arch.To8Byte([]m16arch.Byte{0x1, addr, 0x0, 0x0}))
	date := make([]byte, 8)
	conn.Read(date)
	return m16arch.From8Byte(date)[0]
}
