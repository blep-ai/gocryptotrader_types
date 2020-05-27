package indicators

import (
	objects "github.com/d5/tengo/v2"
)

type EMA struct {
	objects.Array
	Period int
}
