package main

import (
	"fmt"
	"net/http/httptest"
	"os"

	"github.com/leyafo/nut/examples/router"
	"github.com/leyafo/nut/llib"
	lua "github.com/yuin/gopher-lua"
)

func setPath(L *lua.LState) {
	tb := L.GetGlobal("package")
	p, _ := tb.(*lua.LTable)
	path := p.RawGetString("path")
	p.RawSetString("path", lua.LString(fmt.Sprintf("%s;%s", path, "./?.lua")))
	L.SetGlobal("package", p)
}

func main() {
	ts := httptest.NewServer(router.Routers())
	if ts != nil {
		defer ts.Close()
	}

	//loat lua scripts
	L := lua.NewState()
	defer L.Close()

	setPath(L)
	llib.Loader(L)
	llib.SetHTTPHost(ts.URL)

	args := os.Args
	if len(args) > 1 {
		err := L.DoFile("./examples/" + args[1])
		if err != nil {
			panic(err.Error())
		}
	}
}
