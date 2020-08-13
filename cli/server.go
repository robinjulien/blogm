package cli

import srv "github.com/robinjulien/blogm/server"

func server(args *[]string) {
	if len(*args) == 3 && (*args)[2] == "start" {
		srv.Start()
	} else {
		help()
	}
}
