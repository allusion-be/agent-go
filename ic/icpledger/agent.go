// Package icpledger provides a client for the "icpledger" canister.
// Do NOT edit this file. It was automatically generated by https://github.com/aviate-labs/agent-go.
package icpledger

import (
	"github.com/aviate-labs/agent-go"
	"github.com/aviate-labs/agent-go/candid/idl"
	"github.com/aviate-labs/agent-go/principal"
)

type AccountBalanceArgs struct {
	Account AccountIdentifier `ic:"account"`
}

type AccountIdentifier = []byte

// Agent is a client for the "icpledger" canister.
type Agent struct {
	a          *agent.Agent
	canisterId principal.Principal
}

// NewAgent creates a new agent for the "icpledger" canister.
func NewAgent(canisterId principal.Principal, config agent.Config) (*Agent, error) {
	a, err := agent.New(config)
	if err != nil {
		return nil, err
	}
	return &Agent{
		a:          a,
		canisterId: canisterId,
	}, nil
}

// AccountBalance calls the "account_balance" method on the "icpledger" canister.
func (a Agent) AccountBalance(arg0 AccountBalanceArgs) (*Tokens, error) {
	var r0 Tokens
	if err := a.a.Query(
		a.canisterId,
		"account_balance",
		[]any{arg0},
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

// Archives calls the "archives" method on the "icpledger" canister.
func (a Agent) Archives() (*Archives, error) {
	var r0 Archives
	if err := a.a.Query(
		a.canisterId,
		"archives",
		[]any{},
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

// Decimals calls the "decimals" method on the "icpledger" canister.
func (a Agent) Decimals() (*struct {
	Decimals uint32 `ic:"decimals"`
}, error) {
	var r0 struct {
		Decimals uint32 `ic:"decimals"`
	}
	if err := a.a.Query(
		a.canisterId,
		"decimals",
		[]any{},
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

// Name calls the "name" method on the "icpledger" canister.
func (a Agent) Name() (*struct {
	Name string `ic:"name"`
}, error) {
	var r0 struct {
		Name string `ic:"name"`
	}
	if err := a.a.Query(
		a.canisterId,
		"name",
		[]any{},
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

// QueryBlocks calls the "query_blocks" method on the "icpledger" canister.
func (a Agent) QueryBlocks(arg0 GetBlocksArgs) (*QueryBlocksResponse, error) {
	var r0 QueryBlocksResponse
	if err := a.a.Query(
		a.canisterId,
		"query_blocks",
		[]any{arg0},
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

// Symbol calls the "symbol" method on the "icpledger" canister.
func (a Agent) Symbol() (*struct {
	Symbol string `ic:"symbol"`
}, error) {
	var r0 struct {
		Symbol string `ic:"symbol"`
	}
	if err := a.a.Query(
		a.canisterId,
		"symbol",
		[]any{},
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

// Transfer calls the "transfer" method on the "icpledger" canister.
func (a Agent) Transfer(arg0 TransferArgs) (*TransferResult, error) {
	var r0 TransferResult
	if err := a.a.Call(
		a.canisterId,
		"transfer",
		[]any{arg0},
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

// TransferFee calls the "transfer_fee" method on the "icpledger" canister.
func (a Agent) TransferFee(arg0 TransferFeeArg) (*TransferFee, error) {
	var r0 TransferFee
	if err := a.a.Query(
		a.canisterId,
		"transfer_fee",
		[]any{arg0},
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

type Archive struct {
	CanisterId principal.Principal `ic:"canister_id"`
}

type Archives struct {
	Archives []Archive `ic:"archives"`
}

type Block struct {
	ParentHash  *[]byte     `ic:"parent_hash,omitempty"`
	Transaction Transaction `ic:"transaction"`
	Timestamp   TimeStamp   `ic:"timestamp"`
}

type BlockIndex = uint64

type BlockRange struct {
	Blocks []Block `ic:"blocks"`
}

type GetBlocksArgs struct {
	Start  BlockIndex `ic:"start"`
	Length uint64     `ic:"length"`
}

type Memo = uint64

type Operation struct {
	Mint *struct {
		To     AccountIdentifier `ic:"to"`
		Amount Tokens            `ic:"amount"`
	} `ic:"Mint,variant"`
	Burn *struct {
		From   AccountIdentifier `ic:"from"`
		Amount Tokens            `ic:"amount"`
	} `ic:"Burn,variant"`
	Transfer *struct {
		From   AccountIdentifier `ic:"from"`
		To     AccountIdentifier `ic:"to"`
		Amount Tokens            `ic:"amount"`
		Fee    Tokens            `ic:"fee"`
	} `ic:"Transfer,variant"`
	Approve *struct {
		From         AccountIdentifier `ic:"from"`
		Spender      AccountIdentifier `ic:"spender"`
		AllowanceE8s idl.Int           `ic:"allowance_e8s"`
		Fee          Tokens            `ic:"fee"`
		ExpiresAt    *TimeStamp        `ic:"expires_at,omitempty"`
	} `ic:"Approve,variant"`
	TransferFrom *struct {
		From    AccountIdentifier `ic:"from"`
		To      AccountIdentifier `ic:"to"`
		Spender AccountIdentifier `ic:"spender"`
		Amount  Tokens            `ic:"amount"`
		Fee     Tokens            `ic:"fee"`
	} `ic:"TransferFrom,variant"`
}

type QueryArchiveError struct {
	BadFirstBlockIndex *struct {
		RequestedIndex  BlockIndex `ic:"requested_index"`
		FirstValidIndex BlockIndex `ic:"first_valid_index"`
	} `ic:"BadFirstBlockIndex,variant"`
	Other *struct {
		ErrorCode    uint64 `ic:"error_code"`
		ErrorMessage string `ic:"error_message"`
	} `ic:"Other,variant"`
}

type QueryArchiveFn struct { /* NOT SUPPORTED */
}

type QueryArchiveResult struct {
	Ok  *BlockRange        `ic:"Ok,variant"`
	Err *QueryArchiveError `ic:"Err,variant"`
}

type QueryBlocksResponse struct {
	ChainLength     uint64     `ic:"chain_length"`
	Certificate     *[]byte    `ic:"certificate,omitempty"`
	Blocks          []Block    `ic:"blocks"`
	FirstBlockIndex BlockIndex `ic:"first_block_index"`
	ArchivedBlocks  []struct {
		Start    BlockIndex     `ic:"start"`
		Length   uint64         `ic:"length"`
		Callback QueryArchiveFn `ic:"callback"`
	} `ic:"archived_blocks"`
}

type SubAccount = []byte

type TimeStamp struct {
	TimestampNanos uint64 `ic:"timestamp_nanos"`
}

type Tokens struct {
	E8s uint64 `ic:"e8s"`
}

type Transaction struct {
	Memo          Memo       `ic:"memo"`
	Icrc1Memo     *[]byte    `ic:"icrc1_memo,omitempty"`
	Operation     *Operation `ic:"operation,omitempty"`
	CreatedAtTime TimeStamp  `ic:"created_at_time"`
}

type TransferArgs struct {
	Memo           Memo              `ic:"memo"`
	Amount         Tokens            `ic:"amount"`
	Fee            Tokens            `ic:"fee"`
	FromSubaccount *SubAccount       `ic:"from_subaccount,omitempty"`
	To             AccountIdentifier `ic:"to"`
	CreatedAtTime  *TimeStamp        `ic:"created_at_time,omitempty"`
}

type TransferError struct {
	BadFee *struct {
		ExpectedFee Tokens `ic:"expected_fee"`
	} `ic:"BadFee,variant"`
	InsufficientFunds *struct {
		Balance Tokens `ic:"balance"`
	} `ic:"InsufficientFunds,variant"`
	TxTooOld *struct {
		AllowedWindowNanos uint64 `ic:"allowed_window_nanos"`
	} `ic:"TxTooOld,variant"`
	TxCreatedInFuture *idl.Null `ic:"TxCreatedInFuture,variant"`
	TxDuplicate       *struct {
		DuplicateOf BlockIndex `ic:"duplicate_of"`
	} `ic:"TxDuplicate,variant"`
}

type TransferFee struct {
	TransferFee Tokens `ic:"transfer_fee"`
}

type TransferFeeArg struct {
}

type TransferResult struct {
	Ok  *BlockIndex    `ic:"Ok,variant"`
	Err *TransferError `ic:"Err,variant"`
}
