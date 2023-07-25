package main

import (
	"fmt"
	"net/http"
	"web-server/framework"
	"web-server/router"
)

func main() {
	core := framework.NewCore()
	router.RegisterRouter(core)
	server := &http.Server{
		Handler: core,
		Addr:    ":8888",
	}
	fmt.Print("start......")
	server.ListenAndServe()
}