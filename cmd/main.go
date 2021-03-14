package main

import (
	"flag"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
	"microblog/pkg"
	"microblog/postgres"
	"microblog/server"
	"microblog/server/handlers"
	"microblog/server/service"
	"microblog/types"
	"os"
)

func main() {
	configPath := new(string)
	flag.StringVar(configPath, "configs-path", "configs/configs-host.yaml", "path to yaml configs file")
	flag.Parse()
	f, err := os.Open(*configPath)
	if err != nil {
		pkg.FatalError(errors.Wrap(err, "Err with os.Open"))
	}
	cnf := &types.Config{}
	if err = yaml.NewDecoder(f).Decode(&cnf); err != nil {
		pkg.FatalError(errors.Wrap(err, "Err with yaml.NewDecoder"))
	}
	pg, err := postgres.NewSQL(cnf.PsqlInfo)
	if err != nil {
		pkg.FatalError(errors.Wrap(err, "Err with postgres.NewSQL"))
	}
	srv, err := service.NewService(pg)
	if err != nil {
		pkg.FatalError(errors.Wrap(err, "Err with service.NewService"))
	}
	handl, err := handlers.NewHandler(srv)
	if err != nil {
		pkg.FatalError(errors.Wrap(err, "Err with handlers.NewHandle"))
	}
	server.StartServer(cnf.ServerPort, handl)
}
