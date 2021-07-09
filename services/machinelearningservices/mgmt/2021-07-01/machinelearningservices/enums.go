package machinelearningservices

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// AllocationState enumerates the values for allocation state.
type AllocationState string

const (
	// Resizing ...
	Resizing AllocationState = "Resizing"
	// Steady ...
	Steady AllocationState = "Steady"
)

// PossibleAllocationStateValues returns an array of possible values for the AllocationState const type.
func PossibleAllocationStateValues() []AllocationState {
	return []AllocationState{Resizing, Steady}
}

// ApplicationSharingPolicy enumerates the values for application sharing policy.
type ApplicationSharingPolicy string

const (
	// Personal ...
	Personal ApplicationSharingPolicy = "Personal"
	// Shared ...
	Shared ApplicationSharingPolicy = "Shared"
)

// PossibleApplicationSharingPolicyValues returns an array of possible values for the ApplicationSharingPolicy const type.
func PossibleApplicationSharingPolicyValues() []ApplicationSharingPolicy {
	return []ApplicationSharingPolicy{Personal, Shared}
}

// ClusterPurpose enumerates the values for cluster purpose.
type ClusterPurpose string

const (
	// DenseProd ...
	DenseProd ClusterPurpose = "DenseProd"
	// DevTest ...
	DevTest ClusterPurpose = "DevTest"
	// FastProd ...
	FastProd ClusterPurpose = "FastProd"
)

// PossibleClusterPurposeValues returns an array of possible values for the ClusterPurpose const type.
func PossibleClusterPurposeValues() []ClusterPurpose {
	return []ClusterPurpose{DenseProd, DevTest, FastProd}
}

// ComputeInstanceAuthorizationType enumerates the values for compute instance authorization type.
type ComputeInstanceAuthorizationType string

const (
	// ComputeInstanceAuthorizationTypePersonal ...
	ComputeInstanceAuthorizationTypePersonal ComputeInstanceAuthorizationType = "personal"
)

// PossibleComputeInstanceAuthorizationTypeValues returns an array of possible values for the ComputeInstanceAuthorizationType const type.
func PossibleComputeInstanceAuthorizationTypeValues() []ComputeInstanceAuthorizationType {
	return []ComputeInstanceAuthorizationType{ComputeInstanceAuthorizationTypePersonal}
}

// ComputeInstanceState enumerates the values for compute instance state.
type ComputeInstanceState string

const (
	// CreateFailed ...
	CreateFailed ComputeInstanceState = "CreateFailed"
	// Creating ...
	Creating ComputeInstanceState = "Creating"
	// Deleting ...
	Deleting ComputeInstanceState = "Deleting"
	// JobRunning ...
	JobRunning ComputeInstanceState = "JobRunning"
	// Restarting ...
	Restarting ComputeInstanceState = "Restarting"
	// Running ...
	Running ComputeInstanceState = "Running"
	// SettingUp ...
	SettingUp ComputeInstanceState = "SettingUp"
	// SetupFailed ...
	SetupFailed ComputeInstanceState = "SetupFailed"
	// Starting ...
	Starting ComputeInstanceState = "Starting"
	// Stopped ...
	Stopped ComputeInstanceState = "Stopped"
	// Stopping ...
	Stopping ComputeInstanceState = "Stopping"
	// Unknown ...
	Unknown ComputeInstanceState = "Unknown"
	// Unusable ...
	Unusable ComputeInstanceState = "Unusable"
	// UserSettingUp ...
	UserSettingUp ComputeInstanceState = "UserSettingUp"
	// UserSetupFailed ...
	UserSetupFailed ComputeInstanceState = "UserSetupFailed"
)

// PossibleComputeInstanceStateValues returns an array of possible values for the ComputeInstanceState const type.
func PossibleComputeInstanceStateValues() []ComputeInstanceState {
	return []ComputeInstanceState{CreateFailed, Creating, Deleting, JobRunning, Restarting, Running, SettingUp, SetupFailed, Starting, Stopped, Stopping, Unknown, Unusable, UserSettingUp, UserSetupFailed}
}

// ComputeType enumerates the values for compute type.
type ComputeType string

