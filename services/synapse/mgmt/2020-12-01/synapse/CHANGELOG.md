# Unreleased

## Additive Changes

### New Constants

1. NodeSizeFamily.NodeSizeFamilyHardwareAcceleratedFPGA
1. NodeSizeFamily.NodeSizeFamilyHardwareAcceleratedGPU

### New Funcs

1. *LibraryListResponseIterator.Next() error
1. *LibraryListResponseIterator.NextWithContext(context.Context) error
1. *LibraryListResponsePage.Next() error
1. *LibraryListResponsePage.NextWithContext(context.Context) error
1. *LibraryResource.UnmarshalJSON([]byte) error
1. *SparkConfigurationListResponseIterator.Next() error
1. *SparkConfigurationListResponseIterator.NextWithContext(context.Context) error
1. *SparkConfigurationListResponsePage.Next() error
1. *SparkConfigurationListResponsePage.NextWithContext(context.Context) error
1. *SparkConfigurationResource.UnmarshalJSON([]byte) error
1. LibrariesClient.ListByWorkspace(context.Context, string, string) (LibraryListResponsePage, error)
1. LibrariesClient.ListByWorkspaceComplete(context.Context, string, string) (LibraryListResponseIterator, error)
1. LibrariesClient.ListByWorkspacePreparer(context.Context, string, string) (*http.Request, error)
1. LibrariesClient.ListByWorkspaceResponder(*http.Response) (LibraryListResponse, error)
1. LibrariesClient.ListByWorkspaceSender(*http.Request) (*http.Response, error)
1. LibraryClient.Get(context.Context, string, string, string) (LibraryResource, error)
1. LibraryClient.GetPreparer(context.Context, string, string, string) (*http.Request, error)
1. LibraryClient.GetResponder(*http.Response) (LibraryResource, error)
1. LibraryClient.GetSender(*http.Request) (*http.Response, error)
1. LibraryListResponse.IsEmpty() bool
1. LibraryListResponseIterator.NotDone() bool
1. LibraryListResponseIterator.Response() LibraryListResponse
1. LibraryListResponseIterator.Value() LibraryResource
1. LibraryListResponsePage.NotDone() bool
1. LibraryListResponsePage.Response() LibraryListResponse
1. LibraryListResponsePage.Values() []LibraryResource
1. LibraryResource.MarshalJSON() ([]byte, error)
1. NewLibrariesClient(string) LibrariesClient
1. NewLibrariesClientWithBaseURI(string, string) LibrariesClient
1. NewLibraryClient(string) LibraryClient
1. NewLibraryClientWithBaseURI(string, string) LibraryClient
1. NewLibraryListResponseIterator(LibraryListResponsePage) LibraryListResponseIterator
1. NewLibraryListResponsePage(LibraryListResponse, func(context.Context, LibraryListResponse) (LibraryListResponse, error)) LibraryListResponsePage
1. NewSparkConfigurationClient(string) SparkConfigurationClient
1. NewSparkConfigurationClientWithBaseURI(string, string) SparkConfigurationClient
1. NewSparkConfigurationListResponseIterator(SparkConfigurationListResponsePage) SparkConfigurationListResponseIterator
1. NewSparkConfigurationListResponsePage(SparkConfigurationListResponse, func(context.Context, SparkConfigurationListResponse) (SparkConfigurationListResponse, error)) SparkConfigurationListResponsePage
1. NewSparkConfigurationsClient(string) SparkConfigurationsClient
1. NewSparkConfigurationsClientWithBaseURI(string, string) SparkConfigurationsClient
1. SparkConfigurationClient.Get(context.Context, string, string, string) (SparkConfigurationResource, error)
1. SparkConfigurationClient.GetPreparer(context.Context, string, string, string) (*http.Request, error)
1. SparkConfigurationClient.GetResponder(*http.Response) (SparkConfigurationResource, error)
1. SparkConfigurationClient.GetSender(*http.Request) (*http.Response, error)
1. SparkConfigurationListResponse.IsEmpty() bool
1. SparkConfigurationListResponseIterator.NotDone() bool
1. SparkConfigurationListResponseIterator.Response() SparkConfigurationListResponse
1. SparkConfigurationListResponseIterator.Value() SparkConfigurationResource
1. SparkConfigurationListResponsePage.NotDone() bool
1. SparkConfigurationListResponsePage.Response() SparkConfigurationListResponse
1. SparkConfigurationListResponsePage.Values() []SparkConfigurationResource
1. SparkConfigurationResource.MarshalJSON() ([]byte, error)
1. SparkConfigurationsClient.ListByWorkspace(context.Context, string, string) (SparkConfigurationListResponsePage, error)
1. SparkConfigurationsClient.ListByWorkspaceComplete(context.Context, string, string) (SparkConfigurationListResponseIterator, error)
1. SparkConfigurationsClient.ListByWorkspacePreparer(context.Context, string, string) (*http.Request, error)
1. SparkConfigurationsClient.ListByWorkspaceResponder(*http.Response) (SparkConfigurationListResponse, error)
1. SparkConfigurationsClient.ListByWorkspaceSender(*http.Request) (*http.Response, error)

### Struct Changes

#### New Structs

1. LibrariesClient
1. LibraryClient
1. LibraryListResponse
1. LibraryListResponseIterator
1. LibraryListResponsePage
1. LibraryResource
1. SparkConfigurationClient
1. SparkConfigurationInfo
1. SparkConfigurationListResponse
1. SparkConfigurationListResponseIterator
1. SparkConfigurationListResponsePage
1. SparkConfigurationResource
1. SparkConfigurationsClient
