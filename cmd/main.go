package main

import (
	"blrpc/api"
	"blrpc/config"
	"blrpc/ent"
	"blrpc/repositories"
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"strconv"
	"strings"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
	}
	client := repositories.OpenDB(context.Background(), cfg)

	// httpLner, err := NewListner(cfg.HttpPort)
	// if err != nil {
	// 	log.Fatalf("failed to listen port %d: %v", cfg.HttpPort, err)
	// }
	// httpSrv := NewServer(client, httpLner, cfg)

	// grpcLner, err := NewListner(cfg.GrpcPort)
	// if err != nil {
	// 	log.Fatalf("failed to listen port %d: %v", cfg.GrpcPort, err)
	// }
	// grpcSrv := NewServer(client, grpcLner, cfg)

	// go func() error {
	// 	log.Printf("start with: %v", fmt.Sprintf("http://%s", httpSrv.l.Addr().String()))
	// 	return httpSrv.srv.ListenAndServe()
	// }()

	// log.Printf("start with: %v", fmt.Sprintf("grpc://%s", grpcSrv.l.Addr().String()))
	// grpcSrv.srv.ListenAndServe()

	go func() error {
		router := api.NewHttpRouter(client)
		server := http.Server{
			Addr:    ":9090",
			Handler: router,
		}

		return server.ListenAndServe()
	}()

	mux := api.NewGrpcRouter(client)
	http.ListenAndServe(
		"localhost:8080",
		h2c.NewHandler(mux, &http2.Server{}),
	)
}

type Server struct {
	srv *http.Server
	l   net.Listener
}

func NewServer(client *ent.Client, l net.Listener, cfg *config.Config) *Server {
	strPort := l.Addr().String()[strings.LastIndex(l.Addr().String(), ":")+1:]
	port, err := strconv.Atoi(strPort)
	if err != nil {
		log.Fatalf("failed to parseInt string port %s: %v", strPort, err)
	}

	log.Println("start server")

	if port == cfg.GrpcPort {
		gr := api.NewGrpcRouter(client)

		return &Server{
			srv: &http.Server{Handler: h2c.NewHandler(gr, &http2.Server{})},
			l:   l,
		}
	} else {
		hr := api.NewHttpRouter(client)

		return &Server{
			srv: &http.Server{Handler: hr},
			l:   l,
		}
	}
}

func NewListner(port int) (net.Listener, error) {
	return net.Listen("tcp", fmt.Sprintf(":%d", port))
}
