package main

import (
	"crud_demo/app/router"
	"crud_demo/migrates"
	_ "crud_demo/migrates"
)

func main() {
	// set router
	router.SetRouter()

	//for migrating
	migrates.MigrateTableUser()

	// start server
	router.Server.Logger.Fatal(router.Server.Start(":9000"))
}
