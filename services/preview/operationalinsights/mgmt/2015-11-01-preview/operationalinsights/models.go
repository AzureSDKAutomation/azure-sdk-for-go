package operationalinsights

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
	"github.com/Azure/go-autorest/autorest/date"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// The package's fully qualified name.
const fqdn = "github.com/Azure/azure-sdk-for-go/services/preview/operationalinsights/mgmt/2015-11-01-preview/operationalinsights"

// DataSource datasources under OMS Workspace.
type DataSource struct {
	autorest.Response `json:"-"`
	// Properties - The data source properties in raw json format, each kind of data source have it's own schema.
	Properties interface{} `json:"properties,omitempty"`
	// ETag - The ETag of the data source.
	ETag *string `json:"eTag,omitempty"`
	// Kind - Possible values include: 'AzureActivityLog', 'ChangeTrackingPath', 'ChangeTrackingDefaultPath', 'ChangeTrackingDefaultRegistry', 'ChangeTrackingCustomRegistry', 'CustomLog', 'CustomLogCollection', 'GenericDataSource', 'IISLogs', 'LinuxPerformanceObject', 'LinuxPerformanceCollection', 'LinuxSyslog', 'LinuxSyslogCollection', 'WindowsEvent', 'WindowsPerformanceCounter'
	Kind DataSourceKind `json:"kind,omitempty"`
	// ID - READ-ONLY; Resource ID.
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; Resource name.
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; Resource type.
	Type *string `json:"type,omitempty"`
	// Tags - Resource tags
	Tags map[string]*string `json:"tags"`
}

// MarshalJSON is the custom marshaler for DataSource.
func (ds DataSource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if ds.Properties != nil {
		objectMap["properties"] = ds.Properties
	}
	if ds.ETag != nil {
		objectMap["eTag"] = ds.ETag
	}
	if ds.Kind != "" {
		objectMap["kind"] = ds.Kind
	}
	if ds.Tags != nil {
		objectMap["tags"] = ds.Tags
	}
	return json.Marshal(objectMap)
}

// DataSourceFilter dataSource filter. Right now, only filter by kind is supported.
type DataSourceFilter struct {
	// Kind - Possible values include: 'AzureActivityLog', 'ChangeTrackingPath', 'ChangeTrackingDefaultPath', 'ChangeTrackingDefaultRegistry', 'ChangeTrackingCustomRegistry', 'CustomLog', 'CustomLogCollection', 'GenericDataSource', 'IISLogs', 'LinuxPerformanceObject', 'LinuxPerformanceCollection', 'LinuxSyslog', 'LinuxSyslogCollection', 'WindowsEvent', 'WindowsPerformanceCounter'
	Kind DataSourceKind `json:"kind,omitempty"`
}

// DataSourceListResult the list data source by workspace operation response.
type DataSourceListResult struct {
	autorest.Response `json:"-"`
	// Value - A list of datasources.
	Value *[]DataSource `json:"value,omitempty"`
	// NextLink - The link (url) to the next page of datasources.
	NextLink *string `json:"nextLink,omitempty"`
}

