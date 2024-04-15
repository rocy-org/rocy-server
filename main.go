package main

import (
	_ "github.com/rocy-org/rocy-server/utils/db"
	_ "github.com/rocy-org/rocy-server/utils/init"
	"github.com/rocy-org/rocy-server/server"
)

func initAndRunServer() {
	server.InitAndRun()
}

func main() {
	initAndRunServer()
}
