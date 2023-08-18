package dbo

import (
	"context"
	"fmt"
	"github.com/daewu14/go-db-observer/canal"
	cn "github.com/go-mysql-org/go-mysql/canal"
	"os"
	"os/signal"
	"runtime/debug"
	"syscall"
)

func run(canalConfig canal.Config, observers []Observer) {
	binlog := binlog(canalConfig)
	binlog.Exec(&binlogHandler{observers: observers})
}

type binlogHandler struct {
	observers []Observer
	binLog    Contract
	cnc       context.CancelFunc
	cn.DummyEventHandler
}

func (l binlogHandler) OnRow(e *cn.RowsEvent) error {
	defer func() {
		if r := recover(); r != nil {
			fmt.Print(r, " ", string(debug.Stack()))
		}
	}()

	var ctx context.Context
	ctx, l.cnc = context.WithCancel(context.Background())
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-quit
		l.cnc()
	}()
	var n = 0
	var k = 1

	if e.Action == cn.UpdateAction {
		n = 1
		k = 2
	}

	for i := n; i < len(e.Rows); i += k {
		handler(ctx, l.cnc, e.Action, e, l.observers)
	}

	return nil
}
