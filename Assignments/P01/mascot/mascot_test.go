// Bryce Koch
// 4143 PLC

package mascot_test

import (
	"testing"

	"github.com/BKoch74/4143-PLC/mascot"
)

// Runs a test to verify that the BestMascot is a Lion. If not,
// outputs the error message to the console.
func TestMascot(t *testing.T) {
	if mascot.BestMascot() != "Lion" {
		t.Fatal("Wrong mascot!")
	}
}
