package types

import "github.com/thealonemusk/WarpNet/pkg/types"

type Machine struct {
	types.Machine
	Connected bool
	OnChain   bool
	Online    bool
}
