package mascot_test

import (
	"testing"

	"github.com/BKoch74/4143-PLC/mascot"
)

func TestMascot(t *testing.T) {
	if mascot.BestMascot() != "Lion" {
		t.Fatal("Wrong mascot!")
	}
}
