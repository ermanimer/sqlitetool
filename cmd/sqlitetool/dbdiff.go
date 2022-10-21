package main

import (
	"fmt"
	"os"

	"github.com/ermanimer/sqlitetool/pkg/dbdiff"
	"github.com/ermanimer/sqlitetool/pkg/dbinfo"
	"github.com/ermanimer/sqlitetool/pkg/linebreak"
	"github.com/olekukonko/tablewriter"
)

// DBDiff represents the database diff command.
type DBDiff struct {
	DBA string `arg:"" name:"dbA" type:"path" help:"Database path A."`
	DBB string `arg:"" name:"dbB" type:"path" help:"Database path B."`
}

// Run represents the run method of the database diff command.
func (cmd *DBDiff) Run(cw ContextWrapper) error {
	dbA := cmd.DBA
	dbB := cmd.DBB

	tablesA, err := dbinfo.Tables(cw.CTX, dbA)
	if err != nil {
		return err
	}

	tablesB, err := dbinfo.Tables(cw.CTX, dbB)
	if err != nil {
		return err
	}

	res := dbdiff.Diff(tablesA, tablesB)
	signs := []string{"-", "~", "+"}

	fmt.Printf("SQLite Tool - Database Diff%s", linebreak.DoubleLineBreak)
	fmt.Printf("Database A: %s%s", dbA, linebreak.LineBreak)
	fmt.Printf("Database B: %s%s", dbB, linebreak.DoubleLineBreak)
	fmt.Printf("Tables:%s", linebreak.LineBreak)

	w := tablewriter.NewWriter(os.Stdout)
	w.SetHeader([]string{"Diff", "Name"})
	for i, tables := range res {
		for _, table := range tables {
			w.Append([]string{signs[i], table})
		}
	}

	w.Render()

	return nil
}
