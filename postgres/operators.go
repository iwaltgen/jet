package postgres

import "github.com/go-jet/jet/internal/jet"

// --------- Arithmetic operators -------------//

//var MINUSi = jet.MINUSi
var MINUSf = jet.MINUSf

//----------- Logical operators ---------------//

var NOT = jet.NOT
var BIT_NOT = jet.BIT_NOT

func MINUSi(intExp IntegerExpression) IntegerExpression {
	if intLit, ok := intExp.(jet.LiteralExpression); ok {
		intLit.SetConstant(true)
	}

	return intExp
}