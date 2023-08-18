package dbo

import (
	"context"
	"github.com/go-mysql-org/go-mysql/canal"
)

type Contract interface {
	Exec(event canal.EventHandler)
}

type Observer interface {
	Table() string
	OnInserted(ctx context.Context, cnc context.CancelFunc, affected AffectedTable)
	OnUpdated(ctx context.Context, cnc context.CancelFunc, affected AffectedTable)
	OnDeleted(ctx context.Context, cnc context.CancelFunc, affected AffectedTable)
}
