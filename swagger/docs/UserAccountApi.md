# \UserAccountApi

All URIs are relative to *https://api.pro.coinbase.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersSelfTrailingVolumeGet**](UserAccountApi.md#UsersSelfTrailingVolumeGet) | **Get** /users/self/trailing-volume | Trailing Volume


# **UsersSelfTrailingVolumeGet**
> []VolumeItem UsersSelfTrailingVolumeGet(ctx, cBACCESSKEY, cBACCESSSIGN, cBACCESSTIMESTAMP, cBACCESSPASSPHRASE)
Trailing Volume

This endpoint requires either the \"view\" or \"trade\" permission. This request will return your 30-day trailing volume for all products of the API key's profile. This is a cached value that's calculated every day at midnight UTC.\" 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cBACCESSKEY** | **string**| The api key as a string. | 
  **cBACCESSSIGN** | **string**| The base64-encoded signature (see Signing a Message). | 
  **cBACCESSTIMESTAMP** | **string**| A timestamp for your request. | 
  **cBACCESSPASSPHRASE** | **string**| The passphrase you specified when creating the API key. | 

### Return type

[**[]VolumeItem**](VolumeItem.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

