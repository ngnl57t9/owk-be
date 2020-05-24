package nut

import (
	"fmt"
	"owknight-be/services/nut/parser"
)

var typeToContentType = map[string]string{
	"json": "application/json",
}
var typeToParser = map[string]parser.Parser{
	"json": parser.NewJsonParser(),
}

type recordMethod int

const (
	recordMethodGet recordMethod = iota
	recordMethodPost
	recordMethodDelete
	recordMethodPut
	recordMethodPatch
	recordMethodOptions
)

type Record struct {
	method            recordMethod
	path              string
	guards            []Guard
	bindReq           interface{}
	bindContentType   string
	bindBody          interface{}
	bindParser        parser.Parser
	handler           Handler
	catchErrorHandler CatchErrorHandler
}

func Get(path string, handler Handler) Record {
	r := Record{}
	r.method = recordMethodGet
	r.path = path
	r.handler = handler
	return r
}

func Post(path string, handler Handler) Record {
	r := Record{}
	r.method = recordMethodPost
	r.path = path
	r.handler = handler
	return r
}

func Put(path string, handler Handler) Record {
	r := Record{}
	r.method = recordMethodPut
	r.path = path
	r.handler = handler
	return r
}

func Patch(path string, handler Handler) Record {
	r := Record{}
	r.method = recordMethodPatch
	r.path = path
	r.handler = handler
	return r
}

func Delete(path string, handler Handler) Record {
	r := Record{}
	r.method = recordMethodDelete
	r.path = path
	r.handler = handler
	return r
}

func Options(path string, handler Handler) Record {
	r := Record{}
	r.method = recordMethodOptions
	r.path = path
	r.handler = handler
	return r
}

func (r Record) Guard(guards ...Guard) Record {
	r.guards = append(r.guards, guards...)
	return r
}

func (r Record) BindReq(bind interface{}) Record {
	r.bindReq = bind
	return r
}

func (r Record) BindBody(bodyType string, bind interface{}) Record {
	bindContentType := typeToContentType[bodyType]
	if bindContentType == "" {
		panic(fmt.Sprintf("BindBody(BodyType):%s undefined", bodyType))
	}

	bindParser := typeToParser[bodyType]
	if bindParser == nil {
		panic(fmt.Sprintf("BindBody(BodyType):%s undefined", bodyType))
	}

	r.bindContentType = bindContentType
	r.bindParser = bindParser
	r.bindBody = bind
	return r
}
