# \CalculationsAPI

All URIs are relative to */api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CalculatePost**](CalculationsAPI.md#CalculatePost) | **Post** /calculate | Calculate a simple expression
[**GetHistoryGet**](CalculationsAPI.md#GetHistoryGet) | **Get** /get_history | Get calculation get_history for a user



## CalculatePost

> CalculateResponse CalculatePost(ctx).Authorization(authorization).Calculation(calculation).Execute()

Calculate a simple expression



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "authorization_example" // string | Authorization token (UUID)
	calculation := *openapiclient.NewCalculateRequest() // CalculateRequest | Calculation request body

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CalculationsAPI.CalculatePost(context.Background()).Authorization(authorization).Calculation(calculation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CalculationsAPI.CalculatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CalculatePost`: CalculateResponse
	fmt.Fprintf(os.Stdout, "Response from `CalculationsAPI.CalculatePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCalculatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Authorization token (UUID) | 
 **calculation** | [**CalculateRequest**](CalculateRequest.md) | Calculation request body | 

### Return type

[**CalculateResponse**](CalculateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHistoryGet

> GetHistoryResponse GetHistoryGet(ctx).Authorization(authorization).Execute()

Get calculation get_history for a user



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "authorization_example" // string | Authorization token (UUID)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CalculationsAPI.GetHistoryGet(context.Background()).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CalculationsAPI.GetHistoryGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHistoryGet`: GetHistoryResponse
	fmt.Fprintf(os.Stdout, "Response from `CalculationsAPI.GetHistoryGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetHistoryGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Authorization token (UUID) | 

### Return type

[**GetHistoryResponse**](GetHistoryResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

