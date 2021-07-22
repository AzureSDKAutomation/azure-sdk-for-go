package features

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

// SubscriptionFeatureRegistrationsClient is the client for the SubscriptionFeatureRegistrations methods of the
// Features service.
type SubscriptionFeatureRegistrationsClient struct {
	BaseClient
}

// NewSubscriptionFeatureRegistrationsClient creates an instance of the SubscriptionFeatureRegistrationsClient client.
func NewSubscriptionFeatureRegistrationsClient(subscriptionID string) SubscriptionFeatureRegistrationsClient {
	return NewSubscriptionFeatureRegistrationsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewSubscriptionFeatureRegistrationsClientWithBaseURI creates an instance of the
// SubscriptionFeatureRegistrationsClient client using a custom endpoint.  Use this when interacting with an Azure
// cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewSubscriptionFeatureRegistrationsClientWithBaseURI(baseURI string, subscriptionID string) SubscriptionFeatureRegistrationsClient {
	return SubscriptionFeatureRegistrationsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate create or update a feature registration.
// Parameters:
// APIVersion - the API version to use for this operation.
// providerNamespace - the provider namespace.
// featureName - the feature name.
// subscriptionFeatureRegistrationType - subscription Feature Registration Type details.
func (client SubscriptionFeatureRegistrationsClient) CreateOrUpdate(ctx context.Context, APIVersion string, providerNamespace string, featureName string, subscriptionFeatureRegistrationType *SubscriptionFeatureRegistration) (result SubscriptionFeatureRegistration, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SubscriptionFeatureRegistrationsClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: subscriptionFeatureRegistrationType,
			Constraints: []validation.Constraint{{Target: "subscriptionFeatureRegistrationType", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "subscriptionFeatureRegistrationType.Properties", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "subscriptionFeatureRegistrationType.Properties.DocumentationLink", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "subscriptionFeatureRegistrationType.Properties.DocumentationLink", Name: validation.MaxLength, Rule: 1000, Chain: nil}}},
						{Target: "subscriptionFeatureRegistrationType.Properties.Description", Name: validation.Null, Rule: false,
							Chain: []validation.Constraint{{Target: "subscriptionFeatureRegistrationType.Properties.Description", Name: validation.MaxLength, Rule: 1000, Chain: nil}}},
					}},
				}}}}}); err != nil {
		return result, validation.NewError("features.SubscriptionFeatureRegistrationsClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, APIVersion, providerNamespace, featureName, subscriptionFeatureRegistrationType)
	if err != nil {
		err = autorest.NewErrorWithError(err, "features.SubscriptionFeatureRegistrationsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "features.SubscriptionFeatureRegistrationsClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "features.SubscriptionFeatureRegistrationsClient", "CreateOrUpdate", resp, "Failure responding to request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client SubscriptionFeatureRegistrationsClient) CreateOrUpdatePreparer(ctx context.Context, APIVersion string, providerNamespace string, featureName string, subscriptionFeatureRegistrationType *SubscriptionFeatureRegistration) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"featureName":       autorest.Encode("path", featureName),
		"providerNamespace": autorest.Encode("path", providerNamespace),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Features/featureProviders/{providerNamespace}/subscriptionFeatureRegistrations/{featureName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if subscriptionFeatureRegistrationType != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithJSON(subscriptionFeatureRegistrationType))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client SubscriptionFeatureRegistrationsClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client SubscriptionFeatureRegistrationsClient) CreateOrUpdateResponder(resp *http.Response) (result SubscriptionFeatureRegistration, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a feature registration
// Parameters:
// APIVersion - the API version to use for this operation.
// providerNamespace - the provider namespace.
// featureName - the feature name.
func (client SubscriptionFeatureRegistrationsClient) Delete(ctx context.Context, APIVersion string, providerNamespace string, featureName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SubscriptionFeatureRegistrationsClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, APIVersion, providerNamespace, featureName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "features.SubscriptionFeatureRegistrationsClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "features.SubscriptionFeatureRegistrationsClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "features.SubscriptionFeatureRegistrationsClient", "Delete", resp, "Failure responding to request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client SubscriptionFeatureRegistrationsClient) DeletePreparer(ctx context.Context, APIVersion string, providerNamespace string, featureName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"featureName":       autorest.Encode("path", featureName),
		"providerNamespace": autorest.Encode("path", providerNamespace),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Features/featureProviders/{providerNamespace}/subscriptionFeatureRegistrations/{featureName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client SubscriptionFeatureRegistrationsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client SubscriptionFeatureRegistrationsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get returns a feature registration
// Parameters:
// APIVersion - the API version to use for this operation.
// providerNamespace - the provider namespace.
// featureName - the feature name.
func (client SubscriptionFeatureRegistrationsClient) Get(ctx context.Context, APIVersion string, providerNamespace string, featureName string) (result SubscriptionFeatureRegistration, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SubscriptionFeatureRegistrationsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, APIVersion, providerNamespace, featureName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "features.SubscriptionFeatureRegistrationsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "features.SubscriptionFeatureRegistrationsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "features.SubscriptionFeatureRegistrationsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client SubscriptionFeatureRegistrationsClient) GetPreparer(ctx context.Context, APIVersion string, providerNamespace string, featureName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"featureName":       autorest.Encode("path", featureName),
		"providerNamespace": autorest.Encode("path", providerNamespace),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Features/featureProviders/{providerNamespace}/subscriptionFeatureRegistrations/{featureName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client SubscriptionFeatureRegistrationsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client SubscriptionFeatureRegistrationsClient) GetResponder(resp *http.Response) (result SubscriptionFeatureRegistration, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListAllBySubscription returns subscription feature registrations for given subscription.
// Parameters:
// APIVersion - the API version to use for this operation.
func (client SubscriptionFeatureRegistrationsClient) ListAllBySubscription(ctx context.Context, APIVersion string) (result SubscriptionFeatureRegistrationListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SubscriptionFeatureRegistrationsClient.ListAllBySubscription")
		defer func() {
			sc := -1
			if result.sfrl.Response.Response != nil {
				sc = result.sfrl.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listAllBySubscriptionNextResults
	req, err := client.ListAllBySubscriptionPreparer(ctx, APIVersion)
	if err != nil {
		err = autorest.NewErrorWithError(err, "features.SubscriptionFeatureRegistrationsClient", "ListAllBySubscription", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListAllBySubscriptionSender(req)
	if err != nil {
		result.sfrl.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "features.SubscriptionFeatureRegistrationsClient", "ListAllBySubscription", resp, "Failure sending request")
		return
	}

	result.sfrl, err = client.ListAllBySubscriptionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "features.SubscriptionFeatureRegistrationsClient", "ListAllBySubscription", resp, "Failure responding to request")
		return
	}
	if result.sfrl.hasNextLink() && result.sfrl.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListAllBySubscriptionPreparer prepares the ListAllBySubscription request.
func (client SubscriptionFeatureRegistrationsClient) ListAllBySubscriptionPreparer(ctx context.Context, APIVersion string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Features/subscriptionFeatureRegistrations", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListAllBySubscriptionSender sends the ListAllBySubscription request. The method will close the
// http.Response Body if it receives an error.
func (client SubscriptionFeatureRegistrationsClient) ListAllBySubscriptionSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListAllBySubscriptionResponder handles the response to the ListAllBySubscription request. The method always
// closes the http.Response Body.
func (client SubscriptionFeatureRegistrationsClient) ListAllBySubscriptionResponder(resp *http.Response) (result SubscriptionFeatureRegistrationList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listAllBySubscriptionNextResults retrieves the next set of results, if any.
func (client SubscriptionFeatureRegistrationsClient) listAllBySubscriptionNextResults(ctx context.Context, lastResults SubscriptionFeatureRegistrationList) (result SubscriptionFeatureRegistrationList, err error) {
	req, err := lastResults.subscriptionFeatureRegistrationListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "features.SubscriptionFeatureRegistrationsClient", "listAllBySubscriptionNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListAllBySubscriptionSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "features.SubscriptionFeatureRegistrationsClient", "listAllBySubscriptionNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListAllBySubscriptionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "features.SubscriptionFeatureRegistrationsClient", "listAllBySubscriptionNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListAllBySubscriptionComplete enumerates all values, automatically crossing page boundaries as required.
func (client SubscriptionFeatureRegistrationsClient) ListAllBySubscriptionComplete(ctx context.Context, APIVersion string) (result SubscriptionFeatureRegistrationListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SubscriptionFeatureRegistrationsClient.ListAllBySubscription")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListAllBySubscription(ctx, APIVersion)
	return
}

// ListBySubscription returns subscription feature registrations for given subscription and provider namespace.
// Parameters:
// APIVersion - the API version to use for this operation.
// providerNamespace - the provider namespace.
func (client SubscriptionFeatureRegistrationsClient) ListBySubscription(ctx context.Context, APIVersion string, providerNamespace string) (result SubscriptionFeatureRegistrationListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SubscriptionFeatureRegistrationsClient.ListBySubscription")
		defer func() {
			sc := -1
			if result.sfrl.Response.Response != nil {
				sc = result.sfrl.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listBySubscriptionNextResults
	req, err := client.ListBySubscriptionPreparer(ctx, APIVersion, providerNamespace)
	if err != nil {
		err = autorest.NewErrorWithError(err, "features.SubscriptionFeatureRegistrationsClient", "ListBySubscription", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListBySubscriptionSender(req)
	if err != nil {
		result.sfrl.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "features.SubscriptionFeatureRegistrationsClient", "ListBySubscription", resp, "Failure sending request")
		return
	}

	result.sfrl, err = client.ListBySubscriptionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "features.SubscriptionFeatureRegistrationsClient", "ListBySubscription", resp, "Failure responding to request")
		return
	}
	if result.sfrl.hasNextLink() && result.sfrl.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListBySubscriptionPreparer prepares the ListBySubscription request.
func (client SubscriptionFeatureRegistrationsClient) ListBySubscriptionPreparer(ctx context.Context, APIVersion string, providerNamespace string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"providerNamespace": autorest.Encode("path", providerNamespace),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Features/featureProviders/{providerNamespace}/subscriptionFeatureRegistrations", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListBySubscriptionSender sends the ListBySubscription request. The method will close the
// http.Response Body if it receives an error.
func (client SubscriptionFeatureRegistrationsClient) ListBySubscriptionSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListBySubscriptionResponder handles the response to the ListBySubscription request. The method always
// closes the http.Response Body.
func (client SubscriptionFeatureRegistrationsClient) ListBySubscriptionResponder(resp *http.Response) (result SubscriptionFeatureRegistrationList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listBySubscriptionNextResults retrieves the next set of results, if any.
func (client SubscriptionFeatureRegistrationsClient) listBySubscriptionNextResults(ctx context.Context, lastResults SubscriptionFeatureRegistrationList) (result SubscriptionFeatureRegistrationList, err error) {
	req, err := lastResults.subscriptionFeatureRegistrationListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "features.SubscriptionFeatureRegistrationsClient", "listBySubscriptionNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListBySubscriptionSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "features.SubscriptionFeatureRegistrationsClient", "listBySubscriptionNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListBySubscriptionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "features.SubscriptionFeatureRegistrationsClient", "listBySubscriptionNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListBySubscriptionComplete enumerates all values, automatically crossing page boundaries as required.
func (client SubscriptionFeatureRegistrationsClient) ListBySubscriptionComplete(ctx context.Context, APIVersion string, providerNamespace string) (result SubscriptionFeatureRegistrationListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SubscriptionFeatureRegistrationsClient.ListBySubscription")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListBySubscription(ctx, APIVersion, providerNamespace)
	return
}
