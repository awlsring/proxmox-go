# ApplyVirtualMachineConfigurationSyncRequestContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Acpi** | Pointer to **float32** | Enable ACPI support. Default to enabled. | [optional] 
**Affinity** | Pointer to **string** | List of cores to execute processes. Example value: 1,5,8-11. | [optional] 
**Agent** | Pointer to **string** | The QEMU agent and its configuration. | [optional] 
**Arch** | Pointer to [**VirtualMachineArchitecture**](VirtualMachineArchitecture.md) |  | [optional] 
**Args** | Pointer to **string** | Additional command line arguments passed to the kvm. | [optional] 
**Audio0** | Pointer to **string** | The audio device and its configuration. | [optional] 
**Autostart** | Pointer to **float32** | Start the virtual machine on crash. | [optional] 
**Ballon** | Pointer to **float32** | Amount of RAM for the VM in MB. | [optional] 
**Boot** | Pointer to **string** | The boot order of the virtual machine. | [optional] 
**Bootdisk** | Pointer to **string** | The boot disk of the virtual machine. | [optional] 
**Cdrom** | Pointer to **string** | The CD-ROM device and its configuration. An alias for option ide2 | [optional] 
**Cicustom** | Pointer to **string** | Specify custom cloud-init files to be used at start. | [optional] 
**Cipassword** | Pointer to **string** | The password for the cloud-init user. | [optional] 
**Citype** | Pointer to [**VirtualMachineCloudInitType**](VirtualMachineCloudInitType.md) |  | [optional] 
**Ciuser** | Pointer to **string** | The cloud-init user. | [optional] 
**Cores** | Pointer to **float32** | Number of cores per socket. | [optional] 
**Bios** | Pointer to [**VirtualMachineBios**](VirtualMachineBios.md) |  | [optional] 
**Cpu** | Pointer to **string** | The CPU type. | [optional] 
**Cpulimit** | Pointer to **float32** | CPU usage limit. | [optional] 
**Cpuunits** | Pointer to **float32** | CPU weight for a VM. The higher the value, the more CPU time the VM gets. | [optional] 
**Description** | Pointer to **string** | The description of the virtual machine. | [optional] 
**Digest** | Pointer to **string** | The SHA1 digest of the virtual machine configuration. This can prevent concurrent modifications of the virtual machine configuration. | [optional] 
**Delete** | Pointer to **string** | A list of settings to delete from the configuration. | [optional] 
**Efidisk0** | Pointer to **string** | The EFI disk device and its configuration. | [optional] 
**Freeze** | Pointer to **bool** | Freeze the CPU at virtual machine start. | [optional] 
**Hookscript** | Pointer to **string** | The hook script that is used at various point in the virtual machines lifecycle. | [optional] 
**Hostpci0** | Pointer to **string** | The host PCI device and its configuration. | [optional] 
**Hostpci1** | Pointer to **string** | The host PCI device and its configuration. | [optional] 
**Hostpci2** | Pointer to **string** | The host PCI device and its configuration. | [optional] 
**Hostpci3** | Pointer to **string** | The host PCI device and its configuration. | [optional] 
**Hostpci4** | Pointer to **string** | The host PCI device and its configuration. | [optional] 
**Hostpci5** | Pointer to **string** | The host PCI device and its configuration. | [optional] 
**Hostpci6** | Pointer to **string** | The host PCI device and its configuration. | [optional] 
**Hostpci7** | Pointer to **string** | The host PCI device and its configuration. | [optional] 
**Hostpci8** | Pointer to **string** | The host PCI device and its configuration. | [optional] 
**Hostpci9** | Pointer to **string** | The host PCI device and its configuration. | [optional] 
**Hugepages** | Pointer to [**VirtualMachineHugePages**](VirtualMachineHugePages.md) |  | [optional] 
**Ide0** | Pointer to **string** | Volume used as IDE or harddisk. | [optional] 
**Ide1** | Pointer to **string** | Volume used as IDE or harddisk. | [optional] 
**Ide2** | Pointer to **string** | Volume used as IDE or harddisk. | [optional] 
**Ide3** | Pointer to **string** | Volume used as IDE or harddisk. | [optional] 
**Ipconfig0** | Pointer to **string** | The ip address and gateways for the interface. Ip addresses are in CIDR format. Example value: ip&#x3D;10.0.100.101/24,gw&#x3D;10.0.100.1 | [optional] 
**Ipconfig1** | Pointer to **string** | The ip address and gateways for the interface. Ip addresses are in CIDR format. Example value: ip&#x3D;10.0.100.101/24,gw&#x3D;10.0.100.1 | [optional] 
**Ipconfig2** | Pointer to **string** | The ip address and gateways for the interface. Ip addresses are in CIDR format. Example value: ip&#x3D;10.0.100.101/24,gw&#x3D;10.0.100.1 | [optional] 
**Ipconfig3** | Pointer to **string** | The ip address and gateways for the interface. Ip addresses are in CIDR format. Example value: ip&#x3D;10.0.100.101/24,gw&#x3D;10.0.100.1 | [optional] 
**Ipconfig4** | Pointer to **string** | The ip address and gateways for the interface. Ip addresses are in CIDR format. Example value: ip&#x3D;10.0.100.101/24,gw&#x3D;10.0.100.1 | [optional] 
**Ipconfig5** | Pointer to **string** | The ip address and gateways for the interface. Ip addresses are in CIDR format. Example value: ip&#x3D;10.0.100.101/24,gw&#x3D;10.0.100.1 | [optional] 
**Ipconfig6** | Pointer to **string** | The ip address and gateways for the interface. Ip addresses are in CIDR format. Example value: ip&#x3D;10.0.100.101/24,gw&#x3D;10.0.100.1 | [optional] 
**Ipconfig7** | Pointer to **string** | The ip address and gateways for the interface. Ip addresses are in CIDR format. Example value: ip&#x3D;10.0.100.101/24,gw&#x3D;10.0.100.1 | [optional] 
**Ivshmem** | Pointer to **string** | Inter VM-shared memory. | [optional] 
**Keephugepages** | Pointer to **float32** | Keep hugepages after shutdown. | [optional] 
**Keyboard** | Pointer to [**VirtualMachineKeyboard**](VirtualMachineKeyboard.md) |  | [optional] 
**Kvm** | Pointer to **float32** | Enable KVM support. Default to enabled. | [optional] 
**Localtime** | Pointer to **float32** | Set the real time clock to local time. | [optional] 
**Lock** | Pointer to [**VirtualMachineConfigLock**](VirtualMachineConfigLock.md) |  | [optional] 
**Machine** | Pointer to **string** | The machine type. | [optional] 
**Memory** | Pointer to **float32** | The amount of memory in MB. | [optional] 
**MigrateDowntime** | Pointer to **float32** | The maximum tolerated downtime in seconds during migration. | [optional] 
**MigrateSpeed** | Pointer to **float32** | The maximum speed in MB/s during migration. 0 is no limit | [optional] 
**Name** | Pointer to **string** | The name of the virtual machine. | [optional] 
**Nameserver** | Pointer to **string** | The nameserver for the virtual machine. | [optional] 
**Net0** | Pointer to **string** | The network interface and its configuration. | [optional] 
**Net1** | Pointer to **string** | The network interface and its configuration. | [optional] 
**Net2** | Pointer to **string** | The network interface and its configuration. | [optional] 
**Net3** | Pointer to **string** | The network interface and its configuration. | [optional] 
**Net4** | Pointer to **string** | The network interface and its configuration. | [optional] 
**Net5** | Pointer to **string** | The network interface and its configuration. | [optional] 
**Net6** | Pointer to **string** | The network interface and its configuration. | [optional] 
**Net7** | Pointer to **string** | The network interface and its configuration. | [optional] 
**Numa** | Pointer to **float32** | Enable NUMA support. Default to disabled. | [optional] 
**Numa0** | Pointer to **string** | NUMA topology. | [optional] 
**Numa1** | Pointer to **string** | NUMA topology. | [optional] 
**Numa2** | Pointer to **string** | NUMA topology. | [optional] 
**Numa3** | Pointer to **string** | NUMA topology. | [optional] 
**Numa4** | Pointer to **string** | NUMA topology. | [optional] 
**Numa5** | Pointer to **string** | NUMA topology. | [optional] 
**Numa6** | Pointer to **string** | NUMA topology. | [optional] 
**Numa7** | Pointer to **string** | NUMA topology. | [optional] 
**Onboot** | Pointer to **float32** | Start the virtual machine on boot. | [optional] 
**Ostype** | Pointer to [**VirtualMachineOperatingSystem**](VirtualMachineOperatingSystem.md) |  | [optional] 
**Parallel0** | Pointer to **string** | Host parallel device. | [optional] 
**Parallel1** | Pointer to **string** | Host parallel device. | [optional] 
**Parallel2** | Pointer to **string** | Host parallel device. | [optional] 
**Protection** | Pointer to **float32** | The protection flag on the virtual machine. Disables remove VM and disk operations. | [optional] 
**Reboot** | Pointer to **float32** | Allows reboot. False will have the virtual machine exit on reboot. | [optional] 
**Rng0** | Pointer to **string** | Virtio based random number generator. | [optional] 
**Sata0** | Pointer to **string** | Uses the volume as a sata disk. | [optional] 
**Sata1** | Pointer to **string** | Uses the volume as a sata disk. | [optional] 
**Sata2** | Pointer to **string** | Uses the volume as a sata disk. | [optional] 
**Sata3** | Pointer to **string** | Uses the volume as a sata disk. | [optional] 
**Sata4** | Pointer to **string** | Uses the volume as a sata disk. | [optional] 
**Sata5** | Pointer to **string** | Uses the volume as a sata disk. | [optional] 
**Scsi0** | Pointer to **string** | Uses the volume as a scsi disk. | [optional] 
**Scsi1** | Pointer to **string** | Uses the volume as a scsi disk. | [optional] 
**Scsi2** | Pointer to **string** | Uses the volume as a scsi disk. | [optional] 
**Scsi3** | Pointer to **string** | Uses the volume as a scsi disk. | [optional] 
**Scsi4** | Pointer to **string** | Uses the volume as a scsi disk. | [optional] 
**Scsi5** | Pointer to **string** | Uses the volume as a scsi disk. | [optional] 
**Scsi6** | Pointer to **string** | Uses the volume as a scsi disk. | [optional] 
**Scsi7** | Pointer to **string** | Uses the volume as a scsi disk. | [optional] 
**Scsi8** | Pointer to **string** | Uses the volume as a scsi disk. | [optional] 
**Scsi9** | Pointer to **string** | Uses the volume as a scsi disk. | [optional] 
**Scsi10** | Pointer to **string** | Uses the volume as a scsi disk. | [optional] 
**Scsi11** | Pointer to **string** | Uses the volume as a scsi disk. | [optional] 
**Scsi12** | Pointer to **string** | Uses the volume as a scsi disk. | [optional] 
**Scsi13** | Pointer to **string** | Uses the volume as a scsi disk. | [optional] 
**Scsi14** | Pointer to **string** | Uses the volume as a scsi disk. | [optional] 
**Scsi15** | Pointer to **string** | Uses the volume as a scsi disk. | [optional] 
**Scsi16** | Pointer to **string** | Uses the volume as a scsi disk. | [optional] 
**Scsi17** | Pointer to **string** | Uses the volume as a scsi disk. | [optional] 
**Scsi18** | Pointer to **string** | Uses the volume as a scsi disk. | [optional] 
**Scsi19** | Pointer to **string** | Uses the volume as a scsi disk. | [optional] 
**Scsi20** | Pointer to **string** | Uses the volume as a scsi disk. | [optional] 
**Scsi21** | Pointer to **string** | Uses the volume as a scsi disk. | [optional] 
**Scsi22** | Pointer to **string** | Uses the volume as a scsi disk. | [optional] 
**Scsi23** | Pointer to **string** | Uses the volume as a scsi disk. | [optional] 
**Scsi24** | Pointer to **string** | Uses the volume as a scsi disk. | [optional] 
**Scsi25** | Pointer to **string** | Uses the volume as a scsi disk. | [optional] 
**Scsi26** | Pointer to **string** | Uses the volume as a scsi disk. | [optional] 
**Scsi27** | Pointer to **string** | Uses the volume as a scsi disk. | [optional] 
**Scsi28** | Pointer to **string** | Uses the volume as a scsi disk. | [optional] 
**Scsi29** | Pointer to **string** | Uses the volume as a scsi disk. | [optional] 
**Scsi30** | Pointer to **string** | Uses the volume as a scsi disk. | [optional] 
**Scsihw** | Pointer to [**VirtualMachineScsiControllerType**](VirtualMachineScsiControllerType.md) |  | [optional] 
**Searchdomain** | Pointer to **string** | Cloudinit search domain. | [optional] 
**Serial0** | Pointer to **string** | A serial device on the virtual machine. | [optional] 
**Serial1** | Pointer to **string** | A serial device on the virtual machine. | [optional] 
**Serial2** | Pointer to **string** | A serial device on the virtual machine. | [optional] 
**Serial3** | Pointer to **string** | A serial device on the virtual machine. | [optional] 
**Shares** | Pointer to **float32** | The amount of memory shares for autoballooning. | [optional] 
**Smbios1** | Pointer to **string** | SMBIOS type 1 field. | [optional] 
**Sockets** | Pointer to **float32** | The number of cpu sockets. | [optional] 
**SpiceEnhancements** | Pointer to **string** | Enable spice enhancements. | [optional] 
**Sshkeys** | Pointer to **string** | Cloud init SSH public keys. One per line. | [optional] 
**Startdate** | Pointer to **string** | The start date for the virtual machine real time clock. | [optional] 
**Startup** | Pointer to **string** | The startup policy for the virtual machine. | [optional] 
**Tablet** | Pointer to **float32** | Enable tablet device. | [optional] 
**Tags** | Pointer to **string** | The tags for the virtual machine. | [optional] 
**Template** | Pointer to **float32** | If the virtual machine is a template or not. | [optional] 
**Tpmstate0** | Pointer to **string** | Configure a disk for storing TPM state. | [optional] 
**Usb0** | Pointer to **string** | A usb device on the virtual machine. | [optional] 
**Usb1** | Pointer to **string** | A usb device on the virtual machine. | [optional] 
**Usb2** | Pointer to **string** | A usb device on the virtual machine. | [optional] 
**Usb3** | Pointer to **string** | A usb device on the virtual machine. | [optional] 
**Usb4** | Pointer to **string** | A usb device on the virtual machine. | [optional] 
**Usb5** | Pointer to **string** | A usb device on the virtual machine. | [optional] 
**Usb6** | Pointer to **string** | A usb device on the virtual machine. | [optional] 
**Usb7** | Pointer to **string** | A usb device on the virtual machine. | [optional] 
**Usb8** | Pointer to **string** | A usb device on the virtual machine. | [optional] 
**Usb9** | Pointer to **string** | A usb device on the virtual machine. | [optional] 
**Usb10** | Pointer to **string** | A usb device on the virtual machine. | [optional] 
**Usb11** | Pointer to **string** | A usb device on the virtual machine. | [optional] 
**Usb12** | Pointer to **string** | A usb device on the virtual machine. | [optional] 
**Usb13** | Pointer to **string** | A usb device on the virtual machine. | [optional] 
**Usb14** | Pointer to **string** | A usb device on the virtual machine. | [optional] 
**Hotplug** | Pointer to **string** | Hotplug devices. | [optional] 
**Vcpus** | Pointer to **float32** | The number of hotplugged virtual cpus. | [optional] 
**Vga** | Pointer to **string** | Configuration for the VGA hardware. | [optional] 
**Virtio0** | Pointer to **string** | Uses the volume as a virtio disk. | [optional] 
**Virtio1** | Pointer to **string** | Uses the volume as a virtio disk. | [optional] 
**Virtio2** | Pointer to **string** | Uses the volume as a virtio disk. | [optional] 
**Virtio3** | Pointer to **string** | Uses the volume as a virtio disk. | [optional] 
**Virtio4** | Pointer to **string** | Uses the volume as a virtio disk. | [optional] 
**Virtio5** | Pointer to **string** | Uses the volume as a virtio disk. | [optional] 
**Virtio6** | Pointer to **string** | Uses the volume as a virtio disk. | [optional] 
**Virtio7** | Pointer to **string** | Uses the volume as a virtio disk. | [optional] 
**Virtio8** | Pointer to **string** | Uses the volume as a virtio disk. | [optional] 
**Virtio9** | Pointer to **string** | Uses the volume as a virtio disk. | [optional] 
**Virtio10** | Pointer to **string** | Uses the volume as a virtio disk. | [optional] 
**Virtio11** | Pointer to **string** | Uses the volume as a virtio disk. | [optional] 
**Virtio12** | Pointer to **string** | Uses the volume as a virtio disk. | [optional] 
**Virtio13** | Pointer to **string** | Uses the volume as a virtio disk. | [optional] 
**Virtio14** | Pointer to **string** | Uses the volume as a virtio disk. | [optional] 
**Virtio15** | Pointer to **string** | Uses the volume as a virtio disk. | [optional] 
**Vmgenid** | Pointer to **string** | Enable VM generation id seed. | [optional] 
**Vmstatestorage** | Pointer to **string** | Default location for storing VM state. | [optional] 
**Watchdog** | Pointer to **string** | The watchdog device for the virtual machine. | [optional] 

