type Service struct {}

func NewService() *Service {
	return &service{}
}

func (s *service) DoSomething() error {
	return nil
}
