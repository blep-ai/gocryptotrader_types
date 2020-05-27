package postgres

import "testing"

func TestUpsert(t *testing.T) {
	t.Run("AuditEvents", testAuditEventsUpsert)
