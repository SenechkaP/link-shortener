package stat

import (
	"advpractice/pkg/event"
	"context"
	"log"
)

type StatServiceDeps struct {
	EventBus       *event.EventBus
	StatRepository *StatRepository
}

type StatService struct {
	EventBus       *event.EventBus
	StatRepository *StatRepository
}

func NewStatService(deps *StatServiceDeps) *StatService {
	return &StatService{
		EventBus:       deps.EventBus,
		StatRepository: deps.StatRepository,
	}
}

func (service *StatService) AddClick(ctx context.Context) {
	for {
		select {
		case msg := <-service.EventBus.Subscribe():
			if msg.Type == event.EventLinkVisited {
				id, ok := msg.Data.(uint)
				if !ok {
					log.Fatalln("Bad EventLinkVisited data: ", msg.Data)
					continue
				}
				service.StatRepository.AddClick(id)
			}
		case <-ctx.Done():
			return
		}
	}
}
