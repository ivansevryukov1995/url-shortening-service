package stat

import (
	"log/slog"

	"github.com/ivansevryukov1995/url-shortening-service/pkg/di"
	"github.com/ivansevryukov1995/url-shortening-service/pkg/event"
)

type StatService struct {
	EventBus       *event.EventBus
	StatRepository di.IStatRepository
}

type StatServiceDeps struct {
	EventBus       *event.EventBus
	StatRepository di.IStatRepository
}

func NewStatService(deps StatServiceDeps) *StatService {
	return &StatService{
		EventBus:       deps.EventBus,
		StatRepository: deps.StatRepository,
	}
}

func (s *StatService) AddClick() {
	for msg := range s.EventBus.Subscribe() {
		if msg.Type == event.EventLinkVisited {
			id, ok := msg.Data.(uint)
			if !ok {
				slog.Error("Bad EventLinkVisited Data:", msg.Data)
			}
			s.StatRepository.AddClick(id)
		}
	}
}
