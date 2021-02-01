package database

type DeleteComponent struct {
	table string
	selection string
	selectionArgs []string
}

func MakeDeleteStatement(component DeleteComponent) string {
	return concat(" ", []string {"DELETE",
		"FROM", component.table,
		"WHERE", component.selection})
}