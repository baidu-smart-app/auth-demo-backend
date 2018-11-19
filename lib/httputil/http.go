package httputil

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Client struct {
	Req *http.Request
	Rsp *http.Response
	Cli *http.Client

	err error

	path    string
	RspBody []byte
}

func NewClient() *Client {
	return &Client{
		Req: &http.Request{
			Method:     "GET",
			Proto:      "HTTP/1.1",
			ProtoMajor: 1,
			ProtoMinor: 1,
			Header:     make(http.Header),
			//URL:        u,
			//Host:       u.Host,
			//Body:       rc,
			ContentLength: 0,
			//Body:          http.NoBody,
			GetBody: func() (io.ReadCloser, error) { return http.NoBody, nil },

			// The HTTP client ignores PostForm|Form and uses Body instead, but we can use it
			Form:     make(url.Values),
			PostForm: make(url.Values),
		},

		Rsp: nil,
		Cli: &http.Client{},
	}
}

func (cli *Client) SetTimeOut(timeout time.Duration) *Client {
	cli.Cli.Timeout = timeout
	return cli
}

func (cli *Client) SetPath(path string) *Client {
	cli.path = path
	return cli
}

func (cli *Client) SetMethod(method string) *Client {
	cli.Req.Method = method
	return cli
}

func (cli *Client) PostMethod() *Client {
	cli.Req.Method = http.MethodPost
	return cli
}

func (cli *Client) GetMethod() *Client {
	cli.Req.Method = http.MethodGet
	return cli
}

func (cli *Client) SetHeader(k, v string) *Client {
	cli.Req.Header.Set(k, v)
	return cli
}

func (cli *Client) AddHeader(k, v string) *Client {
	cli.Req.Header.Add(k, v)
	return cli
}

func (cli *Client) AddGetParam(k, v string) *Client {
	cli.Req.Form.Add(k, v)
	return cli
}

func (cli *Client) SetGetParam(k, v string) *Client {
	cli.Req.Form.Set(k, v)
	return cli
}

func (cli *Client) AddPostParam(k, v string) *Client {
	cli.Req.PostForm.Add(k, v)
	return cli
}

func (cli *Client) SetPostParam(k, v string) *Client {
	cli.Req.PostForm.Set(k, v)
	return cli
}

func (cli *Client) Do() *Client {
	if cli.err != nil {
		return cli
	}

	if len(cli.Req.Form) > 0 {
		cli.path += "?" + cli.Req.Form.Encode()
	}

	cli.Req.URL, cli.err = url.Parse(cli.path)
	cli.Req.Host = cli.Req.URL.Host
	if cli.err != nil {
		return cli
	}

	if len(cli.Req.PostForm) > 0 {
		body := strings.NewReader(cli.Req.PostForm.Encode())
		cli.Req.ContentLength = int64(body.Len())
		snapshot := *body
		cli.Req.GetBody = func() (io.ReadCloser, error) {
			r := snapshot
			return ioutil.NopCloser(&r), nil
		}

		cli.Req.Body = ioutil.NopCloser(body)
	}

	cli.Rsp, cli.err = cli.Cli.Do(cli.Req)

	if cli.err != nil {
		return cli
	}

	cli.RspBody, cli.err = ioutil.ReadAll(cli.Rsp.Body)

	return cli
}

func (cli *Client) RspJson(data interface{}, judge func(interface{}) error) *Client {
	if cli.err != nil {
		return cli
	}

	cli.err = json.Unmarshal(cli.RspBody, data)
	if cli.err != nil {
		return cli
	}

	if judge != nil {
		cli.err = judge(data)
	}

	return cli
}

func (cli *Client) Error() *Error {
	if cli.err == nil {
		return nil
	}

	return &Error{
		OriginError: cli.err,
		//Req:         cli.Req,
		//Rsp:         cli.Rsp,
		RspBody: string(cli.RspBody),
	}
}

type Error struct {
	OriginError error
	//Req         *http.Request
	//Rsp         *http.Response

	RspBody string
}

func (err *Error) Error() string {
	bs, _ := json.Marshal(err)
	return string(bs)
}

var BadContent = errors.New("bad content")
