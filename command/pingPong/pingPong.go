package pingpong

type PingPong struct {
}

func (d *PingPong) Run(content string) (response string, err error) {
	response = "pong"
	return
}
