package insights

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

// PrivateLinkScopedResourcesClient is the monitor Management Client
type PrivateLinkScopedResourcesClient struct {
	BaseClient
}

// NewPrivateLinkScopedResourcesClient creates an instance of the PrivateLinkScopedResourcesClient client.
func NewPrivateLinkScopedResourcesClient(subscriptionID string) PrivateLinkScopedResourcesClient {
	return NewPrivateLinkScopedResourcesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewPrivateLinkScopedResourcesClientWithBaseURI creates an instance of the PrivateLinkScopedResourcesClient client
// using a custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign
// clouds, Azure stack).
func NewPrivateLinkScopedResourcesClientWithBaseURI(baseURI string, subscriptionID string) PrivateLinkScopedResourcesClient {
	return PrivateLinkScopedResourcesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate approve or reject a private endpoint connection with a given name.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// scopeName - the name of the Azure Monitor PrivateLinkScope resource.
// name - the name of the scoped resource object.
func (client PrivateLinkScopedResourcesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, scopeName string, name string, parameters ScopedResource) (result PrivateLinkScopedResourcesCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PrivateLinkScopedResourcesClient.CreateOrUpdate")
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
		return result, validation.NewError("insights.PrivateLinkScopedResourcesClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, scopeName, name, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insights.PrivateLinkScopedResourcesClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insights.PrivateLinkScopedResourcesClient", "CreateOrUpdate", nil, "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client PrivateLinkScopedResourcesClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, scopeName string, name string, parameters ScopedResource) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"scopeName":         autorest.Encode("path", scopeName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-10-17-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/privateLinkScopes/{scopeName}/scopedResources/{name}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client PrivateLinkScopedResourcesClient) CreateOrUpdateSender(req *http.Request) (future PrivateLinkScopedResourcesCreateOrUpdateFuture, err error) {
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
func (client PrivateLinkScopedResourcesClient) CreateOrUpdateResponder(resp *http.Response) (result ScopedResource, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a private endpoint connection with a given name.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// scopeName - the name of the Azure Monitor PrivateLinkScope resource.
// name - the name of the scoped resource object.
func (client PrivateLinkScopedResourcesClient) Delete(ctx context.Context, resourceGroupName string, scopeName string, name string) (result PrivateLinkScopedResourcesDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PrivateLinkScopedResourcesClient.Delete")
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
		return result, validation.NewError("insights.PrivateLinkScopedResourcesClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, scopeName, name)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insights.PrivateLinkScopedResourcesClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insights.PrivateLinkScopedResourcesClient", "Delete", nil, "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client PrivateLinkScopedResourcesClient) DeletePreparer(ctx context.Context, resourceGroupName string, scopeName string, name string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"scopeName":         autorest.Encode("path", scopeName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-10-17-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/privateLinkScopes/{scopeName}/scopedResources/{name}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client PrivateLinkScopedResourcesClient) DeleteSender(req *http.Request) (future PrivateLinkScopedResourcesDeleteFuture, err error) {
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

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client PrivateLinkScopedResourcesClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets a scoped resource in a private link scope.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// scopeName - the name of the Azure Monitor PrivateLinkScope resource.
// name - the name of the scoped resource object.
func (client PrivateLinkScopedResourcesClient) Get(ctx context.Context, resourceGroupName string, scopeName string, name string) (result ScopedResource, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PrivateLinkScopedResourcesClient.Get")
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
		return result, validation.NewError("insights.PrivateLinkScopedResourcesClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, scopeName, name)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insights.PrivateLinkScopedResourcesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "insights.PrivateLinkScopedResourcesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insights.PrivateLinkScopedResourcesClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client PrivateLinkScopedResourcesClient) GetPreparer(ctx context.Context, resourceGroupName string, scopeName string, name string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"scopeName":         autorest.Encode("path", scopeName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-10-17-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/privateLinkScopes/{scopeName}/scopedResources/{name}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client PrivateLinkScopedResourcesClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client PrivateLinkScopedResourcesClient) GetResponder(resp *http.Response) (result ScopedResource, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByPrivateLinkScope gets all private endpoint connections on a private link scope.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// scopeName - the name of the Azure Monitor PrivateLinkScope resource.
func (client PrivateLinkScopedResourcesClient) ListByPrivateLinkScope(ctx context.Context, resourceGroupName string, scopeName string) (result ScopedResourceListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PrivateLinkScopedResourcesClient.ListByPrivateLinkScope")
		defer func() {
			sc := -1
			if result.srlr.Response.Response != nil {
				sc = result.srlr.Response.Response.StatusCode
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
		return result, validation.NewError("insights.PrivateLinkScopedResourcesClient", "ListByPrivateLinkScope", err.Error())
	}

	result.fn = client.listByPrivateLinkScopeNextResults
	req, err := client.ListByPrivateLinkScopePreparer(ctx, resourceGroupName, scopeName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insights.PrivateLinkScopedResourcesClient", "ListByPrivateLinkScope", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByPrivateLinkScopeSender(req)
	if err != nil {
		result.srlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "insights.PrivateLinkScopedResourcesClient", "ListByPrivateLinkScope", resp, "Failure sending request")
		return
	}

	result.srlr, err = client.ListByPrivateLinkScopeResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insights.PrivateLinkScopedResourcesClient", "ListByPrivateLinkScope", resp, "Failure responding to request")
		return
	}
	if result.srlr.hasNextLink() && result.srlr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByPrivateLinkScopePreparer prepares the ListByPrivateLinkScope request.
func (client PrivateLinkScopedResourcesClient) ListByPrivateLinkScopePreparer(ctx context.Context, resourceGroupName string, scopeName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"scopeName":         autorest.Encode("path", scopeName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-10-17-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/privateLinkScopes/{scopeName}/scopedResources", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByPrivateLinkScopeSender sends the ListByPrivateLinkScope request. The method will close the
// http.Response Body if it receives an error.
func (client PrivateLinkScopedResourcesClient) ListByPrivateLinkScopeSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByPrivateLinkScopeResponder handles the response to the ListByPrivateLinkScope request. The method always
// closes the http.Response Body.
func (client PrivateLinkScopedResourcesClient) ListByPrivateLinkScopeResponder(resp *http.Response) (result ScopedResourceListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByPrivateLinkScopeNextResults retrieves the next set of results, if any.
func (client PrivateLinkScopedResourcesClient) listByPrivateLinkScopeNextResults(ctx context.Context, lastResults ScopedResourceListResult) (result ScopedResourceListResult, err error) {
	req, err := lastResults.scopedResourceListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "insights.PrivateLinkScopedResourcesClient", "listByPrivateLinkScopeNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByPrivateLinkScopeSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "insights.PrivateLinkScopedResourcesClient", "listByPrivateLinkScopeNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByPrivateLinkScopeResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insights.PrivateLinkScopedResourcesClient", "listByPrivateLinkScopeNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByPrivateLinkScopeComplete enumerates all values, automatically crossing page boundaries as required.
func (client PrivateLinkScopedResourcesClient) ListByPrivateLinkScopeComplete(ctx context.Context, resourceGroupName string, scopeName string) (result ScopedResourceListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PrivateLinkScopedResourcesClient.ListByPrivateLinkScope")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByPrivateLinkScope(ctx, resourceGroupName, scopeName)
	return
}
