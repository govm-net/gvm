package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// ShardBlockKeyPrefix is the prefix to retrieve all ShardBlock
	ShardBlockKeyPrefix = "ShardBlock/value/"
)

// ShardBlockKey returns the store key to retrieve a ShardBlock from the index fields
func ShardBlockKey(
	index uint64,
) []byte {
	var key []byte

	indexBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(indexBytes, index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
