package resources

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// AliasPathAttributes enumerates the values for alias path attributes.
type AliasPathAttributes string

const (
	// Modifiable The token that the alias path is referring to is modifiable by policies with 'modify' effect.
	Modifiable AliasPathAttributes = "Modifiable"
	// None The token that the alias path is referring to has no attributes.
	None AliasPathAttributes = "None"
)

// PossibleAliasPathAttributesValues returns an array of possible values for the AliasPathAttributes const type.
func PossibleAliasPathAttributesValues() []AliasPathAttributes {
	return []AliasPathAttributes{Modifiable, None}
}

// AliasPathTokenType enumerates the values for alias path token type.
type AliasPathTokenType string

const (
	// Any The token type can be anything.
	Any AliasPathTokenType = "Any"
	// Array The token type is array.
	Array AliasPathTokenType = "Array"
	// Boolean The token type is boolean.
	Boolean AliasPathTokenType = "Boolean"
	// Integer The token type is integer.
	Integer AliasPathTokenType = "Integer"
	// NotSpecified The token type is not specified.
	NotSpecified AliasPathTokenType = "NotSpecified"
	// Number The token type is number.
	Number AliasPathTokenType = "Number"
	// Object The token type is object.
	Object AliasPathTokenType = "Object"
	// String The token type is string.
	String AliasPathTokenType = "String"
)

// PossibleAliasPathTokenTypeValues returns an array of possible values for the AliasPathTokenType const type.
func PossibleAliasPathTokenTypeValues() []AliasPathTokenType {
	return []AliasPathTokenType{Any, Array, Boolean, Integer, NotSpecified, Number, Object, String}
}

// AliasPatternType enumerates the values for alias pattern type.
type AliasPatternType string

const (
	// AliasPatternTypeExtract Extract is the only allowed value.
	AliasPatternTypeExtract AliasPatternType = "Extract"
	// AliasPatternTypeNotSpecified NotSpecified is not allowed.
	AliasPatternTypeNotSpecified AliasPatternType = "NotSpecified"
)

// PossibleAliasPatternTypeValues returns an array of possible values for the AliasPatternType const type.
func PossibleAliasPatternTypeValues() []AliasPatternType {
	return []AliasPatternType{AliasPatternTypeExtract, AliasPatternTypeNotSpecified}
}

// AliasType enumerates the values for alias type.
type AliasType string

const (
	// AliasTypeMask Alias value is secret.
	AliasTypeMask AliasType = "Mask"
	// AliasTypeNotSpecified Alias type is unknown (same as not providing alias type).
	AliasTypeNotSpecified AliasType = "NotSpecified"
	// AliasTypePlainText Alias value is not secret.
	AliasTypePlainText AliasType = "PlainText"
)

// PossibleAliasTypeValues returns an array of possible values for the AliasType const type.
func PossibleAliasTypeValues() []AliasType {
	return []AliasType{AliasTypeMask, AliasTypeNotSpecified, AliasTypePlainText}
}

// ChangeType enumerates the values for change type.
type ChangeType string

const (
	// Create The resource does not exist in the current state but is present in the desired state. The
	// resource will be created when the deployment is executed.
	Create ChangeType = "Create"
	// Delete The resource exists in the current state and is missing from the desired state. The resource will
	// be deleted when the deployment is executed.
	Delete ChangeType = "Delete"
	// Deploy The resource exists in the current state and the desired state and will be redeployed when the
	// deployment is executed. The properties of the resource may or may not change.
	Deploy ChangeType = "Deploy"
	// Ignore The resource exists in the current state and is missing from the desired state. The resource will
	// not be deployed or modified when the deployment is executed.
	Ignore ChangeType = "Ignore"
	// Modify The resource exists in the current state and the desired state and will be redeployed when the
	// deployment is executed. The properties of the resource will change.
	Modify ChangeType = "Modify"
	// NoChange The resource exists in the current state and the desired state and will be redeployed when the
	// deployment is executed. The properties of the resource will not change.
	NoChange ChangeType = "NoChange"
)

// PossibleChangeTypeValues returns an array of possible values for the ChangeType const type.
func PossibleChangeTypeValues() []ChangeType {
	return []ChangeType{Create, Delete, Deploy, Ignore, Modify, NoChange}
}

// DeploymentMode enumerates the values for deployment mode.
type DeploymentMode string

const (
	// Complete ...
	Complete DeploymentMode = "Complete"
	// Incremental ...
	Incremental DeploymentMode = "Incremental"
)

