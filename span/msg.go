package span

const (
	TYPE_HEARTBEAT     int32 = 10001
	TYPE_HEARTBEAT_RES int32 = 10002
	TYPE_SPAN          int32 = 10003
)

type Msg struct {
	Type int32
	Body []byte
}
