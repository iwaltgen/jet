//
// Code generated by go-jet DO NOT EDIT.
// Generated at Thursday, 08-Aug-19 16:59:58 CEST
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

import (
	"time"
)

type FilmCategory struct {
	FilmID     int16 `sql:"primary_key"`
	CategoryID int16 `sql:"primary_key"`
	LastUpdate time.Time
}
