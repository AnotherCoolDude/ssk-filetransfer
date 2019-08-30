package proadclient

import (
	"strconv"
)

// StatusCode represents the project's current status
type StatusCode int

const (
	// StatusNone does not filter any StatusCode
	StatusNone StatusCode = 0
	// StatusVorbereitung 100
	StatusVorbereitung StatusCode = 100
	// StatusAngebot 200
	StatusAngebot StatusCode = 200
	// StatusDurchführung 300
	StatusDurchführung StatusCode = 300
	// StatusDurchführungTE Durchführung Teilabrechnung erfolgt 301
	StatusDurchführungTE StatusCode = 301
	// StatusGeliefert 400
	StatusGeliefert StatusCode = 400
	// StatusAbgerechnet 500
	StatusAbgerechnet StatusCode = 500
	// StatusAbgerechnetVP Abgerechnet - Verrechnung mit anderem Project 501
	StatusAbgerechnetVP StatusCode = 501
	// StatusAbgebrochen 600
	StatusAbgebrochen StatusCode = 600
	// StatusAbgebrochenNR Abgebrochen - nicht realisiert 601
	StatusAbgebrochenNR StatusCode = 601
	// StatusAbgebrochenNW Abgebrochen - nicht weiterberechenbar 602
	StatusAbgebrochenNW StatusCode = 602
)

// String converts the statuscode s to a string
func (s StatusCode) String() string {
	return strconv.Itoa(int(s))
}
