/*
modification history
--------------------
2023/06/30, create
*/

// Package main is special.  It defines a
// standalone executable program, not a library.
// Within package main the function main is also
// special-it's where execution of the program begins.
// Whatever main does is what the program does.
package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/NKztq/go-web-api/config"
	"github.com/NKztq/go-web-api/router"
)

// flags
var (
	localConfPath = flag.String("c", "../conf/server.yaml",
		"server local config path, default ../conf/server.yaml")
)

func main() {
	flag.Parse()

	localConf, err := config.Load(*localConfPath)
	if err != nil {
		fmt.Printf("config.Load(): %s\n", err.Error())
		os.Exit(-1)
	}

	router, err := router.InitRouter(localConf)
	if err != nil {
		fmt.Printf("router.InitRouter(): %s\n", err.Error())
		os.Exit(-1)
	}

	// pprof
	go func() {
		fmt.Printf("open pprof server error: %v\n",
			http.ListenAndServe(fmt.Sprintf("localhost:%d", localConf.Server.PprofPort), nil))
	}()

	router.Run(fmt.Sprintf(":%d", localConf.Server.ServicePort))
}
