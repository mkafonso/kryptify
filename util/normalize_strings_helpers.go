package util

import "github.com/jackc/pgx/v5/pgtype"

func GetStringValue(nullString pgtype.Text) string {
	if nullString.Valid {
		return nullString.String
	}
	return ""
}
