package keeper

import (
	"encoding/binary"
	"fmt"
	"strconv"
	"time"

	"github.com/abag/quasarnode/x/osmolpv/types"
	qbanktypes "github.com/abag/quasarnode/x/qbank/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

/*
	LpID                   uint64           `protobuf:"varint,1,opt,name=lpID,proto3" json:"lpID,omitempty"`
	LockID                 uint64           `protobuf:"varint,2,opt,name=lockID,proto3" json:"lockID,omitempty"`
	IsActive               bool             `protobuf:"varint,3,opt,name=isActive,proto3" json:"isActive,omitempty"`
	StartTime              time.Time        `protobuf:"bytes,4,opt,name=startTime,proto3,stdtime" json:"startTime" yaml:"startTime"`
	BondingStartEpochDay   uint64           `protobuf:"varint,5,opt,name=bondingStartEpochDay,proto3" json:"bondingStartEpochDay,omitempty"`
	BondDuration           uint64           `protobuf:"varint,6,opt,name=bondDuration,proto3" json:"bondDuration,omitempty"`
	UnbondingStartEpochDay uint64           `protobuf:"varint,7,opt,name=unbondingStartEpochDay,proto3" json:"unbondingStartEpochDay,omitempty"`
	UnbondingDuration      uint64           `protobuf:"varint,8,opt,name=unbondingDuration,proto3" json:"unbondingDuration,omitempty"`
	PoolID                 uint64           `protobuf:"varint,9,opt,name=poolID,proto3" json:"poolID,omitempty"`
	Lptoken                *types.Coin      `protobuf:"bytes,10,opt,name=lptoken,proto3" json:"lptoken,omitempty"`
	Coins                  []types.Coin     `protobuf:"bytes,11,rep,name=coins,proto3" json:"coins"`
	Gaugelocks
*/

// Zero Value of LpID, lockid means invalid values.
// Write proper unit tests.
func NewLP(lockid, bondingStartEpochday, bondDuration, unbondingStartEpochDay,
	unbondingDuration, poolID uint64, lpToken sdk.Coin, coins sdk.Coins) types.LpPosition {
	lp := types.LpPosition{LpID: 0,
		LockID:                 lockid,
		IsActive:               false,
		StartTime:              time.Now(), // TODO AUDIT
		BondingStartEpochDay:   bondingStartEpochday,
		BondDuration:           bondDuration,
		UnbondingStartEpochDay: unbondingDuration,
		UnbondingDuration:      unbondingDuration,
		PoolID:                 poolID,
		Lptoken:                &lpToken, // TODO AUDIT &
		Coins:                  coins,    // TODO AUDIT types
	}
	return lp
}

// GetDepositCount get the total number of deposit
func (k Keeper) GetLPCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.LPCountKBP)
	byteKey := types.CreateLPCountKey()

	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetDepositCount set the total number of deposit
// Note - This could be a reduntant method.
func (k Keeper) setLPCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.LPCountKBP)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(types.CreateLPCountKey(), bz)
}

// AddNewLPPosition update the LP ID of the newly created lp and set the position data in the KV store.
func (k Keeper) AddNewLPPosition(ctx sdk.Context, lpPosition types.LpPosition) {
	// count := k.GetLPCount(ctx)
	lps, _ := k.GetLpStat(ctx, lpPosition.BondingStartEpochDay)
	count := lps.LpCount
	lpPosition.LpID = count + 1
	k.setLpPosition(ctx, lpPosition)
	k.setLpEpochPosition(ctx, lpPosition.LpID, lpPosition.BondingStartEpochDay)

	lps.LpCount = lpPosition.LpID
	var tmp sdk.Coins
	tmp = lps.TotalLPCoins
	for _, coin := range lpPosition.Coins {
		tmp = tmp.Add(coin)
	}
	lps.TotalLPCoins = tmp
	k.SetLpStat(ctx, lpPosition.BondingStartEpochDay, lps)
	// k.setLPCount(ctx, count)
}

// SetLpPosition set lpPosition created by the strategy in a given epochday in the
// prefixed kv store with key formed using epoch day and lpID.
// key = types.LPPositionKBP + {epochday} + {":"} + {lpID}
// Value = types.LpPosition)
func (k Keeper) setLpPosition(ctx sdk.Context, lpPosition types.LpPosition) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.LPPositionKBP)
	key := types.EpochLPIDKey(lpPosition.BondingStartEpochDay, lpPosition.LpID)
	b := k.cdc.MustMarshal(&lpPosition)
	store.Set(key, b)
}

