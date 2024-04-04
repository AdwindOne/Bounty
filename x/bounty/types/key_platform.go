package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// PlatformKeyPrefix is the prefix to retrieve all Platform
	PlatformKeyPrefix = "Platform/value/"
)

// PlatformKey returns the store key to retrieve a Platform from the index fields
func PlatformKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
