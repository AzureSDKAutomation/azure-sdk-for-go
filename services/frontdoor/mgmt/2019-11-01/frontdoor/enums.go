package frontdoor

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// ActionType enumerates the values for action type.
type ActionType string

const (
	// Allow ...
	Allow ActionType = "Allow"
	// Block ...
	Block ActionType = "Block"
	// Log ...
	Log ActionType = "Log"
	// Redirect ...
	Redirect ActionType = "Redirect"
)

// PossibleActionTypeValues returns an array of possible values for the ActionType const type.
func PossibleActionTypeValues() []ActionType {
	return []ActionType{Allow, Block, Log, Redirect}
}

// AggregationInterval enumerates the values for aggregation interval.
type AggregationInterval string

const (
	// Daily ...
	Daily AggregationInterval = "Daily"
	// Hourly ...
	Hourly AggregationInterval = "Hourly"
)

// PossibleAggregationIntervalValues returns an array of possible values for the AggregationInterval const type.
func PossibleAggregationIntervalValues() []AggregationInterval {
	return []AggregationInterval{Daily, Hourly}
}

// Availability enumerates the values for availability.
type Availability string

const (
	// Available ...
	Available Availability = "Available"
	// Unavailable ...
	Unavailable Availability = "Unavailable"
)

// PossibleAvailabilityValues returns an array of possible values for the Availability const type.
func PossibleAvailabilityValues() []Availability {
	return []Availability{Available, Unavailable}
}

// BackendEnabledState enumerates the values for backend enabled state.
type BackendEnabledState string

const (
	// Disabled ...
	Disabled BackendEnabledState = "Disabled"
	// Enabled ...
	Enabled BackendEnabledState = "Enabled"
)

// PossibleBackendEnabledStateValues returns an array of possible values for the BackendEnabledState const type.
func PossibleBackendEnabledStateValues() []BackendEnabledState {
	return []BackendEnabledState{Disabled, Enabled}
}

// CertificateSource enumerates the values for certificate source.
type CertificateSource string

const (
	// CertificateSourceAzureKeyVault ...
	CertificateSourceAzureKeyVault CertificateSource = "AzureKeyVault"
	// CertificateSourceFrontDoor ...
	CertificateSourceFrontDoor CertificateSource = "FrontDoor"
)

// PossibleCertificateSourceValues returns an array of possible values for the CertificateSource const type.
func PossibleCertificateSourceValues() []CertificateSource {
	return []CertificateSource{CertificateSourceAzureKeyVault, CertificateSourceFrontDoor}
}

// CertificateType enumerates the values for certificate type.
type CertificateType string

const (
	// Dedicated ...
	Dedicated CertificateType = "Dedicated"
)

// PossibleCertificateTypeValues returns an array of possible values for the CertificateType const type.
func PossibleCertificateTypeValues() []CertificateType {
	return []CertificateType{Dedicated}
}

// CustomHTTPSProvisioningState enumerates the values for custom https provisioning state.
type CustomHTTPSProvisioningState string

const (
	// CustomHTTPSProvisioningStateDisabled ...
	CustomHTTPSProvisioningStateDisabled CustomHTTPSProvisioningState = "Disabled"
	// CustomHTTPSProvisioningStateDisabling ...
	CustomHTTPSProvisioningStateDisabling CustomHTTPSProvisioningState = "Disabling"
	// CustomHTTPSProvisioningStateEnabled ...
	CustomHTTPSProvisioningStateEnabled CustomHTTPSProvisioningState = "Enabled"
	// CustomHTTPSProvisioningStateEnabling ...
	CustomHTTPSProvisioningStateEnabling CustomHTTPSProvisioningState = "Enabling"
	// CustomHTTPSProvisioningStateFailed ...
	CustomHTTPSProvisioningStateFailed CustomHTTPSProvisioningState = "Failed"
)

// PossibleCustomHTTPSProvisioningStateValues returns an array of possible values for the CustomHTTPSProvisioningState const type.
func PossibleCustomHTTPSProvisioningStateValues() []CustomHTTPSProvisioningState {
	return []CustomHTTPSProvisioningState{CustomHTTPSProvisioningStateDisabled, CustomHTTPSProvisioningStateDisabling, CustomHTTPSProvisioningStateEnabled, CustomHTTPSProvisioningStateEnabling, CustomHTTPSProvisioningStateFailed}
}

