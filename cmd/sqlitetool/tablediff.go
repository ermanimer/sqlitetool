package main

import (
	"fmt"
	"os"

	"github.com/ermanimer/sqlitetool/pkg/linebreak"
	"github.com/ermanimer/sqlitetool/pkg/tablediff"
	"github.com/ermanimer/sqlitetool/pkg/tableinfo"
	"github.com/olekukonko/tablewriter"
)

// TableDiff represents the table diff command.
type TableDiff struct {
	DBA    string `arg:"" name:"dbA" type:"path" help:"Database path A."`
	TableA string `arg:"" name:"tableA" help:"Table name A."`
	DBB    string `arg:"" name:"dbB" type:"path" help:"Database path B."`
	TableB string `arg:"" name:"tableB" help:"Table name B."`
}

// Run represents the run method of the table diff command.
func (cmd *TableDiff) Run(cw ContextWrapper) error {
	dbA := cmd.DBA
	tableA := cmd.TableA
	dbB := cmd.DBB
	tableB := cmd.TableB

	columnsA, err := tableinfo.Columns(cw.CTX, dbA, tableA)
	if err != nil {
		return err
	}

	columnsB, err := tableinfo.Columns(cw.CTX, dbB, tableB)
	if err != nil {
		return err
	}

	res := tablediff.Diff(columnsA, columnsB)
	signs := []string{"-", "~", "+"}

	fmt.Printf("SQLite Tool - Table Diff%s", linebreak.DoubleLineBreak)
	fmt.Printf("Database A: %s%s", dbA, linebreak.LineBreak)
	fmt.Printf("Table A: %s%s", tableA, linebreak.LineBreak)
	fmt.Printf("Database B: %s%s", dbB, linebreak.LineBreak)
	fmt.Printf("Table B: %s%s", tableB, linebreak.DoubleLineBreak)
	fmt.Printf("Columns:%s", linebreak.LineBreak)

	w := tablewriter.NewWriter(os.Stdout)
	w.SetHeader([]string{"Diff", "Name", "Type"})
	for i, columns := range res {
		for _, column := range columns {
			w.Append([]string{signs[i], column[0], column[1]})
		}
	}

	w.Render()

	return nil
}
