package tablediff

// Diff returns columnsB diff columnsA, columnsA intersect columnsB and columnsa diff columnsB.
func Diff(columnsA [][]string, columnsB [][]string) [][][]string {
	var res [][][]string

	res = append(res, difference(columnsB, columnsA))
	res = append(res, intersection(columnsA, columnsB))
	res = append(res, difference(columnsA, columnsB))

	return res
}

func difference(columnsA [][]string, columnsB [][]string) [][]string {
	var res [][]string
	for _, columnA := range columnsA {
		if !exists(columnsB, columnA) {
			res = append(res, columnA)
		}
	}

	return res
}

func intersection(columnsA [][]string, columnsB [][]string) [][]string {
	var res [][]string
	for _, columnA := range columnsA {
		if exists(columnsB, columnA) {
			res = append(res, columnA)
		}
	}

	return res
}

func exists(columns [][]string, column []string) bool {
	for _, eColumn := range columns {
		if eColumn[0] == column[0] && eColumn[1] == column[1] {
			return true
		}
	}

	return false
}