// CustomHTTPSProvisioningSubstate enumerates the values for custom https provisioning substate.
type CustomHTTPSProvisioningSubstate string

const (
	// CertificateDeleted ...
	CertificateDeleted CustomHTTPSProvisioningSubstate = "CertificateDeleted"
	// CertificateDeployed ...
	CertificateDeployed CustomHTTPSProvisioningSubstate = "CertificateDeployed"
	// DeletingCertificate ...
	DeletingCertificate CustomHTTPSProvisioningSubstate = "DeletingCertificate"
	// DeployingCertificate ...
	DeployingCertificate CustomHTTPSProvisioningSubstate = "DeployingCertificate"
	// DomainControlValidationRequestApproved ...
	DomainControlValidationRequestApproved CustomHTTPSProvisioningSubstate = "DomainControlValidationRequestApproved"
	// DomainControlValidationRequestRejected ...
	DomainControlValidationRequestRejected CustomHTTPSProvisioningSubstate = "DomainControlValidationRequestRejected"
	// DomainControlValidationRequestTimedOut ...
	DomainControlValidationRequestTimedOut CustomHTTPSProvisioningSubstate = "DomainControlValidationRequestTimedOut"
	// IssuingCertificate ...
	IssuingCertificate CustomHTTPSProvisioningSubstate = "IssuingCertificate"
	// PendingDomainControlValidationREquestApproval ...
	PendingDomainControlValidationREquestApproval CustomHTTPSProvisioningSubstate = "PendingDomainControlValidationREquestApproval"
	// SubmittingDomainControlValidationRequest ...
	SubmittingDomainControlValidationRequest CustomHTTPSProvisioningSubstate = "SubmittingDomainControlValidationRequest"
)

// PossibleCustomHTTPSProvisioningSubstateValues returns an array of possible values for the CustomHTTPSProvisioningSubstate const type.
func PossibleCustomHTTPSProvisioningSubstateValues() []CustomHTTPSProvisioningSubstate {
	return []CustomHTTPSProvisioningSubstate{CertificateDeleted, CertificateDeployed, DeletingCertificate, DeployingCertificate, DomainControlValidationRequestApproved, DomainControlValidationRequestRejected, DomainControlValidationRequestTimedOut, IssuingCertificate, PendingDomainControlValidationREquestApproval, SubmittingDomainControlValidationRequest}
}

// CustomRuleEnabledState enumerates the values for custom rule enabled state.
type CustomRuleEnabledState string

const (
	// CustomRuleEnabledStateDisabled ...
	CustomRuleEnabledStateDisabled CustomRuleEnabledState = "Disabled"
	// CustomRuleEnabledStateEnabled ...
	CustomRuleEnabledStateEnabled CustomRuleEnabledState = "Enabled"
)

// PossibleCustomRuleEnabledStateValues returns an array of possible values for the CustomRuleEnabledState const type.
func PossibleCustomRuleEnabledStateValues() []CustomRuleEnabledState {
	return []CustomRuleEnabledState{CustomRuleEnabledStateDisabled, CustomRuleEnabledStateEnabled}
}

// DynamicCompressionEnabled enumerates the values for dynamic compression enabled.
type DynamicCompressionEnabled string

const (
	// DynamicCompressionEnabledDisabled ...
	DynamicCompressionEnabledDisabled DynamicCompressionEnabled = "Disabled"
	// DynamicCompressionEnabledEnabled ...
	DynamicCompressionEnabledEnabled DynamicCompressionEnabled = "Enabled"
)

// PossibleDynamicCompressionEnabledValues returns an array of possible values for the DynamicCompressionEnabled const type.
func PossibleDynamicCompressionEnabledValues() []DynamicCompressionEnabled {
	return []DynamicCompressionEnabled{DynamicCompressionEnabledDisabled, DynamicCompressionEnabledEnabled}
}

// EnabledState enumerates the values for enabled state.
type EnabledState string

const (
	// EnabledStateDisabled ...
	EnabledStateDisabled EnabledState = "Disabled"
	// EnabledStateEnabled ...
	EnabledStateEnabled EnabledState = "Enabled"
)

