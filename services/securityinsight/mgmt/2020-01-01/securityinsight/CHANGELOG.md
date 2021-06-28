# Unreleased

## Additive Changes

### New Constants

1. DataConnectorKind.DataConnectorKindDynamics365
1. KindBasicDataConnector.KindBasicDataConnectorKindDynamics365

### New Funcs

1. *Dynamics365DataConnector.UnmarshalJSON([]byte) error
1. AADDataConnector.AsDynamics365DataConnector() (*Dynamics365DataConnector, bool)
1. AATPDataConnector.AsDynamics365DataConnector() (*Dynamics365DataConnector, bool)
1. ASCDataConnector.AsDynamics365DataConnector() (*Dynamics365DataConnector, bool)
1. AwsCloudTrailDataConnector.AsDynamics365DataConnector() (*Dynamics365DataConnector, bool)
1. DataConnector.AsDynamics365DataConnector() (*Dynamics365DataConnector, bool)
1. Dynamics365DataConnector.AsAADDataConnector() (*AADDataConnector, bool)
1. Dynamics365DataConnector.AsAATPDataConnector() (*AATPDataConnector, bool)
1. Dynamics365DataConnector.AsASCDataConnector() (*ASCDataConnector, bool)
1. Dynamics365DataConnector.AsAwsCloudTrailDataConnector() (*AwsCloudTrailDataConnector, bool)
1. Dynamics365DataConnector.AsBasicDataConnector() (BasicDataConnector, bool)
1. Dynamics365DataConnector.AsDataConnector() (*DataConnector, bool)
1. Dynamics365DataConnector.AsDynamics365DataConnector() (*Dynamics365DataConnector, bool)
1. Dynamics365DataConnector.AsMCASDataConnector() (*MCASDataConnector, bool)
1. Dynamics365DataConnector.AsMDATPDataConnector() (*MDATPDataConnector, bool)
1. Dynamics365DataConnector.AsOfficeDataConnector() (*OfficeDataConnector, bool)
1. Dynamics365DataConnector.AsTIDataConnector() (*TIDataConnector, bool)
1. Dynamics365DataConnector.MarshalJSON() ([]byte, error)
1. MCASDataConnector.AsDynamics365DataConnector() (*Dynamics365DataConnector, bool)
1. MDATPDataConnector.AsDynamics365DataConnector() (*Dynamics365DataConnector, bool)
1. OfficeDataConnector.AsDynamics365DataConnector() (*Dynamics365DataConnector, bool)
1. TIDataConnector.AsDynamics365DataConnector() (*Dynamics365DataConnector, bool)

### Struct Changes

#### New Structs

1. Dynamics365DataConnector
1. Dynamics365DataConnectorDataTypes
1. Dynamics365DataConnectorDataTypesDynamics365CdsActivities
1. Dynamics365DataConnectorProperties
