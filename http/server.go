package http

import (
	"fmt"
	"log"
	"net"
)

var defaultServeMux = &ServeMux{handlers: make(map[string]HandlerFunc)}

type HTTPServer struct {
	listenAddr string
	listener   net.Listener
	Mux        *ServeMux
	reqChan    chan *Request
	quitChan   chan struct{}
}

func NewHTTPServer(listenAddr string, handler *ServeMux) *HTTPServer {
	if handler == nil {
		handler = defaultServeMux
	}
	return &HTTPServer{
		listenAddr: listenAddr,
		quitChan:   make(chan struct{}),
		reqChan:    make(chan *Request, 100),
		Mux:        handler,
	}
}

func (s *HTTPServer) ListenAndServe() error {
	var err error
	s.listener, err = net.Listen("tcp", s.listenAddr)
	if err != nil {
		return err
	}
	fmt.Printf("Listening on %s\n", s.listener.Addr().String())
	defer s.listener.Close()
	go s.acceptLoop()

	<-s.quitChan
	close(s.reqChan)

	return nil
}

func (s *HTTPServer) handleConnection(conn net.Conn) {
	defer conn.Close()
	log.Println("connection accepted from: ", conn.RemoteAddr().String())
	// Read the request from the client
	buf := make([]byte, 2048)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			return
		}
		// Process the request from the client
		req, err := NewRequest(buf[:n])
		if err != nil {
			fmt.Println(err)
			return
		}

		// Read method from request, find handler and call it
		handler, ok := s.Mux.handlers[string(req.URI)]
		if !ok {
			log.Println("No handler found for ", string(req.URI))
			return
		}
		res := NewResponse()

		handler(res, req)

		conn.Write([]byte(res.encode()))
		s.reqChan <- req
	}

}

func (s HTTPServer) acceptLoop() {
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			log.Println(err)
			return
		}
		go s.handleConnection(conn)
	}
}

type HandlerFunc func(ResponseWriter, *Request)

func HandleFunc(pattern string, handler HandlerFunc) {
	defaultServeMux.register(pattern, handler)
}

type ServeMux struct {
	handlers map[string]HandlerFunc
}

func (m *ServeMux) register(pattern string, handler HandlerFunc) {
	m.handlers[pattern] = handler
}
