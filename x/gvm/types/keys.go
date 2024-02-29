package types

const (
	// ModuleName defines the module name
	ModuleName = "gvm"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_gvm"
)

var (
	ParamsKey = []byte("p_gvm")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
