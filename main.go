package main

import (

	// "go.uber.org/zap"

	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/obed/demo-project/userAuth/dig"
	"github.com/obed/demo-project/userAuth/server"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(-1)
	}
}

func run() error {
	g := gin.Default()
	d := dig.BuildProject()
	svr := server.NewServer(g, d)
	svr.MapRoutes()
	if err := svr.SetupDB(); err != nil {
		return err
	}
	return svr.Start()
}
