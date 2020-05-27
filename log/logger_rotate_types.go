package log

import (
	"os"
	"sync"
)

type Rotate struct {
	FileName string
	Rotate   *bool
	MaxSize  int64

	size   int64
	output *os.File
	mu     sync.Mutex
}
