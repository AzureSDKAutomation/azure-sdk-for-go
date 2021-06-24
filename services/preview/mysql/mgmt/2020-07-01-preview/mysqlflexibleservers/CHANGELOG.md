# Unreleased

## Additive Changes

### New Funcs

1. *ConfigurationsBatchUpdateFuture.UnmarshalJSON([]byte) error
1. ConfigurationsClient.BatchUpdate(context.Context, string, string, ConfigurationListResult) (ConfigurationsBatchUpdateFuture, error)
1. ConfigurationsClient.BatchUpdatePreparer(context.Context, string, string, ConfigurationListResult) (*http.Request, error)
1. ConfigurationsClient.BatchUpdateResponder(*http.Response) (ConfigurationListResult, error)
1. ConfigurationsClient.BatchUpdateSender(*http.Request) (ConfigurationsBatchUpdateFuture, error)
1. GetPrivateDNSZoneSuffixClient.Execute(context.Context, string, string) (String, error)
1. GetPrivateDNSZoneSuffixClient.ExecutePreparer(context.Context, string, string) (*http.Request, error)
1. GetPrivateDNSZoneSuffixClient.ExecuteResponder(*http.Response) (String, error)
1. GetPrivateDNSZoneSuffixClient.ExecuteSender(*http.Request) (*http.Response, error)
1. NewGetPrivateDNSZoneSuffixClient(string) GetPrivateDNSZoneSuffixClient
1. NewGetPrivateDNSZoneSuffixClientWithBaseURI(string, string) GetPrivateDNSZoneSuffixClient
1. StorageProfile.MarshalJSON() ([]byte, error)

### Struct Changes

#### New Structs

1. ConfigurationsBatchUpdateFuture
1. GetPrivateDNSZoneSuffixClient
1. String

#### New Struct Fields

1. CapabilityProperties.Status
1. ServerEditionCapability.Status
1. ServerVersionCapability.Status
1. StorageEditionCapability.Status
1. StorageProfile.FileStorageSkuName
1. VcoreCapability.Status
