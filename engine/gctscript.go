package engine

type gctScriptManager struct {
	started  int32
	stopped  int32
	shutdown chan struct{}
}
