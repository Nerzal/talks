package work

type service interface {
	DoSomething() error
}

type Worker struct {
	service service
}

func NewWorker(service service) *Worker {
	return &Worker{service: service}
}

func (w *Worker) DoSomething() error {
	return w.service.DoSomething()
}

