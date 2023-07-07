# \OKPluginGenericJSONIngestAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ManageContextGetAllContexts**](OKPluginGenericJSONIngestAPI.md#ManageContextGetAllContexts) | **Get** /api/v{version}/ingest/genericJSON/manage/context | 
[**ManageContextGetContext**](OKPluginGenericJSONIngestAPI.md#ManageContextGetContext) | **Get** /api/v{version}/ingest/genericJSON/manage/context/{id} | 
[**ManageContextRemoveContext**](OKPluginGenericJSONIngestAPI.md#ManageContextRemoveContext) | **Delete** /api/v{version}/ingest/genericJSON/manage/context/{id} | 
[**ManageContextUpsertContext**](OKPluginGenericJSONIngestAPI.md#ManageContextUpsertContext) | **Post** /api/v{version}/ingest/genericJSON/manage/context | 
[**PassiveDataIngest**](OKPluginGenericJSONIngestAPI.md#PassiveDataIngest) | **Post** /api/v{version}/ingest/genericJSON/data | 



## ManageContextGetAllContexts

> []map[string]interface{} ManageContextGetAllContexts(ctx, version).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/max-bytes/omnikeeper-client-go"
)

func main() {
    version := "version_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OKPluginGenericJSONIngestAPI.ManageContextGetAllContexts(context.Background(), version).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OKPluginGenericJSONIngestAPI.ManageContextGetAllContexts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ManageContextGetAllContexts`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `OKPluginGenericJSONIngestAPI.ManageContextGetAllContexts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**version** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiManageContextGetAllContextsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**[]map[string]interface{}**

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml, application/json;odata.metadata=minimal;odata.streaming=true, application/json;odata.metadata=minimal;odata.streaming=false, application/json;odata.metadata=minimal, application/json;odata.metadata=full;odata.streaming=true, application/json;odata.metadata=full;odata.streaming=false, application/json;odata.metadata=full, application/json;odata.metadata=none;odata.streaming=true, application/json;odata.metadata=none;odata.streaming=false, application/json;odata.metadata=none, application/json;odata.streaming=true, application/json;odata.streaming=false, text/plain, application/octet-stream, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ManageContextGetContext

> map[string]interface{} ManageContextGetContext(ctx, id, version).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/max-bytes/omnikeeper-client-go"
)

func main() {
    id := "id_example" // string | 
    version := "version_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OKPluginGenericJSONIngestAPI.ManageContextGetContext(context.Background(), id, version).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OKPluginGenericJSONIngestAPI.ManageContextGetContext``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ManageContextGetContext`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `OKPluginGenericJSONIngestAPI.ManageContextGetContext`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**version** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiManageContextGetContextRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]interface{}**

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml, application/json;odata.metadata=minimal;odata.streaming=true, application/json;odata.metadata=minimal;odata.streaming=false, application/json;odata.metadata=minimal, application/json;odata.metadata=full;odata.streaming=true, application/json;odata.metadata=full;odata.streaming=false, application/json;odata.metadata=full, application/json;odata.metadata=none;odata.streaming=true, application/json;odata.metadata=none;odata.streaming=false, application/json;odata.metadata=none, application/json;odata.streaming=true, application/json;odata.streaming=false, text/plain, application/octet-stream, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ManageContextRemoveContext

> bool ManageContextRemoveContext(ctx, id, version).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/max-bytes/omnikeeper-client-go"
)

func main() {
    id := "id_example" // string | 
    version := "version_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OKPluginGenericJSONIngestAPI.ManageContextRemoveContext(context.Background(), id, version).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OKPluginGenericJSONIngestAPI.ManageContextRemoveContext``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ManageContextRemoveContext`: bool
    fmt.Fprintf(os.Stdout, "Response from `OKPluginGenericJSONIngestAPI.ManageContextRemoveContext`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**version** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiManageContextRemoveContextRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**bool**

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml, application/json;odata.metadata=minimal;odata.streaming=true, application/json;odata.metadata=minimal;odata.streaming=false, application/json;odata.metadata=minimal, application/json;odata.metadata=full;odata.streaming=true, application/json;odata.metadata=full;odata.streaming=false, application/json;odata.metadata=full, application/json;odata.metadata=none;odata.streaming=true, application/json;odata.metadata=none;odata.streaming=false, application/json;odata.metadata=none, application/json;odata.streaming=true, application/json;odata.streaming=false, text/plain, application/octet-stream, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ManageContextUpsertContext

> map[string]interface{} ManageContextUpsertContext(ctx, version).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/max-bytes/omnikeeper-client-go"
)

func main() {
    version := "version_example" // string | 
    body := map[string]interface{}{ ... } // map[string]interface{} | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OKPluginGenericJSONIngestAPI.ManageContextUpsertContext(context.Background(), version).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OKPluginGenericJSONIngestAPI.ManageContextUpsertContext``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ManageContextUpsertContext`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `OKPluginGenericJSONIngestAPI.ManageContextUpsertContext`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**version** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiManageContextUpsertContextRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** |  | 

### Return type

**map[string]interface{}**

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, application/json;odata.metadata=minimal;odata.streaming=true, application/json;odata.metadata=minimal;odata.streaming=false, application/json;odata.metadata=minimal, application/json;odata.metadata=full;odata.streaming=true, application/json;odata.metadata=full;odata.streaming=false, application/json;odata.metadata=full, application/json;odata.metadata=none;odata.streaming=true, application/json;odata.metadata=none;odata.streaming=false, application/json;odata.metadata=none, application/json;odata.streaming=true, application/json;odata.streaming=false, application/xml, text/plain, text/json, application/*+json
- **Accept**: application/json, application/xml, application/json;odata.metadata=minimal;odata.streaming=true, application/json;odata.metadata=minimal;odata.streaming=false, application/json;odata.metadata=minimal, application/json;odata.metadata=full;odata.streaming=true, application/json;odata.metadata=full;odata.streaming=false, application/json;odata.metadata=full, application/json;odata.metadata=none;odata.streaming=true, application/json;odata.metadata=none;odata.streaming=false, application/json;odata.metadata=none, application/json;odata.streaming=true, application/json;odata.streaming=false, text/plain, application/octet-stream, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PassiveDataIngest

> PassiveDataIngest(ctx, version).ReadLayerIDs(readLayerIDs).WriteLayerID(writeLayerID).GenericInboundData(genericInboundData).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/max-bytes/omnikeeper-client-go"
)

func main() {
    readLayerIDs := []string{"Inner_example"} // []string | 
    writeLayerID := "writeLayerID_example" // string | 
    version := "version_example" // string | 
    genericInboundData := *openapiclient.NewGenericInboundData() // GenericInboundData | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.OKPluginGenericJSONIngestAPI.PassiveDataIngest(context.Background(), version).ReadLayerIDs(readLayerIDs).WriteLayerID(writeLayerID).GenericInboundData(genericInboundData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OKPluginGenericJSONIngestAPI.PassiveDataIngest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**version** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPassiveDataIngestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **readLayerIDs** | **[]string** |  | 
 **writeLayerID** | **string** |  | 

 **genericInboundData** | [**GenericInboundData**](GenericInboundData.md) |  | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, application/json;odata.metadata=minimal;odata.streaming=true, application/json;odata.metadata=minimal;odata.streaming=false, application/json;odata.metadata=minimal, application/json;odata.metadata=full;odata.streaming=true, application/json;odata.metadata=full;odata.streaming=false, application/json;odata.metadata=full, application/json;odata.metadata=none;odata.streaming=true, application/json;odata.metadata=none;odata.streaming=false, application/json;odata.metadata=none, application/json;odata.streaming=true, application/json;odata.streaming=false, application/xml, text/plain, text/json, application/*+json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

