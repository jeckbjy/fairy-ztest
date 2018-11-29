package test

import (
	"time"

	"github.com/jeckbjy/fairy/log"
	"github.com/jeckbjy/fairy/timer"
)

func TestTimer() {
	tt := time.Now().UnixNano() / int64(time.Millisecond)
	timer.Start(10000, func(t *timer.Timer) {
		log.Debug("OnTimer out:%+v", time.Now().UnixNano()/int64(time.Millisecond)-tt)
	})

	time.Sleep(100000 * time.Millisecond)
}
