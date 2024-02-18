/*
Cloud Hypervisor API

Local HTTP based API for managing and inspecting a cloud-hypervisor virtual machine.

API version: 0.3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// VdpaConfig struct for VdpaConfig
type VdpaConfig struct {
	Path       string  `json:"path"`
	NumQueues  int32   `json:"num_queues"`
	Iommu      *bool   `json:"iommu,omitempty"`
	PciSegment *int32  `json:"pci_segment,omitempty"`
	Id         *string `json:"id,omitempty"`
}

// NewVdpaConfig instantiates a new VdpaConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVdpaConfig(path string, numQueues int32) *VdpaConfig {
	this := VdpaConfig{}
	this.Path = path
	this.NumQueues = numQueues
	var iommu bool = false
	this.Iommu = &iommu
	return &this
}

// NewVdpaConfigWithDefaults instantiates a new VdpaConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVdpaConfigWithDefaults() *VdpaConfig {
	this := VdpaConfig{}
	var numQueues int32 = 1
	this.NumQueues = numQueues
	var iommu bool = false
	this.Iommu = &iommu
	return &this
}

// GetPath returns the Path field value
func (o *VdpaConfig) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *VdpaConfig) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *VdpaConfig) SetPath(v string) {
	o.Path = v
}

// GetNumQueues returns the NumQueues field value
func (o *VdpaConfig) GetNumQueues() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NumQueues
}

// GetNumQueuesOk returns a tuple with the NumQueues field value
// and a boolean to check if the value has been set.
func (o *VdpaConfig) GetNumQueuesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NumQueues, true
}

// SetNumQueues sets field value
func (o *VdpaConfig) SetNumQueues(v int32) {
	o.NumQueues = v
}

// GetIommu returns the Iommu field value if set, zero value otherwise.
func (o *VdpaConfig) GetIommu() bool {
	if o == nil || o.Iommu == nil {
		var ret bool
		return ret
	}
	return *o.Iommu
}

// GetIommuOk returns a tuple with the Iommu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VdpaConfig) GetIommuOk() (*bool, bool) {
	if o == nil || o.Iommu == nil {
		return nil, false
	}
	return o.Iommu, true
}

// HasIommu returns a boolean if a field has been set.
func (o *VdpaConfig) HasIommu() bool {
	if o != nil && o.Iommu != nil {
		return true
	}

	return false
}

// SetIommu gets a reference to the given bool and assigns it to the Iommu field.
func (o *VdpaConfig) SetIommu(v bool) {
	o.Iommu = &v
}

// GetPciSegment returns the PciSegment field value if set, zero value otherwise.
func (o *VdpaConfig) GetPciSegment() int32 {
	if o == nil || o.PciSegment == nil {
		var ret int32
		return ret
	}
	return *o.PciSegment
}

// GetPciSegmentOk returns a tuple with the PciSegment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VdpaConfig) GetPciSegmentOk() (*int32, bool) {
	if o == nil || o.PciSegment == nil {
		return nil, false
	}
	return o.PciSegment, true
}

// HasPciSegment returns a boolean if a field has been set.
func (o *VdpaConfig) HasPciSegment() bool {
	if o != nil && o.PciSegment != nil {
		return true
	}

	return false
}

// SetPciSegment gets a reference to the given int32 and assigns it to the PciSegment field.
func (o *VdpaConfig) SetPciSegment(v int32) {
	o.PciSegment = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *VdpaConfig) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VdpaConfig) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *VdpaConfig) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *VdpaConfig) SetId(v string) {
	o.Id = &v
}

func (o VdpaConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["path"] = o.Path
	}
	if true {
		toSerialize["num_queues"] = o.NumQueues
	}
	if o.Iommu != nil {
		toSerialize["iommu"] = o.Iommu
	}
	if o.PciSegment != nil {
		toSerialize["pci_segment"] = o.PciSegment
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableVdpaConfig struct {
	value *VdpaConfig
	isSet bool
}

func (v NullableVdpaConfig) Get() *VdpaConfig {
	return v.value
}

func (v *NullableVdpaConfig) Set(val *VdpaConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableVdpaConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableVdpaConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVdpaConfig(val *VdpaConfig) *NullableVdpaConfig {
	return &NullableVdpaConfig{value: val, isSet: true}
}

func (v NullableVdpaConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVdpaConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
