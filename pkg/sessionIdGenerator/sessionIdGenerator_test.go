package sessionidgenerator

import (
	"testing"
)

func TestSessionIdGenerator(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping session ID test in short mode.")
	} else {
		var saveID string
		for i := 1; i <= 1000000; i++ {
			sessionID := generateID()
			if saveID == sessionID {
				t.FailNow()
			} else {
				saveID = sessionID
			}
		}
	}
}
