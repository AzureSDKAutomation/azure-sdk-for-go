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

// SQLPoolBlobAuditingPoliciesClient is the azure Synapse Analytics Management Client
type SQLPoolBlobAuditingPoliciesClient struct {
	BaseClient
}

// NewSQLPoolBlobAuditingPoliciesClient creates an instance of the SQLPoolBlobAuditingPoliciesClient client.
func NewSQLPoolBlobAuditingPoliciesClient(subscriptionID string) SQLPoolBlobAuditingPoliciesClient {
	return NewSQLPoolBlobAuditingPoliciesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewSQLPoolBlobAuditingPoliciesClientWithBaseURI creates an instance of the SQLPoolBlobAuditingPoliciesClient client
// using a custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign
// clouds, Azure stack).
func NewSQLPoolBlobAuditingPoliciesClientWithBaseURI(baseURI string, subscriptionID string) SQLPoolBlobAuditingPoliciesClient {
	return SQLPoolBlobAuditingPoliciesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or updates a SQL pool's blob auditing policy.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// workspaceName - the name of the workspace
// SQLPoolName - SQL pool name
// parameters - the database blob auditing policy.
func (client SQLPoolBlobAuditingPoliciesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, SQLPoolName string, parameters SQLPoolBlobAuditingPolicy) (result SQLPoolBlobAuditingPolicy, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SQLPoolBlobAuditingPoliciesClient.CreateOrUpdate")
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
		return result, validation.NewError("synapse.SQLPoolBlobAuditingPoliciesClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, workspaceName, SQLPoolName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.SQLPoolBlobAuditingPoliciesClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "synapse.SQLPoolBlobAuditingPoliciesClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.SQLPoolBlobAuditingPoliciesClient", "CreateOrUpdate", resp, "Failure responding to request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client SQLPoolBlobAuditingPoliciesClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, workspaceName string, SQLPoolName string, parameters SQLPoolBlobAuditingPolicy) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"blobAuditingPolicyName": autorest.Encode("path", "default"),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"sqlPoolName":            autorest.Encode("path", SQLPoolName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
		"workspaceName":          autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2020-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	parameters.Kind = nil
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/sqlPools/{sqlPoolName}/auditingSettings/{blobAuditingPolicyName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client SQLPoolBlobAuditingPoliciesClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client SQLPoolBlobAuditingPoliciesClient) CreateOrUpdateResponder(resp *http.Response) (result SQLPoolBlobAuditingPolicy, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get get a SQL pool's blob auditing policy.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// workspaceName - the name of the workspace
// SQLPoolName - SQL pool name
func (client SQLPoolBlobAuditingPoliciesClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, SQLPoolName string) (result SQLPoolBlobAuditingPolicy, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SQLPoolBlobAuditingPoliciesClient.Get")
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
		return result, validation.NewError("synapse.SQLPoolBlobAuditingPoliciesClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, workspaceName, SQLPoolName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.SQLPoolBlobAuditingPoliciesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "synapse.SQLPoolBlobAuditingPoliciesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.SQLPoolBlobAuditingPoliciesClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client SQLPoolBlobAuditingPoliciesClient) GetPreparer(ctx context.Context, resourceGroupName string, workspaceName string, SQLPoolName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"blobAuditingPolicyName": autorest.Encode("path", "default"),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"sqlPoolName":            autorest.Encode("path", SQLPoolName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
		"workspaceName":          autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2020-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/sqlPools/{sqlPoolName}/auditingSettings/{blobAuditingPolicyName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client SQLPoolBlobAuditingPoliciesClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client SQLPoolBlobAuditingPoliciesClient) GetResponder(resp *http.Response) (result SQLPoolBlobAuditingPolicy, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListBySQLPool lists auditing settings of a Sql pool.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// workspaceName - the name of the workspace
// SQLPoolName - SQL pool name
func (client SQLPoolBlobAuditingPoliciesClient) ListBySQLPool(ctx context.Context, resourceGroupName string, workspaceName string, SQLPoolName string) (result SQLPoolBlobAuditingPolicyListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SQLPoolBlobAuditingPoliciesClient.ListBySQLPool")
		defer func() {
			sc := -1
			if result.spbaplr.Response.Response != nil {
				sc = result.spbaplr.Response.Response.StatusCode
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
		return result, validation.NewError("synapse.SQLPoolBlobAuditingPoliciesClient", "ListBySQLPool", err.Error())
	}

	result.fn = client.listBySQLPoolNextResults
	req, err := client.ListBySQLPoolPreparer(ctx, resourceGroupName, workspaceName, SQLPoolName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.SQLPoolBlobAuditingPoliciesClient", "ListBySQLPool", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListBySQLPoolSender(req)
	if err != nil {
		result.spbaplr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "synapse.SQLPoolBlobAuditingPoliciesClient", "ListBySQLPool", resp, "Failure sending request")
		return
	}

	result.spbaplr, err = client.ListBySQLPoolResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.SQLPoolBlobAuditingPoliciesClient", "ListBySQLPool", resp, "Failure responding to request")
		return
	}
	if result.spbaplr.hasNextLink() && result.spbaplr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListBySQLPoolPreparer prepares the ListBySQLPool request.
func (client SQLPoolBlobAuditingPoliciesClient) ListBySQLPoolPreparer(ctx context.Context, resourceGroupName string, workspaceName string, SQLPoolName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"sqlPoolName":       autorest.Encode("path", SQLPoolName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workspaceName":     autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2020-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/sqlPools/{sqlPoolName}/auditingSettings", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListBySQLPoolSender sends the ListBySQLPool request. The method will close the
// http.Response Body if it receives an error.
func (client SQLPoolBlobAuditingPoliciesClient) ListBySQLPoolSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListBySQLPoolResponder handles the response to the ListBySQLPool request. The method always
// closes the http.Response Body.
func (client SQLPoolBlobAuditingPoliciesClient) ListBySQLPoolResponder(resp *http.Response) (result SQLPoolBlobAuditingPolicyListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listBySQLPoolNextResults retrieves the next set of results, if any.
func (client SQLPoolBlobAuditingPoliciesClient) listBySQLPoolNextResults(ctx context.Context, lastResults SQLPoolBlobAuditingPolicyListResult) (result SQLPoolBlobAuditingPolicyListResult, err error) {
	req, err := lastResults.sQLPoolBlobAuditingPolicyListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "synapse.SQLPoolBlobAuditingPoliciesClient", "listBySQLPoolNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListBySQLPoolSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "synapse.SQLPoolBlobAuditingPoliciesClient", "listBySQLPoolNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListBySQLPoolResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.SQLPoolBlobAuditingPoliciesClient", "listBySQLPoolNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListBySQLPoolComplete enumerates all values, automatically crossing page boundaries as required.
func (client SQLPoolBlobAuditingPoliciesClient) ListBySQLPoolComplete(ctx context.Context, resourceGroupName string, workspaceName string, SQLPoolName string) (result SQLPoolBlobAuditingPolicyListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SQLPoolBlobAuditingPoliciesClient.ListBySQLPool")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListBySQLPool(ctx, resourceGroupName, workspaceName, SQLPoolName)
	return
}
