package ipc

import (
	"testing"
)

type EchoServer struct {
}

func (server *EchoServer) Handle(request, params string) *Response {
	tmp := &Response{}
	tmp.Body = params
	tmp.Code = request
	return tmp
}

func (server *EchoServer) Name() string {
	return "EchoServer"
}

func TestIpc(t *testing.T) {
	server := NewIpcServer(&EchoServer{})

	client1 := NewIpcClient(server)
	client2 := NewIpcClient(server)

	resp1, _ := client1.Call("echo", "From Client1")
	resp2, _ := client2.Call("echo", "From Client2")

	if resp1.Body != "From Client1" || resp2.Body != "From Client2" {
		t.Error("ipcclinet call failed.resp1:", resp1, "resp2:", resp2)
	}
	client1.Close()
	client2.Close()
}
