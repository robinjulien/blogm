package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/robinjulien/blogm/entities"
)

var cfg entities.Config

// Start starts the http server
func Start() {
	fmt.Println("Starting the server...")

	bytescfg, err := ioutil.ReadFile("config.json") // get the config into the global variable cfg
	check(err)
	err = json.Unmarshal(bytescfg, &cfg)
	check(err)

	mux := new(BlogmHandler) // create new instance of BlogmHandler, that serve HTTP Request as a multiplexer

	srv := &http.Server{
		Addr:    cfg.Host + ":" + cfg.Port,
		Handler: mux,
	}
	srv.ListenAndServe()
	// TODO Add support for TLS
}

// check is used to check if there is an error that should not happen, and thus that there is no recovery from that error
func check(err error) {
	if err != nil {
		panic(err)
	}
}
