package example

import (
	"github.com/daewu14/go-db-observer/dbo"
	"github.com/daewu14/go-db-observer/example/observer"
)

func RunApp() {

	dbo.Observe(dbo.Config{
		Host:     "localhost",
		User:     "root",
		Password: "password",
		Flavor:   "mysql",
		Port:     "3306",
	}, []dbo.Observer{
		// register struct implemented dbo.Observer
		observer.User{},
	})

}