## Methods

### NewApplyVirtualMachineConfigurationSyncRequestContent

`func NewApplyVirtualMachineConfigurationSyncRequestContent() *ApplyVirtualMachineConfigurationSyncRequestContent`

NewApplyVirtualMachineConfigurationSyncRequestContent instantiates a new ApplyVirtualMachineConfigurationSyncRequestContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplyVirtualMachineConfigurationSyncRequestContentWithDefaults

`func NewApplyVirtualMachineConfigurationSyncRequestContentWithDefaults() *ApplyVirtualMachineConfigurationSyncRequestContent`

NewApplyVirtualMachineConfigurationSyncRequestContentWithDefaults instantiates a new ApplyVirtualMachineConfigurationSyncRequestContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcpi

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetAcpi() float32`

GetAcpi returns the Acpi field if non-nil, zero value otherwise.

### GetAcpiOk

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetAcpiOk() (*float32, bool)`

GetAcpiOk returns a tuple with the Acpi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcpi

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetAcpi(v float32)`

SetAcpi sets Acpi field to given value.

### HasAcpi

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasAcpi() bool`

HasAcpi returns a boolean if a field has been set.

### GetAffinity

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetAffinity() string`

GetAffinity returns the Affinity field if non-nil, zero value otherwise.

### GetAffinityOk

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetAffinityOk() (*string, bool)`

GetAffinityOk returns a tuple with the Affinity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffinity

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetAffinity(v string)`

SetAffinity sets Affinity field to given value.

### HasAffinity

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasAffinity() bool`

HasAffinity returns a boolean if a field has been set.

### GetAgent

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetAgent() string`

GetAgent returns the Agent field if non-nil, zero value otherwise.

### GetAgentOk

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetAgentOk() (*string, bool)`

GetAgentOk returns a tuple with the Agent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgent

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetAgent(v string)`

SetAgent sets Agent field to given value.

### HasAgent

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasAgent() bool`

HasAgent returns a boolean if a field has been set.

### GetArch

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetArch() VirtualMachineArchitecture`

GetArch returns the Arch field if non-nil, zero value otherwise.

### GetArchOk

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetArchOk() (*VirtualMachineArchitecture, bool)`

GetArchOk returns a tuple with the Arch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArch

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetArch(v VirtualMachineArchitecture)`

SetArch sets Arch field to given value.

### HasArch

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasArch() bool`

HasArch returns a boolean if a field has been set.

### GetArgs

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetArgs() string`

GetArgs returns the Args field if non-nil, zero value otherwise.

### GetArgsOk

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetArgsOk() (*string, bool)`

GetArgsOk returns a tuple with the Args field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgs

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetArgs(v string)`

SetArgs sets Args field to given value.

### HasArgs

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasArgs() bool`

HasArgs returns a boolean if a field has been set.

### GetAudio0

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetAudio0() string`

GetAudio0 returns the Audio0 field if non-nil, zero value otherwise.

### GetAudio0Ok

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetAudio0Ok() (*string, bool)`

GetAudio0Ok returns a tuple with the Audio0 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudio0

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetAudio0(v string)`

SetAudio0 sets Audio0 field to given value.

### HasAudio0

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasAudio0() bool`

HasAudio0 returns a boolean if a field has been set.

### GetAutostart

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetAutostart() float32`

GetAutostart returns the Autostart field if non-nil, zero value otherwise.

### GetAutostartOk

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetAutostartOk() (*float32, bool)`

GetAutostartOk returns a tuple with the Autostart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutostart

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetAutostart(v float32)`

SetAutostart sets Autostart field to given value.

### HasAutostart

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasAutostart() bool`

HasAutostart returns a boolean if a field has been set.

### GetBallon

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetBallon() float32`

GetBallon returns the Ballon field if non-nil, zero value otherwise.

### GetBallonOk

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetBallonOk() (*float32, bool)`

GetBallonOk returns a tuple with the Ballon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBallon

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetBallon(v float32)`

SetBallon sets Ballon field to given value.

### HasBallon

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasBallon() bool`

HasBallon returns a boolean if a field has been set.

### GetBoot

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetBoot() string`

GetBoot returns the Boot field if non-nil, zero value otherwise.

### GetBootOk

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetBootOk() (*string, bool)`

GetBootOk returns a tuple with the Boot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoot

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetBoot(v string)`

SetBoot sets Boot field to given value.

### HasBoot

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasBoot() bool`

HasBoot returns a boolean if a field has been set.

### GetBootdisk

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetBootdisk() string`

GetBootdisk returns the Bootdisk field if non-nil, zero value otherwise.

### GetBootdiskOk

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetBootdiskOk() (*string, bool)`

GetBootdiskOk returns a tuple with the Bootdisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootdisk

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetBootdisk(v string)`

SetBootdisk sets Bootdisk field to given value.

### HasBootdisk

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasBootdisk() bool`

HasBootdisk returns a boolean if a field has been set.

### GetCdrom

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetCdrom() string`

GetCdrom returns the Cdrom field if non-nil, zero value otherwise.

### GetCdromOk

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetCdromOk() (*string, bool)`

GetCdromOk returns a tuple with the Cdrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdrom

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetCdrom(v string)`

SetCdrom sets Cdrom field to given value.

### HasCdrom

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasCdrom() bool`

HasCdrom returns a boolean if a field has been set.

### GetCicustom

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetCicustom() string`

GetCicustom returns the Cicustom field if non-nil, zero value otherwise.

### GetCicustomOk

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetCicustomOk() (*string, bool)`

GetCicustomOk returns a tuple with the Cicustom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCicustom

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetCicustom(v string)`

SetCicustom sets Cicustom field to given value.

### HasCicustom

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasCicustom() bool`

HasCicustom returns a boolean if a field has been set.

### GetCipassword

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetCipassword() string`

GetCipassword returns the Cipassword field if non-nil, zero value otherwise.

### GetCipasswordOk

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetCipasswordOk() (*string, bool)`

GetCipasswordOk returns a tuple with the Cipassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCipassword

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetCipassword(v string)`

SetCipassword sets Cipassword field to given value.

### HasCipassword

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasCipassword() bool`

HasCipassword returns a boolean if a field has been set.

### GetCitype

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetCitype() VirtualMachineCloudInitType`

GetCitype returns the Citype field if non-nil, zero value otherwise.

### GetCitypeOk

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetCitypeOk() (*VirtualMachineCloudInitType, bool)`

