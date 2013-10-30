package handlers

import (
	"net/http"
	"strconv"
)

func pathParamInt(req *http.Request, param string) (int, error) {
	return queryParamInt(req, ":"+param)
}

func queryParamInt(req *http.Request, param string) (int, error) {
	strValue := req.URL.Query().Get(param)
	return strconv.Atoi(strValue)
}
