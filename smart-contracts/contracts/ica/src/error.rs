use std::string::FromUtf8Error;
use thiserror::Error;

use cosmwasm_std::StdError;

/// Never is a placeholder to ensure we don't return any errors
#[derive(Error, Debug)]
pub enum Never {}

#[derive(Error, Debug, PartialEq)]
pub enum ContractError {
    #[error("{0}")]
    Std(#[from] StdError),

    #[error("Channel doesn't exist: {id}")]
    NoSuchChannel { id: String },

    #[error("Could not deserialize ica metadata, got {raw_metadata}, error: {error}")]
    InvalidIcaMetadata { raw_metadata: String, error: String },

    #[error("Incorrect ICA version, got {version}, want {contract_version}")]
    InvalidIcaVersion { version: String, contract_version: String},

    #[error("Incorrect ICA version, got {encoding}, want {contract_encoding}")]
    InvalidIcaEncoding { encoding: String, contract_encoding: String },

    #[error("Incorrect ICA version, got {tx_type}, want {contract_tx_type}")]
    InvalidIcaTxType { tx_type: String, contract_tx_type: String },

    #[error("Could not deserialize counterparty ica metadata, got {raw_metadata}, error: {error}")]
    InvalidCounterpartyIcaMetadata { raw_metadata: String, error: String },

    #[error("No Counterparty Ica Address")]
    NoCounterpartyIcaAddress {},

    #[error("No Counterparty Version")]
    NoCounterpartyVersion {},

    #[error("Only supports unordered channel")]
    OnlyOrderedChannel {},

    #[error("Parsed port from denom ({port}) doesn't match packet")]
    FromOtherPort { port: String },

    #[error("Parsed channel from denom ({channel}) doesn't match packet")]
    FromOtherChannel { channel: String },

    #[error("Cannot migrate from different contract type: {previous_contract}")]
    CannotMigrate { previous_contract: String },

    #[error("Cannot migrate from unsupported version: {previous_version}")]
    CannotMigrateVersion { previous_version: String },

    #[error("Failed to proto encode")]
    EncodingFail,

    #[error("Failed to proto decode")]
    DecodingFail,

    #[error("Only the governance contract can do this")]
    Unauthorized,
}

impl From<FromUtf8Error> for ContractError {
    fn from(_: FromUtf8Error) -> Self {
        ContractError::Std(StdError::invalid_utf8("parsing denom key"))
    }
}
