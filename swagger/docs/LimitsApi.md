# \LimitsApi

All URIs are relative to *https://api.pro.coinbase.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersSelfExchangeLimitsGet**](LimitsApi.md#UsersSelfExchangeLimitsGet) | **Get** /users/self/exchange-limits | list fills


# **UsersSelfExchangeLimitsGet**
> []LimitResponse UsersSelfExchangeLimitsGet(ctx, cBACCESSKEY, cBACCESSSIGN, cBACCESSTIMESTAMP, cBACCESSPASSPHRASE, optional)
list fills

This request will return information on your payment method transfer limits, as well as buy/sell limits per currency.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cBACCESSKEY** | **string**| The api key as a string. | 
  **cBACCESSSIGN** | **string**| The base64-encoded signature (see Signing a Message). | 
  **cBACCESSTIMESTAMP** | **string**| A timestamp for your request. | 
  **cBACCESSPASSPHRASE** | **string**| The passphrase you specified when creating the API key. | 
 **optional** | ***LimitsApiUsersSelfExchangeLimitsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LimitsApiUsersSelfExchangeLimitsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **orderId** | **optional.String**|  | [default to all]
 **productId** | **optional.String**|  | [default to all]

### Return type

[**[]LimitResponse**](LimitResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