// PossibleEnabledStateValues returns an array of possible values for the EnabledState const type.
func PossibleEnabledStateValues() []EnabledState {
	return []EnabledState{EnabledStateDisabled, EnabledStateEnabled}
}

// EndpointType enumerates the values for endpoint type.
type EndpointType string

const (
	// AFD ...
	AFD EndpointType = "AFD"
	// ATM ...
	ATM EndpointType = "ATM"
	// AzureRegion ...
	AzureRegion EndpointType = "AzureRegion"
	// CDN ...
	CDN EndpointType = "CDN"
)

// PossibleEndpointTypeValues returns an array of possible values for the EndpointType const type.
func PossibleEndpointTypeValues() []EndpointType {
	return []EndpointType{AFD, ATM, AzureRegion, CDN}
}

// EnforceCertificateNameCheckEnabledState enumerates the values for enforce certificate name check enabled
// state.
type EnforceCertificateNameCheckEnabledState string

const (
	// EnforceCertificateNameCheckEnabledStateDisabled ...
	EnforceCertificateNameCheckEnabledStateDisabled EnforceCertificateNameCheckEnabledState = "Disabled"
	// EnforceCertificateNameCheckEnabledStateEnabled ...
	EnforceCertificateNameCheckEnabledStateEnabled EnforceCertificateNameCheckEnabledState = "Enabled"
)

// PossibleEnforceCertificateNameCheckEnabledStateValues returns an array of possible values for the EnforceCertificateNameCheckEnabledState const type.
func PossibleEnforceCertificateNameCheckEnabledStateValues() []EnforceCertificateNameCheckEnabledState {
	return []EnforceCertificateNameCheckEnabledState{EnforceCertificateNameCheckEnabledStateDisabled, EnforceCertificateNameCheckEnabledStateEnabled}
}

// ForwardingProtocol enumerates the values for forwarding protocol.
type ForwardingProtocol string

const (
	// HTTPOnly ...
	HTTPOnly ForwardingProtocol = "HttpOnly"
	// HTTPSOnly ...
	HTTPSOnly ForwardingProtocol = "HttpsOnly"
	// MatchRequest ...
	MatchRequest ForwardingProtocol = "MatchRequest"
)

// PossibleForwardingProtocolValues returns an array of possible values for the ForwardingProtocol const type.
func PossibleForwardingProtocolValues() []ForwardingProtocol {
	return []ForwardingProtocol{HTTPOnly, HTTPSOnly, MatchRequest}
}

// HealthProbeEnabled enumerates the values for health probe enabled.
type HealthProbeEnabled string

const (
	// HealthProbeEnabledDisabled ...
	HealthProbeEnabledDisabled HealthProbeEnabled = "Disabled"
	// HealthProbeEnabledEnabled ...
	HealthProbeEnabledEnabled HealthProbeEnabled = "Enabled"
)

// PossibleHealthProbeEnabledValues returns an array of possible values for the HealthProbeEnabled const type.
func PossibleHealthProbeEnabledValues() []HealthProbeEnabled {
	return []HealthProbeEnabled{HealthProbeEnabledDisabled, HealthProbeEnabledEnabled}
}

// HealthProbeMethod enumerates the values for health probe method.
type HealthProbeMethod string

const (
	// GET ...
	GET HealthProbeMethod = "GET"
	// HEAD ...
	HEAD HealthProbeMethod = "HEAD"
)

// PossibleHealthProbeMethodValues returns an array of possible values for the HealthProbeMethod const type.
func PossibleHealthProbeMethodValues() []HealthProbeMethod {
	return []HealthProbeMethod{GET, HEAD}
}

// LatencyScorecardAggregationInterval enumerates the values for latency scorecard aggregation interval.
type LatencyScorecardAggregationInterval string

const (
	// LatencyScorecardAggregationIntervalDaily ...
	LatencyScorecardAggregationIntervalDaily LatencyScorecardAggregationInterval = "Daily"
	// LatencyScorecardAggregationIntervalMonthly ...
	LatencyScorecardAggregationIntervalMonthly LatencyScorecardAggregationInterval = "Monthly"
	// LatencyScorecardAggregationIntervalWeekly ...
	LatencyScorecardAggregationIntervalWeekly LatencyScorecardAggregationInterval = "Weekly"
)

