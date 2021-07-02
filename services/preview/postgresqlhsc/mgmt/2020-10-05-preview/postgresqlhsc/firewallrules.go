package postgresqlhsc

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

// FirewallRulesClient is the the Microsoft Azure management API provides create, read, update, and delete
// functionality for Azure PostgreSQL Hyperscale (Citus) resources including server groups, servers, databases,
// firewall rules, VNET rules, security alert policies, log files and configurations.
type FirewallRulesClient struct {
	BaseClient
}

// NewFirewallRulesClient creates an instance of the FirewallRulesClient client.
func NewFirewallRulesClient(subscriptionID string) FirewallRulesClient {
	return NewFirewallRulesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewFirewallRulesClientWithBaseURI creates an instance of the FirewallRulesClient client using a custom endpoint.
// Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewFirewallRulesClientWithBaseURI(baseURI string, subscriptionID string) FirewallRulesClient {
	return FirewallRulesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates a new firewall rule or updates an existing firewall rule.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// serverGroupName - the name of the server group.
// firewallRuleName - the name of the server group firewall rule.
// parameters - the required parameters for creating or updating a firewall rule.
func (client FirewallRulesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serverGroupName string, firewallRuleName string, parameters FirewallRule) (result FirewallRulesCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/FirewallRulesClient.CreateOrUpdate")
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
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: serverGroupName,
			Constraints: []validation.Constraint{{Target: "serverGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "serverGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.FirewallRuleProperties", Name: validation.Null, Rule: true,
				Chain: []validation.Constraint{{Target: "parameters.FirewallRuleProperties.StartIPAddress", Name: validation.Null, Rule: true,
					Chain: []validation.Constraint{{Target: "parameters.FirewallRuleProperties.StartIPAddress", Name: validation.Pattern, Rule: `^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$`, Chain: nil}}},
					{Target: "parameters.FirewallRuleProperties.EndIPAddress", Name: validation.Null, Rule: true,
						Chain: []validation.Constraint{{Target: "parameters.FirewallRuleProperties.EndIPAddress", Name: validation.Pattern, Rule: `^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$`, Chain: nil}}},
				}}}}}); err != nil {
		return result, validation.NewError("postgresqlhsc.FirewallRulesClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, serverGroupName, firewallRuleName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "postgresqlhsc.FirewallRulesClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "postgresqlhsc.FirewallRulesClient", "CreateOrUpdate", nil, "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client FirewallRulesClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, serverGroupName string, firewallRuleName string, parameters FirewallRule) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"firewallRuleName":  autorest.Encode("path", firewallRuleName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverGroupName":   autorest.Encode("path", serverGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-10-05-privatepreview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBForPostgreSql/serverGroupsv2/{serverGroupName}/firewallRules/{firewallRuleName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client FirewallRulesClient) CreateOrUpdateSender(req *http.Request) (future FirewallRulesCreateOrUpdateFuture, err error) {
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
func (client FirewallRulesClient) CreateOrUpdateResponder(resp *http.Response) (result FirewallRule, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a server group firewall rule.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// serverGroupName - the name of the server group.
// firewallRuleName - the name of the server group firewall rule.
func (client FirewallRulesClient) Delete(ctx context.Context, resourceGroupName string, serverGroupName string, firewallRuleName string) (result FirewallRulesDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/FirewallRulesClient.Delete")
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
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: serverGroupName,
			Constraints: []validation.Constraint{{Target: "serverGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "serverGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("postgresqlhsc.FirewallRulesClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, serverGroupName, firewallRuleName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "postgresqlhsc.FirewallRulesClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "postgresqlhsc.FirewallRulesClient", "Delete", nil, "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client FirewallRulesClient) DeletePreparer(ctx context.Context, resourceGroupName string, serverGroupName string, firewallRuleName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"firewallRuleName":  autorest.Encode("path", firewallRuleName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverGroupName":   autorest.Encode("path", serverGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-10-05-privatepreview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBForPostgreSql/serverGroupsv2/{serverGroupName}/firewallRules/{firewallRuleName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client FirewallRulesClient) DeleteSender(req *http.Request) (future FirewallRulesDeleteFuture, err error) {
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
func (client FirewallRulesClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets information about a server group firewall rule.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// serverGroupName - the name of the server group.
// firewallRuleName - the name of the server group firewall rule.
func (client FirewallRulesClient) Get(ctx context.Context, resourceGroupName string, serverGroupName string, firewallRuleName string) (result FirewallRule, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/FirewallRulesClient.Get")
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
		{TargetValue: serverGroupName,
			Constraints: []validation.Constraint{{Target: "serverGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "serverGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("postgresqlhsc.FirewallRulesClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, serverGroupName, firewallRuleName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "postgresqlhsc.FirewallRulesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "postgresqlhsc.FirewallRulesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "postgresqlhsc.FirewallRulesClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client FirewallRulesClient) GetPreparer(ctx context.Context, resourceGroupName string, serverGroupName string, firewallRuleName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"firewallRuleName":  autorest.Encode("path", firewallRuleName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverGroupName":   autorest.Encode("path", serverGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-10-05-privatepreview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBForPostgreSql/serverGroupsv2/{serverGroupName}/firewallRules/{firewallRuleName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client FirewallRulesClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client FirewallRulesClient) GetResponder(resp *http.Response) (result FirewallRule, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByServerGroup list all the firewall rules in a given server group.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// serverGroupName - the name of the server group.
func (client FirewallRulesClient) ListByServerGroup(ctx context.Context, resourceGroupName string, serverGroupName string) (result FirewallRuleListResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/FirewallRulesClient.ListByServerGroup")
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
		{TargetValue: serverGroupName,
			Constraints: []validation.Constraint{{Target: "serverGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "serverGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("postgresqlhsc.FirewallRulesClient", "ListByServerGroup", err.Error())
	}

	req, err := client.ListByServerGroupPreparer(ctx, resourceGroupName, serverGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "postgresqlhsc.FirewallRulesClient", "ListByServerGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByServerGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "postgresqlhsc.FirewallRulesClient", "ListByServerGroup", resp, "Failure sending request")
		return
	}

	result, err = client.ListByServerGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "postgresqlhsc.FirewallRulesClient", "ListByServerGroup", resp, "Failure responding to request")
		return
	}

	return
}

// ListByServerGroupPreparer prepares the ListByServerGroup request.
func (client FirewallRulesClient) ListByServerGroupPreparer(ctx context.Context, resourceGroupName string, serverGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverGroupName":   autorest.Encode("path", serverGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-10-05-privatepreview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBForPostgreSql/serverGroupsv2/{serverGroupName}/firewallRules", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByServerGroupSender sends the ListByServerGroup request. The method will close the
// http.Response Body if it receives an error.
func (client FirewallRulesClient) ListByServerGroupSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByServerGroupResponder handles the response to the ListByServerGroup request. The method always
// closes the http.Response Body.
func (client FirewallRulesClient) ListByServerGroupResponder(resp *http.Response) (result FirewallRuleListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