GetCitypeOk returns a tuple with the Citype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCitype

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetCitype(v VirtualMachineCloudInitType)`

SetCitype sets Citype field to given value.

### HasCitype

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasCitype() bool`

HasCitype returns a boolean if a field has been set.

### GetCiuser

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetCiuser() string`

GetCiuser returns the Ciuser field if non-nil, zero value otherwise.

### GetCiuserOk

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetCiuserOk() (*string, bool)`

GetCiuserOk returns a tuple with the Ciuser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCiuser

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetCiuser(v string)`

SetCiuser sets Ciuser field to given value.

### HasCiuser

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasCiuser() bool`

HasCiuser returns a boolean if a field has been set.

### GetCores

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetCores() float32`

GetCores returns the Cores field if non-nil, zero value otherwise.

### GetCoresOk

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetCoresOk() (*float32, bool)`

GetCoresOk returns a tuple with the Cores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCores

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetCores(v float32)`

SetCores sets Cores field to given value.

### HasCores

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasCores() bool`

HasCores returns a boolean if a field has been set.

### GetBios

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetBios() VirtualMachineBios`

GetBios returns the Bios field if non-nil, zero value otherwise.

### GetBiosOk

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetBiosOk() (*VirtualMachineBios, bool)`

GetBiosOk returns a tuple with the Bios field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBios

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetBios(v VirtualMachineBios)`

SetBios sets Bios field to given value.

### HasBios

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasBios() bool`

HasBios returns a boolean if a field has been set.

### GetCpu

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetCpu() string`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetCpuOk() (*string, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetCpu(v string)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### GetCpulimit

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetCpulimit() float32`

GetCpulimit returns the Cpulimit field if non-nil, zero value otherwise.

### GetCpulimitOk

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetCpulimitOk() (*float32, bool)`

GetCpulimitOk returns a tuple with the Cpulimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpulimit

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetCpulimit(v float32)`

SetCpulimit sets Cpulimit field to given value.

### HasCpulimit

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasCpulimit() bool`

HasCpulimit returns a boolean if a field has been set.

### GetCpuunits

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetCpuunits() float32`

GetCpuunits returns the Cpuunits field if non-nil, zero value otherwise.

### GetCpuunitsOk

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetCpuunitsOk() (*float32, bool)`

GetCpuunitsOk returns a tuple with the Cpuunits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuunits

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetCpuunits(v float32)`

SetCpuunits sets Cpuunits field to given value.

### HasCpuunits

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasCpuunits() bool`

HasCpuunits returns a boolean if a field has been set.

### GetDescription

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDigest

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetDigest() string`

GetDigest returns the Digest field if non-nil, zero value otherwise.

### GetDigestOk

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetDigestOk() (*string, bool)`

GetDigestOk returns a tuple with the Digest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigest

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetDigest(v string)`

SetDigest sets Digest field to given value.

### HasDigest

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasDigest() bool`

HasDigest returns a boolean if a field has been set.

### GetDelete

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetDelete() string`

GetDelete returns the Delete field if non-nil, zero value otherwise.

### GetDeleteOk

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetDeleteOk() (*string, bool)`

GetDeleteOk returns a tuple with the Delete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelete

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetDelete(v string)`

SetDelete sets Delete field to given value.

### HasDelete

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasDelete() bool`

HasDelete returns a boolean if a field has been set.

### GetEfidisk0

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetEfidisk0() string`

GetEfidisk0 returns the Efidisk0 field if non-nil, zero value otherwise.

### GetEfidisk0Ok

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetEfidisk0Ok() (*string, bool)`

GetEfidisk0Ok returns a tuple with the Efidisk0 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEfidisk0

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetEfidisk0(v string)`

SetEfidisk0 sets Efidisk0 field to given value.

### HasEfidisk0

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasEfidisk0() bool`

HasEfidisk0 returns a boolean if a field has been set.

### GetFreeze

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetFreeze() bool`

GetFreeze returns the Freeze field if non-nil, zero value otherwise.

### GetFreezeOk

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetFreezeOk() (*bool, bool)`

GetFreezeOk returns a tuple with the Freeze field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeze

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetFreeze(v bool)`

SetFreeze sets Freeze field to given value.

### HasFreeze

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasFreeze() bool`

HasFreeze returns a boolean if a field has been set.

### GetHookscript

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetHookscript() string`

GetHookscript returns the Hookscript field if non-nil, zero value otherwise.

### GetHookscriptOk

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetHookscriptOk() (*string, bool)`

GetHookscriptOk returns a tuple with the Hookscript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHookscript

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetHookscript(v string)`

SetHookscript sets Hookscript field to given value.

### HasHookscript

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasHookscript() bool`

HasHookscript returns a boolean if a field has been set.

### GetHostpci0

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetHostpci0() string`

GetHostpci0 returns the Hostpci0 field if non-nil, zero value otherwise.

### GetHostpci0Ok

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetHostpci0Ok() (*string, bool)`

GetHostpci0Ok returns a tuple with the Hostpci0 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostpci0

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetHostpci0(v string)`

SetHostpci0 sets Hostpci0 field to given value.

### HasHostpci0

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasHostpci0() bool`

HasHostpci0 returns a boolean if a field has been set.

### GetHostpci1

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetHostpci1() string`

GetHostpci1 returns the Hostpci1 field if non-nil, zero value otherwise.

### GetHostpci1Ok

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetHostpci1Ok() (*string, bool)`

GetHostpci1Ok returns a tuple with the Hostpci1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostpci1

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetHostpci1(v string)`

SetHostpci1 sets Hostpci1 field to given value.

### HasHostpci1

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasHostpci1() bool`

HasHostpci1 returns a boolean if a field has been set.

### GetHostpci2

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetHostpci2() string`

GetHostpci2 returns the Hostpci2 field if non-nil, zero value otherwise.

### GetHostpci2Ok

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetHostpci2Ok() (*string, bool)`

GetHostpci2Ok returns a tuple with the Hostpci2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostpci2

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetHostpci2(v string)`

SetHostpci2 sets Hostpci2 field to given value.

### HasHostpci2

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasHostpci2() bool`

HasHostpci2 returns a boolean if a field has been set.

### GetHostpci3

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetHostpci3() string`

GetHostpci3 returns the Hostpci3 field if non-nil, zero value otherwise.

### GetHostpci3Ok

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetHostpci3Ok() (*string, bool)`

GetHostpci3Ok returns a tuple with the Hostpci3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostpci3

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetHostpci3(v string)`

SetHostpci3 sets Hostpci3 field to given value.

### HasHostpci3

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasHostpci3() bool`

HasHostpci3 returns a boolean if a field has been set.

### GetHostpci4

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetHostpci4() string`

GetHostpci4 returns the Hostpci4 field if non-nil, zero value otherwise.

### GetHostpci4Ok

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetHostpci4Ok() (*string, bool)`

GetHostpci4Ok returns a tuple with the Hostpci4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostpci4

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetHostpci4(v string)`

SetHostpci4 sets Hostpci4 field to given value.

### HasHostpci4

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasHostpci4() bool`

HasHostpci4 returns a boolean if a field has been set.

### GetHostpci5

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetHostpci5() string`

GetHostpci5 returns the Hostpci5 field if non-nil, zero value otherwise.

### GetHostpci5Ok

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetHostpci5Ok() (*string, bool)`

GetHostpci5Ok returns a tuple with the Hostpci5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostpci5

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetHostpci5(v string)`

SetHostpci5 sets Hostpci5 field to given value.

### HasHostpci5

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasHostpci5() bool`

HasHostpci5 returns a boolean if a field has been set.

### GetHostpci6

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetHostpci6() string`

GetHostpci6 returns the Hostpci6 field if non-nil, zero value otherwise.

### GetHostpci6Ok

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetHostpci6Ok() (*string, bool)`

GetHostpci6Ok returns a tuple with the Hostpci6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostpci6

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetHostpci6(v string)`

SetHostpci6 sets Hostpci6 field to given value.

### HasHostpci6

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasHostpci6() bool`

HasHostpci6 returns a boolean if a field has been set.

### GetHostpci7

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetHostpci7() string`

GetHostpci7 returns the Hostpci7 field if non-nil, zero value otherwise.

### GetHostpci7Ok

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetHostpci7Ok() (*string, bool)`

GetHostpci7Ok returns a tuple with the Hostpci7 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostpci7

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetHostpci7(v string)`

SetHostpci7 sets Hostpci7 field to given value.

### HasHostpci7

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasHostpci7() bool`

HasHostpci7 returns a boolean if a field has been set.

### GetHostpci8

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetHostpci8() string`

GetHostpci8 returns the Hostpci8 field if non-nil, zero value otherwise.

### GetHostpci8Ok

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetHostpci8Ok() (*string, bool)`

GetHostpci8Ok returns a tuple with the Hostpci8 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostpci8

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetHostpci8(v string)`

SetHostpci8 sets Hostpci8 field to given value.

### HasHostpci8

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasHostpci8() bool`

HasHostpci8 returns a boolean if a field has been set.

### GetHostpci9

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetHostpci9() string`

GetHostpci9 returns the Hostpci9 field if non-nil, zero value otherwise.

### GetHostpci9Ok

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetHostpci9Ok() (*string, bool)`

GetHostpci9Ok returns a tuple with the Hostpci9 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostpci9

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetHostpci9(v string)`

SetHostpci9 sets Hostpci9 field to given value.

### HasHostpci9

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasHostpci9() bool`

HasHostpci9 returns a boolean if a field has been set.

### GetHugepages

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetHugepages() VirtualMachineHugePages`

GetHugepages returns the Hugepages field if non-nil, zero value otherwise.

### GetHugepagesOk

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetHugepagesOk() (*VirtualMachineHugePages, bool)`

GetHugepagesOk returns a tuple with the Hugepages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHugepages

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetHugepages(v VirtualMachineHugePages)`

SetHugepages sets Hugepages field to given value.

### HasHugepages

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasHugepages() bool`

HasHugepages returns a boolean if a field has been set.

### GetIde0

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetIde0() string`

GetIde0 returns the Ide0 field if non-nil, zero value otherwise.

### GetIde0Ok

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetIde0Ok() (*string, bool)`

GetIde0Ok returns a tuple with the Ide0 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIde0

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetIde0(v string)`

SetIde0 sets Ide0 field to given value.

### HasIde0

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasIde0() bool`

HasIde0 returns a boolean if a field has been set.

### GetIde1

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetIde1() string`

GetIde1 returns the Ide1 field if non-nil, zero value otherwise.

### GetIde1Ok

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetIde1Ok() (*string, bool)`

GetIde1Ok returns a tuple with the Ide1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIde1

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetIde1(v string)`

SetIde1 sets Ide1 field to given value.

### HasIde1

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasIde1() bool`

HasIde1 returns a boolean if a field has been set.

### GetIde2

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetIde2() string`

GetIde2 returns the Ide2 field if non-nil, zero value otherwise.

### GetIde2Ok

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetIde2Ok() (*string, bool)`

GetIde2Ok returns a tuple with the Ide2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIde2

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetIde2(v string)`

SetIde2 sets Ide2 field to given value.

### HasIde2

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasIde2() bool`

HasIde2 returns a boolean if a field has been set.

### GetIde3

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetIde3() string`

GetIde3 returns the Ide3 field if non-nil, zero value otherwise.

### GetIde3Ok

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetIde3Ok() (*string, bool)`

GetIde3Ok returns a tuple with the Ide3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIde3

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetIde3(v string)`

SetIde3 sets Ide3 field to given value.

### HasIde3

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasIde3() bool`

HasIde3 returns a boolean if a field has been set.

### GetIpconfig0

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetIpconfig0() string`

GetIpconfig0 returns the Ipconfig0 field if non-nil, zero value otherwise.