// PossibleLatencyScorecardAggregationIntervalValues returns an array of possible values for the LatencyScorecardAggregationInterval const type.
func PossibleLatencyScorecardAggregationIntervalValues() []LatencyScorecardAggregationInterval {
	return []LatencyScorecardAggregationInterval{LatencyScorecardAggregationIntervalDaily, LatencyScorecardAggregationIntervalMonthly, LatencyScorecardAggregationIntervalWeekly}
}

// ManagedRuleEnabledState enumerates the values for managed rule enabled state.
type ManagedRuleEnabledState string

const (
	// ManagedRuleEnabledStateDisabled ...
	ManagedRuleEnabledStateDisabled ManagedRuleEnabledState = "Disabled"
	// ManagedRuleEnabledStateEnabled ...
	ManagedRuleEnabledStateEnabled ManagedRuleEnabledState = "Enabled"
)

// PossibleManagedRuleEnabledStateValues returns an array of possible values for the ManagedRuleEnabledState const type.
func PossibleManagedRuleEnabledStateValues() []ManagedRuleEnabledState {
	return []ManagedRuleEnabledState{ManagedRuleEnabledStateDisabled, ManagedRuleEnabledStateEnabled}
}

// ManagedRuleExclusionMatchVariable enumerates the values for managed rule exclusion match variable.
type ManagedRuleExclusionMatchVariable string

const (
	// QueryStringArgNames ...
	QueryStringArgNames ManagedRuleExclusionMatchVariable = "QueryStringArgNames"
	// RequestBodyPostArgNames ...
	RequestBodyPostArgNames ManagedRuleExclusionMatchVariable = "RequestBodyPostArgNames"
	// RequestCookieNames ...
	RequestCookieNames ManagedRuleExclusionMatchVariable = "RequestCookieNames"
	// RequestHeaderNames ...
	RequestHeaderNames ManagedRuleExclusionMatchVariable = "RequestHeaderNames"
)

// PossibleManagedRuleExclusionMatchVariableValues returns an array of possible values for the ManagedRuleExclusionMatchVariable const type.
func PossibleManagedRuleExclusionMatchVariableValues() []ManagedRuleExclusionMatchVariable {
	return []ManagedRuleExclusionMatchVariable{QueryStringArgNames, RequestBodyPostArgNames, RequestCookieNames, RequestHeaderNames}
}

// ManagedRuleExclusionSelectorMatchOperator enumerates the values for managed rule exclusion selector match
// operator.
type ManagedRuleExclusionSelectorMatchOperator string

const (
	// Contains ...
	Contains ManagedRuleExclusionSelectorMatchOperator = "Contains"
	// EndsWith ...
	EndsWith ManagedRuleExclusionSelectorMatchOperator = "EndsWith"
	// Equals ...
	Equals ManagedRuleExclusionSelectorMatchOperator = "Equals"
	// EqualsAny ...
	EqualsAny ManagedRuleExclusionSelectorMatchOperator = "EqualsAny"
	// StartsWith ...
	StartsWith ManagedRuleExclusionSelectorMatchOperator = "StartsWith"
)

// PossibleManagedRuleExclusionSelectorMatchOperatorValues returns an array of possible values for the ManagedRuleExclusionSelectorMatchOperator const type.
func PossibleManagedRuleExclusionSelectorMatchOperatorValues() []ManagedRuleExclusionSelectorMatchOperator {
	return []ManagedRuleExclusionSelectorMatchOperator{Contains, EndsWith, Equals, EqualsAny, StartsWith}
}

// MatchVariable enumerates the values for match variable.
type MatchVariable string

const (
	// Cookies ...
	Cookies MatchVariable = "Cookies"
	// PostArgs ...
	PostArgs MatchVariable = "PostArgs"
	// QueryString ...
	QueryString MatchVariable = "QueryString"
	// RemoteAddr ...
	RemoteAddr MatchVariable = "RemoteAddr"
	// RequestBody ...
	RequestBody MatchVariable = "RequestBody"
	// RequestHeader ...
	RequestHeader MatchVariable = "RequestHeader"
	// RequestMethod ...
	RequestMethod MatchVariable = "RequestMethod"
	// RequestURI ...
	RequestURI MatchVariable = "RequestUri"
	// SocketAddr ...
	SocketAddr MatchVariable = "SocketAddr"
)

