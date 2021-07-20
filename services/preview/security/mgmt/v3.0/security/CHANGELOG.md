# Unreleased

## Breaking Changes

### Removed Funcs

1. *AssessmentMetadataListIterator.Next() error
1. *AssessmentMetadataListIterator.NextWithContext(context.Context) error
1. *AssessmentMetadataListPage.Next() error
1. *AssessmentMetadataListPage.NextWithContext(context.Context) error
1. AssessmentMetadataList.IsEmpty() bool
1. AssessmentMetadataList.MarshalJSON() ([]byte, error)
1. AssessmentMetadataListIterator.NotDone() bool
1. AssessmentMetadataListIterator.Response() AssessmentMetadataList
1. AssessmentMetadataListIterator.Value() AssessmentMetadata
1. AssessmentMetadataListPage.NotDone() bool
1. AssessmentMetadataListPage.Response() AssessmentMetadataList
1. AssessmentMetadataListPage.Values() []AssessmentMetadata
1. NewAssessmentMetadataListIterator(AssessmentMetadataListPage) AssessmentMetadataListIterator
1. NewAssessmentMetadataListPage(AssessmentMetadataList, func(context.Context, AssessmentMetadataList) (AssessmentMetadataList, error)) AssessmentMetadataListPage

### Struct Changes

#### Removed Structs

1. AssessmentMetadataList
1. AssessmentMetadataListIterator
1. AssessmentMetadataListPage

#### Removed Struct Fields

1. Assessment.autorest.Response
1. AssessmentMetadata.autorest.Response

### Signature Changes

#### Funcs

1. AssessmentListIterator.Value
	- Returns
		- From: Assessment
		- To: AssessmentResponse
1. AssessmentListPage.Values
	- Returns
		- From: []Assessment
		- To: []AssessmentResponse
1. AssessmentsClient.CreateOrUpdate
	- Returns
		- From: Assessment, error
		- To: AssessmentResponse, error
1. AssessmentsClient.CreateOrUpdateResponder
	- Returns
		- From: Assessment, error
		- To: AssessmentResponse, error
1. AssessmentsClient.Get
	- Returns
		- From: Assessment, error
		- To: AssessmentResponse, error
1. AssessmentsClient.GetResponder
	- Returns
		- From: Assessment, error
		- To: AssessmentResponse, error
1. AssessmentsMetadataClient.CreateInSubscription
	- Returns
		- From: AssessmentMetadata, error
		- To: AssessmentMetadataResponse, error
1. AssessmentsMetadataClient.CreateInSubscriptionResponder
	- Returns
		- From: AssessmentMetadata, error
		- To: AssessmentMetadataResponse, error
1. AssessmentsMetadataClient.Get
	- Returns
		- From: AssessmentMetadata, error
		- To: AssessmentMetadataResponse, error
1. AssessmentsMetadataClient.GetInSubscription
	- Returns
		- From: AssessmentMetadata, error
		- To: AssessmentMetadataResponse, error
1. AssessmentsMetadataClient.GetInSubscriptionResponder
	- Returns
		- From: AssessmentMetadata, error
		- To: AssessmentMetadataResponse, error
1. AssessmentsMetadataClient.GetResponder
	- Returns
		- From: AssessmentMetadata, error
		- To: AssessmentMetadataResponse, error
1. AssessmentsMetadataClient.List
	- Returns
		- From: AssessmentMetadataListPage, error
		- To: AssessmentMetadataResponseListPage, error
1. AssessmentsMetadataClient.ListBySubscription
	- Returns
		- From: AssessmentMetadataListPage, error
		- To: AssessmentMetadataResponseListPage, error
1. AssessmentsMetadataClient.ListBySubscriptionComplete
	- Returns
		- From: AssessmentMetadataListIterator, error
		- To: AssessmentMetadataResponseListIterator, error
1. AssessmentsMetadataClient.ListBySubscriptionResponder
	- Returns
		- From: AssessmentMetadataList, error
		- To: AssessmentMetadataResponseList, error
1. AssessmentsMetadataClient.ListComplete
	- Returns
		- From: AssessmentMetadataListIterator, error
		- To: AssessmentMetadataResponseListIterator, error
1. AssessmentsMetadataClient.ListResponder
	- Returns
		- From: AssessmentMetadataList, error
		- To: AssessmentMetadataResponseList, error

#### Struct Fields

1. AssessmentList.Value changed type from *[]Assessment to *[]AssessmentResponse

## Additive Changes

### New Constants

1. Tactics.TacticsCollection
1. Tactics.TacticsCommandAndControl
1. Tactics.TacticsCredentialAccess
1. Tactics.TacticsDefenseEvasion
1. Tactics.TacticsDiscovery
1. Tactics.TacticsExecution
1. Tactics.TacticsExfiltration
1. Tactics.TacticsImpact
1. Tactics.TacticsInitialAccess
1. Tactics.TacticsLateralMovement
1. Tactics.TacticsPersistence
1. Tactics.TacticsPrivilegeEscalation
1. Tactics.TacticsReconnaissance
1. Tactics.TacticsResourceDevelopment

### New Funcs

1. *AssessmentMetadataResponse.UnmarshalJSON([]byte) error
1. *AssessmentMetadataResponseListIterator.Next() error
1. *AssessmentMetadataResponseListIterator.NextWithContext(context.Context) error
1. *AssessmentMetadataResponseListPage.Next() error
1. *AssessmentMetadataResponseListPage.NextWithContext(context.Context) error
1. *AssessmentPropertiesBase.UnmarshalJSON([]byte) error
1. *AssessmentPropertiesResponse.UnmarshalJSON([]byte) error
1. *AssessmentResponse.UnmarshalJSON([]byte) error
1. AssessmentMetadataPropertiesResponse.MarshalJSON() ([]byte, error)
1. AssessmentMetadataResponse.MarshalJSON() ([]byte, error)
1. AssessmentMetadataResponseList.IsEmpty() bool
1. AssessmentMetadataResponseList.MarshalJSON() ([]byte, error)
1. AssessmentMetadataResponseListIterator.NotDone() bool
1. AssessmentMetadataResponseListIterator.Response() AssessmentMetadataResponseList
1. AssessmentMetadataResponseListIterator.Value() AssessmentMetadataResponse
1. AssessmentMetadataResponseListPage.NotDone() bool
1. AssessmentMetadataResponseListPage.Response() AssessmentMetadataResponseList
1. AssessmentMetadataResponseListPage.Values() []AssessmentMetadataResponse
1. AssessmentPropertiesBase.MarshalJSON() ([]byte, error)
1. AssessmentPropertiesResponse.MarshalJSON() ([]byte, error)
1. AssessmentResponse.MarshalJSON() ([]byte, error)
1. AssessmentStatusResponse.MarshalJSON() ([]byte, error)
1. NewAssessmentMetadataResponseListIterator(AssessmentMetadataResponseListPage) AssessmentMetadataResponseListIterator
1. NewAssessmentMetadataResponseListPage(AssessmentMetadataResponseList, func(context.Context, AssessmentMetadataResponseList) (AssessmentMetadataResponseList, error)) AssessmentMetadataResponseListPage
1. PossibleTacticsValues() []Tactics

### Struct Changes

#### New Structs

1. AssessmentMetadataPropertiesResponse
1. AssessmentMetadataPropertiesResponsePublishDates
1. AssessmentMetadataResponse
1. AssessmentMetadataResponseList
1. AssessmentMetadataResponseListIterator
1. AssessmentMetadataResponseListPage
1. AssessmentPropertiesBase
1. AssessmentPropertiesResponse
1. AssessmentResponse
1. AssessmentStatusResponse
