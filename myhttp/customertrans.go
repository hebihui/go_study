package main

import (
	"net/http"
)

type OurCustomerTransport struct { //自定义Transport结构，组合RoundTripper接口
	Transport http.RoundTripper
}

func (t *OurCustomerTransport) transport() http.RoundTripper { //返回Transport
	if t.Transport != nil {
		return t.Transport
	}
	return http.DefaultTransport
}

func (t *OurCustomerTransport) RoundTrip(req *http.Request) (*http.Response, error) { //实现roundtrpper接口的roundtrip方法
	//do something
	return t.transport().RoundTrip(req)
}

func (t *OurCustomerTransport) Client() *http.Client { //返回自定义Client
	return &http.Client{Transport: t}
}

func main() {
	t := &OurCustomerTransport{
	//...
	}

	c := t.Client() //创建自定义Client
	resp, err := c.Get("http://www.baidu.com")
	//...
}