// PossibleMatchVariableValues returns an array of possible values for the MatchVariable const type.
func PossibleMatchVariableValues() []MatchVariable {
	return []MatchVariable{Cookies, PostArgs, QueryString, RemoteAddr, RequestBody, RequestHeader, RequestMethod, RequestURI, SocketAddr}
}

// MinimumTLSVersion enumerates the values for minimum tls version.
type MinimumTLSVersion string

const (
	// OneFullStopTwo ...
	OneFullStopTwo MinimumTLSVersion = "1.2"
	// OneFullStopZero ...
	OneFullStopZero MinimumTLSVersion = "1.0"
)

// PossibleMinimumTLSVersionValues returns an array of possible values for the MinimumTLSVersion const type.
func PossibleMinimumTLSVersionValues() []MinimumTLSVersion {
	return []MinimumTLSVersion{OneFullStopTwo, OneFullStopZero}
}

// NetworkExperimentResourceState enumerates the values for network experiment resource state.
type NetworkExperimentResourceState string

const (
	// NetworkExperimentResourceStateCreating ...
	NetworkExperimentResourceStateCreating NetworkExperimentResourceState = "Creating"
	// NetworkExperimentResourceStateDeleting ...
	NetworkExperimentResourceStateDeleting NetworkExperimentResourceState = "Deleting"
	// NetworkExperimentResourceStateDisabled ...
	NetworkExperimentResourceStateDisabled NetworkExperimentResourceState = "Disabled"
	// NetworkExperimentResourceStateDisabling ...
	NetworkExperimentResourceStateDisabling NetworkExperimentResourceState = "Disabling"
	// NetworkExperimentResourceStateEnabled ...
	NetworkExperimentResourceStateEnabled NetworkExperimentResourceState = "Enabled"
	// NetworkExperimentResourceStateEnabling ...
	NetworkExperimentResourceStateEnabling NetworkExperimentResourceState = "Enabling"
)

// PossibleNetworkExperimentResourceStateValues returns an array of possible values for the NetworkExperimentResourceState const type.
func PossibleNetworkExperimentResourceStateValues() []NetworkExperimentResourceState {
	return []NetworkExperimentResourceState{NetworkExperimentResourceStateCreating, NetworkExperimentResourceStateDeleting, NetworkExperimentResourceStateDisabled, NetworkExperimentResourceStateDisabling, NetworkExperimentResourceStateEnabled, NetworkExperimentResourceStateEnabling}
}

// NetworkOperationStatus enumerates the values for network operation status.
type NetworkOperationStatus string

const (
	// Failed ...
	Failed NetworkOperationStatus = "Failed"
	// InProgress ...
	InProgress NetworkOperationStatus = "InProgress"
	// Succeeded ...
	Succeeded NetworkOperationStatus = "Succeeded"
)

// PossibleNetworkOperationStatusValues returns an array of possible values for the NetworkOperationStatus const type.
func PossibleNetworkOperationStatusValues() []NetworkOperationStatus {
	return []NetworkOperationStatus{Failed, InProgress, Succeeded}
}

// OdataType enumerates the values for odata type.
type OdataType string

const (
	// OdataTypeMicrosoftAzureFrontDoorModelsFrontdoorForwardingConfiguration ...
	OdataTypeMicrosoftAzureFrontDoorModelsFrontdoorForwardingConfiguration OdataType = "#Microsoft.Azure.FrontDoor.Models.FrontdoorForwardingConfiguration"
	// OdataTypeMicrosoftAzureFrontDoorModelsFrontdoorRedirectConfiguration ...
	OdataTypeMicrosoftAzureFrontDoorModelsFrontdoorRedirectConfiguration OdataType = "#Microsoft.Azure.FrontDoor.Models.FrontdoorRedirectConfiguration"
	// OdataTypeRouteConfiguration ...
	OdataTypeRouteConfiguration OdataType = "RouteConfiguration"
)

