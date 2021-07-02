// Package mariadb implements the Azure ARM Mariadb service API version .
//
// The Microsoft Azure management API provides create, read, update, and delete functionality for Azure MariaDB
// resources including servers, databases, firewall rules, VNET rules, log files and configurations with new business
// model.
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

const (
	// DefaultBaseURI is the default URI used for the service Mariadb
	DefaultBaseURI = "https://management.azure.com"
)

// BaseClient is the base client for Mariadb.
type BaseClient struct {
	autorest.Client
	BaseURI        string
	SubscriptionID string
}

// New creates an instance of the BaseClient client.
func New(subscriptionID string) BaseClient {
	return NewWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewWithBaseURI creates an instance of the BaseClient client using a custom endpoint.  Use this when interacting with
// an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return BaseClient{
		Client:         autorest.NewClientWithUserAgent(UserAgent()),
		BaseURI:        baseURI,
		SubscriptionID: subscriptionID,
	}
}

// CreateRecommendedActionSession create recommendation action session for the advisor.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// serverName - the name of the server.
// advisorName - the advisor name for recommendation action.
// databaseName - the name of the database.
func (client BaseClient) CreateRecommendedActionSession(ctx context.Context, resourceGroupName string, serverName string, advisorName string, databaseName string) (result CreateRecommendedActionSessionFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BaseClient.CreateRecommendedActionSession")
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
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("mariadb.BaseClient", "CreateRecommendedActionSession", err.Error())
	}

	req, err := client.CreateRecommendedActionSessionPreparer(ctx, resourceGroupName, serverName, advisorName, databaseName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "mariadb.BaseClient", "CreateRecommendedActionSession", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateRecommendedActionSessionSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "mariadb.BaseClient", "CreateRecommendedActionSession", nil, "Failure sending request")
		return
	}

	return
}

// CreateRecommendedActionSessionPreparer prepares the CreateRecommendedActionSession request.
func (client BaseClient) CreateRecommendedActionSessionPreparer(ctx context.Context, resourceGroupName string, serverName string, advisorName string, databaseName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"advisorName":       autorest.Encode("path", advisorName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverName":        autorest.Encode("path", serverName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-06-01"
	queryParameters := map[string]interface{}{
		"api-version":  APIVersion,
		"databaseName": autorest.Encode("query", databaseName),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMariaDB/servers/{serverName}/advisors/{advisorName}/createRecommendedActionSession", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateRecommendedActionSessionSender sends the CreateRecommendedActionSession request. The method will close the
// http.Response Body if it receives an error.
func (client BaseClient) CreateRecommendedActionSessionSender(req *http.Request) (future CreateRecommendedActionSessionFuture, err error) {
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

// CreateRecommendedActionSessionResponder handles the response to the CreateRecommendedActionSession request. The method always
// closes the http.Response Body.
func (client BaseClient) CreateRecommendedActionSessionResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// ResetQueryPerformanceInsightData reset data for Query Performance Insight.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// serverName - the name of the server.
func (client BaseClient) ResetQueryPerformanceInsightData(ctx context.Context, resourceGroupName string, serverName string) (result QueryPerformanceInsightResetDataResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BaseClient.ResetQueryPerformanceInsightData")
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
		return result, validation.NewError("mariadb.BaseClient", "ResetQueryPerformanceInsightData", err.Error())
	}

	req, err := client.ResetQueryPerformanceInsightDataPreparer(ctx, resourceGroupName, serverName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "mariadb.BaseClient", "ResetQueryPerformanceInsightData", nil, "Failure preparing request")
		return
	}

	resp, err := client.ResetQueryPerformanceInsightDataSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "mariadb.BaseClient", "ResetQueryPerformanceInsightData", resp, "Failure sending request")
		return
	}

	result, err = client.ResetQueryPerformanceInsightDataResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "mariadb.BaseClient", "ResetQueryPerformanceInsightData", resp, "Failure responding to request")
		return
	}

	return
}

// ResetQueryPerformanceInsightDataPreparer prepares the ResetQueryPerformanceInsightData request.
func (client BaseClient) ResetQueryPerformanceInsightDataPreparer(ctx context.Context, resourceGroupName string, serverName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverName":        autorest.Encode("path", serverName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMariaDB/servers/{serverName}/resetQueryPerformanceInsightData", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ResetQueryPerformanceInsightDataSender sends the ResetQueryPerformanceInsightData request. The method will close the
// http.Response Body if it receives an error.
func (client BaseClient) ResetQueryPerformanceInsightDataSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ResetQueryPerformanceInsightDataResponder handles the response to the ResetQueryPerformanceInsightData request. The method always
// closes the http.Response Body.
func (client BaseClient) ResetQueryPerformanceInsightDataResponder(resp *http.Response) (result QueryPerformanceInsightResetDataResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
