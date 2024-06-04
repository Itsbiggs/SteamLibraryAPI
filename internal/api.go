package internal

type Server struct {
	addr string
}

func NewAPIServer(addr string) (*Server, error) {