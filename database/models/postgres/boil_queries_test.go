package postgres

import (
	"bytes"
	"io"
	"regexp"
)

type fKeyDestroyer struct {
	reader io.Reader
	buf    *bytes.Buffer
	rgx    *regexp.Regexp
}
