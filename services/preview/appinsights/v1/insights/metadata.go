package insights

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

// MetadataClient is the composite Swagger for Application Insights Data Client
type MetadataClient struct {
	BaseClient
}

// NewMetadataClient creates an instance of the MetadataClient client.
func NewMetadataClient() MetadataClient {
	return NewMetadataClientWithBaseURI(DefaultBaseURI)
}

// NewMetadataClientWithBaseURI creates an instance of the MetadataClient client using a custom endpoint.  Use this
// when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewMetadataClientWithBaseURI(baseURI string) MetadataClient {
	return MetadataClient{NewWithBaseURI(baseURI)}
}

// Get retrieve the metadata information for the app, including its schema, etc.
// Parameters:
// appID - ID of the application. This is Application ID from the API Access settings blade in the Azure
// portal.
func (client MetadataClient) Get(ctx context.Context, appID string) (result MetadataResults, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MetadataClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, appID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insights.MetadataClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "insights.MetadataClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insights.MetadataClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client MetadataClient) GetPreparer(ctx context.Context, appID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"appId": autorest.Encode("path", appID),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/apps/{appId}/metadata", pathParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client MetadataClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client MetadataClient) GetResponder(resp *http.Response) (result MetadataResults, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Post retrieve the metadata information for the app, including its schema, etc.
// Parameters:
// appID - ID of the application. This is Application ID from the API Access settings blade in the Azure
// portal.
func (client MetadataClient) Post(ctx context.Context, appID string) (result MetadataResults, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MetadataClient.Post")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.PostPreparer(ctx, appID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insights.MetadataClient", "Post", nil, "Failure preparing request")
		return
	}

	resp, err := client.PostSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "insights.MetadataClient", "Post", resp, "Failure sending request")
		return
	}

	result, err = client.PostResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insights.MetadataClient", "Post", resp, "Failure responding to request")
		return
	}

	return
}

// PostPreparer prepares the Post request.
func (client MetadataClient) PostPreparer(ctx context.Context, appID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"appId": autorest.Encode("path", appID),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/apps/{appId}/metadata", pathParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// PostSender sends the Post request. The method will close the
// http.Response Body if it receives an error.
func (client MetadataClient) PostSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// PostResponder handles the response to the Post request. The method always
// closes the http.Response Body.
func (client MetadataClient) PostResponder(resp *http.Response) (result MetadataResults, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
