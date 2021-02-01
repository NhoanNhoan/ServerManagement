package database

type UpdateComponent struct {
	Table         string
	SetClause     string
	Values        []string
	Selection     string
	SelectionArgs []string
}

func MakeUpdateStatement(component UpdateComponent) string {
	return concat(" ", []string{"UPDATE", component.Table,
		"SET", component.SetClause,
		"WHERE", component.Selection})
}

func GetUpdateStatement(comp UpdateComponent) string {
	return MakeUpdateStatement(comp)
}