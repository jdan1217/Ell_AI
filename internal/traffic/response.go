package traffic

import (
	"bytes"
	"io"
	"net/http"
)

func NewResponseData(resp *http.Response) (ResponseData, error) {
	data := ResponseData{
		StatusCode: resp.StatusCode,
		Headers:    make(map[string][]string),
		Body:       nil,
		Protocol:   resp.Proto,
	}
	for name, values := range resp.Header {
		data.Headers[name] = append([]string(nil), values...)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return ResponseData{}, err
	}
	resp.Body = io.NopCloser(bytes.NewReader(body))
	data.Body = body
	return data, nil
}
