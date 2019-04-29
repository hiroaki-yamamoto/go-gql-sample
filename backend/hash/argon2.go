package hash

import (
	"encoding/base64"

	"golang.org/x/crypto/argon2"
)

// Argon2 generates the argon2-hashed string.
func Argon2(raw, salt string) string {
	return base64.StdEncoding.EncodeToString(argon2.IDKey(
		[]byte(raw),
		[]byte(salt),
		50, 128*1024, 9, 32,
	))
}
