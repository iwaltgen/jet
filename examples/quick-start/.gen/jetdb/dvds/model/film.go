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

type Film struct {
	FilmID          int32 `sql:"primary_key"`
	Title           string
	Description     *string
	ReleaseYear     *int32
	LanguageID      int16
	RentalDuration  int16
	RentalRate      float64
	Length          *int16
	ReplacementCost float64
	Rating          *MpaaRating
	LastUpdate      time.Time
	SpecialFeatures *string
	Fulltext        string
}
