package types

const (
	// ModuleName defines the module name
	ModuleName = "spc"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_spc"
)

var (
	ParamsKey = []byte("p_spc")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
