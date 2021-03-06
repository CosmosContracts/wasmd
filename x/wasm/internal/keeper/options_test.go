package keeper

import (
	"github.com/CosmWasm/wasmd/x/wasm/internal/keeper/wasmtesting"
	"github.com/CosmWasm/wasmd/x/wasm/internal/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	distributionkeeper "github.com/cosmos/cosmos-sdk/x/distribution/keeper"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConstructorOptions(t *testing.T) {
	specs := map[string]struct {
		srcOpt Option
		verify func(Keeper)
	}{
		"wasm engine": {
			srcOpt: WithWasmEngine(&wasmtesting.MockWasmer{}),
			verify: func(k Keeper) {
				assert.IsType(t, k.wasmer, &wasmtesting.MockWasmer{})
			},
		},
		"message handler": {
			srcOpt: WithMessageHandler(&wasmtesting.MockMessageHandler{}),
			verify: func(k Keeper) {
				assert.IsType(t, k.messenger, &wasmtesting.MockMessageHandler{})
			},
		},
	}
	for name, spec := range specs {
		t.Run(name, func(t *testing.T) {
			k := NewKeeper(
				nil,
				nil,
				paramtypes.NewSubspace(nil, nil, nil, nil, ""),
				authkeeper.AccountKeeper{},
				nil,
				stakingkeeper.Keeper{},
				distributionkeeper.Keeper{},
				nil,
				nil,
				nil,
				nil,
				"tempDir",
				types.DefaultWasmConfig(),
				"",
				nil,
				nil,
				spec.srcOpt,
			)
			spec.verify(k)
		})
	}

}
