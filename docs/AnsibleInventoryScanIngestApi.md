# \AnsibleInventoryScanIngestApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IngestAnsibleInventoryScan**](AnsibleInventoryScanIngestApi.md#IngestAnsibleInventoryScan) | **Post** /api/v{version}/Ingest/AnsibleInventoryScan | 



## IngestAnsibleInventoryScan

> IngestAnsibleInventoryScan(ctx, version).WriteLayerID(writeLayerID).SearchLayerIDs(searchLayerIDs).AnsibleInventoryScanDTO(ansibleInventoryScanDTO).Execute()



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
    writeLayerID := int64(789) // int64 | 
    searchLayerIDs := []int64{int64(123)} // []int64 | 
    version := "version_example" // string | 
    ansibleInventoryScanDTO := *openapiclient.NewAnsibleInventoryScanDTO(map[string]map[string]interface{}{"key": map[string]interface{}(123)}, map[string]map[string]interface{}{"key": map[string]interface{}(123)}, map[string]map[string]interface{}{"key": map[string]interface{}(123)}, map[string]map[string]interface{}{"key": map[string]interface{}(123)}) // AnsibleInventoryScanDTO | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AnsibleInventoryScanIngestApi.IngestAnsibleInventoryScan(context.Background(), version).WriteLayerID(writeLayerID).SearchLayerIDs(searchLayerIDs).AnsibleInventoryScanDTO(ansibleInventoryScanDTO).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnsibleInventoryScanIngestApi.IngestAnsibleInventoryScan``: %v\n", err)
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

Other parameters are passed through a pointer to a apiIngestAnsibleInventoryScanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **writeLayerID** | **int64** |  | 
 **searchLayerIDs** | **[]int64** |  | 

 **ansibleInventoryScanDTO** | [**AnsibleInventoryScanDTO**](AnsibleInventoryScanDTO.md) |  | 

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

