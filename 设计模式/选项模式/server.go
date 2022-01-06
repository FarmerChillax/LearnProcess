package main

const (
	defaultAddr = "127.0.0.1"
	defaultPort = 5000
)

type Server struct {
	Addr string
	Port int
}

type ServerOptions struct {
	Addr string
	Port int
}

type ServerOption interface {
	apply(*ServerOptions)
}

type FuncServerOption struct {
	f func(*ServerOptions)
}

func (fo FuncServerOption) apply(option *ServerOptions) {
	fo.f(option)
}

func WithAddr(addr string) ServerOption {
	return FuncServerOption{
		f: func(options *ServerOptions) {
			options.Addr = addr
		},
	}
}

func WithPort(port int) ServerOption {
	return FuncServerOption{
		f: func(options *ServerOptions) {
			options.Port = port
		},
	}
}

func NewServer(opts ...ServerOption) *Server {
	options := ServerOptions{
		Addr: defaultAddr,
		Port: defaultPort,
	}

	for _, opt := range opts {
		opt.apply(&options)
	}

	return &Server{
		Addr: options.Addr,
		Port: options.Port,
	}

}
