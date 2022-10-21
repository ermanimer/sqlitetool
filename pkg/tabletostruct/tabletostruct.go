package tabletostruct

import (
	"fmt"
	"strings"

	"github.com/ermanimer/sqlitetool/pkg/linebreak"
)

func TableToStruct(columns [][]string, name string) string {
	sb := &strings.Builder{}
	sb.WriteString("type " + name + " struct {" + linebreak.LineBreak)
	formattedColumns := formatColumns(columns)
	for _, column := range formattedColumns {
		name := column[0]
		dataType := column[1]
		sb.WriteString("\t" + name + " " + dataType + linebreak.LineBreak)
	}
	sb.WriteString("}")

	return sb.String()
}

func formatColumns(columns [][]string) [][]string {
	maxNameLen := calculateMaxNameLen(columns)

	var formattedColumns [][]string
	for _, column := range columns {
		formattedColumn := formatColumn(column, maxNameLen)
		formattedColumns = append(formattedColumns, formattedColumn)
	}

	return formattedColumns
}

func calculateMaxNameLen(columns [][]string) int {
	var maxNameLen int
	for _, column := range columns {
		nameLen := len(column[0])
		if nameLen > maxNameLen {
			maxNameLen = nameLen
		}
	}

	return maxNameLen
}

func formatColumn(column []string, maxNameLen int) []string {
	name := column[0]
	dataType := column[1]

	formattedName := formatName(name, maxNameLen)
	formattedType := formatType(dataType)

	return []string{formattedName, formattedType}
}

func formatName(name string, maxLen int) string {
	camelCaseName := toCamelCase(name)
	return fmt.Sprintf("%-*s", maxLen, camelCaseName)
}

func toCamelCase(s string) string {
	s = strings.ToLower(s)
	ss := strings.Split(s, "_")

	var res string
	for _, es := range ss {
		res += strings.Title(es)
	}

	return res
}

func formatType(dataType string) string {
	switch dataType {
	case "INTEGER":
		return "int64"
	case "TEXT":
		return "string"
	case "BLOB":
		return "interface{}"
	case "REAL":
		return "float64"
	case "DATE", "DATETIME":
		return "time.Time"
	default:
		return "UNSUPPORTED_DATA_TYPE"
	}
}
