package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		PlatformList: []Platform{},
		ClaimList:    []Claim{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in platform
	platformIndexMap := make(map[string]struct{})

	for _, elem := range gs.PlatformList {
		index := string(PlatformKey(elem.Index))
		if _, ok := platformIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for platform")
		}
		platformIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in claim
	claimIndexMap := make(map[string]struct{})

	for _, elem := range gs.ClaimList {
		index := string(ClaimKey(elem.Index))
		if _, ok := claimIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for claim")
		}
		claimIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
