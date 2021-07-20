# Unreleased

## Breaking Changes

### Struct Changes

#### Removed Struct Fields

1. BaseClient.ProviderNamespace

### Signature Changes

#### Funcs

1. New
	- Params
		- From: string, string
		- To: string
1. NewClient
	- Params
		- From: string, string
		- To: string
1. NewClientWithBaseURI
	- Params
		- From: string, string, string
		- To: string, string
1. NewSubscriptionFeatureRegistrationsClient
	- Params
		- From: string, string
		- To: string
1. NewSubscriptionFeatureRegistrationsClientWithBaseURI
	- Params
		- From: string, string, string
		- To: string, string
1. NewWithBaseURI
	- Params
		- From: string, string, string
		- To: string, string
1. SubscriptionFeatureRegistrationsClient.CreateOrUpdate
	- Params
		- From: context.Context, string, string, *SubscriptionFeatureRegistration
		- To: context.Context, string, string, string, *SubscriptionFeatureRegistration
1. SubscriptionFeatureRegistrationsClient.CreateOrUpdatePreparer
	- Params
		- From: context.Context, string, string, *SubscriptionFeatureRegistration
		- To: context.Context, string, string, string, *SubscriptionFeatureRegistration
1. SubscriptionFeatureRegistrationsClient.Delete
	- Params
		- From: context.Context, string, string
		- To: context.Context, string, string, string
1. SubscriptionFeatureRegistrationsClient.DeletePreparer
	- Params
		- From: context.Context, string, string
		- To: context.Context, string, string, string
1. SubscriptionFeatureRegistrationsClient.Get
	- Params
		- From: context.Context, string, string
		- To: context.Context, string, string, string
1. SubscriptionFeatureRegistrationsClient.GetPreparer
	- Params
		- From: context.Context, string, string
		- To: context.Context, string, string, string
1. SubscriptionFeatureRegistrationsClient.ListBySubscription
	- Params
		- From: context.Context, string
		- To: context.Context, string, string
1. SubscriptionFeatureRegistrationsClient.ListBySubscriptionComplete
	- Params
		- From: context.Context, string
		- To: context.Context, string, string
1. SubscriptionFeatureRegistrationsClient.ListBySubscriptionPreparer
	- Params
		- From: context.Context, string
		- To: context.Context, string, string
