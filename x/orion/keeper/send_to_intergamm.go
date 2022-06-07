package keeper

import (
	"fmt"
	"time"

	"github.com/abag/quasarnode/x/orion/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	ibcclienttypes "github.com/cosmos/ibc-go/v3/modules/core/02-client/types"
)

// getOwnerAccStr returns the module account bech32 with which ICA transactions to be done
func (k Keeper) getOwnerAcc() sdk.AccAddress {
	return k.accountKeeper.GetModuleAddress(types.ModuleName)
}

// getOwnerAccStr returns the module account bech32 with which ICA transactions to be done
func (k Keeper) getOwnerAccStr() string {
	// For initial testing use alice address -
	// TODO AUDIT here (which return?)
	//return "quasar1sqlsc5024sszglyh7pswk5hfpc5xtl77gqjwec" // alice on quasar

	accAddr := k.accountKeeper.GetModuleAddress(types.ModuleName)
	accStr, err := sdk.Bech32ifyAddressBytes("quasar", accAddr)
	if err != nil {
		panic(err)
	}
	return accStr
}

func (k Keeper) getDestinationAccStr() string {
	// TODO - Call interchain account getter from the intergamm module
	return "osmo1t8eh66t2w5k67kwurmn5gqhtq6d2ja0vp7jmmq" // alice on osmosis
	// return "osmo1hphwfu3yjf82z8xpcl6e05gzkjwjmu8ts2m97mdk62feuqm77f2skm6qcy"
}

// TODO :  should be a parameter
func (k Keeper) getIntermediateReceiver() string {
	return "cosmos1ppkxa0hxak05tcqq3338k76xqxy2qse96uelcu" // alice on cosmos
}

// getConnectionId returns the connection identifier to osmosis from intergamm module
func (k Keeper) getConnectionId(chainid string) string {
	// TODO - Get it from the param
	// TO osmosis - connection-0
	// TO cosmoshub - connection-0
	// connection could also be self determined by integamm. It should be a param in intergamm
	// But as orion has intergamm keeper access; it can get it from intergamm
	return "connection-0"
}

// Intergamm module method wrappers
func (k Keeper) JoinPool(ctx sdk.Context, poolID uint64, shareOutAmount sdk.Int, tokenInMaxs []sdk.Coin) (uint64, error) {
	k.Logger(ctx).Info(fmt.Sprintf("Entered JoinPool|poolID=%v|shareOutAmount=%v|tokenInMaxs=%v\n",
		poolID, shareOutAmount, tokenInMaxs))

	owner := k.getOwnerAccStr()
	connectionId := k.getConnectionId("osmosis")
	timeoutTimestamp := time.Now().Add(time.Minute).Unix()
	packetSeq, err := k.intergammKeeper.TransmitIbcJoinPool(
		ctx,
		owner,
		connectionId,
		uint64(timeoutTimestamp),
		poolID,
		shareOutAmount,
		tokenInMaxs,
	)
	return packetSeq, err
}

func (k Keeper) ExitPool(ctx sdk.Context, poolID uint64, shareInAmount sdk.Int, tokenOutMins []sdk.Coin) (uint64, error) {
	k.Logger(ctx).Info(fmt.Sprintf("Entered JoinPool|poolID=%v|shareInAmount=%v|tokenOutMins=%v\n",
		poolID, shareInAmount, tokenOutMins))

	owner := k.getOwnerAccStr()
	connectionId := k.getConnectionId("osmosis")
	timeoutTimestamp := time.Now().Add(time.Minute).Unix()
	seq, err := k.intergammKeeper.TransmitIbcExitPool(
		ctx,
		owner,
		connectionId,
		uint64(timeoutTimestamp),
		poolID,
		shareInAmount,
		tokenOutMins,
	)
	return seq, err
}

func (k Keeper) TokenWithdrawFromOsmosis(ctx sdk.Context, coin sdk.Coin) error {
	k.Logger(ctx).Info("TokenWithdrawFromOsmosis", "coin", coin)
	owner := k.getOwnerAccStr()
	receiverAddr := k.getOwnerAccStr() // receiver is same as owner address
	connectionId := k.getConnectionId("osmosis")
	timeoutTimestamp := time.Now().Add(time.Minute).Unix()
	transferPort := "transfer"        // TODO - should be a param
	transferChannel := "channel-1"    // TODO - should be a param
	fwdTransferPort := "transfer"     // TODO - should be a param
	fwdTransferChannel := "channel-1" // TODO - should be a param
	intermediateReceiver := k.getIntermediateReceiver()

	_, err := k.intergammKeeper.TransmitForwardIbcTransfer(
		ctx,
		owner,
		connectionId,
		uint64(timeoutTimestamp),
		transferPort,
		transferChannel,
		coin,
		fwdTransferPort,
		fwdTransferChannel,
		intermediateReceiver,
		receiverAddr,
		ibcclienttypes.ZeroHeight(),
		uint64(timeoutTimestamp),
	)
	return err
}

// IBCTokenTransfer does the multi hop token transfer to the osmosis interchain account via middle chain.
// Returns the packet sequence number of the outgoing packet.
// Logic - if denom is ibc atom then fwd it via cosmos-hub, and so on.
// All these details is supposed to be available in a generic place independent of orion module.
// Intergamm module should be inteligent enough to route packets based on denoms.
// The best orion or any other module can provide to intergam is to give denom string.
// Assumption - IBC token transfer is robust enough to deal with failure. It will return tokens
// in case of failure.
// Can we actually query or determine if ibc token transfer call was successful. And if it failed; tokens
// are returned.
func (k Keeper) IBCTokenTransfer(ctx sdk.Context, coin sdk.Coin) (uint64, error) {
	logger := k.Logger(ctx)
	logger.Info("IBCTokenTransfer",
		"coin", coin,
	)
	destAccStr := k.getDestinationAccStr()
	owner := k.getOwnerAcc()
	return k.intergammKeeper.SendToken(ctx, "osmosis", owner, destAccStr, coin)

	// k.SetIBCTokenTransferRecord(ctx, seqNo, coin)
}

// It should also have State Tx logic.
// Value can be more elegant struct here.
func (k Keeper) SetIBCTokenTransferRecord(ctx sdk.Context, seqNo uint64, coin sdk.Coin) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.IBCTokenTransferKBP)
	key := types.CreateSeqKey(seqNo)
	value := k.cdc.MustMarshal(&coin)
	store.Set(key, value)
}

// DeleteIBCTokenTransferRecord should be called when we receive positve ack.
// It means that we are done with the seq number successfully
func (k Keeper) DeleteIBCTokenTransferRecord(ctx sdk.Context, seqNo uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.IBCTokenTransferKBP)
	key := types.CreateSeqKey(seqNo)
	store.Delete(key)
}
func (k Keeper) GetIBCTokenTransferRecord(ctx sdk.Context, seqNo uint64) (coin sdk.Coin, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.IBCTokenTransferKBP)
	key := types.CreateSeqKey(seqNo)
	b := store.Get(key)
	if b == nil {
		return coin, false
	}
	k.cdc.MustUnmarshal(b, &coin)
	return coin, true
}

func (k Keeper) GetTotalEpochTransffered(ctx sdk.Context, epochNumber uint64) sdk.Coins {
	return nil
}
