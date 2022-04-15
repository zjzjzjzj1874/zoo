package zoo

import (
	"log"
	"net/http"
)

type HandleFunc func(ctx *Context)

type Engine struct {
	router *router
}

func New() *Engine {
	return &Engine{router: newRouter()}
}

func (e *Engine) addRoute(method, pattern string, handle HandleFunc) {
	log.Printf("Route %4s - %s", method, pattern)
	e.router.addRoute(method, pattern, handle)
}

func (e *Engine) GET(pattern string, handle HandleFunc) {
	e.addRoute(http.MethodGet, pattern, handle)
}

func (e *Engine) POST(pattern string, handle HandleFunc) {
	e.addRoute(http.MethodPost, pattern, handle)
}

func (e *Engine) PUT(pattern string, handle HandleFunc) {
	e.addRoute(http.MethodPut, pattern, handle)
}

func (e *Engine) DELETE(pattern string, handle HandleFunc) {
	e.addRoute(http.MethodDelete, pattern, handle)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := newContext(w, r)
	e.router.handle(c)
}

func (e *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, e)
}
