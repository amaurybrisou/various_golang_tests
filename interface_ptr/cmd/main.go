package main

import (
	"context"
	"time"

	api "github.com/amaurybrisou/various_tests/interface_ptr"
	"github.com/sirupsen/logrus"
)

func main() {
	conf := &api.Config{
		Broker: &api.BrokerConfig{
			Name: api.Name("Robinhood.com"),
		},
	}

	secondConf := &api.Config{
		Broker: &api.BrokerConfig{
			Name: api.Name("Binance.com"),
		},
	}

	manager := api.NewManager(conf)
	manager.AddBroker(api.NewBroker(1, conf.Broker))

	ctx, cancel := context.WithCancel(context.Background())

	go manager.Start(ctx, 1)

	time.Sleep(time.Second * 2)

	logrus.Println("switching conf")
	secondConf.CopyTo(conf)

	time.Sleep(time.Second * 2)

	cancel()

	logrus.Println("restarting runner")
	go manager.Start(ctx, 1)

	time.Sleep(time.Second * 2)

}
