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

	// reap.Register()
	reap.Register(ctx context.Context, ssId string, baseDir string, vmmStatePath string, guestMemPath string, memSize int, wsFileDirectIO, wsSingleRead bool) (string, error) 

	reap.Activate(req *http.Request, id string)

}
