package dbo

type AffectedTable struct {
	Table         string
	OnFirstColumn struct {
		Key string
		Val any
	}
	Columns []AffectedColumn
}

type AffectedColumn struct {
	Name     string
	OldValue any
	NewValue any
}

type Config struct {
	Host     string
	User     string
	Password string
	Flavor   string
	Port     string
}
