package main

import (
	"net"
	"net/rpc"
)

type Server struct{}

func (s *Server) Hello(requset string, reply *string) error {
	*reply = "hello " + requset
	return nil
}

func main() {
	_ = rpc.RegisterName("Server", new(Server))
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	if listener == nil {
		panic("listen failed")
	}
	conn, _ := listener.Accept()

	rpc.ServeConn(conn)

}
