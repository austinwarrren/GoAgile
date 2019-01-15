package sessionidgenerator

import (
	"github.com/rs/xid"
)

// sessionIdGenerator
// Generates unique session IDs
func generateID() string {
	sessionID := xid.New().String()
	return sessionID
}
