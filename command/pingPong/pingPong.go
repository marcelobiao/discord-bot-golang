package pingpong

import "regexp"

type PingPong struct {
}

func (d *PingPong) Match(content string) (matched bool, err error) {
	matched, err = regexp.MatchString("^ping$", content)
	return
}

func (d *PingPong) Doc() (documentation string) {
	documentation = "ping: Retorna pong"
	return
}

func (d *PingPong) Run(content string) (response string, err error) {
	response = "pong"
	return
}
