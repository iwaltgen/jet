//
// Code generated by go-jet DO NOT EDIT.
// Generated at Thursday, 08-Aug-19 16:59:58 CEST
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package enum

import "github.com/go-jet/jet/postgres"

var MpaaRating = &struct {
	G    postgres.StringExpression
	Pg   postgres.StringExpression
	Pg13 postgres.StringExpression
	R    postgres.StringExpression
	Nc17 postgres.StringExpression
}{
	G:    postgres.NewEnumValue("G"),
	Pg:   postgres.NewEnumValue("PG"),
	Pg13: postgres.NewEnumValue("PG-13"),
	R:    postgres.NewEnumValue("R"),
	Nc17: postgres.NewEnumValue("NC-17"),
}