// PossibleOdataTypeValues returns an array of possible values for the OdataType const type.
func PossibleOdataTypeValues() []OdataType {
	return []OdataType{OdataTypeMicrosoftAzureFrontDoorModelsFrontdoorForwardingConfiguration, OdataTypeMicrosoftAzureFrontDoorModelsFrontdoorRedirectConfiguration, OdataTypeRouteConfiguration}
}

// Operator enumerates the values for operator.
type Operator string

const (
	// OperatorAny ...
	OperatorAny Operator = "Any"
	// OperatorBeginsWith ...
	OperatorBeginsWith Operator = "BeginsWith"
	// OperatorContains ...
	OperatorContains Operator = "Contains"
	// OperatorEndsWith ...
	OperatorEndsWith Operator = "EndsWith"
	// OperatorEqual ...
	OperatorEqual Operator = "Equal"
	// OperatorGeoMatch ...
	OperatorGeoMatch Operator = "GeoMatch"
	// OperatorGreaterThan ...
	OperatorGreaterThan Operator = "GreaterThan"
	// OperatorGreaterThanOrEqual ...
	OperatorGreaterThanOrEqual Operator = "GreaterThanOrEqual"
	// OperatorIPMatch ...
	OperatorIPMatch Operator = "IPMatch"
	// OperatorLessThan ...
	OperatorLessThan Operator = "LessThan"
	// OperatorLessThanOrEqual ...
	OperatorLessThanOrEqual Operator = "LessThanOrEqual"
	// OperatorRegEx ...
	OperatorRegEx Operator = "RegEx"
)

// PossibleOperatorValues returns an array of possible values for the Operator const type.
func PossibleOperatorValues() []Operator {
	return []Operator{OperatorAny, OperatorBeginsWith, OperatorContains, OperatorEndsWith, OperatorEqual, OperatorGeoMatch, OperatorGreaterThan, OperatorGreaterThanOrEqual, OperatorIPMatch, OperatorLessThan, OperatorLessThanOrEqual, OperatorRegEx}
}

// PolicyEnabledState enumerates the values for policy enabled state.
type PolicyEnabledState string

const (
	// PolicyEnabledStateDisabled ...
	PolicyEnabledStateDisabled PolicyEnabledState = "Disabled"
	// PolicyEnabledStateEnabled ...
	PolicyEnabledStateEnabled PolicyEnabledState = "Enabled"
)

// PossiblePolicyEnabledStateValues returns an array of possible values for the PolicyEnabledState const type.
func PossiblePolicyEnabledStateValues() []PolicyEnabledState {
	return []PolicyEnabledState{PolicyEnabledStateDisabled, PolicyEnabledStateEnabled}
}

// PolicyMode enumerates the values for policy mode.
type PolicyMode string

const (
	// Detection ...
	Detection PolicyMode = "Detection"
	// Prevention ...
	Prevention PolicyMode = "Prevention"
)

// PossiblePolicyModeValues returns an array of possible values for the PolicyMode const type.
func PossiblePolicyModeValues() []PolicyMode {
	return []PolicyMode{Detection, Prevention}
}

// PolicyResourceState enumerates the values for policy resource state.
type PolicyResourceState string

const (
	// PolicyResourceStateCreating ...
	PolicyResourceStateCreating PolicyResourceState = "Creating"
	// PolicyResourceStateDeleting ...
	PolicyResourceStateDeleting PolicyResourceState = "Deleting"
	// PolicyResourceStateDisabled ...
	PolicyResourceStateDisabled PolicyResourceState = "Disabled"
	// PolicyResourceStateDisabling ...
	PolicyResourceStateDisabling PolicyResourceState = "Disabling"
	// PolicyResourceStateEnabled ...
	PolicyResourceStateEnabled PolicyResourceState = "Enabled"
	// PolicyResourceStateEnabling ...
	PolicyResourceStateEnabling PolicyResourceState = "Enabling"
)

// PossiblePolicyResourceStateValues returns an array of possible values for the PolicyResourceState const type.
func PossiblePolicyResourceStateValues() []PolicyResourceState {
	return []PolicyResourceState{PolicyResourceStateCreating, PolicyResourceStateDeleting, PolicyResourceStateDisabled, PolicyResourceStateDisabling, PolicyResourceStateEnabled, PolicyResourceStateEnabling}
}

// Protocol enumerates the values for protocol.
type Protocol string

