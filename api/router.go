package api

import (
	"blrpc/api/interceptor"
	"blrpc/controllers"
	"blrpc/ent"
	"blrpc/services"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewHttpRouter(client *ent.Client) *chi.Mux {
	ser := services.NewAppService(client)
	sCon := controllers.NewTaskController(ser)

	hr := chi.NewRouter()

	hr.Use(interceptor.LoggingInterCeptor)
	hr.Use(interceptor.JsonResponseMiddleware)

	hr.Route("/task", func(hr chi.Router) {
		hr.Get("/", sCon.HttpTaskListHandler)
	})

	return hr
}

func NewGrpcRouter(client *ent.Client) *http.ServeMux {
	log.Println("routing grpc server")

	ns := services.NewAppService(client)
	tc := controllers.NewTaskController(ns)

	path, handler := tc.TaskListHandler()

	gr := http.NewServeMux()
	gr.Handle(path, interceptor.LoggingInterCeptor(handler))

	return gr
}
