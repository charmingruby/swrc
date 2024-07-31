package mapper

import "github.com/charmingruby/swrc/proto/pb"

type PingMessage struct {
	Greeting string
}

func PingPongRequestGRPCToObj(p *pb.PingMessage) *PingMessage {
	return &PingMessage{
		Greeting: p.Greeting,
	}
}

func PingPongReplyObjToGRPC(p *PingMessage) *pb.PingMessage {
	return &pb.PingMessage{
		Greeting: p.Greeting,
	}
}
