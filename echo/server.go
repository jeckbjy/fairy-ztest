package echo

import (
	"fmt"

	"github.com/jeckbjy/fairy/exit"

	"github.com/jeckbjy/fairy"
	"github.com/jeckbjy/fairy-ztest/echo/json"
	"github.com/jeckbjy/fairy-ztest/echo/pb"
	"github.com/jeckbjy/fairy/log"
	"github.com/jeckbjy/fairy/util"
)

func OnServerEcho(ctx *fairy.HandlerCtx) {
	if IsJsonMode() {
		req := ctx.Message().(*json.EchoMsg)
		log.Debug("Recv client echo: %+v", req)
		rsp := &json.EchoMsg{}
		rsp.Info = "server rsp!"
		rsp.Timestamp = util.Now()
		ctx.Send(rsp)
	} else if IsProtobufMode() {
		req := ctx.Message().(*pb.EchoMsg)
		log.Debug("Recv client echo: %+v", req)
		rsp := &pb.EchoMsg{}
		rsp.Info = "server rsp!"
		rsp.Timestamp = util.Now()
		ctx.Send(rsp)
	}
}

func StartServer(net_mode string, msg_mode string) {
	fmt.Printf("start server:net_mode=%v, msg_mode=%v\n", net_mode, msg_mode)

	SetMsgMode(msg_mode)
	RegisterMsg(msg_mode, OnServerEcho)
	tran := NewTransport(net_mode, msg_mode)

	tran.Listen(":8888", 0)
	tran.Start()

	exit.Wait()
	fmt.Sprintf("stop server!\n")
}
