package llib

import (
	lua "github.com/yuin/gopher-lua"
)

//Loader ...
func Loader(L *lua.LState) int {
	L.PreloadModule("http", httpLoader)
	L.SetGlobal("to_json", L.NewFunction(toJSON))
	L.SetGlobal("from_json", L.NewFunction(fromJSON))
	L.SetGlobal("put", L.NewFunction(printLuaVariable))
	L.SetGlobal("check", L.NewFunction(testChecking))
	return 0
}

//SetHTTPHost ...
func SetHTTPHost(host string) {
	HTTPHost = host
}
