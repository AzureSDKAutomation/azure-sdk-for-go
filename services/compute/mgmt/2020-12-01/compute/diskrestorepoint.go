package compute

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

// DiskRestorePointClient is the compute Client
type DiskRestorePointClient struct {
	BaseClient
}

// NewDiskRestorePointClient creates an instance of the DiskRestorePointClient client.
func NewDiskRestorePointClient(subscriptionID string) DiskRestorePointClient {
	return NewDiskRestorePointClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewDiskRestorePointClientWithBaseURI creates an instance of the DiskRestorePointClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewDiskRestorePointClientWithBaseURI(baseURI string, subscriptionID string) DiskRestorePointClient {
	return DiskRestorePointClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get get disk restorePoint resource
// Parameters:
// resourceGroupName - the name of the resource group.
// restorePointCollectionName - the name of the restore point collection that the disk restore point belongs.
// Supported characters for the name are a-z, A-Z, 0-9 and _. The maximum name length is 80 characters.
// VMRestorePointName - the name of the vm restore point that the disk disk restore point belongs. Supported
// characters for the name are a-z, A-Z, 0-9 and _. The maximum name length is 80 characters.
// diskRestorePointName - the name of the disk restore point created. Supported characters for the name are
// a-z, A-Z, 0-9 and _. The maximum name length is 80 characters.
func (client DiskRestorePointClient) Get(ctx context.Context, resourceGroupName string, restorePointCollectionName string, VMRestorePointName string, diskRestorePointName string) (result DiskRestorePoint, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DiskRestorePointClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, restorePointCollectionName, VMRestorePointName, diskRestorePointName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.DiskRestorePointClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "compute.DiskRestorePointClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.DiskRestorePointClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client DiskRestorePointClient) GetPreparer(ctx context.Context, resourceGroupName string, restorePointCollectionName string, VMRestorePointName string, diskRestorePointName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"diskRestorePointName":       autorest.Encode("path", diskRestorePointName),
		"resourceGroupName":          autorest.Encode("path", resourceGroupName),
		"restorePointCollectionName": autorest.Encode("path", restorePointCollectionName),
		"subscriptionId":             autorest.Encode("path", client.SubscriptionID),
		"vmRestorePointName":         autorest.Encode("path", VMRestorePointName),
	}

	const APIVersion = "2020-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/restorePointCollections/{restorePointCollectionName}/restorePoints/{vmRestorePointName}/diskRestorePoints/{diskRestorePointName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client DiskRestorePointClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client DiskRestorePointClient) GetResponder(resp *http.Response) (result DiskRestorePoint, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GrantAccess grants access to a diskRestorePoint.
// Parameters:
// resourceGroupName - the name of the resource group.
// restorePointCollectionName - the name of the restore point collection that the disk restore point belongs.
// Supported characters for the name are a-z, A-Z, 0-9 and _. The maximum name length is 80 characters.
// VMRestorePointName - the name of the vm restore point that the disk disk restore point belongs. Supported
// characters for the name are a-z, A-Z, 0-9 and _. The maximum name length is 80 characters.
// diskRestorePointName - the name of the disk restore point created. Supported characters for the name are
// a-z, A-Z, 0-9 and _. The maximum name length is 80 characters.
// grantAccessData - access data object supplied in the body of the get disk access operation.
func (client DiskRestorePointClient) GrantAccess(ctx context.Context, resourceGroupName string, restorePointCollectionName string, VMRestorePointName string, diskRestorePointName string, grantAccessData GrantAccessData) (result DiskRestorePointGrantAccessFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DiskRestorePointClient.GrantAccess")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: grantAccessData,
			Constraints: []validation.Constraint{{Target: "grantAccessData.DurationInSeconds", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("compute.DiskRestorePointClient", "GrantAccess", err.Error())
	}

	req, err := client.GrantAccessPreparer(ctx, resourceGroupName, restorePointCollectionName, VMRestorePointName, diskRestorePointName, grantAccessData)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.DiskRestorePointClient", "GrantAccess", nil, "Failure preparing request")
		return
	}

	result, err = client.GrantAccessSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.DiskRestorePointClient", "GrantAccess", nil, "Failure sending request")
		return
	}

	return
}

// GrantAccessPreparer prepares the GrantAccess request.
func (client DiskRestorePointClient) GrantAccessPreparer(ctx context.Context, resourceGroupName string, restorePointCollectionName string, VMRestorePointName string, diskRestorePointName string, grantAccessData GrantAccessData) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"diskRestorePointName":       autorest.Encode("path", diskRestorePointName),
		"resourceGroupName":          autorest.Encode("path", resourceGroupName),
		"restorePointCollectionName": autorest.Encode("path", restorePointCollectionName),
		"subscriptionId":             autorest.Encode("path", client.SubscriptionID),
		"vmRestorePointName":         autorest.Encode("path", VMRestorePointName),
	}

	const APIVersion = "2020-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/restorePointCollections/{restorePointCollectionName}/restorePoints/{vmRestorePointName}/diskRestorePoints/{diskRestorePointName}/beginGetAccess", pathParameters),
		autorest.WithJSON(grantAccessData),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GrantAccessSender sends the GrantAccess request. The method will close the