// DataSourceListResultIterator provides access to a complete listing of DataSource values.
type DataSourceListResultIterator struct {
	i    int
	page DataSourceListResultPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *DataSourceListResultIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DataSourceListResultIterator.NextWithContext")
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
func (iter *DataSourceListResultIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter DataSourceListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter DataSourceListResultIterator) Response() DataSourceListResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter DataSourceListResultIterator) Value() DataSource {
	if !iter.page.NotDone() {
		return DataSource{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the DataSourceListResultIterator type.
func NewDataSourceListResultIterator(page DataSourceListResultPage) DataSourceListResultIterator {
	return DataSourceListResultIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (dslr DataSourceListResult) IsEmpty() bool {
	return dslr.Value == nil || len(*dslr.Value) == 0
}

// hasNextLink returns true if the NextLink is not empty.
func (dslr DataSourceListResult) hasNextLink() bool {
	return dslr.NextLink != nil && len(*dslr.NextLink) != 0
}

// dataSourceListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (dslr DataSourceListResult) dataSourceListResultPreparer(ctx context.Context) (*http.Request, error) {
	if !dslr.hasNextLink() {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(dslr.NextLink)))
}

// DataSourceListResultPage contains a page of DataSource values.
type DataSourceListResultPage struct {
	fn   func(context.Context, DataSourceListResult) (DataSourceListResult, error)
	dslr DataSourceListResult
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *DataSourceListResultPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DataSourceListResultPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	for {
		next, err := page.fn(ctx, page.dslr)
		if err != nil {
			return err
		}
		page.dslr = next
		if !next.hasNextLink() || !next.IsEmpty() {
			break
		}
	}
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *DataSourceListResultPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page DataSourceListResultPage) NotDone() bool {
	return !page.dslr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page DataSourceListResultPage) Response() DataSourceListResult {
	return page.dslr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page DataSourceListResultPage) Values() []DataSource {
	if page.dslr.IsEmpty() {
		return nil
	}
	return *page.dslr.Value
}

// Creates a new instance of the DataSourceListResultPage type.
func NewDataSourceListResultPage(cur DataSourceListResult, getNextPage func(context.Context, DataSourceListResult) (DataSourceListResult, error)) DataSourceListResultPage {
	return DataSourceListResultPage{
		fn:   getNextPage,
		dslr: cur,
	}
}

// ErrorResponse describes the format of Error response.
type ErrorResponse struct {
	// Code - Error code
	Code *string `json:"code,omitempty"`
	// Message - Error message indicating why the operation failed.
	Message *string `json:"message,omitempty"`
}

// IntelligencePack intelligence Pack containing a string name and boolean indicating if it's enabled.
type IntelligencePack struct {
	// Name - The name of the intelligence pack.
	Name *string `json:"name,omitempty"`
	// Enabled - The enabled boolean for the intelligence pack.
	Enabled *bool `json:"enabled,omitempty"`
	// DisplayName - The display name of the intelligence pack.
	DisplayName *string `json:"displayName,omitempty"`
}

// LinkedService the top level Linked service resource container.
type LinkedService struct {
	autorest.Response `json:"-"`
	// LinkedServiceProperties - The properties of the linked service.
	*LinkedServiceProperties `json:"properties,omitempty"`
	// ID - READ-ONLY; Resource ID.
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; Resource name.
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; Resource type.
	Type *string `json:"type,omitempty"`
	// Tags - Resource tags
	Tags map[string]*string `json:"tags"`
}

// MarshalJSON is the custom marshaler for LinkedService.
func (ls LinkedService) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if ls.LinkedServiceProperties != nil {
		objectMap["properties"] = ls.LinkedServiceProperties
	}
	if ls.Tags != nil {
		objectMap["tags"] = ls.Tags
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for LinkedService struct.
func (ls *LinkedService) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "properties":
			if v != nil {
				var linkedServiceProperties LinkedServiceProperties
				err = json.Unmarshal(*v, &linkedServiceProperties)
				if err != nil {
					return err
				}
				ls.LinkedServiceProperties = &linkedServiceProperties
			}
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				ls.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				ls.Name = &name
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				ls.Type = &typeVar
			}
		case "tags":
			if v != nil {
				var tags map[string]*string
				err = json.Unmarshal(*v, &tags)
				if err != nil {
					return err
				}
				ls.Tags = tags
			}
		}
	}

	return nil
}

// LinkedServiceListResult the list linked service operation response.
type LinkedServiceListResult struct {
	autorest.Response `json:"-"`
	// Value - Gets or sets a list of linked service instances.
	Value *[]LinkedService `json:"value,omitempty"`
}

// LinkedServiceProperties linked service properties.
type LinkedServiceProperties struct {
	// ResourceID - The resource id of the resource that will be linked to the workspace.
	ResourceID *string `json:"resourceId,omitempty"`
}

// ListIntelligencePack ...
type ListIntelligencePack struct {
	autorest.Response `json:"-"`
	Value             *[]IntelligencePack `json:"value,omitempty"`
}

// ManagementGroup a management group that is connected to a workspace
type ManagementGroup struct {
	// ManagementGroupProperties - The properties of the management group.
	*ManagementGroupProperties `json:"properties,omitempty"`
}

// MarshalJSON is the custom marshaler for ManagementGroup.
func (mg ManagementGroup) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if mg.ManagementGroupProperties != nil {
		objectMap["properties"] = mg.ManagementGroupProperties
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for ManagementGroup struct.
func (mg *ManagementGroup) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "properties":
			if v != nil {
				var managementGroupProperties ManagementGroupProperties
				err = json.Unmarshal(*v, &managementGroupProperties)
				if err != nil {
					return err
				}
				mg.ManagementGroupProperties = &managementGroupProperties
			}
		}
	}

	return nil
}