### GetIpconfig0Ok

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetIpconfig0Ok() (*string, bool)`

GetIpconfig0Ok returns a tuple with the Ipconfig0 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpconfig0

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetIpconfig0(v string)`

SetIpconfig0 sets Ipconfig0 field to given value.

### HasIpconfig0

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasIpconfig0() bool`

HasIpconfig0 returns a boolean if a field has been set.

### GetIpconfig1

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetIpconfig1() string`

GetIpconfig1 returns the Ipconfig1 field if non-nil, zero value otherwise.

### GetIpconfig1Ok

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetIpconfig1Ok() (*string, bool)`

GetIpconfig1Ok returns a tuple with the Ipconfig1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpconfig1

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetIpconfig1(v string)`

SetIpconfig1 sets Ipconfig1 field to given value.

### HasIpconfig1

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasIpconfig1() bool`

HasIpconfig1 returns a boolean if a field has been set.

### GetIpconfig2

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetIpconfig2() string`

GetIpconfig2 returns the Ipconfig2 field if non-nil, zero value otherwise.

### GetIpconfig2Ok

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetIpconfig2Ok() (*string, bool)`

GetIpconfig2Ok returns a tuple with the Ipconfig2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpconfig2

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetIpconfig2(v string)`

SetIpconfig2 sets Ipconfig2 field to given value.

### HasIpconfig2

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasIpconfig2() bool`

HasIpconfig2 returns a boolean if a field has been set.

### GetIpconfig3

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetIpconfig3() string`

GetIpconfig3 returns the Ipconfig3 field if non-nil, zero value otherwise.

### GetIpconfig3Ok

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetIpconfig3Ok() (*string, bool)`

GetIpconfig3Ok returns a tuple with the Ipconfig3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpconfig3

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetIpconfig3(v string)`

SetIpconfig3 sets Ipconfig3 field to given value.

### HasIpconfig3

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasIpconfig3() bool`

HasIpconfig3 returns a boolean if a field has been set.

### GetIpconfig4

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetIpconfig4() string`

GetIpconfig4 returns the Ipconfig4 field if non-nil, zero value otherwise.

### GetIpconfig4Ok

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetIpconfig4Ok() (*string, bool)`

GetIpconfig4Ok returns a tuple with the Ipconfig4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpconfig4

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetIpconfig4(v string)`

SetIpconfig4 sets Ipconfig4 field to given value.

### HasIpconfig4

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasIpconfig4() bool`

HasIpconfig4 returns a boolean if a field has been set.

### GetIpconfig5

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetIpconfig5() string`

GetIpconfig5 returns the Ipconfig5 field if non-nil, zero value otherwise.

### GetIpconfig5Ok

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetIpconfig5Ok() (*string, bool)`

GetIpconfig5Ok returns a tuple with the Ipconfig5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpconfig5

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetIpconfig5(v string)`

SetIpconfig5 sets Ipconfig5 field to given value.

### HasIpconfig5

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasIpconfig5() bool`

HasIpconfig5 returns a boolean if a field has been set.

### GetIpconfig6

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetIpconfig6() string`

GetIpconfig6 returns the Ipconfig6 field if non-nil, zero value otherwise.

### GetIpconfig6Ok

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetIpconfig6Ok() (*string, bool)`

GetIpconfig6Ok returns a tuple with the Ipconfig6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpconfig6

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetIpconfig6(v string)`

SetIpconfig6 sets Ipconfig6 field to given value.

### HasIpconfig6

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasIpconfig6() bool`

HasIpconfig6 returns a boolean if a field has been set.

### GetIpconfig7

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetIpconfig7() string`

GetIpconfig7 returns the Ipconfig7 field if non-nil, zero value otherwise.

### GetIpconfig7Ok

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetIpconfig7Ok() (*string, bool)`

GetIpconfig7Ok returns a tuple with the Ipconfig7 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpconfig7

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetIpconfig7(v string)`

SetIpconfig7 sets Ipconfig7 field to given value.

### HasIpconfig7

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasIpconfig7() bool`

HasIpconfig7 returns a boolean if a field has been set.

### GetIvshmem

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetIvshmem() string`

GetIvshmem returns the Ivshmem field if non-nil, zero value otherwise.

### GetIvshmemOk

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetIvshmemOk() (*string, bool)`

GetIvshmemOk returns a tuple with the Ivshmem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIvshmem

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetIvshmem(v string)`

SetIvshmem sets Ivshmem field to given value.

### HasIvshmem

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasIvshmem() bool`

HasIvshmem returns a boolean if a field has been set.

### GetKeephugepages

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetKeephugepages() float32`

GetKeephugepages returns the Keephugepages field if non-nil, zero value otherwise.

### GetKeephugepagesOk

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetKeephugepagesOk() (*float32, bool)`

GetKeephugepagesOk returns a tuple with the Keephugepages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeephugepages

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetKeephugepages(v float32)`

SetKeephugepages sets Keephugepages field to given value.

### HasKeephugepages

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasKeephugepages() bool`

HasKeephugepages returns a boolean if a field has been set.

### GetKeyboard

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetKeyboard() VirtualMachineKeyboard`

GetKeyboard returns the Keyboard field if non-nil, zero value otherwise.

### GetKeyboardOk

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetKeyboardOk() (*VirtualMachineKeyboard, bool)`

GetKeyboardOk returns a tuple with the Keyboard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyboard

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetKeyboard(v VirtualMachineKeyboard)`

SetKeyboard sets Keyboard field to given value.

### HasKeyboard

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasKeyboard() bool`

HasKeyboard returns a boolean if a field has been set.

### GetKvm

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetKvm() float32`

GetKvm returns the Kvm field if non-nil, zero value otherwise.

### GetKvmOk

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetKvmOk() (*float32, bool)`

GetKvmOk returns a tuple with the Kvm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKvm

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetKvm(v float32)`

SetKvm sets Kvm field to given value.

### HasKvm

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasKvm() bool`

HasKvm returns a boolean if a field has been set.

### GetLocaltime

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetLocaltime() float32`

GetLocaltime returns the Localtime field if non-nil, zero value otherwise.

### GetLocaltimeOk

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetLocaltimeOk() (*float32, bool)`

GetLocaltimeOk returns a tuple with the Localtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocaltime

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetLocaltime(v float32)`

SetLocaltime sets Localtime field to given value.

### HasLocaltime

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasLocaltime() bool`

HasLocaltime returns a boolean if a field has been set.

### GetLock

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetLock() VirtualMachineConfigLock`

GetLock returns the Lock field if non-nil, zero value otherwise.

### GetLockOk

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetLockOk() (*VirtualMachineConfigLock, bool)`

GetLockOk returns a tuple with the Lock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLock

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetLock(v VirtualMachineConfigLock)`

SetLock sets Lock field to given value.

### HasLock

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasLock() bool`

HasLock returns a boolean if a field has been set.

### GetMachine

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetMachine() string`

GetMachine returns the Machine field if non-nil, zero value otherwise.

### GetMachineOk

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetMachineOk() (*string, bool)`

GetMachineOk returns a tuple with the Machine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachine

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetMachine(v string)`

SetMachine sets Machine field to given value.

