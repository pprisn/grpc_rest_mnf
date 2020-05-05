package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/pprisn/grpc_rest_mnf/app/apiserver"
	"log"
)

var (
	configPath string
)

func init() {

	flag.StringVar(&configPath, "config-path", "configs/cfg.toml", "path to config file")
}

func main() {
	flag.Parse()

	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	s := apiserver.New(config)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
	//	if err := apiserver.Start(config); err != nil {
	//		log.Fatal(err)
	//	}

}
