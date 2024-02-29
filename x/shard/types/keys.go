package types

const (
	// ModuleName defines the module name
	ModuleName = "shard"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_shard"
)

var (
	ParamsKey = []byte("p_shard")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	ToParentKey      = "ToParent/value/"
	ToParentCountKey = "ToParent/count/"
)

const (
	ToLeftChildKey      = "ToLeftChild/value/"
	ToLeftChildCountKey = "ToLeftChild/count/"
)

const (
	ToRightChildKey      = "ToRightChild/value/"
	ToRightChildCountKey = "ToRightChild/count/"
)
