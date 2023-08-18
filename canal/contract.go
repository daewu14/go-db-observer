package canal

import "github.com/go-mysql-org/go-mysql/canal"

type Config struct {
	Host     string
	User     string
	Password string
	Flavor   string
	Port     string
}
type Contract interface {
	GetCanal() (*canal.Canal, error)
}