const (
	// ComputeTypeAKS ...
	ComputeTypeAKS ComputeType = "AKS"
	// ComputeTypeAmlCompute ...
	ComputeTypeAmlCompute ComputeType = "AmlCompute"
	// ComputeTypeComputeInstance ...
	ComputeTypeComputeInstance ComputeType = "ComputeInstance"
	// ComputeTypeDatabricks ...
	ComputeTypeDatabricks ComputeType = "Databricks"
	// ComputeTypeDataFactory ...
	ComputeTypeDataFactory ComputeType = "DataFactory"
	// ComputeTypeDataLakeAnalytics ...
	ComputeTypeDataLakeAnalytics ComputeType = "DataLakeAnalytics"
	// ComputeTypeHDInsight ...
	ComputeTypeHDInsight ComputeType = "HDInsight"
	// ComputeTypeKubernetes ...
	ComputeTypeKubernetes ComputeType = "Kubernetes"
	// ComputeTypeSynapseSpark ...
	ComputeTypeSynapseSpark ComputeType = "SynapseSpark"
	// ComputeTypeVirtualMachine ...
	ComputeTypeVirtualMachine ComputeType = "VirtualMachine"
)

// PossibleComputeTypeValues returns an array of possible values for the ComputeType const type.
func PossibleComputeTypeValues() []ComputeType {
	return []ComputeType{ComputeTypeAKS, ComputeTypeAmlCompute, ComputeTypeComputeInstance, ComputeTypeDatabricks, ComputeTypeDataFactory, ComputeTypeDataLakeAnalytics, ComputeTypeHDInsight, ComputeTypeKubernetes, ComputeTypeSynapseSpark, ComputeTypeVirtualMachine}
}

// ComputeTypeBasicCompute enumerates the values for compute type basic compute.
type ComputeTypeBasicCompute string

const (
	// ComputeTypeAKS1 ...
	ComputeTypeAKS1 ComputeTypeBasicCompute = "AKS"
	// ComputeTypeAmlCompute1 ...
	ComputeTypeAmlCompute1 ComputeTypeBasicCompute = "AmlCompute"
	// ComputeTypeCompute ...
	ComputeTypeCompute ComputeTypeBasicCompute = "Compute"
	// ComputeTypeComputeInstance1 ...
	ComputeTypeComputeInstance1 ComputeTypeBasicCompute = "ComputeInstance"
	// ComputeTypeDatabricks1 ...
	ComputeTypeDatabricks1 ComputeTypeBasicCompute = "Databricks"
	// ComputeTypeDataFactory1 ...
	ComputeTypeDataFactory1 ComputeTypeBasicCompute = "DataFactory"
	// ComputeTypeDataLakeAnalytics1 ...
	ComputeTypeDataLakeAnalytics1 ComputeTypeBasicCompute = "DataLakeAnalytics"
	// ComputeTypeHDInsight1 ...
	ComputeTypeHDInsight1 ComputeTypeBasicCompute = "HDInsight"
	// ComputeTypeSynapseSpark1 ...
	ComputeTypeSynapseSpark1 ComputeTypeBasicCompute = "SynapseSpark"
	// ComputeTypeVirtualMachine1 ...
	ComputeTypeVirtualMachine1 ComputeTypeBasicCompute = "VirtualMachine"
)

// PossibleComputeTypeBasicComputeValues returns an array of possible values for the ComputeTypeBasicCompute const type.
func PossibleComputeTypeBasicComputeValues() []ComputeTypeBasicCompute {
	return []ComputeTypeBasicCompute{ComputeTypeAKS1, ComputeTypeAmlCompute1, ComputeTypeCompute, ComputeTypeComputeInstance1, ComputeTypeDatabricks1, ComputeTypeDataFactory1, ComputeTypeDataLakeAnalytics1, ComputeTypeHDInsight1, ComputeTypeSynapseSpark1, ComputeTypeVirtualMachine1}
}

// ComputeTypeBasicComputeNodesInformation enumerates the values for compute type basic compute nodes
// information.
type ComputeTypeBasicComputeNodesInformation string

const (
	// ComputeTypeComputeNodesInformation ...
	ComputeTypeComputeNodesInformation ComputeTypeBasicComputeNodesInformation = "ComputeNodesInformation"
)