// ManagementGroupProperties management group properties.
type ManagementGroupProperties struct {
	// ServerCount - The number of servers connected to the management group.
	ServerCount *int32 `json:"serverCount,omitempty"`
	// IsGateway - Gets or sets a value indicating whether the management group is a gateway.
	IsGateway *bool `json:"isGateway,omitempty"`
	// Name - The name of the management group.
	Name *string `json:"name,omitempty"`
	// ID - The unique ID of the management group.
	ID *string `json:"id,omitempty"`
	// Created - The datetime that the management group was created.
	Created *date.Time `json:"created,omitempty"`
	// DataReceived - The last datetime that the management group received data.
	DataReceived *date.Time `json:"dataReceived,omitempty"`
	// Version - The version of System Center that is managing the management group.
	Version *string `json:"version,omitempty"`
	// Sku - The SKU of System Center that is managing the management group.
	Sku *string `json:"sku,omitempty"`
}

// MetricName the name of a metric.
type MetricName struct {
	// Value - The system name of the metric.
	Value *string `json:"value,omitempty"`
	// LocalizedValue - The localized name of the metric.
	LocalizedValue *string `json:"localizedValue,omitempty"`
}

// Operation supported operation of OperationalInsights resource provider.
type Operation struct {
	// Name - Operation name: {provider}/{resource}/{operation}
	Name *string `json:"name,omitempty"`
	// Display - Display metadata associated with the operation.
	Display *OperationDisplay `json:"display,omitempty"`
}

// OperationDisplay display metadata associated with the operation.
type OperationDisplay struct {
	// Provider - Service provider: Microsoft OperationsManagement.
	Provider *string `json:"provider,omitempty"`
	// Resource - Resource on which the operation is performed etc.
	Resource *string `json:"resource,omitempty"`
	// Operation - Type of operation: get, read, delete, etc.
	Operation *string `json:"operation,omitempty"`
}

// OperationListResult result of the request to list solution operations.
type OperationListResult struct {
	autorest.Response `json:"-"`
	// Value - List of solution operations supported by the OperationsManagement resource provider.
	Value *[]Operation `json:"value,omitempty"`
	// NextLink - READ-ONLY; URL to get the next set of operation list results if there are any.
	NextLink *string `json:"nextLink,omitempty"`
}

// MarshalJSON is the custom marshaler for OperationListResult.
func (olr OperationListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if olr.Value != nil {
		objectMap["value"] = olr.Value
	}
	return json.Marshal(objectMap)
}

