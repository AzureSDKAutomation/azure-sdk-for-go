package containerservice

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"encoding/json"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// The package's fully qualified name.
const fqdn = "github.com/Azure/azure-sdk-for-go/services/containerservice/mgmt/2017-01-31/containerservice"

// AgentPoolProfile profile for the container service agent pool.
type AgentPoolProfile struct {
	// Name - Unique name of the agent pool profile in the context of the subscription and resource group.
	Name *string `json:"name,omitempty"`
	// Count - Number of agents (VMs) to host docker containers. Allowed values must be in the range of 1 to 100 (inclusive). The default value is 1.
	Count *int32 `json:"count,omitempty"`
	// VMSize - Size of agent VMs. Possible values include: 'StandardA0', 'StandardA1', 'StandardA2', 'StandardA3', 'StandardA4', 'StandardA5', 'StandardA6', 'StandardA7', 'StandardA8', 'StandardA9', 'StandardA10', 'StandardA11', 'StandardD1', 'StandardD2', 'StandardD3', 'StandardD4', 'StandardD11', 'StandardD12', 'StandardD13', 'StandardD14', 'StandardD1V2', 'StandardD2V2', 'StandardD3V2', 'StandardD4V2', 'StandardD5V2', 'StandardD11V2', 'StandardD12V2', 'StandardD13V2', 'StandardD14V2', 'StandardG1', 'StandardG2', 'StandardG3', 'StandardG4', 'StandardG5', 'StandardDS1', 'StandardDS2', 'StandardDS3', 'StandardDS4', 'StandardDS11', 'StandardDS12', 'StandardDS13', 'StandardDS14', 'StandardGS1', 'StandardGS2', 'StandardGS3', 'StandardGS4', 'StandardGS5'
	VMSize VMSizeTypes `json:"vmSize,omitempty"`
	// DNSPrefix - DNS prefix to be used to create the FQDN for the agent pool.
	DNSPrefix *string `json:"dnsPrefix,omitempty"`
	// Fqdn - READ-ONLY; FQDN for the agent pool.
	Fqdn *string `json:"fqdn,omitempty"`
}

// MarshalJSON is the custom marshaler for AgentPoolProfile.
func (app AgentPoolProfile) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if app.Name != nil {
		objectMap["name"] = app.Name
	}
	if app.Count != nil {
		objectMap["count"] = app.Count
	}
	if app.VMSize != "" {
		objectMap["vmSize"] = app.VMSize
	}
	if app.DNSPrefix != nil {
		objectMap["dnsPrefix"] = app.DNSPrefix
	}
	return json.Marshal(objectMap)
}

// ContainerService container service.
type ContainerService struct {
	autorest.Response `json:"-"`
	*Properties       `json:"properties,omitempty"`
	// ID - READ-ONLY; Resource Id
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; Resource name
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; Resource type
	Type *string `json:"type,omitempty"`
	// Location - Resource location
	Location *string `json:"location,omitempty"`
	// Tags - Resource tags
	Tags map[string]*string `json:"tags"`
}

// MarshalJSON is the custom marshaler for ContainerService.
func (cs ContainerService) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if cs.Properties != nil {
		objectMap["properties"] = cs.Properties
	}
	if cs.Location != nil {
		objectMap["location"] = cs.Location
	}
	if cs.Tags != nil {
		objectMap["tags"] = cs.Tags
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for ContainerService struct.
func (cs *ContainerService) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "properties":
			if v != nil {
				var properties Properties
				err = json.Unmarshal(*v, &properties)
				if err != nil {
					return err
				}
				cs.Properties = &properties
			}
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				cs.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				cs.Name = &name
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				cs.Type = &typeVar
			}
		case "location":
			if v != nil {
				var location string
				err = json.Unmarshal(*v, &location)
				if err != nil {
					return err
				}
				cs.Location = &location
			}
		case "tags":
			if v != nil {
				var tags map[string]*string
				err = json.Unmarshal(*v, &tags)
				if err != nil {
					return err
				}
				cs.Tags = tags
			}
		}
	}

	return nil
}

