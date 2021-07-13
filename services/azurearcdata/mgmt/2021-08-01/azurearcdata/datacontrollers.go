package azurearcdata

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

// DataControllersClient is the the AzureArcData management API provides a RESTful set of web APIs to manage Azure Data
// Services on Azure Arc Resources.
type DataControllersClient struct {
	BaseClient
}

// NewDataControllersClient creates an instance of the DataControllersClient client.
func NewDataControllersClient(subscriptionID string) DataControllersClient {
	return NewDataControllersClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewDataControllersClientWithBaseURI creates an instance of the DataControllersClient client using a custom endpoint.
// Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewDataControllersClientWithBaseURI(baseURI string, subscriptionID string) DataControllersClient {
	return DataControllersClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// DeleteDataController deletes a dataController resource
// Parameters:
// resourceGroupName - the name of the Azure resource group
func (client DataControllersClient) DeleteDataController(ctx context.Context, resourceGroupName string, dataControllerName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DataControllersClient.DeleteDataController")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeleteDataControllerPreparer(ctx, resourceGroupName, dataControllerName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azurearcdata.DataControllersClient", "DeleteDataController", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteDataControllerSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "azurearcdata.DataControllersClient", "DeleteDataController", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteDataControllerResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azurearcdata.DataControllersClient", "DeleteDataController", resp, "Failure responding to request")
		return
	}

	return
}

// DeleteDataControllerPreparer prepares the DeleteDataController request.
func (client DataControllersClient) DeleteDataControllerPreparer(ctx context.Context, resourceGroupName string, dataControllerName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"dataControllerName": autorest.Encode("path", dataControllerName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureArcData/dataControllers/{dataControllerName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteDataControllerSender sends the DeleteDataController request. The method will close the
// http.Response Body if it receives an error.
func (client DataControllersClient) DeleteDataControllerSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// DeleteDataControllerResponder handles the response to the DeleteDataController request. The method always
// closes the http.Response Body.
func (client DataControllersClient) DeleteDataControllerResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// GetDataController retrieves a dataController resource
// Parameters:
// resourceGroupName - the name of the Azure resource group
func (client DataControllersClient) GetDataController(ctx context.Context, resourceGroupName string, dataControllerName string) (result DataControllerResource, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DataControllersClient.GetDataController")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetDataControllerPreparer(ctx, resourceGroupName, dataControllerName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azurearcdata.DataControllersClient", "GetDataController", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetDataControllerSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "azurearcdata.DataControllersClient", "GetDataController", resp, "Failure sending request")
		return
	}

	result, err = client.GetDataControllerResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azurearcdata.DataControllersClient", "GetDataController", resp, "Failure responding to request")
		return
	}

	return
}

