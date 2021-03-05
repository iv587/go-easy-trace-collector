package client

const (
	typeHeartbeat    int32 = 10001
	typeHeartbeatRes int32 = 10002
	typeSpan         int32 = 10003
)

type Packet struct {
	Type int32
	Body []byte
}