// SetLpEpochPosition set is used to store reverse mapping lpID and epochday as part of key.
// key = types.LPEpochKBP + {lpID} + {":"} + {epochDay}
func (k Keeper) setLpEpochPosition(ctx sdk.Context, lpID uint64, epochday uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.LPEpochKBP)
	key := types.CreateLPEpochKey(lpID, epochday)
	store.Set(key, []byte{0x00})
}

// GetLPEpochDay fetch the epochday of an lp position on which the position was created.
func (k Keeper) GetLPEpochDay(ctx sdk.Context, lpID uint64) (epochday uint64, found bool) {
	bytePrefix := types.LPEpochKBP
	prefixKey := types.CreateLPIDKey(lpID)
	prefixKey = append(bytePrefix, prefixKey...)

	store := ctx.KVStore(k.storeKey)
	iter := sdk.KVStorePrefixIterator(store, prefixKey)
	defer iter.Close()
	// Note : Only one entry;  iteration will have maximum one loop
	for ; iter.Valid(); iter.Next() {
		key, _ := iter.Key(), iter.Value()
		epochStr := string(key)
		epochday, _ = strconv.ParseUint(epochStr, 10, 64)
		found = true
	}
	return epochday, found
}

// GetLpPosition fetch the lpPosition based on the epochday and lpID input
// Note - This could be a reduntant method.
func (k Keeper) GetLpPosition(ctx sdk.Context, epochDay uint64, lpID uint64) (val types.LpPosition, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.LPPositionKBP)
	key := types.EpochLPIDKey(epochDay, lpID)
	b := store.Get(key)
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveLpPosition removes lpPosition from the store prefixed with epochDay and lpID
func (k Keeper) RemoveLpPosition(ctx sdk.Context, epochDay uint64, lpID uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.LPPositionKBP)
	key := types.EpochLPIDKey(epochDay, lpID)
	store.Delete(key)
}

// GetLPIDList fetch the list of lp position lpid for a given epoch day
func (k Keeper) GetLPIDList(ctx sdk.Context, epochday uint64) []uint64 {
	var lpIDs []uint64
	bytePrefix := types.LPPositionKBP
	prefixKey := types.EpochDayKey(epochday)
	prefixKey = append(bytePrefix, prefixKey...)
	prefixKey = append(prefixKey, qbanktypes.SepByte...)

	// prefixKey = types.LPPositionKBP + {epochday} + {":"}
	store := ctx.KVStore(k.storeKey)
	iter := sdk.KVStorePrefixIterator(store, prefixKey)
	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		key, _ := iter.Key(), iter.Value()
		lpIDStr := string(key)
		lpID, _ := strconv.ParseUint(lpIDStr, 10, 64)
		lpIDs = append(lpIDs, lpID)
	}
	return lpIDs
}

// TODO | AUDIT
// CalculateLPWeight calc weight of an Lp position in the current epoch.
// This weight will be used for the approx fair reward distribution.
// Logic -
// 1. Get the lp position
// 2. Get current active gauge of lpID
// 3. Get expected apy
// 4. Get total  tvl*apy
// 5. Calc weight as ( < tvl of this lpid > * <apy of lpID> ) / Sum (tvl*apy)
// NOTE - Weight should be updated on each epoch in the object itself.
// There should be a dedicated field for this.
func (k Keeper) CalculateLPWeight(ctx sdk.Context, epochDay uint64, lpID uint64) (sdk.Dec, error) {
	lpp, found := k.GetLpPosition(ctx, epochDay, lpID)
	if !found {
		return sdk.ZeroDec(), fmt.Errorf("LP position not found")
	}
	k.Logger(ctx).Info(fmt.Sprintf("CalculateLPWeight|epochday=%v|lpp=%v\n", epochDay, lpp))
	g := k.GetCurrentActiveGauge(ctx, epochDay, lpID)
	tvl, _ := k.GetLpTvl(ctx, epochDay, lpID)
	tvlApy := g.ExpectedApy.Mul(tvl)
	totalTvlApy := k.GetTotalTvlApy(ctx, epochDay)
	lpw := tvlApy.Quo(totalTvlApy)
	return lpw, nil
}

// GetTotalTvlApy calculate the total tvl in terms of amount of orion receipt tokens.
// Option #1 Normalize this in terms of one denom. For example - If it is pool#1 <atom, osmo>
// Then calculate this interms of atom or osmo
// Option #2 Calculate in terms of allocated orions amount.
func (k Keeper) GetLpTvl(ctx sdk.Context, epochday uint64, lpID uint64) (sdk.Dec, error) {
	lpp, found := k.GetLpPosition(ctx, epochday, lpID)
	if !found {
		return sdk.ZeroDec(), fmt.Errorf("LP position not found")
	}
	return lpp.Lptoken.Amount.ToDec(), nil
}

