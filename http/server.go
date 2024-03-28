package http

import (
	"fmt"
	"log"
	"net"
)

type HTTPServer struct {
	listenAddr string
	listener   net.Listener
	reqChan    chan *HTTPRequest
	quitChan   chan struct{}
}

func NewHTTPServer(listenAddr string) *HTTPServer {
	return &HTTPServer{
		listenAddr: listenAddr,
		quitChan:   make(chan struct{}),
		reqChan:    make(chan *HTTPRequest, 100),
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
	buf := make([]byte, 2048)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			return
		}
		req, err := NewHTTPRequest(buf[:n])

		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(buf[:n]))
		fmt.Println(req)
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
