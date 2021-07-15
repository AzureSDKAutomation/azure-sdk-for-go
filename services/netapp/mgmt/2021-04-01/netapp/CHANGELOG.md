# Unreleased

## Additive Changes

### New Constants

1. EncryptionType.EncryptionTypeDouble
1. EncryptionType.EncryptionTypeSingle
1. MetricAggregationType.MetricAggregationTypeAverage

### New Funcs

1. AccountsClient.ListBySubscription(context.Context) (AccountListPage, error)
1. AccountsClient.ListBySubscriptionComplete(context.Context) (AccountListIterator, error)
1. AccountsClient.ListBySubscriptionPreparer(context.Context) (*http.Request, error)
1. AccountsClient.ListBySubscriptionResponder(*http.Response) (AccountList, error)
1. AccountsClient.ListBySubscriptionSender(*http.Request) (*http.Response, error)
1. PossibleEncryptionTypeValues() []EncryptionType
1. PossibleMetricAggregationTypeValues() []MetricAggregationType
1. SnapshotPoliciesClient.ListVolumesOldVersion(context.Context, string, string, string) (SnapshotPolicyVolumeList, error)
1. SnapshotPoliciesClient.ListVolumesOldVersionPreparer(context.Context, string, string, string) (*http.Request, error)
1. SnapshotPoliciesClient.ListVolumesOldVersionResponder(*http.Response) (SnapshotPolicyVolumeList, error)
1. SnapshotPoliciesClient.ListVolumesOldVersionSender(*http.Request) (*http.Response, error)

### Struct Changes

#### New Struct Fields

1. Account.Etag
1. BackupPolicy.Etag
1. BackupPolicyProperties.BackupPolicyID
1. CapacityPool.Etag
1. MetricSpecification.InternalMetricName
1. MetricSpecification.SourceMdmAccount
1. MetricSpecification.SourceMdmNamespace
1. MetricSpecification.SupportedAggregationTypes
1. MetricSpecification.SupportedTimeGrainTypes
1. PoolProperties.EncryptionType
1. SnapshotPolicy.Etag
1. Volume.Etag