// http.Response Body if it receives an error.
func (client DiskRestorePointClient) GrantAccessSender(req *http.Request) (future DiskRestorePointGrantAccessFuture, err error) {
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

// GrantAccessResponder handles the response to the GrantAccess request. The method always
// closes the http.Response Body.
func (client DiskRestorePointClient) GrantAccessResponder(resp *http.Response) (result AccessURI, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByRestorePoint lists diskRestorePoints under a vmRestorePoint.
// Parameters:
// resourceGroupName - the name of the resource group.
// restorePointCollectionName - the name of the restore point collection that the disk restore point belongs.
// Supported characters for the name are a-z, A-Z, 0-9 and _. The maximum name length is 80 characters.
// VMRestorePointName - the name of the vm restore point that the disk disk restore point belongs. Supported
// characters for the name are a-z, A-Z, 0-9 and _. The maximum name length is 80 characters.
func (client DiskRestorePointClient) ListByRestorePoint(ctx context.Context, resourceGroupName string, restorePointCollectionName string, VMRestorePointName string) (result DiskRestorePointListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DiskRestorePointClient.ListByRestorePoint")
		defer func() {
			sc := -1
			if result.drpl.Response.Response != nil {
				sc = result.drpl.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByRestorePointNextResults
	req, err := client.ListByRestorePointPreparer(ctx, resourceGroupName, restorePointCollectionName, VMRestorePointName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.DiskRestorePointClient", "ListByRestorePoint", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByRestorePointSender(req)
	if err != nil {
		result.drpl.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "compute.DiskRestorePointClient", "ListByRestorePoint", resp, "Failure sending request")
		return
	}

	result.drpl, err = client.ListByRestorePointResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.DiskRestorePointClient", "ListByRestorePoint", resp, "Failure responding to request")
		return
	}
	if result.drpl.hasNextLink() && result.drpl.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByRestorePointPreparer prepares the ListByRestorePoint request.
func (client DiskRestorePointClient) ListByRestorePointPreparer(ctx context.Context, resourceGroupName string, restorePointCollectionName string, VMRestorePointName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":          autorest.Encode("path", resourceGroupName),
		"restorePointCollectionName": autorest.Encode("path", restorePointCollectionName),
		"subscriptionId":             autorest.Encode("path", client.SubscriptionID),
		"vmRestorePointName":         autorest.Encode("path", VMRestorePointName),
	}

	const APIVersion = "2020-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/restorePointCollections/{restorePointCollectionName}/restorePoints/{vmRestorePointName}/diskRestorePoints", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByRestorePointSender sends the ListByRestorePoint request. The method will close the
// http.Response Body if it receives an error.
func (client DiskRestorePointClient) ListByRestorePointSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByRestorePointResponder handles the response to the ListByRestorePoint request. The method always
// closes the http.Response Body.
func (client DiskRestorePointClient) ListByRestorePointResponder(resp *http.Response) (result DiskRestorePointList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByRestorePointNextResults retrieves the next set of results, if any.
func (client DiskRestorePointClient) listByRestorePointNextResults(ctx context.Context, lastResults DiskRestorePointList) (result DiskRestorePointList, err error) {
	req, err := lastResults.diskRestorePointListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "compute.DiskRestorePointClient", "listByRestorePointNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByRestorePointSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "compute.DiskRestorePointClient", "listByRestorePointNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByRestorePointResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.DiskRestorePointClient", "listByRestorePointNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByRestorePointComplete enumerates all values, automatically crossing page boundaries as required.
func (client DiskRestorePointClient) ListByRestorePointComplete(ctx context.Context, resourceGroupName string, restorePointCollectionName string, VMRestorePointName string) (result DiskRestorePointListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DiskRestorePointClient.ListByRestorePoint")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByRestorePoint(ctx, resourceGroupName, restorePointCollectionName, VMRestorePointName)
	return
}

// RevokeAccess revokes access to a diskRestorePoint.
// Parameters:
// resourceGroupName - the name of the resource group.
// restorePointCollectionName - the name of the restore point collection that the disk restore point belongs.
// Supported characters for the name are a-z, A-Z, 0-9 and _. The maximum name length is 80 characters.
// VMRestorePointName - the name of the vm restore point that the disk disk restore point belongs. Supported
// characters for the name are a-z, A-Z, 0-9 and _. The maximum name length is 80 characters.
// diskRestorePointName - the name of the disk restore point created. Supported characters for the name are
// a-z, A-Z, 0-9 and _. The maximum name length is 80 characters.
func (client DiskRestorePointClient) RevokeAccess(ctx context.Context, resourceGroupName string, restorePointCollectionName string, VMRestorePointName string, diskRestorePointName string) (result DiskRestorePointRevokeAccessFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DiskRestorePointClient.RevokeAccess")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.RevokeAccessPreparer(ctx, resourceGroupName, restorePointCollectionName, VMRestorePointName, diskRestorePointName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.DiskRestorePointClient", "RevokeAccess", nil, "Failure preparing request")
		return
	}

	result, err = client.RevokeAccessSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.DiskRestorePointClient", "RevokeAccess", nil, "Failure sending request")
		return
	}

	return
}

// RevokeAccessPreparer prepares the RevokeAccess request.
func (client DiskRestorePointClient) RevokeAccessPreparer(ctx context.Context, resourceGroupName string, restorePointCollectionName string, VMRestorePointName string, diskRestorePointName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"diskRestorePointName":       autorest.Encode("path", diskRestorePointName),
		"resourceGroupName":          autorest.Encode("path", resourceGroupName),
		"restorePointCollectionName": autorest.Encode("path", restorePointCollectionName),
		"subscriptionId":             autorest.Encode("path", client.SubscriptionID),
		"vmRestorePointName":         autorest.Encode("path", VMRestorePointName),
	}

	const APIVersion = "2020-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/restorePointCollections/{restorePointCollectionName}/restorePoints/{vmRestorePointName}/diskRestorePoints/{diskRestorePointName}/endGetAccess", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// RevokeAccessSender sends the RevokeAccess request. The method will close the
// http.Response Body if it receives an error.
func (client DiskRestorePointClient) RevokeAccessSender(req *http.Request) (future DiskRestorePointRevokeAccessFuture, err error) {
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

// RevokeAccessResponder handles the response to the RevokeAccess request. The method always
// closes the http.Response Body.
func (client DiskRestorePointClient) RevokeAccessResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}
