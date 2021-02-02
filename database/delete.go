package database

type DeleteComponent struct {
	Table string
	Selection string
	SelectionArgs []string
}

func MakeDeleteStatement(component DeleteComponent) string {
	return concat(" ", []string {"DELETE",
		"FROM", component.Table,
		"WHERE", component.Selection})
}