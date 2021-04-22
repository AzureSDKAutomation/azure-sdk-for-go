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
	"strconv"
	"strings"
)

// PipelineRunClient contains the methods for the PipelineRun group.
// Don't use this type directly, use NewPipelineRunClient() instead.
type PipelineRunClient struct {
	con *Connection
}

// NewPipelineRunClient creates a new instance of PipelineRunClient with the specified values.
func NewPipelineRunClient(con *Connection) *PipelineRunClient {
	return &PipelineRunClient{con: con}
}

// CancelPipelineRun - Cancel a pipeline run by its run ID.
func (client *PipelineRunClient) CancelPipelineRun(ctx context.Context, runID string, options *PipelineRunCancelPipelineRunOptions) (*http.Response, error) {
	req, err := client.cancelPipelineRunCreateRequest(ctx, runID, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.cancelPipelineRunHandleError(resp)
	}
	return resp.Response, nil
}

// cancelPipelineRunCreateRequest creates the CancelPipelineRun request.
func (client *PipelineRunClient) cancelPipelineRunCreateRequest(ctx context.Context, runID string, options *PipelineRunCancelPipelineRunOptions) (*azcore.Request, error) {
	urlPath := "/pipelineruns/{runId}/cancel"
	if runID == "" {
		return nil, errors.New("parameter runID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{runId}", url.PathEscape(runID))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	if options != nil && options.IsRecursive != nil {
		reqQP.Set("isRecursive", strconv.FormatBool(*options.IsRecursive))
	}
	reqQP.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// cancelPipelineRunHandleError handles the CancelPipelineRun error response.
func (client *PipelineRunClient) cancelPipelineRunHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err.InnerError); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetPipelineRun - Get a pipeline run by its run ID.
func (client *PipelineRunClient) GetPipelineRun(ctx context.Context, runID string, options *PipelineRunGetPipelineRunOptions) (PipelineRunResponse, error) {
	req, err := client.getPipelineRunCreateRequest(ctx, runID, options)
	if err != nil {
		return PipelineRunResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return PipelineRunResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return PipelineRunResponse{}, client.getPipelineRunHandleError(resp)
	}
	return client.getPipelineRunHandleResponse(resp)
}

// getPipelineRunCreateRequest creates the GetPipelineRun request.
func (client *PipelineRunClient) getPipelineRunCreateRequest(ctx context.Context, runID string, options *PipelineRunGetPipelineRunOptions) (*azcore.Request, error) {
	urlPath := "/pipelineruns/{runId}"
	if runID == "" {
		return nil, errors.New("parameter runID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{runId}", url.PathEscape(runID))
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

// getPipelineRunHandleResponse handles the GetPipelineRun response.
func (client *PipelineRunClient) getPipelineRunHandleResponse(resp *azcore.Response) (PipelineRunResponse, error) {
	var val *PipelineRun
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return PipelineRunResponse{}, err
	}
	return PipelineRunResponse{RawResponse: resp.Response, PipelineRun: val}, nil
}

// getPipelineRunHandleError handles the GetPipelineRun error response.
func (client *PipelineRunClient) getPipelineRunHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err.InnerError); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// QueryActivityRuns - Query activity runs based on input filter conditions.
func (client *PipelineRunClient) QueryActivityRuns(ctx context.Context, pipelineName string, runID string, filterParameters RunFilterParameters, options *PipelineRunQueryActivityRunsOptions) (ActivityRunsQueryResponseResponse, error) {
	req, err := client.queryActivityRunsCreateRequest(ctx, pipelineName, runID, filterParameters, options)
	if err != nil {
		return ActivityRunsQueryResponseResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return ActivityRunsQueryResponseResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return ActivityRunsQueryResponseResponse{}, client.queryActivityRunsHandleError(resp)
	}
	return client.queryActivityRunsHandleResponse(resp)
}

// queryActivityRunsCreateRequest creates the QueryActivityRuns request.
func (client *PipelineRunClient) queryActivityRunsCreateRequest(ctx context.Context, pipelineName string, runID string, filterParameters RunFilterParameters, options *PipelineRunQueryActivityRunsOptions) (*azcore.Request, error) {
	urlPath := "/pipelines/{pipelineName}/pipelineruns/{runId}/queryActivityruns"
	if pipelineName == "" {
		return nil, errors.New("parameter pipelineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{pipelineName}", url.PathEscape(pipelineName))
	if runID == "" {
		return nil, errors.New("parameter runID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{runId}", url.PathEscape(runID))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(filterParameters)
}

// queryActivityRunsHandleResponse handles the QueryActivityRuns response.
func (client *PipelineRunClient) queryActivityRunsHandleResponse(resp *azcore.Response) (ActivityRunsQueryResponseResponse, error) {
	var val *ActivityRunsQueryResponse
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return ActivityRunsQueryResponseResponse{}, err
	}
	return ActivityRunsQueryResponseResponse{RawResponse: resp.Response, ActivityRunsQueryResponse: val}, nil
}

// queryActivityRunsHandleError handles the QueryActivityRuns error response.
func (client *PipelineRunClient) queryActivityRunsHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err.InnerError); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// QueryPipelineRunsByWorkspace - Query pipeline runs in the workspace based on input filter conditions.
func (client *PipelineRunClient) QueryPipelineRunsByWorkspace(ctx context.Context, filterParameters RunFilterParameters, options *PipelineRunQueryPipelineRunsByWorkspaceOptions) (PipelineRunsQueryResponseResponse, error) {
	req, err := client.queryPipelineRunsByWorkspaceCreateRequest(ctx, filterParameters, options)
	if err != nil {
		return PipelineRunsQueryResponseResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return PipelineRunsQueryResponseResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return PipelineRunsQueryResponseResponse{}, client.queryPipelineRunsByWorkspaceHandleError(resp)
	}
	return client.queryPipelineRunsByWorkspaceHandleResponse(resp)
}

// queryPipelineRunsByWorkspaceCreateRequest creates the QueryPipelineRunsByWorkspace request.
func (client *PipelineRunClient) queryPipelineRunsByWorkspaceCreateRequest(ctx context.Context, filterParameters RunFilterParameters, options *PipelineRunQueryPipelineRunsByWorkspaceOptions) (*azcore.Request, error) {
	urlPath := "/queryPipelineRuns"
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(filterParameters)
}

// queryPipelineRunsByWorkspaceHandleResponse handles the QueryPipelineRunsByWorkspace response.
func (client *PipelineRunClient) queryPipelineRunsByWorkspaceHandleResponse(resp *azcore.Response) (PipelineRunsQueryResponseResponse, error) {
	var val *PipelineRunsQueryResponse
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return PipelineRunsQueryResponseResponse{}, err
	}
	return PipelineRunsQueryResponseResponse{RawResponse: resp.Response, PipelineRunsQueryResponse: val}, nil
}

// queryPipelineRunsByWorkspaceHandleError handles the QueryPipelineRunsByWorkspace error response.
func (client *PipelineRunClient) queryPipelineRunsByWorkspaceHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err.InnerError); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
