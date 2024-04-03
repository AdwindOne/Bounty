package types

const (
	// ModuleName defines the module name
	ModuleName = "bounty"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_bounty"
)

var (
	ParamsKey = []byte("p_bounty")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
