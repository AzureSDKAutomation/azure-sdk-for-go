# Unreleased

## Breaking Changes

### Signature Changes

#### Funcs

1. ConnectorsClient.CreateOrUpdate
	- Params
		- From: context.Context, string, ConnectorSetting
		- To: context.Context, string, string, Connector
	- Returns
		- From: ConnectorSetting, error
		- To: Connector, error
1. ConnectorsClient.CreateOrUpdatePreparer
	- Params
		- From: context.Context, string, ConnectorSetting
		- To: context.Context, string, string, Connector
1. ConnectorsClient.CreateOrUpdateResponder
	- Returns
		- From: ConnectorSetting, error
		- To: Connector, error
1. ConnectorsClient.Delete
	- Params
		- From: context.Context, string
		- To: context.Context, string, string
1. ConnectorsClient.DeletePreparer
	- Params
		- From: context.Context, string
		- To: context.Context, string, string
1. ConnectorsClient.Get
	- Params
		- From: context.Context, string
		- To: context.Context, string, string
	- Returns
		- From: ConnectorSetting, error
		- To: Connector, error
1. ConnectorsClient.GetPreparer
	- Params
		- From: context.Context, string
		- To: context.Context, string, string
1. ConnectorsClient.GetResponder
	- Returns
		- From: ConnectorSetting, error
		- To: Connector, error
1. ConnectorsClient.List
	- Returns
		- From: ConnectorSettingListPage, error
		- To: ConnectorsListPage, error
1. ConnectorsClient.ListComplete
	- Returns
		- From: ConnectorSettingListIterator, error
		- To: ConnectorsListIterator, error
1. ConnectorsClient.ListResponder
	- Returns
		- From: ConnectorSettingList, error
		- To: ConnectorsList, error

## Additive Changes

### New Constants

1. MultiCloudName.AWS
1. MultiCloudName.Azure
1. MultiCloudName.GCP
1. OfferingType.OfferingTypeCSPMMonitorAWS
1. OfferingType.OfferingTypeMultiCloudOffering

### New Funcs

1. *Connector.UnmarshalJSON([]byte) error
1. *ConnectorProperties.UnmarshalJSON([]byte) error
1. *ConnectorsListIterator.Next() error
1. *ConnectorsListIterator.NextWithContext(context.Context) error
1. *ConnectorsListPage.Next() error
1. *ConnectorsListPage.NextWithContext(context.Context) error
1. CSPMMonitorAWSOffering.AsBasicMultiCloudOffering() (BasicMultiCloudOffering, bool)
1. CSPMMonitorAWSOffering.AsCSPMMonitorAWSOffering() (*CSPMMonitorAWSOffering, bool)
1. CSPMMonitorAWSOffering.AsMultiCloudOffering() (*MultiCloudOffering, bool)
1. CSPMMonitorAWSOffering.MarshalJSON() ([]byte, error)
1. CSPMMonitorAWSOfferingNativeCloudConnection.MarshalJSON() ([]byte, error)
1. Connector.MarshalJSON() ([]byte, error)
1. ConnectorsClient.ListByResourceGroup(context.Context, string) (ConnectorsListPage, error)
1. ConnectorsClient.ListByResourceGroupComplete(context.Context, string) (ConnectorsListIterator, error)
1. ConnectorsClient.ListByResourceGroupPreparer(context.Context, string) (*http.Request, error)
1. ConnectorsClient.ListByResourceGroupResponder(*http.Response) (ConnectorsList, error)
1. ConnectorsClient.ListByResourceGroupSender(*http.Request) (*http.Response, error)
1. ConnectorsGroupClient.CreateOrUpdate(context.Context, string, ConnectorSetting) (ConnectorSetting, error)
1. ConnectorsGroupClient.CreateOrUpdatePreparer(context.Context, string, ConnectorSetting) (*http.Request, error)
1. ConnectorsGroupClient.CreateOrUpdateResponder(*http.Response) (ConnectorSetting, error)
1. ConnectorsGroupClient.CreateOrUpdateSender(*http.Request) (*http.Response, error)
1. ConnectorsGroupClient.Delete(context.Context, string) (autorest.Response, error)
1. ConnectorsGroupClient.DeletePreparer(context.Context, string) (*http.Request, error)
1. ConnectorsGroupClient.DeleteResponder(*http.Response) (autorest.Response, error)
1. ConnectorsGroupClient.DeleteSender(*http.Request) (*http.Response, error)
1. ConnectorsGroupClient.Get(context.Context, string) (ConnectorSetting, error)
1. ConnectorsGroupClient.GetPreparer(context.Context, string) (*http.Request, error)
1. ConnectorsGroupClient.GetResponder(*http.Response) (ConnectorSetting, error)
1. ConnectorsGroupClient.GetSender(*http.Request) (*http.Response, error)
1. ConnectorsGroupClient.List(context.Context) (ConnectorSettingListPage, error)
1. ConnectorsGroupClient.ListComplete(context.Context) (ConnectorSettingListIterator, error)
1. ConnectorsGroupClient.ListPreparer(context.Context) (*http.Request, error)
1. ConnectorsGroupClient.ListResponder(*http.Response) (ConnectorSettingList, error)
1. ConnectorsGroupClient.ListSender(*http.Request) (*http.Response, error)
1. ConnectorsList.IsEmpty() bool
1. ConnectorsList.MarshalJSON() ([]byte, error)
1. ConnectorsListIterator.NotDone() bool
1. ConnectorsListIterator.Response() ConnectorsList
1. ConnectorsListIterator.Value() Connector
1. ConnectorsListPage.NotDone() bool
1. ConnectorsListPage.Response() ConnectorsList
1. ConnectorsListPage.Values() []Connector
1. MultiCloudOffering.AsBasicMultiCloudOffering() (BasicMultiCloudOffering, bool)
1. MultiCloudOffering.AsCSPMMonitorAWSOffering() (*CSPMMonitorAWSOffering, bool)
1. MultiCloudOffering.AsMultiCloudOffering() (*MultiCloudOffering, bool)
1. MultiCloudOffering.MarshalJSON() ([]byte, error)
1. NewConnectorsGroupClient(string, string) ConnectorsGroupClient
1. NewConnectorsGroupClientWithBaseURI(string, string, string) ConnectorsGroupClient
1. NewConnectorsListIterator(ConnectorsListPage) ConnectorsListIterator
1. NewConnectorsListPage(ConnectorsList, func(context.Context, ConnectorsList) (ConnectorsList, error)) ConnectorsListPage
1. PossibleMultiCloudNameValues() []MultiCloudName
1. PossibleOfferingTypeValues() []OfferingType

### Struct Changes

#### New Structs

1. CSPMMonitorAWSOffering
1. CSPMMonitorAWSOfferingNativeCloudConnection
1. Connector
1. ConnectorProperties
1. ConnectorsGroupClient
1. ConnectorsList
1. ConnectorsListIterator
1. ConnectorsListPage
1. MultiCloudOffering
