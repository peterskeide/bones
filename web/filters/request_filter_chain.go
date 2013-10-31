package filters

import (
	"github.com/gorilla/mux"
	"github.com/gorilla/pat"
	"net/http"
)

type FilterFunc func(http.ResponseWriter, *http.Request) bool

type RequestFilterChain struct {
	Router *pat.Router
}

func NewRequestFilterChain() *RequestFilterChain {
	router := pat.New()
	return &RequestFilterChain{router}
}

func (rfc *RequestFilterChain) Get(pattern string, fn http.HandlerFunc, filters ...FilterFunc) *mux.Route {
	return rfc.Router.Get(pattern, applyFilters(fn, filters...))
}

func (rfc *RequestFilterChain) Post(pattern string, fn http.HandlerFunc, filters ...FilterFunc) *mux.Route {
	return rfc.Router.Post(pattern, applyFilters(fn, filters...))
}

func applyFilters(fn http.HandlerFunc, filters ...FilterFunc) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		for _, filterFn := range filters {
			if !filterFn(res, req) {
				return
			}
		}

		fn(res, req)
	}
}
