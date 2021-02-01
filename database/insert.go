package database

import (
	"strings"
)

type InsertComponent struct {
	Table string
	Columns []string
	Values [][]string
}

func MakeInsert(component InsertComponent) string {
	columnsClause := concat("", []string {"(", strings.Join(component.Columns, ", "), ")"})
	valuesClause := makeValuesClauses(component.Values)
	return concat("", []string {"INSERT INTO ", component.Table, columnsClause,
				" VALUES", valuesClause})
}

func makeValuesClauses(values [][]string) string {
	return "(" + replicateToStr("?", ", ", len(values[0])) + ")"
}