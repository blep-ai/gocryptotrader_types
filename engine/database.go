package engine

type databaseManager struct {
	started  int32
	stopped  int32
	shutdown chan struct{}
}
