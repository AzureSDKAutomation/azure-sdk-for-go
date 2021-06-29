package containerregistryapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/preview/containerregistry/runtime/2019-08-15-preview/containerregistry"
	"github.com/Azure/go-autorest/autorest"
	"io"
)

// V2SupportClientAPI contains the set of methods on the V2SupportClient type.
type V2SupportClientAPI interface {
	Check(ctx context.Context) (result autorest.Response, err error)
}

var _ V2SupportClientAPI = (*containerregistry.V2SupportClient)(nil)

// ManifestsClientAPI contains the set of methods on the ManifestsClient type.
type ManifestsClientAPI interface {
	Create(ctx context.Context, name string, reference string, payload containerregistry.Manifest) (result containerregistry.SetObject, err error)
	Delete(ctx context.Context, name string, reference string) (result autorest.Response, err error)
	Get(ctx context.Context, name string, reference string, accept string) (result containerregistry.ManifestWrapper, err error)
	GetAttributes(ctx context.Context, name string, reference string) (result containerregistry.ManifestAttributes, err error)
	GetList(ctx context.Context, name string, last string, n *int32, orderby string) (result containerregistry.AcrManifests, err error)
	UpdateAttributes(ctx context.Context, name string, reference string, value *containerregistry.ManifestChangeableAttributes) (result autorest.Response, err error)
}

var _ ManifestsClientAPI = (*containerregistry.ManifestsClient)(nil)

// BlobClientAPI contains the set of methods on the BlobClient type.
type BlobClientAPI interface {
	CancelUpload(ctx context.Context, location string) (result autorest.Response, err error)
	Check(ctx context.Context, name string, digest string) (result autorest.Response, err error)
	CheckChunk(ctx context.Context, name string, digest string, rangeParameter string) (result autorest.Response, err error)
	Delete(ctx context.Context, name string, digest string) (result containerregistry.ReadCloser, err error)
	EndUpload(ctx context.Context, digest string, location string, value io.ReadCloser) (result autorest.Response, err error)
	Get(ctx context.Context, name string, digest string) (result containerregistry.ReadCloser, err error)
	GetChunk(ctx context.Context, name string, digest string, rangeParameter string) (result containerregistry.ReadCloser, err error)
	GetStatus(ctx context.Context, location string) (result autorest.Response, err error)
	Mount(ctx context.Context, name string, from string, mount string) (result autorest.Response, err error)
	StartUpload(ctx context.Context, name string) (result autorest.Response, err error)
	Upload(ctx context.Context, value io.ReadCloser, location string) (result autorest.Response, err error)
}

var _ BlobClientAPI = (*containerregistry.BlobClient)(nil)

// RepositoryClientAPI contains the set of methods on the RepositoryClient type.
type RepositoryClientAPI interface {
	Delete(ctx context.Context, name string) (result containerregistry.DeletedRepository, err error)
	GetAttributes(ctx context.Context, name string) (result containerregistry.RepositoryAttributes, err error)
	GetList(ctx context.Context, last string, n *int32) (result containerregistry.Repositories, err error)
	UpdateAttributes(ctx context.Context, name string, value *containerregistry.RepositoryChangeableAttributes) (result autorest.Response, err error)
}

var _ RepositoryClientAPI = (*containerregistry.RepositoryClient)(nil)

// TagClientAPI contains the set of methods on the TagClient type.
type TagClientAPI interface {
	Delete(ctx context.Context, name string, reference string) (result autorest.Response, err error)
	GetAttributes(ctx context.Context, name string, reference string) (result containerregistry.TagAttributes, err error)
	GetList(ctx context.Context, name string, last string, n *int32, orderby string, digest string) (result containerregistry.TagList, err error)
	UpdateAttributes(ctx context.Context, name string, reference string, value *containerregistry.TagChangeableAttributes) (result autorest.Response, err error)
}

var _ TagClientAPI = (*containerregistry.TagClient)(nil)

// RefreshTokensClientAPI contains the set of methods on the RefreshTokensClient type.
type RefreshTokensClientAPI interface {
	GetFromExchange(ctx context.Context, grantType string, service string, tenant string, refreshToken string, accessToken string) (result containerregistry.RefreshToken, err error)
}

var _ RefreshTokensClientAPI = (*containerregistry.RefreshTokensClient)(nil)

// AccessTokensClientAPI contains the set of methods on the AccessTokensClient type.
type AccessTokensClientAPI interface {
	Get(ctx context.Context, service string, scope string, refreshToken string) (result containerregistry.AccessToken, err error)
	GetFromLogin(ctx context.Context, service string, scope string) (result containerregistry.AccessToken, err error)
}

var _ AccessTokensClientAPI = (*containerregistry.AccessTokensClient)(nil)
