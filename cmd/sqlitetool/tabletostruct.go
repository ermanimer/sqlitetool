package main

import (
	"fmt"

	"github.com/ermanimer/sqlitetool/pkg/linebreak"
	"github.com/ermanimer/sqlitetool/pkg/tableinfo"
	"github.com/ermanimer/sqlitetool/pkg/tabletostruct"
)

// TableToStruct represents the table to struct command.
type TableToStruct struct {
	DB     string `arg:"" name:"db" type:"path" help:"Database path."`
	Table  string `arg:"" name:"table" help:"Table name."`
	Struct string `arg:"" name:"struct" help:"Struct name."`
}

// Run represents the run method of the table to struct command.
func (cmd *TableToStruct) Run(cw ContextWrapper) error {
	db := cmd.DB
	table := cmd.Table
	st := cmd.Struct

	columns, err := tableinfo.Columns(cw.CTX, db, table)
	if err != nil {
		return err
	}

	res := tabletostruct.TableToStruct(columns, st)

	fmt.Printf("SQLite Tool - Table To Struct%s", linebreak.DoubleLineBreak)
	fmt.Printf("Database: %s%s", db, linebreak.LineBreak)
	fmt.Printf("Table: %s%s", table, linebreak.LineBreak)
	fmt.Printf("Struct:%s%s", st, linebreak.DoubleLineBreak)
	fmt.Println(res)

	return nil
}
