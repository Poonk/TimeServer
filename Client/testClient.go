package main

import (
	"bailun.com/CT4_quote_server/TimeServer/protocol"
	"github.com/astaxie/beego/logs"
	"github.com/gogo/protobuf/proto"
	zmq "github.com/pebbe/zmq4"
)

func main() {
	logs.EnableFuncCallDepth(true)
	logs.SetLogFuncCallDepth(3)

	requester, _ := zmq.NewSocket(zmq.REQ)
	defer requester.Close()
	requester.Connect("tcp://localhost:65002")

	req := &protocol.TimeReq{
		Cmd: protocol.TimeCMD,
	}
	reqBytes, err := proto.Marshal(req)
	if err != nil {
		logs.Error(err)
		return
	}
	logs.Debug(req)
	requester.SendBytes(reqBytes, 0)

	resBytes, err := requester.RecvBytes(0)
	if err != nil {
		logs.Error(err)
		return
	}
	var res protocol.TimeRes
	err = proto.Unmarshal(resBytes, &res)
	if err != nil {
		logs.Error(err)
		return
	}
	logs.Debug(res)
}
