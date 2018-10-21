package main

import (
	"log"

	"github.com/amivin123/ctor/opts"
	"github.com/amivin123/ctor/server"
)

var VERSION = "1.0"

func main() {
	s := server.Server{
		Title:      "Tores",
		Port:       3000,
		ConfigPath: "config.json",
	}

	o := opts.New(&s)
	o.Version(VERSION)
	o.PkgRepo()
	o.LineWidth = 96
	o.Parse()

	if err := s.Run(VERSION); err != nil {
		log.Fatal(err)
	}
}
