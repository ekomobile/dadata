package dadata

import (
	"testing"
)

func TestNewDaData(t *testing.T) {
	daData := NewClient()

	if daData == nil {
		t.Errorf(`NewClient return nil`)
	}
}
