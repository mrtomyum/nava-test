package models_test

import (
	"testing"
	m "tom/db/nava/models"
)

func TestNewMachine(t *testing.T) {
	if mc := new(m.Machine); mc == nil {
		t.Error("Expect instant new Machine but no new object ")
	}
}

func TestNewLoc(t *testing.T) {
	l := new(m.Loc)
	if l == nil {
		t.Error("Expect new Site instant")
	}
}

func TestLocTypeWhichIsStocking