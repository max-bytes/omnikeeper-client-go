# \TraitApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEffectiveTraitsForTraitName**](TraitApi.md#GetEffectiveTraitsForTraitName) | **Get** /api/v{version}/Trait/getEffectiveTraitsForTraitName | 



## GetEffectiveTraitsForTraitName

> map[string]EffectiveTraitDTO GetEffectiveTraitsForTraitName(ctx, version).LayerIDs(layerIDs).TraitName(traitName).AtTime(atTime).Execute()



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
    layerIDs := []int64{int64(123)} // []int64 | 
    traitName := "traitName_example" // string | 
    version := "version_example" // string | 
    atTime := time.Now() // time.Time |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TraitApi.GetEffectiveTraitsForTraitName(context.Background(), version).LayerIDs(layerIDs).TraitName(traitName).AtTime(atTime).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TraitApi.GetEffectiveTraitsForTraitName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEffectiveTraitsForTraitName`: map[string]EffectiveTraitDTO
    fmt.Fprintf(os.Stdout, "Response from `TraitApi.GetEffectiveTraitsForTraitName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**version** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEffectiveTraitsForTraitNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **layerIDs** | **[]int64** |  | 
 **traitName** | **string** |  | 

 **atTime** | **time.Time** |  | 

### Return type

[**map[string]EffectiveTraitDTO**](EffectiveTraitDTO.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;odata.metadata=minimal;odata.streaming=true, application/json;odata.metadata=minimal;odata.streaming=false, application/json;odata.metadata=minimal, application/json;odata.metadata=full;odata.streaming=true, application/json;odata.metadata=full;odata.streaming=false, application/json;odata.metadata=full, application/json;odata.metadata=none;odata.streaming=true, application/json;odata.metadata=none;odata.streaming=false, application/json;odata.metadata=none, application/json;odata.streaming=true, application/json;odata.streaming=false, application/json, application/xml, application/odata, text/plain, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

