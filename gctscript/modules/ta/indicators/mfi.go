package indicators

import (
	objects "github.com/d5/tengo/v2"
)

type MFI struct {
	objects.Array
	Period int
}
