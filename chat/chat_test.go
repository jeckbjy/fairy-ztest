package chat

import (
	"testing"

	"github.com/jeckbjy/fairy/util"
)

func TestChat(t *testing.T) {
	StartServer()
	StartClient()
	util.Sleep(20 * 1000)
}
