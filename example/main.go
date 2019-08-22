package main

import (
	"github.com/sereiner/parrot/parrot"
)

func main() {

	app := parrot.NewApp(
		parrot.WithPlatName("apiserver"),
		parrot.WithSystemName("apiserver"),
		parrot.WithServerTypes("api"),
		parrot.WithDebug())
	app.Micro("/order/query", NewQueryHandler)
	app.Start()
}
