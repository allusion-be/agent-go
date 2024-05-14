// Package cmc provides a client for the "cmc" canister.
// Do NOT edit this file. It was automatically generated by https://github.com/aviate-labs/agent-go.
package cmc

import (
	"github.com/aviate-labs/agent-go"
	"github.com/aviate-labs/agent-go/candid/idl"
	"github.com/aviate-labs/agent-go/principal"
)

type AccountIdentifier = string

// Agent is a client for the "cmc" canister.
type Agent struct {
	*agent.Agent
	CanisterId principal.Principal
}

// NewAgent creates a new agent for the "cmc" canister.
func NewAgent(canisterId principal.Principal, config agent.Config) (*Agent, error) {
	a, err := agent.New(config)
	if err != nil {
		return nil, err
	}
	return &Agent{
		Agent:      a,
		CanisterId: canisterId,
	}, nil
}

// CreateCanister calls the "create_canister" method on the "cmc" canister.
func (a Agent) CreateCanister(arg0 CreateCanisterArg) (*CreateCanisterResult, error) {
	var r0 CreateCanisterResult
	if err := a.Agent.Call(
		a.CanisterId,
		"create_canister",
		[]any{arg0},
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

// GetBuildMetadata calls the "get_build_metadata" method on the "cmc" canister.
func (a Agent) GetBuildMetadata() (*string, error) {
	var r0 string
	if err := a.Agent.Query(
		a.CanisterId,
		"get_build_metadata",
		[]any{},
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

// GetIcpXdrConversionRate calls the "get_icp_xdr_conversion_rate" method on the "cmc" canister.
func (a Agent) GetIcpXdrConversionRate() (*IcpXdrConversionRateResponse, error) {
	var r0 IcpXdrConversionRateResponse
	if err := a.Agent.Query(
		a.CanisterId,
		"get_icp_xdr_conversion_rate",
		[]any{},
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

// GetPrincipalsAuthorizedToCreateCanistersToSubnets calls the "get_principals_authorized_to_create_canisters_to_subnets" method on the "cmc" canister.
func (a Agent) GetPrincipalsAuthorizedToCreateCanistersToSubnets() (*PrincipalsAuthorizedToCreateCanistersToSubnetsResponse, error) {
	var r0 PrincipalsAuthorizedToCreateCanistersToSubnetsResponse
	if err := a.Agent.Query(
		a.CanisterId,
		"get_principals_authorized_to_create_canisters_to_subnets",
		[]any{},
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

// GetSubnetTypesToSubnets calls the "get_subnet_types_to_subnets" method on the "cmc" canister.
func (a Agent) GetSubnetTypesToSubnets() (*SubnetTypesToSubnetsResponse, error) {
	var r0 SubnetTypesToSubnetsResponse
	if err := a.Agent.Query(
		a.CanisterId,
		"get_subnet_types_to_subnets",
		[]any{},
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

// NotifyCreateCanister calls the "notify_create_canister" method on the "cmc" canister.
func (a Agent) NotifyCreateCanister(arg0 NotifyCreateCanisterArg) (*NotifyCreateCanisterResult, error) {
	var r0 NotifyCreateCanisterResult
	if err := a.Agent.Call(
		a.CanisterId,
		"notify_create_canister",
		[]any{arg0},
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

// NotifyMintCycles calls the "notify_mint_cycles" method on the "cmc" canister.
func (a Agent) NotifyMintCycles(arg0 NotifyMintCyclesArg) (*NotifyMintCyclesResult, error) {
	var r0 NotifyMintCyclesResult
	if err := a.Agent.Call(
		a.CanisterId,
		"notify_mint_cycles",
		[]any{arg0},
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

// NotifyTopUp calls the "notify_top_up" method on the "cmc" canister.
func (a Agent) NotifyTopUp(arg0 NotifyTopUpArg) (*NotifyTopUpResult, error) {
	var r0 NotifyTopUpResult
	if err := a.Agent.Call(
		a.CanisterId,
		"notify_top_up",
		[]any{arg0},
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

type BlockIndex = uint64

type CanisterSettings struct {
	Controller          *principal.Principal   `ic:"controller,omitempty" json:"controller,omitempty"`
	Controllers         *[]principal.Principal `ic:"controllers,omitempty" json:"controllers,omitempty"`
	ComputeAllocation   *idl.Nat               `ic:"compute_allocation,omitempty" json:"compute_allocation,omitempty"`
	MemoryAllocation    *idl.Nat               `ic:"memory_allocation,omitempty" json:"memory_allocation,omitempty"`
	FreezingThreshold   *idl.Nat               `ic:"freezing_threshold,omitempty" json:"freezing_threshold,omitempty"`
	ReservedCyclesLimit *idl.Nat               `ic:"reserved_cycles_limit,omitempty" json:"reserved_cycles_limit,omitempty"`
	LogVisibility       *LogVisibility         `ic:"log_visibility,omitempty" json:"log_visibility,omitempty"`
	WasmMemoryLimit     *idl.Nat               `ic:"wasm_memory_limit,omitempty" json:"wasm_memory_limit,omitempty"`
}

type CreateCanisterArg struct {
	Settings        *CanisterSettings `ic:"settings,omitempty" json:"settings,omitempty"`
	SubnetType      *string           `ic:"subnet_type,omitempty" json:"subnet_type,omitempty"`
	SubnetSelection *SubnetSelection  `ic:"subnet_selection,omitempty" json:"subnet_selection,omitempty"`
}

type CreateCanisterError struct {
	Refunded *struct {
		RefundAmount idl.Nat `ic:"refund_amount" json:"refund_amount"`
		CreateError  string  `ic:"create_error" json:"create_error"`
	} `ic:"Refunded,variant"`
	RefundFailed *struct {
		CreateError string `ic:"create_error" json:"create_error"`
		RefundError string `ic:"refund_error" json:"refund_error"`
	} `ic:"RefundFailed,variant"`
}

type CreateCanisterResult struct {
	Ok  *principal.Principal `ic:"Ok,variant"`
	Err *CreateCanisterError `ic:"Err,variant"`
}

type Cycles = idl.Nat

type CyclesCanisterInitPayload struct {
	LedgerCanisterId       *principal.Principal  `ic:"ledger_canister_id,omitempty" json:"ledger_canister_id,omitempty"`
	GovernanceCanisterId   *principal.Principal  `ic:"governance_canister_id,omitempty" json:"governance_canister_id,omitempty"`
	MintingAccountId       *AccountIdentifier    `ic:"minting_account_id,omitempty" json:"minting_account_id,omitempty"`
	LastPurgedNotification *uint64               `ic:"last_purged_notification,omitempty" json:"last_purged_notification,omitempty"`
	ExchangeRateCanister   *ExchangeRateCanister `ic:"exchange_rate_canister,omitempty" json:"exchange_rate_canister,omitempty"`
	CyclesLedgerCanisterId *principal.Principal  `ic:"cycles_ledger_canister_id,omitempty" json:"cycles_ledger_canister_id,omitempty"`
}

type ExchangeRateCanister struct {
	Set   *principal.Principal `ic:"Set,variant"`
	Unset *idl.Null            `ic:"Unset,variant"`
}

type IcpXdrConversionRate struct {
	TimestampSeconds   uint64 `ic:"timestamp_seconds" json:"timestamp_seconds"`
	XdrPermyriadPerIcp uint64 `ic:"xdr_permyriad_per_icp" json:"xdr_permyriad_per_icp"`
}

type IcpXdrConversionRateResponse struct {
	Data        IcpXdrConversionRate `ic:"data" json:"data"`
	HashTree    []byte               `ic:"hash_tree" json:"hash_tree"`
	Certificate []byte               `ic:"certificate" json:"certificate"`
}

type LogVisibility struct {
	Controllers *idl.Null `ic:"controllers,variant"`
	Public      *idl.Null `ic:"public,variant"`
}

type Memo = *[]byte

type NotifyCreateCanisterArg struct {
	BlockIndex      BlockIndex          `ic:"block_index" json:"block_index"`
	Controller      principal.Principal `ic:"controller" json:"controller"`
	SubnetType      *string             `ic:"subnet_type,omitempty" json:"subnet_type,omitempty"`
	SubnetSelection *SubnetSelection    `ic:"subnet_selection,omitempty" json:"subnet_selection,omitempty"`
	Settings        *CanisterSettings   `ic:"settings,omitempty" json:"settings,omitempty"`
}

type NotifyCreateCanisterResult struct {
	Ok  *principal.Principal `ic:"Ok,variant"`
	Err *NotifyError         `ic:"Err,variant"`
}

type NotifyError struct {
	Refunded *struct {
		Reason     string      `ic:"reason" json:"reason"`
		BlockIndex *BlockIndex `ic:"block_index,omitempty" json:"block_index,omitempty"`
	} `ic:"Refunded,variant"`
	Processing         *idl.Null   `ic:"Processing,variant"`
	TransactionTooOld  *BlockIndex `ic:"TransactionTooOld,variant"`
	InvalidTransaction *string     `ic:"InvalidTransaction,variant"`
	Other              *struct {
		ErrorCode    uint64 `ic:"error_code" json:"error_code"`
		ErrorMessage string `ic:"error_message" json:"error_message"`
	} `ic:"Other,variant"`
}

type NotifyMintCyclesArg struct {
	BlockIndex   BlockIndex `ic:"block_index" json:"block_index"`
	ToSubaccount Subaccount `ic:"to_subaccount" json:"to_subaccount"`
	DepositMemo  Memo       `ic:"deposit_memo" json:"deposit_memo"`
}

type NotifyMintCyclesResult struct {
	Ok  *NotifyMintCyclesSuccess `ic:"Ok,variant"`
	Err *NotifyError             `ic:"Err,variant"`
}

type NotifyMintCyclesSuccess struct {
	BlockIndex idl.Nat `ic:"block_index" json:"block_index"`
	Minted     idl.Nat `ic:"minted" json:"minted"`
	Balance    idl.Nat `ic:"balance" json:"balance"`
}

type NotifyTopUpArg struct {
	BlockIndex BlockIndex          `ic:"block_index" json:"block_index"`
	CanisterId principal.Principal `ic:"canister_id" json:"canister_id"`
}

type NotifyTopUpResult struct {
	Ok  *Cycles      `ic:"Ok,variant"`
	Err *NotifyError `ic:"Err,variant"`
}

type PrincipalsAuthorizedToCreateCanistersToSubnetsResponse struct {
	Data []struct {
		Field0 principal.Principal   `ic:"0" json:"0"`
		Field1 []principal.Principal `ic:"1" json:"1"`
	} `ic:"data" json:"data"`
}

type Subaccount = *[]byte

type SubnetFilter struct {
	SubnetType *string `ic:"subnet_type,omitempty" json:"subnet_type,omitempty"`
}

type SubnetSelection struct {
	Subnet *struct {
		Subnet principal.Principal `ic:"subnet" json:"subnet"`
	} `ic:"Subnet,variant"`
	Filter *SubnetFilter `ic:"Filter,variant"`
}

type SubnetTypesToSubnetsResponse struct {
	Data []struct {
		Field0 string                `ic:"0" json:"0"`
		Field1 []principal.Principal `ic:"1" json:"1"`
	} `ic:"data" json:"data"`
}
