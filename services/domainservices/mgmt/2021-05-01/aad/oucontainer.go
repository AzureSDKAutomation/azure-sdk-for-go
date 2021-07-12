package aad

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

// OuContainerClient is the the AAD Domain Services API.
type OuContainerClient struct {
	BaseClient
}

// NewOuContainerClient creates an instance of the OuContainerClient client.
func NewOuContainerClient(subscriptionID string) OuContainerClient {
	return NewOuContainerClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewOuContainerClientWithBaseURI creates an instance of the OuContainerClient client using a custom endpoint.  Use
// this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewOuContainerClientWithBaseURI(baseURI string, subscriptionID string) OuContainerClient {
	return OuContainerClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Create the Create OuContainer operation creates a new OuContainer under the specified Domain Service instance.
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription. The name is case
// insensitive.
// domainServiceName - the name of the domain service.
// ouContainerName - the name of the OuContainer.
// containerAccount - container Account Description.
func (client OuContainerClient) Create(ctx context.Context, resourceGroupName string, domainServiceName string, ouContainerName string, containerAccount ContainerAccount) (result OuContainerCreateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/OuContainerClient.Create")
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
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("aad.OuContainerClient", "Create", err.Error())
	}

	req, err := client.CreatePreparer(ctx, resourceGroupName, domainServiceName, ouContainerName, containerAccount)
	if err != nil {
		err = autorest.NewErrorWithError(err, "aad.OuContainerClient", "Create", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "aad.OuContainerClient", "Create", nil, "Failure sending request")
		return
	}

	return
}

// CreatePreparer prepares the Create request.
func (client OuContainerClient) CreatePreparer(ctx context.Context, resourceGroupName string, domainServiceName string, ouContainerName string, containerAccount ContainerAccount) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"domainServiceName": autorest.Encode("path", domainServiceName),
		"ouContainerName":   autorest.Encode("path", ouContainerName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Aad/domainServices/{domainServiceName}/ouContainer/{ouContainerName}", pathParameters),
		autorest.WithJSON(containerAccount),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client OuContainerClient) CreateSender(req *http.Request) (future OuContainerCreateFuture, err error) {
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
func (client OuContainerClient) CreateResponder(resp *http.Response) (result OuContainer, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete the Delete OuContainer operation deletes specified OuContainer.
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription. The name is case
// insensitive.
// domainServiceName - the name of the domain service.
// ouContainerName - the name of the OuContainer.
func (client OuContainerClient) Delete(ctx context.Context, resourceGroupName string, domainServiceName string, ouContainerName string) (result OuContainerDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/OuContainerClient.Delete")
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
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("aad.OuContainerClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, domainServiceName, ouContainerName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "aad.OuContainerClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "aad.OuContainerClient", "Delete", nil, "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client OuContainerClient) DeletePreparer(ctx context.Context, resourceGroupName string, domainServiceName string, ouContainerName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"domainServiceName": autorest.Encode("path", domainServiceName),
		"ouContainerName":   autorest.Encode("path", ouContainerName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Aad/domainServices/{domainServiceName}/ouContainer/{ouContainerName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client OuContainerClient) DeleteSender(req *http.Request) (future OuContainerDeleteFuture, err error) {
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
func (client OuContainerClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get get OuContainer in DomainService instance.
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription. The name is case
// insensitive.
// domainServiceName - the name of the domain service.
// ouContainerName - the name of the OuContainer.
func (client OuContainerClient) Get(ctx context.Context, resourceGroupName string, domainServiceName string, ouContainerName string) (result OuContainer, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/OuContainerClient.Get")
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
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("aad.OuContainerClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, domainServiceName, ouContainerName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "aad.OuContainerClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "aad.OuContainerClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "aad.OuContainerClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client OuContainerClient) GetPreparer(ctx context.Context, resourceGroupName string, domainServiceName string, ouContainerName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"domainServiceName": autorest.Encode("path", domainServiceName),
		"ouContainerName":   autorest.Encode("path", ouContainerName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Aad/domainServices/{domainServiceName}/ouContainer/{ouContainerName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client OuContainerClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client OuContainerClient) GetResponder(resp *http.Response) (result OuContainer, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List the List of OuContainers in DomainService instance.
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription. The name is case
// insensitive.
// domainServiceName - the name of the domain service.
func (client OuContainerClient) List(ctx context.Context, resourceGroupName string, domainServiceName string) (result OuContainerListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/OuContainerClient.List")
		defer func() {
			sc := -1
			if result.oclr.Response.Response != nil {
				sc = result.oclr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("aad.OuContainerClient", "List", err.Error())
	}

	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName, domainServiceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "aad.OuContainerClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.oclr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "aad.OuContainerClient", "List", resp, "Failure sending request")
		return
	}

	result.oclr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "aad.OuContainerClient", "List", resp, "Failure responding to request")
		return
	}
	if result.oclr.hasNextLink() && result.oclr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client OuContainerClient) ListPreparer(ctx context.Context, resourceGroupName string, domainServiceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"domainServiceName": autorest.Encode("path", domainServiceName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Aad/domainServices/{domainServiceName}/ouContainer", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client OuContainerClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client OuContainerClient) ListResponder(resp *http.Response) (result OuContainerListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client OuContainerClient) listNextResults(ctx context.Context, lastResults OuContainerListResult) (result OuContainerListResult, err error) {
	req, err := lastResults.ouContainerListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "aad.OuContainerClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "aad.OuContainerClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "aad.OuContainerClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client OuContainerClient) ListComplete(ctx context.Context, resourceGroupName string, domainServiceName string) (result OuContainerListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/OuContainerClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceGroupName, domainServiceName)
	return
}

// Update the Update OuContainer operation can be used to update the existing OuContainers.
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription. The name is case
// insensitive.
// domainServiceName - the name of the domain service.
// ouContainerName - the name of the OuContainer.
// containerAccount - container Account Description.
func (client OuContainerClient) Update(ctx context.Context, resourceGroupName string, domainServiceName string, ouContainerName string, containerAccount ContainerAccount) (result OuContainerUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/OuContainerClient.Update")
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
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("aad.OuContainerClient", "Update", err.Error())
	}

	req, err := client.UpdatePreparer(ctx, resourceGroupName, domainServiceName, ouContainerName, containerAccount)
	if err != nil {
		err = autorest.NewErrorWithError(err, "aad.OuContainerClient", "Update", nil, "Failure preparing request")
		return
	}

	result, err = client.UpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "aad.OuContainerClient", "Update", nil, "Failure sending request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client OuContainerClient) UpdatePreparer(ctx context.Context, resourceGroupName string, domainServiceName string, ouContainerName string, containerAccount ContainerAccount) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"domainServiceName": autorest.Encode("path", domainServiceName),
		"ouContainerName":   autorest.Encode("path", ouContainerName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Aad/domainServices/{domainServiceName}/ouContainer/{ouContainerName}", pathParameters),
		autorest.WithJSON(containerAccount),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client OuContainerClient) UpdateSender(req *http.Request) (future OuContainerUpdateFuture, err error) {
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

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client OuContainerClient) UpdateResponder(resp *http.Response) (result OuContainer, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
