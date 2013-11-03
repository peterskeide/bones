package filters

import (
	"net/http"
)

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

func ApplyTo(fn http.HandlerFunc, filters ...FilterFunc) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		chain := RequestFilterChain{res, req, fn, filters[:]}
		chain.next()
	}
}