// PossibleComputeTypeBasicComputeNodesInformationValues returns an array of possible values for the ComputeTypeBasicComputeNodesInformation const type.
func PossibleComputeTypeBasicComputeNodesInformationValues() []ComputeTypeBasicComputeNodesInformation {
	return []ComputeTypeBasicComputeNodesInformation{ComputeTypeComputeNodesInformation}
}

// ComputeTypeBasicComputeSecrets enumerates the values for compute type basic compute secrets.
type ComputeTypeBasicComputeSecrets string

const (
	// ComputeTypeBasicComputeSecretsComputeTypeComputeSecrets ...
	ComputeTypeBasicComputeSecretsComputeTypeComputeSecrets ComputeTypeBasicComputeSecrets = "ComputeSecrets"
	// ComputeTypeBasicComputeSecretsComputeTypeVirtualMachine ...
	ComputeTypeBasicComputeSecretsComputeTypeVirtualMachine ComputeTypeBasicComputeSecrets = "VirtualMachine"
)

// PossibleComputeTypeBasicComputeSecretsValues returns an array of possible values for the ComputeTypeBasicComputeSecrets const type.
func PossibleComputeTypeBasicComputeSecretsValues() []ComputeTypeBasicComputeSecrets {
	return []ComputeTypeBasicComputeSecrets{ComputeTypeBasicComputeSecretsComputeTypeComputeSecrets, ComputeTypeBasicComputeSecretsComputeTypeVirtualMachine}
}

// ComputeTypeBasicCreateServiceRequest enumerates the values for compute type basic create service request.
type ComputeTypeBasicCreateServiceRequest string

const (
	// ComputeTypeBasicCreateServiceRequestComputeTypeACI ...
	ComputeTypeBasicCreateServiceRequestComputeTypeACI ComputeTypeBasicCreateServiceRequest = "ACI"
	// ComputeTypeBasicCreateServiceRequestComputeTypeAKS ...
	ComputeTypeBasicCreateServiceRequestComputeTypeAKS ComputeTypeBasicCreateServiceRequest = "AKS"
	// ComputeTypeBasicCreateServiceRequestComputeTypeCreateServiceRequest ...
	ComputeTypeBasicCreateServiceRequestComputeTypeCreateServiceRequest ComputeTypeBasicCreateServiceRequest = "CreateServiceRequest"
	// ComputeTypeBasicCreateServiceRequestComputeTypeCustom ...
	ComputeTypeBasicCreateServiceRequestComputeTypeCustom ComputeTypeBasicCreateServiceRequest = "Custom"
)

// PossibleComputeTypeBasicCreateServiceRequestValues returns an array of possible values for the ComputeTypeBasicCreateServiceRequest const type.
func PossibleComputeTypeBasicCreateServiceRequestValues() []ComputeTypeBasicCreateServiceRequest {
	return []ComputeTypeBasicCreateServiceRequest{ComputeTypeBasicCreateServiceRequestComputeTypeACI, ComputeTypeBasicCreateServiceRequestComputeTypeAKS, ComputeTypeBasicCreateServiceRequestComputeTypeCreateServiceRequest, ComputeTypeBasicCreateServiceRequestComputeTypeCustom}
}

// ComputeTypeBasicServiceResponseBase enumerates the values for compute type basic service response base.
type ComputeTypeBasicServiceResponseBase string

const (
	// ComputeTypeBasicServiceResponseBaseComputeTypeACI ...
	ComputeTypeBasicServiceResponseBaseComputeTypeACI ComputeTypeBasicServiceResponseBase = "ACI"
	// ComputeTypeBasicServiceResponseBaseComputeTypeAKS ...
	ComputeTypeBasicServiceResponseBaseComputeTypeAKS ComputeTypeBasicServiceResponseBase = "AKS"
	// ComputeTypeBasicServiceResponseBaseComputeTypeCustom ...
	ComputeTypeBasicServiceResponseBaseComputeTypeCustom ComputeTypeBasicServiceResponseBase = "Custom"
	// ComputeTypeBasicServiceResponseBaseComputeTypeServiceResponseBase ...
	ComputeTypeBasicServiceResponseBaseComputeTypeServiceResponseBase ComputeTypeBasicServiceResponseBase = "ServiceResponseBase"
)

