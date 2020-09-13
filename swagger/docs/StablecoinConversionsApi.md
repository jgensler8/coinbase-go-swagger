# \StablecoinConversionsApi

All URIs are relative to *https://api.pro.coinbase.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConversionsPost**](StablecoinConversionsApi.md#ConversionsPost) | **Post** /conversions | convert to/from stable coins?


# **ConversionsPost**
> []Order ConversionsPost(ctx, cBACCESSKEY, cBACCESSSIGN, cBACCESSTIMESTAMP, cBACCESSPASSPHRASE, order)
convert to/from stable coins?

Convert $10,000.00 to 10,000.00 USDC.

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

