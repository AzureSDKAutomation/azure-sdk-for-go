package hybridkubernetes

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// AuthenticationMethod enumerates the values for authentication method.
type AuthenticationMethod string

const (
	// AAD ...
	AAD AuthenticationMethod = "AAD"
	// Token ...
	Token AuthenticationMethod = "Token"
)

// PossibleAuthenticationMethodValues returns an array of possible values for the AuthenticationMethod const type.
func PossibleAuthenticationMethodValues() []AuthenticationMethod {
	return []AuthenticationMethod{AAD, Token}
}

// ConnectivityStatus enumerates the values for connectivity status.
type ConnectivityStatus string

const (
	// Connected ...
	Connected ConnectivityStatus = "Connected"
	// Connecting ...
	Connecting ConnectivityStatus = "Connecting"
	// Expired ...
	Expired ConnectivityStatus = "Expired"
	// Offline ...
	Offline ConnectivityStatus = "Offline"
)

// PossibleConnectivityStatusValues returns an array of possible values for the ConnectivityStatus const type.
func PossibleConnectivityStatusValues() []ConnectivityStatus {
	return []ConnectivityStatus{Connected, Connecting, Expired, Offline}
}

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

// PrivateLinkState enumerates the values for private link state.
type PrivateLinkState string

const (
	// Disabled ...
	Disabled PrivateLinkState = "Disabled"
	// Enabled ...
	Enabled PrivateLinkState = "Enabled"
)

// PossiblePrivateLinkStateValues returns an array of possible values for the PrivateLinkState const type.
func PossiblePrivateLinkStateValues() []PrivateLinkState {
	return []PrivateLinkState{Disabled, Enabled}
}

// ProvisioningState enumerates the values for provisioning state.
type ProvisioningState string

const (
	// Accepted ...
	Accepted ProvisioningState = "Accepted"
	// Canceled ...
	Canceled ProvisioningState = "Canceled"
	// Deleting ...
	Deleting ProvisioningState = "Deleting"
	// Failed ...
	Failed ProvisioningState = "Failed"
	// Provisioning ...
	Provisioning ProvisioningState = "Provisioning"
	// Succeeded ...
	Succeeded ProvisioningState = "Succeeded"
	// Updating ...
	Updating ProvisioningState = "Updating"
)

// PossibleProvisioningStateValues returns an array of possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{Accepted, Canceled, Deleting, Failed, Provisioning, Succeeded, Updating}
}

// ResourceIdentityType enumerates the values for resource identity type.
type ResourceIdentityType string

const (
	// None ...
	None ResourceIdentityType = "None"
	// SystemAssigned ...
	SystemAssigned ResourceIdentityType = "SystemAssigned"
)

// PossibleResourceIdentityTypeValues returns an array of possible values for the ResourceIdentityType const type.
func PossibleResourceIdentityTypeValues() []ResourceIdentityType {
	return []ResourceIdentityType{None, SystemAssigned}
}