// PossibleComputeTypeBasicServiceResponseBaseValues returns an array of possible values for the ComputeTypeBasicServiceResponseBase const type.
func PossibleComputeTypeBasicServiceResponseBaseValues() []ComputeTypeBasicServiceResponseBase {
	return []ComputeTypeBasicServiceResponseBase{ComputeTypeBasicServiceResponseBaseComputeTypeACI, ComputeTypeBasicServiceResponseBaseComputeTypeAKS, ComputeTypeBasicServiceResponseBaseComputeTypeCustom, ComputeTypeBasicServiceResponseBaseComputeTypeServiceResponseBase}
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

// DeploymentType enumerates the values for deployment type.
type DeploymentType string

const (
	// Batch ...
	Batch DeploymentType = "Batch"
	// GRPCRealtimeEndpoint ...
	GRPCRealtimeEndpoint DeploymentType = "GRPCRealtimeEndpoint"
	// HTTPRealtimeEndpoint ...
	HTTPRealtimeEndpoint DeploymentType = "HttpRealtimeEndpoint"
)

// PossibleDeploymentTypeValues returns an array of possible values for the DeploymentType const type.
func PossibleDeploymentTypeValues() []DeploymentType {
	return []DeploymentType{Batch, GRPCRealtimeEndpoint, HTTPRealtimeEndpoint}
}

// DiagnoseResultLevel enumerates the values for diagnose result level.
type DiagnoseResultLevel string

const (
	// Error ...
	Error DiagnoseResultLevel = "Error"
	// Warning ...
	Warning DiagnoseResultLevel = "Warning"
)

// PossibleDiagnoseResultLevelValues returns an array of possible values for the DiagnoseResultLevel const type.
func PossibleDiagnoseResultLevelValues() []DiagnoseResultLevel {
	return []DiagnoseResultLevel{Error, Warning}
}

// EncryptionStatus enumerates the values for encryption status.
type EncryptionStatus string

const (
	// Disabled ...
	Disabled EncryptionStatus = "Disabled"
	// Enabled ...
	Enabled EncryptionStatus = "Enabled"
)

// PossibleEncryptionStatusValues returns an array of possible values for the EncryptionStatus const type.
func PossibleEncryptionStatusValues() []EncryptionStatus {
	return []EncryptionStatus{Disabled, Enabled}
}

// LoadBalancerType enumerates the values for load balancer type.
type LoadBalancerType string

const (
	// InternalLoadBalancer ...
	InternalLoadBalancer LoadBalancerType = "InternalLoadBalancer"
	// PublicIP ...
	PublicIP LoadBalancerType = "PublicIp"
)

// PossibleLoadBalancerTypeValues returns an array of possible values for the LoadBalancerType const type.
func PossibleLoadBalancerTypeValues() []LoadBalancerType {
	return []LoadBalancerType{InternalLoadBalancer, PublicIP}
}

// NodeState enumerates the values for node state.
type NodeState string

const (
	// NodeStateIdle ...
	NodeStateIdle NodeState = "idle"
	// NodeStateLeaving ...
	NodeStateLeaving NodeState = "leaving"
	// NodeStatePreempted ...
	NodeStatePreempted NodeState = "preempted"
	// NodeStatePreparing ...
	NodeStatePreparing NodeState = "preparing"
	// NodeStateRunning ...
	NodeStateRunning NodeState = "running"
	// NodeStateUnusable ...
	NodeStateUnusable NodeState = "unusable"
)

// PossibleNodeStateValues returns an array of possible values for the NodeState const type.
func PossibleNodeStateValues() []NodeState {
	return []NodeState{NodeStateIdle, NodeStateLeaving, NodeStatePreempted, NodeStatePreparing, NodeStateRunning, NodeStateUnusable}
}

// OperationName enumerates the values for operation name.
type OperationName string

const (
	// Create ...
	Create OperationName = "Create"
	// Delete ...
	Delete OperationName = "Delete"
	// Reimage ...
	Reimage OperationName = "Reimage"
	// Restart ...
	Restart OperationName = "Restart"
	// Start ...
	Start OperationName = "Start"
	// Stop ...
	Stop OperationName = "Stop"
)

// PossibleOperationNameValues returns an array of possible values for the OperationName const type.
func PossibleOperationNameValues() []OperationName {
	return []OperationName{Create, Delete, Reimage, Restart, Start, Stop}
}

// OperationStatus enumerates the values for operation status.
type OperationStatus string

const (
	// OperationStatusCreateFailed ...
	OperationStatusCreateFailed OperationStatus = "CreateFailed"
	// OperationStatusDeleteFailed ...
	OperationStatusDeleteFailed OperationStatus = "DeleteFailed"
	// OperationStatusInProgress ...
	OperationStatusInProgress OperationStatus = "InProgress"
	// OperationStatusReimageFailed ...
	OperationStatusReimageFailed OperationStatus = "ReimageFailed"
	// OperationStatusRestartFailed ...
	OperationStatusRestartFailed OperationStatus = "RestartFailed"
	// OperationStatusStartFailed ...
	OperationStatusStartFailed OperationStatus = "StartFailed"
	// OperationStatusStopFailed ...
	OperationStatusStopFailed OperationStatus = "StopFailed"
	// OperationStatusSucceeded ...
	OperationStatusSucceeded OperationStatus = "Succeeded"
)

// PossibleOperationStatusValues returns an array of possible values for the OperationStatus const type.
func PossibleOperationStatusValues() []OperationStatus {
	return []OperationStatus{OperationStatusCreateFailed, OperationStatusDeleteFailed, OperationStatusInProgress, OperationStatusReimageFailed, OperationStatusRestartFailed, OperationStatusStartFailed, OperationStatusStopFailed, OperationStatusSucceeded}
}

// OrderString enumerates the values for order string.
type OrderString string

const (
	// CreatedAtAsc ...
	CreatedAtAsc OrderString = "CreatedAtAsc"
	// CreatedAtDesc ...
	CreatedAtDesc OrderString = "CreatedAtDesc"
	// UpdatedAtAsc ...
	UpdatedAtAsc OrderString = "UpdatedAtAsc"
	// UpdatedAtDesc ...
	UpdatedAtDesc OrderString = "UpdatedAtDesc"
)

// PossibleOrderStringValues returns an array of possible values for the OrderString const type.
func PossibleOrderStringValues() []OrderString {
	return []OrderString{CreatedAtAsc, CreatedAtDesc, UpdatedAtAsc, UpdatedAtDesc}
}

// OsType enumerates the values for os type.
type OsType string

const (
	// Linux ...
	Linux OsType = "Linux"
	// Windows ...
	Windows OsType = "Windows"
)

// PossibleOsTypeValues returns an array of possible values for the OsType const type.
func PossibleOsTypeValues() []OsType {
	return []OsType{Linux, Windows}
}

// PrivateEndpointConnectionProvisioningState enumerates the values for private endpoint connection
// provisioning state.
type PrivateEndpointConnectionProvisioningState string

const (
	// PrivateEndpointConnectionProvisioningStateCreating ...
	PrivateEndpointConnectionProvisioningStateCreating PrivateEndpointConnectionProvisioningState = "Creating"
	// PrivateEndpointConnectionProvisioningStateDeleting ...
	PrivateEndpointConnectionProvisioningStateDeleting PrivateEndpointConnectionProvisioningState = "Deleting"
	// PrivateEndpointConnectionProvisioningStateFailed ...
	PrivateEndpointConnectionProvisioningStateFailed PrivateEndpointConnectionProvisioningState = "Failed"
	// PrivateEndpointConnectionProvisioningStateSucceeded ...
	PrivateEndpointConnectionProvisioningStateSucceeded PrivateEndpointConnectionProvisioningState = "Succeeded"
)

// PossiblePrivateEndpointConnectionProvisioningStateValues returns an array of possible values for the PrivateEndpointConnectionProvisioningState const type.
func PossiblePrivateEndpointConnectionProvisioningStateValues() []PrivateEndpointConnectionProvisioningState {
	return []PrivateEndpointConnectionProvisioningState{PrivateEndpointConnectionProvisioningStateCreating, PrivateEndpointConnectionProvisioningStateDeleting, PrivateEndpointConnectionProvisioningStateFailed, PrivateEndpointConnectionProvisioningStateSucceeded}
}

// PrivateEndpointServiceConnectionStatus enumerates the values for private endpoint service connection status.
type PrivateEndpointServiceConnectionStatus string

const (
	// Approved ...
	Approved PrivateEndpointServiceConnectionStatus = "Approved"
	// Disconnected ...
	Disconnected PrivateEndpointServiceConnectionStatus = "Disconnected"
	// Pending ...
	Pending PrivateEndpointServiceConnectionStatus = "Pending"
	// Rejected ...
	Rejected PrivateEndpointServiceConnectionStatus = "Rejected"
	// Timeout ...
	Timeout PrivateEndpointServiceConnectionStatus = "Timeout"
)

// PossiblePrivateEndpointServiceConnectionStatusValues returns an array of possible values for the PrivateEndpointServiceConnectionStatus const type.
func PossiblePrivateEndpointServiceConnectionStatusValues() []PrivateEndpointServiceConnectionStatus {
	return []PrivateEndpointServiceConnectionStatus{Approved, Disconnected, Pending, Rejected, Timeout}
}

// ProvisioningState enumerates the values for provisioning state.
type ProvisioningState string

const (
	// ProvisioningStateCanceled ...
	ProvisioningStateCanceled ProvisioningState = "Canceled"
	// ProvisioningStateCreating ...
	ProvisioningStateCreating ProvisioningState = "Creating"
	// ProvisioningStateDeleting ...
	ProvisioningStateDeleting ProvisioningState = "Deleting"
	// ProvisioningStateFailed ...
	ProvisioningStateFailed ProvisioningState = "Failed"
	// ProvisioningStateSucceeded ...
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
	// ProvisioningStateUnknown ...
	ProvisioningStateUnknown ProvisioningState = "Unknown"
	// ProvisioningStateUpdating ...
	ProvisioningStateUpdating ProvisioningState = "Updating"
)

// PossibleProvisioningStateValues returns an array of possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{ProvisioningStateCanceled, ProvisioningStateCreating, ProvisioningStateDeleting, ProvisioningStateFailed, ProvisioningStateSucceeded, ProvisioningStateUnknown, ProvisioningStateUpdating}
}

