package fingerprint

import (
	"fmt"

	"github.com/go-goll/libsignal-protocol-go/util/bytehelper"
)

// NewDisplay will return a new displayable fingerprint.
func NewDisplay() *Display {
	return &Display{}
}

// Display is a structure for displayable fingerprints.
type Display struct{}

// CreateFor create
func (d *Display) CreateFor(localIdentityKey, remoteIdentityKey []byte) string {

	localFingerprint := displayStringFor(localIdentityKey)
	remoteFingerprint := displayStringFor(remoteIdentityKey)
	if localFingerprint < remoteFingerprint {
		return localFingerprint + remoteFingerprint
	}
	return remoteFingerprint + localFingerprint
}

// displayStringFor will return a displayable string representation
// of the given fingerprint.
func displayStringFor(fingerprint []byte) string {
	return encodedChunk(fingerprint, 0) +
		encodedChunk(fingerprint, 5) +
		encodedChunk(fingerprint, 10) +
		encodedChunk(fingerprint, 15) +
		encodedChunk(fingerprint, 20) +
		encodedChunk(fingerprint, 25)
}

// encodedChunk will return an encoded string of the given hash.
func encodedChunk(hash []byte, offset int) string {
	chunk := bytehelper.Bytes5ToInt64(hash, offset) % 100000
	return fmt.Sprintf("%05d", chunk)
}
