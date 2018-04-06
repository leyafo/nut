package main

import (
	"fmt"
	"net/http"

	"github.com/leyafo/nut"
)

func main() {
	r := nut.NewRouter()
	r.Handle("get", "/hello", hello)

	groupRouter := nut.NewGroupRouter(authCallBack, r)
	groupRouter.Handle("get", "/basic_auth", hello)

	port := "9527"
	fmt.Printf("server listening port on %s...", port)
	err := http.ListenAndServe(":"+port, r.Routers())
	if err != nil {
		panic(err.Error())
	}
}
