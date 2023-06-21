package main

// import (
// 	"blrpc/api"
// 	"blrpc/config"
// 	"blrpc/ent"
// 	"fmt"
// 	"log"
// 	"net"
// 	"net/http"
// 	"strconv"
// 	"strings"

// 	"golang.org/x/net/http2"
// 	"golang.org/x/net/http2/h2c"
// )

// type Server struct {
// 	srv *http.Server
// 	l   net.Listener
// }

// func NewServer(client *ent.Client, l net.Listener, cfg *config.Config) *Server {
// 	strPort := l.Addr().String()[strings.LastIndex(l.Addr().String(), ":")+1:]
// 	port, err := strconv.Atoi(strPort)
// 	if err != nil {
// 		log.Fatalf("failed to parseInt string port %s: %v", strPort, err)
// 	}

// 	if port == cfg.GrpcPort {
// 		gr := api.NewGrpcRouter(client)

// 		return &Server{
// 			srv: &http.Server{Handler: h2c.NewHandler(gr, &http2.Server{})},
// 			l:   l,
// 		}
// 	} else {
// 		hr := api.NewHttpRouter(client)

// 		return &Server{
// 			srv: &http.Server{Handler: hr},
// 			l:   l,
// 		}
// 	}
// }

// func NewListner(port int) (net.Listener, error) {
// 	return net.Listen("tcp", fmt.Sprintf(":%d", port))
// }
