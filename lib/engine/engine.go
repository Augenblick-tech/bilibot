package engine

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Engine struct {
	g *gin.Engine
	*RouteGroup
}

type RouteGroup struct {
	Path string
	g    *gin.RouterGroup
	i    InitFunc
	m    []Middleware
}

type Context struct {
	Context   *gin.Context
}

type InitFunc interface {
	Done(*RouteGroup, string, string, Handle) error
}

type Handle func(*Context) (interface{}, error)
type Middleware func(Handle) Handle

func NewDefaultEngine() *Engine {
	g := gin.Default()
	return &Engine{
		g: g,
		RouteGroup: &RouteGroup{
			Path: "",
			g:    g.Group(""),
			m:    make([]Middleware, 0),
		},
	}
}

func SetMode(mode string) {
	gin.SetMode(mode)
}

func (e *Engine) Run(addr string) error {
	return e.g.Run(addr)
}

func (e *Engine) RunTLS(addr string, certFile string, keyFile string) error {
	return e.g.RunTLS(addr, certFile, keyFile)
}

func (e *Engine) Group(path string) *RouteGroup {
	return &RouteGroup{
		Path: path,
		g:    e.g.Group(path),
		m:    make([]Middleware, 0),
	}
}

func (g *RouteGroup) Group(path string) *RouteGroup {
	return &RouteGroup{
		Path: g.Path + path,
		g:    g.g.Group(path),
		m:    make([]Middleware, 0),
	}
}

func (e *Engine) LoadHTMLGlob(pattern string) {
	e.g.LoadHTMLGlob(pattern)
}

func (e *Engine) LoadHTMLFiles(files ...string) {
	e.g.LoadHTMLFiles(files...)
}

func (r *RouteGroup) Use(middleware ...Middleware) *RouteGroup {
	r.m = append(r.m, middleware...)
	return r
}

func (r *RouteGroup) GET(path, name string, h Handle) {
	if r.i != nil {
		r.i.Done(r, path, name, h)
	}
	for _, i := range r.m {
		h = i(h)
	}
	r.g.GET(path, handle(h))
}

func (r *RouteGroup) POST(path, name string, h Handle) {
	if r.i != nil {
		r.i.Done(r, path, name, h)
	}
	for _, i := range r.m {
		h = i(h)
	}
	// h = post(h)
	r.g.POST(path, handle(h))
}

func (r *RouteGroup) StaticFs(path string, fs http.FileSystem) {
	r.g.StaticFS(path, fs)
}

func (r *RouteGroup) Static(path, dir string) {
	r.g.Static(path, dir)
}

func handle(h Handle) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := &Context{
			Context: c,
		}

		h(ctx)
	}
}

func (r *RouteGroup) Init(b InitFunc) *RouteGroup {
	r.i = b
	return r
}

func (c *Context) PostBody() map[string]interface{} {
	var m map[string]interface{}
	json.NewDecoder(c.Context.Request.Body).Decode(&m)
	return m
}

func (c *Context) Bind(i interface{}) error {
	return c.Context.ShouldBind(i)
}

func (c *Context) ImageResult(b []byte, s string) {
	c.Context.Data(200, "image/"+s, b)
}

func (c *Context) TextResult(b []byte) {
	c.Context.Data(200, "text/plain", b)
}