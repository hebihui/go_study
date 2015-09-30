package ipc

import (
	"encoding/json"
	"fmt"
)

type Request struct { //请求对象，包含请求方法和请求参数
	Method string "method"
	Params string "params"
}

type Response struct { //响应对象，包含状态码和返回体
	Code string "code"
	Body string "body"
}

type Server interface { //服务接口，包含处理请求函数，接受请求对象，返回响应对象
	Name() string
	Handle(method, params string) *Response
}

type IpcServer struct { //IPC对象
	Server
}

func NewIpcServer(server Server) *IpcServer {
	return &IpcServer{server}
}

func (server *IpcServer) Connect() chan string { //ipc实现connect方法，返回一个string型chan
	session := make(chan string, 0)

	go func(c chan string) {
		for {
			request := <-c

			if request == "CLOSE" {
				break
			}
			var req Request
			err := json.Unmarshal([]byte(request), &req)
			if err != nil {
				fmt.Println("Invalid request format:", request)
			}

			resp := server.Handle(req.Method, req.Params)
			b, err := json.Marshal(resp)
			c <- string(b) //返回结果
		}
		fmt.Println("session closed.")
	}(session)

	fmt.Println("A new session has been created successfuly.")
	return session
}
