package crypto
import (
	"crypto/hmac"
	"crypto/md5" // nolint // Used by some exchanges
	"crypto/rand"
	"crypto/sha1" // nolint // Used by some exchanges
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"hash"
	"io"
)
