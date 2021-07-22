package desktopvirtualization

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

// UserSessionsClient is the client for the UserSessions methods of the Desktopvirtualization service.
type UserSessionsClient struct {
	BaseClient
}

// NewUserSessionsClient creates an instance of the UserSessionsClient client.
func NewUserSessionsClient(subscriptionID string) UserSessionsClient {
	return NewUserSessionsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewUserSessionsClientWithBaseURI creates an instance of the UserSessionsClient client using a custom endpoint.  Use
// this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewUserSessionsClientWithBaseURI(baseURI string, subscriptionID string) UserSessionsClient {
	return UserSessionsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Delete remove a userSession.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// hostPoolName - the name of the host pool within the specified resource group
// sessionHostName - the name of the session host within the specified host pool
// userSessionID - the name of the user session within the specified session host
// force - force flag to login off userSession.
func (client UserSessionsClient) Delete(ctx context.Context, resourceGroupName string, hostPoolName string, sessionHostName string, userSessionID string, force *bool) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/UserSessionsClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
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
		{TargetValue: hostPoolName,
			Constraints: []validation.Constraint{{Target: "hostPoolName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "hostPoolName", Name: validation.MinLength, Rule: 3, Chain: nil}}},
		{TargetValue: sessionHostName,
			Constraints: []validation.Constraint{{Target: "sessionHostName", Name: validation.MaxLength, Rule: 48, Chain: nil},
				{Target: "sessionHostName", Name: validation.MinLength, Rule: 3, Chain: nil}}},
		{TargetValue: userSessionID,
			Constraints: []validation.Constraint{{Target: "userSessionID", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "userSessionID", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("desktopvirtualization.UserSessionsClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, hostPoolName, sessionHostName, userSessionID, force)
	if err != nil {
		err = autorest.NewErrorWithError(err, "desktopvirtualization.UserSessionsClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "desktopvirtualization.UserSessionsClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "desktopvirtualization.UserSessionsClient", "Delete", resp, "Failure responding to request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client UserSessionsClient) DeletePreparer(ctx context.Context, resourceGroupName string, hostPoolName string, sessionHostName string, userSessionID string, force *bool) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"hostPoolName":      autorest.Encode("path", hostPoolName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"sessionHostName":   autorest.Encode("path", sessionHostName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"userSessionId":     autorest.Encode("path", userSessionID),
	}

	const APIVersion = "2019-01-23-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if force != nil {
		queryParameters["force"] = autorest.Encode("query", *force)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DesktopVirtualization/hostPools/{hostPoolName}/sessionHosts/{sessionHostName}/userSessions/{userSessionId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client UserSessionsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client UserSessionsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Disconnect disconnect a userSession.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// hostPoolName - the name of the host pool within the specified resource group
// sessionHostName - the name of the session host within the specified host pool
// userSessionID - the name of the user session within the specified session host
func (client UserSessionsClient) Disconnect(ctx context.Context, resourceGroupName string, hostPoolName string, sessionHostName string, userSessionID string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/UserSessionsClient.Disconnect")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
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
		{TargetValue: hostPoolName,
			Constraints: []validation.Constraint{{Target: "hostPoolName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "hostPoolName", Name: validation.MinLength, Rule: 3, Chain: nil}}},
		{TargetValue: sessionHostName,
			Constraints: []validation.Constraint{{Target: "sessionHostName", Name: validation.MaxLength, Rule: 48, Chain: nil},
				{Target: "sessionHostName", Name: validation.MinLength, Rule: 3, Chain: nil}}},
		{TargetValue: userSessionID,
			Constraints: []validation.Constraint{{Target: "userSessionID", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "userSessionID", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("desktopvirtualization.UserSessionsClient", "Disconnect", err.Error())
	}

	req, err := client.DisconnectPreparer(ctx, resourceGroupName, hostPoolName, sessionHostName, userSessionID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "desktopvirtualization.UserSessionsClient", "Disconnect", nil, "Failure preparing request")
		return
	}

	resp, err := client.DisconnectSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "desktopvirtualization.UserSessionsClient", "Disconnect", resp, "Failure sending request")
		return
	}

	result, err = client.DisconnectResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "desktopvirtualization.UserSessionsClient", "Disconnect", resp, "Failure responding to request")
		return
	}

	return
}

// DisconnectPreparer prepares the Disconnect request.
func (client UserSessionsClient) DisconnectPreparer(ctx context.Context, resourceGroupName string, hostPoolName string, sessionHostName string, userSessionID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"hostPoolName":      autorest.Encode("path", hostPoolName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"sessionHostName":   autorest.Encode("path", sessionHostName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"userSessionId":     autorest.Encode("path", userSessionID),
	}

	const APIVersion = "2019-01-23-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DesktopVirtualization/hostPools/{hostPoolName}/sessionHosts/{sessionHostName}/userSessions/{userSessionId}/disconnect", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DisconnectSender sends the Disconnect request. The method will close the
// http.Response Body if it receives an error.
func (client UserSessionsClient) DisconnectSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// DisconnectResponder handles the response to the Disconnect request. The method always
// closes the http.Response Body.
func (client UserSessionsClient) DisconnectResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get get a userSession.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// hostPoolName - the name of the host pool within the specified resource group
// sessionHostName - the name of the session host within the specified host pool
// userSessionID - the name of the user session within the specified session host
func (client UserSessionsClient) Get(ctx context.Context, resourceGroupName string, hostPoolName string, sessionHostName string, userSessionID string) (result UserSession, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/UserSessionsClient.Get")
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
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: hostPoolName,
			Constraints: []validation.Constraint{{Target: "hostPoolName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "hostPoolName", Name: validation.MinLength, Rule: 3, Chain: nil}}},
		{TargetValue: sessionHostName,
			Constraints: []validation.Constraint{{Target: "sessionHostName", Name: validation.MaxLength, Rule: 48, Chain: nil},
				{Target: "sessionHostName", Name: validation.MinLength, Rule: 3, Chain: nil}}},
		{TargetValue: userSessionID,
			Constraints: []validation.Constraint{{Target: "userSessionID", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "userSessionID", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("desktopvirtualization.UserSessionsClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, hostPoolName, sessionHostName, userSessionID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "desktopvirtualization.UserSessionsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "desktopvirtualization.UserSessionsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "desktopvirtualization.UserSessionsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client UserSessionsClient) GetPreparer(ctx context.Context, resourceGroupName string, hostPoolName string, sessionHostName string, userSessionID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"hostPoolName":      autorest.Encode("path", hostPoolName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"sessionHostName":   autorest.Encode("path", sessionHostName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"userSessionId":     autorest.Encode("path", userSessionID),
	}

	const APIVersion = "2019-01-23-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DesktopVirtualization/hostPools/{hostPoolName}/sessionHosts/{sessionHostName}/userSessions/{userSessionId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client UserSessionsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client UserSessionsClient) GetResponder(resp *http.Response) (result UserSession, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List list userSessions.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// hostPoolName - the name of the host pool within the specified resource group
// sessionHostName - the name of the session host within the specified host pool
func (client UserSessionsClient) List(ctx context.Context, resourceGroupName string, hostPoolName string, sessionHostName string) (result UserSessionListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/UserSessionsClient.List")
		defer func() {
			sc := -1
			if result.usl.Response.Response != nil {
				sc = result.usl.Response.Response.StatusCode
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
		{TargetValue: hostPoolName,
			Constraints: []validation.Constraint{{Target: "hostPoolName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "hostPoolName", Name: validation.MinLength, Rule: 3, Chain: nil}}},
		{TargetValue: sessionHostName,
			Constraints: []validation.Constraint{{Target: "sessionHostName", Name: validation.MaxLength, Rule: 48, Chain: nil},
				{Target: "sessionHostName", Name: validation.MinLength, Rule: 3, Chain: nil}}}}); err != nil {
		return result, validation.NewError("desktopvirtualization.UserSessionsClient", "List", err.Error())
	}

	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName, hostPoolName, sessionHostName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "desktopvirtualization.UserSessionsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.usl.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "desktopvirtualization.UserSessionsClient", "List", resp, "Failure sending request")
		return
	}

	result.usl, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "desktopvirtualization.UserSessionsClient", "List", resp, "Failure responding to request")
		return
	}
	if result.usl.hasNextLink() && result.usl.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client UserSessionsClient) ListPreparer(ctx context.Context, resourceGroupName string, hostPoolName string, sessionHostName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"hostPoolName":      autorest.Encode("path", hostPoolName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"sessionHostName":   autorest.Encode("path", sessionHostName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-01-23-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DesktopVirtualization/hostPools/{hostPoolName}/sessionHosts/{sessionHostName}/userSessions", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client UserSessionsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client UserSessionsClient) ListResponder(resp *http.Response) (result UserSessionList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client UserSessionsClient) listNextResults(ctx context.Context, lastResults UserSessionList) (result UserSessionList, err error) {
	req, err := lastResults.userSessionListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "desktopvirtualization.UserSessionsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "desktopvirtualization.UserSessionsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "desktopvirtualization.UserSessionsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client UserSessionsClient) ListComplete(ctx context.Context, resourceGroupName string, hostPoolName string, sessionHostName string) (result UserSessionListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/UserSessionsClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceGroupName, hostPoolName, sessionHostName)
	return
}

// ListByHostPool list userSessions.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// hostPoolName - the name of the host pool within the specified resource group
// filter - oData filter expression. Valid properties for filtering are userprincipalname and sessionstate.
func (client UserSessionsClient) ListByHostPool(ctx context.Context, resourceGroupName string, hostPoolName string, filter string) (result UserSessionListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/UserSessionsClient.ListByHostPool")
		defer func() {
			sc := -1
			if result.usl.Response.Response != nil {
				sc = result.usl.Response.Response.StatusCode
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
		{TargetValue: hostPoolName,
			Constraints: []validation.Constraint{{Target: "hostPoolName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "hostPoolName", Name: validation.MinLength, Rule: 3, Chain: nil}}}}); err != nil {
		return result, validation.NewError("desktopvirtualization.UserSessionsClient", "ListByHostPool", err.Error())
	}

	result.fn = client.listByHostPoolNextResults
	req, err := client.ListByHostPoolPreparer(ctx, resourceGroupName, hostPoolName, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "desktopvirtualization.UserSessionsClient", "ListByHostPool", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByHostPoolSender(req)
	if err != nil {
		result.usl.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "desktopvirtualization.UserSessionsClient", "ListByHostPool", resp, "Failure sending request")
		return
	}

	result.usl, err = client.ListByHostPoolResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "desktopvirtualization.UserSessionsClient", "ListByHostPool", resp, "Failure responding to request")
		return
	}
	if result.usl.hasNextLink() && result.usl.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByHostPoolPreparer prepares the ListByHostPool request.
func (client UserSessionsClient) ListByHostPoolPreparer(ctx context.Context, resourceGroupName string, hostPoolName string, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"hostPoolName":      autorest.Encode("path", hostPoolName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-01-23-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DesktopVirtualization/hostPools/{hostPoolName}/userSessions", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByHostPoolSender sends the ListByHostPool request. The method will close the
// http.Response Body if it receives an error.
func (client UserSessionsClient) ListByHostPoolSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByHostPoolResponder handles the response to the ListByHostPool request. The method always
// closes the http.Response Body.
func (client UserSessionsClient) ListByHostPoolResponder(resp *http.Response) (result UserSessionList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByHostPoolNextResults retrieves the next set of results, if any.
func (client UserSessionsClient) listByHostPoolNextResults(ctx context.Context, lastResults UserSessionList) (result UserSessionList, err error) {
	req, err := lastResults.userSessionListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "desktopvirtualization.UserSessionsClient", "listByHostPoolNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByHostPoolSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "desktopvirtualization.UserSessionsClient", "listByHostPoolNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByHostPoolResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "desktopvirtualization.UserSessionsClient", "listByHostPoolNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByHostPoolComplete enumerates all values, automatically crossing page boundaries as required.
func (client UserSessionsClient) ListByHostPoolComplete(ctx context.Context, resourceGroupName string, hostPoolName string, filter string) (result UserSessionListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/UserSessionsClient.ListByHostPool")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByHostPool(ctx, resourceGroupName, hostPoolName, filter)
	return
}

// SendMessageMethod send a message to a user.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// hostPoolName - the name of the host pool within the specified resource group
// sessionHostName - the name of the session host within the specified host pool
// userSessionID - the name of the user session within the specified session host
// sendMessage - object containing message includes title and message body
func (client UserSessionsClient) SendMessageMethod(ctx context.Context, resourceGroupName string, hostPoolName string, sessionHostName string, userSessionID string, sendMessage *SendMessage) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/UserSessionsClient.SendMessageMethod")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
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
		{TargetValue: hostPoolName,
			Constraints: []validation.Constraint{{Target: "hostPoolName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "hostPoolName", Name: validation.MinLength, Rule: 3, Chain: nil}}},
		{TargetValue: sessionHostName,
			Constraints: []validation.Constraint{{Target: "sessionHostName", Name: validation.MaxLength, Rule: 48, Chain: nil},
				{Target: "sessionHostName", Name: validation.MinLength, Rule: 3, Chain: nil}}},
		{TargetValue: userSessionID,
			Constraints: []validation.Constraint{{Target: "userSessionID", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "userSessionID", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("desktopvirtualization.UserSessionsClient", "SendMessageMethod", err.Error())
	}

	req, err := client.SendMessageMethodPreparer(ctx, resourceGroupName, hostPoolName, sessionHostName, userSessionID, sendMessage)
	if err != nil {
		err = autorest.NewErrorWithError(err, "desktopvirtualization.UserSessionsClient", "SendMessageMethod", nil, "Failure preparing request")
		return
	}

	resp, err := client.SendMessageMethodSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "desktopvirtualization.UserSessionsClient", "SendMessageMethod", resp, "Failure sending request")
		return
	}

	result, err = client.SendMessageMethodResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "desktopvirtualization.UserSessionsClient", "SendMessageMethod", resp, "Failure responding to request")
		return
	}

	return
}

// SendMessageMethodPreparer prepares the SendMessageMethod request.
func (client UserSessionsClient) SendMessageMethodPreparer(ctx context.Context, resourceGroupName string, hostPoolName string, sessionHostName string, userSessionID string, sendMessage *SendMessage) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"hostPoolName":      autorest.Encode("path", hostPoolName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"sessionHostName":   autorest.Encode("path", sessionHostName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"userSessionId":     autorest.Encode("path", userSessionID),
	}

	const APIVersion = "2019-01-23-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DesktopVirtualization/hostPools/{hostPoolName}/sessionHosts/{sessionHostName}/userSessions/{userSessionId}/sendMessage", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if sendMessage != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithJSON(sendMessage))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// SendMessageMethodSender sends the SendMessageMethod request. The method will close the
// http.Response Body if it receives an error.
func (client UserSessionsClient) SendMessageMethodSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// SendMessageMethodResponder handles the response to the SendMessageMethod request. The method always
// closes the http.Response Body.
func (client UserSessionsClient) SendMessageMethodResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}