// PublicNetworkAccess enumerates the values for public network access.
type PublicNetworkAccess string

const (
	// PublicNetworkAccessDisabled ...
	PublicNetworkAccessDisabled PublicNetworkAccess = "Disabled"
	// PublicNetworkAccessEnabled ...
	PublicNetworkAccessEnabled PublicNetworkAccess = "Enabled"
)

// PossiblePublicNetworkAccessValues returns an array of possible values for the PublicNetworkAccess const type.
func PossiblePublicNetworkAccessValues() []PublicNetworkAccess {
	return []PublicNetworkAccess{PublicNetworkAccessDisabled, PublicNetworkAccessEnabled}
}

// QuotaUnit enumerates the values for quota unit.
type QuotaUnit string

const (
	// Count ...
	Count QuotaUnit = "Count"
)

// PossibleQuotaUnitValues returns an array of possible values for the QuotaUnit const type.
func PossibleQuotaUnitValues() []QuotaUnit {
	return []QuotaUnit{Count}
}

// ReasonCode enumerates the values for reason code.
type ReasonCode string

const (
	// NotAvailableForRegion ...
	NotAvailableForRegion ReasonCode = "NotAvailableForRegion"
	// NotAvailableForSubscription ...
	NotAvailableForSubscription ReasonCode = "NotAvailableForSubscription"
	// NotSpecified ...
	NotSpecified ReasonCode = "NotSpecified"
)

