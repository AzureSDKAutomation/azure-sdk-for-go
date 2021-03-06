package hanaonazure

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// HanaHardwareTypeNamesEnum enumerates the values for hana hardware type names enum.
type HanaHardwareTypeNamesEnum string

const (
	// CiscoUCS ...
	CiscoUCS HanaHardwareTypeNamesEnum = "Cisco_UCS"
	// HPE ...
	HPE HanaHardwareTypeNamesEnum = "HPE"
)

// PossibleHanaHardwareTypeNamesEnumValues returns an array of possible values for the HanaHardwareTypeNamesEnum const type.
func PossibleHanaHardwareTypeNamesEnumValues() []HanaHardwareTypeNamesEnum {
	return []HanaHardwareTypeNamesEnum{CiscoUCS, HPE}
}

// HanaInstancePowerStateEnum enumerates the values for hana instance power state enum.
type HanaInstancePowerStateEnum string

const (
	// Restarting ...
	Restarting HanaInstancePowerStateEnum = "restarting"
	// Started ...
	Started HanaInstancePowerStateEnum = "started"
	// Starting ...
	Starting HanaInstancePowerStateEnum = "starting"
	// Stopped ...
	Stopped HanaInstancePowerStateEnum = "stopped"
	// Stopping ...
	Stopping HanaInstancePowerStateEnum = "stopping"
	// Unknown ...
	Unknown HanaInstancePowerStateEnum = "unknown"
)

// PossibleHanaInstancePowerStateEnumValues returns an array of possible values for the HanaInstancePowerStateEnum const type.
func PossibleHanaInstancePowerStateEnumValues() []HanaInstancePowerStateEnum {
	return []HanaInstancePowerStateEnum{Restarting, Started, Starting, Stopped, Stopping, Unknown}
}

// HanaInstanceSizeNamesEnum enumerates the values for hana instance size names enum.
type HanaInstanceSizeNamesEnum string

const (
	// S112 ...
	S112 HanaInstanceSizeNamesEnum = "S112"
	// S144 ...
	S144 HanaInstanceSizeNamesEnum = "S144"
	// S144m ...
	S144m HanaInstanceSizeNamesEnum = "S144m"
	// S192 ...
	S192 HanaInstanceSizeNamesEnum = "S192"
	// S192m ...
	S192m HanaInstanceSizeNamesEnum = "S192m"
	// S192xm ...
	S192xm HanaInstanceSizeNamesEnum = "S192xm"
	// S224 ...
	S224 HanaInstanceSizeNamesEnum = "S224"
	// S224m ...
	S224m HanaInstanceSizeNamesEnum = "S224m"
	// S224om ...
	S224om HanaInstanceSizeNamesEnum = "S224om"
	// S224oo ...
	S224oo HanaInstanceSizeNamesEnum = "S224oo"
	// S224oom ...
	S224oom HanaInstanceSizeNamesEnum = "S224oom"
	// S224ooo ...
	S224ooo HanaInstanceSizeNamesEnum = "S224ooo"
	// S384 ...
	S384 HanaInstanceSizeNamesEnum = "S384"
	// S384m ...
	S384m HanaInstanceSizeNamesEnum = "S384m"
	// S384xm ...
	S384xm HanaInstanceSizeNamesEnum = "S384xm"
	// S384xxm ...
	S384xxm HanaInstanceSizeNamesEnum = "S384xxm"
	// S448 ...
	S448 HanaInstanceSizeNamesEnum = "S448"
	// S448m ...
	S448m HanaInstanceSizeNamesEnum = "S448m"
	// S448om ...
	S448om HanaInstanceSizeNamesEnum = "S448om"
	// S448oo ...
	S448oo HanaInstanceSizeNamesEnum = "S448oo"
	// S448oom ...
	S448oom HanaInstanceSizeNamesEnum = "S448oom"
	// S448ooo ...
	S448ooo HanaInstanceSizeNamesEnum = "S448ooo"
	// S576m ...
	S576m HanaInstanceSizeNamesEnum = "S576m"
	// S576xm ...
	S576xm HanaInstanceSizeNamesEnum = "S576xm"
	// S672 ...
	S672 HanaInstanceSizeNamesEnum = "S672"
	// S672m ...
	S672m HanaInstanceSizeNamesEnum = "S672m"
	// S672om ...
	S672om HanaInstanceSizeNamesEnum = "S672om"
	// S672oo ...
	S672oo HanaInstanceSizeNamesEnum = "S672oo"
	// S672oom ...
	S672oom HanaInstanceSizeNamesEnum = "S672oom"
	// S672ooo ...
	S672ooo HanaInstanceSizeNamesEnum = "S672ooo"
	// S72 ...
	S72 HanaInstanceSizeNamesEnum = "S72"
	// S72m ...
	S72m HanaInstanceSizeNamesEnum = "S72m"
	// S768 ...
	S768 HanaInstanceSizeNamesEnum = "S768"
	// S768m ...
	S768m HanaInstanceSizeNamesEnum = "S768m"
	// S768xm ...
	S768xm HanaInstanceSizeNamesEnum = "S768xm"
	// S896 ...
	S896 HanaInstanceSizeNamesEnum = "S896"
	// S896m ...
	S896m HanaInstanceSizeNamesEnum = "S896m"
	// S896om ...
	S896om HanaInstanceSizeNamesEnum = "S896om"
	// S896oo ...
	S896oo HanaInstanceSizeNamesEnum = "S896oo"
	// S896oom ...
	S896oom HanaInstanceSizeNamesEnum = "S896oom"
	// S896ooo ...
	S896ooo HanaInstanceSizeNamesEnum = "S896ooo"
	// S96 ...
	S96 HanaInstanceSizeNamesEnum = "S96"
	// S960m ...
	S960m HanaInstanceSizeNamesEnum = "S960m"
)

// PossibleHanaInstanceSizeNamesEnumValues returns an array of possible values for the HanaInstanceSizeNamesEnum const type.
func PossibleHanaInstanceSizeNamesEnumValues() []HanaInstanceSizeNamesEnum {
	return []HanaInstanceSizeNamesEnum{S112, S144, S144m, S192, S192m, S192xm, S224, S224m, S224om, S224oo, S224oom, S224ooo, S384, S384m, S384xm, S384xxm, S448, S448m, S448om, S448oo, S448oom, S448ooo, S576m, S576xm, S672, S672m, S672om, S672oo, S672oom, S672ooo, S72, S72m, S768, S768m, S768xm, S896, S896m, S896om, S896oo, S896oom, S896ooo, S96, S960m}
}

// HanaProvisioningStatesEnum enumerates the values for hana provisioning states enum.
type HanaProvisioningStatesEnum string

const (
	// Accepted ...
	Accepted HanaProvisioningStatesEnum = "Accepted"
	// Creating ...
	Creating HanaProvisioningStatesEnum = "Creating"
	// Deleting ...
	Deleting HanaProvisioningStatesEnum = "Deleting"
	// Failed ...
	Failed HanaProvisioningStatesEnum = "Failed"
	// Migrating ...
	Migrating HanaProvisioningStatesEnum = "Migrating"
	// Succeeded ...
	Succeeded HanaProvisioningStatesEnum = "Succeeded"
	// Updating ...
	Updating HanaProvisioningStatesEnum = "Updating"
)

// PossibleHanaProvisioningStatesEnumValues returns an array of possible values for the HanaProvisioningStatesEnum const type.
func PossibleHanaProvisioningStatesEnumValues() []HanaProvisioningStatesEnum {
	return []HanaProvisioningStatesEnum{Accepted, Creating, Deleting, Failed, Migrating, Succeeded, Updating}
}
