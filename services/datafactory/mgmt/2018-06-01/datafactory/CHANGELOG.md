# Unreleased

## Breaking Changes

### Signature Changes

#### Struct Fields

1. AvroDatasetTypeProperties.AvroCompressionCodec changed type from AvroCompressionCodec to interface{}
1. OrcDatasetTypeProperties.OrcCompressionCodec changed type from OrcCompressionCodec to interface{}

## Additive Changes

### New Constants

1. SQLDWWriteBehaviorEnum.SQLDWWriteBehaviorEnumInsert
1. SQLDWWriteBehaviorEnum.SQLDWWriteBehaviorEnumUpsert
1. SQLWriteBehaviorEnum.SQLWriteBehaviorEnumInsert
1. SQLWriteBehaviorEnum.SQLWriteBehaviorEnumStoredProcedure
1. SQLWriteBehaviorEnum.SQLWriteBehaviorEnumUpsert

### New Funcs

1. PossibleSQLDWWriteBehaviorEnumValues() []SQLDWWriteBehaviorEnum
1. PossibleSQLWriteBehaviorEnumValues() []SQLWriteBehaviorEnum

### Struct Changes

#### New Structs

1. SQLDWUpsertSettings
1. SQLUpsertSettings

#### New Struct Fields

1. AzureSQLSink.SQLWriterUseTableLock
1. AzureSQLSink.UpsertSettings
1. AzureSQLSink.WriteBehavior
1. PipelineRunInvokedBy.PipelineName
1. PipelineRunInvokedBy.PipelineRunID
1. SQLDWSink.SQLWriterUseTableLock
1. SQLDWSink.UpsertSettings
1. SQLDWSink.WriteBehavior
1. SQLMISink.SQLWriterUseTableLock
1. SQLMISink.UpsertSettings
1. SQLMISink.WriteBehavior
1. SQLSink.SQLWriterUseTableLock
1. SQLSink.UpsertSettings
1. SQLSink.WriteBehavior
