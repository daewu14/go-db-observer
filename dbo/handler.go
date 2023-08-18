package dbo

import (
	"context"
	"encoding/json"
	"github.com/go-mysql-org/go-mysql/canal"
)

func handler(ctx context.Context, cnc context.CancelFunc, action string, e *canal.RowsEvent, observers []Observer) {
	for _, observer := range observers {
		if observer.Table() == e.Table.Name {
			switch action {
			case canal.UpdateAction:
				u := actionUpdate(e)
				observer.OnUpdated(ctx, cnc, u)
				break
			case canal.InsertAction:
				i := actionInsert(e)
				observer.OnInserted(ctx, cnc, i)
				break
			case canal.DeleteAction:
				d := actionDelete(e)
				observer.OnDeleted(ctx, cnc, d)
				break
			}
		}
	}
}

// actionDelete : listen deleted data
func actionDelete(e *canal.RowsEvent) AffectedTable {
	deletedDataSlice := e.Rows[0]

	firstIndex := 0
	firstColumn := e.Table.Columns[firstIndex].Name
	firstColumnVal := deletedDataSlice[firstIndex]

	affected := AffectedTable{
		Table: e.Table.Name,
		OnFirstColumn: struct {
			Key string
			Val any
		}{Key: firstColumn, Val: firstColumnVal},
	}

	for i, column := range e.Table.Columns {
		affected.Columns = append(affected.Columns, AffectedColumn{
			Name:     column.Name,
			OldValue: deletedDataSlice[i],
		})
	}

	return affected
}

// actionInsert : listen inserted data
func actionInsert(e *canal.RowsEvent) AffectedTable {
	newDataSlice := e.Rows[0]

	firstIndex := 0
	firstColumn := e.Table.Columns[firstIndex].Name
	firstColumnVal := newDataSlice[firstIndex]

	affected := AffectedTable{
		Table: e.Table.Name,
		OnFirstColumn: struct {
			Key string
			Val any
		}{Key: firstColumn, Val: firstColumnVal},
	}

	for i, column := range e.Table.Columns {
		affected.Columns = append(affected.Columns, AffectedColumn{
			Name:     column.Name,
			NewValue: newDataSlice[i],
		})
	}

	return affected
}

// actionUpdate : listen updated data
func actionUpdate(e *canal.RowsEvent) AffectedTable {
	oldDataSlice := e.Rows[0]
	newDataSlice := e.Rows[1]

	firstIndex := 0
	firstColumn := e.Table.Columns[firstIndex].Name
	firstColumnVal := oldDataSlice[firstIndex]

	affected := AffectedTable{
		Table: e.Table.Name,
		OnFirstColumn: struct {
			Key string
			Val any
		}{Key: firstColumn, Val: firstColumnVal},
	}

	for i, column := range e.Table.Columns {
		marshal1, _ := json.Marshal(oldDataSlice[i])
		marshal2, _ := json.Marshal(newDataSlice[i])
		oldVal := string(marshal1)
		newVal := string(marshal2)
		if oldVal != newVal {
			affected.Columns = append(affected.Columns, AffectedColumn{
				Name:     column.Name,
				OldValue: oldDataSlice[i],
				NewValue: newDataSlice[i],
			})
		}
	}

	return affected
}
