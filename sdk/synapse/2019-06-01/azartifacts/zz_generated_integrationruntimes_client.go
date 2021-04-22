// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azartifacts

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
)

// IntegrationRuntimesClient contains the methods for the IntegrationRuntimes group.
// Don't use this type directly, use NewIntegrationRuntimesClient() instead.
type IntegrationRuntimesClient struct {
	con *Connection
}

// NewIntegrationRuntimesClient creates a new instance of IntegrationRuntimesClient with the specified values.
func NewIntegrationRuntimesClient(con *Connection) *IntegrationRuntimesClient {
	return &IntegrationRuntimesClient{con: con}
}

// Get - Get Integration Runtime
func (client *IntegrationRuntimesClient) Get(ctx context.Context, integrationRuntimeName string, options *IntegrationRuntimesGetOptions) (IntegrationRuntimeResourceResponse, error) {
	req, err := client.getCreateRequest(ctx, integrationRuntimeName, options)
	if err != nil {
		return IntegrationRuntimeResourceResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return IntegrationRuntimeResourceResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return IntegrationRuntimeResourceResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *IntegrationRuntimesClient) getCreateRequest(ctx context.Context, integrationRuntimeName string, options *IntegrationRuntimesGetOptions) (*azcore.Request, error) {
	urlPath := "/integrationRuntimes/{integrationRuntimeName}"
	if integrationRuntimeName == "" {
		return nil, errors.New("parameter integrationRuntimeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{integrationRuntimeName}", url.PathEscape(integrationRuntimeName))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *IntegrationRuntimesClient) getHandleResponse(resp *azcore.Response) (IntegrationRuntimeResourceResponse, error) {
	var val *IntegrationRuntimeResource
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return IntegrationRuntimeResourceResponse{}, err
	}
	return IntegrationRuntimeResourceResponse{RawResponse: resp.Response, IntegrationRuntimeResource: val}, nil
}

// getHandleError handles the Get error response.
func (client *IntegrationRuntimesClient) getHandleError(resp *azcore.Response) error {
	var err ErrorContract
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// List - List Integration Runtimes
func (client *IntegrationRuntimesClient) List(ctx context.Context, options *IntegrationRuntimesListOptions) (IntegrationRuntimeListResponseResponse, error) {
	req, err := client.listCreateRequest(ctx, options)
	if err != nil {
		return IntegrationRuntimeListResponseResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return IntegrationRuntimeListResponseResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return IntegrationRuntimeListResponseResponse{}, client.listHandleError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *IntegrationRuntimesClient) listCreateRequest(ctx context.Context, options *IntegrationRuntimesListOptions) (*azcore.Request, error) {
	urlPath := "/integrationRuntimes"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *IntegrationRuntimesClient) listHandleResponse(resp *azcore.Response) (IntegrationRuntimeListResponseResponse, error) {
	var val *IntegrationRuntimeListResponse
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return IntegrationRuntimeListResponseResponse{}, err
	}
	return IntegrationRuntimeListResponseResponse{RawResponse: resp.Response, IntegrationRuntimeListResponse: val}, nil
}

// listHandleError handles the List error response.
func (client *IntegrationRuntimesClient) listHandleError(resp *azcore.Response) error {
	var err ErrorContract
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
