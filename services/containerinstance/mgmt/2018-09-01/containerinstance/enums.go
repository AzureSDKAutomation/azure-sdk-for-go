package containerinstance

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// ContainerGroupIPAddressType enumerates the values for container group ip address type.
type ContainerGroupIPAddressType string

const (
	// Private ...
	Private ContainerGroupIPAddressType = "Private"
	// Public ...
	Public ContainerGroupIPAddressType = "Public"
)

// PossibleContainerGroupIPAddressTypeValues returns an array of possible values for the ContainerGroupIPAddressType const type.
func PossibleContainerGroupIPAddressTypeValues() []ContainerGroupIPAddressType {
	return []ContainerGroupIPAddressType{Private, Public}
}

// ContainerGroupNetworkProtocol enumerates the values for container group network protocol.
type ContainerGroupNetworkProtocol string

const (
	// TCP ...
	TCP ContainerGroupNetworkProtocol = "TCP"
	// UDP ...
	UDP ContainerGroupNetworkProtocol = "UDP"
)

// PossibleContainerGroupNetworkProtocolValues returns an array of possible values for the ContainerGroupNetworkProtocol const type.
func PossibleContainerGroupNetworkProtocolValues() []ContainerGroupNetworkProtocol {
	return []ContainerGroupNetworkProtocol{TCP, UDP}
}

// ContainerGroupRestartPolicy enumerates the values for container group restart policy.
type ContainerGroupRestartPolicy string

const (
	// Always ...
	Always ContainerGroupRestartPolicy = "Always"
	// Never ...
	Never ContainerGroupRestartPolicy = "Never"
	// OnFailure ...
	OnFailure ContainerGroupRestartPolicy = "OnFailure"
)

// PossibleContainerGroupRestartPolicyValues returns an array of possible values for the ContainerGroupRestartPolicy const type.
func PossibleContainerGroupRestartPolicyValues() []ContainerGroupRestartPolicy {
	return []ContainerGroupRestartPolicy{Always, Never, OnFailure}
}

// ContainerNetworkProtocol enumerates the values for container network protocol.
type ContainerNetworkProtocol string

const (
	// ContainerNetworkProtocolTCP ...
	ContainerNetworkProtocolTCP ContainerNetworkProtocol = "TCP"
	// ContainerNetworkProtocolUDP ...
	ContainerNetworkProtocolUDP ContainerNetworkProtocol = "UDP"
)

// PossibleContainerNetworkProtocolValues returns an array of possible values for the ContainerNetworkProtocol const type.
func PossibleContainerNetworkProtocolValues() []ContainerNetworkProtocol {
	return []ContainerNetworkProtocol{ContainerNetworkProtocolTCP, ContainerNetworkProtocolUDP}
}

// LogAnalyticsLogType enumerates the values for log analytics log type.
type LogAnalyticsLogType string

const (
	// ContainerInsights ...
	ContainerInsights LogAnalyticsLogType = "ContainerInsights"
	// ContainerInstanceLogs ...
	ContainerInstanceLogs LogAnalyticsLogType = "ContainerInstanceLogs"
)

// PossibleLogAnalyticsLogTypeValues returns an array of possible values for the LogAnalyticsLogType const type.
func PossibleLogAnalyticsLogTypeValues() []LogAnalyticsLogType {
	return []LogAnalyticsLogType{ContainerInsights, ContainerInstanceLogs}
}

// OperatingSystemTypes enumerates the values for operating system types.
type OperatingSystemTypes string

const (
	// Linux ...
	Linux OperatingSystemTypes = "Linux"
	// Windows ...
	Windows OperatingSystemTypes = "Windows"
)

// PossibleOperatingSystemTypesValues returns an array of possible values for the OperatingSystemTypes const type.
func PossibleOperatingSystemTypesValues() []OperatingSystemTypes {
	return []OperatingSystemTypes{Linux, Windows}
}

// OperationsOrigin enumerates the values for operations origin.
type OperationsOrigin string

const (
	// System ...
	System OperationsOrigin = "System"
	// User ...
	User OperationsOrigin = "User"
)

// PossibleOperationsOriginValues returns an array of possible values for the OperationsOrigin const type.
func PossibleOperationsOriginValues() []OperationsOrigin {
	return []OperationsOrigin{System, User}
}

// Scheme enumerates the values for scheme.
type Scheme string

const (
	// HTTP ...
	HTTP Scheme = "http"
	// HTTPS ...
	HTTPS Scheme = "https"
)

// PossibleSchemeValues returns an array of possible values for the Scheme const type.
func PossibleSchemeValues() []Scheme {
	return []Scheme{HTTP, HTTPS}
}
