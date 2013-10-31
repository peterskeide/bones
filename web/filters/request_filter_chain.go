package filters

import (
	"github.com/gorilla/mux"
	"github.com/gorilla/pat"
	"net/http"
)

var router *pat.Router

func SetRouter(r *pat.Router) {
	router = r
}

func Get(pattern string, fn http.HandlerFunc, filters ...FilterFunc) *mux.Route {
	return router.Get(pattern, applyFilters(fn, filters...))
}

func Post(pattern string, fn http.HandlerFunc, filters ...FilterFunc) *mux.Route {
	return router.Post(pattern, applyFilters(fn, filters...))
}

type FilterFunc func(http.ResponseWriter, *http.Request, *RequestFilterChain)

type RequestFilterChain struct {
	res     http.ResponseWriter
	req     *http.Request
	handler http.HandlerFunc
	filters []FilterFunc
}

func (chain *RequestFilterChain) next() {
	if len(chain.filters) > 0 {
		filter := chain.filters[0]
		chain.filters = chain.filters[1:]
		filter(chain.res, chain.req, chain)
	} else {
		chain.handler(chain.res, chain.req)
	}
}

func applyFilters(fn http.HandlerFunc, filters ...FilterFunc) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		chain := RequestFilterChain{res, req, fn, filters[:]}
		chain.next()
	}
}
