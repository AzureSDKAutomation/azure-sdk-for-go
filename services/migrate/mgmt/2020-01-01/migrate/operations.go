package migrate

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

// OperationsClient is the discover your workloads for Azure.
type OperationsClient struct {
	BaseClient
}

// NewOperationsClient creates an instance of the OperationsClient client.
func NewOperationsClient() OperationsClient {
	return NewOperationsClientWithBaseURI(DefaultBaseURI)
}

// NewOperationsClientWithBaseURI creates an instance of the OperationsClient client using a custom endpoint.  Use this
// when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewOperationsClientWithBaseURI(baseURI string) OperationsClient {
	return OperationsClient{NewWithBaseURI(baseURI)}
}

// List sends the list request.
// Parameters:
// APIVersion - the API version to use for this operation.
func (client OperationsClient) List(ctx context.Context, APIVersion string) (result OperationResultListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/OperationsClient.List")
		defer func() {
			sc := -1
			if result.orl.Response.Response != nil {
				sc = result.orl.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, APIVersion)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrate.OperationsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.orl.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "migrate.OperationsClient", "List", resp, "Failure sending request")
		return
	}

	result.orl, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrate.OperationsClient", "List", resp, "Failure responding to request")
		return
	}
	if result.orl.hasNextLink() && result.orl.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client OperationsClient) ListPreparer(ctx context.Context, APIVersion string) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/providers/Microsoft.OffAzure/operations"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client OperationsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client OperationsClient) ListResponder(resp *http.Response) (result OperationResultList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client OperationsClient) listNextResults(ctx context.Context, lastResults OperationResultList) (result OperationResultList, err error) {
	req, err := lastResults.operationResultListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "migrate.OperationsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "migrate.OperationsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrate.OperationsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client OperationsClient) ListComplete(ctx context.Context, APIVersion string) (result OperationResultListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/OperationsClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, APIVersion)
	return
}
