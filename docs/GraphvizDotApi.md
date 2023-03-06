# \GraphvizDotApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GraphvizDotLayerCentric**](GraphvizDotApi.md#GraphvizDotLayerCentric) | **Get** /api/v{version}/GraphvizDot/layerCentric | 
[**GraphvizDotTraitCentric**](GraphvizDotApi.md#GraphvizDotTraitCentric) | **Get** /api/v{version}/GraphvizDot/traitCentric | 



## GraphvizDotLayerCentric

> GraphvizDotLayerCentric(ctx, version).LayerIDs(layerIDs).From(from).To(to).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "github.com/max-bytes/omnikeeper-client-go"
)

func main() {
    layerIDs := []string{"Inner_example"} // []string | 
    from := time.Now() // time.Time | 
    to := time.Now() // time.Time | 
    version := "version_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.GraphvizDotApi.GraphvizDotLayerCentric(context.Background(), version).LayerIDs(layerIDs).From(from).To(to).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GraphvizDotApi.GraphvizDotLayerCentric``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGraphvizDotLayerCentricRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **layerIDs** | **[]string** |  | 
 **from** | **time.Time** |  | 
 **to** | **time.Time** |  | 


### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GraphvizDotTraitCentric

> GraphvizDotTraitCentric(ctx, version).LayerIDs(layerIDs).TraitIDs(traitIDs).TraitIDsRegex(traitIDsRegex).Execute()



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
    layerIDs := []string{"Inner_example"} // []string | 
    version := "version_example" // string | 
    traitIDs := []string{"Inner_example"} // []string |  (optional)
    traitIDsRegex := "traitIDsRegex_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.GraphvizDotApi.GraphvizDotTraitCentric(context.Background(), version).LayerIDs(layerIDs).TraitIDs(traitIDs).TraitIDsRegex(traitIDsRegex).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GraphvizDotApi.GraphvizDotTraitCentric``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGraphvizDotTraitCentricRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **layerIDs** | **[]string** |  | 

 **traitIDs** | **[]string** |  | 
 **traitIDsRegex** | **string** |  | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

