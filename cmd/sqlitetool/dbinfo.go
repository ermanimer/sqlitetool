package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/ermanimer/sqlitetool/pkg/dbinfo"
	"github.com/ermanimer/sqlitetool/pkg/linebreak"
	"github.com/olekukonko/tablewriter"
)

// DBInfo represents the database info command.
type DBInfo struct {
	DB string `arg:"" name:"db" type:"path" help:"Database path."`
}

// Run represents the run method of the database info command.
func (cmd *DBInfo) Run(cw ContextWrapper) error {
	db := cmd.DB

	tables, err := dbinfo.Tables(cw.CTX, db)
	if err != nil {
		return err
	}

	fmt.Printf("SQLite Tool - Database Info%s", linebreak.DoubleLineBreak)
	fmt.Printf("Database: %s%s", db, linebreak.DoubleLineBreak)
	fmt.Printf("Tables:%s", linebreak.LineBreak)

	w := tablewriter.NewWriter(os.Stdout)
	w.SetHeader([]string{"#", "Name"})
	for i, table := range tables {
		w.Append([]string{strconv.Itoa(i + 1), table})
	}
	w.Render()

	return nil
}
