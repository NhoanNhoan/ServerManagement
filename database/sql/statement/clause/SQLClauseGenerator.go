package database

import (
	"strings"
)

const DELIMETER = "?"

type SQLClauseUtil struct {

}

func (util *SQLClauseUtil) ReplaceDelimeterByValues(s string, 
													delimeter string, 
													values []string) string {
	if nil == values || 0 == len(values) || "" == s {
		return ""
	}

	res := s
	valuesIdx := 0
	for _, letter := range s {
		if string(letter) == delimeter {
			res = strings.Replace(res, delimeter, values[valuesIdx], 1)
			valuesIdx += 1
		}
	}

	return res
}

type SQLClauseGenerator struct {

}

func (generator *SQLClauseGenerator) SelectionClause(columns []string) string {
	if nil == columns || 0 == len(columns) {
		return "*"
	}

	return strings.Join(columns, ", ")
}

func (generator *SQLClauseGenerator) FromClause(tableName string) string {
	return tableName
}

func (generator *SQLClauseGenerator) WhereClause(whereClause string, whereArgs []string) string {
	if "" == whereClause || nil == whereArgs || 0 == len(whereArgs) {
		return ""
	}

	return (&SQLClauseUtil{}).ReplaceDelimeterByValues(whereClause, DELIMETER, whereArgs)
}