// GetDataControllerPreparer prepares the GetDataController request.
func (client DataControllersClient) GetDataControllerPreparer(ctx context.Context, resourceGroupName string, dataControllerName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"dataControllerName": autorest.Encode("path", dataControllerName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureArcData/dataControllers/{dataControllerName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetDataControllerSender sends the GetDataController request. The method will close the
// http.Response Body if it receives an error.
func (client DataControllersClient) GetDataControllerSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetDataControllerResponder handles the response to the GetDataController request. The method always
// closes the http.Response Body.
func (client DataControllersClient) GetDataControllerResponder(resp *http.Response) (result DataControllerResource, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListInGroup sends the list in group request.
// Parameters:
// resourceGroupName - the name of the Azure resource group
func (client DataControllersClient) ListInGroup(ctx context.Context, resourceGroupName string) (result PageOfDataControllerResourcePage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DataControllersClient.ListInGroup")
		defer func() {
			sc := -1
			if result.podcr.Response.Response != nil {
				sc = result.podcr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listInGroupNextResults
	req, err := client.ListInGroupPreparer(ctx, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azurearcdata.DataControllersClient", "ListInGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListInGroupSender(req)
	if err != nil {
		result.podcr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "azurearcdata.DataControllersClient", "ListInGroup", resp, "Failure sending request")
		return
	}

	result.podcr, err = client.ListInGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azurearcdata.DataControllersClient", "ListInGroup", resp, "Failure responding to request")
		return
	}
	if result.podcr.hasNextLink() && result.podcr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListInGroupPreparer prepares the ListInGroup request.
func (client DataControllersClient) ListInGroupPreparer(ctx context.Context, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureArcData/dataControllers", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListInGroupSender sends the ListInGroup request. The method will close the
// http.Response Body if it receives an error.
func (client DataControllersClient) ListInGroupSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListInGroupResponder handles the response to the ListInGroup request. The method always
// closes the http.Response Body.
func (client DataControllersClient) ListInGroupResponder(resp *http.Response) (result PageOfDataControllerResource, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listInGroupNextResults retrieves the next set of results, if any.
func (client DataControllersClient) listInGroupNextResults(ctx context.Context, lastResults PageOfDataControllerResource) (result PageOfDataControllerResource, err error) {
	req, err := lastResults.pageOfDataControllerResourcePreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "azurearcdata.DataControllersClient", "listInGroupNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListInGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "azurearcdata.DataControllersClient", "listInGroupNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListInGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azurearcdata.DataControllersClient", "listInGroupNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListInGroupComplete enumerates all values, automatically crossing page boundaries as required.
func (client DataControllersClient) ListInGroupComplete(ctx context.Context, resourceGroupName string) (result PageOfDataControllerResourceIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DataControllersClient.ListInGroup")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListInGroup(ctx, resourceGroupName)
	return
}

// ListInSubscription sends the list in subscription request.
func (client DataControllersClient) ListInSubscription(ctx context.Context) (result PageOfDataControllerResourcePage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DataControllersClient.ListInSubscription")
		defer func() {
			sc := -1
			if result.podcr.Response.Response != nil {
				sc = result.podcr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listInSubscriptionNextResults
	req, err := client.ListInSubscriptionPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azurearcdata.DataControllersClient", "ListInSubscription", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListInSubscriptionSender(req)
	if err != nil {
		result.podcr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "azurearcdata.DataControllersClient", "ListInSubscription", resp, "Failure sending request")
		return
	}

	result.podcr, err = client.ListInSubscriptionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azurearcdata.DataControllersClient", "ListInSubscription", resp, "Failure responding to request")
		return
	}
	if result.podcr.hasNextLink() && result.podcr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListInSubscriptionPreparer prepares the ListInSubscription request.
func (client DataControllersClient) ListInSubscriptionPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.AzureArcData/dataControllers", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListInSubscriptionSender sends the ListInSubscription request. The method will close the
// http.Response Body if it receives an error.
func (client DataControllersClient) ListInSubscriptionSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListInSubscriptionResponder handles the response to the ListInSubscription request. The method always
// closes the http.Response Body.
func (client DataControllersClient) ListInSubscriptionResponder(resp *http.Response) (result PageOfDataControllerResource, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listInSubscriptionNextResults retrieves the next set of results, if any.
func (client DataControllersClient) listInSubscriptionNextResults(ctx context.Context, lastResults PageOfDataControllerResource) (result PageOfDataControllerResource, err error) {
	req, err := lastResults.pageOfDataControllerResourcePreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "azurearcdata.DataControllersClient", "listInSubscriptionNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListInSubscriptionSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "azurearcdata.DataControllersClient", "listInSubscriptionNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListInSubscriptionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azurearcdata.DataControllersClient", "listInSubscriptionNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListInSubscriptionComplete enumerates all values, automatically crossing page boundaries as required.
func (client DataControllersClient) ListInSubscriptionComplete(ctx context.Context) (result PageOfDataControllerResourceIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DataControllersClient.ListInSubscription")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListInSubscription(ctx)
	return
}

// PatchDataController updates a dataController resource
// Parameters:
// resourceGroupName - the name of the Azure resource group
// dataControllerResource - the update data controller resource
func (client DataControllersClient) PatchDataController(ctx context.Context, resourceGroupName string, dataControllerName string, dataControllerResource DataControllerUpdate) (result DataControllerResource, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DataControllersClient.PatchDataController")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.PatchDataControllerPreparer(ctx, resourceGroupName, dataControllerName, dataControllerResource)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azurearcdata.DataControllersClient", "PatchDataController", nil, "Failure preparing request")
		return
	}

	resp, err := client.PatchDataControllerSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "azurearcdata.DataControllersClient", "PatchDataController", resp, "Failure sending request")
		return
	}

	result, err = client.PatchDataControllerResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azurearcdata.DataControllersClient", "PatchDataController", resp, "Failure responding to request")
		return
	}

	return
}

// PatchDataControllerPreparer prepares the PatchDataController request.
func (client DataControllersClient) PatchDataControllerPreparer(ctx context.Context, resourceGroupName string, dataControllerName string, dataControllerResource DataControllerUpdate) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"dataControllerName": autorest.Encode("path", dataControllerName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureArcData/dataControllers/{dataControllerName}", pathParameters),
		autorest.WithJSON(dataControllerResource),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// PatchDataControllerSender sends the PatchDataController request. The method will close the
// http.Response Body if it receives an error.
func (client DataControllersClient) PatchDataControllerSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// PatchDataControllerResponder handles the response to the PatchDataController request. The method always
// closes the http.Response Body.
func (client DataControllersClient) PatchDataControllerResponder(resp *http.Response) (result DataControllerResource, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// PutDataController creates or replaces a dataController resource
// Parameters:
// resourceGroupName - the name of the Azure resource group
// dataControllerResource - desc
func (client DataControllersClient) PutDataController(ctx context.Context, resourceGroupName string, dataControllerResource DataControllerResource, dataControllerName string) (result DataControllersPutDataControllerFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DataControllersClient.PutDataController")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: dataControllerResource,
			Constraints: []validation.Constraint{{Target: "dataControllerResource.Properties", Name: validation.Null, Rule: true,
				Chain: []validation.Constraint{{Target: "dataControllerResource.Properties.OnPremiseProperty", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "dataControllerResource.Properties.OnPremiseProperty.ID", Name: validation.Null, Rule: true, Chain: nil},
						{Target: "dataControllerResource.Properties.OnPremiseProperty.PublicSigningKey", Name: validation.Null, Rule: true, Chain: nil},
					}},
				}}}}}); err != nil {
		return result, validation.NewError("azurearcdata.DataControllersClient", "PutDataController", err.Error())
	}

	req, err := client.PutDataControllerPreparer(ctx, resourceGroupName, dataControllerResource, dataControllerName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azurearcdata.DataControllersClient", "PutDataController", nil, "Failure preparing request")
		return
	}

	result, err = client.PutDataControllerSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azurearcdata.DataControllersClient", "PutDataController", nil, "Failure sending request")
		return
	}

	return
}

// PutDataControllerPreparer prepares the PutDataController request.
func (client DataControllersClient) PutDataControllerPreparer(ctx context.Context, resourceGroupName string, dataControllerResource DataControllerResource, dataControllerName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"dataControllerName": autorest.Encode("path", dataControllerName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureArcData/dataControllers/{dataControllerName}", pathParameters),
		autorest.WithJSON(dataControllerResource),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// PutDataControllerSender sends the PutDataController request. The method will close the
// http.Response Body if it receives an error.
func (client DataControllersClient) PutDataControllerSender(req *http.Request) (future DataControllersPutDataControllerFuture, err error) {
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

// PutDataControllerResponder handles the response to the PutDataController request. The method always
// closes the http.Response Body.
func (client DataControllersClient) PutDataControllerResponder(resp *http.Response) (result DataControllerResource, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
