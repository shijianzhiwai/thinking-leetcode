package main

import (
	"fmt"
	"github.com/valyala/fasthttp"
	"log"
	"net"
	"net/http"
	_ "net/http/pprof"
	"time"
)

type testConn struct {
	net.Conn
}

func (t *testConn) Close() error {
	fmt.Println("=======")
	return t.Conn.Close()
}

func main() {
	go func() {
		if err := http.ListenAndServe(":8898", nil); err != nil {
			log.Fatal(err)
		}
	}()

	req := fasthttp.AcquireRequest()
	req.SetRequestURI("http://127.0.0.1:8899/hello")
	resp := fasthttp.AcquireResponse()
	err := fasthttp.DoTimeout(req, resp, 5*time.Second)
	if err != nil {
		fmt.Println(err)
	}

	s := make(chan struct{}, 0)
	<-s
}