// PossibleDeploymentModeValues returns an array of possible values for the DeploymentMode const type.
func PossibleDeploymentModeValues() []DeploymentMode {
	return []DeploymentMode{Complete, Incremental}
}

// OnErrorDeploymentType enumerates the values for on error deployment type.
type OnErrorDeploymentType string

const (
	// LastSuccessful ...
	LastSuccessful OnErrorDeploymentType = "LastSuccessful"
	// SpecificDeployment ...
	SpecificDeployment OnErrorDeploymentType = "SpecificDeployment"
)

// PossibleOnErrorDeploymentTypeValues returns an array of possible values for the OnErrorDeploymentType const type.
func PossibleOnErrorDeploymentTypeValues() []OnErrorDeploymentType {
	return []OnErrorDeploymentType{LastSuccessful, SpecificDeployment}
}

// PropertyChangeType enumerates the values for property change type.
type PropertyChangeType string

const (
	// PropertyChangeTypeArray The property is an array and contains nested changes.
	PropertyChangeTypeArray PropertyChangeType = "Array"
	// PropertyChangeTypeCreate The property does not exist in the current state but is present in the desired
	// state. The property will be created when the deployment is executed.
	PropertyChangeTypeCreate PropertyChangeType = "Create"
	// PropertyChangeTypeDelete The property exists in the current state and is missing from the desired state.
	// It will be deleted when the deployment is executed.
	PropertyChangeTypeDelete PropertyChangeType = "Delete"
	// PropertyChangeTypeModify The property exists in both current and desired state and is different. The
	// value of the property will change when the deployment is executed.
	PropertyChangeTypeModify PropertyChangeType = "Modify"
)

// PossiblePropertyChangeTypeValues returns an array of possible values for the PropertyChangeType const type.
func PossiblePropertyChangeTypeValues() []PropertyChangeType {
	return []PropertyChangeType{PropertyChangeTypeArray, PropertyChangeTypeCreate, PropertyChangeTypeDelete, PropertyChangeTypeModify}
}

// ProvisioningOperation enumerates the values for provisioning operation.
type ProvisioningOperation string

const (
	// ProvisioningOperationAction The provisioning operation is action.
	ProvisioningOperationAction ProvisioningOperation = "Action"
	// ProvisioningOperationAzureAsyncOperationWaiting The provisioning operation is waiting Azure async
	// operation.
	ProvisioningOperationAzureAsyncOperationWaiting ProvisioningOperation = "AzureAsyncOperationWaiting"
	// ProvisioningOperationCreate The provisioning operation is create.
	ProvisioningOperationCreate ProvisioningOperation = "Create"
	// ProvisioningOperationDelete The provisioning operation is delete.
	ProvisioningOperationDelete ProvisioningOperation = "Delete"
	// ProvisioningOperationDeploymentCleanup The provisioning operation is cleanup. This operation is part of
	// the 'complete' mode deployment.
	ProvisioningOperationDeploymentCleanup ProvisioningOperation = "DeploymentCleanup"
	// ProvisioningOperationEvaluateDeploymentOutput The provisioning operation is evaluate output.
	ProvisioningOperationEvaluateDeploymentOutput ProvisioningOperation = "EvaluateDeploymentOutput"
	// ProvisioningOperationNotSpecified The provisioning operation is not specified.
	ProvisioningOperationNotSpecified ProvisioningOperation = "NotSpecified"
	// ProvisioningOperationRead The provisioning operation is read.
	ProvisioningOperationRead ProvisioningOperation = "Read"
	// ProvisioningOperationResourceCacheWaiting The provisioning operation is waiting for resource cache.
	ProvisioningOperationResourceCacheWaiting ProvisioningOperation = "ResourceCacheWaiting"
	// ProvisioningOperationWaiting The provisioning operation is waiting.
	ProvisioningOperationWaiting ProvisioningOperation = "Waiting"
)

// PossibleProvisioningOperationValues returns an array of possible values for the ProvisioningOperation const type.
func PossibleProvisioningOperationValues() []ProvisioningOperation {
	return []ProvisioningOperation{ProvisioningOperationAction, ProvisioningOperationAzureAsyncOperationWaiting, ProvisioningOperationCreate, ProvisioningOperationDelete, ProvisioningOperationDeploymentCleanup, ProvisioningOperationEvaluateDeploymentOutput, ProvisioningOperationNotSpecified, ProvisioningOperationRead, ProvisioningOperationResourceCacheWaiting, ProvisioningOperationWaiting}
}