// GetTotalTvlApy calculate the total tvl in terms of amount of orion receipt tokens.
func (k Keeper) GetTotalTvlApy(ctx sdk.Context, epochDay uint64) sdk.Dec {
	lpi, _ := k.GetEpochLPInfo(ctx, epochDay)
	return lpi.TotalTVL.Amount.ToDec()
}

// Note - This maynot be used
// AddEpochLPUser add kv store with key = {epochday} + {":"} + {lpID} + {":"} + {userAccount} + {":"} + {denom}
// value = sdk.Coin
func (k Keeper) AddEpochLPUserDenomAmt(ctx sdk.Context, epochday uint64, lpID uint64, userAcc string, coin sdk.Coin) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.EpochLPUserKBP)
	key := types.CreateEpochLPUserKey(epochday, lpID, userAcc, coin.Denom)
	b := store.Get(key)
	if b == nil {
		value := k.cdc.MustMarshal(&coin)
		store.Set(key, value)
	} else {
		var storedCoin sdk.Coin
		k.cdc.MustUnmarshal(b, &storedCoin)
		storedCoin = storedCoin.Add(coin)
		value := k.cdc.MustMarshal(&storedCoin)
		store.Set(key, value)
	}
}

// Note - This maynot be used
// GetEpochLPUser get user's denom amount used in a given and epoch day and lp id
func (k Keeper) GetEpochLPUserDenomAmt(ctx sdk.Context, epochday uint64, lpID uint64, userAcc string, denom string) (val sdk.Coin, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.EpochLPUserKBP)
	key := types.CreateEpochLPUserKey(epochday, lpID, userAcc, denom)
	b := store.Get(key)
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// Note - This maynot be used
// GetEpochLPUser get list of user denom-amount pair used in a given lp id.
// The output from this function is used to calculate a users contribution towards
// the reward associated with this lp position.
func (k Keeper) GetEpochLPUserCoin(ctx sdk.Context, epochday uint64, lpID uint64) ([]types.UserCoin, sdk.Coins) {
	bytePrefix := types.EpochLPUserKBP
	prefixKey := types.CreateEpochLPKey(epochday, lpID)
	prefixKey = append(bytePrefix, prefixKey...)
	store := ctx.KVStore(k.storeKey)
	iter := sdk.KVStorePrefixIterator(store, prefixKey)
	defer iter.Close()

	var result []types.UserCoin
	// var usersWeight []types.EpochUserDenomWeight
	var totalCoins sdk.Coins
	logger := k.Logger(ctx)
	logger.Info(fmt.Sprintf("GetEpochLPUserCoin|modulename=%s|blockheight=%d|prefixKey=%s",
		types.ModuleName, ctx.BlockHeight(), string(prefixKey)))

	// key = {user} + {":"} + {denom}
	for ; iter.Valid(); iter.Next() {
		key, val := iter.Key(), iter.Value()
		uid, _ := types.ParseUserDenomKey(key)
		var coin sdk.Coin
		k.cdc.MustUnmarshal(val, &coin)
		uc := types.UserCoin{UserAcc: uid, Coin: coin}
		totalCoins = totalCoins.Add(coin) // Test if totalCoins will get sorted internally
		result = append(result, uc)

	}

	return result, totalCoins
}

// TODO | AUDIT
// GetCurrentActiveGauge fetch the currently active gauge of an LP position in the live chain
func (k Keeper) GetCurrentActiveGauge(ctx sdk.Context, epochday uint64, lpID uint64) types.GaugeLockInfo {
	var activeG types.GaugeLockInfo
	lp, _ := k.GetLpPosition(ctx, epochday, lpID)
	e, _ := k.GetEpochDayInfo(ctx, epochday)
	for _, g := range lp.Gaugelocks {
		if e.BlockTime.After(g.StartTime) && g.StartTime.Add(g.LockupDuration).After(e.BlockTime) {
			activeG = *g
		}
	}
	return activeG
}

// GetEpochLPDenomAmt returns the total amount of a particular denom used
func (k Keeper) GetEpochLPDenomAmt(ctx sdk.Context, epochday uint64, denom string) (sdk.Int, error) {
	lps, _ := k.GetLpStat(ctx, epochday)
	var tmp sdk.Coins = lps.TotalLPCoins
	// tmp should be sorted
	if !tmp.IsValid() {
		return sdk.ZeroInt(), fmt.Errorf("lps.TotalLPCoins is not sorted")
	}
	amt := tmp.AmountOf(denom)
	return amt, nil
}