// OperationListResultIterator provides access to a complete listing of Operation values.
type OperationListResultIterator struct {
	i    int
	page OperationListResultPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *OperationListResultIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/OperationListResultIterator.NextWithContext")
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
func (iter *OperationListResultIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter OperationListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter OperationListResultIterator) Response() OperationListResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter OperationListResultIterator) Value() Operation {
	if !iter.page.NotDone() {
		return Operation{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the OperationListResultIterator type.
func NewOperationListResultIterator(page OperationListResultPage) OperationListResultIterator {
	return OperationListResultIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (olr OperationListResult) IsEmpty() bool {
	return olr.Value == nil || len(*olr.Value) == 0
}

// hasNextLink returns true if the NextLink is not empty.
func (olr OperationListResult) hasNextLink() bool {
	return olr.NextLink != nil && len(*olr.NextLink) != 0
}

// operationListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (olr OperationListResult) operationListResultPreparer(ctx context.Context) (*http.Request, error) {
	if !olr.hasNextLink() {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(olr.NextLink)))
}

// OperationListResultPage contains a page of Operation values.
type OperationListResultPage struct {
	fn  func(context.Context, OperationListResult) (OperationListResult, error)
	olr OperationListResult
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *OperationListResultPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/OperationListResultPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	for {
		next, err := page.fn(ctx, page.olr)
		if err != nil {
			return err
		}
		page.olr = next
		if !next.hasNextLink() || !next.IsEmpty() {
			break
		}
	}
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *OperationListResultPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page OperationListResultPage) NotDone() bool {
	return !page.olr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page OperationListResultPage) Response() OperationListResult {
	return page.olr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page OperationListResultPage) Values() []Operation {
	if page.olr.IsEmpty() {
		return nil
	}
	return *page.olr.Value
}

// Creates a new instance of the OperationListResultPage type.
func NewOperationListResultPage(cur OperationListResult, getNextPage func(context.Context, OperationListResult) (OperationListResult, error)) OperationListResultPage {
	return OperationListResultPage{
		fn:  getNextPage,
		olr: cur,
	}
}

// OperationStatus the status of operation.
type OperationStatus struct {
	autorest.Response `json:"-"`
	// ID - The operation Id.
	ID *string `json:"id,omitempty"`
	// Name - The operation name.
	Name *string `json:"name,omitempty"`
	// StartTime - The start time of the operation.
	StartTime *string `json:"startTime,omitempty"`
	// EndTime - The end time of the operation.
	EndTime *string `json:"endTime,omitempty"`
	// Status - The status of the operation.
	Status *string `json:"status,omitempty"`
	// Error - The error detail of the operation if any.
	Error *ErrorResponse `json:"error,omitempty"`
}

// ProxyResource common properties of proxy resource.
type ProxyResource struct {
	// ID - READ-ONLY; Resource ID.
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; Resource name.
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; Resource type.
	Type *string `json:"type,omitempty"`
	// Tags - Resource tags
	Tags map[string]*string `json:"tags"`
}

// MarshalJSON is the custom marshaler for ProxyResource.
func (pr ProxyResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if pr.Tags != nil {
		objectMap["tags"] = pr.Tags
	}
	return json.Marshal(objectMap)
}

// Resource the resource definition.
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

// SharedKeys the shared keys for a workspace.
type SharedKeys struct {
	autorest.Response `json:"-"`
	// PrimarySharedKey - The primary shared key of a workspace.
	PrimarySharedKey *string `json:"primarySharedKey,omitempty"`
	// SecondarySharedKey - The secondary shared key of a workspace.
	SecondarySharedKey *string `json:"secondarySharedKey,omitempty"`
}

// Sku the SKU (tier) of a workspace.
type Sku struct {
	// Name - The name of the SKU. Possible values include: 'Free', 'Standard', 'Premium', 'PerNode', 'PerGB2018', 'Standalone', 'CapacityReservation'
	Name SkuNameEnum `json:"name,omitempty"`
}

// UsageMetric a metric describing the usage of a resource.
type UsageMetric struct {
	// Name - The name of the metric.
	Name *MetricName `json:"name,omitempty"`
	// Unit - The units used for the metric.
	Unit *string `json:"unit,omitempty"`
	// CurrentValue - The current value of the metric.
	CurrentValue *float64 `json:"currentValue,omitempty"`
	// Limit - The quota limit for the metric.
	Limit *float64 `json:"limit,omitempty"`
	// NextResetTime - The time that the metric's value will reset.
	NextResetTime *date.Time `json:"nextResetTime,omitempty"`
	// QuotaPeriod - The quota period that determines the length of time between value resets.
	QuotaPeriod *string `json:"quotaPeriod,omitempty"`
}

// Workspace the top level Workspace resource container.
type Workspace struct {
	autorest.Response `json:"-"`
	// WorkspaceProperties - Workspace properties.
	*WorkspaceProperties `json:"properties,omitempty"`
	// ETag - The ETag of the workspace.
	ETag *string `json:"eTag,omitempty"`
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

// MarshalJSON is the custom marshaler for Workspace.
func (w Workspace) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if w.WorkspaceProperties != nil {
		objectMap["properties"] = w.WorkspaceProperties
	}
	if w.ETag != nil {
		objectMap["eTag"] = w.ETag
	}
	if w.Location != nil {
		objectMap["location"] = w.Location
	}
	if w.Tags != nil {
		objectMap["tags"] = w.Tags
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for Workspace struct.
func (w *Workspace) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "properties":
			if v != nil {
				var workspaceProperties WorkspaceProperties
				err = json.Unmarshal(*v, &workspaceProperties)
				if err != nil {
					return err
				}
				w.WorkspaceProperties = &workspaceProperties
			}
		case "eTag":
			if v != nil {
				var eTag string
				err = json.Unmarshal(*v, &eTag)
				if err != nil {
					return err
				}
				w.ETag = &eTag
			}
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				w.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				w.Name = &name
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				w.Type = &typeVar
			}
		case "location":
			if v != nil {
				var location string
				err = json.Unmarshal(*v, &location)
				if err != nil {
					return err
				}
				w.Location = &location
			}
		case "tags":
			if v != nil {
				var tags map[string]*string
				err = json.Unmarshal(*v, &tags)
				if err != nil {
					return err
				}
				w.Tags = tags
			}
		}
	}

	return nil
}

