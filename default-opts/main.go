package main

import "fmt"

type Opts struct {
	id   string
	host string
	tls  bool
	ssl  bool
}

type OptFunc func(*Opts)

func defaultOpts() Opts {
	return Opts{
		host: "localhost",
		id:   "default",
		tls:  false,
		ssl:  false,
	}
}

type Server struct {
	Opts
}

func newServer(opts ...OptFunc) *Server {
	o := defaultOpts()
	for _, fn := range opts {
		fn(&o)
	}

	return &Server{
		Opts: o,
	}
}

func withTLS(opts *Opts) {
	opts.tls = true
}

func withSSL(opts *Opts) {
	opts.ssl = true
}

func main() {
	s := newServer(
		withTLS,
		withSSL,
	)
	fmt.Println(s)
}
