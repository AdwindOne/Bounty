package types

const (
	// ModuleName defines the module name
	ModuleName = "platform"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_platform"
)

var (
	ParamsKey = []byte("p_platform")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
