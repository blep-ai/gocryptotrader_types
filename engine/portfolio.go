package engine

type portfolioManager struct {
	started  int32
	stopped  int32
	shutdown chan struct{}
}