// PossibleReasonCodeValues returns an array of possible values for the ReasonCode const type.
func PossibleReasonCodeValues() []ReasonCode {
	return []ReasonCode{NotAvailableForRegion, NotAvailableForSubscription, NotSpecified}
}

// RemoteLoginPortPublicAccess enumerates the values for remote login port public access.
type RemoteLoginPortPublicAccess string

const (
	// RemoteLoginPortPublicAccessDisabled ...
	RemoteLoginPortPublicAccessDisabled RemoteLoginPortPublicAccess = "Disabled"
	// RemoteLoginPortPublicAccessEnabled ...
	RemoteLoginPortPublicAccessEnabled RemoteLoginPortPublicAccess = "Enabled"
	// RemoteLoginPortPublicAccessNotSpecified ...
	RemoteLoginPortPublicAccessNotSpecified RemoteLoginPortPublicAccess = "NotSpecified"
)

// PossibleRemoteLoginPortPublicAccessValues returns an array of possible values for the RemoteLoginPortPublicAccess const type.
func PossibleRemoteLoginPortPublicAccessValues() []RemoteLoginPortPublicAccess {
	return []RemoteLoginPortPublicAccess{RemoteLoginPortPublicAccessDisabled, RemoteLoginPortPublicAccessEnabled, RemoteLoginPortPublicAccessNotSpecified}
}

