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
    writeLayerID := "writeLayerID_example" // string | 
    searchLayerIDs := []string{"Inner_example"} // []string | 
    version := "version_example" // string | 
    ansibleInventoryScanDTO := *openapiclient.NewAnsibleInventoryScanDTO(map[string]interface{}{"key": interface{}(123)}, map[string]interface{}{"key": interface{}(123)}, map[string]interface{}{"key": interface{}(123)}, map[string]interface{}{"key": interface{}(123)}) // AnsibleInventoryScanDTO | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AnsibleInventoryScanIngestApi.IngestAnsibleInventoryScan(context.Background(), version).WriteLayerID(writeLayerID).SearchLayerIDs(searchLayerIDs).AnsibleInventoryScanDTO(ansibleInventoryScanDTO).Execute()
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
 **writeLayerID** | **string** |  | 
 **searchLayerIDs** | **[]string** |  | 

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

