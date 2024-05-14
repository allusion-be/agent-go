// Package icparchive provides a client for the "icparchive" canister.
// Do NOT edit this file. It was automatically generated by https://github.com/aviate-labs/agent-go.
package icparchive

import (
	"github.com/aviate-labs/agent-go"
	"github.com/aviate-labs/agent-go/candid/idl"
	"github.com/aviate-labs/agent-go/principal"
)

type AccountIdentifier = []byte

// Agent is a client for the "icparchive" canister.
type Agent struct {
	*agent.Agent
	CanisterId principal.Principal
}

// NewAgent creates a new agent for the "icparchive" canister.
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

// GetBlocks calls the "get_blocks" method on the "icparchive" canister.
func (a Agent) GetBlocks(arg0 GetBlocksArgs) (*GetBlocksResult, error) {
	var r0 GetBlocksResult
	if err := a.Agent.Query(
		a.CanisterId,
		"get_blocks",
		[]any{arg0},
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

// GetEncodedBlocks calls the "get_encoded_blocks" method on the "icparchive" canister.
func (a Agent) GetEncodedBlocks(arg0 GetBlocksArgs) (*GetEncodedBlocksResult, error) {
	var r0 GetEncodedBlocksResult
	if err := a.Agent.Query(
		a.CanisterId,
		"get_encoded_blocks",
		[]any{arg0},
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

type Block struct {
	ParentHash  *[]byte     `ic:"parent_hash,omitempty" json:"parent_hash,omitempty"`
	Transaction Transaction `ic:"transaction" json:"transaction"`
	Timestamp   Timestamp   `ic:"timestamp" json:"timestamp"`
}

type BlockIndex = uint64

type BlockRange struct {
	Blocks []Block `ic:"blocks" json:"blocks"`
}

type GetBlocksArgs struct {
	Start  BlockIndex `ic:"start" json:"start"`
	Length uint64     `ic:"length" json:"length"`
}

type GetBlocksError struct {
	BadFirstBlockIndex *struct {
		RequestedIndex  BlockIndex `ic:"requested_index" json:"requested_index"`
		FirstValidIndex BlockIndex `ic:"first_valid_index" json:"first_valid_index"`
	} `ic:"BadFirstBlockIndex,variant"`
	Other *struct {
		ErrorCode    uint64 `ic:"error_code" json:"error_code"`
		ErrorMessage string `ic:"error_message" json:"error_message"`
	} `ic:"Other,variant"`
}

type GetBlocksResult struct {
	Ok  *BlockRange     `ic:"Ok,variant"`
	Err *GetBlocksError `ic:"Err,variant"`
}

type GetEncodedBlocksResult struct {
	Ok  *[][]byte       `ic:"Ok,variant"`
	Err *GetBlocksError `ic:"Err,variant"`
}

type Memo = uint64

type Operation struct {
	Mint *struct {
		To     AccountIdentifier `ic:"to" json:"to"`
		Amount Tokens            `ic:"amount" json:"amount"`
	} `ic:"Mint,variant"`
	Burn *struct {
		From    AccountIdentifier  `ic:"from" json:"from"`
		Spender *AccountIdentifier `ic:"spender,omitempty" json:"spender,omitempty"`
		Amount  Tokens             `ic:"amount" json:"amount"`
	} `ic:"Burn,variant"`
	Transfer *struct {
		From    AccountIdentifier `ic:"from" json:"from"`
		To      AccountIdentifier `ic:"to" json:"to"`
		Amount  Tokens            `ic:"amount" json:"amount"`
		Fee     Tokens            `ic:"fee" json:"fee"`
		Spender *[]uint8          `ic:"spender,omitempty" json:"spender,omitempty"`
	} `ic:"Transfer,variant"`
	Approve *struct {
		From              AccountIdentifier `ic:"from" json:"from"`
		Spender           AccountIdentifier `ic:"spender" json:"spender"`
		AllowanceE8s      idl.Int           `ic:"allowance_e8s" json:"allowance_e8s"`
		Allowance         Tokens            `ic:"allowance" json:"allowance"`
		Fee               Tokens            `ic:"fee" json:"fee"`
		ExpiresAt         *Timestamp        `ic:"expires_at,omitempty" json:"expires_at,omitempty"`
		ExpectedAllowance *Tokens           `ic:"expected_allowance,omitempty" json:"expected_allowance,omitempty"`
	} `ic:"Approve,variant"`
}

type Timestamp struct {
	TimestampNanos uint64 `ic:"timestamp_nanos" json:"timestamp_nanos"`
}

type Tokens struct {
	E8s uint64 `ic:"e8s" json:"e8s"`
}

type Transaction struct {
	Memo          Memo       `ic:"memo" json:"memo"`
	Icrc1Memo     *[]byte    `ic:"icrc1_memo,omitempty" json:"icrc1_memo,omitempty"`
	Operation     *Operation `ic:"operation,omitempty" json:"operation,omitempty"`
	CreatedAtTime Timestamp  `ic:"created_at_time" json:"created_at_time"`
}
