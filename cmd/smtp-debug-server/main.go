package main

import (
	"crypto/tls"
	"flag"
	"io"
	"log"
	"os"

	"github.com/emersion/go-smtp"
)

var addr = "127.0.0.1:1025"
var tlsCert = ""
var tlsKey = ""

func init() {
	flag.StringVar(&addr, "l", addr, "Listen address")
	flag.StringVar(&tlsCert, "tls-cert", tlsCert, "tls cert file path")
	flag.StringVar(&tlsKey, "tls-key", tlsKey, "tls key file path")
}

type backend struct{}

func (bkd *backend) NewSession(c *smtp.Conn) (smtp.Session, error) {
	return &session{}, nil
}

type session struct{}

func (s *session) AuthPlain(username, password string) error {
	return nil
}

func (s *session) Mail(from string, opts *smtp.MailOptions) error {
	return nil
}

func (s *session) Rcpt(to string, opts *smtp.RcptOptions) error {
	return nil
}

func (s *session) Data(r io.Reader) error {
	return nil
}

func (s *session) Reset() {}

func (s *session) Logout() error {
	return nil
}

func main() {
	flag.Parse()

	s := smtp.NewServer(&backend{})

	s.Addr = addr
	s.Domain = "localhost"
	s.AllowInsecureAuth = true
	s.Debug = os.Stdout

	if tlsCert != "" && tlsKey != "" {
		cert, err := tls.LoadX509KeyPair(tlsCert, tlsKey)
		if err != nil {
			log.Fatalf("can not load tls key %v", err)
		}
		s.TLSConfig = &tls.Config{
			Certificates: []tls.Certificate{cert},
		}
	}

	log.Println("Starting SMTP server at", addr)
	log.Fatal(s.ListenAndServe())
}
