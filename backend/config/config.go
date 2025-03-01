package config

import "os"

type HTTPServerConfig struct {
	Network *string `env:"NETWORK,required"`
	Addr    *string `env:"ADDR,required"`
}

func (httpConfig *HTTPServerConfig) Load() {
  network := os.Getenv("NETWORK")
	httpConfig.Network = &network
  addr := os.Getenv("ADDR")
	httpConfig.Addr = &addr
}
