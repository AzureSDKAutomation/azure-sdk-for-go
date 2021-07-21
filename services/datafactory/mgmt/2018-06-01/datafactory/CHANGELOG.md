# Unreleased

## Breaking Changes

### Removed Constants

1. IntegrationRuntimeEntityReferenceType.IntegrationRuntimeEntityReferenceTypeCredentialReference

### Struct Changes

#### Removed Struct Fields

1. IntegrationRuntimeSsisProperties.ManagedCredential

### Signature Changes

#### Struct Fields

1. AvroDatasetTypeProperties.AvroCompressionCodec changed type from AvroCompressionCodec to interface{}
1. OrcDatasetTypeProperties.OrcCompressionCodec changed type from OrcCompressionCodec to interface{}

## Additive Changes

### New Constants

1. TypeBasicCredential.TypeBasicCredentialTypeCredential
1. TypeBasicCredential.TypeBasicCredentialTypeManagedIdentity
1. TypeBasicCredential.TypeBasicCredentialTypeServicePrincipal

### New Funcs

1. *Credential.UnmarshalJSON([]byte) error
1. *CredentialReference.UnmarshalJSON([]byte) error
1. *CredentialResource.UnmarshalJSON([]byte) error
1. *ManagedIdentityCredential.UnmarshalJSON([]byte) error
1. *ServicePrincipalCredential.UnmarshalJSON([]byte) error
1. Credential.AsBasicCredential() (BasicCredential, bool)
1. Credential.AsCredential() (*Credential, bool)
1. Credential.AsManagedIdentityCredential() (*ManagedIdentityCredential, bool)
1. Credential.AsServicePrincipalCredential() (*ServicePrincipalCredential, bool)
1. Credential.MarshalJSON() ([]byte, error)
1. CredentialReference.MarshalJSON() ([]byte, error)
1. CredentialResource.MarshalJSON() ([]byte, error)
1. ManagedIdentityCredential.AsBasicCredential() (BasicCredential, bool)
1. ManagedIdentityCredential.AsCredential() (*Credential, bool)
1. ManagedIdentityCredential.AsManagedIdentityCredential() (*ManagedIdentityCredential, bool)
1. ManagedIdentityCredential.AsServicePrincipalCredential() (*ServicePrincipalCredential, bool)
1. ManagedIdentityCredential.MarshalJSON() ([]byte, error)
1. PossibleTypeBasicCredentialValues() []TypeBasicCredential
1. ServicePrincipalCredential.AsBasicCredential() (BasicCredential, bool)
1. ServicePrincipalCredential.AsCredential() (*Credential, bool)
1. ServicePrincipalCredential.AsManagedIdentityCredential() (*ManagedIdentityCredential, bool)
1. ServicePrincipalCredential.AsServicePrincipalCredential() (*ServicePrincipalCredential, bool)
1. ServicePrincipalCredential.MarshalJSON() ([]byte, error)

### Struct Changes

#### New Structs

1. Credential
1. CredentialReference
1. CredentialResource
1. GitHubClientSecret
1. ManagedIdentityCredential
1. ManagedIdentityTypeProperties
1. ServicePrincipalCredential
1. ServicePrincipalCredentialTypeProperties

#### New Struct Fields

1. AzureBatchLinkedServiceTypeProperties.Credential
1. AzureBlobFSLinkedServiceTypeProperties.Credential
1. AzureBlobStorageLinkedServiceTypeProperties.Credential
1. AzureDataExplorerLinkedServiceTypeProperties.Credential
1. AzureDataLakeStoreLinkedServiceTypeProperties.Credential
1. AzureDatabricksLinkedServiceTypeProperties.Credential
1. AzureFunctionLinkedServiceTypeProperties.Authentication
1. AzureFunctionLinkedServiceTypeProperties.Credential
1. AzureFunctionLinkedServiceTypeProperties.ResourceID
1. AzureKeyVaultLinkedServiceTypeProperties.Credential
1. AzureMLLinkedServiceTypeProperties.Authentication
1. AzureSQLDWLinkedServiceTypeProperties.Credential
1. AzureSQLDatabaseLinkedServiceTypeProperties.Credential
1. AzureSQLMILinkedServiceTypeProperties.Credential
1. FactoryGitHubConfiguration.ClientID
1. FactoryGitHubConfiguration.ClientSecret
1. GitHubAccessTokenRequest.GitHubClientSecret
1. HDInsightOnDemandLinkedServiceTypeProperties.Credential
1. IntegrationRuntimeSsisProperties.Credential
1. PipelineRunInvokedBy.PipelineName
1. PipelineRunInvokedBy.PipelineRunID
1. RestServiceLinkedServiceTypeProperties.Credential
1. WebActivityAuthentication.Credential
