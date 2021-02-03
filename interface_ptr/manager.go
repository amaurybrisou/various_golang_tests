package interface_ptr

import (
	"context"
)

type Manager struct {
	conf    *Config
	runners map[int]BrokerInterface
}

func (me *Manager) AddBroker(r BrokerInterface) {
	if me.runners[r.GetID()] == nil {
		me.runners[r.GetID()] = r
	}
}

func (me *Manager) Start(ctx context.Context, id int) {
	if me.runners[id] == nil {
		return
	}

	go me.runners[id].Start(ctx)

	<-ctx.Done()

}

func NewManager(conf *Config) Manager {
	return Manager{
		conf:    conf,
		runners: make(map[int]BrokerInterface),
	}
}
