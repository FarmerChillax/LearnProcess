package service

type HelloService struct{}

func (hs *HelloService) Hello(request string, reply *string) error {
	*reply = "Hello " + request
	return nil
}
