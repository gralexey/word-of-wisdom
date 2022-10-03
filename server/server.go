package server

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"log"
	"net"

	"word-of-wisdom/pow"
)

type Server struct {
	db Db
}

func Start(tcpPort uint) {
	addr := fmt.Sprintf(":%d", tcpPort)
	server := Server{
		db: newDb(),
	}
	server.Listen(addr)
}

func (s *Server) Listen(addr string) {
	log.Printf("Starting server")

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Panic(err)
	}

	log.Printf("Listening connection %s ...", addr)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Accept error: %v", err)
			continue
		}

		go func() {
			if err := s.handle(conn); err != nil {
				log.Printf("client %s: error: %v", conn.RemoteAddr(), err)
			}
			conn.Close()
		}()
	}
}

func (s *Server) handle(conn net.Conn) (err error) {
	addr := conn.RemoteAddr().String()
	log.Printf("New request: %s", addr)

	// Just 'HELLO'
	_, err = bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		return
	}

	token := pow.NewToken()
	_, err = fmt.Fprintf(conn, "%x\n", token)
	if err != nil {
		return
	}

	proofString, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		return
	}

	proof, _ := hex.DecodeString(proofString)

	if !pow.Verify(token, proof) {
		err = fmt.Errorf("wrong proof")
		return
	}

	wisdom, err := s.db.getWisdom()
	if err != nil {
		return
	}

	_, err = fmt.Fprintf(conn, "%s (%s)\r", wisdom.Text, wisdom.Author)
	return err
}
