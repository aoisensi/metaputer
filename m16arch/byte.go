package m16arch

import "math"

type Byte uint16

const MaxPointer = math.MaxUint16

func To8Byte(input []Byte) []byte {
	out := make([]byte, len(input)*2)
	for i, b := range input {
		out[i*2] = byte(b >> 8)
		out[i*2+1] = byte(b & 0xff)
	}
	return out
}

func From8Byte(input []byte) []Byte {
	out := make([]Byte, len(input)/2)
	for i := range out {
		out[i] = Byte(int(input[i*2])<<8 + int(input[i*2+1]))
	}
	return out
}
