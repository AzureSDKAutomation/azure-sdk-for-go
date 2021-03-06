package resourcemover

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// UnresolvedDependenciesClient is the a first party Azure service orchestrating the move of Azure resources from one
// Azure region to another or between zones within a region.
type UnresolvedDependenciesClient struct {
	BaseClient
}

// NewUnresolvedDependenciesClient creates an instance of the UnresolvedDependenciesClient client.
func NewUnresolvedDependenciesClient(subscriptionID string) UnresolvedDependenciesClient {
	return NewUnresolvedDependenciesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewUnresolvedDependenciesClientWithBaseURI creates an instance of the UnresolvedDependenciesClient client using a
// custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds,
// Azure stack).
func NewUnresolvedDependenciesClientWithBaseURI(baseURI string, subscriptionID string) UnresolvedDependenciesClient {
	return UnresolvedDependenciesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get gets a list of unresolved dependencies.
// Parameters:
// resourceGroupName - the Resource Group Name.
// moveCollectionName - the Move Collection Name.
// dependencyLevel - defines the dependency level.
// orderby - oData order by query option. For example, you can use $orderby=Count desc.
// filter - the filter to apply on the operation. For example, $apply=filter(count eq 2).
func (client UnresolvedDependenciesClient) Get(ctx context.Context, resourceGroupName string, moveCollectionName string, dependencyLevel DependencyLevel, orderby string, filter string) (result UnresolvedDependencyCollectionPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/UnresolvedDependenciesClient.Get")
		defer func() {
			sc := -1
			if result.udc.Response.Response != nil {
				sc = result.udc.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.getNextResults
	req, err := client.GetPreparer(ctx, resourceGroupName, moveCollectionName, dependencyLevel, orderby, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourcemover.UnresolvedDependenciesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.udc.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "resourcemover.UnresolvedDependenciesClient", "Get", resp, "Failure sending request")
		return
	}

	result.udc, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourcemover.UnresolvedDependenciesClient", "Get", resp, "Failure responding to request")
		return
	}
	if result.udc.hasNextLink() && result.udc.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client UnresolvedDependenciesClient) GetPreparer(ctx context.Context, resourceGroupName string, moveCollectionName string, dependencyLevel DependencyLevel, orderby string, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"moveCollectionName": autorest.Encode("path", moveCollectionName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(string(dependencyLevel)) > 0 {
		queryParameters["dependencyLevel"] = autorest.Encode("query", dependencyLevel)
	}
	if len(orderby) > 0 {
		queryParameters["$orderby"] = autorest.Encode("query", orderby)
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/moveCollections/{moveCollectionName}/unresolvedDependencies", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client UnresolvedDependenciesClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client UnresolvedDependenciesClient) GetResponder(resp *http.Response) (result UnresolvedDependencyCollection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// getNextResults retrieves the next set of results, if any.
func (client UnresolvedDependenciesClient) getNextResults(ctx context.Context, lastResults UnresolvedDependencyCollection) (result UnresolvedDependencyCollection, err error) {
	req, err := lastResults.unresolvedDependencyCollectionPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "resourcemover.UnresolvedDependenciesClient", "getNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "resourcemover.UnresolvedDependenciesClient", "getNextResults", resp, "Failure sending next results request")
	}
	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourcemover.UnresolvedDependenciesClient", "getNextResults", resp, "Failure responding to next results request")
	}
	return
}

// GetComplete enumerates all values, automatically crossing page boundaries as required.
func (client UnresolvedDependenciesClient) GetComplete(ctx context.Context, resourceGroupName string, moveCollectionName string, dependencyLevel DependencyLevel, orderby string, filter string) (result UnresolvedDependencyCollectionIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/UnresolvedDependenciesClient.Get")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.Get(ctx, resourceGroupName, moveCollectionName, dependencyLevel, orderby, filter)
	return
}
