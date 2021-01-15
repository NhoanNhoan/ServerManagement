 package database

import (
	_"errors"
	_"strings"
	"database/sql_clause"
)

type SQLStatement interface {
	InvalidParameters() error
	Generate() (string, error)
}

type SQLGeneratorUtil struct {

}

type SQLGenerator struct {
	statement *SQLStatement
}

func (generator SQLGenerator) GenerateQuery(
									tableName string,
									columns []string,
									selection string,
									selectionsArgs []string,
									groupBy string,
									having string,
									orderBy string,
									limit string) (string, error) {
	generator.statement = QueryStatement{
		tableName,
		columns,
		selection,
		selectionsArgs,
		groupBy,
		having,
		orderBy,
		limit
	}

	return generator.Generate()
}

// func (generator *SQLGenerator) Query(tableName string, 
// 									columns []string, 
// 									selection string, 
// 									selectionArgs []string, 
// 									groupBy string, 
// 									having string , 
// 									orderBy string, 
// 									limit string) (string, error) {
// 	if "" == groupBy && "" != having {
// 		return "", errors.New("HAVING clauses are only permitted when using a groupBy clause")
// 	}

// 	if "" == tableName {
// 		return "", errors.New("No table name for query!")
// 	}

// 	query := "SELECT"
// 	query = query + generator.makeSelectClause(columns)
// 	query = query + "FROM " + tableName + " "

// 	if "" != selection {
// 		query = query + "WHERE " + generator.makeWhereClause(selection, selectionArgs)
// 	}

// 	if 0 != len(groupBy) {
// 		query = query + "GROUP BY " + groupBy +  " "
// 	}

// 	if 0 != len(having) {
// 		query = query + "HAVING " + having + " "
// 	}

// 	if 0 != len(orderBy) {
// 		query = query + "ORDER BY " + orderBy + " "
// 	}

// 	if 0 != len(limit) {
// 		query = query + "LIMIT " + limit
// 	}

// 	return query, nil

// }

// func (generator *SQLGenerator) makeSelectClause(columns []string) string {
// 	if nil != columns && 0 != len(columns) {
// 		return " " + strings.Join(columns[:], ", ") + " "
// 	}
// 	return " * "
// }

// func (generator *SQLGenerator) makeWhereClause(selection string, selectionArgs []string) string {
// 	return replaceDelimeter(selection, "?", selectionArgs)
// }

// func replaceDelimeter(s string, delimeter string, words []string) string {
// 	if 0 == len(words) || 0 == len(s) {
// 		return ""
// 	}

// 	words_idx := 0
// 	res := s

// 	for _, letter := range s {
// 		if string(letter) == delimeter {
// 			res = strings.Replace(res, delimeter, words[words_idx], 1)
// 			words_idx += 1
// 		}
// 	}

// 	return res
// }

// func (generator *SQLGenerator) Insert(tableName string, values map[string]string) (string, error) {
// 	columnsClause, valuesClause := generator.makeInsertionClause(values)
// 	return "INSERT INTO " + tableName + " " + columnsClause + " VALUES" + valuesClause, nil
// }

// func (generator *SQLGenerator) makeInsertionClause(values map[string]string) (string, string) {
// 	columns := make([]string, 1)
// 	val := make([]string, 1)

// 	for column, value := range values {
// 		columns = append(columns, column)
// 		val = append(val, value)
// 	}

// 	return ("(" + strings.Join(columns, ", ") + ")"), ("(" + strings.Join(val, ", ") + ")")
// }

// func (generator *SQLGenerator) Update(tableName string, 
// 									selection string, 
// 									selectionArgs []string, 
// 									values map[string] string) (string, error) {
// 	if "" == tableName {
// 		return "", errors.New("No table name for query!")
// 	}

// 	return "", nil

// }