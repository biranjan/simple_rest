package utils

import (
	"encoding/json"
	"io/util"
	"net/http"
)

func ParseBody(r *http.Request, x interface{}) {
	if body, err := ioutill.ReadAll(r.Body); err == nil {
		if err := json.Unmarshall([]byte(body), x); err != nil {
			return
		}
	}
}
