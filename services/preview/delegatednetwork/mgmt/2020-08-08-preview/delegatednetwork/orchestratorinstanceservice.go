package delegatednetwork

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

// OrchestratorInstanceServiceClient is the DNC web api provides way to create, get and delete dnc controller
type OrchestratorInstanceServiceClient struct {
	BaseClient
}

// NewOrchestratorInstanceServiceClient creates an instance of the OrchestratorInstanceServiceClient client.
func NewOrchestratorInstanceServiceClient(subscriptionID string) OrchestratorInstanceServiceClient {
	return NewOrchestratorInstanceServiceClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewOrchestratorInstanceServiceClientWithBaseURI creates an instance of the OrchestratorInstanceServiceClient client
// using a custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign
// clouds, Azure stack).
func NewOrchestratorInstanceServiceClientWithBaseURI(baseURI string, subscriptionID string) OrchestratorInstanceServiceClient {
	return OrchestratorInstanceServiceClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Create create a orchestrator instance
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// resourceName - the name of the resource. It must be a minimum of 3 characters, and a maximum of 63.
// parameters - orchestratorInstance type parameters
func (client OrchestratorInstanceServiceClient) Create(ctx context.Context, resourceGroupName string, resourceName string, parameters Orchestrator) (result OrchestratorInstanceServiceCreateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/OrchestratorInstanceServiceClient.Create")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceName,
			Constraints: []validation.Constraint{{Target: "resourceName", Name: validation.MaxLength, Rule: 63, Chain: nil},
				{Target: "resourceName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "resourceName", Name: validation.Pattern, Rule: `^[a-z][a-z0-9]*$`, Chain: nil}}},
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.OrchestratorResourceProperties", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "parameters.OrchestratorResourceProperties.ControllerDetails", Name: validation.Null, Rule: true, Chain: nil}}}}}}); err != nil {
		return result, validation.NewError("delegatednetwork.OrchestratorInstanceServiceClient", "Create", err.Error())
	}

	req, err := client.CreatePreparer(ctx, resourceGroupName, resourceName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "delegatednetwork.OrchestratorInstanceServiceClient", "Create", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "delegatednetwork.OrchestratorInstanceServiceClient", "Create", nil, "Failure sending request")
		return
	}

	return
}

// CreatePreparer prepares the Create request.
func (client OrchestratorInstanceServiceClient) CreatePreparer(ctx context.Context, resourceGroupName string, resourceName string, parameters Orchestrator) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"resourceName":      autorest.Encode("path", resourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-08-08-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DelegatedNetwork/orchestrators/{resourceName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client OrchestratorInstanceServiceClient) CreateSender(req *http.Request) (future OrchestratorInstanceServiceCreateFuture, err error) {
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

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client OrchestratorInstanceServiceClient) CreateResponder(resp *http.Response) (result Orchestrator, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes the Orchestrator Instance
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// resourceName - the name of the resource. It must be a minimum of 3 characters, and a maximum of 63.
func (client OrchestratorInstanceServiceClient) Delete(ctx context.Context, resourceGroupName string, resourceName string) (result OrchestratorInstanceServiceDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/OrchestratorInstanceServiceClient.Delete")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceName,
			Constraints: []validation.Constraint{{Target: "resourceName", Name: validation.MaxLength, Rule: 63, Chain: nil},
				{Target: "resourceName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "resourceName", Name: validation.Pattern, Rule: `^[a-z][a-z0-9]*$`, Chain: nil}}},
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("delegatednetwork.OrchestratorInstanceServiceClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, resourceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "delegatednetwork.OrchestratorInstanceServiceClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "delegatednetwork.OrchestratorInstanceServiceClient", "Delete", nil, "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client OrchestratorInstanceServiceClient) DeletePreparer(ctx context.Context, resourceGroupName string, resourceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"resourceName":      autorest.Encode("path", resourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-08-08-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DelegatedNetwork/orchestrators/{resourceName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client OrchestratorInstanceServiceClient) DeleteSender(req *http.Request) (future OrchestratorInstanceServiceDeleteFuture, err error) {
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
func (client OrchestratorInstanceServiceClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// GetDetails gets details about the orchestrator instance.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// resourceName - the name of the resource. It must be a minimum of 3 characters, and a maximum of 63.
func (client OrchestratorInstanceServiceClient) GetDetails(ctx context.Context, resourceGroupName string, resourceName string) (result Orchestrator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/OrchestratorInstanceServiceClient.GetDetails")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceName,
			Constraints: []validation.Constraint{{Target: "resourceName", Name: validation.MaxLength, Rule: 63, Chain: nil},
				{Target: "resourceName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "resourceName", Name: validation.Pattern, Rule: `^[a-z][a-z0-9]*$`, Chain: nil}}},
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("delegatednetwork.OrchestratorInstanceServiceClient", "GetDetails", err.Error())
	}

	req, err := client.GetDetailsPreparer(ctx, resourceGroupName, resourceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "delegatednetwork.OrchestratorInstanceServiceClient", "GetDetails", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetDetailsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "delegatednetwork.OrchestratorInstanceServiceClient", "GetDetails", resp, "Failure sending request")
		return
	}

	result, err = client.GetDetailsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "delegatednetwork.OrchestratorInstanceServiceClient", "GetDetails", resp, "Failure responding to request")
		return
	}

	return
}

// GetDetailsPreparer prepares the GetDetails request.
func (client OrchestratorInstanceServiceClient) GetDetailsPreparer(ctx context.Context, resourceGroupName string, resourceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"resourceName":      autorest.Encode("path", resourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-08-08-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DelegatedNetwork/orchestrators/{resourceName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetDetailsSender sends the GetDetails request. The method will close the
// http.Response Body if it receives an error.
func (client OrchestratorInstanceServiceClient) GetDetailsSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetDetailsResponder handles the response to the GetDetails request. The method always
// closes the http.Response Body.
func (client OrchestratorInstanceServiceClient) GetDetailsResponder(resp *http.Response) (result Orchestrator, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByResourceGroup get all the OrchestratorInstances resources in a resource group.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
func (client OrchestratorInstanceServiceClient) ListByResourceGroup(ctx context.Context, resourceGroupName string) (result OrchestratorsPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/OrchestratorInstanceServiceClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.o.Response.Response != nil {
				sc = result.o.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("delegatednetwork.OrchestratorInstanceServiceClient", "ListByResourceGroup", err.Error())
	}

	result.fn = client.listByResourceGroupNextResults
	req, err := client.ListByResourceGroupPreparer(ctx, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "delegatednetwork.OrchestratorInstanceServiceClient", "ListByResourceGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.o.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "delegatednetwork.OrchestratorInstanceServiceClient", "ListByResourceGroup", resp, "Failure sending request")
		return
	}

	result.o, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "delegatednetwork.OrchestratorInstanceServiceClient", "ListByResourceGroup", resp, "Failure responding to request")
		return
	}
	if result.o.hasNextLink() && result.o.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByResourceGroupPreparer prepares the ListByResourceGroup request.
func (client OrchestratorInstanceServiceClient) ListByResourceGroupPreparer(ctx context.Context, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-08-08-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DelegatedNetwork/orchestrators", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByResourceGroupSender sends the ListByResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client OrchestratorInstanceServiceClient) ListByResourceGroupSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByResourceGroupResponder handles the response to the ListByResourceGroup request. The method always
// closes the http.Response Body.
func (client OrchestratorInstanceServiceClient) ListByResourceGroupResponder(resp *http.Response) (result Orchestrators, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByResourceGroupNextResults retrieves the next set of results, if any.
func (client OrchestratorInstanceServiceClient) listByResourceGroupNextResults(ctx context.Context, lastResults Orchestrators) (result Orchestrators, err error) {
	req, err := lastResults.orchestratorsPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "delegatednetwork.OrchestratorInstanceServiceClient", "listByResourceGroupNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "delegatednetwork.OrchestratorInstanceServiceClient", "listByResourceGroupNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "delegatednetwork.OrchestratorInstanceServiceClient", "listByResourceGroupNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByResourceGroupComplete enumerates all values, automatically crossing page boundaries as required.
func (client OrchestratorInstanceServiceClient) ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result OrchestratorsIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/OrchestratorInstanceServiceClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByResourceGroup(ctx, resourceGroupName)
	return
}

// ListBySubscription get all the orchestratorInstance resources in a subscription.
func (client OrchestratorInstanceServiceClient) ListBySubscription(ctx context.Context) (result OrchestratorsPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/OrchestratorInstanceServiceClient.ListBySubscription")
		defer func() {
			sc := -1
			if result.o.Response.Response != nil {
				sc = result.o.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("delegatednetwork.OrchestratorInstanceServiceClient", "ListBySubscription", err.Error())
	}

	result.fn = client.listBySubscriptionNextResults
	req, err := client.ListBySubscriptionPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "delegatednetwork.OrchestratorInstanceServiceClient", "ListBySubscription", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListBySubscriptionSender(req)
	if err != nil {
		result.o.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "delegatednetwork.OrchestratorInstanceServiceClient", "ListBySubscription", resp, "Failure sending request")
		return
	}

	result.o, err = client.ListBySubscriptionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "delegatednetwork.OrchestratorInstanceServiceClient", "ListBySubscription", resp, "Failure responding to request")
		return
	}
	if result.o.hasNextLink() && result.o.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListBySubscriptionPreparer prepares the ListBySubscription request.
func (client OrchestratorInstanceServiceClient) ListBySubscriptionPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-08-08-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.DelegatedNetwork/orchestrators", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListBySubscriptionSender sends the ListBySubscription request. The method will close the
// http.Response Body if it receives an error.
func (client OrchestratorInstanceServiceClient) ListBySubscriptionSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListBySubscriptionResponder handles the response to the ListBySubscription request. The method always
// closes the http.Response Body.
func (client OrchestratorInstanceServiceClient) ListBySubscriptionResponder(resp *http.Response) (result Orchestrators, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listBySubscriptionNextResults retrieves the next set of results, if any.
func (client OrchestratorInstanceServiceClient) listBySubscriptionNextResults(ctx context.Context, lastResults Orchestrators) (result Orchestrators, err error) {
	req, err := lastResults.orchestratorsPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "delegatednetwork.OrchestratorInstanceServiceClient", "listBySubscriptionNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListBySubscriptionSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "delegatednetwork.OrchestratorInstanceServiceClient", "listBySubscriptionNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListBySubscriptionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "delegatednetwork.OrchestratorInstanceServiceClient", "listBySubscriptionNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListBySubscriptionComplete enumerates all values, automatically crossing page boundaries as required.
func (client OrchestratorInstanceServiceClient) ListBySubscriptionComplete(ctx context.Context) (result OrchestratorsIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/OrchestratorInstanceServiceClient.ListBySubscription")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListBySubscription(ctx)
	return
}

// Patch update Orchestrator Instance
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// resourceName - the name of the resource. It must be a minimum of 3 characters, and a maximum of 63.
// parameters - orchestratorInstance update parameters
func (client OrchestratorInstanceServiceClient) Patch(ctx context.Context, resourceGroupName string, resourceName string, parameters OrchestratorResourceUpdateParameters) (result Orchestrator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/OrchestratorInstanceServiceClient.Patch")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceName,
			Constraints: []validation.Constraint{{Target: "resourceName", Name: validation.MaxLength, Rule: 63, Chain: nil},
				{Target: "resourceName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "resourceName", Name: validation.Pattern, Rule: `^[a-z][a-z0-9]*$`, Chain: nil}}},
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("delegatednetwork.OrchestratorInstanceServiceClient", "Patch", err.Error())
	}

	req, err := client.PatchPreparer(ctx, resourceGroupName, resourceName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "delegatednetwork.OrchestratorInstanceServiceClient", "Patch", nil, "Failure preparing request")
		return
	}

	resp, err := client.PatchSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "delegatednetwork.OrchestratorInstanceServiceClient", "Patch", resp, "Failure sending request")
		return
	}

	result, err = client.PatchResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "delegatednetwork.OrchestratorInstanceServiceClient", "Patch", resp, "Failure responding to request")
		return
	}

	return
}

// PatchPreparer prepares the Patch request.
func (client OrchestratorInstanceServiceClient) PatchPreparer(ctx context.Context, resourceGroupName string, resourceName string, parameters OrchestratorResourceUpdateParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"resourceName":      autorest.Encode("path", resourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-08-08-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DelegatedNetwork/orchestrators/{resourceName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// PatchSender sends the Patch request. The method will close the
// http.Response Body if it receives an error.
func (client OrchestratorInstanceServiceClient) PatchSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// PatchResponder handles the response to the Patch request. The method always
// closes the http.Response Body.
func (client OrchestratorInstanceServiceClient) PatchResponder(resp *http.Response) (result Orchestrator, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
