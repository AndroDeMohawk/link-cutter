package stat

import (
	"log"

	"github.com/AndroDeMohawk/link-cutter/pkg/event"
)

type ServiceDeps struct {
	EventBus   *event.EventBus
	Repository *Repository
}

type Service struct {
	EventBus   *event.EventBus
	Repository *Repository
}

func NewService(deps *ServiceDeps) *Service {
	return &Service{
		EventBus:   event.NewEventBus(),
		Repository: &Repository{},
	}
}

func (s *Service) AddClick() {
	for msg := range s.EventBus.Subscribe() {
		if msg.Type == event.LinkVisited {
			id, ok := msg.Data.(uint)
			if !ok {
				log.Printf("Bad eventLinkVisited data")
				continue
			}
			s.Repository.AddClick(id)
		}
	}
}
