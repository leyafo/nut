package nut

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

//Handler ...
type Handler func(ctx *Context)

//CallbackHandle ...
type CallbackHandle func(ctx *Context) bool

//Router ...
type Router struct {
	proxy          *httprouter.Router
	beforeCallback CallbackHandle
}

//NewGroupRouter ...
func NewGroupRouter(beforeCallback CallbackHandle, r *Router) *Router {
	return &Router{
		beforeCallback: beforeCallback,
		proxy:          r.proxy,
	}
}

//NewRouter
func NewRouter() *Router {
	r := httprouter.New()
	return &Router{
		beforeCallback: nil,
		proxy:          r,
	}
}

//Routers return all handled router
func (r Router) Routers() *httprouter.Router {
	return r.proxy
}

//Handle ...
func (r *Router) Handle(method, path string, handle Handler) {
	r.proxy.Handle(method, path, r.generateHandle(handle))
}

func (r *Router) generateHandle(handle Handler) httprouter.Handle {
	return func(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
		if req.Method == "POST" && req.ContentLength <= 0 {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Request body can't be empty")
			return
		}
		ctx := NewContext(w, req, p)
		if r.beforeCallback != nil && !r.beforeCallback(ctx) {
			return
		}
		handle(ctx)
		header := ctx.ResponseWriter.Header()
		if len(header.Get("Content-Type")) == 0 {
			header.Set("Content-Type", "application/json")
		}
	}
}
