package canal

import (
	"fmt"

	"github.com/go-mysql-org/go-mysql/canal"
)

type canalMySQL struct {
	config *Config
}

func NewCanal(cfg *Config) Contract {
	return &canalMySQL{
		config: cfg,
	}
}

func (c *canalMySQL) GetCanal() (*canal.Canal, error) {
	cfg := canal.NewDefaultConfig()
	cfg.Addr = fmt.Sprintf("%s:%s", c.config.Host, c.config.Port)
	cfg.User = c.config.User
	cfg.Password = c.config.Password
	cfg.Flavor = c.config.Flavor
	cfg.Dump.ExecutionPath = ""
	cfg.Dump.Tables = []string{"logs"}
	cnl, err := canal.NewCanal(cfg)
	if err != nil {
		return nil, err
	}
	cnl.GetDelay()
	return cnl, nil
}
