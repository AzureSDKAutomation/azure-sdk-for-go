package redis

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

// FirewallRuleClient is the REST API for Azure Redis Cache Service.
type FirewallRuleClient struct {
	BaseClient
}

// NewFirewallRuleClient creates an instance of the FirewallRuleClient client.
func NewFirewallRuleClient(subscriptionID string) FirewallRuleClient {
	return NewFirewallRuleClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewFirewallRuleClientWithBaseURI creates an instance of the FirewallRuleClient client using a custom endpoint.  Use
// this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewFirewallRuleClientWithBaseURI(baseURI string, subscriptionID string) FirewallRuleClient {
	return FirewallRuleClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate create or update a redis cache firewall rule
// Parameters:
// resourceGroupName - the name of the resource group.
// cacheName - the name of the Redis cache.
// ruleName - the name of the firewall rule.
// parameters - parameters supplied to the create or update redis firewall rule operation.
func (client FirewallRuleClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, cacheName string, ruleName string, parameters FirewallRule) (result FirewallRule, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/FirewallRuleClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.FirewallRuleProperties", Name: validation.Null, Rule: true,
				Chain: []validation.Constraint{{Target: "parameters.FirewallRuleProperties.StartIP", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "parameters.FirewallRuleProperties.EndIP", Name: validation.Null, Rule: true, Chain: nil},
				}}}}}); err != nil {
		return result, validation.NewError("redis.FirewallRuleClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, cacheName, ruleName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "redis.FirewallRuleClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "redis.FirewallRuleClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "redis.FirewallRuleClient", "CreateOrUpdate", resp, "Failure responding to request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client FirewallRuleClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, cacheName string, ruleName string, parameters FirewallRule) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"cacheName":         autorest.Encode("path", cacheName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"ruleName":          autorest.Encode("path", ruleName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	parameters.ID = nil
	parameters.Name = nil
	parameters.Type = nil
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/Redis/{cacheName}/firewallRules/{ruleName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client FirewallRuleClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client FirewallRuleClient) CreateOrUpdateResponder(resp *http.Response) (result FirewallRule, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a single firewall rule in a specified redis cache.
// Parameters:
// resourceGroupName - the name of the resource group.
// cacheName - the name of the Redis cache.
// ruleName - the name of the firewall rule.
func (client FirewallRuleClient) Delete(ctx context.Context, resourceGroupName string, cacheName string, ruleName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/FirewallRuleClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, cacheName, ruleName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "redis.FirewallRuleClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "redis.FirewallRuleClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "redis.FirewallRuleClient", "Delete", resp, "Failure responding to request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client FirewallRuleClient) DeletePreparer(ctx context.Context, resourceGroupName string, cacheName string, ruleName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"cacheName":         autorest.Encode("path", cacheName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"ruleName":          autorest.Encode("path", ruleName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/Redis/{cacheName}/firewallRules/{ruleName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client FirewallRuleClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client FirewallRuleClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets a single firewall rule in a specified redis cache.
// Parameters:
// resourceGroupName - the name of the resource group.
// cacheName - the name of the Redis cache.
// ruleName - the name of the firewall rule.
func (client FirewallRuleClient) Get(ctx context.Context, resourceGroupName string, cacheName string, ruleName string) (result FirewallRule, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/FirewallRuleClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, cacheName, ruleName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "redis.FirewallRuleClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "redis.FirewallRuleClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "redis.FirewallRuleClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client FirewallRuleClient) GetPreparer(ctx context.Context, resourceGroupName string, cacheName string, ruleName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"cacheName":         autorest.Encode("path", cacheName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"ruleName":          autorest.Encode("path", ruleName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/Redis/{cacheName}/firewallRules/{ruleName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client FirewallRuleClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client FirewallRuleClient) GetResponder(resp *http.Response) (result FirewallRule, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
