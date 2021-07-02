package mariadb

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

// QueryTextsClient is the the Microsoft Azure management API provides create, read, update, and delete functionality
// for Azure MariaDB resources including servers, databases, firewall rules, VNET rules, log files and configurations
// with new business model.
type QueryTextsClient struct {
	BaseClient
}

// NewQueryTextsClient creates an instance of the QueryTextsClient client.
func NewQueryTextsClient(subscriptionID string) QueryTextsClient {
	return NewQueryTextsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewQueryTextsClientWithBaseURI creates an instance of the QueryTextsClient client using a custom endpoint.  Use this
// when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewQueryTextsClientWithBaseURI(baseURI string, subscriptionID string) QueryTextsClient {
	return QueryTextsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get retrieve the Query-Store query texts for the queryId.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// serverName - the name of the server.
// queryID - the Query-Store query identifier.
func (client QueryTextsClient) Get(ctx context.Context, resourceGroupName string, serverName string, queryID string) (result QueryText, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/QueryTextsClient.Get")
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
		return result, validation.NewError("mariadb.QueryTextsClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, serverName, queryID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "mariadb.QueryTextsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "mariadb.QueryTextsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "mariadb.QueryTextsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client QueryTextsClient) GetPreparer(ctx context.Context, resourceGroupName string, serverName string, queryID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"queryId":           autorest.Encode("path", queryID),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverName":        autorest.Encode("path", serverName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMariaDB/servers/{serverName}/queryTexts/{queryId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client QueryTextsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client QueryTextsClient) GetResponder(resp *http.Response) (result QueryText, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByServer retrieve the Query-Store query texts for specified queryIds.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// serverName - the name of the server.
// queryIds - the query identifiers
func (client QueryTextsClient) ListByServer(ctx context.Context, resourceGroupName string, serverName string, queryIds []string) (result QueryTextsResultListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/QueryTextsClient.ListByServer")
		defer func() {
			sc := -1
			if result.qtrl.Response.Response != nil {
				sc = result.qtrl.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: queryIds,
			Constraints: []validation.Constraint{{Target: "queryIds", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("mariadb.QueryTextsClient", "ListByServer", err.Error())
	}

	result.fn = client.listByServerNextResults
	req, err := client.ListByServerPreparer(ctx, resourceGroupName, serverName, queryIds)
	if err != nil {
		err = autorest.NewErrorWithError(err, "mariadb.QueryTextsClient", "ListByServer", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByServerSender(req)
	if err != nil {
		result.qtrl.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "mariadb.QueryTextsClient", "ListByServer", resp, "Failure sending request")
		return
	}

	result.qtrl, err = client.ListByServerResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "mariadb.QueryTextsClient", "ListByServer", resp, "Failure responding to request")
		return
	}
	if result.qtrl.hasNextLink() && result.qtrl.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByServerPreparer prepares the ListByServer request.
func (client QueryTextsClient) ListByServerPreparer(ctx context.Context, resourceGroupName string, serverName string, queryIds []string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverName":        autorest.Encode("path", serverName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
		"queryIds":    queryIds,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMariaDB/servers/{serverName}/queryTexts", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByServerSender sends the ListByServer request. The method will close the
// http.Response Body if it receives an error.
func (client QueryTextsClient) ListByServerSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByServerResponder handles the response to the ListByServer request. The method always
// closes the http.Response Body.
func (client QueryTextsClient) ListByServerResponder(resp *http.Response) (result QueryTextsResultList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByServerNextResults retrieves the next set of results, if any.
func (client QueryTextsClient) listByServerNextResults(ctx context.Context, lastResults QueryTextsResultList) (result QueryTextsResultList, err error) {
	req, err := lastResults.queryTextsResultListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "mariadb.QueryTextsClient", "listByServerNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByServerSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "mariadb.QueryTextsClient", "listByServerNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByServerResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "mariadb.QueryTextsClient", "listByServerNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByServerComplete enumerates all values, automatically crossing page boundaries as required.
func (client QueryTextsClient) ListByServerComplete(ctx context.Context, resourceGroupName string, serverName string, queryIds []string) (result QueryTextsResultListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/QueryTextsClient.ListByServer")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByServer(ctx, resourceGroupName, serverName, queryIds)
	return
}
