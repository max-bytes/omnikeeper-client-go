# \CytoscapeApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CytoscapeTraitCentric**](CytoscapeApi.md#CytoscapeTraitCentric) | **Get** /api/v{version}/Cytoscape/traitCentric | 



## CytoscapeTraitCentric

> CytoscapeTraitCentric(ctx, version).LayerIDs(layerIDs).TraitIDs(traitIDs).TraitIDsRegex(traitIDsRegex).Execute()



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
    r, err := apiClient.CytoscapeApi.CytoscapeTraitCentric(context.Background(), version).LayerIDs(layerIDs).TraitIDs(traitIDs).TraitIDsRegex(traitIDsRegex).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CytoscapeApi.CytoscapeTraitCentric``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCytoscapeTraitCentricRequest struct via the builder pattern


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

