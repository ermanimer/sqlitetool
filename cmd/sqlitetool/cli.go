package main

// CLI represents the command line interface.
type CLI struct {
	DBInfo        DBInfo        `cmd:"" help:"Print the database info."`
	TableInfo     TableInfo     `cmd:"" help:"Print the table info."`
	DBDiff        DBDiff        `cmd:"" help:"Print the database diff."`
	TableDiff     TableDiff     `cmd:"" help:"Print the table diff."`
	TableToStruct TableToStruct `cmd:"" help:"Convert the table to a Go struct."`
}
