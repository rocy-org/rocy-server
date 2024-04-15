package main

import (
	"github.com/rocy-org/rocy-server/database/conn"
	"github.com/rocy-org/rocy-server/server"
)

func initAndRunServer() {
	server.InitAndRun()
}

func main() {
	conn.ConnInit()
	initAndRunServer()
}
