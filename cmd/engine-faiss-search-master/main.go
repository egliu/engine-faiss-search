package main

import (
	"engine-faiss-search/version"
	"flag"
	"fmt"

	log "github.com/sirupsen/logrus"
)

var (
	configFile   = flag.String("config", "examples/manager.json", "manager's config file")
	grpcEndpoint = flag.String("grpc_endpoint", ":8091", "grpc endpoint")
	httpEndpoint = flag.String("http_endpoint", ":8081", "http endpoint")
	managerID    = flag.String("manager_id", "", "manager name or ip address")
	printVersion = flag.Bool("version", false, "print version of this build")
	verbose      = flag.Bool("verbose", false, "verbose output")
)

func main() {
	flag.Parse()
	if *printVersion {
		version.PrintFullVersionInfo()
		return
	}
	fmt.Println("engine-faiss-search-manager")

	if *verbose {
		log.SetLevel(log.DebugLevel)
	}
}
