package database/sql/sql_statement

import (
	"errors"
)

type QueryStatement struct {
	tableName string
	columns []string
	selection string
	selectionArgs []string
	groupBy string
	having string
	orderBy string
	limit string
}

func (statement QueryStatement) InvalidParameters() error {
	if "" == groupBy && "" != having {
		return errors.New("HAVING clauses are only permitted when using a groupBy clause")
	}

	if "" == tableName {
		return errors.New("No table name for query!")
	}

	return nil
}

func (statement QueryStatement) Generate() (string, error) {
	invalid := statement.InvalidParameters()
	if nil != invalid {
		return "", invalid
	}

	clauseGenerator := SQLClauseGenerator{}
	selectionClause := clauseGenerator.SelectionClause(statement.columns)
	fromClause := clauseGenerator.FromClause(statement.tableName)
	whereClause := clauseGenerator.WhereClause(statement.selection, statement.selectionArgs)

	query := "SELECT " + selectionClause + " FROM " + fromClause

	if "" != whereClause {
		query = query + " WHERE " + whereClause
	}

	if 0 != len(groupBy) {
		query = query + " GROUP BY " + groupBy
	}

	if 0 != len(having) {
		query = query + " HAVING " + having
	}

	if 0 != len(orderBy) {
		query = query + " ORDER BY " + orderBy
	}

	if 0 != len(limit) {
		query = query + " LIMIT " + limit
	}

	return query, nil
}