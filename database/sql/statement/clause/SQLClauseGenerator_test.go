package database

import (
	"testing"
)

var generator SQLClauseGenerator = SQLClauseGenerator{}

func TestSQLUtil(t *testing.T) {
	util := SQLClauseUtil{}
	testCases := map[string]string {
		util.ReplaceDelimeterByValues("", DELIMETER, []string{"one", "two", "three"}) : "",
		util.ReplaceDelimeterByValues("? ? ?", DELIMETER, []string{}) : "",
		util.ReplaceDelimeterByValues("? ? ?", DELIMETER, []string{"one", "two", "three"}) : "one two three",
		util.ReplaceDelimeterByValues("?-?-?", DELIMETER, []string{"one", "two", "three"}) : "one-two-three",
		util.ReplaceDelimeterByValues("? AND ? AND ?", DELIMETER, []string{"one", "two", "three"}) : "one AND two AND three",
		util.ReplaceDelimeterByValues("name like ? AND age = ?  NOT address = ?", DELIMETER, []string{"Washington", "31", "USA"}) : "name like Washington AND age = 31  NOT address = USA",
		util.ReplaceDelimeterByValues("name like ? AND age = ?  OR address = ?", DELIMETER, []string{"Washington", "31", "USA"}) : "name like Washington AND age = 31  OR address = USA",
	}

	for value, res := range testCases {
		if value != res {
			t.Error(value, " should equals ", res)
		}
	}
}

func TestSelectionClause(t *testing.T) {
	testCases := map[string]string {
		generator.SelectionClause([]string{}): "*",
		generator.SelectionClause([]string{"id", "name"}): "id, name",
		generator.SelectionClause([]string{"value", "abc", "xyz","asc", "tfs"}): "value, abc, xyz, asc, tfs",
		generator.SelectionClause([]string{"room"}): "room",
	}

	for value, res := range testCases {
		if value != res {
			t.Error(value, " should equals ", res)
		}
	}
}

func TestWhereClause(t *testing.T) {
	testCases := map[string]string {
		generator.WhereClause("", []string{"abc", "12", "f1s"}): "",
		generator.WhereClause("id = ?", []string{}): "",
		generator.WhereClause("id = ?", []string{"1"}): "id = 1",
		generator.WhereClause("id = ? AND name like ?", []string{"1", "Washington"}): "id = 1 AND name like Washington",
		generator.WhereClause("id = ? AND name = ? OR address = ?", []string{"1", "Washington", "USA"}): "id = 1 AND name = Washington OR address = USA",
	}

	for value, res := range testCases {
		if value != res {
			t.Error(value, " should equals ", res)
		}
	}
}