package chat

import (
	"github.com/jeckbjy/fairy"
	"github.com/jeckbjy/fairy/codecs"
	"github.com/jeckbjy/fairy/filters"
	"github.com/jeckbjy/fairy/frames"
	"github.com/jeckbjy/fairy/identities"
	"github.com/jeckbjy/fairy/log"
	"github.com/jeckbjy/fairy/tcp"
	"github.com/jeckbjy/fairy/timer"
	"github.com/jeckbjy/fairy/util"
)

type ChatMsg struct {
	Content   string
	Timestamp int64
}

func StartServer() {
	log.Debug("start server")
	// step1: register message
	fairy.RegisterMessage(&ChatMsg{}, nil)

	// step2: register handler
	fairy.RegisterHandler(&ChatMsg{}, func(ctx *fairy.HandlerCtx) {
		req := ctx.Message().(*ChatMsg)
		log.Debug("client msg:%+v", req)

		rsp := &ChatMsg{}
		rsp.Content = "welcome boy!"
		rsp.Timestamp = util.Now()
		ctx.Send(rsp)
	})

	// step3: create transport and add filters
	tran := tcp.NewTran()
	tran.AddFilters(
		filters.NewLogging(),
		filters.NewFrame(frames.NewLine()),
		filters.NewPacket(identities.NewString(), codecs.NewJson()),
		filters.NewExecutor())

	// step4: listen or connect
	tran.Listen(":8080", 0)
}

func StartClient() {
	log.Debug("start client")
	// step1: register message
	fairy.RegisterMessage(&ChatMsg{}, nil)

	// step2: register handler
	fairy.RegisterHandler(&ChatMsg{}, func(ctx *fairy.HandlerCtx) {
		req := ctx.Message().(*ChatMsg)
		log.Debug("server msg:%+v", req)
	})

	var gConn fairy.IConn
	// step3: create transport and add filters
	tran := tcp.NewTran()
	tran.AddFilters(
		filters.NewLogging(),
		filters.NewFrame(frames.NewLine()),
		filters.NewPacket(identities.NewString(), codecs.NewJson()),
		filters.NewExecutor())

	tran.AddFilters(filters.NewConnect(func(conn fairy.IConn) {
		// send msg to server
		req := &ChatMsg{}
		req.Content = "hello word!"
		conn.Send(req)
		gConn = conn
	}))

	// add timer for send message
	timer.Start(timer.ModeLoop, 1000, func() {
		log.Debug("Ontimeout")
		req := &ChatMsg{}
		req.Content = "hello word!"
		req.Timestamp = util.Now()
		gConn.Send(req)
	})

	// step4: listen or connect
	tran.Connect("localhost:8080", 0)
}
