# ApplyVirtualMachineConfigurationAsyncRequestContent

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
**Bios** | Pointer to [**VirtualMachineBios**](VirtualMachineBios.md) |  | [optional] 
**Bootdisk** | Pointer to **string** | The boot disk of the virtual machine. | [optional] 
**Cdrom** | Pointer to **string** | The CD-ROM device and its configuration. An alias for option ide2 | [optional] 
**Cicustom** | Pointer to **string** | Specify custom cloud-init files to be used at start. | [optional] 
**Cipassword** | Pointer to **string** | The password for the cloud-init user. | [optional] 
**Citype** | Pointer to [**VirtualMachineCloudInitType**](VirtualMachineCloudInitType.md) |  | [optional] 
**Ciuser** | Pointer to **string** | The cloud-init user. | [optional] 
**Cores** | Pointer to **float32** | Number of cores per socket. | [optional] 
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

### NewApplyVirtualMachineConfigurationAsyncRequestContent

`func NewApplyVirtualMachineConfigurationAsyncRequestContent() *ApplyVirtualMachineConfigurationAsyncRequestContent`

NewApplyVirtualMachineConfigurationAsyncRequestContent instantiates a new ApplyVirtualMachineConfigurationAsyncRequestContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplyVirtualMachineConfigurationAsyncRequestContentWithDefaults

`func NewApplyVirtualMachineConfigurationAsyncRequestContentWithDefaults() *ApplyVirtualMachineConfigurationAsyncRequestContent`

NewApplyVirtualMachineConfigurationAsyncRequestContentWithDefaults instantiates a new ApplyVirtualMachineConfigurationAsyncRequestContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcpi

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetAcpi() float32`

GetAcpi returns the Acpi field if non-nil, zero value otherwise.

### GetAcpiOk

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetAcpiOk() (*float32, bool)`

GetAcpiOk returns a tuple with the Acpi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcpi

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetAcpi(v float32)`

SetAcpi sets Acpi field to given value.

### HasAcpi

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasAcpi() bool`

HasAcpi returns a boolean if a field has been set.

### GetAffinity

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetAffinity() string`

GetAffinity returns the Affinity field if non-nil, zero value otherwise.

### GetAffinityOk

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetAffinityOk() (*string, bool)`

GetAffinityOk returns a tuple with the Affinity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffinity

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetAffinity(v string)`

SetAffinity sets Affinity field to given value.

### HasAffinity

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasAffinity() bool`

HasAffinity returns a boolean if a field has been set.

### GetAgent

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetAgent() string`

GetAgent returns the Agent field if non-nil, zero value otherwise.

### GetAgentOk

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetAgentOk() (*string, bool)`

GetAgentOk returns a tuple with the Agent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgent

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetAgent(v string)`

SetAgent sets Agent field to given value.

### HasAgent

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasAgent() bool`

HasAgent returns a boolean if a field has been set.

### GetArch

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetArch() VirtualMachineArchitecture`

GetArch returns the Arch field if non-nil, zero value otherwise.

### GetArchOk

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetArchOk() (*VirtualMachineArchitecture, bool)`

GetArchOk returns a tuple with the Arch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArch

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetArch(v VirtualMachineArchitecture)`

SetArch sets Arch field to given value.

### HasArch

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasArch() bool`

HasArch returns a boolean if a field has been set.

### GetArgs

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetArgs() string`

GetArgs returns the Args field if non-nil, zero value otherwise.

### GetArgsOk

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetArgsOk() (*string, bool)`

GetArgsOk returns a tuple with the Args field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgs

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetArgs(v string)`

SetArgs sets Args field to given value.

### HasArgs

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasArgs() bool`

HasArgs returns a boolean if a field has been set.

### GetAudio0

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetAudio0() string`

GetAudio0 returns the Audio0 field if non-nil, zero value otherwise.

### GetAudio0Ok

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetAudio0Ok() (*string, bool)`

GetAudio0Ok returns a tuple with the Audio0 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudio0

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetAudio0(v string)`

SetAudio0 sets Audio0 field to given value.

### HasAudio0

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasAudio0() bool`

HasAudio0 returns a boolean if a field has been set.

### GetAutostart

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetAutostart() float32`

GetAutostart returns the Autostart field if non-nil, zero value otherwise.

### GetAutostartOk

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetAutostartOk() (*float32, bool)`

GetAutostartOk returns a tuple with the Autostart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutostart

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetAutostart(v float32)`

SetAutostart sets Autostart field to given value.

### HasAutostart

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasAutostart() bool`

HasAutostart returns a boolean if a field has been set.

### GetBallon

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetBallon() float32`

GetBallon returns the Ballon field if non-nil, zero value otherwise.

### GetBallonOk

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetBallonOk() (*float32, bool)`

GetBallonOk returns a tuple with the Ballon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBallon

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetBallon(v float32)`

SetBallon sets Ballon field to given value.

### HasBallon

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasBallon() bool`

HasBallon returns a boolean if a field has been set.

### GetBoot

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetBoot() string`

GetBoot returns the Boot field if non-nil, zero value otherwise.

### GetBootOk

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetBootOk() (*string, bool)`

GetBootOk returns a tuple with the Boot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoot

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetBoot(v string)`

SetBoot sets Boot field to given value.

### HasBoot

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasBoot() bool`

HasBoot returns a boolean if a field has been set.

### GetBios

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetBios() VirtualMachineBios`

GetBios returns the Bios field if non-nil, zero value otherwise.

### GetBiosOk

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetBiosOk() (*VirtualMachineBios, bool)`

GetBiosOk returns a tuple with the Bios field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBios

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetBios(v VirtualMachineBios)`

SetBios sets Bios field to given value.

### HasBios

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasBios() bool`

HasBios returns a boolean if a field has been set.

### GetBootdisk

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetBootdisk() string`

GetBootdisk returns the Bootdisk field if non-nil, zero value otherwise.

### GetBootdiskOk

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetBootdiskOk() (*string, bool)`

GetBootdiskOk returns a tuple with the Bootdisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootdisk

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetBootdisk(v string)`

SetBootdisk sets Bootdisk field to given value.

### HasBootdisk

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasBootdisk() bool`

HasBootdisk returns a boolean if a field has been set.

### GetCdrom

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetCdrom() string`

GetCdrom returns the Cdrom field if non-nil, zero value otherwise.

### GetCdromOk

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetCdromOk() (*string, bool)`

GetCdromOk returns a tuple with the Cdrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdrom

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetCdrom(v string)`

SetCdrom sets Cdrom field to given value.

### HasCdrom

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasCdrom() bool`

HasCdrom returns a boolean if a field has been set.

### GetCicustom

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetCicustom() string`

GetCicustom returns the Cicustom field if non-nil, zero value otherwise.

### GetCicustomOk

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetCicustomOk() (*string, bool)`

GetCicustomOk returns a tuple with the Cicustom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCicustom

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetCicustom(v string)`

SetCicustom sets Cicustom field to given value.

### HasCicustom

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasCicustom() bool`

HasCicustom returns a boolean if a field has been set.

### GetCipassword

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetCipassword() string`

GetCipassword returns the Cipassword field if non-nil, zero value otherwise.

### GetCipasswordOk

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetCipasswordOk() (*string, bool)`

GetCipasswordOk returns a tuple with the Cipassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCipassword

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetCipassword(v string)`

SetCipassword sets Cipassword field to given value.

### HasCipassword

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasCipassword() bool`

HasCipassword returns a boolean if a field has been set.

### GetCitype

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetCitype() VirtualMachineCloudInitType`

GetCitype returns the Citype field if non-nil, zero value otherwise.

### GetCitypeOk

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetCitypeOk() (*VirtualMachineCloudInitType, bool)`

GetCitypeOk returns a tuple with the Citype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCitype

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetCitype(v VirtualMachineCloudInitType)`

SetCitype sets Citype field to given value.

### HasCitype

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasCitype() bool`

HasCitype returns a boolean if a field has been set.

### GetCiuser

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetCiuser() string`

GetCiuser returns the Ciuser field if non-nil, zero value otherwise.

### GetCiuserOk

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetCiuserOk() (*string, bool)`

GetCiuserOk returns a tuple with the Ciuser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCiuser

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetCiuser(v string)`

SetCiuser sets Ciuser field to given value.

### HasCiuser

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasCiuser() bool`

HasCiuser returns a boolean if a field has been set.

### GetCores

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetCores() float32`

GetCores returns the Cores field if non-nil, zero value otherwise.

### GetCoresOk

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetCoresOk() (*float32, bool)`

GetCoresOk returns a tuple with the Cores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCores

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetCores(v float32)`

SetCores sets Cores field to given value.

### HasCores

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasCores() bool`

HasCores returns a boolean if a field has been set.

### GetCpu

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetCpu() string`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetCpuOk() (*string, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetCpu(v string)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### GetCpulimit

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetCpulimit() float32`

GetCpulimit returns the Cpulimit field if non-nil, zero value otherwise.

### GetCpulimitOk

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetCpulimitOk() (*float32, bool)`

GetCpulimitOk returns a tuple with the Cpulimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpulimit

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetCpulimit(v float32)`

SetCpulimit sets Cpulimit field to given value.

### HasCpulimit

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasCpulimit() bool`

HasCpulimit returns a boolean if a field has been set.

### GetCpuunits

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetCpuunits() float32`

GetCpuunits returns the Cpuunits field if non-nil, zero value otherwise.

### GetCpuunitsOk

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetCpuunitsOk() (*float32, bool)`

GetCpuunitsOk returns a tuple with the Cpuunits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuunits

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetCpuunits(v float32)`

SetCpuunits sets Cpuunits field to given value.

### HasCpuunits

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasCpuunits() bool`

HasCpuunits returns a boolean if a field has been set.

### GetDescription

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDigest

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetDigest() string`

