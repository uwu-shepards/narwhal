package types

const (
	// ModuleName defines the module name
	ModuleName = "narwhal"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_narwhal"
)

var (
	ParamsKey = []byte("p_narwhal")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
