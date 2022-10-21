package dbdiff

// Diff returns tablesB diff tablesB, tablesA intersect tablesB and tablesA diff tablesB.
func Diff(tablesA []string, tablesB []string) [][]string {
	var res [][]string

	res = append(res, difference(tablesB, tablesA))
	res = append(res, intersection(tablesA, tablesB))
	res = append(res, difference(tablesA, tablesB))

	return res
}

func difference(tablesA []string, tablesB []string) []string {
	var res []string
	for _, tableA := range tablesA {
		if !exists(tablesB, tableA) {
			res = append(res, tableA)
		}
	}

	return res
}

func intersection(tablesA []string, tablesB []string) []string {
	var res []string
	for _, tableA := range tablesA {
		if exists(tablesB, tableA) {
			res = append(res, tableA)
		}
	}

	return res
}

func exists(tables []string, table string) bool {
	for _, eTable := range tables {
		if eTable == table {
			return true
		}
	}

	return false
}
