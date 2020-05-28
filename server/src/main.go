package main

import (
	"io/ioutil"
	"log"

	"github.com/team-gleam/kiwi-basket/server/src/infra/db/handler"
	"github.com/team-gleam/kiwi-basket/server/src/infra/router"
	"gopkg.in/yaml.v3"
)

func main() {
	b, err := ioutil.ReadFile("./config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	var c handler.Config

	err = yaml.Unmarshal(b, &c)
	if err != nil {
		log.Fatal(err)
	}

	router.Run(c)
}
