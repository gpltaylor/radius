//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package radclient

import "encoding/json"

func unmarshalComponentTraitClassification(rawMsg json.RawMessage) (ComponentTraitClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b ComponentTraitClassification
	switch m["kind"] {
	case "dapr.io/Sidecar@v1alpha1":
		b = &DaprSidecarTrait{}
	case "radius.dev/ManualScaling@v1alpha1":
		b = &ManualScalingTrait{}
	default:
		b = &ComponentTrait{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalComponentTraitClassificationArray(rawMsg json.RawMessage) ([]ComponentTraitClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]ComponentTraitClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalComponentTraitClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalComponentTraitClassificationMap(rawMsg json.RawMessage) (map[string]ComponentTraitClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages map[string]json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fMap := make(map[string]ComponentTraitClassification, len(rawMessages))
	for key, rawMessage := range rawMessages {
		f, err := unmarshalComponentTraitClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fMap[key] = f
	}
	return fMap, nil
}

func unmarshalHealthProbePropertiesClassification(rawMsg json.RawMessage) (HealthProbePropertiesClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b HealthProbePropertiesClassification
	switch m["kind"] {
	case "exec":
		b = &ExecHealthProbeProperties{}
	case "httpGet":
		b = &HTTPGetHealthProbeProperties{}
	case "tcp":
		b = &TCPHealthProbeProperties{}
	default:
		b = &HealthProbeProperties{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalHealthProbePropertiesClassificationArray(rawMsg json.RawMessage) ([]HealthProbePropertiesClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]HealthProbePropertiesClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalHealthProbePropertiesClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalHealthProbePropertiesClassificationMap(rawMsg json.RawMessage) (map[string]HealthProbePropertiesClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages map[string]json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fMap := make(map[string]HealthProbePropertiesClassification, len(rawMessages))
	for key, rawMessage := range rawMessages {
		f, err := unmarshalHealthProbePropertiesClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fMap[key] = f
	}
	return fMap, nil
}

func unmarshalVolumeClassification(rawMsg json.RawMessage) (VolumeClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b VolumeClassification
	switch m["kind"] {
	case "ephemeral":
		b = &EphemeralVolume{}
	case "persistent":
		b = &PersistentVolume{}
	default:
		b = &Volume{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalVolumeClassificationArray(rawMsg json.RawMessage) ([]VolumeClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]VolumeClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalVolumeClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalVolumeClassificationMap(rawMsg json.RawMessage) (map[string]VolumeClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages map[string]json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fMap := make(map[string]VolumeClassification, len(rawMessages))
	for key, rawMessage := range rawMessages {
		f, err := unmarshalVolumeClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fMap[key] = f
	}
	return fMap, nil
}

func unmarshalVolumePropertiesClassification(rawMsg json.RawMessage) (VolumePropertiesClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b VolumePropertiesClassification
	switch m["kind"] {
	case "azure.com.fileshare":
		b = &AzureFileShareVolumeProperties{}
	case "azure.com.keyvault":
		b = &AzureKeyVaultVolumeProperties{}
	default:
		b = &VolumeProperties{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalVolumePropertiesClassificationArray(rawMsg json.RawMessage) ([]VolumePropertiesClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]VolumePropertiesClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalVolumePropertiesClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalVolumePropertiesClassificationMap(rawMsg json.RawMessage) (map[string]VolumePropertiesClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages map[string]json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fMap := make(map[string]VolumePropertiesClassification, len(rawMessages))
	for key, rawMessage := range rawMessages {
		f, err := unmarshalVolumePropertiesClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fMap[key] = f
	}
	return fMap, nil
}

