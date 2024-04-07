package err

import (
	"fmt"
)

const (
	// "incorrect params value" 参数值错误
	ERR_PARAMS = "incorrect params value"
	// "lack of necessary parameters" 必填参数不可为空
	ERR_NECESSARY_PARAMS = "lack of necessary parameters"
	// "item info not exist" 记录信息不存在
	ERR_ITEM_NOT_EXIST = "item info not exist"
	// "effect rows is 0" 操作影响数为 0
	ERR_EFFECT_ROWS_0 = "effect rows is 0"

	// 请勿重复请求
	ERR_LIMIT_DUPLICATED_CODE = 10001
	ERR_LIMIT_DUPLICATED_ZN   = "请勿重复请求"
	// 请求去重校验失败
	ERR_LIMIT_DUPLICATION_FAILED_CODE = 10002
	ERR_LIMIT_DUPLICATION_FAILED_ZN   = "请求去重校验失败"

	// 参数有误，请检查
	ERR_INVALIDED_PARAMS_CODE = 10010
	ERR_INVALIDED_PARAMS_ZN   = "参数有误，请检查"
	// 记录已存在
	ERR_ALREADY_EXIST_CODE = 10011
	ERR_ALREADY_EXIST_ZN   = "记录已存在"

	// 用户名或密码错误
	ERR_INVALIDED_USERNAME_PASSWORD_CODE = 10020
	ERR_INVALIDED_USERNAME_PASSWORD_ZN   = "用户名或密码错误"
	// Token 无效
	ERR_INVALIDED_TOKEN_CODE = 10021
	ERR_INVALIDED_TOKEN_ZN   = "Token 无效"
	// "仅支持系统管理员"
	ERR_ONLY_FOR_ADMIN_CODE = 10022
	ERR_ONLY_FOR_ADMIN_ZN   = "仅支持系统管理员"

	// 内部系统错误
	ERR_SYS_BUSY_CODE  = 50001
	ERR_SYS_BUSY_ZN    = "系统繁忙请稍后"
	ERR_SYS_INNER_CODE = 50010
	ERR_SYS_INNER_ZN   = "服务异常，请稍后"
)

type Err interface {
	// set or add error for debug
	Err(cue string, detail ...interface{}) Err
	// attach string: faield, to explain the call path
	Failed(cue string) Err
	// set or add context value for debug
	Info(cue string, detail ...interface{}) Err
	// set or add error code for front
	Code(code uint32) Err
	// set or add cue msg  for front
	Msg(msg string) Err
	// return full error detail for debug
	Detail() error
	GetMsg() string
	GetCode() string
	GetCodeUint32() uint32
	GetCodeInt32() int32

	Error() string
}

func New() Err {
	return &errs{}
}

type errs struct {
	err  error
	code uint32
	msg  string
}

func (e *errs) Error() string {
	return fmt.Sprintf("{\"err\":\"%+v\",\"msg\":\"%v\",\"code\":\"%v\"}", e.err, e.msg, e.code)
}
func (e *errs) GetMsg() string {
	return fmt.Sprintf("%v", e.msg)
}

func (e *errs) GetCode() string {
	return fmt.Sprintf("%v", e.code)
}
func (e *errs) GetCodeUint32() uint32 {
	return e.code
}
func (e *errs) GetCodeInt32() int32 {
	return int32(e.code)
}

func (e *errs) Detail() error {
	return fmt.Errorf("{\"err\":\"%+v\",\"msg\":\"%v\",\"code\":\"%v\"}", e.err, e.msg, e.code)
}

func (e *errs) Err(cue string, detail ...interface{}) Err {
	if e.err != nil {
		return e.add(cue, detail...)
	}
	return e.set(cue, detail...)
}

func (e *errs) Failed(cue string) Err {
	return e.Err(fmt.Sprintf("%v failed", cue))
}

func (e *errs) set(cue string, detail ...interface{}) *errs {
	if len(detail) == 0 {
		return e.setSolo(cue)
	}
	return e.setMulti(cue, detail[0])
}

func (e *errs) setSolo(cue string) *errs {
	e.err = fmt.Errorf("%v;", cue)
	return e
}

func (e *errs) setMulti(cue string, detail interface{}) *errs {
	e.err = fmt.Errorf("%v err: %v;", cue, detail)
	return e
}

func (e *errs) add(cue string, detail ...interface{}) *errs {
	if len(detail) == 0 {
		return e.addSolo(cue)
	}
	return e.addMulti(cue, detail[0])
}

func (e *errs) addSolo(cue string) *errs {
	e.err = fmt.Errorf("%v; %v", cue, e.err)
	return e
}

func (e *errs) addMulti(cue string, detail interface{}) *errs {
	e.err = fmt.Errorf(" %v err: %+v; %v ", cue, detail, e.err)
	return e
}

func (e *errs) Info(cue string, detail ...interface{}) Err {
	if e.err != nil {
		return e.addInfo(cue, detail)
	}
	return e.setInfo(cue, detail)
}

func (e *errs) setInfo(cue string, detail interface{}) *errs {
	e.err = fmt.Errorf("%v: %+v;", cue, detail)
	return e
}

func (e *errs) addInfo(cue string, detail interface{}) *errs {
	e.err = fmt.Errorf("%v %v: %+v;", e.err, cue, detail)
	return e
}
