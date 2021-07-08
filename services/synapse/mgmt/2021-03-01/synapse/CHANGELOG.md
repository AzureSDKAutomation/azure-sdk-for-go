# Unreleased

## Additive Changes

### New Constants

1. NodeSizeFamily.NodeSizeFamilyHardwareAcceleratedFPGA
1. NodeSizeFamily.NodeSizeFamilyHardwareAcceleratedGPU

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
