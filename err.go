package err

import (
	"fmt"
)

type Err interface {
	// 设置 code 值，多次调用会覆盖；可以使用 GetCode() 获取，
	// 会在打印 error 时一并打印出来;
	Code(code uint32) Err
	// 增加 error 信息，多次调用会追加至已存在的 error
	Err(cue string, detail ...interface{}) Err
	// 增加 error 信息 相当于调用 err.errs(cue+" failed")
	Failed(cue string) Err
	// 增加错误相关业务信息， 会在打印error时一并打印出来
	Info(cue string, detail ...interface{}) Err
	// 增加 Message 信息，多次调用会追加至已存在的 Message ；
	// 可以使用 GetMsg() 获取，会在打印 error 时一并打印出来;
	Msg(msg string) Err

	// 获取 Code 值
	GetCode() string
	// 获取 Code 值
	GetCodeInt32() int32
	// 获取 Code 值
	GetCodeUint32() uint32
	// 获取 Message 信息
	GetMsg() string

	// 实现 String 接口 打印结构体时会自动调用
	String() string
	// 实现 errsor 接口
	Error() string
}

type errs struct {
	err  error
	code uint32
	msg  string
}

// 创建一个 Err 实例, cues 选填,
// New("handle", "invalid params") 等同于 New().Err(handle,"invalid params")
func New(cues ...interface{}) Err {
	if len(cues) <= 0 {
		return &errs{}
	}

	err := &errs{}

	return err.Err(
		fmt.Sprint(cues[0]), cues[1:]...,
	)
}

// 增加 error 信息，多次调用会追加至已存在的 error
func (e *errs) Err(
	cue string, detail ...interface{},
) Err {
	if e.err != nil {
		return e.add(cue, detail...)
	}

	return e.set(cue, detail...)
}
func (e *errs) add(
	cue string, detail ...interface{},
) Err {
	if len(detail) == 0 {
		return e.addSolo(cue)
	}

	return e.addMulti(cue, detail[0])
}
func (e *errs) addSolo(cue string) Err {
	e.err = fmt.Errorf("%v; %v", cue, e.err)
	return e
}
func (e *errs) addMulti(
	cue string, detail interface{},
) Err {
	e.err = fmt.Errorf(
		" %v err: %+v; %v ",
		cue, detail, e.err,
	)

	return e
}
func (e *errs) set(
	cue string, detail ...interface{},
) Err {
	if len(detail) == 0 {
		return e.setSolo(cue)
	}

	return e.setMulti(cue, detail[0])
}
func (e *errs) setSolo(cue string) Err {
	e.err = fmt.Errorf("%v;", cue)
	return e
}
func (e *errs) setMulti(
	cue string, detail interface{},
) Err {
	e.err = fmt.Errorf(
		"%v err: %v;",
		cue, detail,
	)

	return e
}

// 增加 error 信息 相当于调用 err.errs(cue+" failed")
func (e *errs) Failed(cue string) Err {
	return e.Err(
		fmt.Sprintf("%v failed", cue),
	)
}

// 增加错误相关业务信息， 会在打印error时一并打印出来
func (e *errs) Info(
	cue string, detail ...interface{},
) Err {
	if e.err != nil {
		return e.addInfo(cue, detail)
	}

	return e.setInfo(cue, detail)
}
func (e *errs) setInfo(
	cue string, detail interface{},
) *errs {
	e.err = fmt.Errorf("%v: %+v;", cue, detail)
	return e
}
func (e *errs) addInfo(
	cue string, detail interface{},
) *errs {
	e.err = fmt.Errorf(
		"%v %v: %+v;",
		e.err, cue, detail,
	)

	return e
}

// 设置 code 值，多次调用会覆盖；可以使用 GetCode() 获取，
// 会在打印 error 时一并打印出来;
func (e *errs) Code(code uint32) Err {
	e.code = code

	return e
}

// 增加 Message 信息，多次调用会追加至已存在的 Message ；
// 可以使用 GetMsg() 获取，会在打印 error 时一并打印出来;
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

// 获取 Message 信息
func (e *errs) GetMsg() string {
	return fmt.Sprintf("%v", e.msg)
}

// 获取 Code 值
func (e *errs) GetCode() string {
	return fmt.Sprintf("%v", e.code)
}

// 获取 Code 值
func (e *errs) GetCodeUint32() uint32 {
	return e.code
}

// 获取 Code 值
func (e *errs) GetCodeInt32() int32 {
	return int32(e.code)
}

// 实现 errsor 接口
func (e *errs) Error() string {
	return e.json()
}

// 实现 String 接口 打印结构体时会自动调用
func (e *errs) String() string {
	return e.json()
}
func (e *errs) json() string {
	err := `"nil"`
	if e.err != nil {
		err = "'" + e.err.Error() + "'"
	}

	return fmt.Sprint(
		"{",
		`"err": `, err, ",",
		`"code": `, e.code, ",",
		`"msg": "`, e.msg, `"`,
		"}",
	)
}
