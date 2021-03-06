package face

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// AccessoryType enumerates the values for accessory type.
type AccessoryType string

const (
	// Glasses ...
	Glasses AccessoryType = "glasses"
	// HeadWear ...
	HeadWear AccessoryType = "headWear"
	// Mask ...
	Mask AccessoryType = "mask"
)

// PossibleAccessoryTypeValues returns an array of possible values for the AccessoryType const type.
func PossibleAccessoryTypeValues() []AccessoryType {
	return []AccessoryType{Glasses, HeadWear, Mask}
}

// AttributeType enumerates the values for attribute type.
type AttributeType string

const (
	// AttributeTypeAccessories ...
	AttributeTypeAccessories AttributeType = "accessories"
	// AttributeTypeAge ...
	AttributeTypeAge AttributeType = "age"
	// AttributeTypeBlur ...
	AttributeTypeBlur AttributeType = "blur"
	// AttributeTypeEmotion ...
	AttributeTypeEmotion AttributeType = "emotion"
	// AttributeTypeExposure ...
	AttributeTypeExposure AttributeType = "exposure"
	// AttributeTypeFacialHair ...
	AttributeTypeFacialHair AttributeType = "facialHair"
	// AttributeTypeGender ...
	AttributeTypeGender AttributeType = "gender"
	// AttributeTypeGlasses ...
	AttributeTypeGlasses AttributeType = "glasses"
	// AttributeTypeHair ...
	AttributeTypeHair AttributeType = "hair"
	// AttributeTypeHeadPose ...
	AttributeTypeHeadPose AttributeType = "headPose"
	// AttributeTypeMakeup ...
	AttributeTypeMakeup AttributeType = "makeup"
	// AttributeTypeNoise ...
	AttributeTypeNoise AttributeType = "noise"
	// AttributeTypeOcclusion ...
	AttributeTypeOcclusion AttributeType = "occlusion"
	// AttributeTypeSmile ...
	AttributeTypeSmile AttributeType = "smile"
)

// PossibleAttributeTypeValues returns an array of possible values for the AttributeType const type.
func PossibleAttributeTypeValues() []AttributeType {
	return []AttributeType{AttributeTypeAccessories, AttributeTypeAge, AttributeTypeBlur, AttributeTypeEmotion, AttributeTypeExposure, AttributeTypeFacialHair, AttributeTypeGender, AttributeTypeGlasses, AttributeTypeHair, AttributeTypeHeadPose, AttributeTypeMakeup, AttributeTypeNoise, AttributeTypeOcclusion, AttributeTypeSmile}
}

// BlurLevel enumerates the values for blur level.
type BlurLevel string

const (
	// High ...
	High BlurLevel = "High"
	// Low ...
	Low BlurLevel = "Low"
	// Medium ...
	Medium BlurLevel = "Medium"
)

// PossibleBlurLevelValues returns an array of possible values for the BlurLevel const type.
func PossibleBlurLevelValues() []BlurLevel {
	return []BlurLevel{High, Low, Medium}
}

// DetectionModel enumerates the values for detection model.
type DetectionModel string

const (
	// Detection01 ...
	Detection01 DetectionModel = "detection_01"
	// Detection02 ...
	Detection02 DetectionModel = "detection_02"
)

// PossibleDetectionModelValues returns an array of possible values for the DetectionModel const type.
func PossibleDetectionModelValues() []DetectionModel {
	return []DetectionModel{Detection01, Detection02}
}

// ExposureLevel enumerates the values for exposure level.
type ExposureLevel string

const (
	// GoodExposure ...
	GoodExposure ExposureLevel = "GoodExposure"
	// OverExposure ...
	OverExposure ExposureLevel = "OverExposure"
	// UnderExposure ...
	UnderExposure ExposureLevel = "UnderExposure"
)

// PossibleExposureLevelValues returns an array of possible values for the ExposureLevel const type.
func PossibleExposureLevelValues() []ExposureLevel {
	return []ExposureLevel{GoodExposure, OverExposure, UnderExposure}
}

// FindSimilarMatchMode enumerates the values for find similar match mode.
type FindSimilarMatchMode string

const (
	// MatchFace ...
	MatchFace FindSimilarMatchMode = "matchFace"
	// MatchPerson ...
	MatchPerson FindSimilarMatchMode = "matchPerson"
)

// PossibleFindSimilarMatchModeValues returns an array of possible values for the FindSimilarMatchMode const type.
func PossibleFindSimilarMatchModeValues() []FindSimilarMatchMode {
	return []FindSimilarMatchMode{MatchFace, MatchPerson}
}

// Gender enumerates the values for gender.
type Gender string

const (
	// Female ...
	Female Gender = "female"
	// Male ...
	Male Gender = "male"
)

// PossibleGenderValues returns an array of possible values for the Gender const type.
func PossibleGenderValues() []Gender {
	return []Gender{Female, Male}
}

// GlassesType enumerates the values for glasses type.
type GlassesType string

const (
	// NoGlasses ...
	NoGlasses GlassesType = "noGlasses"
	// ReadingGlasses ...
	ReadingGlasses GlassesType = "readingGlasses"
	// Sunglasses ...
	Sunglasses GlassesType = "sunglasses"
	// SwimmingGoggles ...
	SwimmingGoggles GlassesType = "swimmingGoggles"
)

