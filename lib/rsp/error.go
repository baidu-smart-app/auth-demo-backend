package rsp

import (
	"fmt"
)

type Error struct {
	Code int64
	Msg  string
}

func (e *Error) Error() string {
	return fmt.Sprintf("code=%d msg=%s", e.Code, e.Msg)
}

var (
	SysFail      = &Error{10000, "sys fail"}
	ParamIllegal = &Error{10001, "param illegal"}
)

var (
	OrderNotPay     = &Error{20000, "order not paid"}
	OrderNotExisted = &Error{20001, "order not existed"}
)

var (
	AuthNoSessionKey = &Error{30001, "session not existed"}
)
