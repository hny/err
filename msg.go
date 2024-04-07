package err

import "fmt"

func (e *errs) Msg(msg string) Err {
	if e.msg == "" {
		return e.setMsg(msg)
	}
	return e.addMsg(msg)
}

func (e *errs) setMsg(msg string) *errs {
	e.msg = msg
	return e
}

func (e *errs) addMsg(msg string) *errs {
	e.msg = fmt.Sprintf("%s; %s", e.msg, msg)
	return e
}
