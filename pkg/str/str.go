package str

import "github.com/iancoleman/strcase"

// Snake => chance_fyi
func Snake(s string) string {
	return strcase.ToSnake(s)
}

// Camel => ChanceFyi
func Camel(s string) string {
	return strcase.ToCamel(s)
}

// LowerCamel => chanceFyi
func LowerCamel(s string) string {
	return strcase.ToLowerCamel(s)
}
