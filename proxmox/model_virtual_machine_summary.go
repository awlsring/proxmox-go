/*
Proxmox

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2023-01-01
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package proxmox

import (
	"encoding/json"
)

// checks if the VirtualMachineSummary type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VirtualMachineSummary{}

// VirtualMachineSummary struct for VirtualMachineSummary
type VirtualMachineSummary struct {
	// The ID of the virtual machine. Unique across cluster
	Vmid float32 `json:"vmid"`
	Status VirtualMachineStatus `json:"status"`
	// An integer used to represent a boolean. 0 is false, 1 is true.
	Template *float32 `json:"template,omitempty"`
	// Current memory utilization in bytes
	Mem *float32 `json:"mem,omitempty"`
	// Max memory allocated in bytes
	Maxmem *float32 `json:"maxmem,omitempty"`
	// Max root disk size in bytes
	Maxdisk *float32 `json:"maxdisk,omitempty"`
	// The name of the virtual machine
	Name *string `json:"name,omitempty"`
	// The uptime of the virtual machine in seconds
	Uptime *float32 `json:"uptime,omitempty"`
	// Current incoming network traffic in bytes
	Netin *float32 `json:"netin,omitempty"`
	// Current outgoing network traffic in bytes
	Netout *float32 `json:"netout,omitempty"`
	// Current disk read in bytes
	Diskread *float32 `json:"diskread,omitempty"`
	// Current disk write in bytes
	Diskwrite *float32 `json:"diskwrite,omitempty"`
	// The virtual machines cpu utilization in percent
	Cpu *float32 `json:"cpu,omitempty"`
	// Amount of CPU cores allocated to the virtual machine
	Cpus *float32 `json:"cpus,omitempty"`
	// PID of the qemu process on the host node
	Pid *float32 `json:"pid,omitempty"`
	// Tags assigned to the virtual machine. Comma separated list of tags returned as a string.
	Tags *string `json:"tags,omitempty"`
	// The current configuration lock
	Lock *string `json:"lock,omitempty"`
	// Qemu QMP agent status
	Qmpstatus *string `json:"qmpstatus,omitempty"`
}

// NewVirtualMachineSummary instantiates a new VirtualMachineSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualMachineSummary(vmid float32, status VirtualMachineStatus) *VirtualMachineSummary {
	this := VirtualMachineSummary{}
	this.Vmid = vmid
	this.Status = status
	return &this
}

// NewVirtualMachineSummaryWithDefaults instantiates a new VirtualMachineSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualMachineSummaryWithDefaults() *VirtualMachineSummary {
	this := VirtualMachineSummary{}
	return &this
}

// GetVmid returns the Vmid field value
func (o *VirtualMachineSummary) GetVmid() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Vmid
}

// GetVmidOk returns a tuple with the Vmid field value
// and a boolean to check if the value has been set.
func (o *VirtualMachineSummary) GetVmidOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Vmid, true
}

// SetVmid sets field value
func (o *VirtualMachineSummary) SetVmid(v float32) {
	o.Vmid = v
}

// GetStatus returns the Status field value
func (o *VirtualMachineSummary) GetStatus() VirtualMachineStatus {
	if o == nil {
		var ret VirtualMachineStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *VirtualMachineSummary) GetStatusOk() (*VirtualMachineStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *VirtualMachineSummary) SetStatus(v VirtualMachineStatus) {
	o.Status = v
}

// GetTemplate returns the Template field value if set, zero value otherwise.
func (o *VirtualMachineSummary) GetTemplate() float32 {
	if o == nil || IsNil(o.Template) {
		var ret float32
		return ret
	}
	return *o.Template
}

// GetTemplateOk returns a tuple with the Template field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualMachineSummary) GetTemplateOk() (*float32, bool) {
	if o == nil || IsNil(o.Template) {
		return nil, false
	}
	return o.Template, true
}

// HasTemplate returns a boolean if a field has been set.
func (o *VirtualMachineSummary) HasTemplate() bool {
	if o != nil && !IsNil(o.Template) {
		return true
	}

	return false
}

// SetTemplate gets a reference to the given float32 and assigns it to the Template field.
func (o *VirtualMachineSummary) SetTemplate(v float32) {
	o.Template = &v
}

// GetMem returns the Mem field value if set, zero value otherwise.
func (o *VirtualMachineSummary) GetMem() float32 {
	if o == nil || IsNil(o.Mem) {
		var ret float32
		return ret
	}
	return *o.Mem
}

// GetMemOk returns a tuple with the Mem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualMachineSummary) GetMemOk() (*float32, bool) {
	if o == nil || IsNil(o.Mem) {
		return nil, false
	}
	return o.Mem, true
}

// HasMem returns a boolean if a field has been set.
func (o *VirtualMachineSummary) HasMem() bool {
	if o != nil && !IsNil(o.Mem) {
		return true
	}

	return false
}

// SetMem gets a reference to the given float32 and assigns it to the Mem field.
func (o *VirtualMachineSummary) SetMem(v float32) {
	o.Mem = &v
}

// GetMaxmem returns the Maxmem field value if set, zero value otherwise.
func (o *VirtualMachineSummary) GetMaxmem() float32 {
	if o == nil || IsNil(o.Maxmem) {
		var ret float32
		return ret
	}
	return *o.Maxmem
}

// GetMaxmemOk returns a tuple with the Maxmem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualMachineSummary) GetMaxmemOk() (*float32, bool) {
	if o == nil || IsNil(o.Maxmem) {
		return nil, false
	}
	return o.Maxmem, true
}

// HasMaxmem returns a boolean if a field has been set.
func (o *VirtualMachineSummary) HasMaxmem() bool {
	if o != nil && !IsNil(o.Maxmem) {
		return true
	}

	return false
}

// SetMaxmem gets a reference to the given float32 and assigns it to the Maxmem field.
func (o *VirtualMachineSummary) SetMaxmem(v float32) {
	o.Maxmem = &v
}

// GetMaxdisk returns the Maxdisk field value if set, zero value otherwise.
func (o *VirtualMachineSummary) GetMaxdisk() float32 {
	if o == nil || IsNil(o.Maxdisk) {
		var ret float32
		return ret
	}
	return *o.Maxdisk
}

// GetMaxdiskOk returns a tuple with the Maxdisk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualMachineSummary) GetMaxdiskOk() (*float32, bool) {
	if o == nil || IsNil(o.Maxdisk) {
		return nil, false
	}
	return o.Maxdisk, true
}

// HasMaxdisk returns a boolean if a field has been set.
func (o *VirtualMachineSummary) HasMaxdisk() bool {
	if o != nil && !IsNil(o.Maxdisk) {
		return true
	}

	return false
}

// SetMaxdisk gets a reference to the given float32 and assigns it to the Maxdisk field.
func (o *VirtualMachineSummary) SetMaxdisk(v float32) {
	o.Maxdisk = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *VirtualMachineSummary) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualMachineSummary) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *VirtualMachineSummary) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *VirtualMachineSummary) SetName(v string) {
	o.Name = &v
}

// GetUptime returns the Uptime field value if set, zero value otherwise.
func (o *VirtualMachineSummary) GetUptime() float32 {
	if o == nil || IsNil(o.Uptime) {
		var ret float32
		return ret
	}
	return *o.Uptime
}

// GetUptimeOk returns a tuple with the Uptime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualMachineSummary) GetUptimeOk() (*float32, bool) {
	if o == nil || IsNil(o.Uptime) {
		return nil, false
	}
	return o.Uptime, true
}

// HasUptime returns a boolean if a field has been set.
func (o *VirtualMachineSummary) HasUptime() bool {
	if o != nil && !IsNil(o.Uptime) {
		return true
	}

	return false
}

// SetUptime gets a reference to the given float32 and assigns it to the Uptime field.
func (o *VirtualMachineSummary) SetUptime(v float32) {
	o.Uptime = &v
}

// GetNetin returns the Netin field value if set, zero value otherwise.
func (o *VirtualMachineSummary) GetNetin() float32 {
	if o == nil || IsNil(o.Netin) {
		var ret float32
		return ret
	}
	return *o.Netin
}

// GetNetinOk returns a tuple with the Netin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualMachineSummary) GetNetinOk() (*float32, bool) {
	if o == nil || IsNil(o.Netin) {
		return nil, false
	}
	return o.Netin, true
}

// HasNetin returns a boolean if a field has been set.
func (o *VirtualMachineSummary) HasNetin() bool {
	if o != nil && !IsNil(o.Netin) {
		return true
	}

	return false
}

// SetNetin gets a reference to the given float32 and assigns it to the Netin field.
func (o *VirtualMachineSummary) SetNetin(v float32) {
	o.Netin = &v
}

// GetNetout returns the Netout field value if set, zero value otherwise.
func (o *VirtualMachineSummary) GetNetout() float32 {
	if o == nil || IsNil(o.Netout) {
		var ret float32
		return ret
	}
	return *o.Netout
}

// GetNetoutOk returns a tuple with the Netout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualMachineSummary) GetNetoutOk() (*float32, bool) {
	if o == nil || IsNil(o.Netout) {
		return nil, false
	}
	return o.Netout, true
}

// HasNetout returns a boolean if a field has been set.
func (o *VirtualMachineSummary) HasNetout() bool {
	if o != nil && !IsNil(o.Netout) {
		return true
	}

	return false
}

// SetNetout gets a reference to the given float32 and assigns it to the Netout field.
func (o *VirtualMachineSummary) SetNetout(v float32) {
	o.Netout = &v
}

// GetDiskread returns the Diskread field value if set, zero value otherwise.
func (o *VirtualMachineSummary) GetDiskread() float32 {
	if o == nil || IsNil(o.Diskread) {
		var ret float32
		return ret
	}
	return *o.Diskread
}

// GetDiskreadOk returns a tuple with the Diskread field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualMachineSummary) GetDiskreadOk() (*float32, bool) {
	if o == nil || IsNil(o.Diskread) {
		return nil, false
	}
	return o.Diskread, true
}

// HasDiskread returns a boolean if a field has been set.
func (o *VirtualMachineSummary) HasDiskread() bool {
	if o != nil && !IsNil(o.Diskread) {
		return true
	}

	return false
}

// SetDiskread gets a reference to the given float32 and assigns it to the Diskread field.
func (o *VirtualMachineSummary) SetDiskread(v float32) {
	o.Diskread = &v
}

// GetDiskwrite returns the Diskwrite field value if set, zero value otherwise.
func (o *VirtualMachineSummary) GetDiskwrite() float32 {
	if o == nil || IsNil(o.Diskwrite) {
		var ret float32
		return ret
	}
	return *o.Diskwrite
}

// GetDiskwriteOk returns a tuple with the Diskwrite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualMachineSummary) GetDiskwriteOk() (*float32, bool) {
	if o == nil || IsNil(o.Diskwrite) {
		return nil, false
	}
	return o.Diskwrite, true
}

// HasDiskwrite returns a boolean if a field has been set.
func (o *VirtualMachineSummary) HasDiskwrite() bool {
	if o != nil && !IsNil(o.Diskwrite) {
		return true
	}

	return false
}

// SetDiskwrite gets a reference to the given float32 and assigns it to the Diskwrite field.
func (o *VirtualMachineSummary) SetDiskwrite(v float32) {
	o.Diskwrite = &v
}

// GetCpu returns the Cpu field value if set, zero value otherwise.
func (o *VirtualMachineSummary) GetCpu() float32 {
	if o == nil || IsNil(o.Cpu) {
		var ret float32
		return ret
	}
	return *o.Cpu
}

// GetCpuOk returns a tuple with the Cpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualMachineSummary) GetCpuOk() (*float32, bool) {
	if o == nil || IsNil(o.Cpu) {
		return nil, false
	}
	return o.Cpu, true
}

// HasCpu returns a boolean if a field has been set.
func (o *VirtualMachineSummary) HasCpu() bool {
	if o != nil && !IsNil(o.Cpu) {
		return true
	}

	return false
}

// SetCpu gets a reference to the given float32 and assigns it to the Cpu field.
func (o *VirtualMachineSummary) SetCpu(v float32) {
	o.Cpu = &v
}

// GetCpus returns the Cpus field value if set, zero value otherwise.
func (o *VirtualMachineSummary) GetCpus() float32 {
	if o == nil || IsNil(o.Cpus) {
		var ret float32
		return ret
	}
	return *o.Cpus
}

// GetCpusOk returns a tuple with the Cpus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualMachineSummary) GetCpusOk() (*float32, bool) {
	if o == nil || IsNil(o.Cpus) {
		return nil, false
	}
	return o.Cpus, true
}

// HasCpus returns a boolean if a field has been set.
func (o *VirtualMachineSummary) HasCpus() bool {
	if o != nil && !IsNil(o.Cpus) {
		return true
	}

	return false
}

// SetCpus gets a reference to the given float32 and assigns it to the Cpus field.
func (o *VirtualMachineSummary) SetCpus(v float32) {
	o.Cpus = &v
}

// GetPid returns the Pid field value if set, zero value otherwise.
func (o *VirtualMachineSummary) GetPid() float32 {
	if o == nil || IsNil(o.Pid) {
		var ret float32
		return ret
	}
	return *o.Pid
}

// GetPidOk returns a tuple with the Pid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualMachineSummary) GetPidOk() (*float32, bool) {
	if o == nil || IsNil(o.Pid) {
		return nil, false
	}
	return o.Pid, true
}

// HasPid returns a boolean if a field has been set.
func (o *VirtualMachineSummary) HasPid() bool {
	if o != nil && !IsNil(o.Pid) {
		return true
	}

	return false
}

// SetPid gets a reference to the given float32 and assigns it to the Pid field.
func (o *VirtualMachineSummary) SetPid(v float32) {
	o.Pid = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *VirtualMachineSummary) GetTags() string {
	if o == nil || IsNil(o.Tags) {
		var ret string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualMachineSummary) GetTagsOk() (*string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *VirtualMachineSummary) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given string and assigns it to the Tags field.
func (o *VirtualMachineSummary) SetTags(v string) {
	o.Tags = &v
}

// GetLock returns the Lock field value if set, zero value otherwise.
func (o *VirtualMachineSummary) GetLock() string {
	if o == nil || IsNil(o.Lock) {
		var ret string
		return ret
	}
	return *o.Lock
}

// GetLockOk returns a tuple with the Lock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualMachineSummary) GetLockOk() (*string, bool) {
	if o == nil || IsNil(o.Lock) {
		return nil, false
	}
	return o.Lock, true
}

// HasLock returns a boolean if a field has been set.
func (o *VirtualMachineSummary) HasLock() bool {
	if o != nil && !IsNil(o.Lock) {
		return true
	}

	return false
}

// SetLock gets a reference to the given string and assigns it to the Lock field.
func (o *VirtualMachineSummary) SetLock(v string) {
	o.Lock = &v
}

// GetQmpstatus returns the Qmpstatus field value if set, zero value otherwise.
func (o *VirtualMachineSummary) GetQmpstatus() string {
	if o == nil || IsNil(o.Qmpstatus) {
		var ret string
		return ret
	}
	return *o.Qmpstatus
}

// GetQmpstatusOk returns a tuple with the Qmpstatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualMachineSummary) GetQmpstatusOk() (*string, bool) {
	if o == nil || IsNil(o.Qmpstatus) {
		return nil, false
	}
	return o.Qmpstatus, true
}

// HasQmpstatus returns a boolean if a field has been set.
func (o *VirtualMachineSummary) HasQmpstatus() bool {
	if o != nil && !IsNil(o.Qmpstatus) {
		return true
	}

	return false
}

// SetQmpstatus gets a reference to the given string and assigns it to the Qmpstatus field.
func (o *VirtualMachineSummary) SetQmpstatus(v string) {
	o.Qmpstatus = &v
}

func (o VirtualMachineSummary) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VirtualMachineSummary) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["vmid"] = o.Vmid
	toSerialize["status"] = o.Status
	if !IsNil(o.Template) {
		toSerialize["template"] = o.Template
	}
	if !IsNil(o.Mem) {
		toSerialize["mem"] = o.Mem
	}
	if !IsNil(o.Maxmem) {
		toSerialize["maxmem"] = o.Maxmem
	}
	if !IsNil(o.Maxdisk) {
		toSerialize["maxdisk"] = o.Maxdisk
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Uptime) {
		toSerialize["uptime"] = o.Uptime
	}
	if !IsNil(o.Netin) {
		toSerialize["netin"] = o.Netin
	}
	if !IsNil(o.Netout) {
		toSerialize["netout"] = o.Netout
	}
	if !IsNil(o.Diskread) {
		toSerialize["diskread"] = o.Diskread
	}
	if !IsNil(o.Diskwrite) {
		toSerialize["diskwrite"] = o.Diskwrite
	}
	if !IsNil(o.Cpu) {
		toSerialize["cpu"] = o.Cpu
	}
	if !IsNil(o.Cpus) {
		toSerialize["cpus"] = o.Cpus
	}
	if !IsNil(o.Pid) {
		toSerialize["pid"] = o.Pid
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.Lock) {
		toSerialize["lock"] = o.Lock
	}
	if !IsNil(o.Qmpstatus) {
		toSerialize["qmpstatus"] = o.Qmpstatus
	}
	return toSerialize, nil
}

type NullableVirtualMachineSummary struct {
	value *VirtualMachineSummary
	isSet bool
}

func (v NullableVirtualMachineSummary) Get() *VirtualMachineSummary {
	return v.value
}

func (v *NullableVirtualMachineSummary) Set(val *VirtualMachineSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualMachineSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualMachineSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualMachineSummary(val *VirtualMachineSummary) *NullableVirtualMachineSummary {
	return &NullableVirtualMachineSummary{value: val, isSet: true}
}

func (v NullableVirtualMachineSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualMachineSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


