# \RelationApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAllMergedRelations**](RelationApi.md#GetAllMergedRelations) | **Get** /api/v{version}/Relation/getAllMergedRelations | 
[**GetMergedRelation**](RelationApi.md#GetMergedRelation) | **Get** /api/v{version}/Relation/getMergedRelation | 
[**GetMergedRelationsFromOrToCI**](RelationApi.md#GetMergedRelationsFromOrToCI) | **Get** /api/v{version}/Relation/getMergedRelationsFromOrToCI | 
[**GetMergedRelationsOutgoingFromCI**](RelationApi.md#GetMergedRelationsOutgoingFromCI) | **Get** /api/v{version}/Relation/getMergedRelationsOutgoingFromCI | 
[**GetMergedRelationsWithPredicate**](RelationApi.md#GetMergedRelationsWithPredicate) | **Get** /api/v{version}/Relation/getMergedRelationsWithPredicate | 



## GetAllMergedRelations

> []RelationDTO GetAllMergedRelations(ctx, version).LayerIDs(layerIDs).AtTime(atTime).Execute()



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
    layerIDs := []string{"Inner_example"} // []string | 
    version := "version_example" // string | 
    atTime := time.Now() // time.Time |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RelationApi.GetAllMergedRelations(context.Background(), version).LayerIDs(layerIDs).AtTime(atTime).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RelationApi.GetAllMergedRelations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllMergedRelations`: []RelationDTO
    fmt.Fprintf(os.Stdout, "Response from `RelationApi.GetAllMergedRelations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**version** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllMergedRelationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **layerIDs** | **[]string** |  | 

 **atTime** | **time.Time** |  | 

### Return type

[**[]RelationDTO**](RelationDTO.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;odata.metadata=minimal;odata.streaming=true, application/json;odata.metadata=minimal;odata.streaming=false, application/json;odata.metadata=minimal, application/json;odata.metadata=full;odata.streaming=true, application/json;odata.metadata=full;odata.streaming=false, application/json;odata.metadata=full, application/json;odata.metadata=none;odata.streaming=true, application/json;odata.metadata=none;odata.streaming=false, application/json;odata.metadata=none, application/json;odata.streaming=true, application/json;odata.streaming=false, application/json, application/xml, application/odata, text/plain, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMergedRelation

> RelationDTO GetMergedRelation(ctx, version).FromCIID(fromCIID).ToCIID(toCIID).PredicateID(predicateID).LayerIDs(layerIDs).AtTime(atTime).Execute()



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
    fromCIID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    toCIID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    predicateID := "predicateID_example" // string | 
    layerIDs := []string{"Inner_example"} // []string | 
    version := "version_example" // string | 
    atTime := time.Now() // time.Time |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RelationApi.GetMergedRelation(context.Background(), version).FromCIID(fromCIID).ToCIID(toCIID).PredicateID(predicateID).LayerIDs(layerIDs).AtTime(atTime).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RelationApi.GetMergedRelation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMergedRelation`: RelationDTO
    fmt.Fprintf(os.Stdout, "Response from `RelationApi.GetMergedRelation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**version** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMergedRelationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fromCIID** | **string** |  | 
 **toCIID** | **string** |  | 
 **predicateID** | **string** |  | 
 **layerIDs** | **[]string** |  | 

 **atTime** | **time.Time** |  | 

### Return type

[**RelationDTO**](RelationDTO.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;odata.metadata=minimal;odata.streaming=true, application/json;odata.metadata=minimal;odata.streaming=false, application/json;odata.metadata=minimal, application/json;odata.metadata=full;odata.streaming=true, application/json;odata.metadata=full;odata.streaming=false, application/json;odata.metadata=full, application/json;odata.metadata=none;odata.streaming=true, application/json;odata.metadata=none;odata.streaming=false, application/json;odata.metadata=none, application/json;odata.streaming=true, application/json;odata.streaming=false, application/json, application/xml, application/odata, text/plain, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMergedRelationsFromOrToCI

> []RelationDTO GetMergedRelationsFromOrToCI(ctx, version).Ciid(ciid).LayerIDs(layerIDs).AtTime(atTime).Execute()



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
    layerIDs := []string{"Inner_example"} // []string | 
    version := "version_example" // string | 
    atTime := time.Now() // time.Time |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RelationApi.GetMergedRelationsFromOrToCI(context.Background(), version).Ciid(ciid).LayerIDs(layerIDs).AtTime(atTime).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RelationApi.GetMergedRelationsFromOrToCI``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMergedRelationsFromOrToCI`: []RelationDTO
    fmt.Fprintf(os.Stdout, "Response from `RelationApi.GetMergedRelationsFromOrToCI`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**version** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMergedRelationsFromOrToCIRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ciid** | **string** |  | 
 **layerIDs** | **[]string** |  | 

 **atTime** | **time.Time** |  | 

### Return type

[**[]RelationDTO**](RelationDTO.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;odata.metadata=minimal;odata.streaming=true, application/json;odata.metadata=minimal;odata.streaming=false, application/json;odata.metadata=minimal, application/json;odata.metadata=full;odata.streaming=true, application/json;odata.metadata=full;odata.streaming=false, application/json;odata.metadata=full, application/json;odata.metadata=none;odata.streaming=true, application/json;odata.metadata=none;odata.streaming=false, application/json;odata.metadata=none, application/json;odata.streaming=true, application/json;odata.streaming=false, application/json, application/xml, application/odata, text/plain, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMergedRelationsOutgoingFromCI

> []RelationDTO GetMergedRelationsOutgoingFromCI(ctx, version).FromCIID(fromCIID).LayerIDs(layerIDs).AtTime(atTime).Execute()



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
    fromCIID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    layerIDs := []string{"Inner_example"} // []string | 
    version := "version_example" // string | 
    atTime := time.Now() // time.Time |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RelationApi.GetMergedRelationsOutgoingFromCI(context.Background(), version).FromCIID(fromCIID).LayerIDs(layerIDs).AtTime(atTime).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RelationApi.GetMergedRelationsOutgoingFromCI``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMergedRelationsOutgoingFromCI`: []RelationDTO
    fmt.Fprintf(os.Stdout, "Response from `RelationApi.GetMergedRelationsOutgoingFromCI`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**version** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMergedRelationsOutgoingFromCIRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fromCIID** | **string** |  | 
 **layerIDs** | **[]string** |  | 

 **atTime** | **time.Time** |  | 

### Return type

[**[]RelationDTO**](RelationDTO.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;odata.metadata=minimal;odata.streaming=true, application/json;odata.metadata=minimal;odata.streaming=false, application/json;odata.metadata=minimal, application/json;odata.metadata=full;odata.streaming=true, application/json;odata.metadata=full;odata.streaming=false, application/json;odata.metadata=full, application/json;odata.metadata=none;odata.streaming=true, application/json;odata.metadata=none;odata.streaming=false, application/json;odata.metadata=none, application/json;odata.streaming=true, application/json;odata.streaming=false, application/json, application/xml, application/odata, text/plain, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMergedRelationsWithPredicate

> []RelationDTO GetMergedRelationsWithPredicate(ctx, version).PredicateID(predicateID).LayerIDs(layerIDs).AtTime(atTime).Execute()



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
    predicateID := "predicateID_example" // string | 
    layerIDs := []string{"Inner_example"} // []string | 
    version := "version_example" // string | 
    atTime := time.Now() // time.Time |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RelationApi.GetMergedRelationsWithPredicate(context.Background(), version).PredicateID(predicateID).LayerIDs(layerIDs).AtTime(atTime).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RelationApi.GetMergedRelationsWithPredicate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMergedRelationsWithPredicate`: []RelationDTO
    fmt.Fprintf(os.Stdout, "Response from `RelationApi.GetMergedRelationsWithPredicate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**version** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMergedRelationsWithPredicateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **predicateID** | **string** |  | 
 **layerIDs** | **[]string** |  | 

 **atTime** | **time.Time** |  | 

### Return type

[**[]RelationDTO**](RelationDTO.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;odata.metadata=minimal;odata.streaming=true, application/json;odata.metadata=minimal;odata.streaming=false, application/json;odata.metadata=minimal, application/json;odata.metadata=full;odata.streaming=true, application/json;odata.metadata=full;odata.streaming=false, application/json;odata.metadata=full, application/json;odata.metadata=none;odata.streaming=true, application/json;odata.metadata=none;odata.streaming=false, application/json;odata.metadata=none, application/json;odata.streaming=true, application/json;odata.streaming=false, application/json, application/xml, application/odata, text/plain, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

