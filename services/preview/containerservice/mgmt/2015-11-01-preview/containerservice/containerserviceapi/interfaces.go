package containerserviceapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/preview/containerservice/mgmt/2015-11-01-preview/containerservice"
)

// ClientAPI contains the set of methods on the Client type.
type ClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, containerServiceName string, parameters containerservice.ContainerService) (result containerservice.CreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, containerServiceName string) (result containerservice.DeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, containerServiceName string) (result containerservice.ContainerService, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result containerservice.ListResult, err error)
}

var _ ClientAPI = (*containerservice.Client)(nil)
