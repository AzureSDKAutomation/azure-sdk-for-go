# Unreleased

## Additive Changes

### New Funcs

1. *AntiMalwareSetting.UnmarshalJSON([]byte) error
1. AntiMalwareClient.Create(context.Context, string, AntiMalwareSetting) (AntiMalwareSetting, error)
1. AntiMalwareClient.CreatePreparer(context.Context, string, AntiMalwareSetting) (*http.Request, error)
1. AntiMalwareClient.CreateResponder(*http.Response) (AntiMalwareSetting, error)
1. AntiMalwareClient.CreateSender(*http.Request) (*http.Response, error)
1. AntiMalwareClient.Get(context.Context, string) (AntiMalwareSetting, error)
1. AntiMalwareClient.GetPreparer(context.Context, string) (*http.Request, error)
1. AntiMalwareClient.GetResponder(*http.Response) (AntiMalwareSetting, error)
1. AntiMalwareClient.GetSender(*http.Request) (*http.Response, error)
1. AntiMalwareSetting.MarshalJSON() ([]byte, error)
1. NewAntiMalwareClient(string, string) AntiMalwareClient
1. NewAntiMalwareClientWithBaseURI(string, string, string) AntiMalwareClient

### Struct Changes

#### New Structs

1. AntiMalwareClient
1. AntiMalwareProperties
1. AntiMalwareSetting
