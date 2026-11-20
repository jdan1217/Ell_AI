package traffic

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func NewRequestData(req *http.Request) (RequestData, error) {
	data := RequestData{
		Method:   req.Method,
		URL:      req.URL.String(),
		Headers:  make(map[string][]string),
		Body:     nil,
		Protocol: req.Proto,
	}
	for name, values := range req.Header {
		data.Headers[name] = append([]string(nil), values...)
	}
	body, err := io.ReadAll(req.Body)
	if err != nil {
		return RequestData{}, err
	}
	req.Body = io.NopCloser(bytes.NewReader(body))
	data.Body = body
	fmt.Printf("DEBUG: req.Proto=%q, data.Protocol=%q\n", req.Proto, data.Protocol)
	return data, nil
}
