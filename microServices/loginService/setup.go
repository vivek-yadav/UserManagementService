package loginService

import (
	"github.com/vivek-yadav/UserManagementService/models/user"
	"net"
	"net/http"
	"net/rpc"
)

func StartService(hostPort string) (e error) {
	// ===== workaround ==========
	oldMux := http.DefaultServeMux
	mux := http.NewServeMux()
	http.DefaultServeMux = mux
	// ===========================

	user := new(models.User)
	handler := rpc.NewServer()
	handler.Register(user)
	handler.HandleHTTP(rpc.DefaultRPCPath, rpc.DefaultDebugPath)
	var l net.Listener
	l, e = net.Listen("tcp", hostPort)
	if e != nil {
		return
	}

	go http.Serve(l, nil)
	// ===== workaround ==========
	http.DefaultServeMux = oldMux
	// ===========================
	return
}
