package purview

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// CreatedByType enumerates the values for created by type.
type CreatedByType string

const (
	// Application ...
	Application CreatedByType = "Application"
	// Key ...
	Key CreatedByType = "Key"
	// ManagedIdentity ...
	ManagedIdentity CreatedByType = "ManagedIdentity"
	// User ...
	User CreatedByType = "User"
)

// PossibleCreatedByTypeValues returns an array of possible values for the CreatedByType const type.
func PossibleCreatedByTypeValues() []CreatedByType {
	return []CreatedByType{Application, Key, ManagedIdentity, User}
}

// LastModifiedByType enumerates the values for last modified by type.
type LastModifiedByType string

const (
	// LastModifiedByTypeApplication ...
	LastModifiedByTypeApplication LastModifiedByType = "Application"
	// LastModifiedByTypeKey ...
	LastModifiedByTypeKey LastModifiedByType = "Key"
	// LastModifiedByTypeManagedIdentity ...
	LastModifiedByTypeManagedIdentity LastModifiedByType = "ManagedIdentity"
	// LastModifiedByTypeUser ...
	LastModifiedByTypeUser LastModifiedByType = "User"
)

// PossibleLastModifiedByTypeValues returns an array of possible values for the LastModifiedByType const type.
func PossibleLastModifiedByTypeValues() []LastModifiedByType {
	return []LastModifiedByType{LastModifiedByTypeApplication, LastModifiedByTypeKey, LastModifiedByTypeManagedIdentity, LastModifiedByTypeUser}
}

// Name enumerates the values for name.
type Name string

const (
	// Standard ...
	Standard Name = "Standard"
)

// PossibleNameValues returns an array of possible values for the Name const type.
func PossibleNameValues() []Name {
	return []Name{Standard}
}

// ProvisioningState enumerates the values for provisioning state.
type ProvisioningState string

const (
	// Canceled ...
	Canceled ProvisioningState = "Canceled"
	// Creating ...
	Creating ProvisioningState = "Creating"
	// Deleting ...
	Deleting ProvisioningState = "Deleting"
	// Failed ...
	Failed ProvisioningState = "Failed"
	// Moving ...
	Moving ProvisioningState = "Moving"
	// SoftDeleted ...
	SoftDeleted ProvisioningState = "SoftDeleted"
	// SoftDeleting ...
	SoftDeleting ProvisioningState = "SoftDeleting"
	// Succeeded ...
	Succeeded ProvisioningState = "Succeeded"
	// Unknown ...
	Unknown ProvisioningState = "Unknown"
)

// PossibleProvisioningStateValues returns an array of possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{Canceled, Creating, Deleting, Failed, Moving, SoftDeleted, SoftDeleting, Succeeded, Unknown}
}

// PublicNetworkAccess enumerates the values for public network access.
type PublicNetworkAccess string

const (
	// Disabled ...
	Disabled PublicNetworkAccess = "Disabled"
	// Enabled ...
	Enabled PublicNetworkAccess = "Enabled"
	// NotSpecified ...
	NotSpecified PublicNetworkAccess = "NotSpecified"
)

// PossiblePublicNetworkAccessValues returns an array of possible values for the PublicNetworkAccess const type.
func PossiblePublicNetworkAccessValues() []PublicNetworkAccess {
	return []PublicNetworkAccess{Disabled, Enabled, NotSpecified}
}

// Reason enumerates the values for reason.
type Reason string

const (
	// AlreadyExists ...
	AlreadyExists Reason = "AlreadyExists"
	// Invalid ...
	Invalid Reason = "Invalid"
)

// PossibleReasonValues returns an array of possible values for the Reason const type.
func PossibleReasonValues() []Reason {
	return []Reason{AlreadyExists, Invalid}
}

// ScopeType enumerates the values for scope type.
type ScopeType string

const (
	// Subscription ...
	Subscription ScopeType = "Subscription"
	// Tenant ...
	Tenant ScopeType = "Tenant"
)

// PossibleScopeTypeValues returns an array of possible values for the ScopeType const type.
func PossibleScopeTypeValues() []ScopeType {
	return []ScopeType{Subscription, Tenant}
}

// Status enumerates the values for status.
type Status string

const (
	// StatusApproved ...
	StatusApproved Status = "Approved"
	// StatusDisconnected ...
	StatusDisconnected Status = "Disconnected"
	// StatusPending ...
	StatusPending Status = "Pending"
	// StatusRejected ...
	StatusRejected Status = "Rejected"
	// StatusUnknown ...
	StatusUnknown Status = "Unknown"
)

// PossibleStatusValues returns an array of possible values for the Status const type.
func PossibleStatusValues() []Status {
	return []Status{StatusApproved, StatusDisconnected, StatusPending, StatusRejected, StatusUnknown}
}

// Type enumerates the values for type.
type Type string

const (
	// SystemAssigned ...
	SystemAssigned Type = "SystemAssigned"
)

// PossibleTypeValues returns an array of possible values for the Type const type.
func PossibleTypeValues() []Type {
	return []Type{SystemAssigned}
}
