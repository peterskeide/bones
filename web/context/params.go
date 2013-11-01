package context

import (
	"github.com/gorilla/context"
	"net/http"
	"strconv"
	"strings"
)

type RequestParam struct {
	Request *http.Request
}

func Params(req *http.Request) *RequestParam {
	if rp, ok := context.Get(req, params).(*RequestParam); ok {
		return rp
	}

	return nil
}

func InitParams(req *http.Request) error {
	err := req.ParseForm()

	if err != nil {
		return err
	}

	rp := RequestParam{req}

	context.Set(req, params, &rp)

	return nil
}

func (rp *RequestParam) Get(param string) string {
	if strings.HasPrefix(param, ":") {
		return rp.Request.URL.Query().Get(param)
	}

	return rp.Request.Form.Get(param)
}

func (rp *RequestParam) GetInt(param string) (int, error) {
	return strconv.Atoi(rp.Get(param))
}