const (
	// HTTP ...
	HTTP Protocol = "Http"
	// HTTPS ...
	HTTPS Protocol = "Https"
)

// PossibleProtocolValues returns an array of possible values for the Protocol const type.
func PossibleProtocolValues() []Protocol {
	return []Protocol{HTTP, HTTPS}
}

// Query enumerates the values for query.
type Query string

const (
	// StripAll ...
	StripAll Query = "StripAll"
	// StripNone ...
	StripNone Query = "StripNone"
)

// PossibleQueryValues returns an array of possible values for the Query const type.
func PossibleQueryValues() []Query {
	return []Query{StripAll, StripNone}
}

// RedirectProtocol enumerates the values for redirect protocol.
type RedirectProtocol string

const (
	// RedirectProtocolHTTPOnly ...
	RedirectProtocolHTTPOnly RedirectProtocol = "HttpOnly"
	// RedirectProtocolHTTPSOnly ...
	RedirectProtocolHTTPSOnly RedirectProtocol = "HttpsOnly"
	// RedirectProtocolMatchRequest ...
	RedirectProtocolMatchRequest RedirectProtocol = "MatchRequest"
)

// PossibleRedirectProtocolValues returns an array of possible values for the RedirectProtocol const type.
func PossibleRedirectProtocolValues() []RedirectProtocol {
	return []RedirectProtocol{RedirectProtocolHTTPOnly, RedirectProtocolHTTPSOnly, RedirectProtocolMatchRequest}
}

// RedirectType enumerates the values for redirect type.
type RedirectType string

const (
	// Found ...
	Found RedirectType = "Found"
	// Moved ...
	Moved RedirectType = "Moved"
	// PermanentRedirect ...
	PermanentRedirect RedirectType = "PermanentRedirect"
	// TemporaryRedirect ...
	TemporaryRedirect RedirectType = "TemporaryRedirect"
)

// PossibleRedirectTypeValues returns an array of possible values for the RedirectType const type.
func PossibleRedirectTypeValues() []RedirectType {
	return []RedirectType{Found, Moved, PermanentRedirect, TemporaryRedirect}
}

// ResourceState enumerates the values for resource state.
type ResourceState string

const (
	// ResourceStateCreating ...
	ResourceStateCreating ResourceState = "Creating"
	// ResourceStateDeleting ...
	ResourceStateDeleting ResourceState = "Deleting"
	// ResourceStateDisabled ...
	ResourceStateDisabled ResourceState = "Disabled"
	// ResourceStateDisabling ...
	ResourceStateDisabling ResourceState = "Disabling"
	// ResourceStateEnabled ...
	ResourceStateEnabled ResourceState = "Enabled"
	// ResourceStateEnabling ...
	ResourceStateEnabling ResourceState = "Enabling"
)

// PossibleResourceStateValues returns an array of possible values for the ResourceState const type.
func PossibleResourceStateValues() []ResourceState {
	return []ResourceState{ResourceStateCreating, ResourceStateDeleting, ResourceStateDisabled, ResourceStateDisabling, ResourceStateEnabled, ResourceStateEnabling}
}

// ResourceType enumerates the values for resource type.
type ResourceType string

const (
	// MicrosoftNetworkfrontDoors ...
	MicrosoftNetworkfrontDoors ResourceType = "Microsoft.Network/frontDoors"
	// MicrosoftNetworkfrontDoorsfrontendEndpoints ...
	MicrosoftNetworkfrontDoorsfrontendEndpoints ResourceType = "Microsoft.Network/frontDoors/frontendEndpoints"
)

// PossibleResourceTypeValues returns an array of possible values for the ResourceType const type.
func PossibleResourceTypeValues() []ResourceType {
	return []ResourceType{MicrosoftNetworkfrontDoors, MicrosoftNetworkfrontDoorsfrontendEndpoints}
}

// RoutingRuleEnabledState enumerates the values for routing rule enabled state.
type RoutingRuleEnabledState string

const (
	// RoutingRuleEnabledStateDisabled ...
	RoutingRuleEnabledStateDisabled RoutingRuleEnabledState = "Disabled"
	// RoutingRuleEnabledStateEnabled ...
	RoutingRuleEnabledStateEnabled RoutingRuleEnabledState = "Enabled"
)