// ContainerServicesCreateOrUpdateFutureType an abstraction for monitoring and retrieving the results of a
// long-running operation.
type ContainerServicesCreateOrUpdateFutureType struct {
	azure.FutureAPI
	// Result returns the result of the asynchronous operation.
	// If the operation has not completed it will return an error.
	Result func(ContainerServicesClient) (ContainerService, error)
}

// UnmarshalJSON is the custom unmarshaller for CreateFuture.
func (future *ContainerServicesCreateOrUpdateFutureType) UnmarshalJSON(body []byte) error {
	var azFuture azure.Future
	if err := json.Unmarshal(body, &azFuture); err != nil {
		return err
	}
	future.FutureAPI = &azFuture
	future.Result = future.result
	return nil
}

// result is the default implementation for ContainerServicesCreateOrUpdateFutureType.Result.
func (future *ContainerServicesCreateOrUpdateFutureType) result(client ContainerServicesClient) (cs ContainerService, err error) {
	var done bool
	done, err = future.DoneWithContext(context.Background(), client)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerservice.ContainerServicesCreateOrUpdateFutureType", "Result", future.Response(), "Polling failure")
		return
	}
	if !done {
		cs.Response.Response = future.Response()
		err = azure.NewAsyncOpIncompleteError("containerservice.ContainerServicesCreateOrUpdateFutureType")
		return
	}
	sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if cs.Response.Response, err = future.GetResult(sender); err == nil && cs.Response.Response.StatusCode != http.StatusNoContent {
		cs, err = client.CreateOrUpdateResponder(cs.Response.Response)
		if err != nil {
			err = autorest.NewErrorWithError(err, "containerservice.ContainerServicesCreateOrUpdateFutureType", "Result", cs.Response.Response, "Failure responding to request")
		}
	}
	return
}

// ContainerServicesDeleteFutureType an abstraction for monitoring and retrieving the results of a
// long-running operation.
type ContainerServicesDeleteFutureType struct {
	azure.FutureAPI
	// Result returns the result of the asynchronous operation.
	// If the operation has not completed it will return an error.
	Result func(ContainerServicesClient) (autorest.Response, error)
}

// UnmarshalJSON is the custom unmarshaller for CreateFuture.
func (future *ContainerServicesDeleteFutureType) UnmarshalJSON(body []byte) error {
	var azFuture azure.Future
	if err := json.Unmarshal(body, &azFuture); err != nil {
		return err
	}
	future.FutureAPI = &azFuture
	future.Result = future.result
	return nil
}

// result is the default implementation for ContainerServicesDeleteFutureType.Result.
func (future *ContainerServicesDeleteFutureType) result(client ContainerServicesClient) (ar autorest.Response, err error) {
	var done bool
	done, err = future.DoneWithContext(context.Background(), client)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerservice.ContainerServicesDeleteFutureType", "Result", future.Response(), "Polling failure")
		return
	}
	if !done {
		ar.Response = future.Response()
		err = azure.NewAsyncOpIncompleteError("containerservice.ContainerServicesDeleteFutureType")
		return
	}
	ar.Response = future.Response()
	return
}

// CustomProfile properties to configure a custom container service cluster.
type CustomProfile struct {
	// Orchestrator - The name of the custom orchestrator to use.
	Orchestrator *string `json:"orchestrator,omitempty"`
}

// DiagnosticsProfile ...
type DiagnosticsProfile struct {
	// VMDiagnostics - Profile for the container service VM diagnostic agent.
	VMDiagnostics *VMDiagnostics `json:"vmDiagnostics,omitempty"`
}

// LinuxProfile profile for Linux VMs in the container service cluster.
type LinuxProfile struct {
	// AdminUsername - The administrator username to use for Linux VMs.
	AdminUsername *string `json:"adminUsername,omitempty"`
	// SSH - The ssh key configuration for Linux VMs.
	SSH *SSHConfiguration `json:"ssh,omitempty"`
}

// ListResult the response from the List Container Services operation.
type ListResult struct {
	autorest.Response `json:"-"`
	// Value - the list of container services.
	Value *[]ContainerService `json:"value,omitempty"`
	// NextLink - The URL to get the next set of container service results.
	NextLink *string `json:"nextLink,omitempty"`
}

