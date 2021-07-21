package containerregistry

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

// ImportPipelinesClient is the client for the ImportPipelines methods of the Containerregistry service.
type ImportPipelinesClient struct {
	BaseClient
}

// NewImportPipelinesClient creates an instance of the ImportPipelinesClient client.
func NewImportPipelinesClient(subscriptionID string) ImportPipelinesClient {
	return NewImportPipelinesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewImportPipelinesClientWithBaseURI creates an instance of the ImportPipelinesClient client using a custom endpoint.
// Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewImportPipelinesClientWithBaseURI(baseURI string, subscriptionID string) ImportPipelinesClient {
	return ImportPipelinesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Create creates an import pipeline for a container registry with the specified parameters.
// Parameters:
// resourceGroupName - the name of the resource group to which the container registry belongs.
// registryName - the name of the container registry.
// importPipelineName - the name of the import pipeline.
// importPipelineCreateParameters - the parameters for creating an import pipeline.
func (client ImportPipelinesClient) Create(ctx context.Context, resourceGroupName string, registryName string, importPipelineName string, importPipelineCreateParameters ImportPipeline) (result ImportPipelinesCreateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ImportPipelinesClient.Create")
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
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: registryName,
			Constraints: []validation.Constraint{{Target: "registryName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "registryName", Name: validation.MinLength, Rule: 5, Chain: nil},
				{Target: "registryName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]*$`, Chain: nil}}},
		{TargetValue: importPipelineName,
			Constraints: []validation.Constraint{{Target: "importPipelineName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "importPipelineName", Name: validation.MinLength, Rule: 5, Chain: nil},
				{Target: "importPipelineName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]*$`, Chain: nil}}},
		{TargetValue: importPipelineCreateParameters,
			Constraints: []validation.Constraint{{Target: "importPipelineCreateParameters.ImportPipelineProperties", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "importPipelineCreateParameters.ImportPipelineProperties.Source", Name: validation.Null, Rule: true,
					Chain: []validation.Constraint{{Target: "importPipelineCreateParameters.ImportPipelineProperties.Source.KeyVaultURI", Name: validation.Null, Rule: true, Chain: nil}}},
				}}}}}); err != nil {
		return result, validation.NewError("containerregistry.ImportPipelinesClient", "Create", err.Error())
	}

	req, err := client.CreatePreparer(ctx, resourceGroupName, registryName, importPipelineName, importPipelineCreateParameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.ImportPipelinesClient", "Create", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.ImportPipelinesClient", "Create", nil, "Failure sending request")
		return
	}

	return
}

// CreatePreparer prepares the Create request.
func (client ImportPipelinesClient) CreatePreparer(ctx context.Context, resourceGroupName string, registryName string, importPipelineName string, importPipelineCreateParameters ImportPipeline) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"importPipelineName": autorest.Encode("path", importPipelineName),
		"registryName":       autorest.Encode("path", registryName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/importPipelines/{importPipelineName}", pathParameters),
		autorest.WithJSON(importPipelineCreateParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client ImportPipelinesClient) CreateSender(req *http.Request) (future ImportPipelinesCreateFuture, err error) {
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
func (client ImportPipelinesClient) CreateResponder(resp *http.Response) (result ImportPipeline, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes an import pipeline from a container registry.
// Parameters:
// resourceGroupName - the name of the resource group to which the container registry belongs.
// registryName - the name of the container registry.
// importPipelineName - the name of the import pipeline.
func (client ImportPipelinesClient) Delete(ctx context.Context, resourceGroupName string, registryName string, importPipelineName string) (result ImportPipelinesDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ImportPipelinesClient.Delete")
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
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: registryName,
			Constraints: []validation.Constraint{{Target: "registryName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "registryName", Name: validation.MinLength, Rule: 5, Chain: nil},
				{Target: "registryName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]*$`, Chain: nil}}},
		{TargetValue: importPipelineName,
			Constraints: []validation.Constraint{{Target: "importPipelineName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "importPipelineName", Name: validation.MinLength, Rule: 5, Chain: nil},
				{Target: "importPipelineName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("containerregistry.ImportPipelinesClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, registryName, importPipelineName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.ImportPipelinesClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.ImportPipelinesClient", "Delete", nil, "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client ImportPipelinesClient) DeletePreparer(ctx context.Context, resourceGroupName string, registryName string, importPipelineName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"importPipelineName": autorest.Encode("path", importPipelineName),
		"registryName":       autorest.Encode("path", registryName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/importPipelines/{importPipelineName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ImportPipelinesClient) DeleteSender(req *http.Request) (future ImportPipelinesDeleteFuture, err error) {
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
func (client ImportPipelinesClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets the properties of the import pipeline.
// Parameters:
// resourceGroupName - the name of the resource group to which the container registry belongs.
// registryName - the name of the container registry.
// importPipelineName - the name of the import pipeline.
func (client ImportPipelinesClient) Get(ctx context.Context, resourceGroupName string, registryName string, importPipelineName string) (result ImportPipeline, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ImportPipelinesClient.Get")
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
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: registryName,
			Constraints: []validation.Constraint{{Target: "registryName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "registryName", Name: validation.MinLength, Rule: 5, Chain: nil},
				{Target: "registryName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]*$`, Chain: nil}}},
		{TargetValue: importPipelineName,
			Constraints: []validation.Constraint{{Target: "importPipelineName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "importPipelineName", Name: validation.MinLength, Rule: 5, Chain: nil},
				{Target: "importPipelineName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("containerregistry.ImportPipelinesClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, registryName, importPipelineName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.ImportPipelinesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "containerregistry.ImportPipelinesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.ImportPipelinesClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client ImportPipelinesClient) GetPreparer(ctx context.Context, resourceGroupName string, registryName string, importPipelineName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"importPipelineName": autorest.Encode("path", importPipelineName),
		"registryName":       autorest.Encode("path", registryName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/importPipelines/{importPipelineName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ImportPipelinesClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ImportPipelinesClient) GetResponder(resp *http.Response) (result ImportPipeline, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List lists all import pipelines for the specified container registry.
// Parameters:
// resourceGroupName - the name of the resource group to which the container registry belongs.
// registryName - the name of the container registry.
func (client ImportPipelinesClient) List(ctx context.Context, resourceGroupName string, registryName string) (result ImportPipelineListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ImportPipelinesClient.List")
		defer func() {
			sc := -1
			if result.iplr.Response.Response != nil {
				sc = result.iplr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: registryName,
			Constraints: []validation.Constraint{{Target: "registryName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "registryName", Name: validation.MinLength, Rule: 5, Chain: nil},
				{Target: "registryName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("containerregistry.ImportPipelinesClient", "List", err.Error())
	}

	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName, registryName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.ImportPipelinesClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.iplr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "containerregistry.ImportPipelinesClient", "List", resp, "Failure sending request")
		return
	}

	result.iplr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.ImportPipelinesClient", "List", resp, "Failure responding to request")
		return
	}
	if result.iplr.hasNextLink() && result.iplr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client ImportPipelinesClient) ListPreparer(ctx context.Context, resourceGroupName string, registryName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"registryName":      autorest.Encode("path", registryName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/importPipelines", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ImportPipelinesClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ImportPipelinesClient) ListResponder(resp *http.Response) (result ImportPipelineListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client ImportPipelinesClient) listNextResults(ctx context.Context, lastResults ImportPipelineListResult) (result ImportPipelineListResult, err error) {
	req, err := lastResults.importPipelineListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "containerregistry.ImportPipelinesClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "containerregistry.ImportPipelinesClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.ImportPipelinesClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client ImportPipelinesClient) ListComplete(ctx context.Context, resourceGroupName string, registryName string) (result ImportPipelineListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ImportPipelinesClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceGroupName, registryName)
	return
}
