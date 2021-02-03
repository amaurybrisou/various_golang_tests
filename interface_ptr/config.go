package interface_ptr

type BrokerConfig struct {
	Name NameInterface
}

func (src *BrokerConfig) CopyTo(dst *BrokerConfig) {
	if src.Name != nil {
		(*dst).Name = (*src).Name
	}
}

type Config struct {
	Broker *BrokerConfig
}

func (src *Config) CopyTo(dst *Config) {
	if src.Broker != nil {
		// src.Broker.CopyTo(dst.Broker)
		*dst.Broker = *src.Broker
	}
}
