package interface_ptr

import (
	"context"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
)

type BrokerInterface interface {
	GetName() NameInterface
	GetID() int
	Start(context.Context)
}

type NameInterface interface {
	String() string
}

type Name string

func (me Name) String() string {
	return fmt.Sprintf("%p : %s", &me, string(me))
}

type Broker struct {
	ID   int
	Name *NameInterface
}

func (me *Broker) GetName() NameInterface {
	return NameInterface(*me.Name)
}

func (me *Broker) GetID() int {
	return me.ID
}

func (me *Broker) Start(ctx context.Context) {
	for {
		logrus.Printf("%p : %s", me.Name, *(me).Name)
		time.Sleep(time.Second * 1)
		select {
		case <-ctx.Done():
			break
		default:
		}
	}
}

func NewBroker(id int, config *BrokerConfig) BrokerInterface {
	return &Broker{
		Name: &config.Name,
		ID:   id,
	}
}