GetDigest returns the Digest field if non-nil, zero value otherwise.

### GetDigestOk

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetDigestOk() (*string, bool)`

GetDigestOk returns a tuple with the Digest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigest

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetDigest(v string)`

SetDigest sets Digest field to given value.

### HasDigest

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasDigest() bool`

HasDigest returns a boolean if a field has been set.

### GetDelete

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetDelete() string`

GetDelete returns the Delete field if non-nil, zero value otherwise.

### GetDeleteOk

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetDeleteOk() (*string, bool)`

GetDeleteOk returns a tuple with the Delete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelete

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetDelete(v string)`

SetDelete sets Delete field to given value.

### HasDelete

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasDelete() bool`

HasDelete returns a boolean if a field has been set.

### GetEfidisk0

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetEfidisk0() string`

GetEfidisk0 returns the Efidisk0 field if non-nil, zero value otherwise.

### GetEfidisk0Ok

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetEfidisk0Ok() (*string, bool)`

GetEfidisk0Ok returns a tuple with the Efidisk0 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEfidisk0

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetEfidisk0(v string)`

SetEfidisk0 sets Efidisk0 field to given value.

### HasEfidisk0

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasEfidisk0() bool`

HasEfidisk0 returns a boolean if a field has been set.

### GetFreeze

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetFreeze() bool`

GetFreeze returns the Freeze field if non-nil, zero value otherwise.

### GetFreezeOk

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetFreezeOk() (*bool, bool)`

GetFreezeOk returns a tuple with the Freeze field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeze

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetFreeze(v bool)`

SetFreeze sets Freeze field to given value.

### HasFreeze

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasFreeze() bool`

HasFreeze returns a boolean if a field has been set.

### GetHookscript

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetHookscript() string`

GetHookscript returns the Hookscript field if non-nil, zero value otherwise.

### GetHookscriptOk

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetHookscriptOk() (*string, bool)`

GetHookscriptOk returns a tuple with the Hookscript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHookscript

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetHookscript(v string)`

SetHookscript sets Hookscript field to given value.

### HasHookscript

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasHookscript() bool`

HasHookscript returns a boolean if a field has been set.

### GetHostpci0

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetHostpci0() string`

GetHostpci0 returns the Hostpci0 field if non-nil, zero value otherwise.

### GetHostpci0Ok

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetHostpci0Ok() (*string, bool)`

GetHostpci0Ok returns a tuple with the Hostpci0 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostpci0

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetHostpci0(v string)`

SetHostpci0 sets Hostpci0 field to given value.

### HasHostpci0

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasHostpci0() bool`

HasHostpci0 returns a boolean if a field has been set.

### GetHostpci1

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetHostpci1() string`

GetHostpci1 returns the Hostpci1 field if non-nil, zero value otherwise.

### GetHostpci1Ok

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetHostpci1Ok() (*string, bool)`

GetHostpci1Ok returns a tuple with the Hostpci1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostpci1

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetHostpci1(v string)`

SetHostpci1 sets Hostpci1 field to given value.

### HasHostpci1

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasHostpci1() bool`

HasHostpci1 returns a boolean if a field has been set.

### GetHostpci2

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetHostpci2() string`

GetHostpci2 returns the Hostpci2 field if non-nil, zero value otherwise.

### GetHostpci2Ok

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetHostpci2Ok() (*string, bool)`

GetHostpci2Ok returns a tuple with the Hostpci2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostpci2

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetHostpci2(v string)`

SetHostpci2 sets Hostpci2 field to given value.

### HasHostpci2

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasHostpci2() bool`

HasHostpci2 returns a boolean if a field has been set.

### GetHostpci3

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetHostpci3() string`

GetHostpci3 returns the Hostpci3 field if non-nil, zero value otherwise.

### GetHostpci3Ok

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetHostpci3Ok() (*string, bool)`

GetHostpci3Ok returns a tuple with the Hostpci3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostpci3

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetHostpci3(v string)`

SetHostpci3 sets Hostpci3 field to given value.

### HasHostpci3

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasHostpci3() bool`

HasHostpci3 returns a boolean if a field has been set.

### GetHostpci4

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetHostpci4() string`

GetHostpci4 returns the Hostpci4 field if non-nil, zero value otherwise.

### GetHostpci4Ok

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetHostpci4Ok() (*string, bool)`

GetHostpci4Ok returns a tuple with the Hostpci4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostpci4

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetHostpci4(v string)`

SetHostpci4 sets Hostpci4 field to given value.

### HasHostpci4

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasHostpci4() bool`

HasHostpci4 returns a boolean if a field has been set.

### GetHostpci5

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetHostpci5() string`

GetHostpci5 returns the Hostpci5 field if non-nil, zero value otherwise.

### GetHostpci5Ok

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetHostpci5Ok() (*string, bool)`

GetHostpci5Ok returns a tuple with the Hostpci5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostpci5

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetHostpci5(v string)`

SetHostpci5 sets Hostpci5 field to given value.

### HasHostpci5

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasHostpci5() bool`

HasHostpci5 returns a boolean if a field has been set.

### GetHostpci6

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetHostpci6() string`

GetHostpci6 returns the Hostpci6 field if non-nil, zero value otherwise.

### GetHostpci6Ok

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetHostpci6Ok() (*string, bool)`

GetHostpci6Ok returns a tuple with the Hostpci6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostpci6

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetHostpci6(v string)`

SetHostpci6 sets Hostpci6 field to given value.

### HasHostpci6

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasHostpci6() bool`

HasHostpci6 returns a boolean if a field has been set.

### GetHostpci7

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetHostpci7() string`

GetHostpci7 returns the Hostpci7 field if non-nil, zero value otherwise.

### GetHostpci7Ok

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetHostpci7Ok() (*string, bool)`

GetHostpci7Ok returns a tuple with the Hostpci7 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostpci7

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetHostpci7(v string)`

SetHostpci7 sets Hostpci7 field to given value.

### HasHostpci7

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasHostpci7() bool`

HasHostpci7 returns a boolean if a field has been set.

### GetHostpci8

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetHostpci8() string`

GetHostpci8 returns the Hostpci8 field if non-nil, zero value otherwise.

### GetHostpci8Ok

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetHostpci8Ok() (*string, bool)`

GetHostpci8Ok returns a tuple with the Hostpci8 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostpci8

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetHostpci8(v string)`

SetHostpci8 sets Hostpci8 field to given value.

### HasHostpci8

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasHostpci8() bool`

HasHostpci8 returns a boolean if a field has been set.

### GetHostpci9

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetHostpci9() string`

GetHostpci9 returns the Hostpci9 field if non-nil, zero value otherwise.

### GetHostpci9Ok

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetHostpci9Ok() (*string, bool)`

GetHostpci9Ok returns a tuple with the Hostpci9 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostpci9

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetHostpci9(v string)`

SetHostpci9 sets Hostpci9 field to given value.

### HasHostpci9

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasHostpci9() bool`

HasHostpci9 returns a boolean if a field has been set.

### GetHugepages

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetHugepages() VirtualMachineHugePages`

GetHugepages returns the Hugepages field if non-nil, zero value otherwise.

### GetHugepagesOk

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetHugepagesOk() (*VirtualMachineHugePages, bool)`

GetHugepagesOk returns a tuple with the Hugepages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHugepages

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetHugepages(v VirtualMachineHugePages)`

SetHugepages sets Hugepages field to given value.

### HasHugepages

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasHugepages() bool`

HasHugepages returns a boolean if a field has been set.

### GetIde0

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetIde0() string`

GetIde0 returns the Ide0 field if non-nil, zero value otherwise.

### GetIde0Ok

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetIde0Ok() (*string, bool)`

GetIde0Ok returns a tuple with the Ide0 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIde0

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetIde0(v string)`

SetIde0 sets Ide0 field to given value.

### HasIde0

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasIde0() bool`

HasIde0 returns a boolean if a field has been set.

### GetIde1

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetIde1() string`

GetIde1 returns the Ide1 field if non-nil, zero value otherwise.

### GetIde1Ok

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetIde1Ok() (*string, bool)`

GetIde1Ok returns a tuple with the Ide1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIde1

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetIde1(v string)`

SetIde1 sets Ide1 field to given value.

### HasIde1

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasIde1() bool`

HasIde1 returns a boolean if a field has been set.

### GetIde2

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetIde2() string`

GetIde2 returns the Ide2 field if non-nil, zero value otherwise.

### GetIde2Ok

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetIde2Ok() (*string, bool)`

GetIde2Ok returns a tuple with the Ide2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIde2

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetIde2(v string)`

SetIde2 sets Ide2 field to given value.

### HasIde2

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasIde2() bool`

HasIde2 returns a boolean if a field has been set.

### GetIde3

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetIde3() string`

GetIde3 returns the Ide3 field if non-nil, zero value otherwise.

### GetIde3Ok

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetIde3Ok() (*string, bool)`

GetIde3Ok returns a tuple with the Ide3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIde3

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetIde3(v string)`

SetIde3 sets Ide3 field to given value.

### HasIde3

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasIde3() bool`

HasIde3 returns a boolean if a field has been set.

### GetIpconfig0

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetIpconfig0() string`

GetIpconfig0 returns the Ipconfig0 field if non-nil, zero value otherwise.

### GetIpconfig0Ok

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetIpconfig0Ok() (*string, bool)`

GetIpconfig0Ok returns a tuple with the Ipconfig0 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpconfig0

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetIpconfig0(v string)`

SetIpconfig0 sets Ipconfig0 field to given value.

### HasIpconfig0

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasIpconfig0() bool`

HasIpconfig0 returns a boolean if a field has been set.

### GetIpconfig1

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetIpconfig1() string`

GetIpconfig1 returns the Ipconfig1 field if non-nil, zero value otherwise.