// ProvisioningState enumerates the values for provisioning state.
type ProvisioningState string

const (
	// ProvisioningStateAccepted ...
	ProvisioningStateAccepted ProvisioningState = "Accepted"
	// ProvisioningStateCanceled ...
	ProvisioningStateCanceled ProvisioningState = "Canceled"
	// ProvisioningStateCreated ...
	ProvisioningStateCreated ProvisioningState = "Created"
	// ProvisioningStateCreating ...
	ProvisioningStateCreating ProvisioningState = "Creating"
	// ProvisioningStateDeleted ...
	ProvisioningStateDeleted ProvisioningState = "Deleted"
	// ProvisioningStateDeleting ...
	ProvisioningStateDeleting ProvisioningState = "Deleting"
	// ProvisioningStateFailed ...
	ProvisioningStateFailed ProvisioningState = "Failed"
	// ProvisioningStateNotSpecified ...
	ProvisioningStateNotSpecified ProvisioningState = "NotSpecified"
	// ProvisioningStateReady ...
	ProvisioningStateReady ProvisioningState = "Ready"
	// ProvisioningStateRunning ...
	ProvisioningStateRunning ProvisioningState = "Running"
	// ProvisioningStateSucceeded ...
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
	// ProvisioningStateUpdating ...
	ProvisioningStateUpdating ProvisioningState = "Updating"
)

// PossibleProvisioningStateValues returns an array of possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{ProvisioningStateAccepted, ProvisioningStateCanceled, ProvisioningStateCreated, ProvisioningStateCreating, ProvisioningStateDeleted, ProvisioningStateDeleting, ProvisioningStateFailed, ProvisioningStateNotSpecified, ProvisioningStateReady, ProvisioningStateRunning, ProvisioningStateSucceeded, ProvisioningStateUpdating}
}

// ResourceIdentityType enumerates the values for resource identity type.
type ResourceIdentityType string

const (
	// ResourceIdentityTypeNone ...
	ResourceIdentityTypeNone ResourceIdentityType = "None"
	// ResourceIdentityTypeSystemAssigned ...
	ResourceIdentityTypeSystemAssigned ResourceIdentityType = "SystemAssigned"
	// ResourceIdentityTypeSystemAssignedUserAssigned ...
	ResourceIdentityTypeSystemAssignedUserAssigned ResourceIdentityType = "SystemAssigned, UserAssigned"
	// ResourceIdentityTypeUserAssigned ...
	ResourceIdentityTypeUserAssigned ResourceIdentityType = "UserAssigned"
)

// PossibleResourceIdentityTypeValues returns an array of possible values for the ResourceIdentityType const type.
func PossibleResourceIdentityTypeValues() []ResourceIdentityType {
	return []ResourceIdentityType{ResourceIdentityTypeNone, ResourceIdentityTypeSystemAssigned, ResourceIdentityTypeSystemAssignedUserAssigned, ResourceIdentityTypeUserAssigned}
}

// TagsPatchOperation enumerates the values for tags patch operation.
type TagsPatchOperation string

const (
	// TagsPatchOperationDelete The 'delete' option allows selectively deleting tags based on given names or
	// name/value pairs.
	TagsPatchOperationDelete TagsPatchOperation = "Delete"
	// TagsPatchOperationMerge The 'merge' option allows adding tags with new names and updating the values of
	// tags with existing names.
	TagsPatchOperationMerge TagsPatchOperation = "Merge"
	// TagsPatchOperationReplace The 'replace' option replaces the entire set of existing tags with a new set.
	TagsPatchOperationReplace TagsPatchOperation = "Replace"
)

// PossibleTagsPatchOperationValues returns an array of possible values for the TagsPatchOperation const type.
func PossibleTagsPatchOperationValues() []TagsPatchOperation {
	return []TagsPatchOperation{TagsPatchOperationDelete, TagsPatchOperationMerge, TagsPatchOperationReplace}
}

// WhatIfResultFormat enumerates the values for what if result format.
type WhatIfResultFormat string

const (
	// FullResourcePayloads ...
	FullResourcePayloads WhatIfResultFormat = "FullResourcePayloads"
	// ResourceIDOnly ...
	ResourceIDOnly WhatIfResultFormat = "ResourceIdOnly"
)

// PossibleWhatIfResultFormatValues returns an array of possible values for the WhatIfResultFormat const type.
func PossibleWhatIfResultFormatValues() []WhatIfResultFormat {
	return []WhatIfResultFormat{FullResourcePayloads, ResourceIDOnly}
}
