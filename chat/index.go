package chat

import (
	"flag"

	"github.com/jeckbjy/fairy/exit"
)

func Test() {
	// example: ./test -s server
	pside := flag.String("s", "server", "test mode:server or client")
	flag.Parse()

	if *pside == "server" {
		StartServer()
	} else {
		StartClient()
	}

	exit.Wait()
}
