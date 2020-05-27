package indicators

import (
	objects "github.com/d5/tengo/v2"
)

type ATR struct {
	objects.Array
	Period int
}
