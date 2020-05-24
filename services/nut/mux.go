package nut

import (
	"github.com/go-chi/chi"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

type Mux struct {
	router            chi.Router
	catchErrorHandler CatchErrorHandler
	records           []Record
}

func NewMux(router chi.Router, handler CatchErrorHandler) *Mux {
	return &Mux{router: router, catchErrorHandler: handler, records: []Record{}}
}

func (m *Mux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	m.router.ServeHTTP(w, r)
}

func (m *Mux) Use(middlewares ...func(http.Handler) http.Handler) {
	m.router.Use(middlewares...)
}

func (m *Mux) With(middlewares ...func(http.Handler) http.Handler) Router {
	router := m.router.With(middlewares...)
	return NewMux(router, m.catchErrorHandler)
}

func (m *Mux) Group(fn func(r Router)) Router {
	router := m.router.Group(func(r chi.Router) {
		fn(NewMux(r, m.catchErrorHandler))
	})
	return NewMux(router, m.catchErrorHandler)
}

func (m *Mux) Route(pattern string, fn func(r Router)) Router {
	router := m.router.Route(pattern, func(r chi.Router) {
		fn(NewMux(r, m.catchErrorHandler))
	})
	return NewMux(router, m.catchErrorHandler)
}

func (m *Mux) Mount(pattern string, h http.Handler) {
	m.router.Mount(pattern, h)
}

func (m *Mux) Handle(pattern string, h http.Handler) {
	m.router.Handle(pattern, h)
}

func (m *Mux) HandleFunc(pattern string, h http.HandlerFunc) {
	m.router.HandleFunc(pattern, h)
}

func (m *Mux) Method(method, pattern string, h http.Handler) {
	m.router.Method(method, pattern, h)
}

func (m *Mux) MethodFunc(method, pattern string, h http.HandlerFunc) {
	m.router.MethodFunc(method, pattern, h)
}

func (m *Mux) NotFound(h http.HandlerFunc) {
	m.router.NotFound(h)
}

func (m *Mux) MethodNotAllowed(h http.HandlerFunc) {
	m.router.MethodNotAllowed(h)
}

func (m *Mux) Connect(pattern string, h Handler) {
	m.router.Connect(pattern, m.sweepHandler(h, m.catchErrorHandler))
}

func (m *Mux) Delete(pattern string, h Handler) {
	m.router.Delete(pattern, m.sweepHandler(h, m.catchErrorHandler))
}

func (m *Mux) Get(pattern string, h Handler) {
	m.router.Get(pattern, m.sweepHandler(h, m.catchErrorHandler))
}

func (m *Mux) Head(pattern string, h Handler) {
	m.router.Head(pattern, m.sweepHandler(h, m.catchErrorHandler))
}

func (m *Mux) Options(pattern string, h Handler) {
	m.router.Options(pattern, m.sweepHandler(h, m.catchErrorHandler))
}

func (m *Mux) Patch(pattern string, h Handler) {
	m.router.Patch(pattern, m.sweepHandler(h, m.catchErrorHandler))
}

func (m *Mux) Post(pattern string, h Handler) {
	m.router.Post(pattern, m.sweepHandler(h, m.catchErrorHandler))
}

func (m *Mux) Put(pattern string, h Handler) {
	m.router.Put(pattern, m.sweepHandler(h, m.catchErrorHandler))
}

func (m *Mux) Trace(pattern string, h Handler) {
	m.router.Trace(pattern, m.sweepHandler(h, m.catchErrorHandler))
}

func (m *Mux) Resource(pattern string, ctl ResourceController) {
	for _, record := range ctl.Routes() {
		vr := reflect.ValueOf(record)
		if h := vr.MethodByName("CatchAllError"); h.IsValid() {
			record.catchErrorHandler = h.Interface().(CatchErrorHandler)
		}
		m.records = append(m.records, record)
	}

	m.router.Route(pattern, func(r chi.Router) {
		records := ctl.Routes()
		for _, record := range records {
			ceh := m.catchErrorHandler
			if record.catchErrorHandler != nil {
				ceh = record.catchErrorHandler
			}
			h := m.sweepHandler(m.resourceGuard(m.resourceBind(record.handler, record), record), ceh)

			switch record.method {
			case recordMethodGet:
				r.Get(record.path, h)
			case recordMethodPost:
				r.Post(record.path, h)
			case recordMethodPut:
				r.Put(record.path, h)
			case recordMethodPatch:
				r.Patch(record.path, h)
			case recordMethodDelete:
				r.Delete(record.path, h)
			case recordMethodOptions:
				r.Options(record.path, h)
			}
		}
	})
}

func (m *Mux) CatchAllError(handler CatchErrorHandler) {
	m.catchErrorHandler = handler
}

func (m *Mux) sweepHandler(h Handler, ceh CatchErrorHandler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := NewContext(r, nil, nil)
		err := h(ctx)

		if err != nil && ceh != nil {
			err = ceh(err, r)
		}

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			if _, err := w.Write([]byte(err.Error())); err != nil {
				panic(err)
			}
			return
		}

		for key, value := range ctx.Response.Header {
			w.Header().Set(key, strings.Join(value, ";"))
		}

		if ctx.Response.StatusCode == 0 {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(ctx.Response.StatusCode)
		}

		if _, err := w.Write(ctx.Response.Body); err != nil {
			panic(err)
		}
	})
}

