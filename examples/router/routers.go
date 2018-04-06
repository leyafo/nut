package router

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/leyafo/nut"
)

func Routers() *httprouter.Router {
	r := nut.NewRouter()
	r.Handle("GET", "/hello", hello)
	r.Handle("POST", "/post_json", postJSON)

	groupRouter := nut.NewGroupRouter(authCallBack, r)
	groupRouter.Handle("GET", "/basic_auth", hello)
	return r.Routers()
}

func authCallBack(ctx *nut.Context) bool {
	user, password, _ := ctx.Request.BasicAuth()
	if user == "john" && password == "abcdefg" {
		ctx.Set("user", user)
		fmt.Printf("User %s login\n", user)
		return true
	}
	retCode := http.StatusUnauthorized
	http.Error(ctx.ResponseWriter, http.StatusText(retCode), retCode)
	return false
}

func hello(ctx *nut.Context) {
	user := ctx.Get("user")
	if user != nil {
		ctx.Response("hello "+user.(string), http.StatusOK)
	} else {
		ctx.Response("hello world", http.StatusOK)
	}
	ctx.ResponseWriter.Header().Set("Content-Type", "text")
}

func postJSON(ctx *nut.Context) {
	params := make(map[string]string)
	if err := ctx.DecodeJSONBody(&params); err != nil {
		ctx.Response(err.Error(), http.StatusBadRequest)
		return
	}
	t := time.Now()
	params["time"] = t.Format(time.UnixDate)
	backBody, _ := json.Marshal(params) //default back content-type is json
	ctx.ResponseBytes(backBody, http.StatusOK)
}