### HasMachine

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasMachine() bool`

HasMachine returns a boolean if a field has been set.

### GetMemory

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetMemory() float32`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetMemoryOk() (*float32, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetMemory(v float32)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### GetMigrateDowntime

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetMigrateDowntime() float32`

GetMigrateDowntime returns the MigrateDowntime field if non-nil, zero value otherwise.

### GetMigrateDowntimeOk

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetMigrateDowntimeOk() (*float32, bool)`

GetMigrateDowntimeOk returns a tuple with the MigrateDowntime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigrateDowntime

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetMigrateDowntime(v float32)`

SetMigrateDowntime sets MigrateDowntime field to given value.

### HasMigrateDowntime

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasMigrateDowntime() bool`

HasMigrateDowntime returns a boolean if a field has been set.

### GetMigrateSpeed

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetMigrateSpeed() float32`

GetMigrateSpeed returns the MigrateSpeed field if non-nil, zero value otherwise.

### GetMigrateSpeedOk

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetMigrateSpeedOk() (*float32, bool)`

GetMigrateSpeedOk returns a tuple with the MigrateSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigrateSpeed

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetMigrateSpeed(v float32)`

SetMigrateSpeed sets MigrateSpeed field to given value.

### HasMigrateSpeed

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasMigrateSpeed() bool`

HasMigrateSpeed returns a boolean if a field has been set.

### GetName

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNameserver

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetNameserver() string`

GetNameserver returns the Nameserver field if non-nil, zero value otherwise.

### GetNameserverOk

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetNameserverOk() (*string, bool)`

GetNameserverOk returns a tuple with the Nameserver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameserver

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetNameserver(v string)`

SetNameserver sets Nameserver field to given value.

### HasNameserver

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasNameserver() bool`

HasNameserver returns a boolean if a field has been set.

### GetNet0

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetNet0() string`

GetNet0 returns the Net0 field if non-nil, zero value otherwise.

### GetNet0Ok

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetNet0Ok() (*string, bool)`

GetNet0Ok returns a tuple with the Net0 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet0

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetNet0(v string)`

SetNet0 sets Net0 field to given value.

### HasNet0

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasNet0() bool`

HasNet0 returns a boolean if a field has been set.

### GetNet1

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetNet1() string`

GetNet1 returns the Net1 field if non-nil, zero value otherwise.

### GetNet1Ok

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetNet1Ok() (*string, bool)`

GetNet1Ok returns a tuple with the Net1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet1

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetNet1(v string)`

SetNet1 sets Net1 field to given value.

### HasNet1

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasNet1() bool`

HasNet1 returns a boolean if a field has been set.

### GetNet2

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetNet2() string`

GetNet2 returns the Net2 field if non-nil, zero value otherwise.

### GetNet2Ok

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetNet2Ok() (*string, bool)`

GetNet2Ok returns a tuple with the Net2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet2

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetNet2(v string)`

SetNet2 sets Net2 field to given value.

### HasNet2

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasNet2() bool`

HasNet2 returns a boolean if a field has been set.

### GetNet3

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetNet3() string`

GetNet3 returns the Net3 field if non-nil, zero value otherwise.

### GetNet3Ok

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetNet3Ok() (*string, bool)`

GetNet3Ok returns a tuple with the Net3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet3

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetNet3(v string)`

SetNet3 sets Net3 field to given value.

### HasNet3

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasNet3() bool`

HasNet3 returns a boolean if a field has been set.

### GetNet4

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetNet4() string`

GetNet4 returns the Net4 field if non-nil, zero value otherwise.

### GetNet4Ok

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetNet4Ok() (*string, bool)`

GetNet4Ok returns a tuple with the Net4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet4

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetNet4(v string)`

SetNet4 sets Net4 field to given value.

### HasNet4

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasNet4() bool`

HasNet4 returns a boolean if a field has been set.

### GetNet5

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetNet5() string`

GetNet5 returns the Net5 field if non-nil, zero value otherwise.

### GetNet5Ok

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetNet5Ok() (*string, bool)`

GetNet5Ok returns a tuple with the Net5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet5

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetNet5(v string)`

SetNet5 sets Net5 field to given value.

### HasNet5

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasNet5() bool`

HasNet5 returns a boolean if a field has been set.

### GetNet6

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetNet6() string`

GetNet6 returns the Net6 field if non-nil, zero value otherwise.

### GetNet6Ok

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetNet6Ok() (*string, bool)`

GetNet6Ok returns a tuple with the Net6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet6

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetNet6(v string)`

SetNet6 sets Net6 field to given value.

### HasNet6

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasNet6() bool`

HasNet6 returns a boolean if a field has been set.

### GetNet7

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetNet7() string`

GetNet7 returns the Net7 field if non-nil, zero value otherwise.

### GetNet7Ok

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetNet7Ok() (*string, bool)`

GetNet7Ok returns a tuple with the Net7 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet7

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetNet7(v string)`

SetNet7 sets Net7 field to given value.

### HasNet7

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasNet7() bool`

HasNet7 returns a boolean if a field has been set.

### GetNuma

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetNuma() float32`

GetNuma returns the Numa field if non-nil, zero value otherwise.

### GetNumaOk

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetNumaOk() (*float32, bool)`

GetNumaOk returns a tuple with the Numa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNuma

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetNuma(v float32)`

SetNuma sets Numa field to given value.

### HasNuma

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasNuma() bool`

HasNuma returns a boolean if a field has been set.

### GetNuma0

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetNuma0() string`

GetNuma0 returns the Numa0 field if non-nil, zero value otherwise.

### GetNuma0Ok

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetNuma0Ok() (*string, bool)`

GetNuma0Ok returns a tuple with the Numa0 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNuma0

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetNuma0(v string)`

SetNuma0 sets Numa0 field to given value.

### HasNuma0

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasNuma0() bool`

HasNuma0 returns a boolean if a field has been set.

### GetNuma1

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetNuma1() string`

GetNuma1 returns the Numa1 field if non-nil, zero value otherwise.

### GetNuma1Ok

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetNuma1Ok() (*string, bool)`

GetNuma1Ok returns a tuple with the Numa1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNuma1

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetNuma1(v string)`

SetNuma1 sets Numa1 field to given value.

### HasNuma1

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasNuma1() bool`

HasNuma1 returns a boolean if a field has been set.

### GetNuma2

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetNuma2() string`

GetNuma2 returns the Numa2 field if non-nil, zero value otherwise.

### GetNuma2Ok

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetNuma2Ok() (*string, bool)`

GetNuma2Ok returns a tuple with the Numa2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNuma2

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetNuma2(v string)`

SetNuma2 sets Numa2 field to given value.

### HasNuma2

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasNuma2() bool`

HasNuma2 returns a boolean if a field has been set.

### GetNuma3

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetNuma3() string`

GetNuma3 returns the Numa3 field if non-nil, zero value otherwise.

### GetNuma3Ok

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetNuma3Ok() (*string, bool)`

GetNuma3Ok returns a tuple with the Numa3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNuma3

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetNuma3(v string)`

SetNuma3 sets Numa3 field to given value.

### HasNuma3

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasNuma3() bool`

HasNuma3 returns a boolean if a field has been set.

### GetNuma4

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetNuma4() string`

GetNuma4 returns the Numa4 field if non-nil, zero value otherwise.

### GetNuma4Ok

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetNuma4Ok() (*string, bool)`

GetNuma4Ok returns a tuple with the Numa4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNuma4

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetNuma4(v string)`

SetNuma4 sets Numa4 field to given value.

### HasNuma4

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasNuma4() bool`

HasNuma4 returns a boolean if a field has been set.

### GetNuma5

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetNuma5() string`

GetNuma5 returns the Numa5 field if non-nil, zero value otherwise.

### GetNuma5Ok

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetNuma5Ok() (*string, bool)`

GetNuma5Ok returns a tuple with the Numa5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNuma5

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetNuma5(v string)`

SetNuma5 sets Numa5 field to given value.

### HasNuma5

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasNuma5() bool`

HasNuma5 returns a boolean if a field has been set.

### GetNuma6

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetNuma6() string`

GetNuma6 returns the Numa6 field if non-nil, zero value otherwise.

### GetNuma6Ok

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetNuma6Ok() (*string, bool)`

GetNuma6Ok returns a tuple with the Numa6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNuma6

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetNuma6(v string)`

SetNuma6 sets Numa6 field to given value.

### HasNuma6

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasNuma6() bool`

HasNuma6 returns a boolean if a field has been set.

### GetNuma7

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetNuma7() string`

GetNuma7 returns the Numa7 field if non-nil, zero value otherwise.

### GetNuma7Ok

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetNuma7Ok() (*string, bool)`

GetNuma7Ok returns a tuple with the Numa7 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNuma7

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetNuma7(v string)`

SetNuma7 sets Numa7 field to given value.

### HasNuma7

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasNuma7() bool`

HasNuma7 returns a boolean if a field has been set.

### GetOnboot

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetOnboot() float32`

GetOnboot returns the Onboot field if non-nil, zero value otherwise.

### GetOnbootOk

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetOnbootOk() (*float32, bool)`

GetOnbootOk returns a tuple with the Onboot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnboot

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetOnboot(v float32)`

SetOnboot sets Onboot field to given value.

### HasOnboot

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasOnboot() bool`

HasOnboot returns a boolean if a field has been set.

### GetOstype

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetOstype() VirtualMachineOperatingSystem`

GetOstype returns the Ostype field if non-nil, zero value otherwise.

### GetOstypeOk

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetOstypeOk() (*VirtualMachineOperatingSystem, bool)`

GetOstypeOk returns a tuple with the Ostype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOstype

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetOstype(v VirtualMachineOperatingSystem)`

SetOstype sets Ostype field to given value.

### HasOstype

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasOstype() bool`

HasOstype returns a boolean if a field has been set.

### GetParallel0

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetParallel0() string`

GetParallel0 returns the Parallel0 field if non-nil, zero value otherwise.

### GetParallel0Ok

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetParallel0Ok() (*string, bool)`

GetParallel0Ok returns a tuple with the Parallel0 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParallel0

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetParallel0(v string)`

SetParallel0 sets Parallel0 field to given value.

### HasParallel0

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasParallel0() bool`

HasParallel0 returns a boolean if a field has been set.

### GetParallel1

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetParallel1() string`

GetParallel1 returns the Parallel1 field if non-nil, zero value otherwise.

### GetParallel1Ok

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetParallel1Ok() (*string, bool)`

GetParallel1Ok returns a tuple with the Parallel1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParallel1

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetParallel1(v string)`

SetParallel1 sets Parallel1 field to given value.

### HasParallel1

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasParallel1() bool`

HasParallel1 returns a boolean if a field has been set.

### GetParallel2

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetParallel2() string`

GetParallel2 returns the Parallel2 field if non-nil, zero value otherwise.

### GetParallel2Ok

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetParallel2Ok() (*string, bool)`

GetParallel2Ok returns a tuple with the Parallel2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParallel2

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetParallel2(v string)`

SetParallel2 sets Parallel2 field to given value.

### HasParallel2

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasParallel2() bool`

HasParallel2 returns a boolean if a field has been set.

### GetProtection

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetProtection() float32`

GetProtection returns the Protection field if non-nil, zero value otherwise.

### GetProtectionOk

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetProtectionOk() (*float32, bool)`

GetProtectionOk returns a tuple with the Protection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtection

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetProtection(v float32)`

SetProtection sets Protection field to given value.

### HasProtection

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasProtection() bool`

HasProtection returns a boolean if a field has been set.

### GetReboot

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetReboot() float32`

GetReboot returns the Reboot field if non-nil, zero value otherwise.

### GetRebootOk

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetRebootOk() (*float32, bool)`

GetRebootOk returns a tuple with the Reboot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReboot

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetReboot(v float32)`

SetReboot sets Reboot field to given value.

### HasReboot

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasReboot() bool`

HasReboot returns a boolean if a field has been set.

### GetRng0

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetRng0() string`

GetRng0 returns the Rng0 field if non-nil, zero value otherwise.

### GetRng0Ok

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetRng0Ok() (*string, bool)`

GetRng0Ok returns a tuple with the Rng0 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRng0

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetRng0(v string)`

SetRng0 sets Rng0 field to given value.

### HasRng0

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasRng0() bool`

HasRng0 returns a boolean if a field has been set.

### GetSata0

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetSata0() string`

GetSata0 returns the Sata0 field if non-nil, zero value otherwise.

### GetSata0Ok

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetSata0Ok() (*string, bool)`

GetSata0Ok returns a tuple with the Sata0 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSata0

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetSata0(v string)`

SetSata0 sets Sata0 field to given value.

### HasSata0

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasSata0() bool`

HasSata0 returns a boolean if a field has been set.

### GetSata1

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetSata1() string`

GetSata1 returns the Sata1 field if non-nil, zero value otherwise.

### GetSata1Ok

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetSata1Ok() (*string, bool)`

GetSata1Ok returns a tuple with the Sata1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSata1

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetSata1(v string)`

SetSata1 sets Sata1 field to given value.

### HasSata1

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasSata1() bool`

HasSata1 returns a boolean if a field has been set.

### GetSata2

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetSata2() string`

GetSata2 returns the Sata2 field if non-nil, zero value otherwise.

### GetSata2Ok

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetSata2Ok() (*string, bool)`

GetSata2Ok returns a tuple with the Sata2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSata2

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetSata2(v string)`

SetSata2 sets Sata2 field to given value.

### HasSata2

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasSata2() bool`

HasSata2 returns a boolean if a field has been set.

### GetSata3

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetSata3() string`

GetSata3 returns the Sata3 field if non-nil, zero value otherwise.

### GetSata3Ok

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetSata3Ok() (*string, bool)`

GetSata3Ok returns a tuple with the Sata3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSata3

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetSata3(v string)`

SetSata3 sets Sata3 field to given value.

### HasSata3

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasSata3() bool`

HasSata3 returns a boolean if a field has been set.

### GetSata4

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetSata4() string`

GetSata4 returns the Sata4 field if non-nil, zero value otherwise.

### GetSata4Ok

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetSata4Ok() (*string, bool)`

GetSata4Ok returns a tuple with the Sata4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSata4

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetSata4(v string)`

SetSata4 sets Sata4 field to given value.

### HasSata4

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasSata4() bool`

HasSata4 returns a boolean if a field has been set.

### GetSata5

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetSata5() string`

GetSata5 returns the Sata5 field if non-nil, zero value otherwise.

### GetSata5Ok

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetSata5Ok() (*string, bool)`

GetSata5Ok returns a tuple with the Sata5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSata5

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetSata5(v string)`

SetSata5 sets Sata5 field to given value.

### HasSata5

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasSata5() bool`

HasSata5 returns a boolean if a field has been set.

### GetScsi0

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetScsi0() string`

GetScsi0 returns the Scsi0 field if non-nil, zero value otherwise.

### GetScsi0Ok

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetScsi0Ok() (*string, bool)`

GetScsi0Ok returns a tuple with the Scsi0 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi0

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetScsi0(v string)`

SetScsi0 sets Scsi0 field to given value.

### HasScsi0

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasScsi0() bool`

HasScsi0 returns a boolean if a field has been set.

### GetScsi1

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetScsi1() string`

GetScsi1 returns the Scsi1 field if non-nil, zero value otherwise.

### GetScsi1Ok

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetScsi1Ok() (*string, bool)`

GetScsi1Ok returns a tuple with the Scsi1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi1

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetScsi1(v string)`

SetScsi1 sets Scsi1 field to given value.

### HasScsi1

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasScsi1() bool`

HasScsi1 returns a boolean if a field has been set.

### GetScsi2

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetScsi2() string`

GetScsi2 returns the Scsi2 field if non-nil, zero value otherwise.

### GetScsi2Ok

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetScsi2Ok() (*string, bool)`

GetScsi2Ok returns a tuple with the Scsi2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi2

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetScsi2(v string)`

SetScsi2 sets Scsi2 field to given value.

### HasScsi2

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasScsi2() bool`

HasScsi2 returns a boolean if a field has been set.

### GetScsi3

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetScsi3() string`

GetScsi3 returns the Scsi3 field if non-nil, zero value otherwise.

### GetScsi3Ok

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetScsi3Ok() (*string, bool)`

GetScsi3Ok returns a tuple with the Scsi3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi3

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetScsi3(v string)`

SetScsi3 sets Scsi3 field to given value.

### HasScsi3

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasScsi3() bool`

HasScsi3 returns a boolean if a field has been set.

### GetScsi4

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetScsi4() string`

GetScsi4 returns the Scsi4 field if non-nil, zero value otherwise.

### GetScsi4Ok

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetScsi4Ok() (*string, bool)`

GetScsi4Ok returns a tuple with the Scsi4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi4

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetScsi4(v string)`

SetScsi4 sets Scsi4 field to given value.

### HasScsi4

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasScsi4() bool`

HasScsi4 returns a boolean if a field has been set.

### GetScsi5

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetScsi5() string`

GetScsi5 returns the Scsi5 field if non-nil, zero value otherwise.

### GetScsi5Ok

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetScsi5Ok() (*string, bool)`

GetScsi5Ok returns a tuple with the Scsi5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi5

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetScsi5(v string)`

SetScsi5 sets Scsi5 field to given value.

### HasScsi5

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasScsi5() bool`

HasScsi5 returns a boolean if a field has been set.

### GetScsi6

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetScsi6() string`

GetScsi6 returns the Scsi6 field if non-nil, zero value otherwise.

### GetScsi6Ok

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetScsi6Ok() (*string, bool)`

GetScsi6Ok returns a tuple with the Scsi6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi6

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetScsi6(v string)`

SetScsi6 sets Scsi6 field to given value.

### HasScsi6

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasScsi6() bool`

HasScsi6 returns a boolean if a field has been set.

### GetScsi7

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetScsi7() string`

GetScsi7 returns the Scsi7 field if non-nil, zero value otherwise.

### GetScsi7Ok

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetScsi7Ok() (*string, bool)`

GetScsi7Ok returns a tuple with the Scsi7 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi7

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetScsi7(v string)`

SetScsi7 sets Scsi7 field to given value.

### HasScsi7

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasScsi7() bool`

HasScsi7 returns a boolean if a field has been set.

### GetScsi8

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetScsi8() string`

GetScsi8 returns the Scsi8 field if non-nil, zero value otherwise.

### GetScsi8Ok

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetScsi8Ok() (*string, bool)`

GetScsi8Ok returns a tuple with the Scsi8 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi8

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetScsi8(v string)`

SetScsi8 sets Scsi8 field to given value.

### HasScsi8

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasScsi8() bool`

HasScsi8 returns a boolean if a field has been set.

### GetScsi9

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetScsi9() string`

GetScsi9 returns the Scsi9 field if non-nil, zero value otherwise.

### GetScsi9Ok

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetScsi9Ok() (*string, bool)`

GetScsi9Ok returns a tuple with the Scsi9 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi9

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetScsi9(v string)`

SetScsi9 sets Scsi9 field to given value.

### HasScsi9

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasScsi9() bool`

HasScsi9 returns a boolean if a field has been set.

### GetScsi10

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetScsi10() string`

GetScsi10 returns the Scsi10 field if non-nil, zero value otherwise.

### GetScsi10Ok

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetScsi10Ok() (*string, bool)`

GetScsi10Ok returns a tuple with the Scsi10 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi10

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetScsi10(v string)`

SetScsi10 sets Scsi10 field to given value.

### HasScsi10

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasScsi10() bool`

HasScsi10 returns a boolean if a field has been set.

### GetScsi11

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetScsi11() string`

GetScsi11 returns the Scsi11 field if non-nil, zero value otherwise.

### GetScsi11Ok

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetScsi11Ok() (*string, bool)`

GetScsi11Ok returns a tuple with the Scsi11 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi11

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetScsi11(v string)`

SetScsi11 sets Scsi11 field to given value.

### HasScsi11

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasScsi11() bool`

HasScsi11 returns a boolean if a field has been set.

### GetScsi12

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetScsi12() string`

GetScsi12 returns the Scsi12 field if non-nil, zero value otherwise.

### GetScsi12Ok

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetScsi12Ok() (*string, bool)`

GetScsi12Ok returns a tuple with the Scsi12 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi12

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetScsi12(v string)`

SetScsi12 sets Scsi12 field to given value.

### HasScsi12

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasScsi12() bool`

HasScsi12 returns a boolean if a field has been set.

### GetScsi13

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetScsi13() string`

GetScsi13 returns the Scsi13 field if non-nil, zero value otherwise.

### GetScsi13Ok

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetScsi13Ok() (*string, bool)`

GetScsi13Ok returns a tuple with the Scsi13 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi13

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetScsi13(v string)`

SetScsi13 sets Scsi13 field to given value.

### HasScsi13

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasScsi13() bool`

HasScsi13 returns a boolean if a field has been set.

### GetScsi14

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetScsi14() string`

GetScsi14 returns the Scsi14 field if non-nil, zero value otherwise.

### GetScsi14Ok

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetScsi14Ok() (*string, bool)`

GetScsi14Ok returns a tuple with the Scsi14 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi14

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetScsi14(v string)`

SetScsi14 sets Scsi14 field to given value.

### HasScsi14

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasScsi14() bool`

HasScsi14 returns a boolean if a field has been set.

### GetScsi15

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetScsi15() string`

GetScsi15 returns the Scsi15 field if non-nil, zero value otherwise.

### GetScsi15Ok

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetScsi15Ok() (*string, bool)`

GetScsi15Ok returns a tuple with the Scsi15 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi15

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetScsi15(v string)`

SetScsi15 sets Scsi15 field to given value.

### HasScsi15

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasScsi15() bool`

HasScsi15 returns a boolean if a field has been set.

### GetScsi16

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetScsi16() string`

GetScsi16 returns the Scsi16 field if non-nil, zero value otherwise.

### GetScsi16Ok

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetScsi16Ok() (*string, bool)`

GetScsi16Ok returns a tuple with the Scsi16 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi16

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetScsi16(v string)`

SetScsi16 sets Scsi16 field to given value.

### HasScsi16

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasScsi16() bool`

HasScsi16 returns a boolean if a field has been set.

### GetScsi17

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetScsi17() string`

GetScsi17 returns the Scsi17 field if non-nil, zero value otherwise.

### GetScsi17Ok

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetScsi17Ok() (*string, bool)`

GetScsi17Ok returns a tuple with the Scsi17 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi17

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetScsi17(v string)`

SetScsi17 sets Scsi17 field to given value.

### HasScsi17

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasScsi17() bool`

HasScsi17 returns a boolean if a field has been set.

### GetScsi18

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetScsi18() string`

GetScsi18 returns the Scsi18 field if non-nil, zero value otherwise.

### GetScsi18Ok

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetScsi18Ok() (*string, bool)`

GetScsi18Ok returns a tuple with the Scsi18 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi18

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetScsi18(v string)`

SetScsi18 sets Scsi18 field to given value.

### HasScsi18

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasScsi18() bool`

HasScsi18 returns a boolean if a field has been set.

### GetScsi19

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetScsi19() string`

GetScsi19 returns the Scsi19 field if non-nil, zero value otherwise.

### GetScsi19Ok

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetScsi19Ok() (*string, bool)`

GetScsi19Ok returns a tuple with the Scsi19 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi19

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetScsi19(v string)`

SetScsi19 sets Scsi19 field to given value.

### HasScsi19

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasScsi19() bool`

HasScsi19 returns a boolean if a field has been set.

### GetScsi20

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetScsi20() string`

GetScsi20 returns the Scsi20 field if non-nil, zero value otherwise.

### GetScsi20Ok

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetScsi20Ok() (*string, bool)`

GetScsi20Ok returns a tuple with the Scsi20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi20

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetScsi20(v string)`

SetScsi20 sets Scsi20 field to given value.

### HasScsi20

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasScsi20() bool`

HasScsi20 returns a boolean if a field has been set.

### GetScsi21

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetScsi21() string`

GetScsi21 returns the Scsi21 field if non-nil, zero value otherwise.

### GetScsi21Ok

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetScsi21Ok() (*string, bool)`

GetScsi21Ok returns a tuple with the Scsi21 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi21

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetScsi21(v string)`

SetScsi21 sets Scsi21 field to given value.

### HasScsi21

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasScsi21() bool`

HasScsi21 returns a boolean if a field has been set.

### GetScsi22

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetScsi22() string`

GetScsi22 returns the Scsi22 field if non-nil, zero value otherwise.

### GetScsi22Ok

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetScsi22Ok() (*string, bool)`

GetScsi22Ok returns a tuple with the Scsi22 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi22

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetScsi22(v string)`

SetScsi22 sets Scsi22 field to given value.

### HasScsi22

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasScsi22() bool`

HasScsi22 returns a boolean if a field has been set.

### GetScsi23

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetScsi23() string`

GetScsi23 returns the Scsi23 field if non-nil, zero value otherwise.

### GetScsi23Ok

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetScsi23Ok() (*string, bool)`

GetScsi23Ok returns a tuple with the Scsi23 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi23

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetScsi23(v string)`

SetScsi23 sets Scsi23 field to given value.

### HasScsi23

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasScsi23() bool`

HasScsi23 returns a boolean if a field has been set.

### GetScsi24

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetScsi24() string`

GetScsi24 returns the Scsi24 field if non-nil, zero value otherwise.

### GetScsi24Ok

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetScsi24Ok() (*string, bool)`

GetScsi24Ok returns a tuple with the Scsi24 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi24

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetScsi24(v string)`

SetScsi24 sets Scsi24 field to given value.

### HasScsi24

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasScsi24() bool`

HasScsi24 returns a boolean if a field has been set.

### GetScsi25

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetScsi25() string`

GetScsi25 returns the Scsi25 field if non-nil, zero value otherwise.

### GetScsi25Ok

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetScsi25Ok() (*string, bool)`

GetScsi25Ok returns a tuple with the Scsi25 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi25

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetScsi25(v string)`

SetScsi25 sets Scsi25 field to given value.

### HasScsi25

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasScsi25() bool`

HasScsi25 returns a boolean if a field has been set.

### GetScsi26

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetScsi26() string`

GetScsi26 returns the Scsi26 field if non-nil, zero value otherwise.

### GetScsi26Ok

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetScsi26Ok() (*string, bool)`

GetScsi26Ok returns a tuple with the Scsi26 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi26

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetScsi26(v string)`

SetScsi26 sets Scsi26 field to given value.

### HasScsi26

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasScsi26() bool`

HasScsi26 returns a boolean if a field has been set.

### GetScsi27

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetScsi27() string`

GetScsi27 returns the Scsi27 field if non-nil, zero value otherwise.

### GetScsi27Ok

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetScsi27Ok() (*string, bool)`

GetScsi27Ok returns a tuple with the Scsi27 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi27

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetScsi27(v string)`

SetScsi27 sets Scsi27 field to given value.

### HasScsi27

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasScsi27() bool`

HasScsi27 returns a boolean if a field has been set.

### GetScsi28

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetScsi28() string`

GetScsi28 returns the Scsi28 field if non-nil, zero value otherwise.

### GetScsi28Ok

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetScsi28Ok() (*string, bool)`

GetScsi28Ok returns a tuple with the Scsi28 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi28

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetScsi28(v string)`

SetScsi28 sets Scsi28 field to given value.

### HasScsi28

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasScsi28() bool`

HasScsi28 returns a boolean if a field has been set.

### GetScsi29

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetScsi29() string`

GetScsi29 returns the Scsi29 field if non-nil, zero value otherwise.

### GetScsi29Ok

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetScsi29Ok() (*string, bool)`

GetScsi29Ok returns a tuple with the Scsi29 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi29

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetScsi29(v string)`

SetScsi29 sets Scsi29 field to given value.

### HasScsi29

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasScsi29() bool`

HasScsi29 returns a boolean if a field has been set.

### GetScsi30

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetScsi30() string`

GetScsi30 returns the Scsi30 field if non-nil, zero value otherwise.

### GetScsi30Ok

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetScsi30Ok() (*string, bool)`

GetScsi30Ok returns a tuple with the Scsi30 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi30

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetScsi30(v string)`

SetScsi30 sets Scsi30 field to given value.

### HasScsi30

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasScsi30() bool`

HasScsi30 returns a boolean if a field has been set.

### GetScsihw

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetScsihw() VirtualMachineScsiControllerType`

GetScsihw returns the Scsihw field if non-nil, zero value otherwise.

### GetScsihwOk

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetScsihwOk() (*VirtualMachineScsiControllerType, bool)`

GetScsihwOk returns a tuple with the Scsihw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsihw

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetScsihw(v VirtualMachineScsiControllerType)`

SetScsihw sets Scsihw field to given value.

### HasScsihw

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasScsihw() bool`

HasScsihw returns a boolean if a field has been set.

### GetSearchdomain

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetSearchdomain() string`

GetSearchdomain returns the Searchdomain field if non-nil, zero value otherwise.

### GetSearchdomainOk

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetSearchdomainOk() (*string, bool)`

GetSearchdomainOk returns a tuple with the Searchdomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchdomain

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetSearchdomain(v string)`

SetSearchdomain sets Searchdomain field to given value.

### HasSearchdomain

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasSearchdomain() bool`

HasSearchdomain returns a boolean if a field has been set.

### GetSerial0

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetSerial0() string`

GetSerial0 returns the Serial0 field if non-nil, zero value otherwise.

### GetSerial0Ok

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetSerial0Ok() (*string, bool)`

GetSerial0Ok returns a tuple with the Serial0 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial0

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetSerial0(v string)`

SetSerial0 sets Serial0 field to given value.

### HasSerial0

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasSerial0() bool`

HasSerial0 returns a boolean if a field has been set.

### GetSerial1

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetSerial1() string`

GetSerial1 returns the Serial1 field if non-nil, zero value otherwise.

### GetSerial1Ok

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetSerial1Ok() (*string, bool)`

GetSerial1Ok returns a tuple with the Serial1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial1

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetSerial1(v string)`

SetSerial1 sets Serial1 field to given value.

### HasSerial1

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasSerial1() bool`

HasSerial1 returns a boolean if a field has been set.

### GetSerial2

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetSerial2() string`

GetSerial2 returns the Serial2 field if non-nil, zero value otherwise.

### GetSerial2Ok

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetSerial2Ok() (*string, bool)`

GetSerial2Ok returns a tuple with the Serial2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial2

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetSerial2(v string)`

SetSerial2 sets Serial2 field to given value.

### HasSerial2

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasSerial2() bool`

HasSerial2 returns a boolean if a field has been set.

### GetSerial3

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetSerial3() string`

GetSerial3 returns the Serial3 field if non-nil, zero value otherwise.

### GetSerial3Ok

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetSerial3Ok() (*string, bool)`

GetSerial3Ok returns a tuple with the Serial3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial3

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetSerial3(v string)`

SetSerial3 sets Serial3 field to given value.

### HasSerial3

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasSerial3() bool`

HasSerial3 returns a boolean if a field has been set.

### GetShares

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetShares() float32`

GetShares returns the Shares field if non-nil, zero value otherwise.

### GetSharesOk

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetSharesOk() (*float32, bool)`

GetSharesOk returns a tuple with the Shares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShares

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetShares(v float32)`

SetShares sets Shares field to given value.

### HasShares

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasShares() bool`

HasShares returns a boolean if a field has been set.

### GetSmbios1

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetSmbios1() string`

GetSmbios1 returns the Smbios1 field if non-nil, zero value otherwise.

### GetSmbios1Ok

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetSmbios1Ok() (*string, bool)`

GetSmbios1Ok returns a tuple with the Smbios1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmbios1

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetSmbios1(v string)`

SetSmbios1 sets Smbios1 field to given value.

### HasSmbios1

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasSmbios1() bool`

HasSmbios1 returns a boolean if a field has been set.

### GetSockets

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetSockets() float32`

GetSockets returns the Sockets field if non-nil, zero value otherwise.

### GetSocketsOk

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetSocketsOk() (*float32, bool)`

GetSocketsOk returns a tuple with the Sockets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSockets

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetSockets(v float32)`

SetSockets sets Sockets field to given value.

### HasSockets

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasSockets() bool`

HasSockets returns a boolean if a field has been set.

### GetSpiceEnhancements

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetSpiceEnhancements() string`

GetSpiceEnhancements returns the SpiceEnhancements field if non-nil, zero value otherwise.

### GetSpiceEnhancementsOk

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetSpiceEnhancementsOk() (*string, bool)`

GetSpiceEnhancementsOk returns a tuple with the SpiceEnhancements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpiceEnhancements

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetSpiceEnhancements(v string)`

SetSpiceEnhancements sets SpiceEnhancements field to given value.

### HasSpiceEnhancements

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasSpiceEnhancements() bool`

HasSpiceEnhancements returns a boolean if a field has been set.

### GetSshkeys

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetSshkeys() string`

GetSshkeys returns the Sshkeys field if non-nil, zero value otherwise.

### GetSshkeysOk

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetSshkeysOk() (*string, bool)`

GetSshkeysOk returns a tuple with the Sshkeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshkeys

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetSshkeys(v string)`

SetSshkeys sets Sshkeys field to given value.

### HasSshkeys

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasSshkeys() bool`

HasSshkeys returns a boolean if a field has been set.

### GetStartdate

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetStartdate() string`

GetStartdate returns the Startdate field if non-nil, zero value otherwise.

### GetStartdateOk

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetStartdateOk() (*string, bool)`

GetStartdateOk returns a tuple with the Startdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartdate

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetStartdate(v string)`

SetStartdate sets Startdate field to given value.

### HasStartdate

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasStartdate() bool`

HasStartdate returns a boolean if a field has been set.

### GetStartup

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetStartup() string`

GetStartup returns the Startup field if non-nil, zero value otherwise.

### GetStartupOk

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetStartupOk() (*string, bool)`

GetStartupOk returns a tuple with the Startup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartup

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetStartup(v string)`

SetStartup sets Startup field to given value.

### HasStartup

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasStartup() bool`

HasStartup returns a boolean if a field has been set.

### GetTablet

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetTablet() float32`

GetTablet returns the Tablet field if non-nil, zero value otherwise.

### GetTabletOk

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetTabletOk() (*float32, bool)`

GetTabletOk returns a tuple with the Tablet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTablet

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetTablet(v float32)`

SetTablet sets Tablet field to given value.

### HasTablet

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasTablet() bool`

HasTablet returns a boolean if a field has been set.

### GetTags

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTemplate

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetTemplate() float32`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetTemplateOk() (*float32, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetTemplate(v float32)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetTpmstate0

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetTpmstate0() string`

GetTpmstate0 returns the Tpmstate0 field if non-nil, zero value otherwise.

### GetTpmstate0Ok

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetTpmstate0Ok() (*string, bool)`

GetTpmstate0Ok returns a tuple with the Tpmstate0 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTpmstate0

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetTpmstate0(v string)`

SetTpmstate0 sets Tpmstate0 field to given value.

### HasTpmstate0

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasTpmstate0() bool`

HasTpmstate0 returns a boolean if a field has been set.

### GetUsb0

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetUsb0() string`

GetUsb0 returns the Usb0 field if non-nil, zero value otherwise.

### GetUsb0Ok

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetUsb0Ok() (*string, bool)`

GetUsb0Ok returns a tuple with the Usb0 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsb0

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetUsb0(v string)`

SetUsb0 sets Usb0 field to given value.

### HasUsb0

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasUsb0() bool`

HasUsb0 returns a boolean if a field has been set.

### GetUsb1

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetUsb1() string`

GetUsb1 returns the Usb1 field if non-nil, zero value otherwise.

### GetUsb1Ok

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetUsb1Ok() (*string, bool)`

GetUsb1Ok returns a tuple with the Usb1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsb1

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetUsb1(v string)`

SetUsb1 sets Usb1 field to given value.

### HasUsb1

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasUsb1() bool`

HasUsb1 returns a boolean if a field has been set.

### GetUsb2

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetUsb2() string`

GetUsb2 returns the Usb2 field if non-nil, zero value otherwise.

### GetUsb2Ok

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetUsb2Ok() (*string, bool)`

GetUsb2Ok returns a tuple with the Usb2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsb2

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetUsb2(v string)`

SetUsb2 sets Usb2 field to given value.

### HasUsb2

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasUsb2() bool`

HasUsb2 returns a boolean if a field has been set.

### GetUsb3

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetUsb3() string`

GetUsb3 returns the Usb3 field if non-nil, zero value otherwise.

### GetUsb3Ok

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetUsb3Ok() (*string, bool)`

GetUsb3Ok returns a tuple with the Usb3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsb3

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetUsb3(v string)`

SetUsb3 sets Usb3 field to given value.

### HasUsb3

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasUsb3() bool`

HasUsb3 returns a boolean if a field has been set.

### GetUsb4

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetUsb4() string`

GetUsb4 returns the Usb4 field if non-nil, zero value otherwise.

### GetUsb4Ok

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetUsb4Ok() (*string, bool)`

GetUsb4Ok returns a tuple with the Usb4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsb4

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetUsb4(v string)`

SetUsb4 sets Usb4 field to given value.

### HasUsb4

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasUsb4() bool`

HasUsb4 returns a boolean if a field has been set.

### GetUsb5

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetUsb5() string`

GetUsb5 returns the Usb5 field if non-nil, zero value otherwise.

### GetUsb5Ok

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetUsb5Ok() (*string, bool)`

GetUsb5Ok returns a tuple with the Usb5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsb5

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetUsb5(v string)`

SetUsb5 sets Usb5 field to given value.

### HasUsb5

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasUsb5() bool`

HasUsb5 returns a boolean if a field has been set.

### GetUsb6

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetUsb6() string`

GetUsb6 returns the Usb6 field if non-nil, zero value otherwise.

### GetUsb6Ok

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetUsb6Ok() (*string, bool)`

GetUsb6Ok returns a tuple with the Usb6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsb6

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetUsb6(v string)`

SetUsb6 sets Usb6 field to given value.

### HasUsb6

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasUsb6() bool`

HasUsb6 returns a boolean if a field has been set.

### GetUsb7

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetUsb7() string`

GetUsb7 returns the Usb7 field if non-nil, zero value otherwise.

### GetUsb7Ok

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetUsb7Ok() (*string, bool)`

GetUsb7Ok returns a tuple with the Usb7 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsb7

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetUsb7(v string)`

SetUsb7 sets Usb7 field to given value.

### HasUsb7

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasUsb7() bool`

HasUsb7 returns a boolean if a field has been set.

### GetUsb8

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetUsb8() string`

GetUsb8 returns the Usb8 field if non-nil, zero value otherwise.

### GetUsb8Ok

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetUsb8Ok() (*string, bool)`

GetUsb8Ok returns a tuple with the Usb8 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsb8

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetUsb8(v string)`

SetUsb8 sets Usb8 field to given value.

### HasUsb8

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasUsb8() bool`

HasUsb8 returns a boolean if a field has been set.

### GetUsb9

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetUsb9() string`

GetUsb9 returns the Usb9 field if non-nil, zero value otherwise.

### GetUsb9Ok

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetUsb9Ok() (*string, bool)`

GetUsb9Ok returns a tuple with the Usb9 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsb9

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetUsb9(v string)`

SetUsb9 sets Usb9 field to given value.

### HasUsb9

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasUsb9() bool`

HasUsb9 returns a boolean if a field has been set.

### GetUsb10

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetUsb10() string`

GetUsb10 returns the Usb10 field if non-nil, zero value otherwise.

### GetUsb10Ok

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetUsb10Ok() (*string, bool)`

GetUsb10Ok returns a tuple with the Usb10 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsb10

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetUsb10(v string)`

SetUsb10 sets Usb10 field to given value.

### HasUsb10

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasUsb10() bool`

HasUsb10 returns a boolean if a field has been set.

### GetUsb11

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetUsb11() string`

GetUsb11 returns the Usb11 field if non-nil, zero value otherwise.

### GetUsb11Ok

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetUsb11Ok() (*string, bool)`

GetUsb11Ok returns a tuple with the Usb11 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsb11

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetUsb11(v string)`

SetUsb11 sets Usb11 field to given value.

### HasUsb11

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasUsb11() bool`

HasUsb11 returns a boolean if a field has been set.

### GetUsb12

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetUsb12() string`

GetUsb12 returns the Usb12 field if non-nil, zero value otherwise.

### GetUsb12Ok

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetUsb12Ok() (*string, bool)`

GetUsb12Ok returns a tuple with the Usb12 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsb12

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetUsb12(v string)`

SetUsb12 sets Usb12 field to given value.

### HasUsb12

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasUsb12() bool`

HasUsb12 returns a boolean if a field has been set.

### GetUsb13

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetUsb13() string`

GetUsb13 returns the Usb13 field if non-nil, zero value otherwise.

### GetUsb13Ok

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetUsb13Ok() (*string, bool)`

GetUsb13Ok returns a tuple with the Usb13 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsb13

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetUsb13(v string)`

SetUsb13 sets Usb13 field to given value.

### HasUsb13

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasUsb13() bool`

HasUsb13 returns a boolean if a field has been set.

### GetUsb14

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetUsb14() string`

GetUsb14 returns the Usb14 field if non-nil, zero value otherwise.

### GetUsb14Ok

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetUsb14Ok() (*string, bool)`

GetUsb14Ok returns a tuple with the Usb14 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsb14

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetUsb14(v string)`

SetUsb14 sets Usb14 field to given value.

### HasUsb14

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasUsb14() bool`

HasUsb14 returns a boolean if a field has been set.

### GetHotplug

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetHotplug() string`

GetHotplug returns the Hotplug field if non-nil, zero value otherwise.

### GetHotplugOk

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetHotplugOk() (*string, bool)`

GetHotplugOk returns a tuple with the Hotplug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHotplug

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetHotplug(v string)`

SetHotplug sets Hotplug field to given value.

### HasHotplug

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasHotplug() bool`

HasHotplug returns a boolean if a field has been set.

### GetVcpus

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetVcpus() float32`

GetVcpus returns the Vcpus field if non-nil, zero value otherwise.

### GetVcpusOk

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetVcpusOk() (*float32, bool)`

GetVcpusOk returns a tuple with the Vcpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcpus

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetVcpus(v float32)`

SetVcpus sets Vcpus field to given value.

### HasVcpus

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasVcpus() bool`

HasVcpus returns a boolean if a field has been set.

### GetVga

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetVga() string`

GetVga returns the Vga field if non-nil, zero value otherwise.

### GetVgaOk

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetVgaOk() (*string, bool)`

GetVgaOk returns a tuple with the Vga field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVga

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetVga(v string)`

SetVga sets Vga field to given value.

### HasVga

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasVga() bool`

HasVga returns a boolean if a field has been set.

### GetVirtio0

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetVirtio0() string`

GetVirtio0 returns the Virtio0 field if non-nil, zero value otherwise.

### GetVirtio0Ok

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetVirtio0Ok() (*string, bool)`

GetVirtio0Ok returns a tuple with the Virtio0 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtio0

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetVirtio0(v string)`

SetVirtio0 sets Virtio0 field to given value.

### HasVirtio0

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasVirtio0() bool`

HasVirtio0 returns a boolean if a field has been set.

### GetVirtio1

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetVirtio1() string`

GetVirtio1 returns the Virtio1 field if non-nil, zero value otherwise.

### GetVirtio1Ok

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetVirtio1Ok() (*string, bool)`

GetVirtio1Ok returns a tuple with the Virtio1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtio1

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetVirtio1(v string)`

SetVirtio1 sets Virtio1 field to given value.

### HasVirtio1

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasVirtio1() bool`

HasVirtio1 returns a boolean if a field has been set.

### GetVirtio2

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetVirtio2() string`

GetVirtio2 returns the Virtio2 field if non-nil, zero value otherwise.

### GetVirtio2Ok

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetVirtio2Ok() (*string, bool)`

GetVirtio2Ok returns a tuple with the Virtio2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtio2

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetVirtio2(v string)`

SetVirtio2 sets Virtio2 field to given value.

### HasVirtio2

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasVirtio2() bool`

HasVirtio2 returns a boolean if a field has been set.

### GetVirtio3

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetVirtio3() string`

GetVirtio3 returns the Virtio3 field if non-nil, zero value otherwise.

### GetVirtio3Ok

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetVirtio3Ok() (*string, bool)`

GetVirtio3Ok returns a tuple with the Virtio3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtio3

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetVirtio3(v string)`

SetVirtio3 sets Virtio3 field to given value.

### HasVirtio3

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasVirtio3() bool`

HasVirtio3 returns a boolean if a field has been set.

### GetVirtio4

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetVirtio4() string`

GetVirtio4 returns the Virtio4 field if non-nil, zero value otherwise.

### GetVirtio4Ok

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetVirtio4Ok() (*string, bool)`

GetVirtio4Ok returns a tuple with the Virtio4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtio4

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetVirtio4(v string)`

SetVirtio4 sets Virtio4 field to given value.

### HasVirtio4

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasVirtio4() bool`

HasVirtio4 returns a boolean if a field has been set.

### GetVirtio5

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetVirtio5() string`

GetVirtio5 returns the Virtio5 field if non-nil, zero value otherwise.

### GetVirtio5Ok

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetVirtio5Ok() (*string, bool)`

GetVirtio5Ok returns a tuple with the Virtio5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtio5

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetVirtio5(v string)`

SetVirtio5 sets Virtio5 field to given value.

### HasVirtio5

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasVirtio5() bool`

HasVirtio5 returns a boolean if a field has been set.

### GetVirtio6

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetVirtio6() string`

GetVirtio6 returns the Virtio6 field if non-nil, zero value otherwise.

### GetVirtio6Ok

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetVirtio6Ok() (*string, bool)`

GetVirtio6Ok returns a tuple with the Virtio6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtio6

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetVirtio6(v string)`

SetVirtio6 sets Virtio6 field to given value.

### HasVirtio6

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasVirtio6() bool`

HasVirtio6 returns a boolean if a field has been set.

### GetVirtio7

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetVirtio7() string`

GetVirtio7 returns the Virtio7 field if non-nil, zero value otherwise.

### GetVirtio7Ok

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetVirtio7Ok() (*string, bool)`

GetVirtio7Ok returns a tuple with the Virtio7 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtio7

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetVirtio7(v string)`

SetVirtio7 sets Virtio7 field to given value.

### HasVirtio7

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasVirtio7() bool`

HasVirtio7 returns a boolean if a field has been set.

### GetVirtio8

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetVirtio8() string`

GetVirtio8 returns the Virtio8 field if non-nil, zero value otherwise.

### GetVirtio8Ok

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetVirtio8Ok() (*string, bool)`

GetVirtio8Ok returns a tuple with the Virtio8 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtio8

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetVirtio8(v string)`

SetVirtio8 sets Virtio8 field to given value.

### HasVirtio8

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasVirtio8() bool`

HasVirtio8 returns a boolean if a field has been set.

### GetVirtio9

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetVirtio9() string`

GetVirtio9 returns the Virtio9 field if non-nil, zero value otherwise.

### GetVirtio9Ok

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetVirtio9Ok() (*string, bool)`

GetVirtio9Ok returns a tuple with the Virtio9 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtio9

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetVirtio9(v string)`

SetVirtio9 sets Virtio9 field to given value.

### HasVirtio9

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasVirtio9() bool`

HasVirtio9 returns a boolean if a field has been set.

### GetVirtio10

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetVirtio10() string`

GetVirtio10 returns the Virtio10 field if non-nil, zero value otherwise.

### GetVirtio10Ok

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetVirtio10Ok() (*string, bool)`

GetVirtio10Ok returns a tuple with the Virtio10 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtio10

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetVirtio10(v string)`

SetVirtio10 sets Virtio10 field to given value.

### HasVirtio10

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasVirtio10() bool`

HasVirtio10 returns a boolean if a field has been set.

### GetVirtio11

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetVirtio11() string`

GetVirtio11 returns the Virtio11 field if non-nil, zero value otherwise.

### GetVirtio11Ok

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetVirtio11Ok() (*string, bool)`

GetVirtio11Ok returns a tuple with the Virtio11 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtio11

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetVirtio11(v string)`

SetVirtio11 sets Virtio11 field to given value.

### HasVirtio11

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasVirtio11() bool`

HasVirtio11 returns a boolean if a field has been set.

### GetVirtio12

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetVirtio12() string`

GetVirtio12 returns the Virtio12 field if non-nil, zero value otherwise.

### GetVirtio12Ok

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetVirtio12Ok() (*string, bool)`

GetVirtio12Ok returns a tuple with the Virtio12 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtio12

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetVirtio12(v string)`

SetVirtio12 sets Virtio12 field to given value.

### HasVirtio12

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasVirtio12() bool`

HasVirtio12 returns a boolean if a field has been set.

### GetVirtio13

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetVirtio13() string`

GetVirtio13 returns the Virtio13 field if non-nil, zero value otherwise.

### GetVirtio13Ok

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetVirtio13Ok() (*string, bool)`

GetVirtio13Ok returns a tuple with the Virtio13 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtio13

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetVirtio13(v string)`

SetVirtio13 sets Virtio13 field to given value.

### HasVirtio13

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasVirtio13() bool`

HasVirtio13 returns a boolean if a field has been set.

### GetVirtio14

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetVirtio14() string`

GetVirtio14 returns the Virtio14 field if non-nil, zero value otherwise.

### GetVirtio14Ok

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetVirtio14Ok() (*string, bool)`

GetVirtio14Ok returns a tuple with the Virtio14 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtio14

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetVirtio14(v string)`

SetVirtio14 sets Virtio14 field to given value.

### HasVirtio14

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasVirtio14() bool`

HasVirtio14 returns a boolean if a field has been set.

### GetVirtio15

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetVirtio15() string`

GetVirtio15 returns the Virtio15 field if non-nil, zero value otherwise.

### GetVirtio15Ok

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetVirtio15Ok() (*string, bool)`

GetVirtio15Ok returns a tuple with the Virtio15 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtio15

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetVirtio15(v string)`

SetVirtio15 sets Virtio15 field to given value.

### HasVirtio15

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasVirtio15() bool`

HasVirtio15 returns a boolean if a field has been set.

### GetVmgenid

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetVmgenid() string`

GetVmgenid returns the Vmgenid field if non-nil, zero value otherwise.

### GetVmgenidOk

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetVmgenidOk() (*string, bool)`

GetVmgenidOk returns a tuple with the Vmgenid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmgenid

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetVmgenid(v string)`

SetVmgenid sets Vmgenid field to given value.

### HasVmgenid

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasVmgenid() bool`

HasVmgenid returns a boolean if a field has been set.

### GetVmstatestorage

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetVmstatestorage() string`

GetVmstatestorage returns the Vmstatestorage field if non-nil, zero value otherwise.

### GetVmstatestorageOk

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetVmstatestorageOk() (*string, bool)`

GetVmstatestorageOk returns a tuple with the Vmstatestorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmstatestorage

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetVmstatestorage(v string)`

SetVmstatestorage sets Vmstatestorage field to given value.

### HasVmstatestorage

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasVmstatestorage() bool`

HasVmstatestorage returns a boolean if a field has been set.

### GetWatchdog

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetWatchdog() string`

GetWatchdog returns the Watchdog field if non-nil, zero value otherwise.

### GetWatchdogOk

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) GetWatchdogOk() (*string, bool)`

GetWatchdogOk returns a tuple with the Watchdog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatchdog

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) SetWatchdog(v string)`

SetWatchdog sets Watchdog field to given value.

### HasWatchdog

`func (o *ApplyVirtualMachineConfigurationSyncRequestContent) HasWatchdog() bool`

HasWatchdog returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


