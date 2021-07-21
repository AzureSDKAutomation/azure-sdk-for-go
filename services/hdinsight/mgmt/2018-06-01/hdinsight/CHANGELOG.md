# Unreleased

## Breaking Changes

### Removed Funcs

1. ClusterIdentityUserAssignedIdentitiesValue.MarshalJSON() ([]byte, error)

### Struct Changes

#### Removed Structs

1. ClusterIdentityUserAssignedIdentitiesValue

### Signature Changes

#### Struct Fields

1. ClusterIdentity.UserAssignedIdentities changed type from map[string]*ClusterIdentityUserAssignedIdentitiesValue to map[string]*UserAssignedIdentity

## Additive Changes

### New Funcs

1. UserAssignedIdentity.MarshalJSON() ([]byte, error)

### Struct Changes

#### New Structs

1. UserAssignedIdentity
