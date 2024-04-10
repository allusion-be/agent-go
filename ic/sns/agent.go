// Package sns provides a client for the "sns" canister.
// Do NOT edit this file. It was automatically generated by https://github.com/aviate-labs/agent-go.
package sns

import (
	"github.com/aviate-labs/agent-go"
	"github.com/aviate-labs/agent-go/candid/idl"
	"github.com/aviate-labs/agent-go/principal"
)

type AddWasmRequest struct {
	Hash []byte   `ic:"hash" json:"hash"`
	Wasm *SnsWasm `ic:"wasm,omitempty" json:"wasm,omitempty"`
}

type AddWasmResponse struct {
	Result *Result `ic:"result,omitempty" json:"result,omitempty"`
}

// Agent is a client for the "sns" canister.
type Agent struct {
	a          *agent.Agent
	canisterId principal.Principal
}

// NewAgent creates a new agent for the "sns" canister.
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

// AddWasm calls the "add_wasm" method on the "sns" canister.
func (a Agent) AddWasm(arg0 AddWasmRequest) (*AddWasmResponse, error) {
	var r0 AddWasmResponse
	if err := a.a.Call(
		a.canisterId,
		"add_wasm",
		[]any{arg0},
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

// DeployNewSns calls the "deploy_new_sns" method on the "sns" canister.
func (a Agent) DeployNewSns(arg0 DeployNewSnsRequest) (*DeployNewSnsResponse, error) {
	var r0 DeployNewSnsResponse
	if err := a.a.Call(
		a.canisterId,
		"deploy_new_sns",
		[]any{arg0},
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

// GetAllowedPrincipals calls the "get_allowed_principals" method on the "sns" canister.
func (a Agent) GetAllowedPrincipals(arg0 struct {
}) (*GetAllowedPrincipalsResponse, error) {
	var r0 GetAllowedPrincipalsResponse
	if err := a.a.Query(
		a.canisterId,
		"get_allowed_principals",
		[]any{arg0},
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

// GetDeployedSnsByProposalId calls the "get_deployed_sns_by_proposal_id" method on the "sns" canister.
func (a Agent) GetDeployedSnsByProposalId(arg0 GetDeployedSnsByProposalIdRequest) (*GetDeployedSnsByProposalIdResponse, error) {
	var r0 GetDeployedSnsByProposalIdResponse
	if err := a.a.Query(
		a.canisterId,
		"get_deployed_sns_by_proposal_id",
		[]any{arg0},
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

// GetLatestSnsVersionPretty calls the "get_latest_sns_version_pretty" method on the "sns" canister.
func (a Agent) GetLatestSnsVersionPretty(arg0 idl.Null) (*[]struct {
	Field0 string `ic:"0" json:"0"`
	Field1 string `ic:"1" json:"1"`
}, error) {
	var r0 []struct {
		Field0 string `ic:"0" json:"0"`
		Field1 string `ic:"1" json:"1"`
	}
	if err := a.a.Query(
		a.canisterId,
		"get_latest_sns_version_pretty",
		[]any{arg0},
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

// GetNextSnsVersion calls the "get_next_sns_version" method on the "sns" canister.
func (a Agent) GetNextSnsVersion(arg0 GetNextSnsVersionRequest) (*GetNextSnsVersionResponse, error) {
	var r0 GetNextSnsVersionResponse
	if err := a.a.Query(
		a.canisterId,
		"get_next_sns_version",
		[]any{arg0},
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

// GetSnsSubnetIds calls the "get_sns_subnet_ids" method on the "sns" canister.
func (a Agent) GetSnsSubnetIds(arg0 struct {
}) (*GetSnsSubnetIdsResponse, error) {
	var r0 GetSnsSubnetIdsResponse
	if err := a.a.Query(
		a.canisterId,
		"get_sns_subnet_ids",
		[]any{arg0},
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

// GetWasm calls the "get_wasm" method on the "sns" canister.
func (a Agent) GetWasm(arg0 GetWasmRequest) (*GetWasmResponse, error) {
	var r0 GetWasmResponse
	if err := a.a.Query(
		a.canisterId,
		"get_wasm",
		[]any{arg0},
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

// GetWasmMetadata calls the "get_wasm_metadata" method on the "sns" canister.
func (a Agent) GetWasmMetadata(arg0 GetWasmMetadataRequest) (*GetWasmMetadataResponse, error) {
	var r0 GetWasmMetadataResponse
	if err := a.a.Query(
		a.canisterId,
		"get_wasm_metadata",
		[]any{arg0},
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

// InsertUpgradePathEntries calls the "insert_upgrade_path_entries" method on the "sns" canister.
func (a Agent) InsertUpgradePathEntries(arg0 InsertUpgradePathEntriesRequest) (*InsertUpgradePathEntriesResponse, error) {
	var r0 InsertUpgradePathEntriesResponse
	if err := a.a.Call(
		a.canisterId,
		"insert_upgrade_path_entries",
		[]any{arg0},
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

// ListDeployedSnses calls the "list_deployed_snses" method on the "sns" canister.
func (a Agent) ListDeployedSnses(arg0 struct {
}) (*ListDeployedSnsesResponse, error) {
	var r0 ListDeployedSnsesResponse
	if err := a.a.Query(
		a.canisterId,
		"list_deployed_snses",
		[]any{arg0},
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

// ListUpgradeSteps calls the "list_upgrade_steps" method on the "sns" canister.
func (a Agent) ListUpgradeSteps(arg0 ListUpgradeStepsRequest) (*ListUpgradeStepsResponse, error) {
	var r0 ListUpgradeStepsResponse
	if err := a.a.Query(
		a.canisterId,
		"list_upgrade_steps",
		[]any{arg0},
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

// UpdateAllowedPrincipals calls the "update_allowed_principals" method on the "sns" canister.
func (a Agent) UpdateAllowedPrincipals(arg0 UpdateAllowedPrincipalsRequest) (*UpdateAllowedPrincipalsResponse, error) {
	var r0 UpdateAllowedPrincipalsResponse
	if err := a.a.Call(
		a.canisterId,
		"update_allowed_principals",
		[]any{arg0},
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

// UpdateSnsSubnetList calls the "update_sns_subnet_list" method on the "sns" canister.
func (a Agent) UpdateSnsSubnetList(arg0 UpdateSnsSubnetListRequest) (*UpdateSnsSubnetListResponse, error) {
	var r0 UpdateSnsSubnetListResponse
	if err := a.a.Call(
		a.canisterId,
		"update_sns_subnet_list",
		[]any{arg0},
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

type AirdropDistribution struct {
	AirdropNeurons []NeuronDistribution `ic:"airdrop_neurons" json:"airdrop_neurons"`
}

type Canister struct {
	Id *principal.Principal `ic:"id,omitempty" json:"id,omitempty"`
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

type DappCanisters struct {
	Canisters []Canister `ic:"canisters" json:"canisters"`
}

type DappCanistersTransferResult struct {
	RestoredDappCanisters      []Canister `ic:"restored_dapp_canisters" json:"restored_dapp_canisters"`
	NnsControlledDappCanisters []Canister `ic:"nns_controlled_dapp_canisters" json:"nns_controlled_dapp_canisters"`
	SnsControlledDappCanisters []Canister `ic:"sns_controlled_dapp_canisters" json:"sns_controlled_dapp_canisters"`
}

type DeployNewSnsRequest struct {
	SnsInitPayload *SnsInitPayload `ic:"sns_init_payload,omitempty" json:"sns_init_payload,omitempty"`
}

type DeployNewSnsResponse struct {
	DappCanistersTransferResult *DappCanistersTransferResult `ic:"dapp_canisters_transfer_result,omitempty" json:"dapp_canisters_transfer_result,omitempty"`
	SubnetId                    *principal.Principal         `ic:"subnet_id,omitempty" json:"subnet_id,omitempty"`
	Error                       *SnsWasmError                `ic:"error,omitempty" json:"error,omitempty"`
	Canisters                   *SnsCanisterIds              `ic:"canisters,omitempty" json:"canisters,omitempty"`
}

type DeployedSns struct {
	RootCanisterId       *principal.Principal `ic:"root_canister_id,omitempty" json:"root_canister_id,omitempty"`
	GovernanceCanisterId *principal.Principal `ic:"governance_canister_id,omitempty" json:"governance_canister_id,omitempty"`
	IndexCanisterId      *principal.Principal `ic:"index_canister_id,omitempty" json:"index_canister_id,omitempty"`
	SwapCanisterId       *principal.Principal `ic:"swap_canister_id,omitempty" json:"swap_canister_id,omitempty"`
	LedgerCanisterId     *principal.Principal `ic:"ledger_canister_id,omitempty" json:"ledger_canister_id,omitempty"`
}

type DeveloperDistribution struct {
	DeveloperNeurons []NeuronDistribution `ic:"developer_neurons" json:"developer_neurons"`
}

type FractionalDeveloperVotingPower struct {
	TreasuryDistribution  *TreasuryDistribution  `ic:"treasury_distribution,omitempty" json:"treasury_distribution,omitempty"`
	DeveloperDistribution *DeveloperDistribution `ic:"developer_distribution,omitempty" json:"developer_distribution,omitempty"`
	AirdropDistribution   *AirdropDistribution   `ic:"airdrop_distribution,omitempty" json:"airdrop_distribution,omitempty"`
	SwapDistribution      *SwapDistribution      `ic:"swap_distribution,omitempty" json:"swap_distribution,omitempty"`
}

type GetAllowedPrincipalsResponse struct {
	AllowedPrincipals []principal.Principal `ic:"allowed_principals" json:"allowed_principals"`
}

type GetDeployedSnsByProposalIdRequest struct {
	ProposalId uint64 `ic:"proposal_id" json:"proposal_id"`
}

type GetDeployedSnsByProposalIdResponse struct {
	GetDeployedSnsByProposalIdResult *GetDeployedSnsByProposalIdResult `ic:"get_deployed_sns_by_proposal_id_result,omitempty" json:"get_deployed_sns_by_proposal_id_result,omitempty"`
}

type GetDeployedSnsByProposalIdResult struct {
	Error       *SnsWasmError `ic:"Error,variant"`
	DeployedSns *DeployedSns  `ic:"DeployedSns,variant"`
}

type GetNextSnsVersionRequest struct {
	GovernanceCanisterId *principal.Principal `ic:"governance_canister_id,omitempty" json:"governance_canister_id,omitempty"`
	CurrentVersion       *SnsVersion          `ic:"current_version,omitempty" json:"current_version,omitempty"`
}

type GetNextSnsVersionResponse struct {
	NextVersion *SnsVersion `ic:"next_version,omitempty" json:"next_version,omitempty"`
}

type GetSnsSubnetIdsResponse struct {
	SnsSubnetIds []principal.Principal `ic:"sns_subnet_ids" json:"sns_subnet_ids"`
}

type GetWasmMetadataRequest struct {
	Hash *[]byte `ic:"hash,omitempty" json:"hash,omitempty"`
}

type GetWasmMetadataResponse struct {
	Result *Result1 `ic:"result,omitempty" json:"result,omitempty"`
}

type GetWasmRequest struct {
	Hash []byte `ic:"hash" json:"hash"`
}

type GetWasmResponse struct {
	Wasm *SnsWasm `ic:"wasm,omitempty" json:"wasm,omitempty"`
}

type IdealMatchedParticipationFunction struct {
	SerializedRepresentation *string `ic:"serialized_representation,omitempty" json:"serialized_representation,omitempty"`
}

type InitialTokenDistribution struct {
	FractionalDeveloperVotingPower *FractionalDeveloperVotingPower `ic:"FractionalDeveloperVotingPower,variant"`
}

type InsertUpgradePathEntriesRequest struct {
	UpgradePath             []SnsUpgrade         `ic:"upgrade_path" json:"upgrade_path"`
	SnsGovernanceCanisterId *principal.Principal `ic:"sns_governance_canister_id,omitempty" json:"sns_governance_canister_id,omitempty"`
}

type InsertUpgradePathEntriesResponse struct {
	Error *SnsWasmError `ic:"error,omitempty" json:"error,omitempty"`
}

type LinearScalingCoefficient struct {
	SlopeNumerator                *uint64 `ic:"slope_numerator,omitempty" json:"slope_numerator,omitempty"`
	InterceptIcpE8s               *uint64 `ic:"intercept_icp_e8s,omitempty" json:"intercept_icp_e8s,omitempty"`
	FromDirectParticipationIcpE8s *uint64 `ic:"from_direct_participation_icp_e8s,omitempty" json:"from_direct_participation_icp_e8s,omitempty"`
	SlopeDenominator              *uint64 `ic:"slope_denominator,omitempty" json:"slope_denominator,omitempty"`
	ToDirectParticipationIcpE8s   *uint64 `ic:"to_direct_participation_icp_e8s,omitempty" json:"to_direct_participation_icp_e8s,omitempty"`
}

type ListDeployedSnsesResponse struct {
	Instances []DeployedSns `ic:"instances" json:"instances"`
}

type ListUpgradeStep struct {
	PrettyVersion *PrettySnsVersion `ic:"pretty_version,omitempty" json:"pretty_version,omitempty"`
	Version       *SnsVersion       `ic:"version,omitempty" json:"version,omitempty"`
}

type ListUpgradeStepsRequest struct {
	Limit                   uint32               `ic:"limit" json:"limit"`
	StartingAt              *SnsVersion          `ic:"starting_at,omitempty" json:"starting_at,omitempty"`
	SnsGovernanceCanisterId *principal.Principal `ic:"sns_governance_canister_id,omitempty" json:"sns_governance_canister_id,omitempty"`
}

type ListUpgradeStepsResponse struct {
	Steps []ListUpgradeStep `ic:"steps" json:"steps"`
}

type MetadataSection struct {
	Contents   *[]byte `ic:"contents,omitempty" json:"contents,omitempty"`
	Name       *string `ic:"name,omitempty" json:"name,omitempty"`
	Visibility *string `ic:"visibility,omitempty" json:"visibility,omitempty"`
}

type NeuronBasketConstructionParameters struct {
	DissolveDelayIntervalSeconds uint64 `ic:"dissolve_delay_interval_seconds" json:"dissolve_delay_interval_seconds"`
	Count                        uint64 `ic:"count" json:"count"`
}

type NeuronDistribution struct {
	Controller           *principal.Principal `ic:"controller,omitempty" json:"controller,omitempty"`
	DissolveDelaySeconds uint64               `ic:"dissolve_delay_seconds" json:"dissolve_delay_seconds"`
	Memo                 uint64               `ic:"memo" json:"memo"`
	StakeE8s             uint64               `ic:"stake_e8s" json:"stake_e8s"`
	VestingPeriodSeconds *uint64              `ic:"vesting_period_seconds,omitempty" json:"vesting_period_seconds,omitempty"`
}

type NeuronsFundParticipants struct {
	Participants []CfParticipant `ic:"participants" json:"participants"`
}

type NeuronsFundParticipationConstraints struct {
	CoefficientIntervals                  []LinearScalingCoefficient         `ic:"coefficient_intervals" json:"coefficient_intervals"`
	MaxNeuronsFundParticipationIcpE8s     *uint64                            `ic:"max_neurons_fund_participation_icp_e8s,omitempty" json:"max_neurons_fund_participation_icp_e8s,omitempty"`
	MinDirectParticipationThresholdIcpE8s *uint64                            `ic:"min_direct_participation_threshold_icp_e8s,omitempty" json:"min_direct_participation_threshold_icp_e8s,omitempty"`
	IdealMatchedParticipationFunction     *IdealMatchedParticipationFunction `ic:"ideal_matched_participation_function,omitempty" json:"ideal_matched_participation_function,omitempty"`
}

type Ok struct {
	Sections []MetadataSection `ic:"sections" json:"sections"`
}

type PrettySnsVersion struct {
	ArchiveWasmHash    string `ic:"archive_wasm_hash" json:"archive_wasm_hash"`
	RootWasmHash       string `ic:"root_wasm_hash" json:"root_wasm_hash"`
	SwapWasmHash       string `ic:"swap_wasm_hash" json:"swap_wasm_hash"`
	LedgerWasmHash     string `ic:"ledger_wasm_hash" json:"ledger_wasm_hash"`
	GovernanceWasmHash string `ic:"governance_wasm_hash" json:"governance_wasm_hash"`
	IndexWasmHash      string `ic:"index_wasm_hash" json:"index_wasm_hash"`
}

type Result struct {
	Error *SnsWasmError `ic:"Error,variant"`
	Hash  *[]byte       `ic:"Hash,variant"`
}

type Result1 struct {
	Ok    *Ok           `ic:"Ok,variant"`
	Error *SnsWasmError `ic:"Error,variant"`
}

type SnsCanisterIds struct {
	Root       *principal.Principal `ic:"root,omitempty" json:"root,omitempty"`
	Swap       *principal.Principal `ic:"swap,omitempty" json:"swap,omitempty"`
	Ledger     *principal.Principal `ic:"ledger,omitempty" json:"ledger,omitempty"`
	Index      *principal.Principal `ic:"index,omitempty" json:"index,omitempty"`
	Governance *principal.Principal `ic:"governance,omitempty" json:"governance,omitempty"`
}

type SnsInitPayload struct {
	Url                                     *string                              `ic:"url,omitempty" json:"url,omitempty"`
	MaxDissolveDelaySeconds                 *uint64                              `ic:"max_dissolve_delay_seconds,omitempty" json:"max_dissolve_delay_seconds,omitempty"`
	MaxDissolveDelayBonusPercentage         *uint64                              `ic:"max_dissolve_delay_bonus_percentage,omitempty" json:"max_dissolve_delay_bonus_percentage,omitempty"`
	NnsProposalId                           *uint64                              `ic:"nns_proposal_id,omitempty" json:"nns_proposal_id,omitempty"`
	NeuronsFundParticipation                *bool                                `ic:"neurons_fund_participation,omitempty" json:"neurons_fund_participation,omitempty"`
	MinParticipantIcpE8s                    *uint64                              `ic:"min_participant_icp_e8s,omitempty" json:"min_participant_icp_e8s,omitempty"`
	NeuronBasketConstructionParameters      *NeuronBasketConstructionParameters  `ic:"neuron_basket_construction_parameters,omitempty" json:"neuron_basket_construction_parameters,omitempty"`
	FallbackControllerPrincipalIds          []string                             `ic:"fallback_controller_principal_ids" json:"fallback_controller_principal_ids"`
	TokenSymbol                             *string                              `ic:"token_symbol,omitempty" json:"token_symbol,omitempty"`
	FinalRewardRateBasisPoints              *uint64                              `ic:"final_reward_rate_basis_points,omitempty" json:"final_reward_rate_basis_points,omitempty"`
	MaxIcpE8s                               *uint64                              `ic:"max_icp_e8s,omitempty" json:"max_icp_e8s,omitempty"`
	NeuronMinimumStakeE8s                   *uint64                              `ic:"neuron_minimum_stake_e8s,omitempty" json:"neuron_minimum_stake_e8s,omitempty"`
	ConfirmationText                        *string                              `ic:"confirmation_text,omitempty" json:"confirmation_text,omitempty"`
	Logo                                    *string                              `ic:"logo,omitempty" json:"logo,omitempty"`
	Name                                    *string                              `ic:"name,omitempty" json:"name,omitempty"`
	SwapStartTimestampSeconds               *uint64                              `ic:"swap_start_timestamp_seconds,omitempty" json:"swap_start_timestamp_seconds,omitempty"`
	SwapDueTimestampSeconds                 *uint64                              `ic:"swap_due_timestamp_seconds,omitempty" json:"swap_due_timestamp_seconds,omitempty"`
	InitialVotingPeriodSeconds              *uint64                              `ic:"initial_voting_period_seconds,omitempty" json:"initial_voting_period_seconds,omitempty"`
	NeuronMinimumDissolveDelayToVoteSeconds *uint64                              `ic:"neuron_minimum_dissolve_delay_to_vote_seconds,omitempty" json:"neuron_minimum_dissolve_delay_to_vote_seconds,omitempty"`
	Description                             *string                              `ic:"description,omitempty" json:"description,omitempty"`
	MaxNeuronAgeSecondsForAgeBonus          *uint64                              `ic:"max_neuron_age_seconds_for_age_bonus,omitempty" json:"max_neuron_age_seconds_for_age_bonus,omitempty"`
	MinParticipants                         *uint64                              `ic:"min_participants,omitempty" json:"min_participants,omitempty"`
	InitialRewardRateBasisPoints            *uint64                              `ic:"initial_reward_rate_basis_points,omitempty" json:"initial_reward_rate_basis_points,omitempty"`
	WaitForQuietDeadlineIncreaseSeconds     *uint64                              `ic:"wait_for_quiet_deadline_increase_seconds,omitempty" json:"wait_for_quiet_deadline_increase_seconds,omitempty"`
	TransactionFeeE8s                       *uint64                              `ic:"transaction_fee_e8s,omitempty" json:"transaction_fee_e8s,omitempty"`
	DappCanisters                           *DappCanisters                       `ic:"dapp_canisters,omitempty" json:"dapp_canisters,omitempty"`
	NeuronsFundParticipationConstraints     *NeuronsFundParticipationConstraints `ic:"neurons_fund_participation_constraints,omitempty" json:"neurons_fund_participation_constraints,omitempty"`
	NeuronsFundParticipants                 *NeuronsFundParticipants             `ic:"neurons_fund_participants,omitempty" json:"neurons_fund_participants,omitempty"`
	MaxAgeBonusPercentage                   *uint64                              `ic:"max_age_bonus_percentage,omitempty" json:"max_age_bonus_percentage,omitempty"`
	InitialTokenDistribution                *InitialTokenDistribution            `ic:"initial_token_distribution,omitempty" json:"initial_token_distribution,omitempty"`
	RewardRateTransitionDurationSeconds     *uint64                              `ic:"reward_rate_transition_duration_seconds,omitempty" json:"reward_rate_transition_duration_seconds,omitempty"`
	TokenLogo                               *string                              `ic:"token_logo,omitempty" json:"token_logo,omitempty"`
	TokenName                               *string                              `ic:"token_name,omitempty" json:"token_name,omitempty"`
	MaxParticipantIcpE8s                    *uint64                              `ic:"max_participant_icp_e8s,omitempty" json:"max_participant_icp_e8s,omitempty"`
	MinDirectParticipationIcpE8s            *uint64                              `ic:"min_direct_participation_icp_e8s,omitempty" json:"min_direct_participation_icp_e8s,omitempty"`
	ProposalRejectCostE8s                   *uint64                              `ic:"proposal_reject_cost_e8s,omitempty" json:"proposal_reject_cost_e8s,omitempty"`
	RestrictedCountries                     *Countries                           `ic:"restricted_countries,omitempty" json:"restricted_countries,omitempty"`
	MinIcpE8s                               *uint64                              `ic:"min_icp_e8s,omitempty" json:"min_icp_e8s,omitempty"`
	MaxDirectParticipationIcpE8s            *uint64                              `ic:"max_direct_participation_icp_e8s,omitempty" json:"max_direct_participation_icp_e8s,omitempty"`
}

type SnsUpgrade struct {
	NextVersion    *SnsVersion `ic:"next_version,omitempty" json:"next_version,omitempty"`
	CurrentVersion *SnsVersion `ic:"current_version,omitempty" json:"current_version,omitempty"`
}

type SnsVersion struct {
	ArchiveWasmHash    []byte `ic:"archive_wasm_hash" json:"archive_wasm_hash"`
	RootWasmHash       []byte `ic:"root_wasm_hash" json:"root_wasm_hash"`
	SwapWasmHash       []byte `ic:"swap_wasm_hash" json:"swap_wasm_hash"`
	LedgerWasmHash     []byte `ic:"ledger_wasm_hash" json:"ledger_wasm_hash"`
	GovernanceWasmHash []byte `ic:"governance_wasm_hash" json:"governance_wasm_hash"`
	IndexWasmHash      []byte `ic:"index_wasm_hash" json:"index_wasm_hash"`
}

type SnsWasm struct {
	Wasm         []byte `ic:"wasm" json:"wasm"`
	CanisterType int32  `ic:"canister_type" json:"canister_type"`
}

type SnsWasmCanisterInitPayload struct {
	AllowedPrincipals     []principal.Principal `ic:"allowed_principals" json:"allowed_principals"`
	AccessControlsEnabled bool                  `ic:"access_controls_enabled" json:"access_controls_enabled"`
	SnsSubnetIds          []principal.Principal `ic:"sns_subnet_ids" json:"sns_subnet_ids"`
}

type SnsWasmError struct {
	Message string `ic:"message" json:"message"`
}

type SwapDistribution struct {
	TotalE8s             uint64 `ic:"total_e8s" json:"total_e8s"`
	InitialSwapAmountE8s uint64 `ic:"initial_swap_amount_e8s" json:"initial_swap_amount_e8s"`
}

type TreasuryDistribution struct {
	TotalE8s uint64 `ic:"total_e8s" json:"total_e8s"`
}

type UpdateAllowedPrincipalsRequest struct {
	AddedPrincipals   []principal.Principal `ic:"added_principals" json:"added_principals"`
	RemovedPrincipals []principal.Principal `ic:"removed_principals" json:"removed_principals"`
}

type UpdateAllowedPrincipalsResponse struct {
	UpdateAllowedPrincipalsResult *UpdateAllowedPrincipalsResult `ic:"update_allowed_principals_result,omitempty" json:"update_allowed_principals_result,omitempty"`
}

type UpdateAllowedPrincipalsResult struct {
	Error             *SnsWasmError                 `ic:"Error,variant"`
	AllowedPrincipals *GetAllowedPrincipalsResponse `ic:"AllowedPrincipals,variant"`
}

type UpdateSnsSubnetListRequest struct {
	SnsSubnetIdsToAdd    []principal.Principal `ic:"sns_subnet_ids_to_add" json:"sns_subnet_ids_to_add"`
	SnsSubnetIdsToRemove []principal.Principal `ic:"sns_subnet_ids_to_remove" json:"sns_subnet_ids_to_remove"`
}

type UpdateSnsSubnetListResponse struct {
	Error *SnsWasmError `ic:"error,omitempty" json:"error,omitempty"`
}