func (m *Mux) resourceGuard(next Handler, record Record) Handler {
	return func(ctx *Context) error {
		for _, guard := range record.guards {
			if err := guard(ctx.Request); err != nil {
				return err
			}
		}
		return next(ctx)
	}
}

func (m *Mux) resourceBind(next Handler, record Record) Handler {
	return func(ctx *Context) error {
		if record.bindReq != nil {
			// 从 Chi 的 Context 中读取 URL Params
			urlParams := map[string]string{}
			chiCtx := ctx.Request.Context().Value(chi.RouteCtxKey).(*chi.Context)
			for i, key := range chiCtx.URLParams.Keys {
				urlParams[key] = chiCtx.URLParams.Values[i]
			}

			bindReqCopy := reflect.New(reflect.TypeOf(record.bindReq).Elem())
			err := travelStructTag(bindReqCopy, func(tagType, tagName string) (string, bool) {
				name := strings.ToLower(tagName[:1]) + tagName[1:]
				switch tagType {
				case "path":
					if val, ok := urlParams[name]; ok {
						return val, false
					} else {
						return "", true
					}
				case "query":
					if val, ok := ctx.Request.URL.Query()[name]; ok {
						return strings.Join(val, ","), false
					} else {
						return "", true
					}
				case "header":
					if val, ok := ctx.Request.Header[name]; ok {
						return strings.Join(val, ";"), false
					} else {
						return "", true
					}
				}
				return "", true
			})
			if err != nil {
				return err
			}

			ctx.BindReq = bindReqCopy.Interface()
		}

		if record.bindBody != nil && record.bindParser != nil {
			bindBodyCopy := reflect.New(reflect.TypeOf(record.bindBody).Elem()).Interface()
			if err := record.bindParser.Parse(ctx.Request, bindBodyCopy); err != nil {
				return err
			}

			ctx.BindBody = bindBodyCopy
		}

		return next(ctx)
	}
}

var tagTypes = []string{"query", "header", "path"}

func travelStructTag(vv reflect.Value, cb func(tagType, tagName string) (string, bool)) error {
	switch vv.Kind() {
	case reflect.Ptr:
		return travelStructTag(vv.Elem(), cb)
	case reflect.Struct:
		vt := vv.Type()
		for i := 0; i < vv.NumField(); i++ {
			fv := vv.Field(i)
			ft := vt.Field(i)
			if fv.Kind() == reflect.Struct {
				if err := travelStructTag(fv, cb); err != nil {
					return err
				}
				continue
			}
			if fv.Kind() == reflect.Ptr && fv.Elem().Kind() == reflect.Struct {
				if err := travelStructTag(fv.Elem(), cb); err != nil {
					return err
				}
				continue
			}

			for _, tagType := range tagTypes {
				if ft.Tag.Get(tagType) != "" {
					resultStr, isEmpty := cb(tagType, ft.Name)
					rv, err := strConvert(fv, resultStr, isEmpty)
					if err != nil {
						return err
					}
					fv.Set(rv)
				}
			}
		}
	}
	return nil
}

func strConvert(v reflect.Value, value string, isEmpty bool) (reflect.Value, error) {
	switch v.Kind() {
	// 支持 Ptr, 但当数据不存在时， 则为 nil
	// 不支持指针的指针
	// 不支持 map
	// 不支持 slice
	case reflect.Ptr:
		if isEmpty {
			return reflect.ValueOf(nil), nil
		}

		switch v.Elem().Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			if value == "" {
				n := 0
				return reflect.ValueOf(&n), nil
			}
			num, err := strconv.Atoi(value)
			if err != nil {
				return reflect.Value{}, err
			}
			return reflect.ValueOf(&num), nil

		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			if value == "" {
				n := 0
				return reflect.ValueOf(&n), nil
			}
			num, err := strconv.Atoi(value)
			if err != nil {
				return reflect.Value{}, err
			}
			return reflect.ValueOf(&num), nil

		case reflect.String:
			return reflect.ValueOf(&value), nil

		case reflect.Bool:
			b := value != ""
			return reflect.ValueOf(&b), nil
		}

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if value == "" {
			return reflect.ValueOf(0), nil
		}
		num, err := strconv.Atoi(value)
		if err != nil {
			return reflect.Value{}, err
		}
		return reflect.ValueOf(num), nil

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if value == "" {
			return reflect.ValueOf(0), nil
		}
		num, err := strconv.Atoi(value)
		if err != nil {
			return reflect.Value{}, err
		}
		return reflect.ValueOf(num), nil

	case reflect.String:
		return reflect.ValueOf(value), nil

	case reflect.Bool:
		b := value != ""
		return reflect.ValueOf(b), nil
	}

	return reflect.Value{}, nil
}
