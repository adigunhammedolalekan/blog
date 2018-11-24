package main

import (
	"github.com/adigunhammedolalekan/blog/api/service"
	"github.com/emicklei/go-restful"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-web"
	"log"
)

var (
	serviceName = "api.blog.com"
)
func main() {

	srv := web.NewService(
		web.Name(serviceName), web.Version("latest::1"))

	srv.Init()
	handler := service.NewApiService(client.DefaultClient)

	wc := restful.NewContainer()
	webService := new(restful.WebService)

	webService.Path("/api").Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON)
	webService.Route(webService.POST("/account/new").To(handler.NewAccount)).Doc("Create a new account")
	webService.Route(webService.GET("/").To(handler.Hello))

	wc.Add(webService)
	srv.Handle("/", wc)

	if err := srv.Run(); err != nil {
		log.Fatal("Error starting server, ", err)
	}
}
