package contract

import "github.com/charmingruby/swrc/proto/pb"

type PingMessage struct {
	Greeting string
}

func PingMessageObjToGRPC(obj PingMessage) pb.PingMessage {
	return pb.PingMessage{
		Greeting: obj.Greeting,
	}
}

func PingMessageGRPCToObj(obj *pb.PingMessage) PingMessage {
	return PingMessage{
		Greeting: obj.Greeting,
	}
}
