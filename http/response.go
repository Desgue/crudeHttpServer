package http

import "fmt"

type ResponseWriter interface {
	Header() Header
	Write([]byte) (int, error)
	WriteHeader(statusCode int)
}

type response struct {
	Version      string
	StatusCode   int
	ReasonPhrase string
	Headers      Header
	Body         []byte
}

func (r *response) Header() Header {
	return r.Headers
}

func (r *response) Write(b []byte) (int, error) {
	r.Body = append(r.Body, b...)
	return len(b), nil
}

func (r *response) WriteHeader(statusCode int) {
	r.StatusCode = statusCode
	r.ReasonPhrase = StatusToText(statusCode)

}

func NewResponse() *response {
	return &response{
		Version: "HTTP/1.1",
		Headers: make(Header),
	}
}

func (r response) encode() string {
	var statusLine string
	statusLine += fmt.Sprintf("%s %d %s", r.Version, r.StatusCode, r.ReasonPhrase)
	headers := r.Headers.encode()
	return fmt.Sprintf(statusLine + "\r\n" + headers + "\r\n\r\n" + string(r.Body))
}