### GetIpconfig1Ok

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetIpconfig1Ok() (*string, bool)`

GetIpconfig1Ok returns a tuple with the Ipconfig1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpconfig1

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetIpconfig1(v string)`

SetIpconfig1 sets Ipconfig1 field to given value.

### HasIpconfig1

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasIpconfig1() bool`

HasIpconfig1 returns a boolean if a field has been set.

### GetIpconfig2

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetIpconfig2() string`

GetIpconfig2 returns the Ipconfig2 field if non-nil, zero value otherwise.

### GetIpconfig2Ok

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetIpconfig2Ok() (*string, bool)`

GetIpconfig2Ok returns a tuple with the Ipconfig2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpconfig2

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetIpconfig2(v string)`

SetIpconfig2 sets Ipconfig2 field to given value.

### HasIpconfig2

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasIpconfig2() bool`

HasIpconfig2 returns a boolean if a field has been set.

### GetIpconfig3

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetIpconfig3() string`

GetIpconfig3 returns the Ipconfig3 field if non-nil, zero value otherwise.

### GetIpconfig3Ok

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetIpconfig3Ok() (*string, bool)`

GetIpconfig3Ok returns a tuple with the Ipconfig3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpconfig3

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetIpconfig3(v string)`

SetIpconfig3 sets Ipconfig3 field to given value.

### HasIpconfig3

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasIpconfig3() bool`

HasIpconfig3 returns a boolean if a field has been set.

### GetIpconfig4

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetIpconfig4() string`

GetIpconfig4 returns the Ipconfig4 field if non-nil, zero value otherwise.

### GetIpconfig4Ok

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetIpconfig4Ok() (*string, bool)`

GetIpconfig4Ok returns a tuple with the Ipconfig4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpconfig4

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetIpconfig4(v string)`

SetIpconfig4 sets Ipconfig4 field to given value.

### HasIpconfig4

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasIpconfig4() bool`

HasIpconfig4 returns a boolean if a field has been set.

### GetIpconfig5

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetIpconfig5() string`

GetIpconfig5 returns the Ipconfig5 field if non-nil, zero value otherwise.

### GetIpconfig5Ok

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetIpconfig5Ok() (*string, bool)`

GetIpconfig5Ok returns a tuple with the Ipconfig5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpconfig5

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetIpconfig5(v string)`

SetIpconfig5 sets Ipconfig5 field to given value.

### HasIpconfig5

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasIpconfig5() bool`

HasIpconfig5 returns a boolean if a field has been set.

### GetIpconfig6

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetIpconfig6() string`

GetIpconfig6 returns the Ipconfig6 field if non-nil, zero value otherwise.

### GetIpconfig6Ok

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetIpconfig6Ok() (*string, bool)`

GetIpconfig6Ok returns a tuple with the Ipconfig6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpconfig6

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetIpconfig6(v string)`

SetIpconfig6 sets Ipconfig6 field to given value.

### HasIpconfig6

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasIpconfig6() bool`

HasIpconfig6 returns a boolean if a field has been set.

### GetIpconfig7

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetIpconfig7() string`

GetIpconfig7 returns the Ipconfig7 field if non-nil, zero value otherwise.

### GetIpconfig7Ok

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetIpconfig7Ok() (*string, bool)`

GetIpconfig7Ok returns a tuple with the Ipconfig7 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpconfig7

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetIpconfig7(v string)`

SetIpconfig7 sets Ipconfig7 field to given value.

### HasIpconfig7

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasIpconfig7() bool`

HasIpconfig7 returns a boolean if a field has been set.

### GetIvshmem

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetIvshmem() string`

GetIvshmem returns the Ivshmem field if non-nil, zero value otherwise.

### GetIvshmemOk

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetIvshmemOk() (*string, bool)`

GetIvshmemOk returns a tuple with the Ivshmem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIvshmem

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetIvshmem(v string)`

SetIvshmem sets Ivshmem field to given value.

### HasIvshmem

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasIvshmem() bool`

HasIvshmem returns a boolean if a field has been set.

### GetKeephugepages

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetKeephugepages() float32`

GetKeephugepages returns the Keephugepages field if non-nil, zero value otherwise.

### GetKeephugepagesOk

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetKeephugepagesOk() (*float32, bool)`

GetKeephugepagesOk returns a tuple with the Keephugepages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeephugepages

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetKeephugepages(v float32)`

SetKeephugepages sets Keephugepages field to given value.

### HasKeephugepages

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasKeephugepages() bool`

HasKeephugepages returns a boolean if a field has been set.

### GetKeyboard

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetKeyboard() VirtualMachineKeyboard`

GetKeyboard returns the Keyboard field if non-nil, zero value otherwise.

### GetKeyboardOk

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetKeyboardOk() (*VirtualMachineKeyboard, bool)`

GetKeyboardOk returns a tuple with the Keyboard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyboard

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetKeyboard(v VirtualMachineKeyboard)`

SetKeyboard sets Keyboard field to given value.

### HasKeyboard

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasKeyboard() bool`

HasKeyboard returns a boolean if a field has been set.

### GetKvm

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetKvm() float32`

GetKvm returns the Kvm field if non-nil, zero value otherwise.

### GetKvmOk

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetKvmOk() (*float32, bool)`

GetKvmOk returns a tuple with the Kvm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKvm

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetKvm(v float32)`

SetKvm sets Kvm field to given value.

### HasKvm

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasKvm() bool`

HasKvm returns a boolean if a field has been set.

### GetLocaltime

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetLocaltime() float32`

GetLocaltime returns the Localtime field if non-nil, zero value otherwise.

### GetLocaltimeOk

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetLocaltimeOk() (*float32, bool)`

GetLocaltimeOk returns a tuple with the Localtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocaltime

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetLocaltime(v float32)`

SetLocaltime sets Localtime field to given value.

### HasLocaltime

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasLocaltime() bool`

HasLocaltime returns a boolean if a field has been set.

### GetLock

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetLock() VirtualMachineConfigLock`

GetLock returns the Lock field if non-nil, zero value otherwise.

### GetLockOk

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetLockOk() (*VirtualMachineConfigLock, bool)`

GetLockOk returns a tuple with the Lock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLock

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetLock(v VirtualMachineConfigLock)`

SetLock sets Lock field to given value.

### HasLock

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasLock() bool`

HasLock returns a boolean if a field has been set.

### GetMachine

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetMachine() string`

GetMachine returns the Machine field if non-nil, zero value otherwise.

### GetMachineOk

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetMachineOk() (*string, bool)`

GetMachineOk returns a tuple with the Machine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachine

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetMachine(v string)`

SetMachine sets Machine field to given value.

