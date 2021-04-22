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
	"time"
)

// LinkedServiceClient contains the methods for the LinkedService group.
// Don't use this type directly, use NewLinkedServiceClient() instead.
type LinkedServiceClient struct {
	con *Connection
}

// NewLinkedServiceClient creates a new instance of LinkedServiceClient with the specified values.
func NewLinkedServiceClient(con *Connection) *LinkedServiceClient {
	return &LinkedServiceClient{con: con}
}

// BeginCreateOrUpdateLinkedService - Creates or updates a linked service.
func (client *LinkedServiceClient) BeginCreateOrUpdateLinkedService(ctx context.Context, linkedServiceName string, linkedService LinkedServiceResource, options *LinkedServiceBeginCreateOrUpdateLinkedServiceOptions) (LinkedServiceResourcePollerResponse, error) {
	resp, err := client.createOrUpdateLinkedService(ctx, linkedServiceName, linkedService, options)
	if err != nil {
		return LinkedServiceResourcePollerResponse{}, err
	}
	result := LinkedServiceResourcePollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := azcore.NewLROPoller("LinkedServiceClient.CreateOrUpdateLinkedService", resp, client.con.Pipeline(), client.createOrUpdateLinkedServiceHandleError)
	if err != nil {
		return LinkedServiceResourcePollerResponse{}, err
	}
	poller := &linkedServiceResourcePoller{
		pt: pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (LinkedServiceResourceResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeCreateOrUpdateLinkedService creates a new LinkedServiceResourcePoller from the specified resume token.
// token - The value must come from a previous call to LinkedServiceResourcePoller.ResumeToken().
func (client *LinkedServiceClient) ResumeCreateOrUpdateLinkedService(token string) (LinkedServiceResourcePoller, error) {
	pt, err := azcore.NewLROPollerFromResumeToken("LinkedServiceClient.CreateOrUpdateLinkedService", token, client.con.Pipeline(), client.createOrUpdateLinkedServiceHandleError)
	if err != nil {
		return nil, err
	}
	return &linkedServiceResourcePoller{
		pt: pt,
	}, nil
}

// CreateOrUpdateLinkedService - Creates or updates a linked service.
func (client *LinkedServiceClient) createOrUpdateLinkedService(ctx context.Context, linkedServiceName string, linkedService LinkedServiceResource, options *LinkedServiceBeginCreateOrUpdateLinkedServiceOptions) (*azcore.Response, error) {
	req, err := client.createOrUpdateLinkedServiceCreateRequest(ctx, linkedServiceName, linkedService, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.createOrUpdateLinkedServiceHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateLinkedServiceCreateRequest creates the CreateOrUpdateLinkedService request.
func (client *LinkedServiceClient) createOrUpdateLinkedServiceCreateRequest(ctx context.Context, linkedServiceName string, linkedService LinkedServiceResource, options *LinkedServiceBeginCreateOrUpdateLinkedServiceOptions) (*azcore.Request, error) {
	urlPath := "/linkedservices/{linkedServiceName}"
	if linkedServiceName == "" {
		return nil, errors.New("parameter linkedServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{linkedServiceName}", url.PathEscape(linkedServiceName))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfMatch != nil {
		req.Header.Set("If-Match", *options.IfMatch)
	}
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(linkedService)
}

// createOrUpdateLinkedServiceHandleResponse handles the CreateOrUpdateLinkedService response.
func (client *LinkedServiceClient) createOrUpdateLinkedServiceHandleResponse(resp *azcore.Response) (LinkedServiceResourceResponse, error) {
	var val *LinkedServiceResource
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return LinkedServiceResourceResponse{}, err
	}
	return LinkedServiceResourceResponse{RawResponse: resp.Response, LinkedServiceResource: val}, nil
}

// createOrUpdateLinkedServiceHandleError handles the CreateOrUpdateLinkedService error response.
func (client *LinkedServiceClient) createOrUpdateLinkedServiceHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err.InnerError); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// BeginDeleteLinkedService - Deletes a linked service.
func (client *LinkedServiceClient) BeginDeleteLinkedService(ctx context.Context, linkedServiceName string, options *LinkedServiceBeginDeleteLinkedServiceOptions) (HTTPPollerResponse, error) {
	resp, err := client.deleteLinkedService(ctx, linkedServiceName, options)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	result := HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := azcore.NewLROPoller("LinkedServiceClient.DeleteLinkedService", resp, client.con.Pipeline(), client.deleteLinkedServiceHandleError)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	poller := &httpPoller{
		pt: pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*http.Response, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeDeleteLinkedService creates a new HTTPPoller from the specified resume token.
// token - The value must come from a previous call to HTTPPoller.ResumeToken().
func (client *LinkedServiceClient) ResumeDeleteLinkedService(token string) (HTTPPoller, error) {
	pt, err := azcore.NewLROPollerFromResumeToken("LinkedServiceClient.DeleteLinkedService", token, client.con.Pipeline(), client.deleteLinkedServiceHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pt: pt,
	}, nil
}

// DeleteLinkedService - Deletes a linked service.
func (client *LinkedServiceClient) deleteLinkedService(ctx context.Context, linkedServiceName string, options *LinkedServiceBeginDeleteLinkedServiceOptions) (*azcore.Response, error) {
	req, err := client.deleteLinkedServiceCreateRequest(ctx, linkedServiceName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteLinkedServiceHandleError(resp)
	}
	return resp, nil
}

// deleteLinkedServiceCreateRequest creates the DeleteLinkedService request.
func (client *LinkedServiceClient) deleteLinkedServiceCreateRequest(ctx context.Context, linkedServiceName string, options *LinkedServiceBeginDeleteLinkedServiceOptions) (*azcore.Request, error) {
	urlPath := "/linkedservices/{linkedServiceName}"
	if linkedServiceName == "" {
		return nil, errors.New("parameter linkedServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{linkedServiceName}", url.PathEscape(linkedServiceName))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.con.Endpoint(), urlPath))
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

// deleteLinkedServiceHandleError handles the DeleteLinkedService error response.
func (client *LinkedServiceClient) deleteLinkedServiceHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err.InnerError); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetLinkedService - Gets a linked service.
func (client *LinkedServiceClient) GetLinkedService(ctx context.Context, linkedServiceName string, options *LinkedServiceGetLinkedServiceOptions) (LinkedServiceResourceResponse, error) {
	req, err := client.getLinkedServiceCreateRequest(ctx, linkedServiceName, options)
	if err != nil {
		return LinkedServiceResourceResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return LinkedServiceResourceResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusNotModified) {
		return LinkedServiceResourceResponse{}, client.getLinkedServiceHandleError(resp)
	}
	return client.getLinkedServiceHandleResponse(resp)
}

// getLinkedServiceCreateRequest creates the GetLinkedService request.
func (client *LinkedServiceClient) getLinkedServiceCreateRequest(ctx context.Context, linkedServiceName string, options *LinkedServiceGetLinkedServiceOptions) (*azcore.Request, error) {
	urlPath := "/linkedservices/{linkedServiceName}"
	if linkedServiceName == "" {
		return nil, errors.New("parameter linkedServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{linkedServiceName}", url.PathEscape(linkedServiceName))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfNoneMatch != nil {
		req.Header.Set("If-None-Match", *options.IfNoneMatch)
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getLinkedServiceHandleResponse handles the GetLinkedService response.
func (client *LinkedServiceClient) getLinkedServiceHandleResponse(resp *azcore.Response) (LinkedServiceResourceResponse, error) {
	var val *LinkedServiceResource
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return LinkedServiceResourceResponse{}, err
	}
	return LinkedServiceResourceResponse{RawResponse: resp.Response, LinkedServiceResource: val}, nil
}

// getLinkedServiceHandleError handles the GetLinkedService error response.
func (client *LinkedServiceClient) getLinkedServiceHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err.InnerError); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetLinkedServicesByWorkspace - Lists linked services.
func (client *LinkedServiceClient) GetLinkedServicesByWorkspace(options *LinkedServiceGetLinkedServicesByWorkspaceOptions) LinkedServiceListResponsePager {
	return &linkedServiceListResponsePager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.getLinkedServicesByWorkspaceCreateRequest(ctx, options)
		},
		responder: client.getLinkedServicesByWorkspaceHandleResponse,
		errorer:   client.getLinkedServicesByWorkspaceHandleError,
		advancer: func(ctx context.Context, resp LinkedServiceListResponseResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.LinkedServiceListResponse.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// getLinkedServicesByWorkspaceCreateRequest creates the GetLinkedServicesByWorkspace request.
func (client *LinkedServiceClient) getLinkedServicesByWorkspaceCreateRequest(ctx context.Context, options *LinkedServiceGetLinkedServicesByWorkspaceOptions) (*azcore.Request, error) {
	urlPath := "/linkedservices"
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

// getLinkedServicesByWorkspaceHandleResponse handles the GetLinkedServicesByWorkspace response.
func (client *LinkedServiceClient) getLinkedServicesByWorkspaceHandleResponse(resp *azcore.Response) (LinkedServiceListResponseResponse, error) {
	var val *LinkedServiceListResponse
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return LinkedServiceListResponseResponse{}, err
	}
	return LinkedServiceListResponseResponse{RawResponse: resp.Response, LinkedServiceListResponse: val}, nil
}

// getLinkedServicesByWorkspaceHandleError handles the GetLinkedServicesByWorkspace error response.
func (client *LinkedServiceClient) getLinkedServicesByWorkspaceHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err.InnerError); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// BeginRenameLinkedService - Renames a linked service.
func (client *LinkedServiceClient) BeginRenameLinkedService(ctx context.Context, linkedServiceName string, request ArtifactRenameRequest, options *LinkedServiceBeginRenameLinkedServiceOptions) (HTTPPollerResponse, error) {
	resp, err := client.renameLinkedService(ctx, linkedServiceName, request, options)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	result := HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := azcore.NewLROPoller("LinkedServiceClient.RenameLinkedService", resp, client.con.Pipeline(), client.renameLinkedServiceHandleError)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	poller := &httpPoller{
		pt: pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*http.Response, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeRenameLinkedService creates a new HTTPPoller from the specified resume token.
// token - The value must come from a previous call to HTTPPoller.ResumeToken().
func (client *LinkedServiceClient) ResumeRenameLinkedService(token string) (HTTPPoller, error) {
	pt, err := azcore.NewLROPollerFromResumeToken("LinkedServiceClient.RenameLinkedService", token, client.con.Pipeline(), client.renameLinkedServiceHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pt: pt,
	}, nil
}

// RenameLinkedService - Renames a linked service.
func (client *LinkedServiceClient) renameLinkedService(ctx context.Context, linkedServiceName string, request ArtifactRenameRequest, options *LinkedServiceBeginRenameLinkedServiceOptions) (*azcore.Response, error) {
	req, err := client.renameLinkedServiceCreateRequest(ctx, linkedServiceName, request, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.renameLinkedServiceHandleError(resp)
	}
	return resp, nil
}

// renameLinkedServiceCreateRequest creates the RenameLinkedService request.
func (client *LinkedServiceClient) renameLinkedServiceCreateRequest(ctx context.Context, linkedServiceName string, request ArtifactRenameRequest, options *LinkedServiceBeginRenameLinkedServiceOptions) (*azcore.Request, error) {
	urlPath := "/linkedservices/{linkedServiceName}/rename"
	if linkedServiceName == "" {
		return nil, errors.New("parameter linkedServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{linkedServiceName}", url.PathEscape(linkedServiceName))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(request)
}

// renameLinkedServiceHandleError handles the RenameLinkedService error response.
func (client *LinkedServiceClient) renameLinkedServiceHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err.InnerError); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
