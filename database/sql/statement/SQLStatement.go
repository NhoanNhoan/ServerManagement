package database/sql/sql_statement

import (
	"errors"
)

type SQLStatement interface {
	InvalidParameters() error
	Generate() (string, error)
}