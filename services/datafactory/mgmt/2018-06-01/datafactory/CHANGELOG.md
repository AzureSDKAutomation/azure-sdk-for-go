# Unreleased

## Breaking Changes

### Signature Changes

#### Struct Fields

1. AvroDatasetTypeProperties.AvroCompressionCodec changed type from AvroCompressionCodec to interface{}
1. OrcDatasetTypeProperties.OrcCompressionCodec changed type from OrcCompressionCodec to interface{}

## Additive Changes

### New Funcs

1. IntegrationRuntimesClient.OutboundNetworkDependenciesEndpoints(context.Context, string, string, string) (IntegrationRuntimeOutboundNetworkDependenciesEndpointsResponse, error)
1. IntegrationRuntimesClient.OutboundNetworkDependenciesEndpointsPreparer(context.Context, string, string, string) (*http.Request, error)
1. IntegrationRuntimesClient.OutboundNetworkDependenciesEndpointsResponder(*http.Response) (IntegrationRuntimeOutboundNetworkDependenciesEndpointsResponse, error)
1. IntegrationRuntimesClient.OutboundNetworkDependenciesEndpointsSender(*http.Request) (*http.Response, error)

### Struct Changes

#### New Structs

1. IntegrationRuntimeOutboundNetworkDependenciesCategoryEndpoint
1. IntegrationRuntimeOutboundNetworkDependenciesEndpoint
1. IntegrationRuntimeOutboundNetworkDependenciesEndpointDetails
1. IntegrationRuntimeOutboundNetworkDependenciesEndpointsResponse

#### New Struct Fields

1. IntegrationRuntimeVNetProperties.SubnetID
1. PipelineRunInvokedBy.PipelineName
1. PipelineRunInvokedBy.PipelineRunID
