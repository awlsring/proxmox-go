# \DefaultApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddCorosyncNode**](DefaultApi.md#AddCorosyncNode) | **Post** /cluster/config/nodes/{node} | 
[**AddCustomNodeCertificate**](DefaultApi.md#AddCustomNodeCertificate) | **Post** /nodes/{node}/certificates/custom | 
[**AddRepository**](DefaultApi.md#AddRepository) | **Put** /nodes/{node}/apt/repositories | 
[**ApplyNetworkInterfaceConfiguration**](DefaultApi.md#ApplyNetworkInterfaceConfiguration) | **Put** /nodes/{node}/network | 
[**ApplyVirtualMachineConfigurationAsync**](DefaultApi.md#ApplyVirtualMachineConfigurationAsync) | **Post** /nodes/{node}/qemu/{vmId}/config | 
[**ApplyVirtualMachineConfigurationSync**](DefaultApi.md#ApplyVirtualMachineConfigurationSync) | **Put** /nodes/{node}/qemu/{vmId}/config | 
[**ChangePassword**](DefaultApi.md#ChangePassword) | **Put** /access/password | 
[**ChangeRepositoryProperties**](DefaultApi.md#ChangeRepositoryProperties) | **Post** /nodes/{node}/apt/repositories | 
[**CloneVirtualMachine**](DefaultApi.md#CloneVirtualMachine) | **Post** /nodes/{node}/qemu/{vmId}/clone | 
[**CreateCephManager**](DefaultApi.md#CreateCephManager) | **Post** /nodes/{node}/ceph/mgr/{manager} | 
[**CreateCephMonitor**](DefaultApi.md#CreateCephMonitor) | **Post** /nodes/{node}/ceph/mon/{monitor} | 
[**CreateClusterConfig**](DefaultApi.md#CreateClusterConfig) | **Post** /cluster/config | 
[**CreateDirectory**](DefaultApi.md#CreateDirectory) | **Post** /nodes/{node}/disks/directory | 
[**CreateGroup**](DefaultApi.md#CreateGroup) | **Post** /access/groups | 
[**CreateLVM**](DefaultApi.md#CreateLVM) | **Post** /nodes/{node}/disks/lvm | 
[**CreateLVMThin**](DefaultApi.md#CreateLVMThin) | **Post** /nodes/{node}/disks/lvmthin | 
[**CreateNetworkInterface**](DefaultApi.md#CreateNetworkInterface) | **Post** /nodes/{node}/network | 
[**CreatePool**](DefaultApi.md#CreatePool) | **Post** /pools | 
[**CreateRealm**](DefaultApi.md#CreateRealm) | **Post** /access/domains | 
[**CreateRole**](DefaultApi.md#CreateRole) | **Post** /access/roles | 
[**CreateSnapshot**](DefaultApi.md#CreateSnapshot) | **Post** /nodes/{node}/qemu/{vmId}/snapshot | 
[**CreateStorage**](DefaultApi.md#CreateStorage) | **Post** /storage | 
[**CreateStorageVolume**](DefaultApi.md#CreateStorageVolume) | **Post** /nodes/{node}/storage/{storage}/content | 
[**CreateTicket**](DefaultApi.md#CreateTicket) | **Post** /access/ticket | 
[**CreateUser**](DefaultApi.md#CreateUser) | **Post** /access/users | 
[**CreateUserToken**](DefaultApi.md#CreateUserToken) | **Post** /access/users/{userId}/token/{tokenId} | 
[**CreateVirtualMachine**](DefaultApi.md#CreateVirtualMachine) | **Post** /nodes/{node}/qemu | 
[**CreateVirtualMachineTemplate**](DefaultApi.md#CreateVirtualMachineTemplate) | **Post** /nodes/{node}/qemu/{vmId}/template | 
[**CreateZFSPool**](DefaultApi.md#CreateZFSPool) | **Post** /nodes/{node}/disks/zfs | 
[**DeleteCephManager**](DefaultApi.md#DeleteCephManager) | **Delete** /nodes/{node}/ceph/mgr/{manager} | 
[**DeleteCephMonitor**](DefaultApi.md#DeleteCephMonitor) | **Delete** /nodes/{node}/ceph/mon/{monitor} | 
[**DeleteDirectory**](DefaultApi.md#DeleteDirectory) | **Delete** /nodes/{node}/disks/directory/{name} | 
[**DeleteGroup**](DefaultApi.md#DeleteGroup) | **Delete** /access/groups/{groupId} | 
[**DeleteLVM**](DefaultApi.md#DeleteLVM) | **Delete** /nodes/{node}/disks/lvm/{name} | 
[**DeleteLVMThin**](DefaultApi.md#DeleteLVMThin) | **Delete** /nodes/{node}/disks/lvmthin/{name} | 
[**DeleteNetworkInterface**](DefaultApi.md#DeleteNetworkInterface) | **Delete** /nodes/{node}/network/{interface} | 
[**DeleteNodeCertificate**](DefaultApi.md#DeleteNodeCertificate) | **Delete** /nodes/{node}/certificates/acme/certificate | 
[**DeletePool**](DefaultApi.md#DeletePool) | **Delete** /pools/{poolId} | 
[**DeleteRealm**](DefaultApi.md#DeleteRealm) | **Delete** /access/domains/{realm} | 
[**DeleteRole**](DefaultApi.md#DeleteRole) | **Delete** /access/roles/{roleId} | 
[**DeleteStorage**](DefaultApi.md#DeleteStorage) | **Delete** /storage/{storage} | 
[**DeleteStorageVolume**](DefaultApi.md#DeleteStorageVolume) | **Delete** /nodes/{node}/storage/{storage}/content/{volume} | 
[**DeleteUser**](DefaultApi.md#DeleteUser) | **Delete** /access/users/{userId} | 
[**DeleteUserToken**](DefaultApi.md#DeleteUserToken) | **Delete** /access/users/{userId}/token/{tokenId} | 
[**DeleteVirtualMachine**](DefaultApi.md#DeleteVirtualMachine) | **Delete** /nodes/{node}/qemu/{vmId} | 
[**DeleteZFSPool**](DefaultApi.md#DeleteZFSPool) | **Delete** /nodes/{node}/disks/zfs/{name} | 
[**DownloadFromUrlToStorage**](DefaultApi.md#DownloadFromUrlToStorage) | **Post** /nodes/{node}/storage/{storage}/download-url | 
[**GetAccessControlList**](DefaultApi.md#GetAccessControlList) | **Get** /access/acl | 
[**GetClusterApiVersion**](DefaultApi.md#GetClusterApiVersion) | **Get** /cluster/config/apiversion | 
[**GetClusterJoinInformation**](DefaultApi.md#GetClusterJoinInformation) | **Get** /cluster/config/join | 
[**GetClusterTotemSettings**](DefaultApi.md#GetClusterTotemSettings) | **Get** /cluster/config/totem | 
[**GetGroupDetails**](DefaultApi.md#GetGroupDetails) | **Get** /access/groups/{groupId} | 
[**GetNetworkInterface**](DefaultApi.md#GetNetworkInterface) | **Get** /nodes/{node}/network/{interface} | 
[**GetNodeDns**](DefaultApi.md#GetNodeDns) | **Get** /nodes/{node}/dns | 
[**GetPackageChangelog**](DefaultApi.md#GetPackageChangelog) | **Get** /nodes/{node}/apt/changelog | 
[**GetPendingVirtualMachineCloudInitChanges**](DefaultApi.md#GetPendingVirtualMachineCloudInitChanges) | **Get** /nodes/{node}/qemu/{vmId}/cloudinit | 
[**GetPermissions**](DefaultApi.md#GetPermissions) | **Get** /access/permissions | 
[**GetPool**](DefaultApi.md#GetPool) | **Get** /pools/{poolId} | 
[**GetRealm**](DefaultApi.md#GetRealm) | **Get** /access/domains/{realm} | 
[**GetRole**](DefaultApi.md#GetRole) | **Get** /access/roles/{roleId} | 
[**GetSmartHealth**](DefaultApi.md#GetSmartHealth) | **Get** /nodes/{node}/disks/smart | 
[**GetStorage**](DefaultApi.md#GetStorage) | **Get** /storage/{storage} | 
[**GetStorageStatus**](DefaultApi.md#GetStorageStatus) | **Get** /nodes/{node}/storage/{storage}/status | 
[**GetUserConfiguration**](DefaultApi.md#GetUserConfiguration) | **Get** /access/users/{userId} | 
[**GetUserToken**](DefaultApi.md#GetUserToken) | **Get** /access/users/{userId}/token/{tokenId} | 
[**GetVersion**](DefaultApi.md#GetVersion) | **Get** /version | 
[**GetVirtualMachineAgentInfo**](DefaultApi.md#GetVirtualMachineAgentInfo) | **Get** /nodes/{node}/qemu/{vmId}/agent/get-info | 
[**GetVirtualMachineCloudInit**](DefaultApi.md#GetVirtualMachineCloudInit) | **Get** /nodes/{node}/qemu/{vmId}/cloudinit/dump | 
[**GetVirtualMachineCommandStatus**](DefaultApi.md#GetVirtualMachineCommandStatus) | **Get** /nodes/{node}/qemu/{vmId}/agent/exec-status | 
[**GetVirtualMachineConfiguration**](DefaultApi.md#GetVirtualMachineConfiguration) | **Get** /nodes/{node}/qemu/{vmId}/config | 
[**GetVirtualMachineFeatureSupport**](DefaultApi.md#GetVirtualMachineFeatureSupport) | **Get** /nodes/{node}/qemu/{vmId}/feature | 
[**GetVirtualMachineFileSystemInformation**](DefaultApi.md#GetVirtualMachineFileSystemInformation) | **Get** /nodes/{node}/qemu/{vmId}/agent/get-fsinfo | 
[**GetVirtualMachineHostname**](DefaultApi.md#GetVirtualMachineHostname) | **Get** /nodes/{node}/qemu/{vmId}/agent/get-host-name | 
[**GetVirtualMachineMemoryBlockInformation**](DefaultApi.md#GetVirtualMachineMemoryBlockInformation) | **Get** /nodes/{node}/qemu/{vmId}/agent/get-memory-block-info | 
[**GetVirtualMachineMemoryBlocks**](DefaultApi.md#GetVirtualMachineMemoryBlocks) | **Get** /nodes/{node}/qemu/{vmId}/agent/get-memory-blocks | 
[**GetVirtualMachineNetworkInterfaces**](DefaultApi.md#GetVirtualMachineNetworkInterfaces) | **Get** /nodes/{node}/qemu/{vmId}/agent/network-get-interfaces | 
[**GetVirtualMachineOperatingSystemInformation**](DefaultApi.md#GetVirtualMachineOperatingSystemInformation) | **Get** /nodes/{node}/qemu/{vmId}/agent/get-osinfo | 
[**GetVirtualMachineStatus**](DefaultApi.md#GetVirtualMachineStatus) | **Get** /nodes/{node}/qemu/{vmId}/status/current | 
[**GetVirtualMachineTime**](DefaultApi.md#GetVirtualMachineTime) | **Get** /nodes/{node}/qemu/{vmId}/agent/get-time | 
[**GetVirtualMachineTimeZone**](DefaultApi.md#GetVirtualMachineTimeZone) | **Get** /nodes/{node}/qemu/{vmId}/agent/get-timezone | 
[**GetVirtualMachineVcpus**](DefaultApi.md#GetVirtualMachineVcpus) | **Get** /nodes/{node}/qemu/{vmId}/agent/get-vcpus | 
[**GetZFSPoolStatus**](DefaultApi.md#GetZFSPoolStatus) | **Get** /nodes/{node}/disks/zfs/{name} | 
[**InitCephCluster**](DefaultApi.md#InitCephCluster) | **Post** /nodes/{node}/ceph/init | 
[**InitializeGPT**](DefaultApi.md#InitializeGPT) | **Post** /nodes/{node}/disks/smart | 
[**JoinCluster**](DefaultApi.md#JoinCluster) | **Post** /cluster/config/join | 
[**ListCephManagers**](DefaultApi.md#ListCephManagers) | **Get** /nodes/{node}/ceph/mgr | 
[**ListCephMonitors**](DefaultApi.md#ListCephMonitors) | **Get** /nodes/{node}/ceph/mon | 
[**ListCorosyncNodes**](DefaultApi.md#ListCorosyncNodes) | **Get** /cluster/config/nodes | 
[**ListCpuCapabilities**](DefaultApi.md#ListCpuCapabilities) | **Get** /nodes/{node}/capabilities/qemu/cpu | 
[**ListDirectories**](DefaultApi.md#ListDirectories) | **Get** /nodes/{node}/disks/directory | 
[**ListDisks**](DefaultApi.md#ListDisks) | **Get** /nodes/{node}/disks/list | 
[**ListGroups**](DefaultApi.md#ListGroups) | **Get** /access/groups | 
[**ListLVMThins**](DefaultApi.md#ListLVMThins) | **Get** /nodes/{node}/disks/lvmthin | 
[**ListLVMs**](DefaultApi.md#ListLVMs) | **Get** /nodes/{node}/disks/lvm | 
[**ListMachineCapabilities**](DefaultApi.md#ListMachineCapabilities) | **Get** /nodes/{node}/capabilities/qemu/machines | 
[**ListNetworkInterfaces**](DefaultApi.md#ListNetworkInterfaces) | **Get** /nodes/{node}/network | 
[**ListNodeCertificates**](DefaultApi.md#ListNodeCertificates) | **Get** /nodes/{node}/certificates/info | 
[**ListNodeStorage**](DefaultApi.md#ListNodeStorage) | **Get** /nodes/{node}/storage | 
[**ListNodes**](DefaultApi.md#ListNodes) | **Get** /nodes | 
[**ListPackages**](DefaultApi.md#ListPackages) | **Get** /nodes/{node}/apt/versions | 
[**ListPciDeviceMediatedDevices**](DefaultApi.md#ListPciDeviceMediatedDevices) | **Get** /nodes/{node}/hardware/pci/{deviceId} | 
[**ListPciDevices**](DefaultApi.md#ListPciDevices) | **Get** /nodes/{node}/hardware/pci | 
[**ListPendingVirtualMachineConfigurationChanges**](DefaultApi.md#ListPendingVirtualMachineConfigurationChanges) | **Get** /nodes/{node}/qemu/{vmId}/pending | 
[**ListPools**](DefaultApi.md#ListPools) | **Get** /pools | 
[**ListRealms**](DefaultApi.md#ListRealms) | **Get** /access/domains | 
[**ListRepositoriesInformation**](DefaultApi.md#ListRepositoriesInformation) | **Get** /nodes/{node}/apt/repository | 
[**ListRoles**](DefaultApi.md#ListRoles) | **Get** /access/roles | 
[**ListSnapshots**](DefaultApi.md#ListSnapshots) | **Get** /nodes/{node}/qemu/{vmId}/snapshot | 
[**ListStorage**](DefaultApi.md#ListStorage) | **Get** /storage | 
[**ListStorageVolumes**](DefaultApi.md#ListStorageVolumes) | **Get** /nodes/{node}/storage/{storage}/content | 
[**ListUpdates**](DefaultApi.md#ListUpdates) | **Get** /nodes/{node}/apt/update | 
[**ListUsbDevices**](DefaultApi.md#ListUsbDevices) | **Get** /nodes/{node}/hardware/usb | 
[**ListUserTokens**](DefaultApi.md#ListUserTokens) | **Get** /access/users/{userId}/token | 
[**ListUsers**](DefaultApi.md#ListUsers) | **Get** /access/users | 
[**ListVirtualMachineFirewallReferences**](DefaultApi.md#ListVirtualMachineFirewallReferences) | **Get** /nodes/{node}/qemu/{vmId}/firewall/refs | 
[**ListVirtualMachines**](DefaultApi.md#ListVirtualMachines) | **Get** /nodes/{node}/qemu | 
[**ListZFSPools**](DefaultApi.md#ListZFSPools) | **Get** /nodes/{node}/disks/zfs | 
[**ModifyPool**](DefaultApi.md#ModifyPool) | **Put** /pools/{poolId} | 
[**ModifyRole**](DefaultApi.md#ModifyRole) | **Put** /access/roles/{roleId} | 
[**ModifyStorage**](DefaultApi.md#ModifyStorage) | **Put** /storage/{storage} | 
[**ModifyUser**](DefaultApi.md#ModifyUser) | **Put** /access/users/{userId} | 
[**ModifyUserToken**](DefaultApi.md#ModifyUserToken) | **Put** /access/users/{userId}/token/{tokenId} | 
[**OrderNodeCertificate**](DefaultApi.md#OrderNodeCertificate) | **Post** /nodes/{node}/certificates/acme/certificate | 
[**PingVirtualMachine**](DefaultApi.md#PingVirtualMachine) | **Post** /nodes/{node}/qemu/{vmId}/agent/ping | 
[**RebootVirtualMachine**](DefaultApi.md#RebootVirtualMachine) | **Post** /nodes/{node}/qemu/{vmId}/status/reboot | 
[**RegenerateVirtualMachineCloudInit**](DefaultApi.md#RegenerateVirtualMachineCloudInit) | **Put** /nodes/{node}/qemu/{vmId}/cloudinit | 
[**RemoveCorosyncNode**](DefaultApi.md#RemoveCorosyncNode) | **Delete** /cluster/config/nodes/{node} | 
[**RenewNodeCertificate**](DefaultApi.md#RenewNodeCertificate) | **Put** /nodes/{node}/certificates/acme/certificate | 
[**ResetVirtualMachine**](DefaultApi.md#ResetVirtualMachine) | **Post** /nodes/{node}/qemu/{vmId}/status/reset | 
[**ResizeVirtualMachineDisk**](DefaultApi.md#ResizeVirtualMachineDisk) | **Put** /nodes/{node}/qemu/{vmId}/resize | 
[**ResumeVirtualMachine**](DefaultApi.md#ResumeVirtualMachine) | **Post** /nodes/{node}/qemu/{vmId}/status/resume | 
[**RevertNetworkInterfaceConfiguration**](DefaultApi.md#RevertNetworkInterfaceConfiguration) | **Delete** /nodes/{node}/network | 
[**ShutdownVirtualMachine**](DefaultApi.md#ShutdownVirtualMachine) | **Post** /nodes/{node}/qemu/{vmId}/status/shutdown | 
[**StartVirtualMachine**](DefaultApi.md#StartVirtualMachine) | **Post** /nodes/{node}/qemu/{vmId}/status/start | 
[**StopVirtualMachine**](DefaultApi.md#StopVirtualMachine) | **Post** /nodes/{node}/qemu/{vmId}/status/stop | 
[**SuspendVirtualMachine**](DefaultApi.md#SuspendVirtualMachine) | **Post** /nodes/{node}/qemu/{vmId}/status/suspend | 
[**SyncRealm**](DefaultApi.md#SyncRealm) | **Post** /access/domains/{realm} | 
[**UnlinkVirtualMachineDisks**](DefaultApi.md#UnlinkVirtualMachineDisks) | **Put** /nodes/{node}/qemu/{vmId}/unlink | 
[**UpdateAccessControlList**](DefaultApi.md#UpdateAccessControlList) | **Put** /access/acl | 
[**UpdateGroup**](DefaultApi.md#UpdateGroup) | **Put** /access/groups/{groupId} | 
[**UpdateNetworkInterface**](DefaultApi.md#UpdateNetworkInterface) | **Put** /nodes/{node}/network/{interface} | 
[**UpdateNodeDns**](DefaultApi.md#UpdateNodeDns) | **Put** /nodes/{node}/dns | 
[**UpdateRealm**](DefaultApi.md#UpdateRealm) | **Put** /access/domains/{realm} | 
[**UpdateStorageVolume**](DefaultApi.md#UpdateStorageVolume) | **Put** /nodes/{node}/storage/{storage}/content/{volume} | 
[**UploadToStorage**](DefaultApi.md#UploadToStorage) | **Post** /nodes/{node}/storage/{storage}/upload | 
[**VirtualMachineExecuteCommand**](DefaultApi.md#VirtualMachineExecuteCommand) | **Post** /nodes/{node}/qemu/{vmId}/agent/exec | 
[**VirtualMachineReadFile**](DefaultApi.md#VirtualMachineReadFile) | **Get** /nodes/{node}/qemu/{vmId}/agent/file-read | 
[**VirtualMachineWriteFile**](DefaultApi.md#VirtualMachineWriteFile) | **Post** /nodes/{node}/qemu/{vmId}/agent/file-write | 
[**WipeDisk**](DefaultApi.md#WipeDisk) | **Put** /nodes/{node}/disks/smart | 



## AddCorosyncNode

> AddCorosyncNodeResponseContent AddCorosyncNode(ctx, node).AddCorosyncNodeRequestContent(addCorosyncNodeRequestContent).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    node := "node_example" // string | 
    addCorosyncNodeRequestContent := *openapiclient.NewAddCorosyncNodeRequestContent() // AddCorosyncNodeRequestContent |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.AddCorosyncNode(context.Background(), node).AddCorosyncNodeRequestContent(addCorosyncNodeRequestContent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.AddCorosyncNode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddCorosyncNode`: AddCorosyncNodeResponseContent
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.AddCorosyncNode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddCorosyncNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addCorosyncNodeRequestContent** | [**AddCorosyncNodeRequestContent**](AddCorosyncNodeRequestContent.md) |  | 

### Return type

[**AddCorosyncNodeResponseContent**](AddCorosyncNodeResponseContent.md)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddCustomNodeCertificate

> AddCustomNodeCertificateResponseContent AddCustomNodeCertificate(ctx, node).AddCustomNodeCertificateRequestContent(addCustomNodeCertificateRequestContent).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    node := "node_example" // string | 
    addCustomNodeCertificateRequestContent := *openapiclient.NewAddCustomNodeCertificateRequestContent("Certificates_example") // AddCustomNodeCertificateRequestContent | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.AddCustomNodeCertificate(context.Background(), node).AddCustomNodeCertificateRequestContent(addCustomNodeCertificateRequestContent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.AddCustomNodeCertificate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddCustomNodeCertificate`: AddCustomNodeCertificateResponseContent
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.AddCustomNodeCertificate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddCustomNodeCertificateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addCustomNodeCertificateRequestContent** | [**AddCustomNodeCertificateRequestContent**](AddCustomNodeCertificateRequestContent.md) |  | 

### Return type

[**AddCustomNodeCertificateResponseContent**](AddCustomNodeCertificateResponseContent.md)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddRepository

> AddRepository(ctx, node).Handle(handle).Digest(digest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    node := "node_example" // string | 
    handle := "handle_example" // string | Handle that identifies the repository
    digest := "digest_example" // string | Digest to detect modifications (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultApi.AddRepository(context.Background(), node).Handle(handle).Digest(digest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.AddRepository``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddRepositoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **handle** | **string** | Handle that identifies the repository | 
 **digest** | **string** | Digest to detect modifications | 

### Return type

 (empty response body)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplyNetworkInterfaceConfiguration

> ApplyNetworkInterfaceConfigurationResponseContent ApplyNetworkInterfaceConfiguration(ctx, node).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    node := "node_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ApplyNetworkInterfaceConfiguration(context.Background(), node).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApplyNetworkInterfaceConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApplyNetworkInterfaceConfiguration`: ApplyNetworkInterfaceConfigurationResponseContent
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ApplyNetworkInterfaceConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplyNetworkInterfaceConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApplyNetworkInterfaceConfigurationResponseContent**](ApplyNetworkInterfaceConfigurationResponseContent.md)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplyVirtualMachineConfigurationAsync

> ApplyVirtualMachineConfigurationAsyncResponseContent ApplyVirtualMachineConfigurationAsync(ctx, node, vmId).ApplyVirtualMachineConfigurationAsyncRequestContent(applyVirtualMachineConfigurationAsyncRequestContent).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    node := "node_example" // string | 
    vmId := "vmId_example" // string | The id of the virtual machine as a string
    applyVirtualMachineConfigurationAsyncRequestContent := *openapiclient.NewApplyVirtualMachineConfigurationAsyncRequestContent() // ApplyVirtualMachineConfigurationAsyncRequestContent |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ApplyVirtualMachineConfigurationAsync(context.Background(), node, vmId).ApplyVirtualMachineConfigurationAsyncRequestContent(applyVirtualMachineConfigurationAsyncRequestContent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApplyVirtualMachineConfigurationAsync``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApplyVirtualMachineConfigurationAsync`: ApplyVirtualMachineConfigurationAsyncResponseContent
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ApplyVirtualMachineConfigurationAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** |  | 
**vmId** | **string** | The id of the virtual machine as a string | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplyVirtualMachineConfigurationAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **applyVirtualMachineConfigurationAsyncRequestContent** | [**ApplyVirtualMachineConfigurationAsyncRequestContent**](ApplyVirtualMachineConfigurationAsyncRequestContent.md) |  | 

### Return type

[**ApplyVirtualMachineConfigurationAsyncResponseContent**](ApplyVirtualMachineConfigurationAsyncResponseContent.md)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplyVirtualMachineConfigurationSync

> ApplyVirtualMachineConfigurationSync(ctx, node, vmId).ApplyVirtualMachineConfigurationSyncRequestContent(applyVirtualMachineConfigurationSyncRequestContent).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    node := "node_example" // string | 
    vmId := "vmId_example" // string | The id of the virtual machine as a string
    applyVirtualMachineConfigurationSyncRequestContent := *openapiclient.NewApplyVirtualMachineConfigurationSyncRequestContent() // ApplyVirtualMachineConfigurationSyncRequestContent |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultApi.ApplyVirtualMachineConfigurationSync(context.Background(), node, vmId).ApplyVirtualMachineConfigurationSyncRequestContent(applyVirtualMachineConfigurationSyncRequestContent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApplyVirtualMachineConfigurationSync``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** |  | 
**vmId** | **string** | The id of the virtual machine as a string | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplyVirtualMachineConfigurationSyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **applyVirtualMachineConfigurationSyncRequestContent** | [**ApplyVirtualMachineConfigurationSyncRequestContent**](ApplyVirtualMachineConfigurationSyncRequestContent.md) |  | 

### Return type

 (empty response body)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChangePassword

> ChangePassword(ctx).ChangePasswordRequestContent(changePasswordRequestContent).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    changePasswordRequestContent := *openapiclient.NewChangePasswordRequestContent("Userid_example", "Password_example") // ChangePasswordRequestContent | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultApi.ChangePassword(context.Background()).ChangePasswordRequestContent(changePasswordRequestContent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ChangePassword``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChangePasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **changePasswordRequestContent** | [**ChangePasswordRequestContent**](ChangePasswordRequestContent.md) |  | 

### Return type

 (empty response body)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChangeRepositoryProperties

> ChangeRepositoryProperties(ctx, node).Path(path).Index(index).Digest(digest).Enabled(enabled).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    node := "node_example" // string | 
    path := "path_example" // string | Path to the containing file
    index := float32(8.14) // float32 | Index within the file
    digest := "digest_example" // string | Digest to detect modifications (optional)
    enabled := true // bool | Wether the repository is enabled (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultApi.ChangeRepositoryProperties(context.Background(), node).Path(path).Index(index).Digest(digest).Enabled(enabled).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ChangeRepositoryProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiChangeRepositoryPropertiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **path** | **string** | Path to the containing file | 
 **index** | **float32** | Index within the file | 
 **digest** | **string** | Digest to detect modifications | 
 **enabled** | **bool** | Wether the repository is enabled | 

### Return type

 (empty response body)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloneVirtualMachine

> CloneVirtualMachineResponseContent CloneVirtualMachine(ctx, node, vmId).CloneVirtualMachineRequestContent(cloneVirtualMachineRequestContent).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    node := "node_example" // string | 
    vmId := "vmId_example" // string | The id of the virtual machine as a string
    cloneVirtualMachineRequestContent := *openapiclient.NewCloneVirtualMachineRequestContent(float32(123)) // CloneVirtualMachineRequestContent | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CloneVirtualMachine(context.Background(), node, vmId).CloneVirtualMachineRequestContent(cloneVirtualMachineRequestContent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CloneVirtualMachine``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CloneVirtualMachine`: CloneVirtualMachineResponseContent
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CloneVirtualMachine`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** |  | 
**vmId** | **string** | The id of the virtual machine as a string | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloneVirtualMachineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cloneVirtualMachineRequestContent** | [**CloneVirtualMachineRequestContent**](CloneVirtualMachineRequestContent.md) |  | 

### Return type

[**CloneVirtualMachineResponseContent**](CloneVirtualMachineResponseContent.md)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCephManager

> CreateCephManagerResponseContent CreateCephManager(ctx, node, manager).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    node := "node_example" // string | 
    manager := "manager_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CreateCephManager(context.Background(), node, manager).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateCephManager``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCephManager`: CreateCephManagerResponseContent
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateCephManager`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** |  | 
**manager** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCephManagerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CreateCephManagerResponseContent**](CreateCephManagerResponseContent.md)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCephMonitor

> CreateCephMonitorResponseContent CreateCephMonitor(ctx, node, monitor).CreateCephMonitorRequestContent(createCephMonitorRequestContent).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    node := "node_example" // string | 
    monitor := "monitor_example" // string | 
    createCephMonitorRequestContent := *openapiclient.NewCreateCephMonitorRequestContent() // CreateCephMonitorRequestContent |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CreateCephMonitor(context.Background(), node, monitor).CreateCephMonitorRequestContent(createCephMonitorRequestContent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateCephMonitor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCephMonitor`: CreateCephMonitorResponseContent
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateCephMonitor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** |  | 
**monitor** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCephMonitorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createCephMonitorRequestContent** | [**CreateCephMonitorRequestContent**](CreateCephMonitorRequestContent.md) |  | 

### Return type

[**CreateCephMonitorResponseContent**](CreateCephMonitorResponseContent.md)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateClusterConfig

> CreateClusterConfigResponseContent CreateClusterConfig(ctx).CreateClusterConfigRequestContent(createClusterConfigRequestContent).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    createClusterConfigRequestContent := *openapiclient.NewCreateClusterConfigRequestContent("Clustername_example") // CreateClusterConfigRequestContent | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CreateClusterConfig(context.Background()).CreateClusterConfigRequestContent(createClusterConfigRequestContent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateClusterConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateClusterConfig`: CreateClusterConfigResponseContent
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateClusterConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateClusterConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createClusterConfigRequestContent** | [**CreateClusterConfigRequestContent**](CreateClusterConfigRequestContent.md) |  | 

### Return type

[**CreateClusterConfigResponseContent**](CreateClusterConfigResponseContent.md)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDirectory

> CreateDirectoryResponseContent CreateDirectory(ctx, node).CreateDirectoryRequestContent(createDirectoryRequestContent).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    node := "node_example" // string | 
    createDirectoryRequestContent := *openapiclient.NewCreateDirectoryRequestContent("Device_example", "Name_example") // CreateDirectoryRequestContent | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CreateDirectory(context.Background(), node).CreateDirectoryRequestContent(createDirectoryRequestContent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateDirectory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDirectory`: CreateDirectoryResponseContent
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateDirectory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDirectoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createDirectoryRequestContent** | [**CreateDirectoryRequestContent**](CreateDirectoryRequestContent.md) |  | 

### Return type

[**CreateDirectoryResponseContent**](CreateDirectoryResponseContent.md)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateGroup

> CreateGroup(ctx).CreateGroupRequestContent(createGroupRequestContent).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    createGroupRequestContent := *openapiclient.NewCreateGroupRequestContent("Groupid_example") // CreateGroupRequestContent | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultApi.CreateGroup(context.Background()).CreateGroupRequestContent(createGroupRequestContent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createGroupRequestContent** | [**CreateGroupRequestContent**](CreateGroupRequestContent.md) |  | 

### Return type

 (empty response body)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateLVM

> CreateLVMResponseContent CreateLVM(ctx, node).CreateLVMRequestContent(createLVMRequestContent).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    node := "node_example" // string | 
    createLVMRequestContent := *openapiclient.NewCreateLVMRequestContent("Device_example", "Name_example") // CreateLVMRequestContent | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CreateLVM(context.Background(), node).CreateLVMRequestContent(createLVMRequestContent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateLVM``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateLVM`: CreateLVMResponseContent
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateLVM`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateLVMRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createLVMRequestContent** | [**CreateLVMRequestContent**](CreateLVMRequestContent.md) |  | 

### Return type

[**CreateLVMResponseContent**](CreateLVMResponseContent.md)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateLVMThin

> CreateLVMThinResponseContent CreateLVMThin(ctx, node).CreateLVMThinRequestContent(createLVMThinRequestContent).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    node := "node_example" // string | 
    createLVMThinRequestContent := *openapiclient.NewCreateLVMThinRequestContent("Device_example", "Name_example") // CreateLVMThinRequestContent | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CreateLVMThin(context.Background(), node).CreateLVMThinRequestContent(createLVMThinRequestContent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateLVMThin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateLVMThin`: CreateLVMThinResponseContent
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateLVMThin`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateLVMThinRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createLVMThinRequestContent** | [**CreateLVMThinRequestContent**](CreateLVMThinRequestContent.md) |  | 

### Return type

[**CreateLVMThinResponseContent**](CreateLVMThinResponseContent.md)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNetworkInterface

> CreateNetworkInterface(ctx, node).CreateNetworkInterfaceRequestContent(createNetworkInterfaceRequestContent).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    node := "node_example" // string | 
    createNetworkInterfaceRequestContent := *openapiclient.NewCreateNetworkInterfaceRequestContent("Iface_example", openapiclient.NetworkInterfaceType("bridge")) // CreateNetworkInterfaceRequestContent | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultApi.CreateNetworkInterface(context.Background(), node).CreateNetworkInterfaceRequestContent(createNetworkInterfaceRequestContent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateNetworkInterface``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkInterfaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkInterfaceRequestContent** | [**CreateNetworkInterfaceRequestContent**](CreateNetworkInterfaceRequestContent.md) |  | 

### Return type

 (empty response body)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePool

> CreatePool(ctx).CreatePoolRequestContent(createPoolRequestContent).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    createPoolRequestContent := *openapiclient.NewCreatePoolRequestContent("Poolid_example") // CreatePoolRequestContent | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultApi.CreatePool(context.Background()).CreatePoolRequestContent(createPoolRequestContent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreatePool``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createPoolRequestContent** | [**CreatePoolRequestContent**](CreatePoolRequestContent.md) |  | 

### Return type

 (empty response body)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRealm

> CreateRealm(ctx).CreateRealmRequestContent(createRealmRequestContent).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    createRealmRequestContent := *openapiclient.NewCreateRealmRequestContent("Realm_example", openapiclient.RealmType("ad")) // CreateRealmRequestContent | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultApi.CreateRealm(context.Background()).CreateRealmRequestContent(createRealmRequestContent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateRealm``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRealmRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createRealmRequestContent** | [**CreateRealmRequestContent**](CreateRealmRequestContent.md) |  | 

### Return type

 (empty response body)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRole

> CreateRole(ctx).CreateRoleRequestContent(createRoleRequestContent).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    createRoleRequestContent := *openapiclient.NewCreateRoleRequestContent("Roleid_example") // CreateRoleRequestContent | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultApi.CreateRole(context.Background()).CreateRoleRequestContent(createRoleRequestContent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createRoleRequestContent** | [**CreateRoleRequestContent**](CreateRoleRequestContent.md) |  | 

### Return type

 (empty response body)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSnapshot

> CreateSnapshotResponseContent CreateSnapshot(ctx, node, vmId).CreateSnapshotRequestContent(createSnapshotRequestContent).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    node := "node_example" // string | 
    vmId := "vmId_example" // string | The id of the virtual machine as a string
    createSnapshotRequestContent := *openapiclient.NewCreateSnapshotRequestContent("Snapname_example") // CreateSnapshotRequestContent | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CreateSnapshot(context.Background(), node, vmId).CreateSnapshotRequestContent(createSnapshotRequestContent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateSnapshot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSnapshot`: CreateSnapshotResponseContent
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateSnapshot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** |  | 
**vmId** | **string** | The id of the virtual machine as a string | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createSnapshotRequestContent** | [**CreateSnapshotRequestContent**](CreateSnapshotRequestContent.md) |  | 

### Return type

[**CreateSnapshotResponseContent**](CreateSnapshotResponseContent.md)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateStorage

> CreateStorageResponseContent CreateStorage(ctx).CreateStorageRequestContent(createStorageRequestContent).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    createStorageRequestContent := *openapiclient.NewCreateStorageRequestContent("Storage_example", openapiclient.StorageType("zfspool")) // CreateStorageRequestContent | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CreateStorage(context.Background()).CreateStorageRequestContent(createStorageRequestContent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateStorage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateStorage`: CreateStorageResponseContent
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateStorage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateStorageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createStorageRequestContent** | [**CreateStorageRequestContent**](CreateStorageRequestContent.md) |  | 

### Return type

[**CreateStorageResponseContent**](CreateStorageResponseContent.md)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateStorageVolume

> CreateStorageVolumeResponseContent CreateStorageVolume(ctx, node, storage).CreateStorageVolumeRequestContent(createStorageVolumeRequestContent).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    node := "node_example" // string | 
    storage := "storage_example" // string | 
    createStorageVolumeRequestContent := *openapiclient.NewCreateStorageVolumeRequestContent("Filename_example", float32(123), "Vmid_example") // CreateStorageVolumeRequestContent | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CreateStorageVolume(context.Background(), node, storage).CreateStorageVolumeRequestContent(createStorageVolumeRequestContent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateStorageVolume``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateStorageVolume`: CreateStorageVolumeResponseContent
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateStorageVolume`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** |  | 
**storage** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateStorageVolumeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createStorageVolumeRequestContent** | [**CreateStorageVolumeRequestContent**](CreateStorageVolumeRequestContent.md) |  | 

### Return type

[**CreateStorageVolumeResponseContent**](CreateStorageVolumeResponseContent.md)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTicket

> CreateTicketResponseContent CreateTicket(ctx).CreateTicketRequestContent(createTicketRequestContent).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    createTicketRequestContent := *openapiclient.NewCreateTicketRequestContent("Username_example", "Password_example") // CreateTicketRequestContent | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CreateTicket(context.Background()).CreateTicketRequestContent(createTicketRequestContent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateTicket``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTicket`: CreateTicketResponseContent
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateTicket`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTicketRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createTicketRequestContent** | [**CreateTicketRequestContent**](CreateTicketRequestContent.md) |  | 

### Return type

[**CreateTicketResponseContent**](CreateTicketResponseContent.md)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateUser

> CreateUser(ctx).CreateUserRequestContent(createUserRequestContent).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    createUserRequestContent := *openapiclient.NewCreateUserRequestContent("Userid_example") // CreateUserRequestContent | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultApi.CreateUser(context.Background()).CreateUserRequestContent(createUserRequestContent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createUserRequestContent** | [**CreateUserRequestContent**](CreateUserRequestContent.md) |  | 

### Return type

 (empty response body)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateUserToken

> CreateUserTokenResponseContent CreateUserToken(ctx, userId, tokenId).CreateUserTokenRequestContent(createUserTokenRequestContent).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    userId := "userId_example" // string | 
    tokenId := "tokenId_example" // string | 
    createUserTokenRequestContent := *openapiclient.NewCreateUserTokenRequestContent() // CreateUserTokenRequestContent |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CreateUserToken(context.Background(), userId, tokenId).CreateUserTokenRequestContent(createUserTokenRequestContent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateUserToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateUserToken`: CreateUserTokenResponseContent
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateUserToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 
**tokenId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateUserTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createUserTokenRequestContent** | [**CreateUserTokenRequestContent**](CreateUserTokenRequestContent.md) |  | 

### Return type

[**CreateUserTokenResponseContent**](CreateUserTokenResponseContent.md)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateVirtualMachine

> CreateVirtualMachineResponseContent CreateVirtualMachine(ctx, node).CreateVirtualMachineRequestContent(createVirtualMachineRequestContent).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    node := "node_example" // string | 
    createVirtualMachineRequestContent := *openapiclient.NewCreateVirtualMachineRequestContent("Vmid_example") // CreateVirtualMachineRequestContent | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CreateVirtualMachine(context.Background(), node).CreateVirtualMachineRequestContent(createVirtualMachineRequestContent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateVirtualMachine``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVirtualMachine`: CreateVirtualMachineResponseContent
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateVirtualMachine`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVirtualMachineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createVirtualMachineRequestContent** | [**CreateVirtualMachineRequestContent**](CreateVirtualMachineRequestContent.md) |  | 

### Return type

[**CreateVirtualMachineResponseContent**](CreateVirtualMachineResponseContent.md)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateVirtualMachineTemplate

> CreateVirtualMachineTemplateResponseContent CreateVirtualMachineTemplate(ctx, node, vmId).CreateVirtualMachineTemplateRequestContent(createVirtualMachineTemplateRequestContent).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    node := "node_example" // string | 
    vmId := "vmId_example" // string | The id of the virtual machine as a string
    createVirtualMachineTemplateRequestContent := *openapiclient.NewCreateVirtualMachineTemplateRequestContent() // CreateVirtualMachineTemplateRequestContent |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CreateVirtualMachineTemplate(context.Background(), node, vmId).CreateVirtualMachineTemplateRequestContent(createVirtualMachineTemplateRequestContent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateVirtualMachineTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVirtualMachineTemplate`: CreateVirtualMachineTemplateResponseContent
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateVirtualMachineTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** |  | 
**vmId** | **string** | The id of the virtual machine as a string | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVirtualMachineTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createVirtualMachineTemplateRequestContent** | [**CreateVirtualMachineTemplateRequestContent**](CreateVirtualMachineTemplateRequestContent.md) |  | 

### Return type

[**CreateVirtualMachineTemplateResponseContent**](CreateVirtualMachineTemplateResponseContent.md)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateZFSPool

> CreateZFSPoolResponseContent CreateZFSPool(ctx, node).CreateZFSPoolRequestContent(createZFSPoolRequestContent).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    node := "node_example" // string | 
    createZFSPoolRequestContent := *openapiclient.NewCreateZFSPoolRequestContent("Devices_example", "Name_example", openapiclient.ZFSRaidLevel("single")) // CreateZFSPoolRequestContent | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CreateZFSPool(context.Background(), node).CreateZFSPoolRequestContent(createZFSPoolRequestContent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateZFSPool``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateZFSPool`: CreateZFSPoolResponseContent
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateZFSPool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateZFSPoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createZFSPoolRequestContent** | [**CreateZFSPoolRequestContent**](CreateZFSPoolRequestContent.md) |  | 

### Return type

[**CreateZFSPoolResponseContent**](CreateZFSPoolResponseContent.md)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCephManager

> DeleteCephManagerResponseContent DeleteCephManager(ctx, node, manager).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    node := "node_example" // string | 
    manager := "manager_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.DeleteCephManager(context.Background(), node, manager).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteCephManager``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteCephManager`: DeleteCephManagerResponseContent
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DeleteCephManager`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** |  | 
**manager** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCephManagerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DeleteCephManagerResponseContent**](DeleteCephManagerResponseContent.md)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCephMonitor

> DeleteCephMonitorResponseContent DeleteCephMonitor(ctx, node, monitor).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    node := "node_example" // string | 
    monitor := "monitor_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.DeleteCephMonitor(context.Background(), node, monitor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteCephMonitor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteCephMonitor`: DeleteCephMonitorResponseContent
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DeleteCephMonitor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** |  | 
**monitor** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCephMonitorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DeleteCephMonitorResponseContent**](DeleteCephMonitorResponseContent.md)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDirectory

> DeleteDirectoryResponseContent DeleteDirectory(ctx, node, name).CleanupConfig(cleanupConfig).CleanupDisks(cleanupDisks).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    node := "node_example" // string | 
    name := "name_example" // string | The storage identifier.
    cleanupConfig := float32(8.14) // float32 | Marks the associated storage as not available on this node anr removes them from the config. Takes a boolean integer value (0 false, 1 true). (optional)
    cleanupDisks := float32(8.14) // float32 | Wipes the disk. Takes a boolean integer value (0 false, 1 true). (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.DeleteDirectory(context.Background(), node, name).CleanupConfig(cleanupConfig).CleanupDisks(cleanupDisks).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteDirectory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteDirectory`: DeleteDirectoryResponseContent
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DeleteDirectory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** |  | 
**name** | **string** | The storage identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDirectoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cleanupConfig** | **float32** | Marks the associated storage as not available on this node anr removes them from the config. Takes a boolean integer value (0 false, 1 true). | 
 **cleanupDisks** | **float32** | Wipes the disk. Takes a boolean integer value (0 false, 1 true). | 

### Return type

[**DeleteDirectoryResponseContent**](DeleteDirectoryResponseContent.md)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteGroup

> DeleteGroup(ctx, groupId).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    groupId := "groupId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultApi.DeleteGroup(context.Background(), groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteLVM

> DeleteLVMResponseContent DeleteLVM(ctx, node, name).CleanupConfig(cleanupConfig).CleanupDisks(cleanupDisks).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    node := "node_example" // string | 
    name := "name_example" // string | The storage identifier.
    cleanupConfig := float32(8.14) // float32 | Marks the associated storage as not available on this node anr removes them from the config. Takes a boolean integer value (0 false, 1 true). (optional)
    cleanupDisks := float32(8.14) // float32 | Wipes the disk. Takes a boolean integer value (0 false, 1 true). (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.DeleteLVM(context.Background(), node, name).CleanupConfig(cleanupConfig).CleanupDisks(cleanupDisks).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteLVM``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteLVM`: DeleteLVMResponseContent
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DeleteLVM`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** |  | 
**name** | **string** | The storage identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLVMRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cleanupConfig** | **float32** | Marks the associated storage as not available on this node anr removes them from the config. Takes a boolean integer value (0 false, 1 true). | 
 **cleanupDisks** | **float32** | Wipes the disk. Takes a boolean integer value (0 false, 1 true). | 

### Return type

[**DeleteLVMResponseContent**](DeleteLVMResponseContent.md)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteLVMThin

> DeleteLVMThinResponseContent DeleteLVMThin(ctx, node, name).VolumeGroup(volumeGroup).CleanupConfig(cleanupConfig).CleanupDisks(cleanupDisks).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    node := "node_example" // string | 
    name := "name_example" // string | The storage identifier.
    volumeGroup := "volumeGroup_example" // string | The volume group name.
    cleanupConfig := float32(8.14) // float32 | Marks the associated storage as not available on this node anr removes them from the config. Takes a boolean integer value (0 false, 1 true). (optional)
    cleanupDisks := float32(8.14) // float32 | Wipes the disk. Takes a boolean integer value (0 false, 1 true). (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.DeleteLVMThin(context.Background(), node, name).VolumeGroup(volumeGroup).CleanupConfig(cleanupConfig).CleanupDisks(cleanupDisks).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteLVMThin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteLVMThin`: DeleteLVMThinResponseContent
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DeleteLVMThin`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** |  | 
**name** | **string** | The storage identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLVMThinRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **volumeGroup** | **string** | The volume group name. | 
 **cleanupConfig** | **float32** | Marks the associated storage as not available on this node anr removes them from the config. Takes a boolean integer value (0 false, 1 true). | 
 **cleanupDisks** | **float32** | Wipes the disk. Takes a boolean integer value (0 false, 1 true). | 

### Return type

[**DeleteLVMThinResponseContent**](DeleteLVMThinResponseContent.md)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetworkInterface

> DeleteNetworkInterface(ctx, node, interface_).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    node := "node_example" // string | 
    interface_ := "interface__example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultApi.DeleteNetworkInterface(context.Background(), node, interface_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteNetworkInterface``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** |  | 
**interface_** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkInterfaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNodeCertificate

> DeleteNodeCertificateResponseContent DeleteNodeCertificate(ctx, node).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    node := "node_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.DeleteNodeCertificate(context.Background(), node).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteNodeCertificate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteNodeCertificate`: DeleteNodeCertificateResponseContent
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DeleteNodeCertificate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNodeCertificateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteNodeCertificateResponseContent**](DeleteNodeCertificateResponseContent.md)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePool

> DeletePool(ctx, poolId).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    poolId := "poolId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultApi.DeletePool(context.Background(), poolId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeletePool``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**poolId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRealm

> DeleteRealm(ctx, realm).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    realm := "realm_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultApi.DeleteRealm(context.Background(), realm).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteRealm``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRealmRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRole

> DeleteRole(ctx, roleId).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    roleId := "roleId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultApi.DeleteRole(context.Background(), roleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteStorage

> DeleteStorage(ctx, storage).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    storage := "storage_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultApi.DeleteStorage(context.Background(), storage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteStorage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storage** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteStorageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteStorageVolume

> DeleteStorageVolumeResponseContent DeleteStorageVolume(ctx, node, storage, volume).Delay(delay).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    node := "node_example" // string | 
    storage := "storage_example" // string | 
    volume := "volume_example" // string | 
    delay := float32(8.14) // float32 | Delay in seconds to wait for task to finish. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.DeleteStorageVolume(context.Background(), node, storage, volume).Delay(delay).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteStorageVolume``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteStorageVolume`: DeleteStorageVolumeResponseContent
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DeleteStorageVolume`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** |  | 
**storage** | **string** |  | 
**volume** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteStorageVolumeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **delay** | **float32** | Delay in seconds to wait for task to finish. | 

### Return type

[**DeleteStorageVolumeResponseContent**](DeleteStorageVolumeResponseContent.md)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUser

> DeleteUser(ctx, userId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    userId := "userId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultApi.DeleteUser(context.Background(), userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUserToken

> DeleteUserToken(ctx, userId, tokenId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    userId := "userId_example" // string | 
    tokenId := "tokenId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultApi.DeleteUserToken(context.Background(), userId, tokenId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteUserToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 
**tokenId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteVirtualMachine

> DeleteVirtualMachineResponseContent DeleteVirtualMachine(ctx, node, vmId).DestoryUnreferencedDisks(destoryUnreferencedDisks).Purge(purge).Skiplock(skiplock).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    node := "node_example" // string | 
    vmId := "vmId_example" // string | The id of the virtual machine as a string
    destoryUnreferencedDisks := float32(8.14) // float32 | Destroy disks that are not referenced in the config. (optional)
    purge := float32(8.14) // float32 | Purge the VM from the configurations, backups, jobs, and HA. (optional)
    skiplock := float32(8.14) // float32 | Skip the lock check. Only valid for root. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.DeleteVirtualMachine(context.Background(), node, vmId).DestoryUnreferencedDisks(destoryUnreferencedDisks).Purge(purge).Skiplock(skiplock).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteVirtualMachine``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteVirtualMachine`: DeleteVirtualMachineResponseContent
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DeleteVirtualMachine`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** |  | 
**vmId** | **string** | The id of the virtual machine as a string | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVirtualMachineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **destoryUnreferencedDisks** | **float32** | Destroy disks that are not referenced in the config. | 
 **purge** | **float32** | Purge the VM from the configurations, backups, jobs, and HA. | 
 **skiplock** | **float32** | Skip the lock check. Only valid for root. | 

### Return type

[**DeleteVirtualMachineResponseContent**](DeleteVirtualMachineResponseContent.md)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteZFSPool

> DeleteZFSPoolResponseContent DeleteZFSPool(ctx, node, name).CleanupConfig(cleanupConfig).CleanupDisks(cleanupDisks).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    node := "node_example" // string | 
    name := "name_example" // string | The storage identifier.
    cleanupConfig := float32(8.14) // float32 | Marks the associated storage as not available on this node anr removes them from the config. Takes a boolean integer value (0 false, 1 true). (optional)
    cleanupDisks := float32(8.14) // float32 | Wipes the disk. Takes a boolean integer value (0 false, 1 true). (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.DeleteZFSPool(context.Background(), node, name).CleanupConfig(cleanupConfig).CleanupDisks(cleanupDisks).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteZFSPool``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteZFSPool`: DeleteZFSPoolResponseContent
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DeleteZFSPool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** |  | 
**name** | **string** | The storage identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteZFSPoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cleanupConfig** | **float32** | Marks the associated storage as not available on this node anr removes them from the config. Takes a boolean integer value (0 false, 1 true). | 
 **cleanupDisks** | **float32** | Wipes the disk. Takes a boolean integer value (0 false, 1 true). | 

### Return type

[**DeleteZFSPoolResponseContent**](DeleteZFSPoolResponseContent.md)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadFromUrlToStorage

> DownloadFromUrlToStorageResponseContent DownloadFromUrlToStorage(ctx, node, storage).DownloadFromUrlToStorageRequestContent(downloadFromUrlToStorageRequestContent).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    node := "node_example" // string | 
    storage := "storage_example" // string | 
    downloadFromUrlToStorageRequestContent := *openapiclient.NewDownloadFromUrlToStorageRequestContent(openapiclient.UploadContentType("iso"), "Filename_example", "Url_example") // DownloadFromUrlToStorageRequestContent | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.DownloadFromUrlToStorage(context.Background(), node, storage).DownloadFromUrlToStorageRequestContent(downloadFromUrlToStorageRequestContent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DownloadFromUrlToStorage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DownloadFromUrlToStorage`: DownloadFromUrlToStorageResponseContent
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DownloadFromUrlToStorage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** |  | 
**storage** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadFromUrlToStorageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **downloadFromUrlToStorageRequestContent** | [**DownloadFromUrlToStorageRequestContent**](DownloadFromUrlToStorageRequestContent.md) |  | 

### Return type

[**DownloadFromUrlToStorageResponseContent**](DownloadFromUrlToStorageResponseContent.md)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccessControlList

> GetAccessControlListResponseContent GetAccessControlList(ctx).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetAccessControlList(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetAccessControlList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccessControlList`: GetAccessControlListResponseContent
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetAccessControlList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccessControlListRequest struct via the builder pattern


### Return type

[**GetAccessControlListResponseContent**](GetAccessControlListResponseContent.md)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClusterApiVersion

> GetClusterApiVersionResponseContent GetClusterApiVersion(ctx).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetClusterApiVersion(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetClusterApiVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetClusterApiVersion`: GetClusterApiVersionResponseContent
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetClusterApiVersion`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetClusterApiVersionRequest struct via the builder pattern


### Return type

[**GetClusterApiVersionResponseContent**](GetClusterApiVersionResponseContent.md)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClusterJoinInformation

> GetClusterJoinInformationResponseContent GetClusterJoinInformation(ctx).Node(node).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    node := "node_example" // string | The node which to join. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetClusterJoinInformation(context.Background()).Node(node).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetClusterJoinInformation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetClusterJoinInformation`: GetClusterJoinInformationResponseContent
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetClusterJoinInformation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetClusterJoinInformationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **node** | **string** | The node which to join. | 

### Return type

[**GetClusterJoinInformationResponseContent**](GetClusterJoinInformationResponseContent.md)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClusterTotemSettings

> GetClusterTotemSettingsResponseContent GetClusterTotemSettings(ctx).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetClusterTotemSettings(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetClusterTotemSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetClusterTotemSettings`: GetClusterTotemSettingsResponseContent
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetClusterTotemSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetClusterTotemSettingsRequest struct via the builder pattern


### Return type

[**GetClusterTotemSettingsResponseContent**](GetClusterTotemSettingsResponseContent.md)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroupDetails

> GetGroupDetailsResponseContent GetGroupDetails(ctx, groupId).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    groupId := "groupId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetGroupDetails(context.Background(), groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetGroupDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGroupDetails`: GetGroupDetailsResponseContent
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetGroupDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetGroupDetailsResponseContent**](GetGroupDetailsResponseContent.md)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkInterface

> GetNetworkInterfaceResponseContent GetNetworkInterface(ctx, node, interface_).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    node := "node_example" // string | 
    interface_ := "interface__example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetNetworkInterface(context.Background(), node, interface_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetNetworkInterface``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkInterface`: GetNetworkInterfaceResponseContent
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetNetworkInterface`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** |  | 
**interface_** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkInterfaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetNetworkInterfaceResponseContent**](GetNetworkInterfaceResponseContent.md)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNodeDns

> GetNodeDnsResponseContent GetNodeDns(ctx, node).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    node := "node_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetNodeDns(context.Background(), node).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetNodeDns``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNodeDns`: GetNodeDnsResponseContent
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetNodeDns`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNodeDnsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetNodeDnsResponseContent**](GetNodeDnsResponseContent.md)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPackageChangelog

> GetPackageChangelogResponseContent GetPackageChangelog(ctx, node).Name(name).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    node := "node_example" // string | 
    name := "name_example" // string | The name of the package to get the changelog for. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetPackageChangelog(context.Background(), node).Name(name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetPackageChangelog``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPackageChangelog`: GetPackageChangelogResponseContent
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetPackageChangelog`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPackageChangelogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **string** | The name of the package to get the changelog for. | 

### Return type

[**GetPackageChangelogResponseContent**](GetPackageChangelogResponseContent.md)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPendingVirtualMachineCloudInitChanges

> GetPendingVirtualMachineCloudInitChangesResponseContent GetPendingVirtualMachineCloudInitChanges(ctx, node, vmId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    node := "node_example" // string | 
    vmId := "vmId_example" // string | The id of the virtual machine as a string

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetPendingVirtualMachineCloudInitChanges(context.Background(), node, vmId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetPendingVirtualMachineCloudInitChanges``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPendingVirtualMachineCloudInitChanges`: GetPendingVirtualMachineCloudInitChangesResponseContent
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetPendingVirtualMachineCloudInitChanges`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** |  | 
**vmId** | **string** | The id of the virtual machine as a string | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPendingVirtualMachineCloudInitChangesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetPendingVirtualMachineCloudInitChangesResponseContent**](GetPendingVirtualMachineCloudInitChangesResponseContent.md)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPermissions

> GetPermissionsResponseContent GetPermissions(ctx).Path(path).Userid(userid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    path := "path_example" // string |  (optional)
    userid := "userid_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetPermissions(context.Background()).Path(path).Userid(userid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetPermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPermissions`: GetPermissionsResponseContent
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetPermissions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string** |  | 
 **userid** | **string** |  | 

### Return type

[**GetPermissionsResponseContent**](GetPermissionsResponseContent.md)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPool

> GetPoolResponseContent GetPool(ctx, poolId).Type_(type_).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    poolId := "poolId_example" // string | 
    type_ := openapiclient.PoolMemberType("qemu") // PoolMemberType |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetPool(context.Background(), poolId).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetPool``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPool`: GetPoolResponseContent
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetPool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**poolId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | [**PoolMemberType**](PoolMemberType.md) |  | 

### Return type

[**GetPoolResponseContent**](GetPoolResponseContent.md)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRealm

> GetRealmResponseContent GetRealm(ctx, realm).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    realm := "realm_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetRealm(context.Background(), realm).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetRealm``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRealm`: GetRealmResponseContent
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetRealm`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRealmRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetRealmResponseContent**](GetRealmResponseContent.md)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRole

> GetRoleResponseContent GetRole(ctx, roleId).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    roleId := "roleId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetRole(context.Background(), roleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRole`: GetRoleResponseContent
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetRoleResponseContent**](GetRoleResponseContent.md)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSmartHealth

> GetSmartHealthResponseContent GetSmartHealth(ctx, node).Disk(disk).Healthonly(healthonly).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    node := "node_example" // string | 
    disk := "disk_example" // string | Disk to check. (optional)
    healthonly := float32(8.14) // float32 | Return only health status. Takes a boolean integer value (0 false, 1 true). (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetSmartHealth(context.Background(), node).Disk(disk).Healthonly(healthonly).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetSmartHealth``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSmartHealth`: GetSmartHealthResponseContent
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetSmartHealth`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSmartHealthRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **disk** | **string** | Disk to check. | 
 **healthonly** | **float32** | Return only health status. Takes a boolean integer value (0 false, 1 true). | 

### Return type

[**GetSmartHealthResponseContent**](GetSmartHealthResponseContent.md)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStorage

> GetStorageResponseContent GetStorage(ctx, storage).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    storage := "storage_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetStorage(context.Background(), storage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetStorage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStorage`: GetStorageResponseContent
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetStorage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storage** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStorageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetStorageResponseContent**](GetStorageResponseContent.md)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStorageStatus

> GetStorageStatusResponseContent GetStorageStatus(ctx, node, storage).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    node := "node_example" // string | 
    storage := "storage_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetStorageStatus(context.Background(), node, storage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetStorageStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStorageStatus`: GetStorageStatusResponseContent
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetStorageStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** |  | 
**storage** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStorageStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetStorageStatusResponseContent**](GetStorageStatusResponseContent.md)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserConfiguration

> GetUserConfigurationResponseContent GetUserConfiguration(ctx, userId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    userId := "userId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetUserConfiguration(context.Background(), userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetUserConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserConfiguration`: GetUserConfigurationResponseContent
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetUserConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetUserConfigurationResponseContent**](GetUserConfigurationResponseContent.md)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserToken

> GetUserTokenResponseContent GetUserToken(ctx, userId, tokenId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    userId := "userId_example" // string | 
    tokenId := "tokenId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetUserToken(context.Background(), userId, tokenId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetUserToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserToken`: GetUserTokenResponseContent
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetUserToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 
**tokenId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetUserTokenResponseContent**](GetUserTokenResponseContent.md)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVersion

> GetVersionResponseContent GetVersion(ctx).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetVersion(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVersion`: GetVersionResponseContent
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetVersion`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetVersionRequest struct via the builder pattern


### Return type

[**GetVersionResponseContent**](GetVersionResponseContent.md)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVirtualMachineAgentInfo

> GetVirtualMachineAgentInfoResponseContent GetVirtualMachineAgentInfo(ctx, node, vmId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    node := "node_example" // string | 
    vmId := "vmId_example" // string | The id of the virtual machine as a string

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetVirtualMachineAgentInfo(context.Background(), node, vmId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetVirtualMachineAgentInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVirtualMachineAgentInfo`: GetVirtualMachineAgentInfoResponseContent
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetVirtualMachineAgentInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** |  | 
**vmId** | **string** | The id of the virtual machine as a string | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVirtualMachineAgentInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetVirtualMachineAgentInfoResponseContent**](GetVirtualMachineAgentInfoResponseContent.md)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVirtualMachineCloudInit

> GetVirtualMachineCloudInitResponseContent GetVirtualMachineCloudInit(ctx, node, vmId).Type_(type_).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    node := "node_example" // string | 
    vmId := "vmId_example" // string | The id of the virtual machine as a string
    type_ := openapiclient.CloudInitType("user") // CloudInitType | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetVirtualMachineCloudInit(context.Background(), node, vmId).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetVirtualMachineCloudInit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVirtualMachineCloudInit`: GetVirtualMachineCloudInitResponseContent
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetVirtualMachineCloudInit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** |  | 
**vmId** | **string** | The id of the virtual machine as a string | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVirtualMachineCloudInitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **type_** | [**CloudInitType**](CloudInitType.md) |  | 

### Return type

[**GetVirtualMachineCloudInitResponseContent**](GetVirtualMachineCloudInitResponseContent.md)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVirtualMachineCommandStatus

> GetVirtualMachineCommandStatusResponseContent GetVirtualMachineCommandStatus(ctx, node, vmId).Pid(pid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    node := "node_example" // string | 
    vmId := "vmId_example" // string | The id of the virtual machine as a string
    pid := float32(8.14) // float32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetVirtualMachineCommandStatus(context.Background(), node, vmId).Pid(pid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetVirtualMachineCommandStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVirtualMachineCommandStatus`: GetVirtualMachineCommandStatusResponseContent
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetVirtualMachineCommandStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** |  | 
**vmId** | **string** | The id of the virtual machine as a string | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVirtualMachineCommandStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pid** | **float32** |  | 

### Return type

[**GetVirtualMachineCommandStatusResponseContent**](GetVirtualMachineCommandStatusResponseContent.md)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVirtualMachineConfiguration

> GetVirtualMachineConfigurationResponseContent GetVirtualMachineConfiguration(ctx, node, vmId).Current(current).Snapshot(snapshot).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    node := "node_example" // string | 
    vmId := "vmId_example" // string | The id of the virtual machine as a string
    current := float32(8.14) // float32 | If specified, the configuration returned will be the current configuration. Otherwise, the configuration returned will be the pending configuration. (optional)
    snapshot := "snapshot_example" // string | Fetch the configuration from the specified snapshot. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetVirtualMachineConfiguration(context.Background(), node, vmId).Current(current).Snapshot(snapshot).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetVirtualMachineConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVirtualMachineConfiguration`: GetVirtualMachineConfigurationResponseContent
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetVirtualMachineConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** |  | 
**vmId** | **string** | The id of the virtual machine as a string | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVirtualMachineConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **current** | **float32** | If specified, the configuration returned will be the current configuration. Otherwise, the configuration returned will be the pending configuration. | 
 **snapshot** | **string** | Fetch the configuration from the specified snapshot. | 

### Return type

[**GetVirtualMachineConfigurationResponseContent**](GetVirtualMachineConfigurationResponseContent.md)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVirtualMachineFeatureSupport

> GetVirtualMachineFeatureSupportResponseContent GetVirtualMachineFeatureSupport(ctx, node, vmId).Feature(feature).Snapname(snapname).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    node := "node_example" // string | 
    vmId := "vmId_example" // string | The id of the virtual machine as a string
    feature := "feature_example" // string | 
    snapname := "snapname_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetVirtualMachineFeatureSupport(context.Background(), node, vmId).Feature(feature).Snapname(snapname).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetVirtualMachineFeatureSupport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVirtualMachineFeatureSupport`: GetVirtualMachineFeatureSupportResponseContent
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetVirtualMachineFeatureSupport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** |  | 
**vmId** | **string** | The id of the virtual machine as a string | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVirtualMachineFeatureSupportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **feature** | **string** |  | 
 **snapname** | **string** |  | 

### Return type

[**GetVirtualMachineFeatureSupportResponseContent**](GetVirtualMachineFeatureSupportResponseContent.md)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVirtualMachineFileSystemInformation

> GetVirtualMachineFileSystemInformationResponseContent GetVirtualMachineFileSystemInformation(ctx, node, vmId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    node := "node_example" // string | 
    vmId := "vmId_example" // string | The id of the virtual machine as a string

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetVirtualMachineFileSystemInformation(context.Background(), node, vmId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetVirtualMachineFileSystemInformation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVirtualMachineFileSystemInformation`: GetVirtualMachineFileSystemInformationResponseContent
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetVirtualMachineFileSystemInformation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** |  | 
**vmId** | **string** | The id of the virtual machine as a string | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVirtualMachineFileSystemInformationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetVirtualMachineFileSystemInformationResponseContent**](GetVirtualMachineFileSystemInformationResponseContent.md)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVirtualMachineHostname

> GetVirtualMachineHostnameResponseContent GetVirtualMachineHostname(ctx, node, vmId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    node := "node_example" // string | 
    vmId := "vmId_example" // string | The id of the virtual machine as a string

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetVirtualMachineHostname(context.Background(), node, vmId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetVirtualMachineHostname``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVirtualMachineHostname`: GetVirtualMachineHostnameResponseContent
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetVirtualMachineHostname`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** |  | 
**vmId** | **string** | The id of the virtual machine as a string | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVirtualMachineHostnameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetVirtualMachineHostnameResponseContent**](GetVirtualMachineHostnameResponseContent.md)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVirtualMachineMemoryBlockInformation

> GetVirtualMachineMemoryBlockInformationResponseContent GetVirtualMachineMemoryBlockInformation(ctx, node, vmId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    node := "node_example" // string | 
    vmId := "vmId_example" // string | The id of the virtual machine as a string

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetVirtualMachineMemoryBlockInformation(context.Background(), node, vmId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetVirtualMachineMemoryBlockInformation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVirtualMachineMemoryBlockInformation`: GetVirtualMachineMemoryBlockInformationResponseContent
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetVirtualMachineMemoryBlockInformation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** |  | 
**vmId** | **string** | The id of the virtual machine as a string | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVirtualMachineMemoryBlockInformationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetVirtualMachineMemoryBlockInformationResponseContent**](GetVirtualMachineMemoryBlockInformationResponseContent.md)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVirtualMachineMemoryBlocks

> GetVirtualMachineMemoryBlocksResponseContent GetVirtualMachineMemoryBlocks(ctx, node, vmId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    node := "node_example" // string | 
    vmId := "vmId_example" // string | The id of the virtual machine as a string

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetVirtualMachineMemoryBlocks(context.Background(), node, vmId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetVirtualMachineMemoryBlocks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVirtualMachineMemoryBlocks`: GetVirtualMachineMemoryBlocksResponseContent
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetVirtualMachineMemoryBlocks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** |  | 
**vmId** | **string** | The id of the virtual machine as a string | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVirtualMachineMemoryBlocksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetVirtualMachineMemoryBlocksResponseContent**](GetVirtualMachineMemoryBlocksResponseContent.md)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVirtualMachineNetworkInterfaces

> GetVirtualMachineNetworkInterfacesResponseContent GetVirtualMachineNetworkInterfaces(ctx, node, vmId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    node := "node_example" // string | 
    vmId := "vmId_example" // string | The id of the virtual machine as a string

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetVirtualMachineNetworkInterfaces(context.Background(), node, vmId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetVirtualMachineNetworkInterfaces``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVirtualMachineNetworkInterfaces`: GetVirtualMachineNetworkInterfacesResponseContent
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetVirtualMachineNetworkInterfaces`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** |  | 
**vmId** | **string** | The id of the virtual machine as a string | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVirtualMachineNetworkInterfacesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetVirtualMachineNetworkInterfacesResponseContent**](GetVirtualMachineNetworkInterfacesResponseContent.md)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVirtualMachineOperatingSystemInformation

> GetVirtualMachineOperatingSystemInformationResponseContent GetVirtualMachineOperatingSystemInformation(ctx, node, vmId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    node := "node_example" // string | 
    vmId := "vmId_example" // string | The id of the virtual machine as a string

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetVirtualMachineOperatingSystemInformation(context.Background(), node, vmId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetVirtualMachineOperatingSystemInformation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVirtualMachineOperatingSystemInformation`: GetVirtualMachineOperatingSystemInformationResponseContent
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetVirtualMachineOperatingSystemInformation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** |  | 
**vmId** | **string** | The id of the virtual machine as a string | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVirtualMachineOperatingSystemInformationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetVirtualMachineOperatingSystemInformationResponseContent**](GetVirtualMachineOperatingSystemInformationResponseContent.md)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVirtualMachineStatus

> GetVirtualMachineStatusResponseContent GetVirtualMachineStatus(ctx, node, vmId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    node := "node_example" // string | 
    vmId := "vmId_example" // string | The id of the virtual machine as a string

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetVirtualMachineStatus(context.Background(), node, vmId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetVirtualMachineStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVirtualMachineStatus`: GetVirtualMachineStatusResponseContent
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetVirtualMachineStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** |  | 
**vmId** | **string** | The id of the virtual machine as a string | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVirtualMachineStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetVirtualMachineStatusResponseContent**](GetVirtualMachineStatusResponseContent.md)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVirtualMachineTime

> GetVirtualMachineTimeResponseContent GetVirtualMachineTime(ctx, node, vmId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    node := "node_example" // string | 
    vmId := "vmId_example" // string | The id of the virtual machine as a string

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetVirtualMachineTime(context.Background(), node, vmId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetVirtualMachineTime``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVirtualMachineTime`: GetVirtualMachineTimeResponseContent
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetVirtualMachineTime`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** |  | 
**vmId** | **string** | The id of the virtual machine as a string | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVirtualMachineTimeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetVirtualMachineTimeResponseContent**](GetVirtualMachineTimeResponseContent.md)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVirtualMachineTimeZone

> GetVirtualMachineTimeZoneResponseContent GetVirtualMachineTimeZone(ctx, node, vmId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    node := "node_example" // string | 
    vmId := "vmId_example" // string | The id of the virtual machine as a string

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetVirtualMachineTimeZone(context.Background(), node, vmId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetVirtualMachineTimeZone``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVirtualMachineTimeZone`: GetVirtualMachineTimeZoneResponseContent
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetVirtualMachineTimeZone`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** |  | 
**vmId** | **string** | The id of the virtual machine as a string | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVirtualMachineTimeZoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetVirtualMachineTimeZoneResponseContent**](GetVirtualMachineTimeZoneResponseContent.md)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVirtualMachineVcpus

> GetVirtualMachineVcpusResponseContent GetVirtualMachineVcpus(ctx, node, vmId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    node := "node_example" // string | 
    vmId := "vmId_example" // string | The id of the virtual machine as a string

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetVirtualMachineVcpus(context.Background(), node, vmId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetVirtualMachineVcpus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVirtualMachineVcpus`: GetVirtualMachineVcpusResponseContent
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetVirtualMachineVcpus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** |  | 
**vmId** | **string** | The id of the virtual machine as a string | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVirtualMachineVcpusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetVirtualMachineVcpusResponseContent**](GetVirtualMachineVcpusResponseContent.md)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetZFSPoolStatus

> GetZFSPoolStatusResponseContent GetZFSPoolStatus(ctx, node, name).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    node := "node_example" // string | 
    name := "name_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetZFSPoolStatus(context.Background(), node, name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetZFSPoolStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetZFSPoolStatus`: GetZFSPoolStatusResponseContent
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetZFSPoolStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** |  | 
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetZFSPoolStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetZFSPoolStatusResponseContent**](GetZFSPoolStatusResponseContent.md)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InitCephCluster

> InitCephCluster(ctx, node).InitCephClusterRequestContent(initCephClusterRequestContent).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    node := "node_example" // string | 
    initCephClusterRequestContent := *openapiclient.NewInitCephClusterRequestContent() // InitCephClusterRequestContent |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultApi.InitCephCluster(context.Background(), node).InitCephClusterRequestContent(initCephClusterRequestContent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.InitCephCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiInitCephClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **initCephClusterRequestContent** | [**InitCephClusterRequestContent**](InitCephClusterRequestContent.md) |  | 

### Return type

 (empty response body)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InitializeGPT

> InitializeGPTResponseContent InitializeGPT(ctx, node).InitializeGPTRequestContent(initializeGPTRequestContent).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    node := "node_example" // string | 
    initializeGPTRequestContent := *openapiclient.NewInitializeGPTRequestContent() // InitializeGPTRequestContent |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.InitializeGPT(context.Background(), node).InitializeGPTRequestContent(initializeGPTRequestContent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.InitializeGPT``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InitializeGPT`: InitializeGPTResponseContent
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.InitializeGPT`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiInitializeGPTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **initializeGPTRequestContent** | [**InitializeGPTRequestContent**](InitializeGPTRequestContent.md) |  | 

### Return type

[**InitializeGPTResponseContent**](InitializeGPTResponseContent.md)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## JoinCluster

> JoinClusterResponseContent JoinCluster(ctx).JoinClusterRequestContent(joinClusterRequestContent).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    joinClusterRequestContent := *openapiclient.NewJoinClusterRequestContent("Fingerprint_example", "Hostname_example", "Password_example") // JoinClusterRequestContent | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.JoinCluster(context.Background()).JoinClusterRequestContent(joinClusterRequestContent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.JoinCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `JoinCluster`: JoinClusterResponseContent
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.JoinCluster`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiJoinClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **joinClusterRequestContent** | [**JoinClusterRequestContent**](JoinClusterRequestContent.md) |  | 

### Return type

[**JoinClusterResponseContent**](JoinClusterResponseContent.md)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCephManagers

> ListCephManagersResponseContent ListCephManagers(ctx, node).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    node := "node_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListCephManagers(context.Background(), node).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListCephManagers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCephManagers`: ListCephManagersResponseContent
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListCephManagers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListCephManagersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListCephManagersResponseContent**](ListCephManagersResponseContent.md)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCephMonitors

> ListCephMonitorsResponseContent ListCephMonitors(ctx, node).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    node := "node_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListCephMonitors(context.Background(), node).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListCephMonitors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCephMonitors`: ListCephMonitorsResponseContent
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListCephMonitors`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListCephMonitorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListCephMonitorsResponseContent**](ListCephMonitorsResponseContent.md)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCorosyncNodes

> ListCorosyncNodesResponseContent ListCorosyncNodes(ctx).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListCorosyncNodes(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListCorosyncNodes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCorosyncNodes`: ListCorosyncNodesResponseContent
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListCorosyncNodes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListCorosyncNodesRequest struct via the builder pattern


### Return type

[**ListCorosyncNodesResponseContent**](ListCorosyncNodesResponseContent.md)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCpuCapabilities

> ListCpuCapabilitiesResponseContent ListCpuCapabilities(ctx, node).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    node := "node_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListCpuCapabilities(context.Background(), node).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListCpuCapabilities``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCpuCapabilities`: ListCpuCapabilitiesResponseContent
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListCpuCapabilities`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListCpuCapabilitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListCpuCapabilitiesResponseContent**](ListCpuCapabilitiesResponseContent.md)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDirectories

> ListDirectoriesResponseContent ListDirectories(ctx, node).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    node := "node_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListDirectories(context.Background(), node).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListDirectories``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDirectories`: ListDirectoriesResponseContent
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListDirectories`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListDirectoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListDirectoriesResponseContent**](ListDirectoriesResponseContent.md)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDisks

> ListDisksResponseContent ListDisks(ctx, node).IncludePartitions(includePartitions).Skipsmart(skipsmart).Type_(type_).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    node := "node_example" // string | 
    includePartitions := float32(8.14) // float32 | Include partitions in response list. Takes a boolean integer value (0 false, 1 true). (optional)
    skipsmart := float32(8.14) // float32 | Skip SMART checks. Takes a boolean integer value (0 false, 1 true). (optional)
    type_ := openapiclient.DiskTypeFilter("unused") // DiskTypeFilter |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListDisks(context.Background(), node).IncludePartitions(includePartitions).Skipsmart(skipsmart).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListDisks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDisks`: ListDisksResponseContent
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListDisks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListDisksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includePartitions** | **float32** | Include partitions in response list. Takes a boolean integer value (0 false, 1 true). | 
 **skipsmart** | **float32** | Skip SMART checks. Takes a boolean integer value (0 false, 1 true). | 
 **type_** | [**DiskTypeFilter**](DiskTypeFilter.md) |  | 

### Return type

[**ListDisksResponseContent**](ListDisksResponseContent.md)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListGroups

> ListGroupsResponseContent ListGroups(ctx).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListGroups(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListGroups`: ListGroupsResponseContent
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListGroups`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListGroupsRequest struct via the builder pattern


### Return type

[**ListGroupsResponseContent**](ListGroupsResponseContent.md)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLVMThins

> ListLVMThinsResponseContent ListLVMThins(ctx, node).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    node := "node_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListLVMThins(context.Background(), node).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListLVMThins``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLVMThins`: ListLVMThinsResponseContent
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListLVMThins`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListLVMThinsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListLVMThinsResponseContent**](ListLVMThinsResponseContent.md)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLVMs

> ListLVMsResponseContent ListLVMs(ctx, node).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    node := "node_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListLVMs(context.Background(), node).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListLVMs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLVMs`: ListLVMsResponseContent
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListLVMs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListLVMsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListLVMsResponseContent**](ListLVMsResponseContent.md)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMachineCapabilities

> ListMachineCapabilitiesResponseContent ListMachineCapabilities(ctx, node).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    node := "node_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListMachineCapabilities(context.Background(), node).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListMachineCapabilities``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListMachineCapabilities`: ListMachineCapabilitiesResponseContent
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListMachineCapabilities`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListMachineCapabilitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListMachineCapabilitiesResponseContent**](ListMachineCapabilitiesResponseContent.md)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNetworkInterfaces

> ListNetworkInterfacesResponseContent ListNetworkInterfaces(ctx, node).Type_(type_).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    node := "node_example" // string | 
    type_ := openapiclient.NetworkInterfaceType("bridge") // NetworkInterfaceType |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListNetworkInterfaces(context.Background(), node).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListNetworkInterfaces``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListNetworkInterfaces`: ListNetworkInterfacesResponseContent
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListNetworkInterfaces`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListNetworkInterfacesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | [**NetworkInterfaceType**](NetworkInterfaceType.md) |  | 

### Return type

[**ListNetworkInterfacesResponseContent**](ListNetworkInterfacesResponseContent.md)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNodeCertificates

> ListNodeCertificatesResponseContent ListNodeCertificates(ctx, node).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    node := "node_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListNodeCertificates(context.Background(), node).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListNodeCertificates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListNodeCertificates`: ListNodeCertificatesResponseContent
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListNodeCertificates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListNodeCertificatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListNodeCertificatesResponseContent**](ListNodeCertificatesResponseContent.md)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNodeStorage

> ListNodeStorageResponseContent ListNodeStorage(ctx, node).Content(content).Enabled(enabled).Storage(storage).Target(target).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    node := "node_example" // string | 
    content := "content_example" // string |  (optional)
    enabled := float32(8.14) // float32 | An integer used to represent a boolean. 0 is false, 1 is true. (optional)
    storage := "storage_example" // string |  (optional)
    target := "target_example" // string | If target and node differ, only return storage that is available on both nodes. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListNodeStorage(context.Background(), node).Content(content).Enabled(enabled).Storage(storage).Target(target).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListNodeStorage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListNodeStorage`: ListNodeStorageResponseContent
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListNodeStorage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListNodeStorageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **content** | **string** |  | 
 **enabled** | **float32** | An integer used to represent a boolean. 0 is false, 1 is true. | 
 **storage** | **string** |  | 
 **target** | **string** | If target and node differ, only return storage that is available on both nodes. | 

### Return type

[**ListNodeStorageResponseContent**](ListNodeStorageResponseContent.md)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNodes

> ListNodesResponseContent ListNodes(ctx).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListNodes(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListNodes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListNodes`: ListNodesResponseContent
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListNodes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListNodesRequest struct via the builder pattern


### Return type

[**ListNodesResponseContent**](ListNodesResponseContent.md)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPackages

> ListPackagesResponseContent ListPackages(ctx, node).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    node := "node_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListPackages(context.Background(), node).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListPackages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPackages`: ListPackagesResponseContent
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListPackages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListPackagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListPackagesResponseContent**](ListPackagesResponseContent.md)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPciDeviceMediatedDevices

> ListPciDeviceMediatedDevicesResponseContent ListPciDeviceMediatedDevices(ctx, node, deviceId).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    node := "node_example" // string | 
    deviceId := "deviceId_example" // string | The PCI device to get mediated devices for.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListPciDeviceMediatedDevices(context.Background(), node, deviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListPciDeviceMediatedDevices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPciDeviceMediatedDevices`: ListPciDeviceMediatedDevicesResponseContent
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListPciDeviceMediatedDevices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** |  | 
**deviceId** | **string** | The PCI device to get mediated devices for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListPciDeviceMediatedDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ListPciDeviceMediatedDevicesResponseContent**](ListPciDeviceMediatedDevicesResponseContent.md)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPciDevices

> ListPciDevicesResponseContent ListPciDevices(ctx, node).PciClassBlacklist(pciClassBlacklist).Verbose(verbose).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    node := "node_example" // string | 
    pciClassBlacklist := "pciClassBlacklist_example" // string | Comma seperated list of PCI class IDs to exclude from the list (optional)
    verbose := float32(8.14) // float32 | An integer used to represent a boolean. 0 is false, 1 is true. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListPciDevices(context.Background(), node).PciClassBlacklist(pciClassBlacklist).Verbose(verbose).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListPciDevices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPciDevices`: ListPciDevicesResponseContent
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListPciDevices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListPciDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pciClassBlacklist** | **string** | Comma seperated list of PCI class IDs to exclude from the list | 
 **verbose** | **float32** | An integer used to represent a boolean. 0 is false, 1 is true. | 

### Return type

[**ListPciDevicesResponseContent**](ListPciDevicesResponseContent.md)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPendingVirtualMachineConfigurationChanges

> ListPendingVirtualMachineConfigurationChangesResponseContent ListPendingVirtualMachineConfigurationChanges(ctx, node, vmId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    node := "node_example" // string | 
    vmId := "vmId_example" // string | The id of the virtual machine as a string

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListPendingVirtualMachineConfigurationChanges(context.Background(), node, vmId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListPendingVirtualMachineConfigurationChanges``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPendingVirtualMachineConfigurationChanges`: ListPendingVirtualMachineConfigurationChangesResponseContent
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListPendingVirtualMachineConfigurationChanges`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** |  | 
**vmId** | **string** | The id of the virtual machine as a string | 

### Other Parameters

Other parameters are passed through a pointer to a apiListPendingVirtualMachineConfigurationChangesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ListPendingVirtualMachineConfigurationChangesResponseContent**](ListPendingVirtualMachineConfigurationChangesResponseContent.md)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPools

> ListPoolsResponseContent ListPools(ctx).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListPools(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListPools``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPools`: ListPoolsResponseContent
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListPools`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListPoolsRequest struct via the builder pattern


### Return type

[**ListPoolsResponseContent**](ListPoolsResponseContent.md)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRealms

> ListRealmsResponseContent ListRealms(ctx).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListRealms(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListRealms``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRealms`: ListRealmsResponseContent
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListRealms`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListRealmsRequest struct via the builder pattern


### Return type

[**ListRealmsResponseContent**](ListRealmsResponseContent.md)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRepositoriesInformation

> ListRepositoriesInformationResponseContent ListRepositoriesInformation(ctx, node).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    node := "node_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListRepositoriesInformation(context.Background(), node).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListRepositoriesInformation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRepositoriesInformation`: ListRepositoriesInformationResponseContent
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListRepositoriesInformation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListRepositoriesInformationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListRepositoriesInformationResponseContent**](ListRepositoriesInformationResponseContent.md)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRoles

> ListRolesResponseContent ListRoles(ctx).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListRoles(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListRoles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRoles`: ListRolesResponseContent
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListRoles`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListRolesRequest struct via the builder pattern


### Return type

[**ListRolesResponseContent**](ListRolesResponseContent.md)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSnapshots

> ListSnapshotsResponseContent ListSnapshots(ctx, node, vmId).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    node := "node_example" // string | 
    vmId := "vmId_example" // string | The id of the virtual machine as a string

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListSnapshots(context.Background(), node, vmId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListSnapshots``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSnapshots`: ListSnapshotsResponseContent
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListSnapshots`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** |  | 
**vmId** | **string** | The id of the virtual machine as a string | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSnapshotsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ListSnapshotsResponseContent**](ListSnapshotsResponseContent.md)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListStorage

> ListStorageResponseContent ListStorage(ctx).Type_(type_).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    type_ := openapiclient.StorageType("zfspool") // StorageType |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListStorage(context.Background()).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListStorage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListStorage`: ListStorageResponseContent
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListStorage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListStorageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | [**StorageType**](StorageType.md) |  | 

### Return type

[**ListStorageResponseContent**](ListStorageResponseContent.md)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListStorageVolumes

> ListStorageVolumesResponseContent ListStorageVolumes(ctx, node, storage).Content(content).Vmid(vmid).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    node := "node_example" // string | 
    storage := "storage_example" // string | 
    content := "content_example" // string |  (optional)
    vmid := "vmid_example" // string | The id of the virtual machine as a string (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListStorageVolumes(context.Background(), node, storage).Content(content).Vmid(vmid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListStorageVolumes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListStorageVolumes`: ListStorageVolumesResponseContent
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListStorageVolumes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** |  | 
**storage** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListStorageVolumesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **content** | **string** |  | 
 **vmid** | **string** | The id of the virtual machine as a string | 

### Return type

[**ListStorageVolumesResponseContent**](ListStorageVolumesResponseContent.md)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUpdates

> ListUpdatesResponseContent ListUpdates(ctx, node).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    node := "node_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListUpdates(context.Background(), node).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListUpdates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUpdates`: ListUpdatesResponseContent
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListUpdates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListUpdatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListUpdatesResponseContent**](ListUpdatesResponseContent.md)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUsbDevices

> ListUsbDevicesResponseContent ListUsbDevices(ctx, node).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    node := "node_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListUsbDevices(context.Background(), node).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListUsbDevices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUsbDevices`: ListUsbDevicesResponseContent
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListUsbDevices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListUsbDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListUsbDevicesResponseContent**](ListUsbDevicesResponseContent.md)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUserTokens

> ListUserTokensResponseContent ListUserTokens(ctx, userId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    userId := "userId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListUserTokens(context.Background(), userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListUserTokens``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUserTokens`: ListUserTokensResponseContent
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListUserTokens`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListUserTokensRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListUserTokensResponseContent**](ListUserTokensResponseContent.md)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUsers

> ListUsersResponseContent ListUsers(ctx).Enabled(enabled).Full(full).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    enabled := float32(8.14) // float32 | An integer used to represent a boolean. 0 is false, 1 is true. (optional)
    full := float32(8.14) // float32 | An integer used to represent a boolean. 0 is false, 1 is true. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListUsers(context.Background()).Enabled(enabled).Full(full).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUsers`: ListUsersResponseContent
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **enabled** | **float32** | An integer used to represent a boolean. 0 is false, 1 is true. | 
 **full** | **float32** | An integer used to represent a boolean. 0 is false, 1 is true. | 

### Return type

[**ListUsersResponseContent**](ListUsersResponseContent.md)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVirtualMachineFirewallReferences

> ListVirtualMachineFirewallReferencesResponseContent ListVirtualMachineFirewallReferences(ctx, node, vmId).Type_(type_).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    node := "node_example" // string | 
    vmId := "vmId_example" // string | The id of the virtual machine as a string
    type_ := openapiclient.FirewallReferenceType("ipset") // FirewallReferenceType |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListVirtualMachineFirewallReferences(context.Background(), node, vmId).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListVirtualMachineFirewallReferences``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVirtualMachineFirewallReferences`: ListVirtualMachineFirewallReferencesResponseContent
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListVirtualMachineFirewallReferences`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** |  | 
**vmId** | **string** | The id of the virtual machine as a string | 

### Other Parameters

Other parameters are passed through a pointer to a apiListVirtualMachineFirewallReferencesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **type_** | [**FirewallReferenceType**](FirewallReferenceType.md) |  | 

### Return type

[**ListVirtualMachineFirewallReferencesResponseContent**](ListVirtualMachineFirewallReferencesResponseContent.md)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVirtualMachines

> ListVirtualMachinesResponseContent ListVirtualMachines(ctx, node).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    node := "node_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListVirtualMachines(context.Background(), node).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListVirtualMachines``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVirtualMachines`: ListVirtualMachinesResponseContent
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListVirtualMachines`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListVirtualMachinesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListVirtualMachinesResponseContent**](ListVirtualMachinesResponseContent.md)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListZFSPools

> ListZFSPoolsResponseContent ListZFSPools(ctx, node).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    node := "node_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListZFSPools(context.Background(), node).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListZFSPools``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListZFSPools`: ListZFSPoolsResponseContent
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListZFSPools`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListZFSPoolsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListZFSPoolsResponseContent**](ListZFSPoolsResponseContent.md)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyPool

> ModifyPool(ctx, poolId).ModifyPoolRequestContent(modifyPoolRequestContent).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    poolId := "poolId_example" // string | 
    modifyPoolRequestContent := *openapiclient.NewModifyPoolRequestContent() // ModifyPoolRequestContent |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultApi.ModifyPool(context.Background(), poolId).ModifyPoolRequestContent(modifyPoolRequestContent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ModifyPool``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**poolId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyPoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **modifyPoolRequestContent** | [**ModifyPoolRequestContent**](ModifyPoolRequestContent.md) |  | 

### Return type

 (empty response body)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyRole

> ModifyRole(ctx, roleId).ModifyRoleRequestContent(modifyRoleRequestContent).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    roleId := "roleId_example" // string | 
    modifyRoleRequestContent := *openapiclient.NewModifyRoleRequestContent() // ModifyRoleRequestContent |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultApi.ModifyRole(context.Background(), roleId).ModifyRoleRequestContent(modifyRoleRequestContent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ModifyRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **modifyRoleRequestContent** | [**ModifyRoleRequestContent**](ModifyRoleRequestContent.md) |  | 

### Return type

 (empty response body)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyStorage

> ModifyStorageResponseContent ModifyStorage(ctx, storage).ModifyStorageRequestContent(modifyStorageRequestContent).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    storage := "storage_example" // string | 
    modifyStorageRequestContent := *openapiclient.NewModifyStorageRequestContent() // ModifyStorageRequestContent |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ModifyStorage(context.Background(), storage).ModifyStorageRequestContent(modifyStorageRequestContent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ModifyStorage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ModifyStorage`: ModifyStorageResponseContent
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ModifyStorage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storage** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyStorageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **modifyStorageRequestContent** | [**ModifyStorageRequestContent**](ModifyStorageRequestContent.md) |  | 

### Return type

[**ModifyStorageResponseContent**](ModifyStorageResponseContent.md)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyUser

> ModifyUser(ctx, userId).ModifyUserRequestContent(modifyUserRequestContent).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    userId := "userId_example" // string | 
    modifyUserRequestContent := *openapiclient.NewModifyUserRequestContent() // ModifyUserRequestContent |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultApi.ModifyUser(context.Background(), userId).ModifyUserRequestContent(modifyUserRequestContent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ModifyUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **modifyUserRequestContent** | [**ModifyUserRequestContent**](ModifyUserRequestContent.md) |  | 

### Return type

 (empty response body)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyUserToken

> ModifyUserTokenResponseContent ModifyUserToken(ctx, userId, tokenId).ModifyUserTokenRequestContent(modifyUserTokenRequestContent).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    userId := "userId_example" // string | 
    tokenId := "tokenId_example" // string | 
    modifyUserTokenRequestContent := *openapiclient.NewModifyUserTokenRequestContent() // ModifyUserTokenRequestContent |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ModifyUserToken(context.Background(), userId, tokenId).ModifyUserTokenRequestContent(modifyUserTokenRequestContent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ModifyUserToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ModifyUserToken`: ModifyUserTokenResponseContent
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ModifyUserToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 
**tokenId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyUserTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **modifyUserTokenRequestContent** | [**ModifyUserTokenRequestContent**](ModifyUserTokenRequestContent.md) |  | 

### Return type

[**ModifyUserTokenResponseContent**](ModifyUserTokenResponseContent.md)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrderNodeCertificate

> OrderNodeCertificateResponseContent OrderNodeCertificate(ctx, node).OrderNodeCertificateRequestContent(orderNodeCertificateRequestContent).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    node := "node_example" // string | 
    orderNodeCertificateRequestContent := *openapiclient.NewOrderNodeCertificateRequestContent() // OrderNodeCertificateRequestContent |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.OrderNodeCertificate(context.Background(), node).OrderNodeCertificateRequestContent(orderNodeCertificateRequestContent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.OrderNodeCertificate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrderNodeCertificate`: OrderNodeCertificateResponseContent
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.OrderNodeCertificate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrderNodeCertificateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **orderNodeCertificateRequestContent** | [**OrderNodeCertificateRequestContent**](OrderNodeCertificateRequestContent.md) |  | 

### Return type

[**OrderNodeCertificateResponseContent**](OrderNodeCertificateResponseContent.md)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PingVirtualMachine

> PingVirtualMachineResponseContent PingVirtualMachine(ctx, node, vmId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    node := "node_example" // string | 
    vmId := "vmId_example" // string | The id of the virtual machine as a string

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.PingVirtualMachine(context.Background(), node, vmId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.PingVirtualMachine``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PingVirtualMachine`: PingVirtualMachineResponseContent
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.PingVirtualMachine`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** |  | 
**vmId** | **string** | The id of the virtual machine as a string | 

### Other Parameters

Other parameters are passed through a pointer to a apiPingVirtualMachineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PingVirtualMachineResponseContent**](PingVirtualMachineResponseContent.md)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RebootVirtualMachine

> RebootVirtualMachineResponseContent RebootVirtualMachine(ctx, node, vmId).RebootVirtualMachineRequestContent(rebootVirtualMachineRequestContent).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    node := "node_example" // string | 
    vmId := "vmId_example" // string | The id of the virtual machine as a string
    rebootVirtualMachineRequestContent := *openapiclient.NewRebootVirtualMachineRequestContent() // RebootVirtualMachineRequestContent |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.RebootVirtualMachine(context.Background(), node, vmId).RebootVirtualMachineRequestContent(rebootVirtualMachineRequestContent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.RebootVirtualMachine``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RebootVirtualMachine`: RebootVirtualMachineResponseContent
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.RebootVirtualMachine`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** |  | 
**vmId** | **string** | The id of the virtual machine as a string | 

### Other Parameters

Other parameters are passed through a pointer to a apiRebootVirtualMachineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **rebootVirtualMachineRequestContent** | [**RebootVirtualMachineRequestContent**](RebootVirtualMachineRequestContent.md) |  | 

### Return type

[**RebootVirtualMachineResponseContent**](RebootVirtualMachineResponseContent.md)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegenerateVirtualMachineCloudInit

> RegenerateVirtualMachineCloudInit(ctx, node, vmId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    node := "node_example" // string | 
    vmId := "vmId_example" // string | The id of the virtual machine as a string

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultApi.RegenerateVirtualMachineCloudInit(context.Background(), node, vmId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.RegenerateVirtualMachineCloudInit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** |  | 
**vmId** | **string** | The id of the virtual machine as a string | 

### Other Parameters

Other parameters are passed through a pointer to a apiRegenerateVirtualMachineCloudInitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveCorosyncNode

> RemoveCorosyncNode(ctx, node).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    node := "node_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultApi.RemoveCorosyncNode(context.Background(), node).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.RemoveCorosyncNode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveCorosyncNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RenewNodeCertificate

> RenewNodeCertificateResponseContent RenewNodeCertificate(ctx, node).RenewNodeCertificateRequestContent(renewNodeCertificateRequestContent).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    node := "node_example" // string | 
    renewNodeCertificateRequestContent := *openapiclient.NewRenewNodeCertificateRequestContent() // RenewNodeCertificateRequestContent |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.RenewNodeCertificate(context.Background(), node).RenewNodeCertificateRequestContent(renewNodeCertificateRequestContent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.RenewNodeCertificate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RenewNodeCertificate`: RenewNodeCertificateResponseContent
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.RenewNodeCertificate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRenewNodeCertificateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **renewNodeCertificateRequestContent** | [**RenewNodeCertificateRequestContent**](RenewNodeCertificateRequestContent.md) |  | 

### Return type

[**RenewNodeCertificateResponseContent**](RenewNodeCertificateResponseContent.md)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResetVirtualMachine

> ResetVirtualMachineResponseContent ResetVirtualMachine(ctx, node, vmId).ResetVirtualMachineRequestContent(resetVirtualMachineRequestContent).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    node := "node_example" // string | 
    vmId := "vmId_example" // string | The id of the virtual machine as a string
    resetVirtualMachineRequestContent := *openapiclient.NewResetVirtualMachineRequestContent() // ResetVirtualMachineRequestContent |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ResetVirtualMachine(context.Background(), node, vmId).ResetVirtualMachineRequestContent(resetVirtualMachineRequestContent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ResetVirtualMachine``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ResetVirtualMachine`: ResetVirtualMachineResponseContent
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ResetVirtualMachine`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** |  | 
**vmId** | **string** | The id of the virtual machine as a string | 

### Other Parameters

Other parameters are passed through a pointer to a apiResetVirtualMachineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **resetVirtualMachineRequestContent** | [**ResetVirtualMachineRequestContent**](ResetVirtualMachineRequestContent.md) |  | 

### Return type

[**ResetVirtualMachineResponseContent**](ResetVirtualMachineResponseContent.md)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResizeVirtualMachineDisk

> ResizeVirtualMachineDisk(ctx, node, vmId).Disk(disk).Size(size).Digest(digest).Skiplock(skiplock).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    node := "node_example" // string | 
    vmId := "vmId_example" // string | The id of the virtual machine as a string
    disk := openapiclient.VirtualMachineDiskTarget("ide0") // VirtualMachineDiskTarget | The name of the disk to resize.
    size := "size_example" // string | The new size of the disk in bytes, or with a suffix of K, M, G, or T for kilobytes, megabytes, gigabytes, or terabytes. If + is specified, the size is increased by the given amount.
    digest := "digest_example" // string | The SHA1 digest of the current configuration. Used to prevent concurrent operations. (optional)
    skiplock := true // bool | Ignore lock. Only valid if authenticated as root user. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultApi.ResizeVirtualMachineDisk(context.Background(), node, vmId).Disk(disk).Size(size).Digest(digest).Skiplock(skiplock).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ResizeVirtualMachineDisk``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** |  | 
**vmId** | **string** | The id of the virtual machine as a string | 

### Other Parameters

Other parameters are passed through a pointer to a apiResizeVirtualMachineDiskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **disk** | [**VirtualMachineDiskTarget**](VirtualMachineDiskTarget.md) | The name of the disk to resize. | 
 **size** | **string** | The new size of the disk in bytes, or with a suffix of K, M, G, or T for kilobytes, megabytes, gigabytes, or terabytes. If + is specified, the size is increased by the given amount. | 
 **digest** | **string** | The SHA1 digest of the current configuration. Used to prevent concurrent operations. | 
 **skiplock** | **bool** | Ignore lock. Only valid if authenticated as root user. | 

### Return type

 (empty response body)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResumeVirtualMachine

> ResumeVirtualMachineResponseContent ResumeVirtualMachine(ctx, node, vmId).ResumeVirtualMachineRequestContent(resumeVirtualMachineRequestContent).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    node := "node_example" // string | 
    vmId := "vmId_example" // string | The id of the virtual machine as a string
    resumeVirtualMachineRequestContent := *openapiclient.NewResumeVirtualMachineRequestContent() // ResumeVirtualMachineRequestContent |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ResumeVirtualMachine(context.Background(), node, vmId).ResumeVirtualMachineRequestContent(resumeVirtualMachineRequestContent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ResumeVirtualMachine``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ResumeVirtualMachine`: ResumeVirtualMachineResponseContent
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ResumeVirtualMachine`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** |  | 
**vmId** | **string** | The id of the virtual machine as a string | 

### Other Parameters

Other parameters are passed through a pointer to a apiResumeVirtualMachineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **resumeVirtualMachineRequestContent** | [**ResumeVirtualMachineRequestContent**](ResumeVirtualMachineRequestContent.md) |  | 

### Return type

[**ResumeVirtualMachineResponseContent**](ResumeVirtualMachineResponseContent.md)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RevertNetworkInterfaceConfiguration

> RevertNetworkInterfaceConfiguration(ctx, node).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    node := "node_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultApi.RevertNetworkInterfaceConfiguration(context.Background(), node).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.RevertNetworkInterfaceConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRevertNetworkInterfaceConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShutdownVirtualMachine

> ShutdownVirtualMachineResponseContent ShutdownVirtualMachine(ctx, node, vmId).ShutdownVirtualMachineRequestContent(shutdownVirtualMachineRequestContent).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    node := "node_example" // string | 
    vmId := "vmId_example" // string | The id of the virtual machine as a string
    shutdownVirtualMachineRequestContent := *openapiclient.NewShutdownVirtualMachineRequestContent() // ShutdownVirtualMachineRequestContent |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ShutdownVirtualMachine(context.Background(), node, vmId).ShutdownVirtualMachineRequestContent(shutdownVirtualMachineRequestContent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ShutdownVirtualMachine``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ShutdownVirtualMachine`: ShutdownVirtualMachineResponseContent
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ShutdownVirtualMachine`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** |  | 
**vmId** | **string** | The id of the virtual machine as a string | 

### Other Parameters

Other parameters are passed through a pointer to a apiShutdownVirtualMachineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **shutdownVirtualMachineRequestContent** | [**ShutdownVirtualMachineRequestContent**](ShutdownVirtualMachineRequestContent.md) |  | 

### Return type

[**ShutdownVirtualMachineResponseContent**](ShutdownVirtualMachineResponseContent.md)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartVirtualMachine

> StartVirtualMachineResponseContent StartVirtualMachine(ctx, node, vmId).StartVirtualMachineRequestContent(startVirtualMachineRequestContent).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    node := "node_example" // string | 
    vmId := "vmId_example" // string | The id of the virtual machine as a string
    startVirtualMachineRequestContent := *openapiclient.NewStartVirtualMachineRequestContent() // StartVirtualMachineRequestContent |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.StartVirtualMachine(context.Background(), node, vmId).StartVirtualMachineRequestContent(startVirtualMachineRequestContent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.StartVirtualMachine``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StartVirtualMachine`: StartVirtualMachineResponseContent
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.StartVirtualMachine`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** |  | 
**vmId** | **string** | The id of the virtual machine as a string | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartVirtualMachineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **startVirtualMachineRequestContent** | [**StartVirtualMachineRequestContent**](StartVirtualMachineRequestContent.md) |  | 

### Return type

[**StartVirtualMachineResponseContent**](StartVirtualMachineResponseContent.md)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StopVirtualMachine

> StopVirtualMachineResponseContent StopVirtualMachine(ctx, node, vmId).StopVirtualMachineRequestContent(stopVirtualMachineRequestContent).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    node := "node_example" // string | 
    vmId := "vmId_example" // string | The id of the virtual machine as a string
    stopVirtualMachineRequestContent := *openapiclient.NewStopVirtualMachineRequestContent() // StopVirtualMachineRequestContent |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.StopVirtualMachine(context.Background(), node, vmId).StopVirtualMachineRequestContent(stopVirtualMachineRequestContent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.StopVirtualMachine``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StopVirtualMachine`: StopVirtualMachineResponseContent
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.StopVirtualMachine`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** |  | 
**vmId** | **string** | The id of the virtual machine as a string | 

### Other Parameters

Other parameters are passed through a pointer to a apiStopVirtualMachineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **stopVirtualMachineRequestContent** | [**StopVirtualMachineRequestContent**](StopVirtualMachineRequestContent.md) |  | 

### Return type

[**StopVirtualMachineResponseContent**](StopVirtualMachineResponseContent.md)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SuspendVirtualMachine

> SuspendVirtualMachineResponseContent SuspendVirtualMachine(ctx, node, vmId).SuspendVirtualMachineRequestContent(suspendVirtualMachineRequestContent).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    node := "node_example" // string | 
    vmId := "vmId_example" // string | The id of the virtual machine as a string
    suspendVirtualMachineRequestContent := *openapiclient.NewSuspendVirtualMachineRequestContent() // SuspendVirtualMachineRequestContent |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.SuspendVirtualMachine(context.Background(), node, vmId).SuspendVirtualMachineRequestContent(suspendVirtualMachineRequestContent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SuspendVirtualMachine``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SuspendVirtualMachine`: SuspendVirtualMachineResponseContent
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.SuspendVirtualMachine`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** |  | 
**vmId** | **string** | The id of the virtual machine as a string | 

### Other Parameters

Other parameters are passed through a pointer to a apiSuspendVirtualMachineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **suspendVirtualMachineRequestContent** | [**SuspendVirtualMachineRequestContent**](SuspendVirtualMachineRequestContent.md) |  | 

### Return type

[**SuspendVirtualMachineResponseContent**](SuspendVirtualMachineResponseContent.md)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SyncRealm

> SyncRealm(ctx, realm).SyncRealmRequestContent(syncRealmRequestContent).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    realm := "realm_example" // string | 
    syncRealmRequestContent := *openapiclient.NewSyncRealmRequestContent() // SyncRealmRequestContent |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultApi.SyncRealm(context.Background(), realm).SyncRealmRequestContent(syncRealmRequestContent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SyncRealm``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSyncRealmRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **syncRealmRequestContent** | [**SyncRealmRequestContent**](SyncRealmRequestContent.md) |  | 

### Return type

 (empty response body)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnlinkVirtualMachineDisks

> UnlinkVirtualMachineDisks(ctx, node, vmId).Idlist(idlist).Force(force).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    node := "node_example" // string | 
    vmId := "vmId_example" // string | The id of the virtual machine as a string
    idlist := "idlist_example" // string | A list of disk ids to unlink.
    force := true // bool | Fore removal of disk. Without this, the disk is replaced with a disk entry of 'unused[n]'.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultApi.UnlinkVirtualMachineDisks(context.Background(), node, vmId).Idlist(idlist).Force(force).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UnlinkVirtualMachineDisks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** |  | 
**vmId** | **string** | The id of the virtual machine as a string | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnlinkVirtualMachineDisksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **idlist** | **string** | A list of disk ids to unlink. | 
 **force** | **bool** | Fore removal of disk. Without this, the disk is replaced with a disk entry of &#39;unused[n]&#39;. | 

### Return type

 (empty response body)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAccessControlList

> UpdateAccessControlList(ctx).UpdateAccessControlListRequestContent(updateAccessControlListRequestContent).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    updateAccessControlListRequestContent := *openapiclient.NewUpdateAccessControlListRequestContent("Path_example", "Roles_example") // UpdateAccessControlListRequestContent | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultApi.UpdateAccessControlList(context.Background()).UpdateAccessControlListRequestContent(updateAccessControlListRequestContent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateAccessControlList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAccessControlListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateAccessControlListRequestContent** | [**UpdateAccessControlListRequestContent**](UpdateAccessControlListRequestContent.md) |  | 

### Return type

 (empty response body)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGroup

> UpdateGroup(ctx, groupId).UpdateGroupRequestContent(updateGroupRequestContent).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    groupId := "groupId_example" // string | 
    updateGroupRequestContent := *openapiclient.NewUpdateGroupRequestContent() // UpdateGroupRequestContent |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultApi.UpdateGroup(context.Background(), groupId).UpdateGroupRequestContent(updateGroupRequestContent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateGroupRequestContent** | [**UpdateGroupRequestContent**](UpdateGroupRequestContent.md) |  | 

### Return type

 (empty response body)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkInterface

> UpdateNetworkInterface(ctx, node, interface_).UpdateNetworkInterfaceRequestContent(updateNetworkInterfaceRequestContent).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    node := "node_example" // string | 
    interface_ := "interface__example" // string | 
    updateNetworkInterfaceRequestContent := *openapiclient.NewUpdateNetworkInterfaceRequestContent(openapiclient.NetworkInterfaceType("bridge")) // UpdateNetworkInterfaceRequestContent | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultApi.UpdateNetworkInterface(context.Background(), node, interface_).UpdateNetworkInterfaceRequestContent(updateNetworkInterfaceRequestContent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateNetworkInterface``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** |  | 
**interface_** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkInterfaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkInterfaceRequestContent** | [**UpdateNetworkInterfaceRequestContent**](UpdateNetworkInterfaceRequestContent.md) |  | 

### Return type

 (empty response body)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNodeDns

> UpdateNodeDns(ctx, node).Search(search).UpdateNodeDnsRequestContent(updateNodeDnsRequestContent).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    node := "node_example" // string | 
    search := "search_example" // string | 
    updateNodeDnsRequestContent := *openapiclient.NewUpdateNodeDnsRequestContent() // UpdateNodeDnsRequestContent |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultApi.UpdateNodeDns(context.Background(), node).Search(search).UpdateNodeDnsRequestContent(updateNodeDnsRequestContent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateNodeDns``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNodeDnsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **search** | **string** |  | 
 **updateNodeDnsRequestContent** | [**UpdateNodeDnsRequestContent**](UpdateNodeDnsRequestContent.md) |  | 

### Return type

 (empty response body)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRealm

> UpdateRealm(ctx, realm).UpdateRealmRequestContent(updateRealmRequestContent).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    realm := "realm_example" // string | 
    updateRealmRequestContent := *openapiclient.NewUpdateRealmRequestContent() // UpdateRealmRequestContent |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultApi.UpdateRealm(context.Background(), realm).UpdateRealmRequestContent(updateRealmRequestContent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateRealm``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRealmRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRealmRequestContent** | [**UpdateRealmRequestContent**](UpdateRealmRequestContent.md) |  | 

### Return type

 (empty response body)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateStorageVolume

> UpdateStorageVolume(ctx, node, storage, volume).UpdateStorageVolumeRequestContent(updateStorageVolumeRequestContent).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    node := "node_example" // string | 
    storage := "storage_example" // string | 
    volume := "volume_example" // string | 
    updateStorageVolumeRequestContent := *openapiclient.NewUpdateStorageVolumeRequestContent(float32(123)) // UpdateStorageVolumeRequestContent | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultApi.UpdateStorageVolume(context.Background(), node, storage, volume).UpdateStorageVolumeRequestContent(updateStorageVolumeRequestContent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateStorageVolume``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** |  | 
**storage** | **string** |  | 
**volume** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateStorageVolumeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **updateStorageVolumeRequestContent** | [**UpdateStorageVolumeRequestContent**](UpdateStorageVolumeRequestContent.md) |  | 

### Return type

 (empty response body)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadToStorage

> UploadToStorageResponseContent UploadToStorage(ctx, node, storage).UploadToStorageRequestContent(uploadToStorageRequestContent).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    node := "node_example" // string | 
    storage := "storage_example" // string | 
    uploadToStorageRequestContent := *openapiclient.NewUploadToStorageRequestContent(openapiclient.UploadContentType("iso"), "Filename_example") // UploadToStorageRequestContent | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.UploadToStorage(context.Background(), node, storage).UploadToStorageRequestContent(uploadToStorageRequestContent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UploadToStorage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UploadToStorage`: UploadToStorageResponseContent
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UploadToStorage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** |  | 
**storage** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadToStorageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **uploadToStorageRequestContent** | [**UploadToStorageRequestContent**](UploadToStorageRequestContent.md) |  | 

### Return type

[**UploadToStorageResponseContent**](UploadToStorageResponseContent.md)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VirtualMachineExecuteCommand

> VirtualMachineExecuteCommandResponseContent VirtualMachineExecuteCommand(ctx, node, vmId).VirtualMachineExecuteCommandRequestContent(virtualMachineExecuteCommandRequestContent).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    node := "node_example" // string | 
    vmId := "vmId_example" // string | The id of the virtual machine as a string
    virtualMachineExecuteCommandRequestContent := *openapiclient.NewVirtualMachineExecuteCommandRequestContent() // VirtualMachineExecuteCommandRequestContent |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.VirtualMachineExecuteCommand(context.Background(), node, vmId).VirtualMachineExecuteCommandRequestContent(virtualMachineExecuteCommandRequestContent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.VirtualMachineExecuteCommand``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VirtualMachineExecuteCommand`: VirtualMachineExecuteCommandResponseContent
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.VirtualMachineExecuteCommand`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** |  | 
**vmId** | **string** | The id of the virtual machine as a string | 

### Other Parameters

Other parameters are passed through a pointer to a apiVirtualMachineExecuteCommandRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **virtualMachineExecuteCommandRequestContent** | [**VirtualMachineExecuteCommandRequestContent**](VirtualMachineExecuteCommandRequestContent.md) |  | 

### Return type

[**VirtualMachineExecuteCommandResponseContent**](VirtualMachineExecuteCommandResponseContent.md)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VirtualMachineReadFile

> VirtualMachineReadFileResponseContent VirtualMachineReadFile(ctx, node, vmId).File(file).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    node := "node_example" // string | 
    vmId := "vmId_example" // string | The id of the virtual machine as a string
    file := "file_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.VirtualMachineReadFile(context.Background(), node, vmId).File(file).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.VirtualMachineReadFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VirtualMachineReadFile`: VirtualMachineReadFileResponseContent
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.VirtualMachineReadFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** |  | 
**vmId** | **string** | The id of the virtual machine as a string | 

### Other Parameters

Other parameters are passed through a pointer to a apiVirtualMachineReadFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **file** | **string** |  | 

### Return type

[**VirtualMachineReadFileResponseContent**](VirtualMachineReadFileResponseContent.md)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VirtualMachineWriteFile

> VirtualMachineWriteFile(ctx, node, vmId).VirtualMachineWriteFileRequestContent(virtualMachineWriteFileRequestContent).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    node := "node_example" // string | 
    vmId := "vmId_example" // string | The id of the virtual machine as a string
    virtualMachineWriteFileRequestContent := *openapiclient.NewVirtualMachineWriteFileRequestContent("File_example", "Content_example") // VirtualMachineWriteFileRequestContent | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultApi.VirtualMachineWriteFile(context.Background(), node, vmId).VirtualMachineWriteFileRequestContent(virtualMachineWriteFileRequestContent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.VirtualMachineWriteFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** |  | 
**vmId** | **string** | The id of the virtual machine as a string | 

### Other Parameters

Other parameters are passed through a pointer to a apiVirtualMachineWriteFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **virtualMachineWriteFileRequestContent** | [**VirtualMachineWriteFileRequestContent**](VirtualMachineWriteFileRequestContent.md) |  | 

### Return type

 (empty response body)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WipeDisk

> WipeDisk(ctx, node).Disk(disk).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/awlsring/proxmox-go"
)

func main() {
    node := "node_example" // string | 
    disk := "disk_example" // string | Disk to wipe. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultApi.WipeDisk(context.Background(), node).Disk(disk).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.WipeDisk``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWipeDiskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **disk** | **string** | Disk to wipe. | 

### Return type

 (empty response body)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth), [smithy.api.httpBasicAuth](../README.md#smithy.api.httpBasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

