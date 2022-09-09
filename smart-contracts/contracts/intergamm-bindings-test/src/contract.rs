use std::env;

#[cfg(not(feature = "library"))]
use cosmwasm_std::entry_point;
use cosmwasm_std::{
    to_binary, Binary, Coin, Deps, DepsMut, Env, IbcMsg, IbcTimeout, MessageInfo, Reply, Response,
    StdResult, Storage, SubMsg, Uint64,
};
use cw2::set_contract_version;
use intergamm_bindings::helper::{create_intergamm_msg, handle_reply};
use intergamm_bindings::msg::IntergammMsg;

use crate::error::ContractError;
use crate::msg::{AckTriggeredResponse, ExecuteMsg, InstantiateMsg, QueryMsg};
use crate::state::{ACKS, ACKTRIGGERED};

// version info for migration info
const CONTRACT_NAME: &str = "crates.io:intergamm-bindings-test-2";
const CONTRACT_VERSION: &str = env!("CARGO_PKG_VERSION");

#[cfg_attr(not(feature = "library"), entry_point)]
pub fn instantiate(
    deps: DepsMut,
    _env: Env,
    _info: MessageInfo,
    _msg: InstantiateMsg,
) -> Result<Response, ContractError> {
    // set the ack triggered to false
    set_contract_version(deps.storage, CONTRACT_NAME, CONTRACT_VERSION)?;
    ACKTRIGGERED.save(deps.storage, &0)?;
    Ok(Response::default())
}

#[entry_point]
pub fn reply(deps: DepsMut, _env: Env, msg: Reply) -> StdResult<Response> {
    handle_reply(deps.storage, msg, ACKS)
}

#[cfg_attr(not(feature = "library"), entry_point)]
pub fn execute(
    deps: DepsMut,
    env: Env,
    info: MessageInfo,
    msg: ExecuteMsg,
) -> Result<Response<IntergammMsg>, ContractError> {
    match msg {
        ExecuteMsg::SendToken {
            destination_local_zone_id,
        } => execute_send_token(destination_local_zone_id, env),
        ExecuteMsg::SendTokenIbc {
            channel_id,
            to_address,
            amount,
        } => execute_send_token_ibc(channel_id, to_address, amount, env),
        ExecuteMsg::RegisterInterchainAccount { connection_id } => {
            execute_register_ica(connection_id, env)
        }
        ExecuteMsg::JoinSinglePool {
            connection_id,
            pool_id,
            share_out_min_amount,
            token_in,
        } => execute_join_pool(
            connection_id,
            pool_id,
            share_out_min_amount,
            token_in,
            deps,
            env,
        ),
        ExecuteMsg::TestIcaScenario {} => execute_test_scenario(env),
        ExecuteMsg::AckTriggered {} => do_ibc_packet_ack(deps, env),
        ExecuteMsg::Deposit {} => execute_deposit(info),
    }
}

// TODO as of 23 august, there is a bug in the go implementation of send token
pub fn execute_send_token(
    destination_local_zone_id: String,
    env: Env,
) -> Result<Response<IntergammMsg>, ContractError> {
    Ok(Response::new()
        .add_message(IntergammMsg::SendToken {
            creator: env.contract.address.to_string(),
            destination_local_zone_id: destination_local_zone_id,
            sender: env.contract.address.to_string(),
            receiver: env.contract.address.to_string(),
            coin: Coin::new(100, "uqsr"),
        })
        .add_attribute("sending tokens", "100 uqsr to osmosis"))
}

pub fn execute_send_token_ibc(
    channel_id: String,
    to_address: String,
    amount: Coin,
    env: Env,
) -> Result<Response<IntergammMsg>, ContractError> {
    // timeout in 600 seconds after current block timestamp
    let timeout = IbcTimeout::with_timestamp(env.block.time.plus_seconds(600));
    Ok(Response::new().add_message(IbcMsg::Transfer {
        channel_id,
        to_address,
        amount,
        timeout,
    }))
}

pub fn execute_test_scenario(env: Env) -> Result<Response<IntergammMsg>, ContractError> {
    Ok(Response::new().add_message(IntergammMsg::TestScenario {
        creator: env.contract.address.to_string(),
        scenario: "registerIca".to_string(),
    }))
}

// join pool requires us to have a pool on the remote chain
pub fn execute_join_pool(
    connection_id: String,
    pool_id: Uint64,
    share_out_min_amount: i64,
    token_in: Coin,
    deps: DepsMut,
    env: Env,
) -> Result<Response<IntergammMsg>, ContractError> {
    let msg = IntergammMsg::JoinSwapExternAmountIn {
        creator: env.contract.address.to_string(),
        connection_id,
        // timeout in 10 minutes
        timeout_timestamp: env.block.time.plus_seconds(600).nanos(),
        pool_id: pool_id.u64(),
        share_out_min_amount,
        token_in,
    };
    create_intergamm_msg(deps.storage, msg).map_err(|E | ContractError::Std(E))
}

pub fn execute_register_ica(
    connection_id: String,
    env: Env,
) -> Result<Response<IntergammMsg>, ContractError> {
    Ok(
        Response::new().add_submessage(SubMsg::new(IntergammMsg::RegisterInterchainAccount {
            creator: env.contract.address.to_string(),
            connection_id: connection_id,
        })),
    )
}

pub fn execute_deposit(info: MessageInfo) -> Result<Response<IntergammMsg>, ContractError> {
    let funds = cw_utils::one_coin(&info)?;
    if funds.denom != "uqsr" && funds.denom != "stake" {
        return Err(ContractError::PaymentError(
            cw_utils::PaymentError::MissingDenom("uqsr/stake".into()),
        ));
    }
    // we dont do anything else with the funds since we solely use them for testing and don't need to deposit
    Ok(Response::new()
        .add_attribute("deposit_amount", funds.amount)
        .add_attribute("deposit_denom", funds.denom))
}

#[cfg_attr(not(feature = "library"), entry_point)]
pub fn query(deps: Deps, _env: Env, msg: QueryMsg) -> StdResult<Binary> {
    match msg {
        QueryMsg::AckTriggered {} => to_binary(&query_ack_triggered(deps)?),
    }
}

pub fn query_ack_triggered(deps: Deps) -> StdResult<AckTriggeredResponse> {
    let triggered = ACKTRIGGERED.load(deps.storage)?;
    Ok(AckTriggeredResponse { triggered })
}

pub fn do_ibc_packet_ack(
    deps: DepsMut,
    _env: Env,
) -> Result<Response<IntergammMsg>, ContractError> {
    let triggered = ACKTRIGGERED.load(deps.storage)?;
    ACKTRIGGERED.save(deps.storage, &(triggered + 1))?;
    Ok(Response::new().add_attribute("ack tiggered", (triggered + 1).to_string()))
}

#[cfg(test)]
mod tests {}
