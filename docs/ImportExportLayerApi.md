# \ImportExportLayerApi

All URIs are relative to *https://localhost:44378*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ImportExportLayerExportLayer**](ImportExportLayerApi.md#ImportExportLayerExportLayer) | **Get** /api/v{version}/ImportExportLayer/exportLayer | 



## ImportExportLayerExportLayer

> ImportExportLayerExportLayer(ctx, version).LayerID(layerID).Ciids(ciids).Execute()



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
    layerID := "layerID_example" // string | 
    version := "version_example" // string | 
    ciids := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImportExportLayerApi.ImportExportLayerExportLayer(context.Background(), version).LayerID(layerID).Ciids(ciids).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportExportLayerApi.ImportExportLayerExportLayer``: %v\n", err)
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

Other parameters are passed through a pointer to a apiImportExportLayerExportLayerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **layerID** | **string** |  | 

 **ciids** | **[]string** |  | 

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

