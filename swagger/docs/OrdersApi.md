# \OrdersApi

All URIs are relative to *https://api.pro.coinbase.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FillsGet**](OrdersApi.md#FillsGet) | **Get** /fills | list fills
[**OrdersDelete**](OrdersApi.md#OrdersDelete) | **Delete** /orders | cancel all orders
[**OrdersGet**](OrdersApi.md#OrdersGet) | **Get** /orders | list orders
[**OrdersOrderIdDelete**](OrdersApi.md#OrdersOrderIdDelete) | **Delete** /orders/{order-id} | cancel an order
[**OrdersPost**](OrdersApi.md#OrdersPost) | **Post** /orders | place a new order


# **FillsGet**
> []Fill FillsGet(ctx, cBACCESSKEY, cBACCESSSIGN, cBACCESSTIMESTAMP, cBACCESSPASSPHRASE, optional)
list fills

Get a list of recent fills of the API key's profile.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cBACCESSKEY** | **string**| The api key as a string. | 
  **cBACCESSSIGN** | **string**| The base64-encoded signature (see Signing a Message). | 
  **cBACCESSTIMESTAMP** | **string**| A timestamp for your request. | 
  **cBACCESSPASSPHRASE** | **string**| The passphrase you specified when creating the API key. | 
 **optional** | ***OrdersApiFillsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrdersApiFillsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **orderId** | **optional.String**|  | [default to all]
 **productId** | **optional.String**|  | [default to all]

### Return type

[**[]Fill**](Fill.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrdersDelete**
> OrdersDelete(ctx, cBACCESSKEY, cBACCESSSIGN, cBACCESSTIMESTAMP, cBACCESSPASSPHRASE, order, optional)
cancel all orders

With best effort, cancel all open orders from the profile that the API key belongs to. The response is a list of ids of the canceled orders.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cBACCESSKEY** | **string**| The api key as a string. | 
  **cBACCESSSIGN** | **string**| The base64-encoded signature (see Signing a Message). | 
  **cBACCESSTIMESTAMP** | **string**| A timestamp for your request. | 
  **cBACCESSPASSPHRASE** | **string**| The passphrase you specified when creating the API key. | 
  **order** | **[]string**|  | 
 **optional** | ***OrdersApiOrdersDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrdersApiOrdersDeleteOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **productId** | **optional.String**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrdersGet**
> []Order OrdersGet(ctx, cBACCESSKEY, cBACCESSSIGN, cBACCESSTIMESTAMP, cBACCESSPASSPHRASE, optional)
list orders

List your current open orders from the profile that the API key belongs to. Only open or un-settled orders are returned. As soon as an order is no longer open and settled, it will no longer appear in the default request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cBACCESSKEY** | **string**| The api key as a string. | 
  **cBACCESSSIGN** | **string**| The base64-encoded signature (see Signing a Message). | 
  **cBACCESSTIMESTAMP** | **string**| A timestamp for your request. | 
  **cBACCESSPASSPHRASE** | **string**| The passphrase you specified when creating the API key. | 
 **optional** | ***OrdersApiOrdersGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrdersApiOrdersGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **productId** | **optional.String**|  | 

### Return type

[**[]Order**](Order.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrdersOrderIdDelete**
> []Order OrdersOrderIdDelete(ctx, cBACCESSKEY, cBACCESSSIGN, cBACCESSTIMESTAMP, cBACCESSPASSPHRASE, orderId, optional)
cancel an order

Cancel a previously placed order. Order must belong to the profile that the API key belongs to. If the order had no matches during its lifetime its record may be purged. This means the order details will not be available with GET /orders/<id>.\" 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cBACCESSKEY** | **string**| The api key as a string. | 
  **cBACCESSSIGN** | **string**| The base64-encoded signature (see Signing a Message). | 
  **cBACCESSTIMESTAMP** | **string**| A timestamp for your request. | 
  **cBACCESSPASSPHRASE** | **string**| The passphrase you specified when creating the API key. | 
  **orderId** | **string**|  | 
 **optional** | ***OrdersApiOrdersOrderIdDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrdersApiOrdersOrderIdDeleteOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **productId** | **optional.String**|  | 

### Return type

[**[]Order**](Order.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrdersPost**
> []Order OrdersPost(ctx, cBACCESSKEY, cBACCESSSIGN, cBACCESSTIMESTAMP, cBACCESSPASSPHRASE, order)
place a new order

You can place two types of orders: limit and market. Orders can only be placed if your account has sufficient funds. Each profile can have a maximum of 500 open orders on a product. Once reached, the profile will not be able to place any new orders until the total number of open orders is below 500. Once an order is placed, your account funds will be put on hold for the duration of the order. How much and which funds are put on hold depends on the order type and parameters specified. See the Holds details below.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cBACCESSKEY** | **string**| The api key as a string. | 
  **cBACCESSSIGN** | **string**| The base64-encoded signature (see Signing a Message). | 
  **cBACCESSTIMESTAMP** | **string**| A timestamp for your request. | 
  **cBACCESSPASSPHRASE** | **string**| The passphrase you specified when creating the API key. | 
  **order** | [**OrderRequest**](OrderRequest.md)|  | 

### Return type

[**[]Order**](Order.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

