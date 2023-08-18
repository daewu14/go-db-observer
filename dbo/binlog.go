package dbo

import (
	"github.com/daewu14/go-db-observer/canal"
	canal2 "github.com/go-mysql-org/go-mysql/canal"
	"log"
)

type binLog struct {
	cnl *canal2.Canal
}

func newBinlog(cnl *canal2.Canal) Contract {
	return &binLog{
		cnl: cnl,
	}
}

func (m *binLog) Exec(event canal2.EventHandler) {
	coords, err := m.cnl.GetMasterPos()
	if err == nil {
		m.cnl.SetEventHandler(event)
		m.cnl.RunFrom(coords)
	}
}

var bl Contract
var blInstance bool = false

func cnl(config canal.Config) canal.Contract {
	return canal.NewCanal(&config)
}

func binlog(canalConfig canal.Config) Contract {
	if blInstance == false {

		cn, err := cnl(canalConfig).GetCanal()

		if err != nil {
			log.Fatal("Error 1 :>", err)
		}
		bl = newBinlog(cn)
		blInstance = true
	}
	return bl
}