// PossibleGlassesTypeValues returns an array of possible values for the GlassesType const type.
func PossibleGlassesTypeValues() []GlassesType {
	return []GlassesType{NoGlasses, ReadingGlasses, Sunglasses, SwimmingGoggles}
}

// HairColorType enumerates the values for hair color type.
type HairColorType string

const (
	// Black ...
	Black HairColorType = "black"
	// Blond ...
	Blond HairColorType = "blond"
	// Brown ...
	Brown HairColorType = "brown"
	// Gray ...
	Gray HairColorType = "gray"
	// Other ...
	Other HairColorType = "other"
	// Red ...
	Red HairColorType = "red"
	// Unknown ...
	Unknown HairColorType = "unknown"
	// White ...
	White HairColorType = "white"
)

// PossibleHairColorTypeValues returns an array of possible values for the HairColorType const type.
func PossibleHairColorTypeValues() []HairColorType {
	return []HairColorType{Black, Blond, Brown, Gray, Other, Red, Unknown, White}
}

// NoiseLevel enumerates the values for noise level.
type NoiseLevel string

const (
	// NoiseLevelHigh ...
	NoiseLevelHigh NoiseLevel = "High"
	// NoiseLevelLow ...
	NoiseLevelLow NoiseLevel = "Low"
	// NoiseLevelMedium ...
	NoiseLevelMedium NoiseLevel = "Medium"
)

// PossibleNoiseLevelValues returns an array of possible values for the NoiseLevel const type.
func PossibleNoiseLevelValues() []NoiseLevel {
	return []NoiseLevel{NoiseLevelHigh, NoiseLevelLow, NoiseLevelMedium}
}

// OperationStatusType enumerates the values for operation status type.
type OperationStatusType string

const (
	// Failed ...
	Failed OperationStatusType = "failed"
	// Notstarted ...
	Notstarted OperationStatusType = "notstarted"
	// Running ...
	Running OperationStatusType = "running"
	// Succeeded ...
	Succeeded OperationStatusType = "succeeded"
)

// PossibleOperationStatusTypeValues returns an array of possible values for the OperationStatusType const type.
func PossibleOperationStatusTypeValues() []OperationStatusType {
	return []OperationStatusType{Failed, Notstarted, Running, Succeeded}
}

// RecognitionModel enumerates the values for recognition model.
type RecognitionModel string

const (
	// Recognition01 ...
	Recognition01 RecognitionModel = "recognition_01"
	// Recognition02 ...
	Recognition02 RecognitionModel = "recognition_02"
	// Recognition03 ...
	Recognition03 RecognitionModel = "recognition_03"
)

// PossibleRecognitionModelValues returns an array of possible values for the RecognitionModel const type.
func PossibleRecognitionModelValues() []RecognitionModel {
	return []RecognitionModel{Recognition01, Recognition02, Recognition03}
}

// SnapshotApplyMode enumerates the values for snapshot apply mode.
type SnapshotApplyMode string

const (
	// CreateNew ...
	CreateNew SnapshotApplyMode = "CreateNew"
)

// PossibleSnapshotApplyModeValues returns an array of possible values for the SnapshotApplyMode const type.
func PossibleSnapshotApplyModeValues() []SnapshotApplyMode {
	return []SnapshotApplyMode{CreateNew}
}

// SnapshotObjectType enumerates the values for snapshot object type.
type SnapshotObjectType string

const (
	// SnapshotObjectTypeFaceList ...
	SnapshotObjectTypeFaceList SnapshotObjectType = "FaceList"
	// SnapshotObjectTypeLargeFaceList ...
	SnapshotObjectTypeLargeFaceList SnapshotObjectType = "LargeFaceList"
	// SnapshotObjectTypeLargePersonGroup ...
	SnapshotObjectTypeLargePersonGroup SnapshotObjectType = "LargePersonGroup"
	// SnapshotObjectTypePersonGroup ...
	SnapshotObjectTypePersonGroup SnapshotObjectType = "PersonGroup"
)

// PossibleSnapshotObjectTypeValues returns an array of possible values for the SnapshotObjectType const type.
func PossibleSnapshotObjectTypeValues() []SnapshotObjectType {
	return []SnapshotObjectType{SnapshotObjectTypeFaceList, SnapshotObjectTypeLargeFaceList, SnapshotObjectTypeLargePersonGroup, SnapshotObjectTypePersonGroup}
}

// TrainingStatusType enumerates the values for training status type.
type TrainingStatusType string

const (
	// TrainingStatusTypeFailed ...
	TrainingStatusTypeFailed TrainingStatusType = "failed"
	// TrainingStatusTypeNonstarted ...
	TrainingStatusTypeNonstarted TrainingStatusType = "nonstarted"
	// TrainingStatusTypeRunning ...
	TrainingStatusTypeRunning TrainingStatusType = "running"
	// TrainingStatusTypeSucceeded ...
	TrainingStatusTypeSucceeded TrainingStatusType = "succeeded"
)

// PossibleTrainingStatusTypeValues returns an array of possible values for the TrainingStatusType const type.
func PossibleTrainingStatusTypeValues() []TrainingStatusType {
	return []TrainingStatusType{TrainingStatusTypeFailed, TrainingStatusTypeNonstarted, TrainingStatusTypeRunning, TrainingStatusTypeSucceeded}
}
