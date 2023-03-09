use cosmwasm_schema::cw_serde;
use cosmwasm_std::{Addr, Coin, Decimal, Timestamp, Uint128};
use cw_controllers::Claims;
use cw_storage_plus::{Item, Map};
use quasar_types::callback::{BondResponse, UnbondResponse};

use crate::msg::PrimitiveConfig;

// constants
pub const FALLBACK_RATIO: Decimal = Decimal::one();

// reply ids
pub const STRATEGY_BOND_ID: u64 = 80085;

// version info for migration info
pub const CONTRACT_NAME: &str = "basic-vault";
pub const CONTRACT_VERSION: &str = env!("CARGO_PKG_VERSION");

pub const CLAIMS: Claims = Claims::new("claims");

/// Investment info is fixed at instantiation, and is used to control the function of the contract
#[cw_serde]
pub struct InvestmentInfo {
    /// Owner created the contract and takes a cut
    pub owner: Addr,
    /// This is the minimum amount we will pull out to reinvest, as well as a minimum
    /// that can be unbonded (to avoid needless staking tx)
    pub min_withdrawal: Uint128,
    /// this is the array of primitives that this vault will subscribe to
    pub primitives: Vec<PrimitiveConfig>,
}

/// Supply is dynamic and tracks the current supply of staked and ERC20 tokens.
#[cw_serde]
#[derive(Default)]
pub struct Supply {
    /// issued is how many derivative tokens this contract has issued
    pub issued: Uint128,
    /// bonded is how many native tokens exist bonded to the validator
    pub bonded: Uint128,
    /// claims is how many tokens need to be reserved paying back those who unbonded
    pub claims: Uint128,
}

pub const INVESTMENT: Item<InvestmentInfo> = Item::new("invest");
pub const TOTAL_SUPPLY: Item<Supply> = Item::new("total_supply");

#[cw_serde]
#[derive(Default)]
pub struct BondingStub {
    // the contract address of the primitive
    pub address: String,
    // the response of the primitive upon successful bond or error
    pub bond_response: Option<BondResponse>,
}

#[cw_serde]
#[derive(Default)]
pub struct Unbond {
    pub stub: Vec<UnbondingStub>,
    pub shares: Uint128,
}

#[cw_serde]
#[derive(Default)]
pub struct UnbondingStub {
    // the contract address of the primitive
    pub address: String,
    // the response of the primitive upon successful bond or error
    pub unlock_time: Option<Timestamp>,
    // response of the unbond, if this is present then we have finished unbonding
    pub unbond_response: Option<UnbondResponse>,
    // funds attached to the unbond_response
    pub unbond_funds: Vec<Coin>,
}

// (un)bonding sequence number (to map primitive responses to the right bond action)
pub const BONDING_SEQ: Item<Uint128> = Item::new("bond_seq");
// mapping from bonding sequence number to depositor/withdrawer address (todo: better way to do this?)
// TODO addresses should be of type ADDR
pub const BONDING_SEQ_TO_ADDR: Map<String, String> = Map::new("bond_seq_to_addr");
// current bonds pending for a user
pub const PENDING_BOND_IDS: Map<Addr, Vec<String>> = Map::new("pending_bond_ids");
// current unbonds pending for a user
pub const PENDING_UNBOND_IDS: Map<Addr, Vec<String>> = Map::new("pending_unbond_ids");
// map of bond id to bond state
// todo: find the type of the vec items here (replace supply obvs)
pub const BOND_STATE: Map<String, Vec<BondingStub>> = Map::new("bond_state");
// map of unbond id to unbond state
pub const UNBOND_STATE: Map<String, Unbond> = Map::new("unbond_state");

pub const DEBUG_TOOL: Item<String> = Item::new("debug_tool");