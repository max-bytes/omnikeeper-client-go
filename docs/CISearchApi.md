# \CISearchApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SearchCIsByTraits**](CISearchApi.md#SearchCIsByTraits) | **Get** /api/v{version}/CISearch/searchCIsByTraits | 



## SearchCIsByTraits

> []CIDTO SearchCIsByTraits(ctx, version).LayerIDs(layerIDs).WithTraits(withTraits).WithoutTraits(withoutTraits).AtTime(atTime).Execute()



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
    withTraits := []string{"Inner_example"} // []string | 
    withoutTraits := []string{"Inner_example"} // []string | 
    version := "version_example" // string | 
    atTime := time.Now() // time.Time |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CISearchApi.SearchCIsByTraits(context.Background(), version).LayerIDs(layerIDs).WithTraits(withTraits).WithoutTraits(withoutTraits).AtTime(atTime).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CISearchApi.SearchCIsByTraits``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchCIsByTraits`: []CIDTO
    fmt.Fprintf(os.Stdout, "Response from `CISearchApi.SearchCIsByTraits`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**version** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchCIsByTraitsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **layerIDs** | **[]string** |  | 
 **withTraits** | **[]string** |  | 
 **withoutTraits** | **[]string** |  | 

 **atTime** | **time.Time** |  | 

### Return type

[**[]CIDTO**](CIDTO.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;odata.metadata=minimal;odata.streaming=true, application/json;odata.metadata=minimal;odata.streaming=false, application/json;odata.metadata=minimal, application/json;odata.metadata=full;odata.streaming=true, application/json;odata.metadata=full;odata.streaming=false, application/json;odata.metadata=full, application/json;odata.metadata=none;odata.streaming=true, application/json;odata.metadata=none;odata.streaming=false, application/json;odata.metadata=none, application/json;odata.streaming=true, application/json;odata.streaming=false, application/json, application/xml, application/odata, text/plain, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