// ListResultIterator provides access to a complete listing of ContainerService values.
type ListResultIterator struct {
	i    int
	page ListResultPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *ListResultIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ListResultIterator.NextWithContext")
		defer func() {
			sc := -1
			if iter.Response().Response.Response != nil {
				sc = iter.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err = iter.page.NextWithContext(ctx)
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (iter *ListResultIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter ListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter ListResultIterator) Response() ListResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter ListResultIterator) Value() ContainerService {
	if !iter.page.NotDone() {
		return ContainerService{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the ListResultIterator type.
func NewListResultIterator(page ListResultPage) ListResultIterator {
	return ListResultIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (lr ListResult) IsEmpty() bool {
	return lr.Value == nil || len(*lr.Value) == 0
}

// hasNextLink returns true if the NextLink is not empty.
func (lr ListResult) hasNextLink() bool {
	return lr.NextLink != nil && len(*lr.NextLink) != 0
}

// listResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (lr ListResult) listResultPreparer(ctx context.Context) (*http.Request, error) {
	if !lr.hasNextLink() {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(lr.NextLink)))
}

// ListResultPage contains a page of ContainerService values.
type ListResultPage struct {
	fn func(context.Context, ListResult) (ListResult, error)
	lr ListResult
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *ListResultPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ListResultPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	for {
		next, err := page.fn(ctx, page.lr)
		if err != nil {
			return err
		}
		page.lr = next
		if !next.hasNextLink() || !next.IsEmpty() {
			break
		}
	}
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *ListResultPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page ListResultPage) NotDone() bool {
	return !page.lr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page ListResultPage) Response() ListResult {
	return page.lr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page ListResultPage) Values() []ContainerService {
	if page.lr.IsEmpty() {
		return nil
	}
	return *page.lr.Value
}

// Creates a new instance of the ListResultPage type.
func NewListResultPage(cur ListResult, getNextPage func(context.Context, ListResult) (ListResult, error)) ListResultPage {
	return ListResultPage{
		fn: getNextPage,
		lr: cur,
	}
}

// MasterProfile profile for the container service master.
type MasterProfile struct {
	// Count - Number of masters (VMs) in the container service cluster. Allowed values are 1, 3, and 5. The default value is 1.
	Count *int32 `json:"count,omitempty"`
	// DNSPrefix - DNS prefix to be used to create the FQDN for master.
	DNSPrefix *string `json:"dnsPrefix,omitempty"`
	// Fqdn - READ-ONLY; FQDN for the master.
	Fqdn *string `json:"fqdn,omitempty"`
}

// MarshalJSON is the custom marshaler for MasterProfile.
func (mp MasterProfile) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if mp.Count != nil {
		objectMap["count"] = mp.Count
	}
	if mp.DNSPrefix != nil {
		objectMap["dnsPrefix"] = mp.DNSPrefix
	}
	return json.Marshal(objectMap)
}

// OrchestratorProfile profile for the container service orchestrator.
type OrchestratorProfile struct {
	// OrchestratorType - The orchestrator to use to manage container service cluster resources. Valid values are Swarm, DCOS, and Custom. Possible values include: 'Swarm', 'DCOS', 'Custom', 'Kubernetes'
	OrchestratorType OrchestratorTypes `json:"orchestratorType,omitempty"`
}

// Properties properties of the container service.
type Properties struct {
	// ProvisioningState - READ-ONLY; the current deployment or provisioning state, which only appears in the response.
	ProvisioningState *string `json:"provisioningState,omitempty"`
	// OrchestratorProfile - Properties of the orchestrator.
	OrchestratorProfile *OrchestratorProfile `json:"orchestratorProfile,omitempty"`
	// CustomProfile - Properties for custom clusters.
	CustomProfile *CustomProfile `json:"customProfile,omitempty"`
	// ServicePrincipalProfile - Properties for cluster service principals.
	ServicePrincipalProfile *ServicePrincipalProfile `json:"servicePrincipalProfile,omitempty"`
	// MasterProfile - Properties of master agents.
	MasterProfile *MasterProfile `json:"masterProfile,omitempty"`
	// AgentPoolProfiles - Properties of the agent pool.
	AgentPoolProfiles *[]AgentPoolProfile `json:"agentPoolProfiles,omitempty"`
	// WindowsProfile - Properties of Windows VMs.
	WindowsProfile *WindowsProfile `json:"windowsProfile,omitempty"`
	// LinuxProfile - Properties of Linux VMs.
	LinuxProfile *LinuxProfile `json:"linuxProfile,omitempty"`
	// DiagnosticsProfile - Properties of the diagnostic agent.
	DiagnosticsProfile *DiagnosticsProfile `json:"diagnosticsProfile,omitempty"`
}

// MarshalJSON is the custom marshaler for Properties.
func (p Properties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if p.OrchestratorProfile != nil {
		objectMap["orchestratorProfile"] = p.OrchestratorProfile
	}
	if p.CustomProfile != nil {
		objectMap["customProfile"] = p.CustomProfile
	}
	if p.ServicePrincipalProfile != nil {
		objectMap["servicePrincipalProfile"] = p.ServicePrincipalProfile
	}
	if p.MasterProfile != nil {
		objectMap["masterProfile"] = p.MasterProfile
	}
	if p.AgentPoolProfiles != nil {
		objectMap["agentPoolProfiles"] = p.AgentPoolProfiles
	}
	if p.WindowsProfile != nil {
		objectMap["windowsProfile"] = p.WindowsProfile
	}
	if p.LinuxProfile != nil {
		objectMap["linuxProfile"] = p.LinuxProfile
	}
	if p.DiagnosticsProfile != nil {
		objectMap["diagnosticsProfile"] = p.DiagnosticsProfile
	}
	return json.Marshal(objectMap)
}

// Resource the Resource model definition.
type Resource struct {
	// ID - READ-ONLY; Resource Id
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; Resource name
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; Resource type
	Type *string `json:"type,omitempty"`
	// Location - Resource location
	Location *string `json:"location,omitempty"`
	// Tags - Resource tags
	Tags map[string]*string `json:"tags"`
}

// MarshalJSON is the custom marshaler for Resource.
func (r Resource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if r.Location != nil {
		objectMap["location"] = r.Location
	}
	if r.Tags != nil {
		objectMap["tags"] = r.Tags
	}
	return json.Marshal(objectMap)
}

// ResourceGroupsCheckExistenceResponse ...
type ResourceGroupsCheckExistenceResponse struct {
	Body *bool `json:"body,omitempty"`
}

// ServicePrincipalProfile information about a service principal identity for the cluster to use for
// manipulating Azure APIs.
type ServicePrincipalProfile struct {
	// ClientID - The ID for the service principal.
	ClientID *string `json:"clientId,omitempty"`
	// Secret - The secret password associated with the service principal.
	Secret *string `json:"secret,omitempty"`
}

// SSHConfiguration SSH configuration for Linux-based VMs running on Azure.
type SSHConfiguration struct {
	// PublicKeys - the list of SSH public keys used to authenticate with Linux-based VMs.
	PublicKeys *[]SSHPublicKey `json:"publicKeys,omitempty"`
}

// SSHPublicKey contains information about SSH certificate public key data.
type SSHPublicKey struct {
	// KeyData - Certificate public key used to authenticate with VMs through SSH. The certificate must be in PEM format with or without headers.
	KeyData *string `json:"keyData,omitempty"`
}

// VMDiagnostics profile for diagnostics on the container service VMs.
type VMDiagnostics struct {
	// Enabled - Whether the VM diagnostic agent is provisioned on the VM.
	Enabled *bool `json:"enabled,omitempty"`
	// StorageURI - READ-ONLY; The URI of the storage account where diagnostics are stored.
	StorageURI *string `json:"storageUri,omitempty"`
}

// MarshalJSON is the custom marshaler for VMDiagnostics.
func (vd VMDiagnostics) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if vd.Enabled != nil {
		objectMap["enabled"] = vd.Enabled
	}
	return json.Marshal(objectMap)
}

// WindowsProfile profile for Windows VMs in the container service cluster.
type WindowsProfile struct {
	// AdminUsername - The administrator username to use for Windows VMs.
	AdminUsername *string `json:"adminUsername,omitempty"`
	// AdminPassword - The administrator password to use for Windows VMs.
	AdminPassword *string `json:"adminPassword,omitempty"`
}
