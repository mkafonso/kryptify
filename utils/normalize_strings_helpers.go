package utils

import "database/sql"

func GetStringValue(nullString sql.NullString) string {
	if nullString.Valid {
		return nullString.String
	}
	return ""
}
