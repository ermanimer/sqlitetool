package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/ermanimer/sqlitetool/pkg/linebreak"
	"github.com/ermanimer/sqlitetool/pkg/tableinfo"
	"github.com/olekukonko/tablewriter"
)

// TableInfo represents the table info command.
type TableInfo struct {
	DB    string `arg:"" name:"db" type:"path" help:"Database path."`
	Table string `arg:"" name:"table" help:"Table name."`
}

// Run represents the run method of the table info command.
func (cmd *TableInfo) Run(cw ContextWrapper) error {
	db := cmd.DB
	table := cmd.Table

	columns, err := tableinfo.Columns(cw.CTX, db, table)
	if err != nil {
		return err
	}

	fmt.Printf("SQLite Tool - Table Info%s", linebreak.DoubleLineBreak)
	fmt.Printf("Database: %s%s", db, linebreak.LineBreak)
	fmt.Printf("Table: %s%s", table, linebreak.DoubleLineBreak)
	fmt.Printf("Columns:%s", linebreak.LineBreak)

	w := tablewriter.NewWriter(os.Stdout)
	w.SetHeader([]string{"#", "Name", "Type"})
	for i, column := range columns {
		name := column[0]
		dataType := column[1]
		w.Append([]string{strconv.Itoa(i + 1), name, dataType})
	}
	w.Render()

	return nil
}
