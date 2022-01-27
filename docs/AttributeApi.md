# \AttributeApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BulkReplaceAttributesInLayer**](AttributeApi.md#BulkReplaceAttributesInLayer) | **Post** /api/v{version}/Attribute/bulkReplaceAttributesInLayer | bulk replace all attributes in specified layer
[**FindMergedAttributesByName**](AttributeApi.md#FindMergedAttributesByName) | **Get** /api/v{version}/Attribute/findMergedAttributesByName | 
[**GetMergedAttribute**](AttributeApi.md#GetMergedAttribute) | **Get** /api/v{version}/Attribute/getMergedAttribute | 
[**GetMergedAttributes**](AttributeApi.md#GetMergedAttributes) | **Get** /api/v{version}/Attribute/getMergedAttributes | 
[**GetMergedAttributesWithName**](AttributeApi.md#GetMergedAttributesWithName) | **Get** /api/v{version}/Attribute/getMergedAttributesWithName | 



## BulkReplaceAttributesInLayer

> BulkReplaceAttributesInLayer(ctx, version).BulkCIAttributeLayerScopeDTO(bulkCIAttributeLayerScopeDTO).Execute()

bulk replace all attributes in specified layer

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
    version := "version_example" // string | 
    bulkCIAttributeLayerScopeDTO := *openapiclient.NewBulkCIAttributeLayerScopeDTO("NamePrefix_example", "LayerID_example", []openapiclient.FragmentDTO{*openapiclient.NewFragmentDTO("Name_example", *openapiclient.NewAttributeValueDTO(openapiclient.AttributeValueType("Text"), false, []string{"Values_example"}), "Ciid_example")}) // BulkCIAttributeLayerScopeDTO | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AttributeApi.BulkReplaceAttributesInLayer(context.Background(), version).BulkCIAttributeLayerScopeDTO(bulkCIAttributeLayerScopeDTO).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AttributeApi.BulkReplaceAttributesInLayer``: %v\n", err)
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

Other parameters are passed through a pointer to a apiBulkReplaceAttributesInLayerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bulkCIAttributeLayerScopeDTO** | [**BulkCIAttributeLayerScopeDTO**](BulkCIAttributeLayerScopeDTO.md) |  | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json;odata.metadata=minimal;odata.streaming=true, application/json;odata.metadata=minimal;odata.streaming=false, application/json;odata.metadata=minimal, application/json;odata.metadata=full;odata.streaming=true, application/json;odata.metadata=full;odata.streaming=false, application/json;odata.metadata=full, application/json;odata.metadata=none;odata.streaming=true, application/json;odata.metadata=none;odata.streaming=false, application/json;odata.metadata=none, application/json;odata.streaming=true, application/json;odata.streaming=false, application/json, application/xml, application/odata, application/json-patch+json, text/json, application/_*+json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindMergedAttributesByName

> []CIAttributeDTO FindMergedAttributesByName(ctx, version).Regex(regex).LayerIDs(layerIDs).Ciids(ciids).AtTime(atTime).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    regex := "regex_example" // string | 
    layerIDs := []string{"Inner_example"} // []string | 
    version := "version_example" // string | 
    ciids := []string{"Inner_example"} // []string |  (optional)
    atTime := time.Now() // time.Time |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AttributeApi.FindMergedAttributesByName(context.Background(), version).Regex(regex).LayerIDs(layerIDs).Ciids(ciids).AtTime(atTime).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AttributeApi.FindMergedAttributesByName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindMergedAttributesByName`: []CIAttributeDTO
    fmt.Fprintf(os.Stdout, "Response from `AttributeApi.FindMergedAttributesByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**version** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindMergedAttributesByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **regex** | **string** |  | 
 **layerIDs** | **[]string** |  | 

 **ciids** | **[]string** |  | 
 **atTime** | **time.Time** |  | 

### Return type

[**[]CIAttributeDTO**](CIAttributeDTO.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;odata.metadata=minimal;odata.streaming=true, application/json;odata.metadata=minimal;odata.streaming=false, application/json;odata.metadata=minimal, application/json;odata.metadata=full;odata.streaming=true, application/json;odata.metadata=full;odata.streaming=false, application/json;odata.metadata=full, application/json;odata.metadata=none;odata.streaming=true, application/json;odata.metadata=none;odata.streaming=false, application/json;odata.metadata=none, application/json;odata.streaming=true, application/json;odata.streaming=false, application/json, application/xml, application/odata, text/plain, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMergedAttribute

> CIAttributeDTO GetMergedAttribute(ctx, version).Ciid(ciid).Name(name).LayerIDs(layerIDs).AtTime(atTime).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    ciid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    name := "name_example" // string | 
    layerIDs := []string{"Inner_example"} // []string | 
    version := "version_example" // string | 
    atTime := time.Now() // time.Time |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AttributeApi.GetMergedAttribute(context.Background(), version).Ciid(ciid).Name(name).LayerIDs(layerIDs).AtTime(atTime).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AttributeApi.GetMergedAttribute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMergedAttribute`: CIAttributeDTO
    fmt.Fprintf(os.Stdout, "Response from `AttributeApi.GetMergedAttribute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**version** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMergedAttributeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ciid** | **string** |  | 
 **name** | **string** |  | 
 **layerIDs** | **[]string** |  | 

 **atTime** | **time.Time** |  | 

### Return type

[**CIAttributeDTO**](CIAttributeDTO.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;odata.metadata=minimal;odata.streaming=true, application/json;odata.metadata=minimal;odata.streaming=false, application/json;odata.metadata=minimal, application/json;odata.metadata=full;odata.streaming=true, application/json;odata.metadata=full;odata.streaming=false, application/json;odata.metadata=full, application/json;odata.metadata=none;odata.streaming=true, application/json;odata.metadata=none;odata.streaming=false, application/json;odata.metadata=none, application/json;odata.streaming=true, application/json;odata.streaming=false, application/json, application/xml, application/odata, text/plain, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMergedAttributes

> []CIAttributeDTO GetMergedAttributes(ctx, version).Ciids(ciids).LayerIDs(layerIDs).AtTime(atTime).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    ciids := []string{"Inner_example"} // []string | 
    layerIDs := []string{"Inner_example"} // []string | 
    version := "version_example" // string | 
    atTime := time.Now() // time.Time |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AttributeApi.GetMergedAttributes(context.Background(), version).Ciids(ciids).LayerIDs(layerIDs).AtTime(atTime).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AttributeApi.GetMergedAttributes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMergedAttributes`: []CIAttributeDTO
    fmt.Fprintf(os.Stdout, "Response from `AttributeApi.GetMergedAttributes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**version** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMergedAttributesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ciids** | **[]string** |  | 
 **layerIDs** | **[]string** |  | 

 **atTime** | **time.Time** |  | 

### Return type

[**[]CIAttributeDTO**](CIAttributeDTO.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;odata.metadata=minimal;odata.streaming=true, application/json;odata.metadata=minimal;odata.streaming=false, application/json;odata.metadata=minimal, application/json;odata.metadata=full;odata.streaming=true, application/json;odata.metadata=full;odata.streaming=false, application/json;odata.metadata=full, application/json;odata.metadata=none;odata.streaming=true, application/json;odata.metadata=none;odata.streaming=false, application/json;odata.metadata=none, application/json;odata.streaming=true, application/json;odata.streaming=false, application/json, application/xml, application/odata, text/plain, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMergedAttributesWithName

> []CIAttributeDTO GetMergedAttributesWithName(ctx, version).Name(name).LayerIDs(layerIDs).AtTime(atTime).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | 
    layerIDs := []string{"Inner_example"} // []string | 
    version := "version_example" // string | 
    atTime := time.Now() // time.Time |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AttributeApi.GetMergedAttributesWithName(context.Background(), version).Name(name).LayerIDs(layerIDs).AtTime(atTime).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AttributeApi.GetMergedAttributesWithName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMergedAttributesWithName`: []CIAttributeDTO
    fmt.Fprintf(os.Stdout, "Response from `AttributeApi.GetMergedAttributesWithName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**version** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMergedAttributesWithNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** |  | 
 **layerIDs** | **[]string** |  | 

 **atTime** | **time.Time** |  | 

### Return type

[**[]CIAttributeDTO**](CIAttributeDTO.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;odata.metadata=minimal;odata.streaming=true, application/json;odata.metadata=minimal;odata.streaming=false, application/json;odata.metadata=minimal, application/json;odata.metadata=full;odata.streaming=true, application/json;odata.metadata=full;odata.streaming=false, application/json;odata.metadata=full, application/json;odata.metadata=none;odata.streaming=true, application/json;odata.metadata=none;odata.streaming=false, application/json;odata.metadata=none, application/json;odata.streaming=true, application/json;odata.streaming=false, application/json, application/xml, application/odata, text/plain, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

