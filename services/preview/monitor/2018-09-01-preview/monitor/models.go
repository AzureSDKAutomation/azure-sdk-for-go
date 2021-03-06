package monitor

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
)

// The package's fully qualified name.
const fqdn = "github.com/Azure/azure-sdk-for-go/services/preview/monitor/2018-09-01-preview/monitor"

// APIError ...
type APIError struct {
	// Code - Gets or sets the azure metrics error code
	Code *string `json:"code,omitempty"`
	// Message - Gets or sets the azure metrics error message
	Message *string `json:"message,omitempty"`
}

// APIFailureResponse ...
type APIFailureResponse struct {
	Error *APIError `json:"error,omitempty"`
}

// AzureMetricsBaseData ...
type AzureMetricsBaseData struct {
	// Metric - Gets or sets the Metric name
	Metric *string `json:"metric,omitempty"`
	// Namespace - Gets or sets the Metric namespace
	Namespace *string `json:"namespace,omitempty"`
	// DimNames - Gets or sets the list of dimension names (optional)
	DimNames *[]string `json:"dimNames,omitempty"`
	// Series - Gets or sets the list of time series data for the metric (one per unique dimension combination)
	Series *[]AzureTimeSeriesData `json:"series,omitempty"`
}

// AzureMetricsData ...
type AzureMetricsData struct {
	BaseData *AzureMetricsBaseData `json:"baseData,omitempty"`
}

// AzureMetricsDocument ...
type AzureMetricsDocument struct {
	// Time - Gets or sets Time property (in ISO 8601 format)
	Time *string           `json:"time,omitempty"`
	Data *AzureMetricsData `json:"data,omitempty"`
}

// AzureMetricsResult ...
type AzureMetricsResult struct {
	autorest.Response `json:"-"`
	// StatusCode - Http status code response
	StatusCode         *int32              `json:"statusCode,omitempty"`
	APIFailureResponse *APIFailureResponse `json:"apiFailureResponse,omitempty"`
}

// AzureTimeSeriesData ...
type AzureTimeSeriesData struct {
	// DimValues - Gets or sets dimension values
	DimValues *[]string `json:"dimValues,omitempty"`
	// Min - Gets or sets Min value
	Min *float64 `json:"min,omitempty"`
	// Max - Gets or sets Max value
	Max *float64 `json:"max,omitempty"`
	// Sum - Gets or sets Sum value
	Sum *float64 `json:"sum,omitempty"`
	// Count - Gets or sets Count value
	Count *int32 `json:"count,omitempty"`
}
