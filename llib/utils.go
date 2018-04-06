package llib

import (
	"fmt"

	"github.com/fatih/color"
	lua "github.com/yuin/gopher-lua"
)

func printLuaVariable(L *lua.LState) int {
	length := L.GetTop()
	for i := 1; i <= length; i++ {
		printLV(0, L.Get(i))
	}
	fmt.Print("\n")
	return 0
}

func testChecking(L *lua.LState) int {
	v1 := L.CheckAny(1)
	v2 := L.CheckAny(2)
	dbg, ok := L.GetStack(1)
	if ok {
		_, err := L.GetInfo("Sl", dbg, lua.LNil)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
	info := fmt.Sprintf("%s:%d ", dbg.Source, dbg.CurrentLine)
	if v1.Type() != v2.Type() || v1 != v2 {
		fmt.Println(info + "Test failed!!!")
		color.Set(color.FgRed)
		fmt.Printf(info+"'%v' is not equal '%v'\n\n", v1, v2)
		L.Push(lua.LBool(false))
	} else {
		color.Set(color.FgGreen)
		fmt.Println(info + "PASS\n")
		L.Push(lua.LBool(true))
	}
	color.Unset()
	return 1
}

func printLV(level int, v lua.LValue) {
	switch v.Type() {
	case lua.LTTable:
		tb, _ := v.(*lua.LTable)
		tb.ForEach(func(key, value lua.LValue) {
			for i := 0; i < level; i++ {
				fmt.Print("  ")
			}
			printLV(level+1, key)
			fmt.Print("=> ")
			if value.Type() == lua.LTTable {
				fmt.Print("\n")
				printLV(level+1, value)
			} else {
				printLV(level+1, value)
				fmt.Print("\n")
			}
		})
	case lua.LTUserData:
		fmt.Printf("%v ", v.(*lua.LUserData).Value)
	default:
		fmt.Printf("%v ", v)
	}
}
