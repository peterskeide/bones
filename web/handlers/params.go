package handlers

import (
	"net/http"
	"strconv"
)

func queryParamInt(req *http.Request, param string) (int, error) {
	strValue := req.URL.Query().Get(":" + param)
	return strconv.Atoi(strValue)
}
