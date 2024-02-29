package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// ShardKeyPrefix is the prefix to retrieve all Shard
	ShardKeyPrefix = "Shard/value/"
)

// ShardKey returns the store key to retrieve a Shard from the index fields
func ShardKey(
	index uint64,
) []byte {
	var key []byte

	indexBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(indexBytes, index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
