# \DefaultApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddCorosyncNode**](DefaultApi.md#AddCorosyncNode) | **Post** /cluster/config/nodes/{node} | 
[**ApplyNetworkInterfaceConfiguration**](DefaultApi.md#ApplyNetworkInterfaceConfiguration) | **Put** /nodes/{node}/network | 
[**CreateClusterConfig**](DefaultApi.md#CreateClusterConfig) | **Post** /cluster/config | 
[**CreateNetworkInterface**](DefaultApi.md#CreateNetworkInterface) | **Post** /nodes/{node}/network | 
[**CreateTicket**](DefaultApi.md#CreateTicket) | **Post** /access/ticket | 
[**DeleteNetworkInterface**](DefaultApi.md#DeleteNetworkInterface) | **Delete** /nodes/{node}/network/{interface} | 
[**GetAccessControlList**](DefaultApi.md#GetAccessControlList) | **Get** /access/acl | 
[**GetClusterApiVersion**](DefaultApi.md#GetClusterApiVersion) | **Get** /cluster/config/apiversion | 
[**GetClusterJoinInformation**](DefaultApi.md#GetClusterJoinInformation) | **Get** /cluster/config/join | 
[**GetClusterTotemSettings**](DefaultApi.md#GetClusterTotemSettings) | **Get** /cluster/config/totem | 
[**GetNetworkInterface**](DefaultApi.md#GetNetworkInterface) | **Get** /nodes/{node}/network/{interface} | 
[**GetVersion**](DefaultApi.md#GetVersion) | **Get** /version | 
[**JoinCluster**](DefaultApi.md#JoinCluster) | **Post** /cluster/config/join | 
[**ListCorosyncNodes**](DefaultApi.md#ListCorosyncNodes) | **Get** /cluster/config/nodes | 
[**ListNetworkInterfaces**](DefaultApi.md#ListNetworkInterfaces) | **Get** /nodes/{node}/network | 
[**ListNodes**](DefaultApi.md#ListNodes) | **Get** /nodes | 
[**ListStorage**](DefaultApi.md#ListStorage) | **Get** /storage | 
[**ListVirtualMachines**](DefaultApi.md#ListVirtualMachines) | **Get** /nodes/{node}/qemu | 
[**RemoveCorosyncNode**](DefaultApi.md#RemoveCorosyncNode) | **Delete** /cluster/config/nodes/{node} | 
[**RevertNetworkInterfaceConfiguration**](DefaultApi.md#RevertNetworkInterfaceConfiguration) | **Delete** /nodes/{node}/network | 
[**UpdateAccessControlList**](DefaultApi.md#UpdateAccessControlList) | **Put** /access/acl | 
[**UpdateNetworkInterface**](DefaultApi.md#UpdateNetworkInterface) | **Put** /nodes/{node}/network/{interface} | 



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

