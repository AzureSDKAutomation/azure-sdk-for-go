# Unreleased

## Breaking Changes

### Signature Changes

#### Funcs

1. TrainClient.TrainVersion
	- Params
		- From: context.Context, uuid.UUID, string
		- To: context.Context, uuid.UUID, string, string
1. TrainClient.TrainVersionPreparer
	- Params
		- From: context.Context, uuid.UUID, string
		- To: context.Context, uuid.UUID, string, string

#### Struct Fields

1. LabelExampleResponse.ExampleID changed type from *int32 to *int64
