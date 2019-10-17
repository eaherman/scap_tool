package scap

import (
	"../pkg/scap"
	"testing"
)

func TestSetHost(t *testing.T) {
	sr := scap.ScapRunner{}
	sr.SetHost()
	if sr.Name != "Name" {
		t.Errorf("Does not equal Name")
	}
}
