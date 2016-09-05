package authService

import (
	"github.com/vivek-yadav/UserManagementService/models/authRequest"
	"net"
	//"net/http"
	"net/rpc"
)

func StartService(hostPort string) (e error) {
	// ===== workaround ==========
	//oldMux := http.DefaultServeMux
	//mux := http.NewServeMux()
	//http.DefaultServeMux = mux
	// ===========================

	auth := new(models.AuthRequest)
	//handler := rpc.NewServer()
	//handler.Register(auth)
	//handler.HandleHTTP(rpc.DefaultRPCPath, rpc.DefaultDebugPath)
	//var l net.Listener
	//l, e = net.Listen("tcp", hostPort)
	//if e != nil {
	//	return
	//}
	//
	//go http.Serve(l, nil)
	// ===== workaround ==========
	//http.DefaultServeMux = oldMux
	// ===========================
	rpc.Register(auth)
	tcpAddr, e := net.ResolveTCPAddr("tcp", hostPort)
	if e != nil {
		return
	}

	listener, e := net.ListenTCP("tcp", tcpAddr)
	if e != nil {
		return
	}

	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				continue
			}
			rpc.ServeConn(conn)
		}
	}()
	return
}
