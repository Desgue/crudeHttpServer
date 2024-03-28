package http

import (
	"bytes"
	"errors"
	"fmt"
)

type HTTPRequest struct {
	Method   []byte
	URI      []byte
	Protocol []byte
	Headers  map[string]string
	Body     []byte
}

func (s HTTPRequest) String() string {
	return fmt.Sprintf("Method: %s\nURI: %s\nProtocol: %s\nHeaders: %v\nBody: %s\n", s.Method, s.URI, s.Protocol, s.Headers, s.Body)
}

func NewHTTPRequest(data []byte) (*HTTPRequest, error) {
	req := &HTTPRequest{
		Headers: make(map[string]string),
	}
	req.decode(data)
	return req, nil
}

func (r *HTTPRequest) decode(data []byte) error {
	// Parse Request Line
	requestLine, headersAndBody, found := bytes.Cut(data, []byte("\r\n"))
	if !found {
		return errors.New("invalid request, missing request line")
	}
	rl := bytes.Split(requestLine, []byte(" "))
	r.Method = rl[0]
	r.URI = rl[1]
	r.Protocol = rl[2]

	// Parse headers
	headers, body, found := bytes.Cut(headersAndBody, []byte("\r\n\r\n"))
	if !found {
		return errors.New("invalid request, missing headers")
	}
	for _, header := range bytes.Split(headers, []byte("\r\n")) {
		key, value, found := bytes.Cut(header, []byte(":"))
		if !found {
			return errors.New("invalid header format")
		}
		r.Headers[string(key)] = string(value)
	}

	// Parse Body
	r.Body = body
	return nil

}
