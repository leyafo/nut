package nut

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
)

//Context ..
type Context struct {
	ResponseWriter http.ResponseWriter
	Request        *http.Request
	Params         httprouter.Params

	data map[string]interface{}
}

//NewContext ...
func NewContext(w http.ResponseWriter, req *http.Request, p httprouter.Params) *Context {
	ctx := &Context{
		ResponseWriter: w,
		Request:        req,
		Params:         p,
	}
	return ctx
}

//Set ...
func (c *Context) Set(key string, value interface{}) {
	if c.data == nil {
		c.data = make(map[string]interface{})
	}
	c.data[key] = value
}

//Get ...
func (c *Context) Get(key string) interface{} {
	if c.data == nil {
		return nil
	}
	_, hasKey := c.data[key]
	if hasKey {
		return c.data[key]
	}
	return nil
}

//DecodeJSONBody ...
func (c *Context) DecodeJSONBody(v interface{}) error {
	err := json.NewDecoder(c.Request.Body).Decode(v)
	if err != nil {
		if ute, ok := err.(*json.UnmarshalTypeError); ok {
			return fmt.Errorf("Unmarshal type error: expected=%v, got=%v, offset=%v", ute.Type, ute.Value, ute.Offset)
		} else if se, ok := err.(*json.SyntaxError); ok {
			return fmt.Errorf("Syntax error: offset=%v, error=%v", se.Offset, se.Error())
		} else {
			return err
		}
	}
	return err
}

//Response ...
func (c *Context) Response(response string, code int) {
	c.ResponseWriter.WriteHeader(code)
	if len(response) != 0 {
		fmt.Fprintf(c.ResponseWriter, response)
	}
}

//ResponseBytes ...
func (c *Context) ResponseBytes(bytes []byte, code int) {
	c.ResponseWriter.WriteHeader(code)
	if len(bytes) != 0 {
		c.ResponseWriter.Write(bytes)
	}
}

//GetPathName ...
func (c Context) GetPathName() string {
	return strings.TrimLeft(c.Request.URL.Path, "/")
}

//GetContentType ...
func (c Context) GetContentType() string {
	return c.Request.Header.Get("content-type")
}

//SetContentType ...
func (c Context) SetContentType(key, value string) {
	c.ResponseWriter.Header().Set(key, value)
}
