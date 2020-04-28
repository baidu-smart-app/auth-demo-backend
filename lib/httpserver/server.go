package httpserver

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"auth-demo-backend/lib/rsp"
)

// Server 一个服务器
type Server struct {
	Routers map[string]HandleFunc
}

// HandleFunc 处理方法
type HandleFunc func(*Context) interface{}

// RunServer 运行Http服务
func (server *Server) RunServer(conf *Config) error {
	for path, handle := range server.Routers {
		func(handle HandleFunc) {
			http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
				r.ParseForm()

				var code int64
				msg := "succ"
				var rspData interface{}

				ctx := NewContext(w, r)
				defer HTTPSavior(ctx)

				data := handle(ctx)
				if data == nil {
					ctx.Info(fmt.Sprintf("uri[%s] method[%s] params[%s] code[%d] msg[%s] rsp[%s]",
						ctx.Req.URL.Path, ctx.Req.Method, ctx.Req.Form.Encode(), -1, "", ""))
					return
				}

				switch data.(type) {
				case *rsp.Error:
					r := data.(*rsp.Error)
					code, msg = r.Code, r.Msg
					rspData = map[string]interface{}{
						"code": r.Code,
						"msg":  r.Msg,
					}
				case error:
					r := data.(error)
					code, msg = rsp.SysFail.Code, rsp.SysFail.Msg
					rspData = map[string]interface{}{
						"code":    rsp.SysFail.Code,
						"msg":     rsp.SysFail.Msg,
						"err_msg": r.Error(),
					}
				case *rsp.Raw:
					rspData = data.(*rsp.Raw).Data
				default:
					rspData = map[string]interface{}{
						"code": 0,
						"msg":  "succ",
						"data": data,
					}
				}
				ctx.JSON(rspData)

				bs, _ := json.Marshal(rspData)

				ctx.Info(fmt.Sprintf("uri[%s] method[%s] params[%s] code[%d] msg[%s] rsp[%s]",
					ctx.Req.URL.Path, ctx.Req.Method, ctx.Req.Form.Encode(), code, msg, string(bs)))
			})
		}(handle)
	}

	log.Printf("server running: %s\n", conf.Port)
	return http.ListenAndServe(conf.Port, nil)

}

// RegisterRouter 注册路由
func (server *Server) RegisterRouter(path string, handler HandleFunc) {
	server.Routers[path] = handler
}

// HTTPSavior 注册救世主
var HTTPSavior = func(ctx *Context) {
	if err := recover(); err != nil {
		ctx.Warning(err)

		ctx.JSON(map[string]interface{}{
			"code":    rsp.SysFail.Code,
			"msg":     rsp.SysFail.Msg,
			"err_msg": err.(error).Error(),
		})
		ctx.Info(fmt.Sprintf("params[%s] code[%d] msg[%s]", ctx.Req.Form.Encode(), rsp.SysFail.Code, rsp.SysFail.Msg))
	}
}

// Config 服务器配置
type Config struct {
	Addr string `json:"addr"`
	Port string `json:"port"`
}
