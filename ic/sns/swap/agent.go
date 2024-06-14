// Package swap provides a client for the "swap" canister.
// Do NOT edit this file. It was automatically generated by https://github.com/aviate-labs/agent-go.
package swap

import (
	"github.com/aviate-labs/agent-go"
	"github.com/aviate-labs/agent-go/candid/idl"
	"github.com/aviate-labs/agent-go/principal"
)

// Agent is a client for the "swap" canister.
type Agent struct {
	*agent.Agent
	CanisterId principal.Principal
}

// NewAgent creates a new agent for the "swap" canister.
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

// ErrorRefundIcp calls the "error_refund_icp" method on the "swap" canister.
func (a Agent) ErrorRefundIcp(arg0 ErrorRefundIcpRequest) (*ErrorRefundIcpResponse, error) {
	var r0 ErrorRefundIcpResponse
	if err := a.Agent.Call(
		a.CanisterId,
		"error_refund_icp",
		[]any{arg0},
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

// FinalizeSwap calls the "finalize_swap" method on the "swap" canister.
func (a Agent) FinalizeSwap(arg0 struct {
}) (*FinalizeSwapResponse, error) {
	var r0 FinalizeSwapResponse
	if err := a.Agent.Call(
		a.CanisterId,
		"finalize_swap",
		[]any{arg0},
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

// GetAutoFinalizationStatus calls the "get_auto_finalization_status" method on the "swap" canister.
func (a Agent) GetAutoFinalizationStatus(arg0 struct {
}) (*GetAutoFinalizationStatusResponse, error) {
	var r0 GetAutoFinalizationStatusResponse
	if err := a.Agent.Query(
		a.CanisterId,
		"get_auto_finalization_status",
		[]any{arg0},
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

// GetBuyerState calls the "get_buyer_state" method on the "swap" canister.
func (a Agent) GetBuyerState(arg0 GetBuyerStateRequest) (*GetBuyerStateResponse, error) {
	var r0 GetBuyerStateResponse
	if err := a.Agent.Query(
		a.CanisterId,
		"get_buyer_state",
		[]any{arg0},
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

// GetBuyersTotal calls the "get_buyers_total" method on the "swap" canister.
func (a Agent) GetBuyersTotal(arg0 struct {
}) (*GetBuyersTotalResponse, error) {
	var r0 GetBuyersTotalResponse
	if err := a.Agent.Call(
		a.CanisterId,
		"get_buyers_total",
		[]any{arg0},
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

// GetCanisterStatus calls the "get_canister_status" method on the "swap" canister.
func (a Agent) GetCanisterStatus(arg0 struct {
}) (*CanisterStatusResultV2, error) {
	var r0 CanisterStatusResultV2
	if err := a.Agent.Call(
		a.CanisterId,
		"get_canister_status",
		[]any{arg0},
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

// GetDerivedState calls the "get_derived_state" method on the "swap" canister.
func (a Agent) GetDerivedState(arg0 struct {
}) (*GetDerivedStateResponse, error) {
	var r0 GetDerivedStateResponse
	if err := a.Agent.Query(
		a.CanisterId,
		"get_derived_state",
		[]any{arg0},
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

// GetInit calls the "get_init" method on the "swap" canister.
func (a Agent) GetInit(arg0 struct {
}) (*GetInitResponse, error) {
	var r0 GetInitResponse
	if err := a.Agent.Query(
		a.CanisterId,
		"get_init",
		[]any{arg0},
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

// GetLifecycle calls the "get_lifecycle" method on the "swap" canister.
func (a Agent) GetLifecycle(arg0 struct {
}) (*GetLifecycleResponse, error) {
	var r0 GetLifecycleResponse
	if err := a.Agent.Query(
		a.CanisterId,
		"get_lifecycle",
		[]any{arg0},
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

// GetOpenTicket calls the "get_open_ticket" method on the "swap" canister.
func (a Agent) GetOpenTicket(arg0 struct {
}) (*GetOpenTicketResponse, error) {
	var r0 GetOpenTicketResponse
	if err := a.Agent.Query(
		a.CanisterId,
		"get_open_ticket",
		[]any{arg0},
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

// GetSaleParameters calls the "get_sale_parameters" method on the "swap" canister.
func (a Agent) GetSaleParameters(arg0 struct {
}) (*GetSaleParametersResponse, error) {
	var r0 GetSaleParametersResponse
	if err := a.Agent.Query(
		a.CanisterId,
		"get_sale_parameters",
		[]any{arg0},
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

// GetState calls the "get_state" method on the "swap" canister.
func (a Agent) GetState(arg0 struct {
}) (*GetStateResponse, error) {
	var r0 GetStateResponse
	if err := a.Agent.Query(
		a.CanisterId,
		"get_state",
		[]any{arg0},
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

// ListCommunityFundParticipants calls the "list_community_fund_participants" method on the "swap" canister.
func (a Agent) ListCommunityFundParticipants(arg0 ListCommunityFundParticipantsRequest) (*NeuronsFundParticipants, error) {
	var r0 NeuronsFundParticipants
	if err := a.Agent.Query(
		a.CanisterId,
		"list_community_fund_participants",
		[]any{arg0},
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

// ListDirectParticipants calls the "list_direct_participants" method on the "swap" canister.
func (a Agent) ListDirectParticipants(arg0 ListDirectParticipantsRequest) (*ListDirectParticipantsResponse, error) {
	var r0 ListDirectParticipantsResponse
	if err := a.Agent.Query(
		a.CanisterId,
		"list_direct_participants",
		[]any{arg0},
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

// ListSnsNeuronRecipes calls the "list_sns_neuron_recipes" method on the "swap" canister.
func (a Agent) ListSnsNeuronRecipes(arg0 ListSnsNeuronRecipesRequest) (*ListSnsNeuronRecipesResponse, error) {
	var r0 ListSnsNeuronRecipesResponse
	if err := a.Agent.Query(
		a.CanisterId,
		"list_sns_neuron_recipes",
		[]any{arg0},
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

// NewSaleTicket calls the "new_sale_ticket" method on the "swap" canister.
func (a Agent) NewSaleTicket(arg0 NewSaleTicketRequest) (*NewSaleTicketResponse, error) {
	var r0 NewSaleTicketResponse
	if err := a.Agent.Call(
		a.CanisterId,
		"new_sale_ticket",
		[]any{arg0},
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

// NotifyPaymentFailure calls the "notify_payment_failure" method on the "swap" canister.
func (a Agent) NotifyPaymentFailure(arg0 struct {
}) (*Ok2, error) {
	var r0 Ok2
	if err := a.Agent.Call(
		a.CanisterId,
		"notify_payment_failure",
		[]any{arg0},
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

// Open calls the "open" method on the "swap" canister.
func (a Agent) Open(arg0 OpenRequest) (*struct {
}, error) {
	var r0 struct {
	}
	if err := a.Agent.Call(
		a.CanisterId,
		"open",
		[]any{arg0},
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

// RefreshBuyerTokens calls the "refresh_buyer_tokens" method on the "swap" canister.
func (a Agent) RefreshBuyerTokens(arg0 RefreshBuyerTokensRequest) (*RefreshBuyerTokensResponse, error) {
	var r0 RefreshBuyerTokensResponse
	if err := a.Agent.Call(
		a.CanisterId,
		"refresh_buyer_tokens",
		[]any{arg0},
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

type BuyerState struct {
	Icp                     *TransferableAmount `ic:"icp,omitempty" json:"icp,omitempty"`
	HasCreatedNeuronRecipes *bool               `ic:"has_created_neuron_recipes,omitempty" json:"has_created_neuron_recipes,omitempty"`
}

type CanisterCallError struct {
	Code        *int32 `ic:"code,omitempty" json:"code,omitempty"`
	Description string `ic:"description" json:"description"`
}

type CanisterStatusResultV2 struct {
	Status                 CanisterStatusType           `ic:"status" json:"status"`
	MemorySize             idl.Nat                      `ic:"memory_size" json:"memory_size"`
	Cycles                 idl.Nat                      `ic:"cycles" json:"cycles"`
	Settings               DefiniteCanisterSettingsArgs `ic:"settings" json:"settings"`
	IdleCyclesBurnedPerDay idl.Nat                      `ic:"idle_cycles_burned_per_day" json:"idle_cycles_burned_per_day"`
	ModuleHash             *[]byte                      `ic:"module_hash,omitempty" json:"module_hash,omitempty"`
}

type CanisterStatusType struct {
	Stopped  *idl.Null `ic:"stopped,variant"`
	Stopping *idl.Null `ic:"stopping,variant"`
	Running  *idl.Null `ic:"running,variant"`
}

type CfInvestment struct {
	HotkeyPrincipal string `ic:"hotkey_principal" json:"hotkey_principal"`
	NnsNeuronId     uint64 `ic:"nns_neuron_id" json:"nns_neuron_id"`
}

type CfNeuron struct {
	HasCreatedNeuronRecipes *bool  `ic:"has_created_neuron_recipes,omitempty" json:"has_created_neuron_recipes,omitempty"`
	NnsNeuronId             uint64 `ic:"nns_neuron_id" json:"nns_neuron_id"`
	AmountIcpE8s            uint64 `ic:"amount_icp_e8s" json:"amount_icp_e8s"`
}

type CfParticipant struct {
	HotkeyPrincipal string     `ic:"hotkey_principal" json:"hotkey_principal"`
	CfNeurons       []CfNeuron `ic:"cf_neurons" json:"cf_neurons"`
}

type Countries struct {
	IsoCodes []string `ic:"iso_codes" json:"iso_codes"`
}

type DefiniteCanisterSettingsArgs struct {
	FreezingThreshold idl.Nat               `ic:"freezing_threshold" json:"freezing_threshold"`
	Controllers       []principal.Principal `ic:"controllers" json:"controllers"`
	MemoryAllocation  idl.Nat               `ic:"memory_allocation" json:"memory_allocation"`
	ComputeAllocation idl.Nat               `ic:"compute_allocation" json:"compute_allocation"`
}

type DerivedState struct {
	SnsTokensPerIcp                float32 `ic:"sns_tokens_per_icp" json:"sns_tokens_per_icp"`
	BuyerTotalIcpE8s               uint64  `ic:"buyer_total_icp_e8s" json:"buyer_total_icp_e8s"`
	CfParticipantCount             *uint64 `ic:"cf_participant_count,omitempty" json:"cf_participant_count,omitempty"`
	NeuronsFundParticipationIcpE8s *uint64 `ic:"neurons_fund_participation_icp_e8s,omitempty" json:"neurons_fund_participation_icp_e8s,omitempty"`
	DirectParticipationIcpE8s      *uint64 `ic:"direct_participation_icp_e8s,omitempty" json:"direct_participation_icp_e8s,omitempty"`
	DirectParticipantCount         *uint64 `ic:"direct_participant_count,omitempty" json:"direct_participant_count,omitempty"`
	CfNeuronCount                  *uint64 `ic:"cf_neuron_count,omitempty" json:"cf_neuron_count,omitempty"`
}

type DirectInvestment struct {
	BuyerPrincipal string `ic:"buyer_principal" json:"buyer_principal"`
}

type Err struct {
	Description *string `ic:"description,omitempty" json:"description,omitempty"`
	ErrorType   *int32  `ic:"error_type,omitempty" json:"error_type,omitempty"`
}

type Err1 struct {
	ErrorType *int32 `ic:"error_type,omitempty" json:"error_type,omitempty"`
}

type Err2 struct {
	InvalidUserAmount *InvalidUserAmount `ic:"invalid_user_amount,omitempty" json:"invalid_user_amount,omitempty"`
	ExistingTicket    *Ticket            `ic:"existing_ticket,omitempty" json:"existing_ticket,omitempty"`
	ErrorType         int32              `ic:"error_type" json:"error_type"`
}

type Error struct {
	Message *string `ic:"message,omitempty" json:"message,omitempty"`
}

type ErrorRefundIcpRequest struct {
	SourcePrincipalId *principal.Principal `ic:"source_principal_id,omitempty" json:"source_principal_id,omitempty"`
}

type ErrorRefundIcpResponse struct {
	Result *Result `ic:"result,omitempty" json:"result,omitempty"`
}

type FailedUpdate struct {
	Err            *CanisterCallError   `ic:"err,omitempty" json:"err,omitempty"`
	DappCanisterId *principal.Principal `ic:"dapp_canister_id,omitempty" json:"dapp_canister_id,omitempty"`
}

type FinalizeSwapResponse struct {
	SetDappControllersCallResult           *SetDappControllersCallResult           `ic:"set_dapp_controllers_call_result,omitempty" json:"set_dapp_controllers_call_result,omitempty"`
	CreateSnsNeuronRecipesResult           *SweepResult                            `ic:"create_sns_neuron_recipes_result,omitempty" json:"create_sns_neuron_recipes_result,omitempty"`
	SettleCommunityFundParticipationResult *SettleCommunityFundParticipationResult `ic:"settle_community_fund_participation_result,omitempty" json:"settle_community_fund_participation_result,omitempty"`
	ErrorMessage                           *string                                 `ic:"error_message,omitempty" json:"error_message,omitempty"`
	SettleNeuronsFundParticipationResult   *SettleNeuronsFundParticipationResult   `ic:"settle_neurons_fund_participation_result,omitempty" json:"settle_neurons_fund_participation_result,omitempty"`
	SetModeCallResult                      *SetModeCallResult                      `ic:"set_mode_call_result,omitempty" json:"set_mode_call_result,omitempty"`
	SweepIcpResult                         *SweepResult                            `ic:"sweep_icp_result,omitempty" json:"sweep_icp_result,omitempty"`
	ClaimNeuronResult                      *SweepResult                            `ic:"claim_neuron_result,omitempty" json:"claim_neuron_result,omitempty"`
	SweepSnsResult                         *SweepResult                            `ic:"sweep_sns_result,omitempty" json:"sweep_sns_result,omitempty"`
}

type GetAutoFinalizationStatusResponse struct {
	AutoFinalizeSwapResponse     *FinalizeSwapResponse `ic:"auto_finalize_swap_response,omitempty" json:"auto_finalize_swap_response,omitempty"`
	HasAutoFinalizeBeenAttempted *bool                 `ic:"has_auto_finalize_been_attempted,omitempty" json:"has_auto_finalize_been_attempted,omitempty"`
	IsAutoFinalizeEnabled        *bool                 `ic:"is_auto_finalize_enabled,omitempty" json:"is_auto_finalize_enabled,omitempty"`
}

type GetBuyerStateRequest struct {
	PrincipalId *principal.Principal `ic:"principal_id,omitempty" json:"principal_id,omitempty"`
}

type GetBuyerStateResponse struct {
	BuyerState *BuyerState `ic:"buyer_state,omitempty" json:"buyer_state,omitempty"`
}

type GetBuyersTotalResponse struct {
	BuyersTotal uint64 `ic:"buyers_total" json:"buyers_total"`
}

type GetDerivedStateResponse struct {
	SnsTokensPerIcp                *float64 `ic:"sns_tokens_per_icp,omitempty" json:"sns_tokens_per_icp,omitempty"`
	BuyerTotalIcpE8s               *uint64  `ic:"buyer_total_icp_e8s,omitempty" json:"buyer_total_icp_e8s,omitempty"`
	CfParticipantCount             *uint64  `ic:"cf_participant_count,omitempty" json:"cf_participant_count,omitempty"`
	NeuronsFundParticipationIcpE8s *uint64  `ic:"neurons_fund_participation_icp_e8s,omitempty" json:"neurons_fund_participation_icp_e8s,omitempty"`
	DirectParticipationIcpE8s      *uint64  `ic:"direct_participation_icp_e8s,omitempty" json:"direct_participation_icp_e8s,omitempty"`
	DirectParticipantCount         *uint64  `ic:"direct_participant_count,omitempty" json:"direct_participant_count,omitempty"`
	CfNeuronCount                  *uint64  `ic:"cf_neuron_count,omitempty" json:"cf_neuron_count,omitempty"`
}

type GetInitResponse struct {
	Init *Init `ic:"init,omitempty" json:"init,omitempty"`
}

type GetLifecycleResponse struct {
	DecentralizationSaleOpenTimestampSeconds        *uint64 `ic:"decentralization_sale_open_timestamp_seconds,omitempty" json:"decentralization_sale_open_timestamp_seconds,omitempty"`
	Lifecycle                                       *int32  `ic:"lifecycle,omitempty" json:"lifecycle,omitempty"`
	DecentralizationSwapTerminationTimestampSeconds *uint64 `ic:"decentralization_swap_termination_timestamp_seconds,omitempty" json:"decentralization_swap_termination_timestamp_seconds,omitempty"`
}

type GetOpenTicketResponse struct {
	Result *Result1 `ic:"result,omitempty" json:"result,omitempty"`
}

type GetSaleParametersResponse struct {
	Params *Params `ic:"params,omitempty" json:"params,omitempty"`
}

type GetStateResponse struct {
	Swap    *Swap         `ic:"swap,omitempty" json:"swap,omitempty"`
	Derived *DerivedState `ic:"derived,omitempty" json:"derived,omitempty"`
}

type GovernanceError struct {
	ErrorMessage string `ic:"error_message" json:"error_message"`
	ErrorType    int32  `ic:"error_type" json:"error_type"`
}

type Icrc1Account struct {
	Owner      *principal.Principal `ic:"owner,omitempty" json:"owner,omitempty"`
	Subaccount *[]byte              `ic:"subaccount,omitempty" json:"subaccount,omitempty"`
}

type IdealMatchedParticipationFunction struct {
	SerializedRepresentation *string `ic:"serialized_representation,omitempty" json:"serialized_representation,omitempty"`
}

type Init struct {
	NnsProposalId                       *uint64                              `ic:"nns_proposal_id,omitempty" json:"nns_proposal_id,omitempty"`
	SnsRootCanisterId                   string                               `ic:"sns_root_canister_id" json:"sns_root_canister_id"`
	NeuronsFundParticipation            *bool                                `ic:"neurons_fund_participation,omitempty" json:"neurons_fund_participation,omitempty"`
	MinParticipantIcpE8s                *uint64                              `ic:"min_participant_icp_e8s,omitempty" json:"min_participant_icp_e8s,omitempty"`
	NeuronBasketConstructionParameters  *NeuronBasketConstructionParameters  `ic:"neuron_basket_construction_parameters,omitempty" json:"neuron_basket_construction_parameters,omitempty"`
	FallbackControllerPrincipalIds      []string                             `ic:"fallback_controller_principal_ids" json:"fallback_controller_principal_ids"`
	MaxIcpE8s                           *uint64                              `ic:"max_icp_e8s,omitempty" json:"max_icp_e8s,omitempty"`
	NeuronMinimumStakeE8s               *uint64                              `ic:"neuron_minimum_stake_e8s,omitempty" json:"neuron_minimum_stake_e8s,omitempty"`
	ConfirmationText                    *string                              `ic:"confirmation_text,omitempty" json:"confirmation_text,omitempty"`
	SwapStartTimestampSeconds           *uint64                              `ic:"swap_start_timestamp_seconds,omitempty" json:"swap_start_timestamp_seconds,omitempty"`
	SwapDueTimestampSeconds             *uint64                              `ic:"swap_due_timestamp_seconds,omitempty" json:"swap_due_timestamp_seconds,omitempty"`
	MinParticipants                     *uint32                              `ic:"min_participants,omitempty" json:"min_participants,omitempty"`
	SnsTokenE8s                         *uint64                              `ic:"sns_token_e8s,omitempty" json:"sns_token_e8s,omitempty"`
	NnsGovernanceCanisterId             string                               `ic:"nns_governance_canister_id" json:"nns_governance_canister_id"`
	TransactionFeeE8s                   *uint64                              `ic:"transaction_fee_e8s,omitempty" json:"transaction_fee_e8s,omitempty"`
	IcpLedgerCanisterId                 string                               `ic:"icp_ledger_canister_id" json:"icp_ledger_canister_id"`
	SnsLedgerCanisterId                 string                               `ic:"sns_ledger_canister_id" json:"sns_ledger_canister_id"`
	NeuronsFundParticipationConstraints *NeuronsFundParticipationConstraints `ic:"neurons_fund_participation_constraints,omitempty" json:"neurons_fund_participation_constraints,omitempty"`
	NeuronsFundParticipants             *NeuronsFundParticipants             `ic:"neurons_fund_participants,omitempty" json:"neurons_fund_participants,omitempty"`
	ShouldAutoFinalize                  *bool                                `ic:"should_auto_finalize,omitempty" json:"should_auto_finalize,omitempty"`
	MaxParticipantIcpE8s                *uint64                              `ic:"max_participant_icp_e8s,omitempty" json:"max_participant_icp_e8s,omitempty"`
	SnsGovernanceCanisterId             string                               `ic:"sns_governance_canister_id" json:"sns_governance_canister_id"`
	MinDirectParticipationIcpE8s        *uint64                              `ic:"min_direct_participation_icp_e8s,omitempty" json:"min_direct_participation_icp_e8s,omitempty"`
	RestrictedCountries                 *Countries                           `ic:"restricted_countries,omitempty" json:"restricted_countries,omitempty"`
	MinIcpE8s                           *uint64                              `ic:"min_icp_e8s,omitempty" json:"min_icp_e8s,omitempty"`
	MaxDirectParticipationIcpE8s        *uint64                              `ic:"max_direct_participation_icp_e8s,omitempty" json:"max_direct_participation_icp_e8s,omitempty"`
}

type InvalidUserAmount struct {
	MinAmountIcpE8sIncluded uint64 `ic:"min_amount_icp_e8s_included" json:"min_amount_icp_e8s_included"`
	MaxAmountIcpE8sIncluded uint64 `ic:"max_amount_icp_e8s_included" json:"max_amount_icp_e8s_included"`
}

type Investor struct {
	CommunityFund *CfInvestment     `ic:"CommunityFund,variant"`
	Direct        *DirectInvestment `ic:"Direct,variant"`
}

type LinearScalingCoefficient struct {
	SlopeNumerator                *uint64 `ic:"slope_numerator,omitempty" json:"slope_numerator,omitempty"`
	InterceptIcpE8s               *uint64 `ic:"intercept_icp_e8s,omitempty" json:"intercept_icp_e8s,omitempty"`
	FromDirectParticipationIcpE8s *uint64 `ic:"from_direct_participation_icp_e8s,omitempty" json:"from_direct_participation_icp_e8s,omitempty"`
	SlopeDenominator              *uint64 `ic:"slope_denominator,omitempty" json:"slope_denominator,omitempty"`
	ToDirectParticipationIcpE8s   *uint64 `ic:"to_direct_participation_icp_e8s,omitempty" json:"to_direct_participation_icp_e8s,omitempty"`
}

type ListCommunityFundParticipantsRequest struct {
	Offset *uint64 `ic:"offset,omitempty" json:"offset,omitempty"`
	Limit  *uint32 `ic:"limit,omitempty" json:"limit,omitempty"`
}

type ListDirectParticipantsRequest struct {
	Offset *uint32 `ic:"offset,omitempty" json:"offset,omitempty"`
	Limit  *uint32 `ic:"limit,omitempty" json:"limit,omitempty"`
}

type ListDirectParticipantsResponse struct {
	Participants []Participant `ic:"participants" json:"participants"`
}

type ListSnsNeuronRecipesRequest struct {
	Offset *uint64 `ic:"offset,omitempty" json:"offset,omitempty"`
	Limit  *uint32 `ic:"limit,omitempty" json:"limit,omitempty"`
}

type ListSnsNeuronRecipesResponse struct {
	SnsNeuronRecipes []SnsNeuronRecipe `ic:"sns_neuron_recipes" json:"sns_neuron_recipes"`
}

type NeuronAttributes struct {
	DissolveDelaySeconds uint64     `ic:"dissolve_delay_seconds" json:"dissolve_delay_seconds"`
	Memo                 uint64     `ic:"memo" json:"memo"`
	Followees            []NeuronId `ic:"followees" json:"followees"`
}

type NeuronBasketConstructionParameters struct {
	DissolveDelayIntervalSeconds uint64 `ic:"dissolve_delay_interval_seconds" json:"dissolve_delay_interval_seconds"`
	Count                        uint64 `ic:"count" json:"count"`
}

type NeuronId struct {
	Id []byte `ic:"id" json:"id"`
}

type NeuronsFundParticipants struct {
	CfParticipants []CfParticipant `ic:"cf_participants" json:"cf_participants"`
}

type NeuronsFundParticipationConstraints struct {
	CoefficientIntervals                  []LinearScalingCoefficient         `ic:"coefficient_intervals" json:"coefficient_intervals"`
	MaxNeuronsFundParticipationIcpE8s     *uint64                            `ic:"max_neurons_fund_participation_icp_e8s,omitempty" json:"max_neurons_fund_participation_icp_e8s,omitempty"`
	MinDirectParticipationThresholdIcpE8s *uint64                            `ic:"min_direct_participation_threshold_icp_e8s,omitempty" json:"min_direct_participation_threshold_icp_e8s,omitempty"`
	IdealMatchedParticipationFunction     *IdealMatchedParticipationFunction `ic:"ideal_matched_participation_function,omitempty" json:"ideal_matched_participation_function,omitempty"`
}

type NewSaleTicketRequest struct {
	Subaccount   *[]byte `ic:"subaccount,omitempty" json:"subaccount,omitempty"`
	AmountIcpE8s uint64  `ic:"amount_icp_e8s" json:"amount_icp_e8s"`
}

type NewSaleTicketResponse struct {
	Result *Result2 `ic:"result,omitempty" json:"result,omitempty"`
}

type Ok struct {
	BlockHeight *uint64 `ic:"block_height,omitempty" json:"block_height,omitempty"`
}

type Ok1 struct {
	NeuronsFundParticipationIcpE8s *uint64 `ic:"neurons_fund_participation_icp_e8s,omitempty" json:"neurons_fund_participation_icp_e8s,omitempty"`
	NeuronsFundNeuronsCount        *uint64 `ic:"neurons_fund_neurons_count,omitempty" json:"neurons_fund_neurons_count,omitempty"`
}

type Ok2 struct {
	Ticket *Ticket `ic:"ticket,omitempty" json:"ticket,omitempty"`
}

type OpenRequest struct {
	CfParticipants             []CfParticipant `ic:"cf_participants" json:"cf_participants"`
	Params                     *Params         `ic:"params,omitempty" json:"params,omitempty"`
	OpenSnsTokenSwapProposalId *uint64         `ic:"open_sns_token_swap_proposal_id,omitempty" json:"open_sns_token_swap_proposal_id,omitempty"`
}

type Params struct {
	MinParticipantIcpE8s               uint64                              `ic:"min_participant_icp_e8s" json:"min_participant_icp_e8s"`
	NeuronBasketConstructionParameters *NeuronBasketConstructionParameters `ic:"neuron_basket_construction_parameters,omitempty" json:"neuron_basket_construction_parameters,omitempty"`
	MaxIcpE8s                          uint64                              `ic:"max_icp_e8s" json:"max_icp_e8s"`
	SwapDueTimestampSeconds            uint64                              `ic:"swap_due_timestamp_seconds" json:"swap_due_timestamp_seconds"`
	MinParticipants                    uint32                              `ic:"min_participants" json:"min_participants"`
	SnsTokenE8s                        uint64                              `ic:"sns_token_e8s" json:"sns_token_e8s"`
	SaleDelaySeconds                   *uint64                             `ic:"sale_delay_seconds,omitempty" json:"sale_delay_seconds,omitempty"`
	MaxParticipantIcpE8s               uint64                              `ic:"max_participant_icp_e8s" json:"max_participant_icp_e8s"`
	MinDirectParticipationIcpE8s       *uint64                             `ic:"min_direct_participation_icp_e8s,omitempty" json:"min_direct_participation_icp_e8s,omitempty"`
	MinIcpE8s                          uint64                              `ic:"min_icp_e8s" json:"min_icp_e8s"`
	MaxDirectParticipationIcpE8s       *uint64                             `ic:"max_direct_participation_icp_e8s,omitempty" json:"max_direct_participation_icp_e8s,omitempty"`
}

type Participant struct {
	Participation *BuyerState          `ic:"participation,omitempty" json:"participation,omitempty"`
	ParticipantId *principal.Principal `ic:"participant_id,omitempty" json:"participant_id,omitempty"`
}

type Possibility struct {
	Ok  *SetDappControllersResponse `ic:"Ok,variant"`
	Err *CanisterCallError          `ic:"Err,variant"`
}

type Possibility1 struct {
	Ok  *Response          `ic:"Ok,variant"`
	Err *CanisterCallError `ic:"Err,variant"`
}

type Possibility2 struct {
	Ok  *Ok1   `ic:"Ok,variant"`
	Err *Error `ic:"Err,variant"`
}

type Possibility3 struct {
	Ok *struct {
	} `ic:"Ok,variant"`
	Err *CanisterCallError `ic:"Err,variant"`
}

type RefreshBuyerTokensRequest struct {
	ConfirmationText *string `ic:"confirmation_text,omitempty" json:"confirmation_text,omitempty"`
	Buyer            string  `ic:"buyer" json:"buyer"`
}

type RefreshBuyerTokensResponse struct {
	IcpAcceptedParticipationE8s uint64 `ic:"icp_accepted_participation_e8s" json:"icp_accepted_participation_e8s"`
	IcpLedgerAccountBalanceE8s  uint64 `ic:"icp_ledger_account_balance_e8s" json:"icp_ledger_account_balance_e8s"`
}

type Response struct {
	GovernanceError *GovernanceError `ic:"governance_error,omitempty" json:"governance_error,omitempty"`
}

type Result struct {
	Ok  *Ok  `ic:"Ok,variant"`
	Err *Err `ic:"Err,variant"`
}

type Result1 struct {
	Ok  *Ok2  `ic:"Ok,variant"`
	Err *Err1 `ic:"Err,variant"`
}

type Result2 struct {
	Ok  *Ok2  `ic:"Ok,variant"`
	Err *Err2 `ic:"Err,variant"`
}

type SetDappControllersCallResult struct {
	Possibility *Possibility `ic:"possibility,omitempty" json:"possibility,omitempty"`
}

type SetDappControllersResponse struct {
	FailedUpdates []FailedUpdate `ic:"failed_updates" json:"failed_updates"`
}

type SetModeCallResult struct {
	Possibility *Possibility3 `ic:"possibility,omitempty" json:"possibility,omitempty"`
}

type SettleCommunityFundParticipationResult struct {
	Possibility *Possibility1 `ic:"possibility,omitempty" json:"possibility,omitempty"`
}

type SettleNeuronsFundParticipationResult struct {
	Possibility *Possibility2 `ic:"possibility,omitempty" json:"possibility,omitempty"`
}

type SnsNeuronRecipe struct {
	Sns              *TransferableAmount `ic:"sns,omitempty" json:"sns,omitempty"`
	ClaimedStatus    *int32              `ic:"claimed_status,omitempty" json:"claimed_status,omitempty"`
	NeuronAttributes *NeuronAttributes   `ic:"neuron_attributes,omitempty" json:"neuron_attributes,omitempty"`
	Investor         *Investor           `ic:"investor,omitempty" json:"investor,omitempty"`
}

type Swap struct {
	AutoFinalizeSwapResponse                          *FinalizeSwapResponse `ic:"auto_finalize_swap_response,omitempty" json:"auto_finalize_swap_response,omitempty"`
	NeuronRecipes                                     []SnsNeuronRecipe     `ic:"neuron_recipes" json:"neuron_recipes"`
	NextTicketId                                      *uint64               `ic:"next_ticket_id,omitempty" json:"next_ticket_id,omitempty"`
	DecentralizationSaleOpenTimestampSeconds          *uint64               `ic:"decentralization_sale_open_timestamp_seconds,omitempty" json:"decentralization_sale_open_timestamp_seconds,omitempty"`
	FinalizeSwapInProgress                            *bool                 `ic:"finalize_swap_in_progress,omitempty" json:"finalize_swap_in_progress,omitempty"`
	CfParticipants                                    []CfParticipant       `ic:"cf_participants" json:"cf_participants"`
	Init                                              *Init                 `ic:"init,omitempty" json:"init,omitempty"`
	AlreadyTriedToAutoFinalize                        *bool                 `ic:"already_tried_to_auto_finalize,omitempty" json:"already_tried_to_auto_finalize,omitempty"`
	NeuronsFundParticipationIcpE8s                    *uint64               `ic:"neurons_fund_participation_icp_e8s,omitempty" json:"neurons_fund_participation_icp_e8s,omitempty"`
	PurgeOldTicketsLastCompletionTimestampNanoseconds *uint64               `ic:"purge_old_tickets_last_completion_timestamp_nanoseconds,omitempty" json:"purge_old_tickets_last_completion_timestamp_nanoseconds,omitempty"`
	DirectParticipationIcpE8s                         *uint64               `ic:"direct_participation_icp_e8s,omitempty" json:"direct_participation_icp_e8s,omitempty"`
	Lifecycle                                         int32                 `ic:"lifecycle" json:"lifecycle"`
	PurgeOldTicketsNextPrincipal                      *[]byte               `ic:"purge_old_tickets_next_principal,omitempty" json:"purge_old_tickets_next_principal,omitempty"`
	DecentralizationSwapTerminationTimestampSeconds   *uint64               `ic:"decentralization_swap_termination_timestamp_seconds,omitempty" json:"decentralization_swap_termination_timestamp_seconds,omitempty"`
	Buyers                                            []struct {
		Field0 string     `ic:"0" json:"0"`
		Field1 BuyerState `ic:"1" json:"1"`
	} `ic:"buyers" json:"buyers"`
	Params                     *Params `ic:"params,omitempty" json:"params,omitempty"`
	OpenSnsTokenSwapProposalId *uint64 `ic:"open_sns_token_swap_proposal_id,omitempty" json:"open_sns_token_swap_proposal_id,omitempty"`
}

type SweepResult struct {
	Failure        uint32 `ic:"failure" json:"failure"`
	Skipped        uint32 `ic:"skipped" json:"skipped"`
	Invalid        uint32 `ic:"invalid" json:"invalid"`
	Success        uint32 `ic:"success" json:"success"`
	GlobalFailures uint32 `ic:"global_failures" json:"global_failures"`
}

type Ticket struct {
	CreationTime uint64        `ic:"creation_time" json:"creation_time"`
	TicketId     uint64        `ic:"ticket_id" json:"ticket_id"`
	Account      *Icrc1Account `ic:"account,omitempty" json:"account,omitempty"`
	AmountIcpE8s uint64        `ic:"amount_icp_e8s" json:"amount_icp_e8s"`
}

type TransferableAmount struct {
	TransferFeePaidE8s              *uint64 `ic:"transfer_fee_paid_e8s,omitempty" json:"transfer_fee_paid_e8s,omitempty"`
	TransferStartTimestampSeconds   uint64  `ic:"transfer_start_timestamp_seconds" json:"transfer_start_timestamp_seconds"`
	AmountE8s                       uint64  `ic:"amount_e8s" json:"amount_e8s"`
	AmountTransferredE8s            *uint64 `ic:"amount_transferred_e8s,omitempty" json:"amount_transferred_e8s,omitempty"`
	TransferSuccessTimestampSeconds uint64  `ic:"transfer_success_timestamp_seconds" json:"transfer_success_timestamp_seconds"`
}
