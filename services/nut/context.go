package nut

import (
	"encoding/json"
	"net/http"
	"strings"
)

type Context struct {
	Request  *http.Request
	BindReq  interface{}
	BindBody interface{}
	Response struct {
		StatusCode int
		Header     http.Header
		Body       []byte
	}
}

func NewContext(r *http.Request, bindReq, bindBody interface{}) *Context {
	return &Context{Request: r, BindReq: bindReq, BindBody: bindBody}
}

func (c *Context) RawHeader(header http.Header) *Context {
	for key, value := range header {
		c.Header(key, strings.Join(value, ";"))
	}

	return c
}

func (c *Context) Header(name string, value string) *Context {
	c.Response.Header.Add(name, value)
	return c
}

func (c *Context) StatusOk(statusCode int) *Context {
	c.Response.StatusCode = statusCode
	return c
}

func (c *Context) Body(bytes []byte) error {
	c.Response.Body = bytes
	return nil
}

func (c *Context) Json(statusCode int, bytes []byte) error {
	c.StatusOk(statusCode)
	return c.Body(bytes)
}

func (c *Context) Ok(obj interface{}) error {
	bytes, err := json.Marshal(obj)
	if err != nil {
		return err
	}

	return c.Json(http.StatusOK, bytes)
}
