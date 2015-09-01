package main

import "log"

//Op is Operation function
type Op func(byte)

var ops = []Op{
	opNope,
	opSwap,
	opSave,
	opLoad,
	opIncN,
	opIncS, //0x05
	opAddN,
	opAddS,
	opJmpP,
	opRdPt,
	opOutN, //0x0a
	opStop,
}

func opNope(n byte) {

}

func opSwap(n byte) {
	a, b := n&0xF>>4, n&0xF
	if a == b {
		rgNumN, rgNumS[a] = rgNumS[a], rgNumN
		return
	}

	rgNumS[b], rgNumS[a] = rgNumS[a], rgNumS[b]
}

func opSave(n byte) {
	rgNumS[n&0xF] = rgNumN
}

func opLoad(n byte) {
	rgNumN = rgNumS[n&0xF]
}

func opIncN(n byte) {
	rgNumN++
}

func opIncS(n byte) {
	rgNumS[n&0xF]++
}

func opAddN(n byte) {
	rgNumN += rgNumS[n&0xF]
}

func opAddS(n byte) {
	rgNumS[n&0xF] += rgNumN
}

func opJmpP(n byte) {
	rgPtrW = rgNumP
}

func opRdPt(n byte) {
	rgNumP = readMem(rgPtrW)
	rgPtrW++
}

func opOutN(n byte) {
	log.Printf("0x%04x, %v\n", rgNumN, rgNumN)
}

func opStop(n byte) {
	for {
	}
}
