package dbo

import (
	"github.com/daewu14/go-db-observer/canal"
)

// Observe : simplify binlog as database observer
func Observe(config Config, observers []Observer) {
	run(canal.Config{
		Host:     config.Host,
		User:     config.User,
		Password: config.Password,
		Flavor:   config.Flavor,
		Port:     config.Port,
	}, observers)
}
