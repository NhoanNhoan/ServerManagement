package database

type QueryComponent struct 	{
	Tables []string
	Columns []string
	Selection string
	SelectionArgs []string
	GroupBy string
	Having string
	OrderBy string
	Limit string
}

func MakeQuery(component QueryComponent) string {
	fromClause := concat(", ", component.Tables)
	ColumnsClause := concat(", ", component.Columns)
	whereClause := makeWhereClause(&component)

	return concat(" ", 
				[]string{"SELECT", ColumnsClause, 
				"FROM", fromClause, 
				whereClause})
}

func makeWhereClause(component *QueryComponent) (clause string) {

	if "" != component.Selection {
		return  "WHERE " + component.Selection
	}

	return clause
}

func GetQueryStatement(comp QueryComponent) string {
	return MakeQuery(comp)
}