// ResourceIdentityType enumerates the values for resource identity type.
type ResourceIdentityType string

const (
	// None ...
	None ResourceIdentityType = "None"
	// SystemAssigned ...
	SystemAssigned ResourceIdentityType = "SystemAssigned"
	// SystemAssignedUserAssigned ...
	SystemAssignedUserAssigned ResourceIdentityType = "SystemAssigned,UserAssigned"
	// UserAssigned ...
	UserAssigned ResourceIdentityType = "UserAssigned"
)

// PossibleResourceIdentityTypeValues returns an array of possible values for the ResourceIdentityType const type.
func PossibleResourceIdentityTypeValues() []ResourceIdentityType {
	return []ResourceIdentityType{None, SystemAssigned, SystemAssignedUserAssigned, UserAssigned}
}

// SSHPublicAccess enumerates the values for ssh public access.
type SSHPublicAccess string

const (
	// SSHPublicAccessDisabled ...
	SSHPublicAccessDisabled SSHPublicAccess = "Disabled"
	// SSHPublicAccessEnabled ...
	SSHPublicAccessEnabled SSHPublicAccess = "Enabled"
)

// PossibleSSHPublicAccessValues returns an array of possible values for the SSHPublicAccess const type.
func PossibleSSHPublicAccessValues() []SSHPublicAccess {
	return []SSHPublicAccess{SSHPublicAccessDisabled, SSHPublicAccessEnabled}
}

// Status enumerates the values for status.
type Status string

const (
	// Failure ...
	Failure Status = "Failure"
	// InvalidQuotaBelowClusterMinimum ...
	InvalidQuotaBelowClusterMinimum Status = "InvalidQuotaBelowClusterMinimum"
	// InvalidQuotaExceedsSubscriptionLimit ...
	InvalidQuotaExceedsSubscriptionLimit Status = "InvalidQuotaExceedsSubscriptionLimit"
	// InvalidVMFamilyName ...
	InvalidVMFamilyName Status = "InvalidVMFamilyName"
	// OperationNotEnabledForRegion ...
	OperationNotEnabledForRegion Status = "OperationNotEnabledForRegion"
	// OperationNotSupportedForSku ...
	OperationNotSupportedForSku Status = "OperationNotSupportedForSku"
	// Success ...
	Success Status = "Success"
	// Undefined ...
	Undefined Status = "Undefined"
)

// PossibleStatusValues returns an array of possible values for the Status const type.
func PossibleStatusValues() []Status {
	return []Status{Failure, InvalidQuotaBelowClusterMinimum, InvalidQuotaExceedsSubscriptionLimit, InvalidVMFamilyName, OperationNotEnabledForRegion, OperationNotSupportedForSku, Success, Undefined}
}

// Status1 enumerates the values for status 1.
type Status1 string

const (
	// Status1Auto ...
	Status1Auto Status1 = "Auto"
	// Status1Disabled ...
	Status1Disabled Status1 = "Disabled"
	// Status1Enabled ...
	Status1Enabled Status1 = "Enabled"
)

