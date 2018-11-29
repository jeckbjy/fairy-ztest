package echo

import (
	"github.com/jeckbjy/fairy"
	"github.com/jeckbjy/fairy-kcp/kcp"
	"github.com/jeckbjy/fairy-protobuf/pbcodec"
	"github.com/jeckbjy/fairy-websocket/ws"
	"github.com/jeckbjy/fairy-ztest/echo/json"
	"github.com/jeckbjy/fairy-ztest/echo/pb"
	"github.com/jeckbjy/fairy/codecs"
	"github.com/jeckbjy/fairy/filters"
	"github.com/jeckbjy/fairy/frames"
	"github.com/jeckbjy/fairy/identities"
	"github.com/jeckbjy/fairy/tcp"
)

var gMsgMode string

func NewTransport(net_mode string, msg_mode string) fairy.ITran {
	var tran fairy.ITran
	switch net_mode {
	case "ws":
		tran = ws.NewTran()
	case "kcp":
		tran = kcp.NewTran()
	default:
		// tcp
		tran = tcp.NewTran()
	}

	var zframe fairy.IFrame
	var zidentity fairy.IIdentity
	var zcodec fairy.ICodec

	switch msg_mode {
	case "pb":
		zframe = frames.NewVarint()
		zidentity = identities.NewFixed16()
		zcodec = pbcodec.New()
	default:
		// json
		zframe = frames.NewLine()
		zidentity = identities.NewString()
		zcodec = codecs.NewJson()
	}

	tran.AddFilters(
		filters.NewLogging(),
		filters.NewFrame(zframe),
		filters.NewPacket(zidentity, zcodec),
		filters.NewExecutor())

	return tran
}

func RegisterMsg(msg_mode string, cb fairy.HandlerCB) {
	switch msg_mode {
	case "pb":
		// protobuf
		Register(cb, &pb.EchoMsg{}, 1)
	default:
		// json
		Register(cb, &json.EchoMsg{}, 0)
	}
}

func Register(cb fairy.HandlerCB, msg interface{}, id int) {
	if id == 0 {
		fairy.RegisterMessage(msg, nil)
		fairy.RegisterHandler(msg, cb)
	} else {
		fairy.RegisterMessage(msg, id)
		fairy.RegisterHandler(id, cb)
	}

}

func SetMsgMode(mode string) {
	gMsgMode = mode
}

func IsJsonMode() bool {
	return gMsgMode == "json"
}

func IsProtobufMode() bool {
	return gMsgMode == "pb"
}
