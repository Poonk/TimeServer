package main

import (
	"flag"
	"strconv"
	"time"

	"bailun.com/CT4_quote_server/TimeServer/conf"
	"bailun.com/CT4_quote_server/TimeServer/protocol"
	"github.com/astaxie/beego/logs"
	"github.com/gogo/protobuf/proto"
	zmq "github.com/pebbe/zmq4"
)

func getGMT() string {
	name, offset := time.Now().Zone()
	logs.Debug("gmt", name, offset)
	offset /= conf.Conf.Time.Second
	str := strconv.Itoa(offset)
	if offset >= 0 {
		str = "+" + str
	}
	return str
}

func main() {
	logs.EnableFuncCallDepth(true)
	logs.SetLogFuncCallDepth(3)

	flag.Parse()
	if err := conf.Init(); err != nil {
		logs.Error(err)
		return
	}

	responder, err := zmq.NewSocket(zmq.REP)
	if err != nil {
		logs.Error(err)
	}
	defer responder.Close()
	responder.Bind(conf.Conf.Server.Proto + conf.Conf.Server.Addr)

	for {
		reqBytes, err := responder.RecvBytes(0)
		if err != nil {
			logs.Error(err)
			continue
		}

		var req protocol.TimeReq
		err = proto.Unmarshal(reqBytes, &req)
		if err != nil {
			logs.Error(req)
			logs.Error(err)
			responder.SendBytes(nil, 0)
			continue
		}
		logs.Debug(req)

		logs.Debug(time.Now())
		t := time.Now().Format(protocol.TIME_LAYOUT)
		gmt := getGMT()
		logs.Debug(t, "\t", gmt)

		res := &protocol.TimeRes{
			Time: t,
			GMT:  gmt,
		}
		logs.Debug(res)

		resBytes, err := proto.Marshal(res)
		if err != nil {
			logs.Error(err)
			continue
		}
		responder.SendBytes(resBytes, 0)
	}
}
