package slashing

import (
	"context"
	"fmt"
	"time"

	"cosmossdk.io/core/comet"

	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/slashing/keeper"
	"github.com/cosmos/cosmos-sdk/x/slashing/types"
)

// BeginBlocker check for infraction evidence or downtime of validators
// on every begin block
func BeginBlocker(ctx context.Context, k keeper.Keeper) error {
	fmt.Println(" ------INSIDE SLASHING BEGIN BLOCKER ------")
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyBeginBlocker)

	// Iterate over all the validators which *should* have signed this block
	// store whether or not they have actually signed it and slash/unbond any
	// which have missed too many blocks in a row (downtime slashing)
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	for _, voteInfo := range sdkCtx.VoteInfos() {
		fmt.Println("calling handle validator signature for: ", voteInfo.Validator.Address, " with power", voteInfo.Validator.Power)
		err := k.HandleValidatorSignature(ctx, voteInfo.Validator.Address, voteInfo.Validator.Power, comet.BlockIDFlag(voteInfo.BlockIdFlag))
		if err != nil {
			return err
		}
	}
	fmt.Println("------------END OF SLASHING BEGIN BLOCKER----------")
	return nil
}