// PossibleStatus1Values returns an array of possible values for the Status1 const type.
func PossibleStatus1Values() []Status1 {
	return []Status1{Status1Auto, Status1Disabled, Status1Enabled}
}

// UnderlyingResourceAction enumerates the values for underlying resource action.
type UnderlyingResourceAction string

const (
	// UnderlyingResourceActionDelete ...
	UnderlyingResourceActionDelete UnderlyingResourceAction = "Delete"
	// UnderlyingResourceActionDetach ...
	UnderlyingResourceActionDetach UnderlyingResourceAction = "Detach"
)

// PossibleUnderlyingResourceActionValues returns an array of possible values for the UnderlyingResourceAction const type.
func PossibleUnderlyingResourceActionValues() []UnderlyingResourceAction {
	return []UnderlyingResourceAction{UnderlyingResourceActionDelete, UnderlyingResourceActionDetach}
}

// UsageUnit enumerates the values for usage unit.
type UsageUnit string

const (
	// UsageUnitCount ...
	UsageUnitCount UsageUnit = "Count"
)

// PossibleUsageUnitValues returns an array of possible values for the UsageUnit const type.
func PossibleUsageUnitValues() []UsageUnit {
	return []UsageUnit{UsageUnitCount}
}

// ValueFormat enumerates the values for value format.
type ValueFormat string

const (
	// JSON ...
	JSON ValueFormat = "JSON"
)

// PossibleValueFormatValues returns an array of possible values for the ValueFormat const type.
func PossibleValueFormatValues() []ValueFormat {
	return []ValueFormat{JSON}
}

// VariantType enumerates the values for variant type.
type VariantType string

const (
	// Control ...
	Control VariantType = "Control"
	// Treatment ...
	Treatment VariantType = "Treatment"
)

// PossibleVariantTypeValues returns an array of possible values for the VariantType const type.
func PossibleVariantTypeValues() []VariantType {
	return []VariantType{Control, Treatment}
}

// VMPriceOSType enumerates the values for vm price os type.
type VMPriceOSType string

const (
	// VMPriceOSTypeLinux ...
	VMPriceOSTypeLinux VMPriceOSType = "Linux"
	// VMPriceOSTypeWindows ...
	VMPriceOSTypeWindows VMPriceOSType = "Windows"
)

// PossibleVMPriceOSTypeValues returns an array of possible values for the VMPriceOSType const type.
func PossibleVMPriceOSTypeValues() []VMPriceOSType {
	return []VMPriceOSType{VMPriceOSTypeLinux, VMPriceOSTypeWindows}
}

// VMPriority enumerates the values for vm priority.
type VMPriority string

const (
	// Dedicated ...
	Dedicated VMPriority = "Dedicated"
	// LowPriority ...
	LowPriority VMPriority = "LowPriority"
)

// PossibleVMPriorityValues returns an array of possible values for the VMPriority const type.
func PossibleVMPriorityValues() []VMPriority {
	return []VMPriority{Dedicated, LowPriority}
}

// VMTier enumerates the values for vm tier.
type VMTier string

const (
	// VMTierLowPriority ...
	VMTierLowPriority VMTier = "LowPriority"
	// VMTierSpot ...
	VMTierSpot VMTier = "Spot"
	// VMTierStandard ...
	VMTierStandard VMTier = "Standard"
)

// PossibleVMTierValues returns an array of possible values for the VMTier const type.
func PossibleVMTierValues() []VMTier {
	return []VMTier{VMTierLowPriority, VMTierSpot, VMTierStandard}
}

// WebServiceState enumerates the values for web service state.
type WebServiceState string

const (
	// Failed ...
	Failed WebServiceState = "Failed"
	// Healthy ...
	Healthy WebServiceState = "Healthy"
	// Transitioning ...
	Transitioning WebServiceState = "Transitioning"
	// Unhealthy ...
	Unhealthy WebServiceState = "Unhealthy"
	// Unschedulable ...
	Unschedulable WebServiceState = "Unschedulable"
)

// PossibleWebServiceStateValues returns an array of possible values for the WebServiceState const type.
func PossibleWebServiceStateValues() []WebServiceState {
	return []WebServiceState{Failed, Healthy, Transitioning, Unhealthy, Unschedulable}
}
