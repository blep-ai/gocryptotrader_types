package postgres

type insertCache struct {
	query        string
	retQuery     string
	valueMapping []uint64
	retMapping   []uint64
}
type updateCache struct {
	query        string
	valueMapping []uint64
}
