package vm

import (
	"context"
	"time"

	"github.com/d5/tengo/v2"
	"github.com/gofrs/uuid"
)

type VM struct {
	ID       uuid.UUID
	Hash     string
	File     string
	Path     string
	Script   *tengo.Script
	Compiled *tengo.Compiled
	ctx      context.Context
	T        time.Duration
	NextRun  time.Time
	S        chan struct{}
}
