package echo

import (
	"fmt"

	"github.com/jeckbjy/fairy/exit"

	"github.com/jeckbjy/fairy"
	"github.com/jeckbjy/fairy-ztest/echo/json"
	"github.com/jeckbjy/fairy-ztest/echo/pb"
	"github.com/jeckbjy/fairy/filters"
	"github.com/jeckbjy/fairy/log"
	"github.com/jeckbjy/fairy/timer"
	"github.com/jeckbjy/fairy/util"
)

var gClient fairy.IConn

func SendEchoToServer() {
	log.Debug("send msg to server!")
	if IsJsonMode() {
		req := &json.EchoMsg{}
		req.Info = "Client json.Echo!"
		req.Timestamp = util.Now()
		gClient.Send(req)
	} else {
		req := &pb.EchoMsg{}
		req.Info = "Client pb.Echo!"
		req.Timestamp = util.Now()
		gClient.Send(req)
	}
}

func OnTimeout() {
	log.Debug("timeout")
	if gClient == nil {
		return
	}

	SendEchoToServer()
}

func OnConnected(conn fairy.IConn) {
	log.Debug("OnConnected")
	gClient = conn
	SendEchoToServer()
}

func OnClientEcho(ctx *fairy.HandlerCtx) {
	rsp := ctx.Message()
	log.Debug("Recv server echo: %+v", rsp)
}

func StartClient(net_mode string, msg_mode string) {
	fmt.Printf("start client:net_mode=%v, msg_mode=%v\n", net_mode, msg_mode)

	SetMsgMode(msg_mode)
	RegisterMsg(msg_mode, OnClientEcho)

	tran := NewTransport(net_mode, msg_mode)
	tran.AddFilters(filters.NewConnect(OnConnected))

	tran.Connect("localhost:8888", 0)
	tran.Start()

	timer.Start(timer.ModeLoop, 500, OnTimeout)
	exit.Wait()
	fmt.Printf("stop client!\n")
}
