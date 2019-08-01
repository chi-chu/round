package round

type Server struct {
    Name			string
    Weight			int
    CurrentWeight	int
    IP              string
    Port            int
    Alive           bool
}

func NewServer() *Server{
    return &Server{Alive:true}
}
