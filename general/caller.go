package general

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func MakeHTTPRequest[R any, A any](method, url string, request *R) (*A, error) {
	var buf *bytes.Buffer
	var r *http.Request
	var err error
	if request != nil {
		buf = new(bytes.Buffer)
		json.NewEncoder(buf).Encode(request)
		r, err = http.NewRequest(method, url, buf)
	} else {
		r, err = http.NewRequest(method, url, nil)
	}
	if err != nil {
		return nil, err
	}
	client := http.Client{}
	res, err := client.Do(r)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	answer := new(A)
	err = json.NewDecoder(res.Body).Decode(answer)
	if err != nil {
		return nil, err
	}
	return answer, nil
}
