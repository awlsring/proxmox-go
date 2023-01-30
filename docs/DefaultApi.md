# \DefaultApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddCorosyncNode**](DefaultApi.md#AddCorosyncNode) | **Post** /cluster/config/nodes/{node} | 
[**ApplyNetworkInterfaceConfiguration**](DefaultApi.md#ApplyNetworkInterfaceConfiguration) | **Put** /nodes/{node}/network | 
[**CreateClusterConfig**](DefaultApi.md#CreateClusterConfig) | **Post** /cluster/config | 
[**CreateDirectory**](DefaultApi.md#CreateDirectory) | **Post** /nodes/{node}/disks/directory | 
[**CreateLVM**](DefaultApi.md#CreateLVM) | **Post** /nodes/{node}/disks/lvm | 
[**CreateLVMThin**](DefaultApi.md#CreateLVMThin) | **Post** /nodes/{node}/disks/lvmthin | 
[**CreateNetworkInterface**](DefaultApi.md#CreateNetworkInterface) | **Post** /nodes/{node}/network | 
[**CreateStorage**](DefaultApi.md#CreateStorage) | **Post** /storage | 
[**CreateTicket**](DefaultApi.md#CreateTicket) | **Post** /access/ticket | 
[**CreateZFSPool**](DefaultApi.md#CreateZFSPool) | **Post** /nodes/{node}/disks/zfs | 
[**DeleteDirectory**](DefaultApi.md#DeleteDirectory) | **Delete** /nodes/{node}/disks/directory/{name} | 
[**DeleteLVM**](DefaultApi.md#DeleteLVM) | **Delete** /nodes/{node}/disks/lvm/{name} | 
[**DeleteLVMThin**](DefaultApi.md#DeleteLVMThin) | **Delete** /nodes/{node}/disks/lvmthin/{name} | 
[**DeleteNetworkInterface**](DefaultApi.md#DeleteNetworkInterface) | **Delete** /nodes/{node}/network/{interface} | 
[**DeleteStorage**](DefaultApi.md#DeleteStorage) | **Delete** /storage/{storage} | 
[**DeleteZFSPool**](DefaultApi.md#DeleteZFSPool) | **Delete** /nodes/{node}/disks/zfs/{name} | 
[**GetAccessControlList**](DefaultApi.md#GetAccessControlList) | **Get** /access/acl | 
[**GetClusterApiVersion**](DefaultApi.md#GetClusterApiVersion) | **Get** /cluster/config/apiversion | 
[**GetClusterJoinInformation**](DefaultApi.md#GetClusterJoinInformation) | **Get** /cluster/config/join | 
[**GetClusterTotemSettings**](DefaultApi.md#GetClusterTotemSettings) | **Get** /cluster/config/totem | 
[**GetNetworkInterface**](DefaultApi.md#GetNetworkInterface) | **Get** /nodes/{node}/network/{interface} | 
[**GetSmartHealth**](DefaultApi.md#GetSmartHealth) | **Get** /nodes/{node}/disks/smart | 
[**GetStorage**](DefaultApi.md#GetStorage) | **Get** /storage/{storage} | 
[**GetVersion**](DefaultApi.md#GetVersion) | **Get** /version | 
[**GetZFSPoolStatus**](DefaultApi.md#GetZFSPoolStatus) | **Get** /nodes/{node}/disks/zfs/{name} | 
[**InitializeGPT**](DefaultApi.md#InitializeGPT) | **Post** /nodes/{node}/disks/smart | 
[**JoinCluster**](DefaultApi.md#JoinCluster) | **Post** /cluster/config/join | 
[**ListCorosyncNodes**](DefaultApi.md#ListCorosyncNodes) | **Get** /cluster/config/nodes | 
[**ListCpuCapabilities**](DefaultApi.md#ListCpuCapabilities) | **Get** /nodes/{node}/capabilities/qemu/cpu | 
[**ListDirectories**](DefaultApi.md#ListDirectories) | **Get** /nodes/{node}/disks/directory | 
[**ListDisks**](DefaultApi.md#ListDisks) | **Get** /nodes/{node}/disks/list | 
[**ListLVMThins**](DefaultApi.md#ListLVMThins) | **Get** /nodes/{node}/disks/lvmthin | 
[**ListLVMs**](DefaultApi.md#ListLVMs) | **Get** /nodes/{node}/disks/lvm | 
[**ListMachineCapabilities**](DefaultApi.md#ListMachineCapabilities) | **Get** /nodes/{node}/capabilities/qemu/machines | 
[**ListNetworkInterfaces**](DefaultApi.md#ListNetworkInterfaces) | **Get** /nodes/{node}/network | 
[**ListNodes**](DefaultApi.md#ListNodes) | **Get** /nodes | 
[**ListStorage**](DefaultApi.md#ListStorage) | **Get** /storage | 
[**ListVirtualMachines**](DefaultApi.md#ListVirtualMachines) | **Get** /nodes/{node}/qemu | 
[**ListZFSPools**](DefaultApi.md#ListZFSPools) | **Get** /nodes/{node}/disks/zfs | 
[**ModifyStorage**](DefaultApi.md#ModifyStorage) | **Put** /storage/{storage} | 
[**RemoveCorosyncNode**](DefaultApi.md#RemoveCorosyncNode) | **Delete** /cluster/config/nodes/{node} | 
[**RevertNetworkInterfaceConfiguration**](DefaultApi.md#RevertNetworkInterfaceConfiguration) | **Delete** /nodes/{node}/network | 
[**UpdateAccessControlList**](DefaultApi.md#UpdateAccessControlList) | **Put** /access/acl | 
[**UpdateNetworkInterface**](DefaultApi.md#UpdateNetworkInterface) | **Put** /nodes/{node}/network/{interface} | 
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
    openapiclient "./openapi"
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

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
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
    openapiclient "./openapi"
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

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
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
    openapiclient "./openapi"
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

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth)

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
    openapiclient "./openapi"
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

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth)

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
    openapiclient "./openapi"
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

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth)

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
    openapiclient "./openapi"
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

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth)

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
    openapiclient "./openapi"
)

func main() {
    node := "node_example" // string | 
    createNetworkInterfaceRequestContent := *openapiclient.NewCreateNetworkInterfaceRequestContent("Iface_example", openapiclient.NetworkInterfaceType("bridge")) // CreateNetworkInterfaceRequestContent | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CreateNetworkInterface(context.Background(), node).CreateNetworkInterfaceRequestContent(createNetworkInterfaceRequestContent).Execute()
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

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth)

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
    openapiclient "./openapi"
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

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth)

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
    openapiclient "./openapi"
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

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth)

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
    openapiclient "./openapi"
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

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
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
    openapiclient "./openapi"
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

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth)

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
    openapiclient "./openapi"
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

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteLVMThin

> DeleteLVMThinResponseContent DeleteLVMThin(ctx, node, name).CleanupConfig(cleanupConfig).CleanupDisks(cleanupDisks).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    node := "node_example" // string | 
    name := "name_example" // string | The storage identifier.
    cleanupConfig := float32(8.14) // float32 | Marks the associated storage as not available on this node anr removes them from the config. Takes a boolean integer value (0 false, 1 true). (optional)
    cleanupDisks := float32(8.14) // float32 | Wipes the disk. Takes a boolean integer value (0 false, 1 true). (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.DeleteLVMThin(context.Background(), node, name).CleanupConfig(cleanupConfig).CleanupDisks(cleanupDisks).Execute()
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


 **cleanupConfig** | **float32** | Marks the associated storage as not available on this node anr removes them from the config. Takes a boolean integer value (0 false, 1 true). | 
 **cleanupDisks** | **float32** | Wipes the disk. Takes a boolean integer value (0 false, 1 true). | 

### Return type

[**DeleteLVMThinResponseContent**](DeleteLVMThinResponseContent.md)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth)

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
    openapiclient "./openapi"
)

func main() {
    node := "node_example" // string | 
    interface_ := "interface__example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.DeleteNetworkInterface(context.Background(), node, interface_).Execute()
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

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth)

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
    openapiclient "./openapi"
)

func main() {
    storage := "storage_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.DeleteStorage(context.Background(), storage).Execute()
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

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth)

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
    openapiclient "./openapi"
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

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
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
    openapiclient "./openapi"
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

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth)

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
    openapiclient "./openapi"
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

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth)

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
    openapiclient "./openapi"
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

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth)

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
    openapiclient "./openapi"
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

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth)

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
    openapiclient "./openapi"
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

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth)

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
    openapiclient "./openapi"
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

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth)

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
    openapiclient "./openapi"
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

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth)

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
    openapiclient "./openapi"
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

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth)

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
    openapiclient "./openapi"
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

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
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
    openapiclient "./openapi"
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

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth)

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
    openapiclient "./openapi"
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

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
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
    openapiclient "./openapi"
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

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth)

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
    openapiclient "./openapi"
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

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth)

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
    openapiclient "./openapi"
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

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth)

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
    openapiclient "./openapi"
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

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth)

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
    openapiclient "./openapi"
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

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth)

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
    openapiclient "./openapi"
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

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth)

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
    openapiclient "./openapi"
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

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth)

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
    openapiclient "./openapi"
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

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth)

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
    openapiclient "./openapi"
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

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListStorage

> ListStorageResponseContent ListStorage(ctx).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListStorage(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListStorage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListStorage`: ListStorageResponseContent
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListStorage`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListStorageRequest struct via the builder pattern


### Return type

[**ListStorageResponseContent**](ListStorageResponseContent.md)

### Authorization

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth)

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
    openapiclient "./openapi"
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

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth)

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
    openapiclient "./openapi"
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

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
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
    openapiclient "./openapi"
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

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
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
    openapiclient "./openapi"
)

func main() {
    node := "node_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.RemoveCorosyncNode(context.Background(), node).Execute()
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

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
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
    openapiclient "./openapi"
)

func main() {
    node := "node_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.RevertNetworkInterfaceConfiguration(context.Background(), node).Execute()
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

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth)

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
    openapiclient "./openapi"
)

func main() {
    updateAccessControlListRequestContent := *openapiclient.NewUpdateAccessControlListRequestContent("Path_example", "Roles_example") // UpdateAccessControlListRequestContent | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.UpdateAccessControlList(context.Background()).UpdateAccessControlListRequestContent(updateAccessControlListRequestContent).Execute()
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

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth)

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
    openapiclient "./openapi"
)

func main() {
    node := "node_example" // string | 
    interface_ := "interface__example" // string | 
    updateNetworkInterfaceRequestContent := *openapiclient.NewUpdateNetworkInterfaceRequestContent(openapiclient.NetworkInterfaceType("bridge")) // UpdateNetworkInterfaceRequestContent | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.UpdateNetworkInterface(context.Background(), node, interface_).UpdateNetworkInterfaceRequestContent(updateNetworkInterfaceRequestContent).Execute()
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

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth)

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
    openapiclient "./openapi"
)

func main() {
    node := "node_example" // string | 
    disk := "disk_example" // string | Disk to wipe. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.WipeDisk(context.Background(), node).Disk(disk).Execute()
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

[smithy.api.httpApiKeyAuth](../README.md#smithy.api.httpApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

