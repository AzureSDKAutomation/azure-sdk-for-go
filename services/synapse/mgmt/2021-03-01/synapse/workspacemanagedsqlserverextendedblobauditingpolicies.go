package synapse

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClient is the azure Synapse Analytics Management Client
type WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClient struct {
	BaseClient
}

// NewWorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClient creates an instance of the
// WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClient client.
func NewWorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClient(subscriptionID string) WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClient {
	return NewWorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewWorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClientWithBaseURI creates an instance of the
// WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClient client using a custom endpoint.  Use this when
// interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewWorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClientWithBaseURI(baseURI string, subscriptionID string) WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClient {
	return WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate create or Update a workspace managed sql server's extended blob auditing policy.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// workspaceName - the name of the workspace
// parameters - properties of extended blob auditing policy.
func (client WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, parameters ExtendedServerBlobAuditingPolicy) (result WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("synapse.WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, workspaceName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClient", "CreateOrUpdate", nil, "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, workspaceName string, parameters ExtendedServerBlobAuditingPolicy) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"blobAuditingPolicyName": autorest.Encode("path", "default"),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
		"workspaceName":          autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2021-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/extendedAuditingSettings/{blobAuditingPolicyName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClient) CreateOrUpdateSender(req *http.Request) (future WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesCreateOrUpdateFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClient) CreateOrUpdateResponder(resp *http.Response) (result ExtendedServerBlobAuditingPolicy, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get get a workspace SQL server's extended blob auditing policy.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// workspaceName - the name of the workspace
func (client WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClient) Get(ctx context.Context, resourceGroupName string, workspaceName string) (result ExtendedServerBlobAuditingPolicy, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("synapse.WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, workspaceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "synapse.WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClient) GetPreparer(ctx context.Context, resourceGroupName string, workspaceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"blobAuditingPolicyName": autorest.Encode("path", "default"),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
		"workspaceName":          autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2021-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/extendedAuditingSettings/{blobAuditingPolicyName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClient) GetResponder(resp *http.Response) (result ExtendedServerBlobAuditingPolicy, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByWorkspace list workspace managed sql server's extended blob auditing policies.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// workspaceName - the name of the workspace
func (client WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClient) ListByWorkspace(ctx context.Context, resourceGroupName string, workspaceName string) (result ExtendedServerBlobAuditingPolicyListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClient.ListByWorkspace")
		defer func() {
			sc := -1
			if result.esbaplr.Response.Response != nil {
				sc = result.esbaplr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("synapse.WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClient", "ListByWorkspace", err.Error())
	}

	result.fn = client.listByWorkspaceNextResults
	req, err := client.ListByWorkspacePreparer(ctx, resourceGroupName, workspaceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClient", "ListByWorkspace", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByWorkspaceSender(req)
	if err != nil {
		result.esbaplr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "synapse.WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClient", "ListByWorkspace", resp, "Failure sending request")
		return
	}

	result.esbaplr, err = client.ListByWorkspaceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClient", "ListByWorkspace", resp, "Failure responding to request")
		return
	}
	if result.esbaplr.hasNextLink() && result.esbaplr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByWorkspacePreparer prepares the ListByWorkspace request.
func (client WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClient) ListByWorkspacePreparer(ctx context.Context, resourceGroupName string, workspaceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workspaceName":     autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2021-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/extendedAuditingSettings", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByWorkspaceSender sends the ListByWorkspace request. The method will close the
// http.Response Body if it receives an error.
func (client WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClient) ListByWorkspaceSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByWorkspaceResponder handles the response to the ListByWorkspace request. The method always
// closes the http.Response Body.
func (client WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClient) ListByWorkspaceResponder(resp *http.Response) (result ExtendedServerBlobAuditingPolicyListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByWorkspaceNextResults retrieves the next set of results, if any.
func (client WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClient) listByWorkspaceNextResults(ctx context.Context, lastResults ExtendedServerBlobAuditingPolicyListResult) (result ExtendedServerBlobAuditingPolicyListResult, err error) {
	req, err := lastResults.extendedServerBlobAuditingPolicyListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "synapse.WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClient", "listByWorkspaceNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByWorkspaceSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "synapse.WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClient", "listByWorkspaceNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByWorkspaceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClient", "listByWorkspaceNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByWorkspaceComplete enumerates all values, automatically crossing page boundaries as required.
func (client WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClient) ListByWorkspaceComplete(ctx context.Context, resourceGroupName string, workspaceName string) (result ExtendedServerBlobAuditingPolicyListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClient.ListByWorkspace")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByWorkspace(ctx, resourceGroupName, workspaceName)
	return
}