// WorkspaceListManagementGroupsResult the list workspace management groups operation response.
type WorkspaceListManagementGroupsResult struct {
	autorest.Response `json:"-"`
	// Value - Gets or sets a list of management groups attached to the workspace.
	Value *[]ManagementGroup `json:"value,omitempty"`
}

// WorkspaceListResult the list workspaces operation response.
type WorkspaceListResult struct {
	autorest.Response `json:"-"`
	// Value - A list of workspaces.
	Value *[]Workspace `json:"value,omitempty"`
}

// WorkspaceListUsagesResult the list workspace usages operation response.
type WorkspaceListUsagesResult struct {
	autorest.Response `json:"-"`
	// Value - Gets or sets a list of usage metrics for a workspace.
	Value *[]UsageMetric `json:"value,omitempty"`
}

// WorkspaceProperties workspace properties.
type WorkspaceProperties struct {
	// ProvisioningState - The provisioning state of the workspace. Possible values include: 'Creating', 'Succeeded', 'Failed', 'Canceled', 'Deleting', 'ProvisioningAccount'
	ProvisioningState EntityStatus `json:"provisioningState,omitempty"`
	// Source - READ-ONLY; This is a read-only legacy property. It is always set to 'Azure' by the service. Kept here for backward compatibility.
	Source *string `json:"source,omitempty"`
	// CustomerID - READ-ONLY; This is a read-only property. Represents the ID associated with the workspace.
	CustomerID *string `json:"customerId,omitempty"`
	// PortalURL - READ-ONLY; This is a legacy property and is not used anymore. Kept here for backward compatibility.
	PortalURL *string `json:"portalUrl,omitempty"`
	// Sku - The SKU of the workspace.
	Sku *Sku `json:"sku,omitempty"`
	// RetentionInDays - The workspace data retention in days. -1 means Unlimited retention for the Unlimited Sku. 730 days is the maximum allowed for all other Skus.
	RetentionInDays *int32 `json:"retentionInDays,omitempty"`
}

// MarshalJSON is the custom marshaler for WorkspaceProperties.
func (wp WorkspaceProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if wp.ProvisioningState != "" {
		objectMap["provisioningState"] = wp.ProvisioningState
	}
	if wp.Sku != nil {
		objectMap["sku"] = wp.Sku
	}
	if wp.RetentionInDays != nil {
		objectMap["retentionInDays"] = wp.RetentionInDays
	}
	return json.Marshal(objectMap)
}

// WorkspacesCreateOrUpdateFuture an abstraction for monitoring and retrieving the results of a
// long-running operation.
type WorkspacesCreateOrUpdateFuture struct {
	azure.FutureAPI
	// Result returns the result of the asynchronous operation.
	// If the operation has not completed it will return an error.
	Result func(WorkspacesClient) (Workspace, error)
}

// UnmarshalJSON is the custom unmarshaller for CreateFuture.
func (future *WorkspacesCreateOrUpdateFuture) UnmarshalJSON(body []byte) error {
	var azFuture azure.Future
	if err := json.Unmarshal(body, &azFuture); err != nil {
		return err
	}
	future.FutureAPI = &azFuture
	future.Result = future.result
	return nil
}

// result is the default implementation for WorkspacesCreateOrUpdateFuture.Result.
func (future *WorkspacesCreateOrUpdateFuture) result(client WorkspacesClient) (w Workspace, err error) {
	var done bool
	done, err = future.DoneWithContext(context.Background(), client)
	if err != nil {
		err = autorest.NewErrorWithError(err, "operationalinsights.WorkspacesCreateOrUpdateFuture", "Result", future.Response(), "Polling failure")
		return
	}
	if !done {
		w.Response.Response = future.Response()
		err = azure.NewAsyncOpIncompleteError("operationalinsights.WorkspacesCreateOrUpdateFuture")
		return
	}
	sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if w.Response.Response, err = future.GetResult(sender); err == nil && w.Response.Response.StatusCode != http.StatusNoContent {
		w, err = client.CreateOrUpdateResponder(w.Response.Response)
		if err != nil {
			err = autorest.NewErrorWithError(err, "operationalinsights.WorkspacesCreateOrUpdateFuture", "Result", w.Response.Response, "Failure responding to request")
		}
	}
	return
}
