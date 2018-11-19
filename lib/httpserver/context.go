package httpserver

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"

	"icode.baidu.com/baidu/searchbox/ma-srvdemo/lib/log"
)

// Context 请求的上下文
type Context struct {
	Rsp http.ResponseWriter
	Req *http.Request

	data map[string]interface{}
}

// NewContext 创建上下文
// TODO 对象池
func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		Rsp:  w,
		Req:  r,
		data: map[string]interface{}{},
	}
}

//********** Request
func (ctx *Context) FormString(key string) string {
	ctx.Req.ParseForm()
	return ctx.Req.PostForm.Get(key)
}

func (ctx *Context) FormInt64(key string) int64 {
	data := ctx.FormString(key)
	if data == "" {
		return 0
	}

	i, _ := strconv.ParseInt(data, 10, 64)
	return i
}

func (ctx *Context) QueryString(key string) string {
	ctx.Req.ParseForm()
	return ctx.Req.URL.Query().Get(key)
}

func (ctx *Context) QueryInt64(key string) int64 {
	data := ctx.QueryString(key)
	if data == "" {
		return 0
	}

	i, _ := strconv.ParseInt(data, 10, 64)
	return i
}

func (ctx *Context) ReqJson(data interface{}) error {
	if ctx.Req.Body == nil {
		return errors.New("nil body")
	}

	return json.NewDecoder(ctx.Req.Body).Decode(data)
}

//********** Response
// Succ 成功的返回方法
func (ctx *Context) Succ(data interface{}) error {
	return ctx.JSON(map[string]interface{}{
		"code": 0,
		"msg":  "succ",
		"data": data,
	})
}

func (ctx *Context) Fail(err error) error {
	return nil
}

func (ctx *Context) JSON(data interface{}) error {
	bs, err := json.Marshal(data)
	if err != nil {
		return err
	}

	ctx.Rsp.WriteHeader(200)
	ctx.Rsp.Header().Add("Content*Type", "application/json")
	ctx.Rsp.Write(bs)
	return nil
}

//********** Log
func (ctx *Context) Warning(data interface{}) {
	log.Warn(ctx.RequestID(), data)
}

func (ctx *Context) Info(data interface{}) {
	log.Info(ctx.RequestID(), data)
}

const (
	requestIDKey = "___request_id"
)

func (ctx *Context) RequestID() string {
	id := ctx.GetString(requestIDKey)
	if id == "" {
		id = strconv.Itoa(time.Now().Nanosecond())
	}
	ctx.data[requestIDKey] = id
	return id
}

//********** Data
func (ctx *Context) Get(key string) (d interface{}, existed bool) {
	d, existed = ctx.data[key]
	return
}

func (ctx *Context) GetString(key string) string {
	data, existed := ctx.Get(key)

	if !existed {
		return ""
	}

	s, ok := data.(string)
	if !ok {
		return ""
	}

	return s
}
