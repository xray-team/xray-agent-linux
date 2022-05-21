package service

type Service struct {
	statGetter Module
	statSender Module
}

func New(getter, sender Module) *Service {
	return &Service{
		statGetter: getter,
		statSender: sender,
	}
}

func (s *Service) Start() {
	RunModules(s.statGetter, s.statSender)
}

func (s *Service) DryRun() {
	DryRunModules(s.statGetter, s.statSender)
}
