package main

import (
	crand "crypto/rand" /* You import this package and name it as crand */
	"encoding/binary"
	"fmt"
	"math/rand"
)

func main() {
	r := seedRand()
	fmt.Println(r.Int())
}

func seedRand() *rand.Rand {
	var b [8]byte
	_, err := crand.Read(b[:])

	if err != nil {
		panic("Cannot seed with cryptographic random number generator")
	}
	r := rand.New(rand.NewSource(int64(binary.LittleEndian.Uint64(b[:]))))
	return r
}