### HasMachine

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasMachine() bool`

HasMachine returns a boolean if a field has been set.

### GetMemory

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetMemory() float32`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetMemoryOk() (*float32, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetMemory(v float32)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### GetMigrateDowntime

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetMigrateDowntime() float32`

GetMigrateDowntime returns the MigrateDowntime field if non-nil, zero value otherwise.

### GetMigrateDowntimeOk

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetMigrateDowntimeOk() (*float32, bool)`

GetMigrateDowntimeOk returns a tuple with the MigrateDowntime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigrateDowntime

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetMigrateDowntime(v float32)`

SetMigrateDowntime sets MigrateDowntime field to given value.

### HasMigrateDowntime

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasMigrateDowntime() bool`

HasMigrateDowntime returns a boolean if a field has been set.

### GetMigrateSpeed

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetMigrateSpeed() float32`

GetMigrateSpeed returns the MigrateSpeed field if non-nil, zero value otherwise.

### GetMigrateSpeedOk

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetMigrateSpeedOk() (*float32, bool)`

GetMigrateSpeedOk returns a tuple with the MigrateSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigrateSpeed

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetMigrateSpeed(v float32)`

SetMigrateSpeed sets MigrateSpeed field to given value.

### HasMigrateSpeed

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasMigrateSpeed() bool`

HasMigrateSpeed returns a boolean if a field has been set.

### GetName

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNameserver

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetNameserver() string`

GetNameserver returns the Nameserver field if non-nil, zero value otherwise.

### GetNameserverOk

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetNameserverOk() (*string, bool)`

GetNameserverOk returns a tuple with the Nameserver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameserver

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetNameserver(v string)`

SetNameserver sets Nameserver field to given value.

### HasNameserver

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasNameserver() bool`

HasNameserver returns a boolean if a field has been set.

### GetNet0

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetNet0() string`

GetNet0 returns the Net0 field if non-nil, zero value otherwise.

### GetNet0Ok

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetNet0Ok() (*string, bool)`

GetNet0Ok returns a tuple with the Net0 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet0

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetNet0(v string)`

SetNet0 sets Net0 field to given value.

### HasNet0

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasNet0() bool`

HasNet0 returns a boolean if a field has been set.

### GetNet1

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetNet1() string`

GetNet1 returns the Net1 field if non-nil, zero value otherwise.

### GetNet1Ok

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetNet1Ok() (*string, bool)`

GetNet1Ok returns a tuple with the Net1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet1

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetNet1(v string)`

SetNet1 sets Net1 field to given value.

### HasNet1

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasNet1() bool`

HasNet1 returns a boolean if a field has been set.

### GetNet2

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetNet2() string`

GetNet2 returns the Net2 field if non-nil, zero value otherwise.

### GetNet2Ok

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetNet2Ok() (*string, bool)`

GetNet2Ok returns a tuple with the Net2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet2

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetNet2(v string)`

SetNet2 sets Net2 field to given value.

### HasNet2

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasNet2() bool`

HasNet2 returns a boolean if a field has been set.

### GetNet3

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetNet3() string`

GetNet3 returns the Net3 field if non-nil, zero value otherwise.

### GetNet3Ok

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetNet3Ok() (*string, bool)`

GetNet3Ok returns a tuple with the Net3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet3

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetNet3(v string)`

SetNet3 sets Net3 field to given value.

### HasNet3

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasNet3() bool`

HasNet3 returns a boolean if a field has been set.

### GetNet4

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetNet4() string`

GetNet4 returns the Net4 field if non-nil, zero value otherwise.

### GetNet4Ok

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetNet4Ok() (*string, bool)`

GetNet4Ok returns a tuple with the Net4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet4

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetNet4(v string)`

SetNet4 sets Net4 field to given value.

### HasNet4

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasNet4() bool`

HasNet4 returns a boolean if a field has been set.

### GetNet5

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetNet5() string`

GetNet5 returns the Net5 field if non-nil, zero value otherwise.

### GetNet5Ok

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetNet5Ok() (*string, bool)`

GetNet5Ok returns a tuple with the Net5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet5

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetNet5(v string)`

SetNet5 sets Net5 field to given value.

### HasNet5

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasNet5() bool`

HasNet5 returns a boolean if a field has been set.

### GetNet6

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetNet6() string`

GetNet6 returns the Net6 field if non-nil, zero value otherwise.

### GetNet6Ok

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetNet6Ok() (*string, bool)`

GetNet6Ok returns a tuple with the Net6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet6

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetNet6(v string)`

SetNet6 sets Net6 field to given value.

### HasNet6

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasNet6() bool`

HasNet6 returns a boolean if a field has been set.

### GetNet7

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetNet7() string`

GetNet7 returns the Net7 field if non-nil, zero value otherwise.

### GetNet7Ok

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetNet7Ok() (*string, bool)`

GetNet7Ok returns a tuple with the Net7 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet7

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetNet7(v string)`

SetNet7 sets Net7 field to given value.

### HasNet7

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasNet7() bool`

HasNet7 returns a boolean if a field has been set.

### GetNuma

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetNuma() float32`

GetNuma returns the Numa field if non-nil, zero value otherwise.

### GetNumaOk

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetNumaOk() (*float32, bool)`

GetNumaOk returns a tuple with the Numa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNuma

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetNuma(v float32)`

SetNuma sets Numa field to given value.

### HasNuma

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasNuma() bool`

HasNuma returns a boolean if a field has been set.

### GetNuma0

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetNuma0() string`

GetNuma0 returns the Numa0 field if non-nil, zero value otherwise.

### GetNuma0Ok

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetNuma0Ok() (*string, bool)`

GetNuma0Ok returns a tuple with the Numa0 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNuma0

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetNuma0(v string)`

SetNuma0 sets Numa0 field to given value.

### HasNuma0

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasNuma0() bool`

HasNuma0 returns a boolean if a field has been set.

### GetNuma1

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetNuma1() string`

GetNuma1 returns the Numa1 field if non-nil, zero value otherwise.

### GetNuma1Ok

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetNuma1Ok() (*string, bool)`

GetNuma1Ok returns a tuple with the Numa1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNuma1

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetNuma1(v string)`

SetNuma1 sets Numa1 field to given value.

### HasNuma1

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasNuma1() bool`

HasNuma1 returns a boolean if a field has been set.

### GetNuma2

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetNuma2() string`

GetNuma2 returns the Numa2 field if non-nil, zero value otherwise.

### GetNuma2Ok

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetNuma2Ok() (*string, bool)`

GetNuma2Ok returns a tuple with the Numa2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNuma2

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetNuma2(v string)`

SetNuma2 sets Numa2 field to given value.

### HasNuma2

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasNuma2() bool`

HasNuma2 returns a boolean if a field has been set.

### GetNuma3

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetNuma3() string`

GetNuma3 returns the Numa3 field if non-nil, zero value otherwise.

### GetNuma3Ok

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetNuma3Ok() (*string, bool)`

GetNuma3Ok returns a tuple with the Numa3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNuma3

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetNuma3(v string)`

SetNuma3 sets Numa3 field to given value.

### HasNuma3

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasNuma3() bool`

HasNuma3 returns a boolean if a field has been set.

### GetNuma4

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetNuma4() string`

GetNuma4 returns the Numa4 field if non-nil, zero value otherwise.

### GetNuma4Ok

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetNuma4Ok() (*string, bool)`

GetNuma4Ok returns a tuple with the Numa4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNuma4

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetNuma4(v string)`

SetNuma4 sets Numa4 field to given value.

### HasNuma4

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasNuma4() bool`

HasNuma4 returns a boolean if a field has been set.

### GetNuma5

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetNuma5() string`

GetNuma5 returns the Numa5 field if non-nil, zero value otherwise.

### GetNuma5Ok

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetNuma5Ok() (*string, bool)`

GetNuma5Ok returns a tuple with the Numa5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNuma5

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetNuma5(v string)`

SetNuma5 sets Numa5 field to given value.

### HasNuma5

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasNuma5() bool`

HasNuma5 returns a boolean if a field has been set.

### GetNuma6

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetNuma6() string`

GetNuma6 returns the Numa6 field if non-nil, zero value otherwise.

### GetNuma6Ok

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetNuma6Ok() (*string, bool)`

GetNuma6Ok returns a tuple with the Numa6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNuma6

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetNuma6(v string)`

SetNuma6 sets Numa6 field to given value.

### HasNuma6

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasNuma6() bool`

HasNuma6 returns a boolean if a field has been set.

### GetNuma7

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetNuma7() string`

GetNuma7 returns the Numa7 field if non-nil, zero value otherwise.

### GetNuma7Ok

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetNuma7Ok() (*string, bool)`

GetNuma7Ok returns a tuple with the Numa7 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNuma7

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetNuma7(v string)`

SetNuma7 sets Numa7 field to given value.

### HasNuma7

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasNuma7() bool`

HasNuma7 returns a boolean if a field has been set.

### GetOnboot

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetOnboot() float32`

GetOnboot returns the Onboot field if non-nil, zero value otherwise.

### GetOnbootOk

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetOnbootOk() (*float32, bool)`

GetOnbootOk returns a tuple with the Onboot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnboot

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetOnboot(v float32)`

SetOnboot sets Onboot field to given value.

### HasOnboot

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasOnboot() bool`

HasOnboot returns a boolean if a field has been set.

### GetOstype

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetOstype() VirtualMachineOperatingSystem`

GetOstype returns the Ostype field if non-nil, zero value otherwise.

### GetOstypeOk

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetOstypeOk() (*VirtualMachineOperatingSystem, bool)`

GetOstypeOk returns a tuple with the Ostype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOstype

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetOstype(v VirtualMachineOperatingSystem)`

SetOstype sets Ostype field to given value.

### HasOstype

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasOstype() bool`

HasOstype returns a boolean if a field has been set.

### GetParallel0

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetParallel0() string`

GetParallel0 returns the Parallel0 field if non-nil, zero value otherwise.

### GetParallel0Ok

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetParallel0Ok() (*string, bool)`

GetParallel0Ok returns a tuple with the Parallel0 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParallel0

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetParallel0(v string)`

SetParallel0 sets Parallel0 field to given value.

### HasParallel0

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasParallel0() bool`

HasParallel0 returns a boolean if a field has been set.

### GetParallel1

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetParallel1() string`

GetParallel1 returns the Parallel1 field if non-nil, zero value otherwise.

### GetParallel1Ok

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetParallel1Ok() (*string, bool)`

GetParallel1Ok returns a tuple with the Parallel1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParallel1

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetParallel1(v string)`

SetParallel1 sets Parallel1 field to given value.

### HasParallel1

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasParallel1() bool`

HasParallel1 returns a boolean if a field has been set.

### GetParallel2

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetParallel2() string`

GetParallel2 returns the Parallel2 field if non-nil, zero value otherwise.

### GetParallel2Ok

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetParallel2Ok() (*string, bool)`

GetParallel2Ok returns a tuple with the Parallel2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParallel2

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetParallel2(v string)`

SetParallel2 sets Parallel2 field to given value.

### HasParallel2

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasParallel2() bool`

HasParallel2 returns a boolean if a field has been set.

### GetProtection

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetProtection() float32`

GetProtection returns the Protection field if non-nil, zero value otherwise.

### GetProtectionOk

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetProtectionOk() (*float32, bool)`

GetProtectionOk returns a tuple with the Protection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtection

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetProtection(v float32)`

SetProtection sets Protection field to given value.

### HasProtection

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasProtection() bool`

HasProtection returns a boolean if a field has been set.

### GetReboot

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetReboot() float32`

GetReboot returns the Reboot field if non-nil, zero value otherwise.

### GetRebootOk

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetRebootOk() (*float32, bool)`

GetRebootOk returns a tuple with the Reboot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReboot

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetReboot(v float32)`

SetReboot sets Reboot field to given value.

### HasReboot

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasReboot() bool`

HasReboot returns a boolean if a field has been set.

### GetRng0

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetRng0() string`

GetRng0 returns the Rng0 field if non-nil, zero value otherwise.

### GetRng0Ok

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetRng0Ok() (*string, bool)`

GetRng0Ok returns a tuple with the Rng0 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRng0

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetRng0(v string)`

SetRng0 sets Rng0 field to given value.

### HasRng0

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasRng0() bool`

HasRng0 returns a boolean if a field has been set.

### GetSata0

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetSata0() string`

GetSata0 returns the Sata0 field if non-nil, zero value otherwise.

### GetSata0Ok

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetSata0Ok() (*string, bool)`

GetSata0Ok returns a tuple with the Sata0 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSata0

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetSata0(v string)`

SetSata0 sets Sata0 field to given value.

### HasSata0

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasSata0() bool`

HasSata0 returns a boolean if a field has been set.

### GetSata1

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetSata1() string`

GetSata1 returns the Sata1 field if non-nil, zero value otherwise.

### GetSata1Ok

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetSata1Ok() (*string, bool)`

GetSata1Ok returns a tuple with the Sata1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSata1

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetSata1(v string)`

SetSata1 sets Sata1 field to given value.

### HasSata1

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasSata1() bool`

HasSata1 returns a boolean if a field has been set.

### GetSata2

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetSata2() string`

GetSata2 returns the Sata2 field if non-nil, zero value otherwise.

### GetSata2Ok

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetSata2Ok() (*string, bool)`

GetSata2Ok returns a tuple with the Sata2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSata2

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetSata2(v string)`

SetSata2 sets Sata2 field to given value.

### HasSata2

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasSata2() bool`

HasSata2 returns a boolean if a field has been set.

### GetSata3

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetSata3() string`

GetSata3 returns the Sata3 field if non-nil, zero value otherwise.

### GetSata3Ok

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetSata3Ok() (*string, bool)`

GetSata3Ok returns a tuple with the Sata3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSata3

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetSata3(v string)`

SetSata3 sets Sata3 field to given value.

### HasSata3

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasSata3() bool`

HasSata3 returns a boolean if a field has been set.

### GetSata4

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetSata4() string`

GetSata4 returns the Sata4 field if non-nil, zero value otherwise.

### GetSata4Ok

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetSata4Ok() (*string, bool)`

GetSata4Ok returns a tuple with the Sata4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSata4

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetSata4(v string)`

SetSata4 sets Sata4 field to given value.

### HasSata4

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasSata4() bool`

HasSata4 returns a boolean if a field has been set.

### GetSata5

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetSata5() string`

GetSata5 returns the Sata5 field if non-nil, zero value otherwise.

### GetSata5Ok

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetSata5Ok() (*string, bool)`

GetSata5Ok returns a tuple with the Sata5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSata5

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetSata5(v string)`

SetSata5 sets Sata5 field to given value.

### HasSata5

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasSata5() bool`

HasSata5 returns a boolean if a field has been set.

### GetScsi0

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetScsi0() string`

GetScsi0 returns the Scsi0 field if non-nil, zero value otherwise.

### GetScsi0Ok

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetScsi0Ok() (*string, bool)`

GetScsi0Ok returns a tuple with the Scsi0 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi0

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetScsi0(v string)`

SetScsi0 sets Scsi0 field to given value.

### HasScsi0

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasScsi0() bool`

HasScsi0 returns a boolean if a field has been set.

### GetScsi1

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetScsi1() string`

GetScsi1 returns the Scsi1 field if non-nil, zero value otherwise.

### GetScsi1Ok

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetScsi1Ok() (*string, bool)`

GetScsi1Ok returns a tuple with the Scsi1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi1

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetScsi1(v string)`

SetScsi1 sets Scsi1 field to given value.

### HasScsi1

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasScsi1() bool`

HasScsi1 returns a boolean if a field has been set.

### GetScsi2

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetScsi2() string`

GetScsi2 returns the Scsi2 field if non-nil, zero value otherwise.

### GetScsi2Ok

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetScsi2Ok() (*string, bool)`

GetScsi2Ok returns a tuple with the Scsi2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi2

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetScsi2(v string)`

SetScsi2 sets Scsi2 field to given value.

### HasScsi2

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasScsi2() bool`

HasScsi2 returns a boolean if a field has been set.

### GetScsi3

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetScsi3() string`

GetScsi3 returns the Scsi3 field if non-nil, zero value otherwise.

### GetScsi3Ok

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetScsi3Ok() (*string, bool)`

GetScsi3Ok returns a tuple with the Scsi3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi3

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetScsi3(v string)`

SetScsi3 sets Scsi3 field to given value.

### HasScsi3

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasScsi3() bool`

HasScsi3 returns a boolean if a field has been set.

### GetScsi4

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetScsi4() string`

GetScsi4 returns the Scsi4 field if non-nil, zero value otherwise.

### GetScsi4Ok

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetScsi4Ok() (*string, bool)`

GetScsi4Ok returns a tuple with the Scsi4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi4

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetScsi4(v string)`

SetScsi4 sets Scsi4 field to given value.

### HasScsi4

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasScsi4() bool`

HasScsi4 returns a boolean if a field has been set.

### GetScsi5

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetScsi5() string`

GetScsi5 returns the Scsi5 field if non-nil, zero value otherwise.

### GetScsi5Ok

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetScsi5Ok() (*string, bool)`

GetScsi5Ok returns a tuple with the Scsi5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi5

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetScsi5(v string)`

SetScsi5 sets Scsi5 field to given value.

### HasScsi5

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasScsi5() bool`

HasScsi5 returns a boolean if a field has been set.

### GetScsi6

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetScsi6() string`

GetScsi6 returns the Scsi6 field if non-nil, zero value otherwise.

### GetScsi6Ok

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetScsi6Ok() (*string, bool)`

GetScsi6Ok returns a tuple with the Scsi6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi6

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetScsi6(v string)`

SetScsi6 sets Scsi6 field to given value.

### HasScsi6

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasScsi6() bool`

HasScsi6 returns a boolean if a field has been set.

### GetScsi7

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetScsi7() string`

GetScsi7 returns the Scsi7 field if non-nil, zero value otherwise.

### GetScsi7Ok

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetScsi7Ok() (*string, bool)`

GetScsi7Ok returns a tuple with the Scsi7 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi7

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetScsi7(v string)`

SetScsi7 sets Scsi7 field to given value.

### HasScsi7

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasScsi7() bool`

HasScsi7 returns a boolean if a field has been set.

### GetScsi8

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetScsi8() string`

GetScsi8 returns the Scsi8 field if non-nil, zero value otherwise.

### GetScsi8Ok

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetScsi8Ok() (*string, bool)`

GetScsi8Ok returns a tuple with the Scsi8 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi8

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetScsi8(v string)`

SetScsi8 sets Scsi8 field to given value.

### HasScsi8

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasScsi8() bool`

HasScsi8 returns a boolean if a field has been set.

### GetScsi9

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetScsi9() string`

GetScsi9 returns the Scsi9 field if non-nil, zero value otherwise.

### GetScsi9Ok

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetScsi9Ok() (*string, bool)`

GetScsi9Ok returns a tuple with the Scsi9 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi9

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetScsi9(v string)`

SetScsi9 sets Scsi9 field to given value.

### HasScsi9

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasScsi9() bool`

HasScsi9 returns a boolean if a field has been set.

### GetScsi10

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetScsi10() string`

GetScsi10 returns the Scsi10 field if non-nil, zero value otherwise.

### GetScsi10Ok

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetScsi10Ok() (*string, bool)`

GetScsi10Ok returns a tuple with the Scsi10 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi10

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetScsi10(v string)`

SetScsi10 sets Scsi10 field to given value.

### HasScsi10

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasScsi10() bool`

HasScsi10 returns a boolean if a field has been set.

### GetScsi11

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetScsi11() string`

GetScsi11 returns the Scsi11 field if non-nil, zero value otherwise.

### GetScsi11Ok

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetScsi11Ok() (*string, bool)`

GetScsi11Ok returns a tuple with the Scsi11 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi11

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetScsi11(v string)`

SetScsi11 sets Scsi11 field to given value.

### HasScsi11

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasScsi11() bool`

HasScsi11 returns a boolean if a field has been set.

### GetScsi12

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetScsi12() string`

GetScsi12 returns the Scsi12 field if non-nil, zero value otherwise.

### GetScsi12Ok

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetScsi12Ok() (*string, bool)`

GetScsi12Ok returns a tuple with the Scsi12 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi12

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetScsi12(v string)`

SetScsi12 sets Scsi12 field to given value.

### HasScsi12

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasScsi12() bool`

HasScsi12 returns a boolean if a field has been set.

### GetScsi13

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetScsi13() string`

GetScsi13 returns the Scsi13 field if non-nil, zero value otherwise.

### GetScsi13Ok

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetScsi13Ok() (*string, bool)`

GetScsi13Ok returns a tuple with the Scsi13 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi13

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetScsi13(v string)`

SetScsi13 sets Scsi13 field to given value.

### HasScsi13

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasScsi13() bool`

HasScsi13 returns a boolean if a field has been set.

### GetScsi14

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetScsi14() string`

GetScsi14 returns the Scsi14 field if non-nil, zero value otherwise.

### GetScsi14Ok

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetScsi14Ok() (*string, bool)`

GetScsi14Ok returns a tuple with the Scsi14 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi14

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetScsi14(v string)`

SetScsi14 sets Scsi14 field to given value.

### HasScsi14

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasScsi14() bool`

HasScsi14 returns a boolean if a field has been set.

### GetScsi15

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetScsi15() string`

GetScsi15 returns the Scsi15 field if non-nil, zero value otherwise.

### GetScsi15Ok

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetScsi15Ok() (*string, bool)`

GetScsi15Ok returns a tuple with the Scsi15 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi15

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetScsi15(v string)`

SetScsi15 sets Scsi15 field to given value.

### HasScsi15

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasScsi15() bool`

HasScsi15 returns a boolean if a field has been set.

### GetScsi16

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetScsi16() string`

GetScsi16 returns the Scsi16 field if non-nil, zero value otherwise.

### GetScsi16Ok

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetScsi16Ok() (*string, bool)`

GetScsi16Ok returns a tuple with the Scsi16 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi16

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetScsi16(v string)`

SetScsi16 sets Scsi16 field to given value.

### HasScsi16

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasScsi16() bool`

HasScsi16 returns a boolean if a field has been set.

### GetScsi17

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetScsi17() string`

GetScsi17 returns the Scsi17 field if non-nil, zero value otherwise.

### GetScsi17Ok

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetScsi17Ok() (*string, bool)`

GetScsi17Ok returns a tuple with the Scsi17 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi17

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetScsi17(v string)`

SetScsi17 sets Scsi17 field to given value.

### HasScsi17

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasScsi17() bool`

HasScsi17 returns a boolean if a field has been set.

### GetScsi18

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetScsi18() string`

GetScsi18 returns the Scsi18 field if non-nil, zero value otherwise.

### GetScsi18Ok

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetScsi18Ok() (*string, bool)`

GetScsi18Ok returns a tuple with the Scsi18 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi18

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetScsi18(v string)`

SetScsi18 sets Scsi18 field to given value.

### HasScsi18

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasScsi18() bool`

HasScsi18 returns a boolean if a field has been set.

### GetScsi19

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetScsi19() string`

GetScsi19 returns the Scsi19 field if non-nil, zero value otherwise.

### GetScsi19Ok

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetScsi19Ok() (*string, bool)`

GetScsi19Ok returns a tuple with the Scsi19 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi19

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetScsi19(v string)`

SetScsi19 sets Scsi19 field to given value.

### HasScsi19

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasScsi19() bool`

HasScsi19 returns a boolean if a field has been set.

### GetScsi20

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetScsi20() string`

GetScsi20 returns the Scsi20 field if non-nil, zero value otherwise.

### GetScsi20Ok

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetScsi20Ok() (*string, bool)`

GetScsi20Ok returns a tuple with the Scsi20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi20

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetScsi20(v string)`

SetScsi20 sets Scsi20 field to given value.

### HasScsi20

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasScsi20() bool`

HasScsi20 returns a boolean if a field has been set.

### GetScsi21

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetScsi21() string`

GetScsi21 returns the Scsi21 field if non-nil, zero value otherwise.

### GetScsi21Ok

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetScsi21Ok() (*string, bool)`

GetScsi21Ok returns a tuple with the Scsi21 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi21

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetScsi21(v string)`

SetScsi21 sets Scsi21 field to given value.

### HasScsi21

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasScsi21() bool`

HasScsi21 returns a boolean if a field has been set.

### GetScsi22

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetScsi22() string`

GetScsi22 returns the Scsi22 field if non-nil, zero value otherwise.

### GetScsi22Ok

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetScsi22Ok() (*string, bool)`

GetScsi22Ok returns a tuple with the Scsi22 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi22

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetScsi22(v string)`

SetScsi22 sets Scsi22 field to given value.

### HasScsi22

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasScsi22() bool`

HasScsi22 returns a boolean if a field has been set.

### GetScsi23

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetScsi23() string`

GetScsi23 returns the Scsi23 field if non-nil, zero value otherwise.

### GetScsi23Ok

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetScsi23Ok() (*string, bool)`

GetScsi23Ok returns a tuple with the Scsi23 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi23

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetScsi23(v string)`

SetScsi23 sets Scsi23 field to given value.

### HasScsi23

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasScsi23() bool`

HasScsi23 returns a boolean if a field has been set.

### GetScsi24

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetScsi24() string`

GetScsi24 returns the Scsi24 field if non-nil, zero value otherwise.

### GetScsi24Ok

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetScsi24Ok() (*string, bool)`

GetScsi24Ok returns a tuple with the Scsi24 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi24

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetScsi24(v string)`

SetScsi24 sets Scsi24 field to given value.

### HasScsi24

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasScsi24() bool`

HasScsi24 returns a boolean if a field has been set.

### GetScsi25

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetScsi25() string`

GetScsi25 returns the Scsi25 field if non-nil, zero value otherwise.

### GetScsi25Ok

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetScsi25Ok() (*string, bool)`

GetScsi25Ok returns a tuple with the Scsi25 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi25

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetScsi25(v string)`

SetScsi25 sets Scsi25 field to given value.

### HasScsi25

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasScsi25() bool`

HasScsi25 returns a boolean if a field has been set.

### GetScsi26

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetScsi26() string`

GetScsi26 returns the Scsi26 field if non-nil, zero value otherwise.

### GetScsi26Ok

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetScsi26Ok() (*string, bool)`

GetScsi26Ok returns a tuple with the Scsi26 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi26

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetScsi26(v string)`

SetScsi26 sets Scsi26 field to given value.

### HasScsi26

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasScsi26() bool`

HasScsi26 returns a boolean if a field has been set.

### GetScsi27

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetScsi27() string`

GetScsi27 returns the Scsi27 field if non-nil, zero value otherwise.

### GetScsi27Ok

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetScsi27Ok() (*string, bool)`

GetScsi27Ok returns a tuple with the Scsi27 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi27

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetScsi27(v string)`

SetScsi27 sets Scsi27 field to given value.

### HasScsi27

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasScsi27() bool`

HasScsi27 returns a boolean if a field has been set.

### GetScsi28

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetScsi28() string`

GetScsi28 returns the Scsi28 field if non-nil, zero value otherwise.

### GetScsi28Ok

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetScsi28Ok() (*string, bool)`

GetScsi28Ok returns a tuple with the Scsi28 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi28

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetScsi28(v string)`

SetScsi28 sets Scsi28 field to given value.

### HasScsi28

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasScsi28() bool`

HasScsi28 returns a boolean if a field has been set.

### GetScsi29

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetScsi29() string`

GetScsi29 returns the Scsi29 field if non-nil, zero value otherwise.

### GetScsi29Ok

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetScsi29Ok() (*string, bool)`

GetScsi29Ok returns a tuple with the Scsi29 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi29

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetScsi29(v string)`

SetScsi29 sets Scsi29 field to given value.

### HasScsi29

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasScsi29() bool`

HasScsi29 returns a boolean if a field has been set.

### GetScsi30

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetScsi30() string`

GetScsi30 returns the Scsi30 field if non-nil, zero value otherwise.

### GetScsi30Ok

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetScsi30Ok() (*string, bool)`

GetScsi30Ok returns a tuple with the Scsi30 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi30

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetScsi30(v string)`

SetScsi30 sets Scsi30 field to given value.

### HasScsi30

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasScsi30() bool`

HasScsi30 returns a boolean if a field has been set.

### GetScsihw

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetScsihw() VirtualMachineScsiControllerType`

GetScsihw returns the Scsihw field if non-nil, zero value otherwise.

### GetScsihwOk

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetScsihwOk() (*VirtualMachineScsiControllerType, bool)`

GetScsihwOk returns a tuple with the Scsihw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsihw

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetScsihw(v VirtualMachineScsiControllerType)`

SetScsihw sets Scsihw field to given value.

### HasScsihw

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasScsihw() bool`

HasScsihw returns a boolean if a field has been set.

### GetSearchdomain

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetSearchdomain() string`

GetSearchdomain returns the Searchdomain field if non-nil, zero value otherwise.

### GetSearchdomainOk

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetSearchdomainOk() (*string, bool)`

GetSearchdomainOk returns a tuple with the Searchdomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchdomain

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetSearchdomain(v string)`

SetSearchdomain sets Searchdomain field to given value.

### HasSearchdomain

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasSearchdomain() bool`

HasSearchdomain returns a boolean if a field has been set.

### GetSerial0

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetSerial0() string`

GetSerial0 returns the Serial0 field if non-nil, zero value otherwise.

### GetSerial0Ok

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetSerial0Ok() (*string, bool)`

GetSerial0Ok returns a tuple with the Serial0 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial0

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetSerial0(v string)`

SetSerial0 sets Serial0 field to given value.

### HasSerial0

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasSerial0() bool`

HasSerial0 returns a boolean if a field has been set.

### GetSerial1

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetSerial1() string`

GetSerial1 returns the Serial1 field if non-nil, zero value otherwise.

### GetSerial1Ok

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetSerial1Ok() (*string, bool)`

GetSerial1Ok returns a tuple with the Serial1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial1

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetSerial1(v string)`

SetSerial1 sets Serial1 field to given value.

### HasSerial1

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasSerial1() bool`

HasSerial1 returns a boolean if a field has been set.

### GetSerial2

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetSerial2() string`

GetSerial2 returns the Serial2 field if non-nil, zero value otherwise.

### GetSerial2Ok

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetSerial2Ok() (*string, bool)`

GetSerial2Ok returns a tuple with the Serial2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial2

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetSerial2(v string)`

SetSerial2 sets Serial2 field to given value.

### HasSerial2

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasSerial2() bool`

HasSerial2 returns a boolean if a field has been set.

### GetSerial3

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetSerial3() string`

GetSerial3 returns the Serial3 field if non-nil, zero value otherwise.

### GetSerial3Ok

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetSerial3Ok() (*string, bool)`

GetSerial3Ok returns a tuple with the Serial3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial3

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetSerial3(v string)`

SetSerial3 sets Serial3 field to given value.

### HasSerial3

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasSerial3() bool`

HasSerial3 returns a boolean if a field has been set.

### GetShares

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetShares() float32`

GetShares returns the Shares field if non-nil, zero value otherwise.

### GetSharesOk

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetSharesOk() (*float32, bool)`

GetSharesOk returns a tuple with the Shares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShares

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetShares(v float32)`

SetShares sets Shares field to given value.

### HasShares

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasShares() bool`

HasShares returns a boolean if a field has been set.

### GetSmbios1

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetSmbios1() string`

GetSmbios1 returns the Smbios1 field if non-nil, zero value otherwise.

### GetSmbios1Ok

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetSmbios1Ok() (*string, bool)`

GetSmbios1Ok returns a tuple with the Smbios1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmbios1

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetSmbios1(v string)`

SetSmbios1 sets Smbios1 field to given value.

### HasSmbios1

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasSmbios1() bool`

HasSmbios1 returns a boolean if a field has been set.

### GetSockets

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetSockets() float32`

GetSockets returns the Sockets field if non-nil, zero value otherwise.

### GetSocketsOk

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetSocketsOk() (*float32, bool)`

GetSocketsOk returns a tuple with the Sockets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSockets

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetSockets(v float32)`

SetSockets sets Sockets field to given value.

### HasSockets

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasSockets() bool`

HasSockets returns a boolean if a field has been set.

### GetSpiceEnhancements

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetSpiceEnhancements() string`

GetSpiceEnhancements returns the SpiceEnhancements field if non-nil, zero value otherwise.

### GetSpiceEnhancementsOk

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetSpiceEnhancementsOk() (*string, bool)`

GetSpiceEnhancementsOk returns a tuple with the SpiceEnhancements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpiceEnhancements

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetSpiceEnhancements(v string)`

SetSpiceEnhancements sets SpiceEnhancements field to given value.

### HasSpiceEnhancements

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasSpiceEnhancements() bool`

HasSpiceEnhancements returns a boolean if a field has been set.

### GetSshkeys

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetSshkeys() string`

GetSshkeys returns the Sshkeys field if non-nil, zero value otherwise.

### GetSshkeysOk

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetSshkeysOk() (*string, bool)`

GetSshkeysOk returns a tuple with the Sshkeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshkeys

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetSshkeys(v string)`

SetSshkeys sets Sshkeys field to given value.

### HasSshkeys

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasSshkeys() bool`

HasSshkeys returns a boolean if a field has been set.

### GetStartdate

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetStartdate() string`

GetStartdate returns the Startdate field if non-nil, zero value otherwise.

### GetStartdateOk

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetStartdateOk() (*string, bool)`

GetStartdateOk returns a tuple with the Startdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartdate

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetStartdate(v string)`

SetStartdate sets Startdate field to given value.

### HasStartdate

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasStartdate() bool`

HasStartdate returns a boolean if a field has been set.

### GetStartup

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetStartup() string`

GetStartup returns the Startup field if non-nil, zero value otherwise.

### GetStartupOk

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetStartupOk() (*string, bool)`

GetStartupOk returns a tuple with the Startup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartup

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetStartup(v string)`

SetStartup sets Startup field to given value.

### HasStartup

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasStartup() bool`

HasStartup returns a boolean if a field has been set.

### GetTablet

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetTablet() float32`

GetTablet returns the Tablet field if non-nil, zero value otherwise.

### GetTabletOk

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetTabletOk() (*float32, bool)`

GetTabletOk returns a tuple with the Tablet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTablet

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetTablet(v float32)`

SetTablet sets Tablet field to given value.

### HasTablet

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasTablet() bool`

HasTablet returns a boolean if a field has been set.

### GetTags

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTemplate

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetTemplate() float32`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetTemplateOk() (*float32, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetTemplate(v float32)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetTpmstate0

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetTpmstate0() string`

GetTpmstate0 returns the Tpmstate0 field if non-nil, zero value otherwise.

### GetTpmstate0Ok

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetTpmstate0Ok() (*string, bool)`

GetTpmstate0Ok returns a tuple with the Tpmstate0 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTpmstate0

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetTpmstate0(v string)`

SetTpmstate0 sets Tpmstate0 field to given value.

### HasTpmstate0

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasTpmstate0() bool`

HasTpmstate0 returns a boolean if a field has been set.

### GetUsb0

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetUsb0() string`

GetUsb0 returns the Usb0 field if non-nil, zero value otherwise.

### GetUsb0Ok

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetUsb0Ok() (*string, bool)`

GetUsb0Ok returns a tuple with the Usb0 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsb0

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetUsb0(v string)`

SetUsb0 sets Usb0 field to given value.

### HasUsb0

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasUsb0() bool`

HasUsb0 returns a boolean if a field has been set.

### GetUsb1

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetUsb1() string`

GetUsb1 returns the Usb1 field if non-nil, zero value otherwise.

### GetUsb1Ok

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetUsb1Ok() (*string, bool)`

GetUsb1Ok returns a tuple with the Usb1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsb1

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetUsb1(v string)`

SetUsb1 sets Usb1 field to given value.

### HasUsb1

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasUsb1() bool`

HasUsb1 returns a boolean if a field has been set.

### GetUsb2

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetUsb2() string`

GetUsb2 returns the Usb2 field if non-nil, zero value otherwise.

### GetUsb2Ok

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetUsb2Ok() (*string, bool)`

GetUsb2Ok returns a tuple with the Usb2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsb2

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetUsb2(v string)`

SetUsb2 sets Usb2 field to given value.

### HasUsb2

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasUsb2() bool`

HasUsb2 returns a boolean if a field has been set.

### GetUsb3

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetUsb3() string`

GetUsb3 returns the Usb3 field if non-nil, zero value otherwise.

### GetUsb3Ok

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetUsb3Ok() (*string, bool)`

GetUsb3Ok returns a tuple with the Usb3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsb3

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetUsb3(v string)`

SetUsb3 sets Usb3 field to given value.

### HasUsb3

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasUsb3() bool`

HasUsb3 returns a boolean if a field has been set.

### GetUsb4

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetUsb4() string`

GetUsb4 returns the Usb4 field if non-nil, zero value otherwise.

### GetUsb4Ok

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetUsb4Ok() (*string, bool)`

GetUsb4Ok returns a tuple with the Usb4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsb4

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetUsb4(v string)`

SetUsb4 sets Usb4 field to given value.

### HasUsb4

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasUsb4() bool`

HasUsb4 returns a boolean if a field has been set.

### GetUsb5

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetUsb5() string`

GetUsb5 returns the Usb5 field if non-nil, zero value otherwise.

### GetUsb5Ok

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetUsb5Ok() (*string, bool)`

GetUsb5Ok returns a tuple with the Usb5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsb5

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetUsb5(v string)`

SetUsb5 sets Usb5 field to given value.

### HasUsb5

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasUsb5() bool`

HasUsb5 returns a boolean if a field has been set.

### GetUsb6

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetUsb6() string`

GetUsb6 returns the Usb6 field if non-nil, zero value otherwise.

### GetUsb6Ok

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetUsb6Ok() (*string, bool)`

GetUsb6Ok returns a tuple with the Usb6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsb6

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetUsb6(v string)`

SetUsb6 sets Usb6 field to given value.

### HasUsb6

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasUsb6() bool`

HasUsb6 returns a boolean if a field has been set.

### GetUsb7

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetUsb7() string`

GetUsb7 returns the Usb7 field if non-nil, zero value otherwise.

### GetUsb7Ok

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetUsb7Ok() (*string, bool)`

GetUsb7Ok returns a tuple with the Usb7 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsb7

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetUsb7(v string)`

SetUsb7 sets Usb7 field to given value.

### HasUsb7

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasUsb7() bool`

HasUsb7 returns a boolean if a field has been set.

### GetUsb8

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetUsb8() string`

GetUsb8 returns the Usb8 field if non-nil, zero value otherwise.

### GetUsb8Ok

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetUsb8Ok() (*string, bool)`

GetUsb8Ok returns a tuple with the Usb8 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsb8

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetUsb8(v string)`

SetUsb8 sets Usb8 field to given value.

### HasUsb8

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasUsb8() bool`

HasUsb8 returns a boolean if a field has been set.

### GetUsb9

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetUsb9() string`

GetUsb9 returns the Usb9 field if non-nil, zero value otherwise.

### GetUsb9Ok

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetUsb9Ok() (*string, bool)`

GetUsb9Ok returns a tuple with the Usb9 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsb9

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetUsb9(v string)`

SetUsb9 sets Usb9 field to given value.

### HasUsb9

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasUsb9() bool`

HasUsb9 returns a boolean if a field has been set.

### GetUsb10

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetUsb10() string`

GetUsb10 returns the Usb10 field if non-nil, zero value otherwise.

### GetUsb10Ok

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetUsb10Ok() (*string, bool)`

GetUsb10Ok returns a tuple with the Usb10 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsb10

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetUsb10(v string)`

SetUsb10 sets Usb10 field to given value.

### HasUsb10

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasUsb10() bool`

HasUsb10 returns a boolean if a field has been set.

### GetUsb11

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetUsb11() string`

GetUsb11 returns the Usb11 field if non-nil, zero value otherwise.

### GetUsb11Ok

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetUsb11Ok() (*string, bool)`

GetUsb11Ok returns a tuple with the Usb11 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsb11

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetUsb11(v string)`

SetUsb11 sets Usb11 field to given value.

### HasUsb11

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasUsb11() bool`

HasUsb11 returns a boolean if a field has been set.

### GetUsb12

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetUsb12() string`

GetUsb12 returns the Usb12 field if non-nil, zero value otherwise.

### GetUsb12Ok

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetUsb12Ok() (*string, bool)`

GetUsb12Ok returns a tuple with the Usb12 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsb12

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetUsb12(v string)`

SetUsb12 sets Usb12 field to given value.

### HasUsb12

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasUsb12() bool`

HasUsb12 returns a boolean if a field has been set.

### GetUsb13

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetUsb13() string`

GetUsb13 returns the Usb13 field if non-nil, zero value otherwise.

### GetUsb13Ok

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetUsb13Ok() (*string, bool)`

GetUsb13Ok returns a tuple with the Usb13 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsb13

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetUsb13(v string)`

SetUsb13 sets Usb13 field to given value.

### HasUsb13

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasUsb13() bool`

HasUsb13 returns a boolean if a field has been set.

### GetUsb14

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetUsb14() string`

GetUsb14 returns the Usb14 field if non-nil, zero value otherwise.

### GetUsb14Ok

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetUsb14Ok() (*string, bool)`

GetUsb14Ok returns a tuple with the Usb14 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsb14

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetUsb14(v string)`

SetUsb14 sets Usb14 field to given value.

### HasUsb14

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasUsb14() bool`

HasUsb14 returns a boolean if a field has been set.

### GetHotplug

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetHotplug() string`

GetHotplug returns the Hotplug field if non-nil, zero value otherwise.

### GetHotplugOk

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetHotplugOk() (*string, bool)`

GetHotplugOk returns a tuple with the Hotplug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHotplug

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetHotplug(v string)`

SetHotplug sets Hotplug field to given value.

### HasHotplug

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasHotplug() bool`

HasHotplug returns a boolean if a field has been set.

### GetVcpus

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetVcpus() float32`

GetVcpus returns the Vcpus field if non-nil, zero value otherwise.

### GetVcpusOk

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetVcpusOk() (*float32, bool)`

GetVcpusOk returns a tuple with the Vcpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcpus

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetVcpus(v float32)`

SetVcpus sets Vcpus field to given value.

### HasVcpus

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasVcpus() bool`

HasVcpus returns a boolean if a field has been set.

### GetVga

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetVga() string`

GetVga returns the Vga field if non-nil, zero value otherwise.

### GetVgaOk

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetVgaOk() (*string, bool)`

GetVgaOk returns a tuple with the Vga field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVga

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetVga(v string)`

SetVga sets Vga field to given value.

### HasVga

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasVga() bool`

HasVga returns a boolean if a field has been set.

### GetVirtio0

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetVirtio0() string`

GetVirtio0 returns the Virtio0 field if non-nil, zero value otherwise.

### GetVirtio0Ok

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetVirtio0Ok() (*string, bool)`

GetVirtio0Ok returns a tuple with the Virtio0 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtio0

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetVirtio0(v string)`

SetVirtio0 sets Virtio0 field to given value.

### HasVirtio0

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasVirtio0() bool`

HasVirtio0 returns a boolean if a field has been set.

### GetVirtio1

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetVirtio1() string`

GetVirtio1 returns the Virtio1 field if non-nil, zero value otherwise.

### GetVirtio1Ok

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetVirtio1Ok() (*string, bool)`

GetVirtio1Ok returns a tuple with the Virtio1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtio1

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetVirtio1(v string)`

SetVirtio1 sets Virtio1 field to given value.

### HasVirtio1

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasVirtio1() bool`

HasVirtio1 returns a boolean if a field has been set.

### GetVirtio2

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetVirtio2() string`

GetVirtio2 returns the Virtio2 field if non-nil, zero value otherwise.

### GetVirtio2Ok

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetVirtio2Ok() (*string, bool)`

GetVirtio2Ok returns a tuple with the Virtio2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtio2

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetVirtio2(v string)`

SetVirtio2 sets Virtio2 field to given value.

### HasVirtio2

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasVirtio2() bool`

HasVirtio2 returns a boolean if a field has been set.

### GetVirtio3

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetVirtio3() string`

GetVirtio3 returns the Virtio3 field if non-nil, zero value otherwise.

### GetVirtio3Ok

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetVirtio3Ok() (*string, bool)`

GetVirtio3Ok returns a tuple with the Virtio3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtio3

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetVirtio3(v string)`

SetVirtio3 sets Virtio3 field to given value.

### HasVirtio3

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasVirtio3() bool`

HasVirtio3 returns a boolean if a field has been set.

### GetVirtio4

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetVirtio4() string`

GetVirtio4 returns the Virtio4 field if non-nil, zero value otherwise.

### GetVirtio4Ok

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetVirtio4Ok() (*string, bool)`

GetVirtio4Ok returns a tuple with the Virtio4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtio4

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetVirtio4(v string)`

SetVirtio4 sets Virtio4 field to given value.

### HasVirtio4

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasVirtio4() bool`

HasVirtio4 returns a boolean if a field has been set.

### GetVirtio5

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetVirtio5() string`

GetVirtio5 returns the Virtio5 field if non-nil, zero value otherwise.

### GetVirtio5Ok

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetVirtio5Ok() (*string, bool)`

GetVirtio5Ok returns a tuple with the Virtio5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtio5

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetVirtio5(v string)`

SetVirtio5 sets Virtio5 field to given value.

### HasVirtio5

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasVirtio5() bool`

HasVirtio5 returns a boolean if a field has been set.

### GetVirtio6

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetVirtio6() string`

GetVirtio6 returns the Virtio6 field if non-nil, zero value otherwise.

### GetVirtio6Ok

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetVirtio6Ok() (*string, bool)`

GetVirtio6Ok returns a tuple with the Virtio6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtio6

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetVirtio6(v string)`

SetVirtio6 sets Virtio6 field to given value.

### HasVirtio6

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasVirtio6() bool`

HasVirtio6 returns a boolean if a field has been set.

### GetVirtio7

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetVirtio7() string`

GetVirtio7 returns the Virtio7 field if non-nil, zero value otherwise.

### GetVirtio7Ok

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetVirtio7Ok() (*string, bool)`

GetVirtio7Ok returns a tuple with the Virtio7 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtio7

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetVirtio7(v string)`

SetVirtio7 sets Virtio7 field to given value.

### HasVirtio7

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasVirtio7() bool`

HasVirtio7 returns a boolean if a field has been set.

### GetVirtio8

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetVirtio8() string`

GetVirtio8 returns the Virtio8 field if non-nil, zero value otherwise.

### GetVirtio8Ok

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetVirtio8Ok() (*string, bool)`

GetVirtio8Ok returns a tuple with the Virtio8 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtio8

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetVirtio8(v string)`

SetVirtio8 sets Virtio8 field to given value.

### HasVirtio8

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasVirtio8() bool`

HasVirtio8 returns a boolean if a field has been set.

### GetVirtio9

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetVirtio9() string`

GetVirtio9 returns the Virtio9 field if non-nil, zero value otherwise.

### GetVirtio9Ok

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetVirtio9Ok() (*string, bool)`

GetVirtio9Ok returns a tuple with the Virtio9 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtio9

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetVirtio9(v string)`

SetVirtio9 sets Virtio9 field to given value.

### HasVirtio9

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasVirtio9() bool`

HasVirtio9 returns a boolean if a field has been set.

### GetVirtio10

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetVirtio10() string`

GetVirtio10 returns the Virtio10 field if non-nil, zero value otherwise.

### GetVirtio10Ok

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetVirtio10Ok() (*string, bool)`

GetVirtio10Ok returns a tuple with the Virtio10 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtio10

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetVirtio10(v string)`

SetVirtio10 sets Virtio10 field to given value.

### HasVirtio10

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasVirtio10() bool`

HasVirtio10 returns a boolean if a field has been set.

### GetVirtio11

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetVirtio11() string`

GetVirtio11 returns the Virtio11 field if non-nil, zero value otherwise.

### GetVirtio11Ok

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetVirtio11Ok() (*string, bool)`

GetVirtio11Ok returns a tuple with the Virtio11 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtio11

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetVirtio11(v string)`

SetVirtio11 sets Virtio11 field to given value.

### HasVirtio11

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasVirtio11() bool`

HasVirtio11 returns a boolean if a field has been set.

### GetVirtio12

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetVirtio12() string`

GetVirtio12 returns the Virtio12 field if non-nil, zero value otherwise.

### GetVirtio12Ok

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetVirtio12Ok() (*string, bool)`

GetVirtio12Ok returns a tuple with the Virtio12 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtio12

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetVirtio12(v string)`

SetVirtio12 sets Virtio12 field to given value.

### HasVirtio12

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasVirtio12() bool`

HasVirtio12 returns a boolean if a field has been set.

### GetVirtio13

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetVirtio13() string`

GetVirtio13 returns the Virtio13 field if non-nil, zero value otherwise.

### GetVirtio13Ok

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetVirtio13Ok() (*string, bool)`

GetVirtio13Ok returns a tuple with the Virtio13 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtio13

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetVirtio13(v string)`

SetVirtio13 sets Virtio13 field to given value.

### HasVirtio13

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasVirtio13() bool`

HasVirtio13 returns a boolean if a field has been set.

### GetVirtio14

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetVirtio14() string`

GetVirtio14 returns the Virtio14 field if non-nil, zero value otherwise.

### GetVirtio14Ok

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetVirtio14Ok() (*string, bool)`

GetVirtio14Ok returns a tuple with the Virtio14 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtio14

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetVirtio14(v string)`

SetVirtio14 sets Virtio14 field to given value.

### HasVirtio14

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasVirtio14() bool`

HasVirtio14 returns a boolean if a field has been set.

### GetVirtio15

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetVirtio15() string`

GetVirtio15 returns the Virtio15 field if non-nil, zero value otherwise.

### GetVirtio15Ok

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetVirtio15Ok() (*string, bool)`

GetVirtio15Ok returns a tuple with the Virtio15 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtio15

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetVirtio15(v string)`

SetVirtio15 sets Virtio15 field to given value.

### HasVirtio15

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasVirtio15() bool`

HasVirtio15 returns a boolean if a field has been set.

### GetVmgenid

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetVmgenid() string`

GetVmgenid returns the Vmgenid field if non-nil, zero value otherwise.

### GetVmgenidOk

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetVmgenidOk() (*string, bool)`

GetVmgenidOk returns a tuple with the Vmgenid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmgenid

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetVmgenid(v string)`

SetVmgenid sets Vmgenid field to given value.

### HasVmgenid

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasVmgenid() bool`

HasVmgenid returns a boolean if a field has been set.

### GetVmstatestorage

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetVmstatestorage() string`

GetVmstatestorage returns the Vmstatestorage field if non-nil, zero value otherwise.

### GetVmstatestorageOk

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetVmstatestorageOk() (*string, bool)`

GetVmstatestorageOk returns a tuple with the Vmstatestorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmstatestorage

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetVmstatestorage(v string)`

SetVmstatestorage sets Vmstatestorage field to given value.

### HasVmstatestorage

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasVmstatestorage() bool`

HasVmstatestorage returns a boolean if a field has been set.

### GetWatchdog

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetWatchdog() string`

GetWatchdog returns the Watchdog field if non-nil, zero value otherwise.

### GetWatchdogOk

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) GetWatchdogOk() (*string, bool)`

GetWatchdogOk returns a tuple with the Watchdog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatchdog

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) SetWatchdog(v string)`

SetWatchdog sets Watchdog field to given value.

### HasWatchdog

`func (o *ApplyVirtualMachineConfigurationAsyncRequestContent) HasWatchdog() bool`

HasWatchdog returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


