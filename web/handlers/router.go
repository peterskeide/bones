package handlers

import (
	"github.com/gorilla/pat"
)

var router *pat.Router

func SetRouter(r *pat.Router) {
	router = r
}

// An error occurring in this function is most likely
// caused by incorrectly configured routes or an incorrect
// routeName, so panic to ensure these errors do not
// go unnoticed
func getRouteURL(routeName string, urlArgs ...string) string {
	url, err := router.GetRoute(routeName).URL(urlArgs...)

	if err != nil {
		panic(err)
	}

	return url.String()
}
