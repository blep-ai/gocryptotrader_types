package dispatch

import (
	"sync"

	"github.com/gofrs/uuid"
)

type Dispatcher struct {
	// routes refers to a subystem uuid ticket map with associated publish
	// channels, a relayer will be given a unique id through its job channel,
	// then publish the data across the full registered channels for that uuid.
	// See relayer() method below.
	routes map[uuid.UUID][]chan interface{}

	// rMtx protects the routes variable ensuring acceptable read/write access
	rMtx sync.RWMutex

	// Persistent buffered job queue for relayers
	jobs chan *job

	// Dynamic channel pool; returns an unbuffered channel for routes map
	outbound sync.Pool

	// MaxWorkers defines max worker ceiling
	maxWorkers int32
	// Atomic values -----------------------
	// Worker counter
	count int32
	// Dispatch status
	running uint32

	// Unbufferd shutdown chan, sync wg for ensuring concurrency when only
	// dropping a single relayer routine
	shutdown chan *sync.WaitGroup

	// Relayer shutdown tracking
	wg sync.WaitGroup
}
type job struct {
	Data interface{}
	ID   uuid.UUID
}
type Mux struct {
	// Reference to the main running dispatch service
	d *Dispatcher
	sync.RWMutex
}
type Pipe struct {
	// Channel to get all our lovely informations
	C chan interface{}
	// ID to tracked system
	id uuid.UUID
	// Reference to multiplexor
	m *Mux
}
