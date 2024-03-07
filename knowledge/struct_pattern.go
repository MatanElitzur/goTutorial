package knowledge

import "fmt"

type OptFunc func(*Opts)

type Opts struct {
	maxConn int
	id      string
	tls     bool
}

type Server struct {
	Opts
}

func withID(id string) OptFunc {
	return func(opts *Opts) {
		opts.id = id
	}
}

func withTLS(opts *Opts) {
	opts.tls = true
}

func withMaxConn(n int) OptFunc {
	return func(opts *Opts) {
		opts.maxConn = n
	}
}

// The newServer func args are an Ellipsis slice of type OptFunc
// The function that use ... is called variadic function
// Or in other words, a user is allowed to pass zero or more arguments in the variadic function
// https://www.geeksforgeeks.org/how-to-use-ellipsis-in-golang/
func newServer(opts ...OptFunc) *Server {
	o := defaultOpts()
	for _, fn := range opts {
		fn(&o)
	}
	return &Server{
		Opts: o,
	}
}

func defaultOpts() Opts {
	return Opts{
		maxConn: 10,
		id:      "default",
		tls:     false,
	}
}

func StructPattern() {
	s := newServer()
	fmt.Printf("%+v\n", s)
	s = newServer(withTLS)
	fmt.Printf("%+v\n", s)
	s = newServer(withTLS, withID("ThisIsMyID"))
	fmt.Printf("%+v\n", s)
	s = newServer(withTLS, withID("ThisNOTMyID"), withMaxConn(99))
	fmt.Printf("%+v\n", s)
}
