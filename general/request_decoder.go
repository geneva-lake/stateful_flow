package general

import (
	"encoding/json"
	"net/http"
)

func RequestDecode[R any](r *http.Request) (*R, error) {
	request := new(R)
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(request)
	if err != nil {
		return nil, err
	}
	return request, nil
}
