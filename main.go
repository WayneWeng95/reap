package main

import (
	"math/rand"

	"github.com/ucsdsysnet/faasnap/reap"
)

func main() {
	for i := 0; i < 5; i++ {
		println(rand.Intn(10))
	}

	reap.Setup()

}
