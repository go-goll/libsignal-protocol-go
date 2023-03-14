package tests

import (
	"github.com/go-goll/libsignal-protocol-go/serialize"
)

// newSerializer will return a JSON serializer for testing.
func newSerializer() *serialize.Serializer {
	serializer := serialize.NewProtoBufSerializer()

	return serializer
}
