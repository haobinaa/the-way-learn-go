package code_block

type Server struct {
	Addr     string
	Port     int
	Protocol string
	Timeout  int
	MaxConns int
	TLS      string
}

type Option func(*Server)
