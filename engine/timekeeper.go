package engine

type ntpManager struct {
	started       int32
	stopped       int32
	inititalCheck bool
	shutdown      chan struct{}
}
