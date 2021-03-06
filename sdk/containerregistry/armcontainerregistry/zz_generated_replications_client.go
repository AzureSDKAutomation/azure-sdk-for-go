// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcontainerregistry

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// ReplicationsClient contains the methods for the Replications group.
// Don't use this type directly, use NewReplicationsClient() instead.
type ReplicationsClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewReplicationsClient creates a new instance of ReplicationsClient with the specified values.
func NewReplicationsClient(con *armcore.Connection, subscriptionID string) *ReplicationsClient {
	return &ReplicationsClient{con: con, subscriptionID: subscriptionID}
}

// BeginCreate - Creates a replication for a container registry with the specified parameters.
// If the operation fails it returns a generic error.
func (client *ReplicationsClient) BeginCreate(ctx context.Context, resourceGroupName string, registryName string, replicationName string, replication Replication, options *ReplicationsBeginCreateOptions) (ReplicationPollerResponse, error) {
	resp, err := client.create(ctx, resourceGroupName, registryName, replicationName, replication, options)
	if err != nil {
		return ReplicationPollerResponse{}, err
	}
	result := ReplicationPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewLROPoller("ReplicationsClient.Create", "", resp, client.con.Pipeline(), client.createHandleError)
	if err != nil {
		return ReplicationPollerResponse{}, err
	}
	poller := &replicationPoller{
		pt: pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (ReplicationResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeCreate creates a new ReplicationPoller from the specified resume token.
// token - The value must come from a previous call to ReplicationPoller.ResumeToken().
func (client *ReplicationsClient) ResumeCreate(ctx context.Context, token string) (ReplicationPollerResponse, error) {
	pt, err := armcore.NewLROPollerFromResumeToken("ReplicationsClient.Create", token, client.con.Pipeline(), client.createHandleError)
	if err != nil {
		return ReplicationPollerResponse{}, err
	}
	poller := &replicationPoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return ReplicationPollerResponse{}, err
	}
	result := ReplicationPollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (ReplicationResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// Create - Creates a replication for a container registry with the specified parameters.
// If the operation fails it returns a generic error.
func (client *ReplicationsClient) create(ctx context.Context, resourceGroupName string, registryName string, replicationName string, replication Replication, options *ReplicationsBeginCreateOptions) (*azcore.Response, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, registryName, replicationName, replication, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated) {
		return nil, client.createHandleError(resp)
	}
	return resp, nil
}

// createCreateRequest creates the Create request.
func (client *ReplicationsClient) createCreateRequest(ctx context.Context, resourceGroupName string, registryName string, replicationName string, replication Replication, options *ReplicationsBeginCreateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/replications/{replicationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if registryName == "" {
		return nil, errors.New("parameter registryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registryName}", url.PathEscape(registryName))
	if replicationName == "" {
		return nil, errors.New("parameter replicationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{replicationName}", url.PathEscape(replicationName))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(replication)
}

// createHandleError handles the Create error response.
func (client *ReplicationsClient) createHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// BeginDelete - Deletes a replication from a container registry.
// If the operation fails it returns a generic error.
func (client *ReplicationsClient) BeginDelete(ctx context.Context, resourceGroupName string, registryName string, replicationName string, options *ReplicationsBeginDeleteOptions) (HTTPPollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, registryName, replicationName, options)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	result := HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewLROPoller("ReplicationsClient.Delete", "", resp, client.con.Pipeline(), client.deleteHandleError)
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

// ResumeDelete creates a new HTTPPoller from the specified resume token.
// token - The value must come from a previous call to HTTPPoller.ResumeToken().
func (client *ReplicationsClient) ResumeDelete(ctx context.Context, token string) (HTTPPollerResponse, error) {
	pt, err := armcore.NewLROPollerFromResumeToken("ReplicationsClient.Delete", token, client.con.Pipeline(), client.deleteHandleError)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	poller := &httpPoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	result := HTTPPollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*http.Response, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// Delete - Deletes a replication from a container registry.
// If the operation fails it returns a generic error.
func (client *ReplicationsClient) deleteOperation(ctx context.Context, resourceGroupName string, registryName string, replicationName string, options *ReplicationsBeginDeleteOptions) (*azcore.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, registryName, replicationName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ReplicationsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, registryName string, replicationName string, options *ReplicationsBeginDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/replications/{replicationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if registryName == "" {
		return nil, errors.New("parameter registryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registryName}", url.PathEscape(registryName))
	if replicationName == "" {
		return nil, errors.New("parameter replicationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{replicationName}", url.PathEscape(replicationName))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *ReplicationsClient) deleteHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// Get - Gets the properties of the specified replication.
// If the operation fails it returns a generic error.
func (client *ReplicationsClient) Get(ctx context.Context, resourceGroupName string, registryName string, replicationName string, options *ReplicationsGetOptions) (ReplicationResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, registryName, replicationName, options)
	if err != nil {
		return ReplicationResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return ReplicationResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return ReplicationResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ReplicationsClient) getCreateRequest(ctx context.Context, resourceGroupName string, registryName string, replicationName string, options *ReplicationsGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/replications/{replicationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if registryName == "" {
		return nil, errors.New("parameter registryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registryName}", url.PathEscape(registryName))
	if replicationName == "" {
		return nil, errors.New("parameter replicationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{replicationName}", url.PathEscape(replicationName))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ReplicationsClient) getHandleResponse(resp *azcore.Response) (ReplicationResponse, error) {
	var val *Replication
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return ReplicationResponse{}, err
	}
	return ReplicationResponse{RawResponse: resp.Response, Replication: val}, nil
}

// getHandleError handles the Get error response.
func (client *ReplicationsClient) getHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// List - Lists all the replications for the specified container registry.
// If the operation fails it returns a generic error.
func (client *ReplicationsClient) List(resourceGroupName string, registryName string, options *ReplicationsListOptions) ReplicationListResultPager {
	return &replicationListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, registryName, options)
		},
		responder: client.listHandleResponse,
		errorer:   client.listHandleError,
		advancer: func(ctx context.Context, resp ReplicationListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.ReplicationListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listCreateRequest creates the List request.
func (client *ReplicationsClient) listCreateRequest(ctx context.Context, resourceGroupName string, registryName string, options *ReplicationsListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/replications"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if registryName == "" {
		return nil, errors.New("parameter registryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registryName}", url.PathEscape(registryName))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ReplicationsClient) listHandleResponse(resp *azcore.Response) (ReplicationListResultResponse, error) {
	var val *ReplicationListResult
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return ReplicationListResultResponse{}, err
	}
	return ReplicationListResultResponse{RawResponse: resp.Response, ReplicationListResult: val}, nil
}

// listHandleError handles the List error response.
func (client *ReplicationsClient) listHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// BeginUpdate - Updates a replication for a container registry with the specified parameters.
// If the operation fails it returns a generic error.
func (client *ReplicationsClient) BeginUpdate(ctx context.Context, resourceGroupName string, registryName string, replicationName string, replicationUpdateParameters ReplicationUpdateParameters, options *ReplicationsBeginUpdateOptions) (ReplicationPollerResponse, error) {
	resp, err := client.update(ctx, resourceGroupName, registryName, replicationName, replicationUpdateParameters, options)
	if err != nil {
		return ReplicationPollerResponse{}, err
	}
	result := ReplicationPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewLROPoller("ReplicationsClient.Update", "", resp, client.con.Pipeline(), client.updateHandleError)
	if err != nil {
		return ReplicationPollerResponse{}, err
	}
	poller := &replicationPoller{
		pt: pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (ReplicationResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeUpdate creates a new ReplicationPoller from the specified resume token.
// token - The value must come from a previous call to ReplicationPoller.ResumeToken().
func (client *ReplicationsClient) ResumeUpdate(ctx context.Context, token string) (ReplicationPollerResponse, error) {
	pt, err := armcore.NewLROPollerFromResumeToken("ReplicationsClient.Update", token, client.con.Pipeline(), client.updateHandleError)
	if err != nil {
		return ReplicationPollerResponse{}, err
	}
	poller := &replicationPoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return ReplicationPollerResponse{}, err
	}
	result := ReplicationPollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (ReplicationResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// Update - Updates a replication for a container registry with the specified parameters.
// If the operation fails it returns a generic error.
func (client *ReplicationsClient) update(ctx context.Context, resourceGroupName string, registryName string, replicationName string, replicationUpdateParameters ReplicationUpdateParameters, options *ReplicationsBeginUpdateOptions) (*azcore.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, registryName, replicationName, replicationUpdateParameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated) {
		return nil, client.updateHandleError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *ReplicationsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, registryName string, replicationName string, replicationUpdateParameters ReplicationUpdateParameters, options *ReplicationsBeginUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/replications/{replicationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if registryName == "" {
		return nil, errors.New("parameter registryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registryName}", url.PathEscape(registryName))
	if replicationName == "" {
		return nil, errors.New("parameter replicationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{replicationName}", url.PathEscape(replicationName))
	req, err := azcore.NewRequest(ctx, http.MethodPatch, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(replicationUpdateParameters)
}

// updateHandleError handles the Update error response.
func (client *ReplicationsClient) updateHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}