// PossibleRoutingRuleEnabledStateValues returns an array of possible values for the RoutingRuleEnabledState const type.
func PossibleRoutingRuleEnabledStateValues() []RoutingRuleEnabledState {
	return []RoutingRuleEnabledState{RoutingRuleEnabledStateDisabled, RoutingRuleEnabledStateEnabled}
}

// RuleType enumerates the values for rule type.
type RuleType string

const (
	// MatchRule ...
	MatchRule RuleType = "MatchRule"
	// RateLimitRule ...
	RateLimitRule RuleType = "RateLimitRule"
)

// PossibleRuleTypeValues returns an array of possible values for the RuleType const type.
func PossibleRuleTypeValues() []RuleType {
	return []RuleType{MatchRule, RateLimitRule}
}

// SessionAffinityEnabledState enumerates the values for session affinity enabled state.
type SessionAffinityEnabledState string

const (
	// SessionAffinityEnabledStateDisabled ...
	SessionAffinityEnabledStateDisabled SessionAffinityEnabledState = "Disabled"
	// SessionAffinityEnabledStateEnabled ...
	SessionAffinityEnabledStateEnabled SessionAffinityEnabledState = "Enabled"
)

// PossibleSessionAffinityEnabledStateValues returns an array of possible values for the SessionAffinityEnabledState const type.
func PossibleSessionAffinityEnabledStateValues() []SessionAffinityEnabledState {
	return []SessionAffinityEnabledState{SessionAffinityEnabledStateDisabled, SessionAffinityEnabledStateEnabled}
}

// State enumerates the values for state.
type State string

const (
	// StateDisabled ...
	StateDisabled State = "Disabled"
	// StateEnabled ...
	StateEnabled State = "Enabled"
)

// PossibleStateValues returns an array of possible values for the State const type.
func PossibleStateValues() []State {
	return []State{StateDisabled, StateEnabled}
}

// TimeseriesAggregationInterval enumerates the values for timeseries aggregation interval.
type TimeseriesAggregationInterval string

const (
	// TimeseriesAggregationIntervalDaily ...
	TimeseriesAggregationIntervalDaily TimeseriesAggregationInterval = "Daily"
	// TimeseriesAggregationIntervalHourly ...
	TimeseriesAggregationIntervalHourly TimeseriesAggregationInterval = "Hourly"
)

// PossibleTimeseriesAggregationIntervalValues returns an array of possible values for the TimeseriesAggregationInterval const type.
func PossibleTimeseriesAggregationIntervalValues() []TimeseriesAggregationInterval {
	return []TimeseriesAggregationInterval{TimeseriesAggregationIntervalDaily, TimeseriesAggregationIntervalHourly}
}

// TimeseriesType enumerates the values for timeseries type.
type TimeseriesType string

const (
	// LatencyP50 ...
	LatencyP50 TimeseriesType = "LatencyP50"
	// LatencyP75 ...
	LatencyP75 TimeseriesType = "LatencyP75"
	// LatencyP95 ...
	LatencyP95 TimeseriesType = "LatencyP95"
	// MeasurementCounts ...
	MeasurementCounts TimeseriesType = "MeasurementCounts"
)

// PossibleTimeseriesTypeValues returns an array of possible values for the TimeseriesType const type.
func PossibleTimeseriesTypeValues() []TimeseriesType {
	return []TimeseriesType{LatencyP50, LatencyP75, LatencyP95, MeasurementCounts}
}

// TransformType enumerates the values for transform type.
type TransformType string

const (
	// Lowercase ...
	Lowercase TransformType = "Lowercase"
	// RemoveNulls ...
	RemoveNulls TransformType = "RemoveNulls"
	// Trim ...
	Trim TransformType = "Trim"
	// Uppercase ...
	Uppercase TransformType = "Uppercase"
	// URLDecode ...
	URLDecode TransformType = "UrlDecode"
	// URLEncode ...
	URLEncode TransformType = "UrlEncode"
)

// PossibleTransformTypeValues returns an array of possible values for the TransformType const type.
func PossibleTransformTypeValues() []TransformType {
	return []TransformType{Lowercase, RemoveNulls, Trim, Uppercase, URLDecode, URLEncode}
}
