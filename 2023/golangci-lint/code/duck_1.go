type Service interface {
	DoSomething() error
}

type service struct {}

func NewService() Service {
	return &service{}
}

func (s *service) DoSomething() error {
	return nil
}